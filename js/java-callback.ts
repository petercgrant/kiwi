import { Schema, Definition, Field } from "./schema";
import { error, quote } from "./util";
import { Options } from "./options";
import { nativeTypes } from "./parser";

interface Argument {
  type: string;
  name: string;
}

function javaType(
  definitions: { [name: string]: Definition },
  field: Field,
  isArray: boolean
): string {
  let type;

  switch (field.type) {
    case "bool":
      type = "Boolean";
      break;
    case "byte":
      type = "Byte";
      break;
    case "int":
      type = "Integer";
      break;
    case "uint":
      type = "Long";
      break;
    case "float":
    case "float32":
      type = "Float";
      break;
    case "string":
      type = "String";
      break;
    case "map":
      type = "LinkedHashMap";
      break;

    default: {
      let definition = definitions[field.type!];

      if (!definition) {
        error(
          "Invalid type " +
            quote(field.type!) +
            " for field " +
            quote(field.name),
          field.line,
          field.column
        );
      }

      if (definition.kind === "ENUM") {
        type = "Long";
      } else {
        type = definition.name;
      }
      break;
    }
  }

  if (isArray) {
    type = "[]" + type;
  }

  return type;
}

function decodeCodeForField(
  field: Field,
  definitions: { [name: string]: Definition },
  value: string,
  indent: string
): string[] {
  let decl = javaType(definitions, field, false) + " " + value;
  let throwError =
    ' throw new IllegalArgumentException("reading ' + field.name + '");';

  switch (field.type) {
    case "bool": {
      return [
        indent + decl + " = this.readBool();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "byte": {
      return [
        indent + decl + " = this.readByte();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "int": {
      return [
        indent + decl + " = this.readVarInt();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "uint": {
      return [
        indent + decl + " = this.readVarUint();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "float":
    case "float32": {
      return [
        indent + decl + " = this.readVarFloat();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "string": {
      return [
        indent + decl + " = this.readString();",
        indent + "if (" + value + " == null)" + throwError
      ];
    }

    case "map": {
      error(
        "Invalid type " +
          quote(field.type!) +
          " for field " +
          quote(field.name),
        field.line,
        field.column
      );
    }
    default: {
      var type = definitions[field.type!];

      if (!type) {
        error(
          "Invalid type " +
            quote(field.type!) +
            " for field " +
            quote(field.name),
          field.line,
          field.column
        );
      } else if (type.kind === "ENUM") {
        return [
          indent + decl + " = this.readVarUint();",
          indent + "if (" + value + " == null)" + throwError
        ];
      } else {
        var jtype = javaType(definitions, field, false);
        return [indent + "parse" + jtype + "(visitor);"];
      }
    }
  }
}

function argumentForField(
  definitions: { [name: string]: Definition },
  field: Field
): Argument | null {
  let name = field.name;
  let type = field.type;
  switch (type) {
    case "bool":
      return { type: "boolean ", name: name! };
    case "byte":
      return { type: "byte ", name: name! };
    case "int":
      return { type: "int ", name: name! };
    case "uint":
      return { type: "long ", name: name! };
    case "float32":
    case "float":
      return { type: "float ", name: name! };
    case "string":
      return { type: "String ", name: name! };
    case "map":
      let key = argumentForField(definitions, {
        type: field.mapKeyType,
        name: field.name + ".key",
        line: field.line,
        column: field.column,
        isArray: false,
        isMap: false,
        isDeprecated: field.isDeprecated,
        value: field.value,
        mapKeyType: null,
        mapValueType: null
      });
      let value = argumentForField(definitions, {
        type: field.mapValueType,
        name: field.name + ".value",
        line: field.line,
        column: field.column,
        isArray: false,
        isMap: false,
        isDeprecated: field.isDeprecated,
        value: field.value,
        mapKeyType: null,
        mapValueType: null
      });
      if (value === null) {
        let definition = definitions[field.mapValueType!];
        value = { type: definition.name + " ", name: name! };
      }
      return {
        type: "Map<" + key!.type! + ", " + value!.type! + "> ",
        name: name!
      };
    default: {
      let definition = definitions[type!];
      if (definition.kind === "ENUM")
        return { type: definition.name + " ", name: name! };
      return null;
    }
  }
}

function argToDeclaration(arg: Argument): string {
  return arg.type + arg.name;
}

function argToName(arg: Argument): string {
  return arg.name;
}

function visitField(
  definitions: { [name: string]: Definition },
  definition: Definition,
  field: Field,
  name: string
): string {
  switch (field.type) {
    case "bool":
    case "byte":
    case "int":
    case "uint":
    case "float":
    case "float32":
    case "string":
      return "visitor." + field.type + "Value(" + name + ");";

    case "map": {
      return "visitor.visitMap" + field.type + "Value(" + name + ");";
    }
    default: {
      var type = definitions[field.type!];

      if (!type) {
        error(
          "Invalid type " +
            quote(field.type!) +
            " for field " +
            quote(field.name),
          field.line,
          field.column
        );
      } else if (type.kind === "ENUM") {
        return "visitor.uintValue(" + name + ");";
      } else {
        return "";
      }
    }
  }
}

function emitReadField(
  definitions: { [name: string]: Definition },
  definition: Definition,
  field: Field,
  indent: string
): string[] {
  let java: string[] = [];
  let name = field.name;

  java.push(
    indent +
      "visitor.startField(Field." +
      definition.name.toUpperCase() +
      "_" +
      field.name.toUpperCase() +
      ");"
  );
  if (field.isArray) {
    let count = "_" + name + "_count";
    java.push(indent + "Long " + count + " = this.readVarUint();");
    java.push(
      indent +
        "if (" +
        count +
        ' == null) throw new IllegalArgumentException("missing array count for ' +
        field.name +
        '");'
    );
    java.push(indent + "visitor.startArray(" + count + ".intValue());");
    java.push(indent + "for (long j = 0; j < " + count + "; j++) {");
    let element = name + "_element";
    java = java.concat(
      decodeCodeForField(field, definitions, element, indent + "  ")
    );
    java.push(
      indent + "  " + visitField(definitions, definition, field, element)
    );
    java.push(indent + "}");
    java.push(indent + "visitor.endArray(" + count + ".intValue());");
  } else if (field.isMap) {
    let count = "_" + name + "_count";
    java.push(indent + "Long " + count + " = this.readVarUint();");
    java.push(
      indent +
        "if (" +
        count +
        ' == null) throw new IllegalArgumentException("missing map entry count for ' +
        field.name +
        '");'
    );
    java.push(indent + "visitor.startMap(" + count + ".intValue());");
    java.push(indent + "for (long j = 0; j < " + count + "; j++) {");
    java.push(indent + "  " + "visitor.startMapEntry();");
    let keyFieldName =
      definition.name.toUpperCase() +
      "_" +
      field.name.toUpperCase() +
      "_MAP_KEY";

    java.push(
      indent + "  " + "visitor.startField(Field." + keyFieldName + ");"
    );
    let keyField = {
      type: field.mapKeyType,
      name: field.name + ".key",
      line: field.line,
      column: field.column,
      isArray: false,
      isMap: false,
      isDeprecated: field.isDeprecated,
      value: field.value,
      mapKeyType: null,
      mapValueType: null
    };
    java = java.concat(
      decodeCodeForField(keyField, definitions, "key", indent + "  ")
    );
    java.push(
      indent + "  " + visitField(definitions, definition, keyField, "key")
    );
    java.push(indent + "  " + "visitor.endField(Field." + keyFieldName + ");");
    let valueFieldName =
      definition.name.toUpperCase() +
      "_" +
      field.name.toUpperCase() +
      "_MAP_VALUE";

    java.push(
      indent + "  " + "visitor.startField(Field." + valueFieldName + ");"
    );
    let valueField = {
      type: field.mapValueType,
      name: field.name + ".value",
      line: field.line,
      column: field.column,
      isArray: false,
      isMap: false,
      isDeprecated: field.isDeprecated,
      value: field.value,
      mapKeyType: null,
      mapValueType: null
    };
    java = java.concat(
      decodeCodeForField(valueField, definitions, "value", indent + "  ")
    );
    java.push(
      indent + "  " + visitField(definitions, definition, valueField, "value")
    );
    java.push(
      indent + "  " + "visitor.endField(Field." + valueFieldName + ");"
    );
    java.push(indent + "  " + "visitor.endMapEntry();");
    java.push(indent + "}");
    java.push(indent + "visitor.endMap(" + count + ".intValue());");
  } else {
    java.push(indent + "{");
    java = java.concat(
      decodeCodeForField(field, definitions, field.name, indent + "  ")
    );
    java.push(
      indent + "  " + visitField(definitions, definition, field, field.name)
    );
    java.push(indent + "}");
  }
  java.push(
    indent +
      "visitor.endField(Field." +
      definition.name.toUpperCase() +
      "_" +
      field.name.toUpperCase() +
      ");"
  );
  return java;
}

interface EnumValue {
  name: string;
  value: number;
  isDeprecated?: boolean;
}

function defineEnum(
  name: string,
  names: EnumValue[],
  indent: string
): string[] {
  let indent0 = indent;
  let indent1 = indent0 + "  ";
  let indent2 = indent1 + "  ";
  let indent3 = indent2 + "  ";
  let java: string[] = [];
  java.push(indent0 + "public enum " + name + " {");
  let values: string[] = [];
  for (let j = names.length - 1; j >= 0; j--) {
    let field = names[j];
    if (field.isDeprecated !== true) {
      let end = values.length > 0 ? "," : ";";
      values.push(indent1 + field.name + "(" + field.value + ")" + end);
    }
  }

  java = java.concat(values.reverse());
  java.push(indent1 + "private int v;");
  java.push(
    indent1 + "private static Map<Integer, " + name + "> map = new HashMap<>();"
  );
  java.push("");
  java.push(indent1 + name + "(int v) {");
  java.push(indent2 + "this.v = v;");
  java.push(indent1 + "}");
  java.push("");
  java.push(indent1 + "static {");
  java.push(indent2 + "for (" + name + " v : " + name + ".values()) {");
  java.push(indent3 + "map.put(v.getValue(), v);");
  java.push(indent2 + "}");
  java.push(indent1 + "}");
  java.push("");
  java.push(indent1 + "public static " + name + " valueOf(int v) {");
  java.push(indent2 + "return map.get(v);");
  java.push(indent1 + "}");
  java.push("");
  java.push(indent1 + "public int getValue() {");
  java.push(indent2 + "return v;");
  java.push(indent1 + "}");
  java.push(indent0 + "}");
  java.push("");
  return java;
}

export function compileSchemaCallbackJava(
  schema: Schema,
  options: Options
): string {
  let definitions: { [name: string]: Definition } = {};
  let java: string[] = [];

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];
    definitions[definition.name] = definition;
  }

  if (schema.package !== null) {
    java.push("package " + schema.package + ";");
    java.push("");
  }

  let runtime = "";
  if (options.runtime !== undefined) {
    runtime = options.runtime;
    java.push("import " + runtime + ".ByteBuffer;");
    java.push("");
  }
  java.push("import java.util.HashMap;");
  java.push("import java.util.Map;");

  java.push("public class Schema {");

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];
    if (definition.kind === "ENUM") {
      java = java.concat(defineEnum(definition.name, definition.fields, "  "));
    }
  }

  java = java.concat(
    defineEnum(
      "StructDefinition",
      schema.definitions
        .filter(d => d.kind === "STRUCT")
        .map((d, j) => ({ name: d.name.toUpperCase(), value: j + 1 })),
      "  "
    )
  );
  java = java.concat(
    defineEnum(
      "MessageDefinition",
      schema.definitions
        .filter(d => d.kind === "MESSAGE")
        .map((d, j) => ({ name: d.name.toUpperCase(), value: j + 1 })),
      "  "
    )
  );
  java = java.concat(
    defineEnum(
      "Field",
      schema.definitions
        .filter(d => d.kind === "MESSAGE" || d.kind === "STRUCT")
        .map(d =>
          d.fields
            .map(f => d.name.toUpperCase() + "_" + f.name.toUpperCase())
            .concat(
              d.fields
                .filter(f => f.isMap)
                .map(
                  f =>
                    d.name.toUpperCase() +
                    "_" +
                    f.name.toUpperCase() +
                    "_MAP_KEY"
                ),
              d.fields
                .filter(f => f.isMap)
                .map(
                  f =>
                    d.name.toUpperCase() +
                    "_" +
                    f.name.toUpperCase() +
                    "_MAP_VALUE"
                )
            )
        )
        .reduce((acc, f) => acc.concat(f), [])
        .map((f, j) => ({ name: f, value: j + 1 })),
      "  "
    )
  );

  java.push("  public interface Visitor {");
  java.push("    void startStruct(StructDefinition def);");
  java.push("");
  java.push("    void endStruct(StructDefinition def);");
  java.push("");
  java.push("    void startMessage(MessageDefinition def);");
  java.push("");
  java.push("    void endMessage(MessageDefinition def);");
  java.push("");
  java.push("    void startArray(int len);");
  java.push("");
  java.push("    void endArray(int len);");
  java.push("");
  java.push("    void startMap(int len);");
  java.push("");
  java.push("    void endMap(int len);");
  java.push("");
  java.push("    void startMapEntry();");
  java.push("");
  java.push("    void endMapEntry();");
  java.push("");
  java.push("    void startField(Field f);");
  java.push("");
  java.push("    void endField(Field f);");
  java.push("");
  java.push("    void boolValue(boolean v);");
  java.push("");
  java.push("    void byteValue(byte v);");
  java.push("");
  java.push("    void intValue(int v);");
  java.push("");
  java.push("    void uintValue(long v);");
  java.push("");
  java.push("    void floatValue(float v);");
  java.push("");
  java.push("    void float32Value(float v);");
  java.push("");
  java.push("    void stringValue(String v);");
  java.push("  }");
  java.push("");
  java.push("  public static class VisitorBase implements Visitor {");
  java.push("    public VisitorBase() {");
  java.push("    }");
  java.push("");
  java.push("    public void startStruct(StructDefinition def) {");
  java.push("    }");
  java.push("");
  java.push("    public void endStruct(StructDefinition def) {");
  java.push("    }");
  java.push("");
  java.push("    public void startMessage(MessageDefinition def) {");
  java.push("    }");
  java.push("");
  java.push("    public void endMessage(MessageDefinition def) {");
  java.push("    }");
  java.push("");
  java.push("    public void startArray(int len) {");
  java.push("    }");
  java.push("");
  java.push("    public void endArray(int len) {");
  java.push("    }");
  java.push("");
  java.push("    public void startMap(int len) {");
  java.push("    }");
  java.push("");
  java.push("    public void endMap(int len) {");
  java.push("    }");
  java.push("");
  java.push("    public void startMapEntry() {");
  java.push("    }");
  java.push("");
  java.push("    public void endMapEntry() {");
  java.push("    }");
  java.push("");
  java.push("    public void startField(Field f) {");
  java.push("    }");
  java.push("");
  java.push("    public void endField(Field f) {");
  java.push("    }");
  java.push("");
  java.push("    public void boolValue(boolean v) {");
  java.push("    }");
  java.push("");
  java.push("    public void byteValue(byte v) {");
  java.push("    }");
  java.push("");
  java.push("    public void intValue(int v) {");
  java.push("    }");
  java.push("");
  java.push("    public void uintValue(long v) {");
  java.push("    }");
  java.push("");
  java.push("    public void floatValue(float v) {");
  java.push("    }");
  java.push("");
  java.push("    public void float32Value(float v) {");
  java.push("    }");
  java.push("");
  java.push("    public void stringValue(String v) {");
  java.push("    }");
  java.push("");
  java.push("  }");
  let suffix = ") {}";

  java.push("");

  java.push("  public static class Parser extends ByteBuffer {");
  java.push("    public Parser(byte[] buf) {");
  java.push("      super(buf);");
  java.push("    }");

  java.push("");

  for (let i = 0; i < schema.definitions.length; i++) {
    let definition = schema.definitions[i];

    if (definition.kind === "STRUCT") {
      java.push(
        "    public void parse" + definition.name + "(Visitor visitor) {"
      );
      java.push(
        "      visitor.startStruct(StructDefinition." +
          definition.name.toUpperCase() +
          ");"
      );
      for (let j = 0; j < definition.fields.length; j++) {
        let field = definition.fields[j];
        java = java.concat(
          emitReadField(definitions, definition, field, "      ")
        );
      }
      java.push(
        "      visitor.endStruct(StructDefinition." +
          definition.name.toUpperCase() +
          ");"
      );
      java.push("    }");
      java.push("");
    } else if (definition.kind === "MESSAGE") {
      java.push(
        "    public void parse" + definition.name + "(Visitor visitor) {"
      );
      java.push(
        "      visitor.startMessage(MessageDefinition." +
          definition.name.toUpperCase() +
          ");"
      );
      java.push("      for(;;) {");
      java.push("        Long _type = this.readVarUint();");
      java.push(
        '        if (_type == null) throw new IllegalArgumentException("premature end of ' +
          definition.name +
          '");'
      );
      java.push("        switch (_type.intValue()) {");
      java.push("          case 0: {");
      java.push(
        "            visitor.endMessage(MessageDefinition." +
          definition.name.toUpperCase() +
          ");"
      );
      java.push("            return;");
      java.push("          }");
      for (let j = 0; j < definition.fields.length; j++) {
        let field = definition.fields[j];
        java.push("          case " + field.value + ": {");
        java = java.concat(
          emitReadField(definitions, definition, field, "            ")
        );
        java.push("            break;");
        java.push("          }");
      }
      java.push("          default:");
      java.push(
        '            throw new IllegalArgumentException("unrecognized field tag: " + _type);'
      );

      java.push("        }");
      java.push("      }");
      java.push("    }");
      java.push("");
    }
  }

  java.push("  }");
  java.push("}");
  java.push("");

  return java.join("\n");
}
