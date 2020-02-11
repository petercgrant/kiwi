import { Schema, Definition, Field } from "./schema";
import { ByteBuffer } from "./bb";
import { error, quote } from "./util";
import { nativeTypes } from "./parser";

function decodeCodeForField(field: Field, definitions: {[name: string]: Definition}): string | string[] {
    let code: string;

    switch (field.type) {
      case 'bool': {
        code = '!!bb.readByte()';
        break;
      }

      case 'byte': {
        code = 'bb.readByte()';  // only used if not array
        break;
      }

      case 'int': {
        code = 'bb.readVarInt()';
        break;
      }

      case 'uint': {
        code = 'bb.readVarUint()';
        break;
      }

      case 'float':
      case 'float32': {
        code = 'bb.readVarFloat()';
        break;
      }

      case 'string': {
        code = 'bb.readString()';
        break;
      }

      case 'map': {
        let keyCode = decodeCodeForField({ type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null }, definitions);
        if (!keyCode || nativeTypes.indexOf(field.mapKeyType as string) < 0) {
          error(
            'Invalid type ' +
              quote(field.mapKeyType as string) +
              ' for map key for ' +
              quote(field.name) +
              '.  Map keys must be native fields.',
            field.line,
            field.column
          );
        }
        let mapValueField: Field = { type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
        let valueCode = decodeCodeForField(mapValueField, definitions);
        if (!valueCode) {
          error(
            'Invalid type ' +
              quote(field.mapValueType as string) +
              ' for map value for ' +
              quote(field.name),
            field.line,
            field.column
          );
        }
        return [keyCode as string, valueCode as string];
      }

      default: {
        let type = definitions[field.type!];
        if (!type) {
          error('Invalid type ' + quote(field.type!) + ' for field ' + quote(field.name), field.line, field.column);
        } else if (type.kind === 'ENUM') {
          code = [
            '(function (t) {',
            '  var byte = bb.readVarUint();',
            '  if (undefined == t[' + quote(type.name) + '][byte]) { throw new Error("Attempted to parse invalid enum"); }',
            '  return t[' + quote(type.name) + '][byte]',
            '})(this)',
          ].join('\n');
        } else {
          code = 'this[' + quote('decode' + type.name) + '](bb)';
        }
      }
    }
    return code;
}

function compileDecode(definition: Definition, definitions: {[name: string]: Definition}): string {
  let lines: string[] = [];
  let indent = '  ';

  lines.push('function(bb) {');
  lines.push('  var result = {};');
  lines.push('  if (!(bb instanceof this.ByteBuffer)) {');
  lines.push('    bb = new this.ByteBuffer(bb);');
  lines.push('  }');
  lines.push('');

  if (definition.kind === 'MESSAGE') {
    lines.push('  while (true) {');
    lines.push('    switch (bb.readVarUint()) {');
    lines.push('    case 0:');
    lines.push('      return result;');
    lines.push('');
    indent = '      ';
  }

  for (let i = 0; i < definition.fields.length; i++) {
    let field = definition.fields[i];
    let code = decodeCodeForField(field, definitions);

    if (definition.kind === 'MESSAGE') {
      lines.push('    case ' + field.value + ':');
    }
    if (field.isArray) {
      code = code as string;
      if (field.isDeprecated) {
        if (field.type === 'byte') {
          lines.push(indent + 'bb.readByteArray();');
        } else {
          lines.push(indent + 'var length = bb.readVarUint();');
          lines.push(indent + 'while (length-- > 0) {' + code + '};');
        }
      } else {
        if (field.type === 'byte') {
          lines.push(indent + 'result[' + quote(field.name) + '] = bb.readByteArray();');
        } else if (field.type === 'float32') {
          lines.push(indent + 'var length = bb.readVarUint();');
          lines.push(indent + 'var values = result[' + quote(field.name) + '] = new Float32Array(length);');
          lines.push(indent + 'var c = 0;');
          lines.push(indent + 'while (length-- > 0) { values[c] = (' + code + '); c++ }');
        } else {
          lines.push(indent + 'var length = bb.readVarUint();');
          lines.push(indent + 'var values = result[' + quote(field.name) + '] = [];');
          lines.push(indent + 'while (length-- > 0) { values.push(' + code + '); }');
        }
      }
    } else if (field.isMap) {
        code = code as string[];
        if (field.isDeprecated) {
          lines.push(indent + 'var length = bb.readVarUint();');
          lines.push(indent + 'while (length-- > 0) { ' + code[0] + ';' + code[1] + '; }');
         } else {
          lines.push(indent + 'var map = result[' + quote(field.name) + '] = {};');
           lines.push(indent + 'var length = bb.readVarUint();');
          lines.push(indent + 'while (length-- > 0) { map[' + code[0] + '] = ' + code[1] + '; }');
         }
    } else {
      code = code as string;
      if (field.isDeprecated) {
        lines.push(indent + code + ';');
      } else {
        lines.push(indent + 'result[' + quote(field.name) + '] = ' + code + ';');
      }
    }

    if (definition.kind === 'MESSAGE') {
      lines.push('      break;');
      lines.push('');
    }
  }

  if (definition.kind === 'MESSAGE') {
    lines.push('    default:');
    lines.push('      throw new Error("Attempted to parse invalid message");');
    lines.push('    }');
    lines.push('  }');
  }

  else {
    lines.push('  return result;');
  }

  lines.push('}');

  return lines.join('\n');
}

function encodeCodeForField(field: Field, definitions: {[name: string]: Definition}): string | string[] {
    let code: string;

    switch (field.type) {
      case 'bool': {
        code = 'bb.writeByte(value);';
        break;
      }

      case 'byte': {
        code = 'bb.writeByte(value);';  // only used if not array
        break;
      }

      case 'int': {
        code = 'bb.writeVarInt(value);';
        break;
      }

      case 'uint': {
        code = 'bb.writeVarUint(value);';
        break;
      }

      case 'float':
      case 'float32': {
        code = 'bb.writeVarFloat(value);';
        break;
      }

      case 'string': {
        code = 'bb.writeString(value);';
        break;
      }

      case 'map': {
        let keyCode = encodeCodeForField({ type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null }, definitions);
        if (!keyCode || nativeTypes.indexOf(field.mapKeyType as string) < 0) {
          error(
            'Invalid type ' +
              quote(field.mapKeyType as string) +
              ' for map key for ' +
              quote(field.name) +
              '.  Map keys must be native fields.',
            field.line,
            field.column
          );
         }
        let mapValueField = { type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
        let valueCode = encodeCodeForField(mapValueField, definitions);
        if (!valueCode) {
          error(
            'Invalid type ' +
              quote(field.mapValueType as string) +
              ' for map value for ' +
              quote(field.name),
            field.line,
            field.column
          );
         }
        return [keyCode as string, valueCode as string];
      }

      default: {
        let type = definitions[field.type!];
        if (!type) {
          throw new Error('Invalid type ' + quote(field.type!) + ' for field ' + quote(field.name));
        } else if (type.kind === 'ENUM') {
          code =
            'var encoded = this[' + quote(type.name) + '][value]; ' +
            'if (encoded === void 0) throw new Error("Invalid value " + JSON.stringify(value) + ' + quote(' for enum ' + quote(type.name)) + '); ' +
            'bb.writeVarUint(encoded);';
        } else {
          code = 'this[' + quote('encode' + type.name) + '](value, bb);';
        }
      }
    }
    return code
}

function compileEncode(definition: Definition, definitions: {[name: string]: Definition}): string {
  let lines: string[] = [];

  lines.push('function(message, bb) {');
  lines.push('  var isTopLevel = !bb;');
  lines.push('  if (isTopLevel) bb = new this.ByteBuffer();');

  for (let j = 0; j < definition.fields.length; j++) {
    let field = definition.fields[j];

    if (field.isDeprecated) {
      continue;
    }

    let code = encodeCodeForField(field, definitions);

    lines.push('');
    lines.push('  var value = message[' + quote(field.name) + '];');
    lines.push('  if (value != null) {'); // Comparing with null using "!=" also checks for undefined

    if (definition.kind === 'MESSAGE') {
      lines.push('    bb.writeVarUint(' + field.value + ');');
    }

    if (field.isArray) {
      if (field.type === 'byte') {
        lines.push('    bb.writeByteArray(value);');
      } else {
        lines.push('    var values = value, n = values.length;');
        lines.push('    bb.writeVarUint(n);');
        lines.push('    for (var i = 0; i < n; i++) {');
        lines.push('      value = values[i];');
        lines.push('      ' + code);
        lines.push('    }');
      }
    } else if (field.isMap) {
        lines.push('    var obj = value, keys = Object.keys(obj), n = keys.length;');
        lines.push('    bb.writeVarUint(n);');
        lines.push('    for (var i = 0; i < keys.length; i++) {');
        lines.push('      value = keys[i];');
        lines.push('      ' + code[0]);
        lines.push('      value = obj[keys[i]];');
        lines.push('      ' + code[1]);
        lines.push('    }');
    } else {
      lines.push('    ' + code);
    }

    if (definition.kind === 'STRUCT') {
      lines.push('  } else {');
      lines.push('    throw new Error(' + quote('Missing required field ' + quote(field.name)) + ');');
    }

    lines.push('  }');
  }

  // A field id of zero is reserved to indicate the end of the message
  if (definition.kind === 'MESSAGE') {
    lines.push('  bb.writeVarUint(0);');
  }

  lines.push('');
  lines.push('  if (isTopLevel) return bb.toUint8Array();');
  lines.push('}');

  return lines.join('\n');
}

export function compileSchemaJS(schema: Schema): string {
  let definitions: {[name: string]: Definition} = {};
  let name = schema.package;
  let js: string[] = [];

  if (name !== null) {
    js.push('var ' + name + ' = exports || ' + name + ' || {}, exports;');
  } else {
    js.push('var exports = exports || {};');
    name = 'exports';
  }

  js.push(name + '.ByteBuffer = ' + name + '.ByteBuffer || require("kiwi-schema").ByteBuffer;');

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];
    definitions[definition.name] = definition;
  }

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];

    switch (definition.kind) {
      case 'ENUM': {
        let value: any = {};
        for (let j = 0; j < definition.fields.length; j++) {
          let field = definition.fields[j];
          value[field.name] = field.value;
          value[field.value] = field.name;
        }
        js.push(name + '[' + quote(definition.name) + '] = ' + JSON.stringify(value, null, 2) + ';');
        break;
      }

      case 'STRUCT':
      case 'MESSAGE': {
        js.push('');
        js.push(name + '[' + quote('decode' + definition.name) + '] = ' + compileDecode(definition, definitions) + ';');
        js.push('');
        js.push(name + '[' + quote('encode' + definition.name) + '] = ' + compileEncode(definition, definitions) + ';');
        break;
      }

      default: {
        error('Invalid definition kind ' + quote(definition.kind), definition.line, definition.column);
        break;
      }
    }
  }

  js.push('');
  return js.join('\n');
}

export function compileSchema(schema: Schema): any {
  let result = {
    ByteBuffer: ByteBuffer,
  };
  new Function('exports', compileSchemaJS(schema))(result);
  return result;
}
