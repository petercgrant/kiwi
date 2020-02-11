import { Field, Schema } from "./schema";
import { error, quote } from "./util";

function tsTypeForField(field: Field): string {
  let type: string;
  switch (field.type) {
    case 'bool':
      type = 'boolean';
      break;
    case 'byte':
    case 'int':
    case 'uint':
    case 'float':
      type = 'number';
      break;
    case 'float32':
      type = 'Float32Array';
      field.isArray = false;
      break;
    case 'map':
      var keyType = tsTypeForField({ type: field.mapKeyType, name: field.name + '.key', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null });
      var valueType = tsTypeForField({ type: field.mapValueType, name: field.name + '.value', line: field.line, column: field.column, isArray: false, isMap: false, isDeprecated: field.isDeprecated, value: field.value, mapKeyType: null, mapValueType: null });
      type = '{ [key: ' + keyType + ']: ' + valueType + ' }'
      break;
    default:
      type = field.type as string;
      break;
  }

  if (field.type === 'byte' && field.isArray) type = 'Uint8Array';
  else if (field.isArray) type += '[]';
  return type;
}

export function compileSchemaTypeScript(schema: Schema): string {
  var indent = '';
  var lines = [];

  if (schema.package !== null) {
    lines.push('export namespace ' + schema.package + ' {');
    indent += '  ';
  }

  for (var i = 0; i < schema.definitions.length; i++) {
    var definition = schema.definitions[i];

    if (definition.kind === 'ENUM') {
      lines.push(indent + 'export enum ' + definition.name + ' {');

      for (var j = 0; j < definition.fields.length; j++) {
        lines.push(indent + '  ' + definition.fields[j].name + ' = ' + JSON.stringify(definition.fields[j].name) + ',');
      }

      lines.push(indent + '}');
      lines.push('')
}
  }

  for (var i = 0; i < schema.definitions.length; i++) {
    var definition = schema.definitions[i];

    if (definition.kind === 'STRUCT' || definition.kind === 'MESSAGE') {
      lines.push(indent + 'export interface ' + definition.name + ' {');

      for (var j = 0; j < definition.fields.length; j++) {
        var field = definition.fields[j];

        if (field.isDeprecated) {
          continue;
        }

        let type = tsTypeForField(field);

        lines.push(indent + '  ' + field.name + (definition.kind === 'MESSAGE' ? '?' : '') + ': ' + type + ';');
      }

      lines.push(indent + '}');
      lines.push('');
    }

    else if (definition.kind !== 'ENUM') {
      error('Invalid definition kind ' + quote(definition.kind), definition.line, definition.column);
    }
  }

  lines.push(indent + 'export interface Schema {');

  for (var i = 0; i < schema.definitions.length; i++) {
    var definition = schema.definitions[i];

    if (definition.kind === 'ENUM') {
      lines.push(indent + '  ' + definition.name + ': any;');
    }

    else if (definition.kind === 'STRUCT' || definition.kind === 'MESSAGE') {
      lines.push(indent + '  encode' + definition.name + '(message: ' + definition.name + '): Uint8Array;');
      lines.push(indent + '  decode' + definition.name + '(buffer: Uint8Array): ' + definition.name + ';');
    }
  }

  lines.push(indent + '}');

  if (schema.package !== null) {
    lines.push('}');
  }

  lines.push('');
  return lines.join('\n');
}
