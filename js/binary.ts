import { ByteBuffer } from "./bb";
import { Schema, Field, Definition, DefinitionKind } from "./schema";

let types: (string | null)[] = ['bool', 'byte', 'int', 'uint', 'float', 'string', 'float32'];
let kinds: DefinitionKind[] = ['ENUM', 'STRUCT', 'MESSAGE'];

export function decodeBinarySchema(buffer: Uint8Array | ByteBuffer): Schema {
  let bb = buffer instanceof ByteBuffer ? buffer : new ByteBuffer(buffer);
  let definitionCount = bb.readVarUint();
  let definitions: Definition[] = [];

  // Read in the schema
  for (let i = 0; i < definitionCount; i++) {
    let definitionName = bb.readString();
    let kind = bb.readByte();
    let fieldCount = bb.readVarUint();
    let fields: Field[] = [];

    for (let j = 0; j < fieldCount; j++) {
      let fieldName = bb.readString();
      let type = bb.readVarInt();
      let flags = bb.readByte();
      let isArray = !!(flags & 1);
      let isMap = !!(flags & 2);
      let value = bb.readVarUint();
      let mapKeyType: string | null = null;
      let mapValueType: string | null = null;
      if (isMap) {
        mapKeyType = type as any;
        type = null as any;
        mapValueType = bb.readVarInt() as any;
      }
      fields.push({
        name: fieldName,
        line: 0,
        column: 0,
        type: (kinds[kind] === 'ENUM' || isMap) ? null : type as any,
        isArray: isArray,
        isMap: isMap,
        isDeprecated: false,
        value: value,
        mapKeyType: mapKeyType as any,
        mapValueType: mapValueType as any,
      });
    }

    definitions.push({
      name: definitionName,
      line: 0,
      column: 0,
      kind: kinds[kind],
      fields: fields,
    });
  }

  // Bind type names afterwards
  for (let i = 0; i < definitionCount; i++) {
    let fields = definitions[i].fields;
    for (let j = 0; j < fields.length; j++) {
      let field = fields[j];
      if (field.isMap) {
        field.type = 'map'
        field.mapKeyType = bind(field.mapKeyType as number | null, definitions)
        field.mapValueType = bind(field.mapValueType as number | null, definitions)
      } else {
        field.type = bind(field.type as number | null, definitions)
      }
    }
  }

  return {
    package: null,
    definitions: definitions,
  };
}

export function bind(type: number | null, definitions: Definition[]): string | null {
  if (type !== null && type < 0) {
    if (~type >= types.length) {
      throw new Error('Invalid type ' + type);
    }
    return types[~type];
  }

  if (type !== null && type >= definitions.length) {
    throw new Error('Invalid type ' + type);
  }
  return type === null ? null : definitions[type].name;
}

export function encodeBinarySchema(schema: Schema): Uint8Array {
  let bb = new ByteBuffer();
  let definitions = schema.definitions;
  let definitionIndex: {[name: string]: number} = {};

  bb.writeVarUint(definitions.length);

  for (let i = 0; i < definitions.length; i++) {
    definitionIndex[definitions[i].name] = i;
  }

  for (let i = 0; i < definitions.length; i++) {
    let definition = definitions[i];

    bb.writeString(definition.name);
    bb.writeByte(kinds.indexOf(definition.kind));
    bb.writeVarUint(definition.fields.length);

    for (let j = 0; j < definition.fields.length; j++) {
      let field = definition.fields[j];

      bb.writeString(field.name);
      if (field.isMap) {
        let type = types.indexOf(field.mapKeyType);
        bb.writeVarInt(type === -1 ? definitionIndex[field.mapKeyType!] : ~type);
      } else {
        let type = types.indexOf(field.type);
        bb.writeVarInt(type === -1 ? definitionIndex[field.type!] : ~type);
      }
      bb.writeByte((field.isArray ? 1 : 0) | (field.isMap ? 2 : 0));
      bb.writeVarUint(field.value);
      if (field.isMap) {
        let type = types.indexOf(field.mapValueType);
        bb.writeVarInt(type === -1 ? definitionIndex[field.mapValueType!] : ~type);
      }
    }
  }

  return bb.toUint8Array();
}
