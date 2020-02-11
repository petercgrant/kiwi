export interface Schema {
  package: string | null;
  definitions: Definition[];
}

export type DefinitionKind = 'ENUM' | 'STRUCT' | 'MESSAGE';

export interface Definition {
  name: string;
  line: number;
  column: number;
  kind: DefinitionKind;
  fields: Field[];
}

export interface Field {
  name: string;
  line: number;
  column: number;
  type: string | null;
  isArray: boolean;
  isMap: boolean;
  isDeprecated: boolean;
  mapKeyType: string | null;
  mapValueType: string | null;
  value: number;
}
