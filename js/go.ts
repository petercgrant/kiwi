import { Schema, Definition, Field } from "./schema";
import { error, quote } from "./util";
import { nativeTypes } from "./parser";

function decodeCodeForField(field: Field, definitions: {[name: string]: Definition}, value: string, indent: string): string | string[] {
  let code;

  switch (field.type) {
    case 'bool': {
      code = 'bb.ReadBool()';
      break;
    }

    case 'byte': {
      code = 'bb.ReadByte()';
      break;
    }

    case 'int': {
      code = 'bb.ReadVarInt()';
      break;
    }

    case 'uint': {
      code = 'bb.ReadVarUint()';
      break;
    }

    case 'float':
    case 'float32': {
      code = 'bb.ReadVarFloat()';
      break;
    }

    case 'string': {
      code = 'bb.ReadString()';
      break;
    }

    case 'map': {
      let keyCode = decodeCodeForField({ type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null }, definitions, '&key', indent);
      if (!keyCode || nativeTypes.indexOf(field.mapKeyType!) < 0) {
        error(
          'Invalid type ' +
          quote(field.mapKeyType!) +
          ' for map key for ' +
          quote(field.name) +
          '.  Map keys must be native fields.',
          field.line,
          field.column
        );
      }
      let mapValueField = { type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
      let valueCode = decodeCodeForField(mapValueField, definitions, '&val', indent);
      if (!valueCode) {
        error(
          'Invalid type ' +
          quote(field.mapValueType!) +
          ' for map value for ' +
          quote(field.name),
          field.line,
          field.column
        );
      }

      code = [keyCode as string, valueCode as string];
      break;
    }
    default: {
      var type = definitions[field.type!];

      if (!type) {
        error('Invalid type ' + quote(field.type!) + ' for field ' + quote(field.name), field.line, field.column);
      }

      else if (type.kind === 'ENUM') {
        var gotype = goType(definitions, field, false);
        var lines = [];
        lines.push('func() (val ' + gotype + ', ok bool) {');
        lines.push('\tvar v uint32');
        lines.push('\tif v, ok = bb.ReadVarUint(); ok {');
        lines.push('\t\tval = ' + gotype + '(v)');
        lines.push('\t}');
        lines.push('\treturn');
        lines.push('}()');
        code = lines.join('\n' + indent);
      }

      else {
        var gotype = goType(definitions, field, false);
        var isPointer = goIsFieldPointer(definitions, field);
        var lines = [];
        lines.push('func() (val ' + (isPointer ? '*' : '') + gotype + ', ok bool) {');
        lines.push('\tv := ' + (isPointer ? '&' : '') + gotype + '{}');
        lines.push('\tif ok = v.Decode(bb, schema); ok {');
        lines.push('\t\tval = v');
        lines.push('\t}');
        lines.push('\treturn');
        lines.push('}()');
        code = lines.join('\n' + indent)
      }
    }
  }

  return code;
}

function encodeCodeForField(field: Field, definitions: {[name: string]: Definition}, value: string, indent: string): string | string[] {
  let code: string;

  switch (field.type) {
    case 'bool': {
      code = 'bb.WriteBool(' + value + ')';
      break;
    }

    case 'byte': {
      code = 'bb.WriteByte(' + value + ')';
      break;
    }

    case 'int': {
      code = 'bb.WriteVarInt(' + value + ')';
      break;
    }

    case 'uint': {
      code = 'bb.WriteVarUint(' + value + ')';
      break;
    }

    case 'float':
    case 'float32': {
      code = 'bb.WriteVarFloat(' + value + ')';
      break;
    }

    case 'string': {
      code = 'bb.WriteString(' + value + ')';
      break;
    }

    case 'map': {
      let keyCode = encodeCodeForField({ type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null }, definitions, 'key', indent);
      if (!keyCode || nativeTypes.indexOf(field.mapKeyType!) < 0) {
        error(
          'Invalid type ' +
          quote(field.mapKeyType!) +
          ' for map key for ' +
          quote(field.name) +
          '.  Map keys must be native fields.',
          field.line,
          field.column
        );
      }
      let mapValueField = { type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
      let valueCode = encodeCodeForField(mapValueField, definitions, 'val', indent);
      if (!valueCode) {
        error(
          'Invalid type ' +
          quote(field.mapValueType!) +
          ' for map value for ' +
          quote(field.name),
          field.line,
          field.column
        );
      }
      return [keyCode as string, valueCode as string];
    }

    default: {
      var type = definitions[field.type!];

      if (!type) {
        error('Invalid type ' + quote(field.type!) + ' for field ' + quote(field.name), field.line, field.column);
      }

      else if (type.kind === 'ENUM') {
        code = 'bb.WriteVarUint(uint32(' + value + '))';
      }

      else {
        var lines = [];
        lines.push('if !' + value + '.Encode(bb) {');
        lines.push('\treturn false');
        lines.push('}');
        code = lines.join('\n' + indent);
      }
    }
  }
  return code;
}

function goType(definitions: {[name: string]: Definition}, field: Field, isArray: boolean): string {
  let type;

  switch (field.type) {
    case 'bool': type = 'bool'; break;
    case 'byte': type = 'byte'; break;
    case 'int': type = 'int32'; break;
    case 'uint': type = 'uint32'; break;
    case 'float':
    case 'float32': type = 'float32'; break;
    case 'string': type = 'string'; break;
    case 'map': type = 'kiwi.LinkedMap'; break;

    default: {
      let definition = definitions[field.type!];

      if (!definition) {
        error('Invalid type ' + quote(field.type!) + ' for field ' + quote(field.name), field.line, field.column);
      }

      type = definition.name;
      break;
    }
  }

  if (isArray) {
    type = '[]' + type;
  }

  return type;
}

function goFieldName(field: Field): string {
  return '_data_' + field.name;
}

function goAccessorName(field: string): string {
  return 'Get' + field
}

function goSetterName(field: string): string {
  return 'Set' + field
}

function goFlagIndex(i: number): number {
  return i >> 5;
}

function goFlagMask(i: number): number {
  return 1 << (i % 32) >>> 0;
}

function goIsFieldPointer(definitions: {[name: string]: Definition}, field: Field): boolean {
  return !field.isArray && field.type! in definitions && definitions[field.type!].kind !== 'ENUM';
}

function printCodeForField(field: Field, definitions: {[name: string]: Definition}, value: string, indent: string): string {
  let key = field.isArray ? "" : (field.name ? field.name : value);
  let val = field.isArray ? "&_it" : (field.name ? ('p.' + goAccessorName(field.name) + '()') : value);
  return 'printer.Print("' + key + '", ' + val + ')';
}

export function compileSchemaGo(schema: Schema): string {
  let definitions: {[name: string]: Definition} = {};
  let go: string[] = [];

  go.push('package ' + (schema.package === null ? 'schema' : schema.package));
  go.push('');

  go.push('import "mnk.ee/kiwi"');
  go.push('');

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];
    definitions[definition.name] = definition;
  }

  go.push('type BinarySchema struct {');
  go.push('\tschema kiwi.BinarySchema');

  for (var i = 0; i < schema.definitions.length; i++) {
    var definition = schema.definitions[i];
    if (definition.kind === 'MESSAGE') {
      go.push('\tindex' + definition.name + ' uint32');
    }
  }

  go.push('}');

  go.push('');

  for (var i = 0; i < schema.definitions.length; i++) {
    var definition = schema.definitions[i];

    if (definition.kind === 'ENUM') {
      go.push('type ' + definition.name + ' uint32');
      go.push('const (');
      go.push('');
      for (var j = 0; j < definition.fields.length; j++) {
        var field = definition.fields[j];
        go.push('\t' + definition.name + '_' + field.name + ' ' + definition.name + ' = ' + field.value);
      }
      go.push(')');
      go.push('');
    }

    else if (definition.kind !== 'STRUCT' && definition.kind !== 'MESSAGE') {
      error('Invalid definition kind ' + quote(definition.kind), definition.line, definition.column);
    }
  }

  for (var pass = 0; pass < 3; pass++) {
    var newline = false;

    if (pass === 2) {
      go.push('');

      go.push('func (bs *BinarySchema) Parse(bb *kiwi.ByteBuffer) bool {');
      go.push('\tif !bs.schema.Parse(bb) {');
      go.push('\t\treturn false');
      go.push('\t}');

      go.push('');

      for (var i = 0; i < schema.definitions.length; i++) {
        var definition = schema.definitions[i];
        if (definition.kind === 'MESSAGE') {
          go.push('\tbs.schema.FindDefinition("' + definition.name + '", &bs.index' + definition.name + ')');
        }
      }

      go.push('\treturn true');
      go.push('}');
      go.push('');

      for (var i = 0; i < schema.definitions.length; i++) {
        var definition = schema.definitions[i];
        if (definition.kind === 'MESSAGE') {
          go.push('func (bs *BinarySchema) Skip' + definition.name + 'Field(bb *kiwi.ByteBuffer, id uint32) bool {');
          go.push('\treturn bs.schema.SkipField(bb, bs.index' + definition.name + ', id)');
          go.push('}');
          go.push('');
        }
      }
    }

    for (var i = 0; i < schema.definitions.length; i++) {
      var definition = schema.definitions[i];

      if (definition.kind === 'ENUM') {
        continue;
      }

      var fields = definition.fields;

      if (pass === 0) {
        newline = true;
      }

      else if (pass === 1) {
        go.push('type ' + definition.name + ' struct {');
        go.push('\t_flags [' + (fields.length + 31 >> 5) + ']uint32 `spew:"-"`');

        // Sort fields by size since that makes the resulting struct smaller
        let sizes: {[name: string]: number} = {'bool': 1, 'byte': 1, 'int': 4, 'uint': 4, 'float': 4, 'float32': 4};
        let sortedFields = fields.slice().sort(function (a, b) {
          var sizeA = !a.isArray && sizes[a.type as string] || 8;
          var sizeB = !b.isArray && sizes[b.type as string] || 8;
          if (sizeA !== sizeB) return sizeB - sizeA;
          return fields.indexOf(a) - fields.indexOf(b); // Make sure the sort is stable
        });

        for (var j = 0; j < sortedFields.length; j++) {
          var field = sortedFields[j];

          if (field.isDeprecated) {
            continue;
          }

          var name = goFieldName(field);
          var type = goType(definitions, field, field.isArray);

          if (goIsFieldPointer(definitions, field)) {
            go.push('\t' + name + ' *' + type + '`spew:"' + field.name + '"`');
          } else {
            go.push('\t' + name + ' ' + type + '`spew:"' + field.name + '"`');
          }
        }

        go.push('}');
        go.push('');
      }

      else {
        for (var j = 0; j < fields.length; j++) {
          var field = fields[j];
          var name = goFieldName(field);
          var type = goType(definitions, field, field.isArray);
          var flagIndex = goFlagIndex(j);
          var flagMask = goFlagMask(j);

          if (field.isDeprecated) {
            continue;
          }

          if (goIsFieldPointer(definitions, field)) {
            go.push('func (p *' + definition.name + ') ' + goAccessorName(field.name) + '() *' + type + ' {');
            go.push('\tif p == nil {');
            go.push('\t\treturn nil');
            go.push('\t}');
            go.push('\treturn p.' + name);
            go.push('}');
            go.push('');

            go.push('func (p *' + definition.name + ') ' + goSetterName(field.name) + '(value *' + type + ') {');
            go.push('\tp.' + name + ' = value;');
            go.push('}');
            go.push('');
          }

          else if (field.isArray) {
            go.push('func (p *' + definition.name + ') ' + goAccessorName(field.name) + '() ' + type + ' {');
            go.push('\tif p._flags[' + flagIndex + '] & ' + flagMask + ' > 0 {');
            go.push('\t\treturn p.' + name);
            go.push('\t}');
            go.push('\treturn nil');
            go.push('}');
            go.push('');

            // TODO: support memory pool?
            go.push('func (p *' + definition.name + ') ' + goSetterName(field.name) + '(count uint32) ' + type + ' {');
            go.push('\tp._flags[' + flagIndex + '] |= ' + flagMask);
            go.push('\tp.' + name + ' = make(' + type + ', int(count))');
            go.push('\treturn p.' + name);
            go.push('}');
            go.push('');
          }

          else {
            go.push('func (p *' + definition.name + ') ' + goAccessorName(field.name) + '() *' + type + ' {');
            go.push('\tif p._flags[' + flagIndex + '] & ' + flagMask + ' > 0 {');
            go.push('\t\treturn &p.' + name);
            go.push('\t}');
            go.push('\treturn nil');
            go.push('}');
            go.push('');

            go.push('func (p *' + definition.name + ') ' + goSetterName(field.name) + '(value ' + type + ') {');
            go.push('\tp._flags[' + flagIndex + '] |= ' + flagMask);
            go.push('\tp.' + name + ' = value');
            go.push('}');
            go.push('');
          }
        }

        go.push('func (p *' + definition.name + ') Encode(bb *kiwi.ByteBuffer) bool {');

        for (var j = 0; j < fields.length; j++) {
          var field = fields[j];

          if (field.isDeprecated) {
            continue;
          }

          var name = goFieldName(field);
          var value = field.isArray ? '_it' : ('p.' + name);
          var flagIndex = goFlagIndex(j);
          var flagMask = goFlagMask(j);
          var indent = '\t';

          if (definition.kind === 'STRUCT') {
            go.push('\tif p.' + goAccessorName(field.name) + '() == nil {');
            go.push('\t\treturn false');
            go.push('\t}');
          } else {
            go.push('\tif p.' + goAccessorName(field.name) + '() != nil {');
            indent = '\t\t';
          }

          let code = encodeCodeForField(field, definitions, value, field.isArray || field.isMap ? indent + '\t' : indent);

          if (definition.kind === 'MESSAGE') {
            go.push(indent + 'bb.WriteVarUint(' + field.value + ')');
          }

          if (field.isArray) {
            go.push(indent + 'bb.WriteVarUint(uint32(len(p.' + name + ')))');
            go.push(indent + 'for _, _it := range p.' + name + '{');
            go.push(indent + '\t' + code);
            go.push(indent + '}')
          } else if (field.isMap) {
            let keyField = { type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
            let valueField = { type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null };
            go.push(indent + 'bb.WriteVarUint(uint32(p.' + name + '.Len()))');
            go.push(indent + 'for it := p.' + name + '.Iter(); it.HasNext(); {');
            go.push(indent + '\tkif, vif := it.Next()');
            go.push(indent + '\tkey := kif.(' + goType(definitions, keyField, false) + ')');
            go.push(indent + '\tval := vif.(' + (goIsFieldPointer(definitions, valueField) ? '*' : '') + goType(definitions, valueField, false) + ')');
            go.push(indent + '\t' + code[0]);
            go.push(indent + '\t' + code[1]);
            go.push(indent + '}')
          } else {
            go.push(indent + code);
          }

          if (definition.kind !== 'STRUCT') {
            go.push('\t}');
          }
        }

        if (definition.kind === 'MESSAGE') {
          go.push('\tbb.WriteVarUint(0)');
        }

        go.push('\treturn true;');
        go.push('}');
        go.push('');

        go.push('func (p *' + definition.name + ') Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {');

        for (var j = 0; j < fields.length; j++) {
          if (fields[j].isArray || fields[j].isMap) {
            go.push('\tvar count uint32');
            break;
          }
        }
        if (fields.length > 0) {
          go.push('\tvar ok bool');
          go.push('\t_ = ok');
        }

        if (definition.kind === 'MESSAGE') {
          go.push('\thighest := uint32(0)');
          go.push('\tfor {');
          go.push('\t\ttyp, ok := bb.ReadVarUint()');
          go.push('\t\tif !ok {');
          go.push('\t\t\treturn false');
          go.push('\t\t}');
          go.push('\t\tif typ < highest && typ > 0 {');
          go.push('\t\t\treturn false');
          go.push('\t\t}');
          go.push('\t\thighest = typ');
          go.push('\t\tswitch typ {');
          go.push('\t\t\tcase 0:');
          go.push('\t\t\t\treturn true');
        }

        for (var j = 0; j < fields.length; j++) {
          let field = fields[j];
          let name = goFieldName(field);
          let isPointer = goIsFieldPointer(definitions, field);
          let value = field.isArray ? '_it' : ('p.' + name);

          let type = goType(definitions, field, false);
          var indent = '\t';

          if (definition.kind === 'MESSAGE') {
            go.push('\t\t\tcase ' + field.value + ':');
            indent = '\t\t\t\t';
          }
          let code = decodeCodeForField(field, definitions, value, field.isArray || field.isMap ? indent + '\t' : indent);

          if (field.isArray) {
            go.push(indent + 'count, ok = bb.ReadVarUint()');
            go.push(indent + 'if !ok {');
            go.push(indent + '\treturn false');
            go.push(indent + '}');

            if (field.isDeprecated) {
              go.push(indent + 'for j := uint32(0); j < count; j++ {');
              go.push(indent + '\tif _, ok = ' + code + '; !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '}');
            } else {
              go.push(indent + field.name + ' := p.' + goSetterName(field.name) + '(count)');
              go.push(indent + 'for j := range ' + field.name + ' {');
              go.push(indent + '\t' + field.name + '[j], ok = ' + code);
              go.push(indent + '\tif !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '}');
            }
          } else if (field.isMap) {
            if (field.isDeprecated) {
              go.push(indent + 'count, ok = bb.ReadVarUint()');
              go.push(indent + 'if !ok {');
              go.push(indent + '\treturn false');
              go.push(indent + '}');
              go.push(indent + 'for j := uint32(0); j < count; j++ {');
              go.push(indent + '\tif _, ok = ' + code[0] + '; !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '\tif _, ok = ' + code[1] + '; !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '}');
            } else {
              go.push(indent + 'count, ok = bb.ReadVarUint()');
              go.push(indent + 'if !ok {');
              go.push(indent + '\treturn false');
              go.push(indent + '}');
              go.push(indent + "p." + goSetterName(field.name) + '(kiwi.NewLinkedMap(int(count)))');
              go.push(indent + 'for j := uint32(0); j < count; j++ {');
              go.push(indent + '\tkey, ok := ' + code[0]);
              go.push(indent + '\tif !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '\tval, ok := ' + code[1]);
              go.push(indent + '\tif !ok {');
              go.push(indent + '\t\treturn false');
              go.push(indent + '\t}');
              go.push(indent + '\t' + value + '.Set(key, val)');
              go.push(indent + '}');
            }
          } else {
            if (field.isDeprecated) {
              go.push(indent + 'if _, ok = ' + code + '; !ok {');
              go.push(indent + '\treturn false');
              go.push(indent + '}');
            }

            else {
              go.push(indent + 'if p.' + name + ', ok = ' + code + '; !ok {');
              go.push(indent + '\treturn false');
              go.push(indent + '}');

              if (!isPointer) {
                go.push(indent + 'p.' + goSetterName(field.name) + '(p.' + name + ')');
              }
            }
          }

          // TODO: make this fallthrough for not-MESSAGE?
          if (definition.kind === 'MESSAGE') {
            go.push('\t\t\t\tbreak');
          }
        }

        if (definition.kind === 'MESSAGE') {
          go.push('\t\t\tdefault:');
          go.push('\t\t\t\tif schema == nil || !schema.Skip' + definition.name + 'Field(bb, typ) {');
          go.push('\t\t\t\t\treturn false');
          go.push('\t\t\t\t}');
          go.push('\t\t\t\tbreak');
          go.push('\t\t}');
          go.push('\t}');
        }

        else {
          go.push('\treturn true');
        }

        go.push('}');
        go.push('');

        go.push('func (p *' + definition.name + ') Print(printer *kiwi.Printer) bool {');

        go.push('\tif p == nil { return true }')
        go.push('\tprinter = printer.With(p)')
        go.push('\tprinter.StartObject()')

        for (var j = 0; j < fields.length; j++) {
          let field = fields[j];

          if (field.isDeprecated) {
            continue;
          }

          let name = goFieldName(field);
          let value = field.isArray ? '_it' : ('p.' + name);
          let flagIndex = goFlagIndex(j);
          let flagMask = goFlagMask(j);
          var indent = '\t';

          go.push('\tif p.' + goAccessorName(field.name) + '() != nil {');
          indent = '\t\t';

          let code = printCodeForField(field, definitions, value, field.isArray || field.isMap ? indent + '\t' : indent);

          if (field.isArray) {
            go.push(indent + 'printer.Field("' + field.name + '")')
            go.push(indent + 'printer.StartArray()');
            go.push(indent + 'for _, _it := range p.' + name + '{');
            go.push(indent + '\t' + code);
            go.push(indent + '}')
            go.push(indent + 'printer.EndArray()');
          } else {
            go.push(indent + code);
          }

          go.push('\t}');
        }

        go.push('\tprinter.EndObject()');

        go.push('\treturn true;');
        go.push('}');
        go.push('');
      }
    }

    if (pass === 2) {
      go.push('');
    }

    else if (newline) go.push('');
  }

  return go.join('\n');
}