package com.google.cloud.teleport.io;

import java.util.HashMap;
import java.util.Map;
public class Schema {
  public static enum Enum {
    A(100),
    B(200);
    private int v;
    private static Map<Integer, Enum> map = new HashMap<>();
    private Enum(int v) {
      this.v = v;
    }
    static {
      for (Enum v : Enum.values()) {
        map.put(v.getValue(), v);
      }
    }
    public static Enum valueOf(int v) {
      return map.get(v);
    }
    public int getValue() {
      return v;
    }
  }

  public static enum StructDefinition {
    ENUMSTRUCT(1),
    BOOLSTRUCT(2),
    BYTESTRUCT(3),
    INTSTRUCT(4),
    UINTSTRUCT(5),
    FLOATSTRUCT(6),
    STRINGSTRUCT(7),
    COMPOUNDSTRUCT(8),
    NESTEDSTRUCT(9),
    ENUMARRAYSTRUCT(10),
    BOOLARRAYSTRUCT(11),
    BYTEARRAYSTRUCT(12),
    INTARRAYSTRUCT(13),
    UINTARRAYSTRUCT(14),
    FLOATARRAYSTRUCT(15),
    STRINGARRAYSTRUCT(16),
    COMPOUNDARRAYSTRUCT(17),
    SORTEDSTRUCT(18);
    private int v;
    private static Map<Integer, StructDefinition> map = new HashMap<>();
    private StructDefinition(int v) {
      this.v = v;
    }
    static {
      for (StructDefinition v : StructDefinition.values()) {
        map.put(v.getValue(), v);
      }
    }
    public static StructDefinition valueOf(int v) {
      return map.get(v);
    }
    public int getValue() {
      return v;
    }
  }

  public static enum MessageDefinition {
    ENUMMESSAGE(1),
    BOOLMESSAGE(2),
    BYTEMESSAGE(3),
    INTMESSAGE(4),
    UINTMESSAGE(5),
    FLOATMESSAGE(6),
    FLOAT32ARRAYMESSAGE(7),
    STRINGMESSAGE(8),
    COMPOUNDMESSAGE(9),
    MAPMESSAGE(10),
    NESTEDMESSAGE(11),
    ENUMARRAYMESSAGE(12),
    BOOLARRAYMESSAGE(13),
    BYTEARRAYMESSAGE(14),
    INTARRAYMESSAGE(15),
    UINTARRAYMESSAGE(16),
    FLOATARRAYMESSAGE(17),
    STRINGARRAYMESSAGE(18),
    COMPOUNDARRAYMESSAGE(19),
    RECURSIVEMESSAGE(20),
    NONDEPRECATEDMESSAGE(21),
    DEPRECATEDMESSAGE(22);
    private int v;
    private static Map<Integer, MessageDefinition> map = new HashMap<>();
    private MessageDefinition(int v) {
      this.v = v;
    }
    static {
      for (MessageDefinition v : MessageDefinition.values()) {
        map.put(v.getValue(), v);
      }
    }
    public static MessageDefinition valueOf(int v) {
      return map.get(v);
    }
    public int getValue() {
      return v;
    }
  }

  public static enum Field {
    ENUMSTRUCT_X(1),
    ENUMSTRUCT_Y(2),
    BOOLSTRUCT_X(3),
    BYTESTRUCT_X(4),
    INTSTRUCT_X(5),
    UINTSTRUCT_X(6),
    FLOATSTRUCT_X(7),
    STRINGSTRUCT_X(8),
    COMPOUNDSTRUCT_X(9),
    COMPOUNDSTRUCT_Y(10),
    NESTEDSTRUCT_A(11),
    NESTEDSTRUCT_B(12),
    NESTEDSTRUCT_C(13),
    NESTEDSTRUCT_D(14),
    ENUMMESSAGE_X(15),
    BOOLMESSAGE_X(16),
    BYTEMESSAGE_X(17),
    INTMESSAGE_X(18),
    UINTMESSAGE_X(19),
    FLOATMESSAGE_X(20),
    FLOAT32ARRAYMESSAGE_X(21),
    STRINGMESSAGE_X(22),
    COMPOUNDMESSAGE_X(23),
    COMPOUNDMESSAGE_Y(24),
    MAPMESSAGE_X(25),
    NESTEDMESSAGE_A(26),
    NESTEDMESSAGE_B(27),
    NESTEDMESSAGE_C(28),
    NESTEDMESSAGE_D(29),
    ENUMARRAYSTRUCT_X(30),
    BOOLARRAYSTRUCT_X(31),
    BYTEARRAYSTRUCT_X(32),
    INTARRAYSTRUCT_X(33),
    UINTARRAYSTRUCT_X(34),
    FLOATARRAYSTRUCT_X(35),
    STRINGARRAYSTRUCT_X(36),
    COMPOUNDARRAYSTRUCT_X(37),
    COMPOUNDARRAYSTRUCT_Y(38),
    ENUMARRAYMESSAGE_X(39),
    BOOLARRAYMESSAGE_X(40),
    BYTEARRAYMESSAGE_X(41),
    INTARRAYMESSAGE_X(42),
    UINTARRAYMESSAGE_X(43),
    FLOATARRAYMESSAGE_X(44),
    STRINGARRAYMESSAGE_X(45),
    COMPOUNDARRAYMESSAGE_X(46),
    COMPOUNDARRAYMESSAGE_Y(47),
    RECURSIVEMESSAGE_X(48),
    NONDEPRECATEDMESSAGE_A(49),
    NONDEPRECATEDMESSAGE_B(50),
    NONDEPRECATEDMESSAGE_C(51),
    NONDEPRECATEDMESSAGE_D(52),
    NONDEPRECATEDMESSAGE_E(53),
    NONDEPRECATEDMESSAGE_F(54),
    NONDEPRECATEDMESSAGE_G(55),
    DEPRECATEDMESSAGE_A(56),
    DEPRECATEDMESSAGE_B(57),
    DEPRECATEDMESSAGE_C(58),
    DEPRECATEDMESSAGE_D(59),
    DEPRECATEDMESSAGE_E(60),
    DEPRECATEDMESSAGE_F(61),
    DEPRECATEDMESSAGE_G(62),
    SORTEDSTRUCT_A1(63),
    SORTEDSTRUCT_B1(64),
    SORTEDSTRUCT_C1(65),
    SORTEDSTRUCT_D1(66),
    SORTEDSTRUCT_E1(67),
    SORTEDSTRUCT_F1(68),
    SORTEDSTRUCT_A2(69),
    SORTEDSTRUCT_B2(70),
    SORTEDSTRUCT_C2(71),
    SORTEDSTRUCT_D2(72),
    SORTEDSTRUCT_E2(73),
    SORTEDSTRUCT_F2(74),
    SORTEDSTRUCT_A3(75),
    SORTEDSTRUCT_B3(76),
    SORTEDSTRUCT_C3(77),
    SORTEDSTRUCT_D3(78),
    SORTEDSTRUCT_E3(79),
    SORTEDSTRUCT_F3(80);
    private int v;
    private static Map<Integer, Field> map = new HashMap<>();
    private Field(int v) {
      this.v = v;
    }
    static {
      for (Field v : Field.values()) {
        map.put(v.getValue(), v);
      }
    }
    public static Field valueOf(int v) {
      return map.get(v);
    }
    public int getValue() {
      return v;
    }
  }

    public interface Visitor {
        void startStruct(StructDefinition def);
        void endStruct(StructDefinition def);

        void startMessage(MessageDefinition def);
        void endMessage(MessageDefinition def);

        void startArray(int len);
        void endArray(int len);

        void startMap(int len);
        void endMap(int len);

        void startMapEntry();
        void endMapEntry();

        void startMapKey();
        void endMapKey();

        void startMapValue();
        void endMapValue();

        void startField(Field f);
        void endField(Field f);

        void boolValue(boolean v);
        void byteValue(byte v);
        void intValue(int v);
        void uintValue(long v);
        void floatValue(float v);
        void float32Value(float v);
        void stringValue(String v);
    }
    public static class VisitorBase implements Visitor {
        public VisitorBase() {}
        public void startStruct(StructDefinition def) { }
        public void endStruct(StructDefinition def) { }

        public void startMessage(MessageDefinition def) { }
        public void endMessage(MessageDefinition def) { }

        public void startArray(int len) { }
        public void endArray(int len) { }

        public void startMap(int len) { }
        public void endMap(int len) { }

        public void startMapEntry() { }
        public void endMapEntry() { }

        public void startMapKey() { }
        public void endMapKey() { }

        public void startMapValue() { }
        public void endMapValue() { }

        public void startField(Field f) { }
        public void endField(Field f) { }

        public void boolValue(boolean v) { }
        public void byteValue(byte v) { }
        public void intValue(int v) { }
        public void uintValue(long v) { }
        public void floatValue(float v) { }
        public void float32Value(float v) { }
        public void stringValue(String v) { }
    }

public static class Parser extends ByteBuffer {
public Parser(byte[] buf) {
  super(buf);
}
public void parseEnumStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.ENUMSTRUCT);
  visitor.startField(Field.ENUMSTRUCT_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x);
}
  visitor.endField(Field.ENUMSTRUCT_X);
  visitor.startField(Field.ENUMSTRUCT_Y);
  Long _y_count = this.readVarUint();
  if (_y_count == null) throw new IllegalArgumentException("missing array count for y");
  visitor.startArray(_y_count.intValue());
  for (long j = 0; j < _y_count; j++) {
Long y_element = this.readVarUint(); if (y_element == null) throw new IllegalArgumentException("reading y");
  visitor.uintValue(y_element);
  }
  visitor.endArray(_y_count.intValue());
  visitor.endField(Field.ENUMSTRUCT_Y);
  visitor.endStruct(StructDefinition.ENUMSTRUCT);
}

public void parseBoolStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.BOOLSTRUCT);
  visitor.startField(Field.BOOLSTRUCT_X);
{
Boolean x = this.readBool(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.boolValue(x);
}
  visitor.endField(Field.BOOLSTRUCT_X);
  visitor.endStruct(StructDefinition.BOOLSTRUCT);
}

public void parseByteStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.BYTESTRUCT);
  visitor.startField(Field.BYTESTRUCT_X);
{
Byte x = this.readByte(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.byteValue(x);
}
  visitor.endField(Field.BYTESTRUCT_X);
  visitor.endStruct(StructDefinition.BYTESTRUCT);
}

public void parseIntStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.INTSTRUCT);
  visitor.startField(Field.INTSTRUCT_X);
{
Integer x = this.readVarInt(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.intValue(x);
}
  visitor.endField(Field.INTSTRUCT_X);
  visitor.endStruct(StructDefinition.INTSTRUCT);
}

public void parseUintStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.UINTSTRUCT);
  visitor.startField(Field.UINTSTRUCT_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x);
}
  visitor.endField(Field.UINTSTRUCT_X);
  visitor.endStruct(StructDefinition.UINTSTRUCT);
}

public void parseFloatStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.FLOATSTRUCT);
  visitor.startField(Field.FLOATSTRUCT_X);
{
Float x = this.readVarFloat(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.floatValue(x);
}
  visitor.endField(Field.FLOATSTRUCT_X);
  visitor.endStruct(StructDefinition.FLOATSTRUCT);
}

public void parseStringStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.STRINGSTRUCT);
  visitor.startField(Field.STRINGSTRUCT_X);
{
String x = this.readString(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.stringValue(x);
}
  visitor.endField(Field.STRINGSTRUCT_X);
  visitor.endStruct(StructDefinition.STRINGSTRUCT);
}

public void parseCompoundStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.COMPOUNDSTRUCT);
  visitor.startField(Field.COMPOUNDSTRUCT_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x);
}
  visitor.endField(Field.COMPOUNDSTRUCT_X);
  visitor.startField(Field.COMPOUNDSTRUCT_Y);
{
Long y = this.readVarUint(); if (y == null) throw new IllegalArgumentException("reading y");
  visitor.uintValue(y);
}
  visitor.endField(Field.COMPOUNDSTRUCT_Y);
  visitor.endStruct(StructDefinition.COMPOUNDSTRUCT);
}

public void parseNestedStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.NESTEDSTRUCT);
  visitor.startField(Field.NESTEDSTRUCT_A);
{
Long a = this.readVarUint(); if (a == null) throw new IllegalArgumentException("reading a");
  visitor.uintValue(a);
}
  visitor.endField(Field.NESTEDSTRUCT_A);
  visitor.startField(Field.NESTEDSTRUCT_B);
{
parseCompoundStruct(visitor);
  
}
  visitor.endField(Field.NESTEDSTRUCT_B);
  visitor.startField(Field.NESTEDSTRUCT_C);
{
Long c = this.readVarUint(); if (c == null) throw new IllegalArgumentException("reading c");
  visitor.uintValue(c);
}
  visitor.endField(Field.NESTEDSTRUCT_C);
  visitor.startField(Field.NESTEDSTRUCT_D);
  Long _d_count = this.readVarUint();
  if (_d_count == null) throw new IllegalArgumentException("missing map entry count for d");
  visitor.startMap(_d_count.intValue());
  for (long j = 0; j < _d_count; j++) {
  visitor.startMapEntry();
  visitor.startMapKey();
String key = this.readString(); if (key == null) throw new IllegalArgumentException("reading d.key");
  visitor.stringValue(key);
  visitor.endMapKey();
  visitor.startMapValue();
parseCompoundStruct(visitor);
  
  visitor.endMapValue();
  visitor.endMapEntry();
  }
  visitor.endMap(_d_count.intValue());
  visitor.endField(Field.NESTEDSTRUCT_D);
  visitor.endStruct(StructDefinition.NESTEDSTRUCT);
}

public void parseEnumMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.ENUMMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of EnumMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.ENUMMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.ENUMMESSAGE_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x);
}
        visitor.endField(Field.ENUMMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseBoolMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.BOOLMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of BoolMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.BOOLMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.BOOLMESSAGE_X);
{
Boolean x = this.readBool(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.boolValue(x);
}
        visitor.endField(Field.BOOLMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseByteMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.BYTEMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of ByteMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.BYTEMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.BYTEMESSAGE_X);
{
Byte x = this.readByte(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.byteValue(x);
}
        visitor.endField(Field.BYTEMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseIntMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.INTMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of IntMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.INTMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.INTMESSAGE_X);
{
Integer x = this.readVarInt(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.intValue(x);
}
        visitor.endField(Field.INTMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseUintMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.UINTMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of UintMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.UINTMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.UINTMESSAGE_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x);
}
        visitor.endField(Field.UINTMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseFloatMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.FLOATMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of FloatMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.FLOATMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.FLOATMESSAGE_X);
{
Float x = this.readVarFloat(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.floatValue(x);
}
        visitor.endField(Field.FLOATMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseFloat32ArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.FLOAT32ARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of Float32ArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.FLOAT32ARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.FLOAT32ARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Float x_element = this.readVarFloat(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.float32Value(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.FLOAT32ARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseStringMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.STRINGMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of StringMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.STRINGMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.STRINGMESSAGE_X);
{
String x = this.readString(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.stringValue(x);
}
        visitor.endField(Field.STRINGMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseCompoundMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.COMPOUNDMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of CompoundMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.COMPOUNDMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.COMPOUNDMESSAGE_X);
{
Long x = this.readVarUint(); if (x == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x);
}
        visitor.endField(Field.COMPOUNDMESSAGE_X);
        break;
      }
      case 2: {
        visitor.startField(Field.COMPOUNDMESSAGE_Y);
{
Long y = this.readVarUint(); if (y == null) throw new IllegalArgumentException("reading y");
        visitor.uintValue(y);
}
        visitor.endField(Field.COMPOUNDMESSAGE_Y);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseMapMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.MAPMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of MapMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.MAPMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.MAPMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing map entry count for x");
        visitor.startMap(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
        visitor.startMapEntry();
        visitor.startMapKey();
String key = this.readString(); if (key == null) throw new IllegalArgumentException("reading x.key");
        visitor.stringValue(key);
        visitor.endMapKey();
        visitor.startMapValue();
Integer value = this.readVarInt(); if (value == null) throw new IllegalArgumentException("reading x.value");
        visitor.intValue(value);
        visitor.endMapValue();
        visitor.endMapEntry();
        }
        visitor.endMap(_x_count.intValue());
        visitor.endField(Field.MAPMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseNestedMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.NESTEDMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of NestedMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.NESTEDMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.NESTEDMESSAGE_A);
{
Long a = this.readVarUint(); if (a == null) throw new IllegalArgumentException("reading a");
        visitor.uintValue(a);
}
        visitor.endField(Field.NESTEDMESSAGE_A);
        break;
      }
      case 2: {
        visitor.startField(Field.NESTEDMESSAGE_B);
{
parseCompoundMessage(visitor);
        
}
        visitor.endField(Field.NESTEDMESSAGE_B);
        break;
      }
      case 3: {
        visitor.startField(Field.NESTEDMESSAGE_C);
{
Long c = this.readVarUint(); if (c == null) throw new IllegalArgumentException("reading c");
        visitor.uintValue(c);
}
        visitor.endField(Field.NESTEDMESSAGE_C);
        break;
      }
      case 4: {
        visitor.startField(Field.NESTEDMESSAGE_D);
        Long _d_count = this.readVarUint();
        if (_d_count == null) throw new IllegalArgumentException("missing map entry count for d");
        visitor.startMap(_d_count.intValue());
        for (long j = 0; j < _d_count; j++) {
        visitor.startMapEntry();
        visitor.startMapKey();
String key = this.readString(); if (key == null) throw new IllegalArgumentException("reading d.key");
        visitor.stringValue(key);
        visitor.endMapKey();
        visitor.startMapValue();
parseCompoundMessage(visitor);
        
        visitor.endMapValue();
        visitor.endMapEntry();
        }
        visitor.endMap(_d_count.intValue());
        visitor.endField(Field.NESTEDMESSAGE_D);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseEnumArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.ENUMARRAYSTRUCT);
  visitor.startField(Field.ENUMARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.ENUMARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.ENUMARRAYSTRUCT);
}

public void parseBoolArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.BOOLARRAYSTRUCT);
  visitor.startField(Field.BOOLARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Boolean x_element = this.readBool(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.boolValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.BOOLARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.BOOLARRAYSTRUCT);
}

public void parseByteArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.BYTEARRAYSTRUCT);
  visitor.startField(Field.BYTEARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Byte x_element = this.readByte(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.byteValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.BYTEARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.BYTEARRAYSTRUCT);
}

public void parseIntArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.INTARRAYSTRUCT);
  visitor.startField(Field.INTARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Integer x_element = this.readVarInt(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.intValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.INTARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.INTARRAYSTRUCT);
}

public void parseUintArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.UINTARRAYSTRUCT);
  visitor.startField(Field.UINTARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.UINTARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.UINTARRAYSTRUCT);
}

public void parseFloatArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.FLOATARRAYSTRUCT);
  visitor.startField(Field.FLOATARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Float x_element = this.readVarFloat(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.floatValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.FLOATARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.FLOATARRAYSTRUCT);
}

public void parseStringArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.STRINGARRAYSTRUCT);
  visitor.startField(Field.STRINGARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
String x_element = this.readString(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.stringValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.STRINGARRAYSTRUCT_X);
  visitor.endStruct(StructDefinition.STRINGARRAYSTRUCT);
}

public void parseCompoundArrayStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.COMPOUNDARRAYSTRUCT);
  visitor.startField(Field.COMPOUNDARRAYSTRUCT_X);
  Long _x_count = this.readVarUint();
  if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
  visitor.startArray(_x_count.intValue());
  for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
  visitor.uintValue(x_element);
  }
  visitor.endArray(_x_count.intValue());
  visitor.endField(Field.COMPOUNDARRAYSTRUCT_X);
  visitor.startField(Field.COMPOUNDARRAYSTRUCT_Y);
  Long _y_count = this.readVarUint();
  if (_y_count == null) throw new IllegalArgumentException("missing array count for y");
  visitor.startArray(_y_count.intValue());
  for (long j = 0; j < _y_count; j++) {
Long y_element = this.readVarUint(); if (y_element == null) throw new IllegalArgumentException("reading y");
  visitor.uintValue(y_element);
  }
  visitor.endArray(_y_count.intValue());
  visitor.endField(Field.COMPOUNDARRAYSTRUCT_Y);
  visitor.endStruct(StructDefinition.COMPOUNDARRAYSTRUCT);
}

public void parseEnumArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.ENUMARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of EnumArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.ENUMARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.ENUMARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.ENUMARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseBoolArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.BOOLARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of BoolArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.BOOLARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.BOOLARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Boolean x_element = this.readBool(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.boolValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.BOOLARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseByteArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.BYTEARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of ByteArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.BYTEARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.BYTEARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Byte x_element = this.readByte(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.byteValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.BYTEARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseIntArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.INTARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of IntArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.INTARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.INTARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Integer x_element = this.readVarInt(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.intValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.INTARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseUintArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.UINTARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of UintArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.UINTARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.UINTARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.UINTARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseFloatArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.FLOATARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of FloatArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.FLOATARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.FLOATARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Float x_element = this.readVarFloat(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.floatValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.FLOATARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseStringArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.STRINGARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of StringArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.STRINGARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.STRINGARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
String x_element = this.readString(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.stringValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.STRINGARRAYMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseCompoundArrayMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.COMPOUNDARRAYMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of CompoundArrayMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.COMPOUNDARRAYMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.COMPOUNDARRAYMESSAGE_X);
        Long _x_count = this.readVarUint();
        if (_x_count == null) throw new IllegalArgumentException("missing array count for x");
        visitor.startArray(_x_count.intValue());
        for (long j = 0; j < _x_count; j++) {
Long x_element = this.readVarUint(); if (x_element == null) throw new IllegalArgumentException("reading x");
        visitor.uintValue(x_element);
        }
        visitor.endArray(_x_count.intValue());
        visitor.endField(Field.COMPOUNDARRAYMESSAGE_X);
        break;
      }
      case 2: {
        visitor.startField(Field.COMPOUNDARRAYMESSAGE_Y);
        Long _y_count = this.readVarUint();
        if (_y_count == null) throw new IllegalArgumentException("missing array count for y");
        visitor.startArray(_y_count.intValue());
        for (long j = 0; j < _y_count; j++) {
Long y_element = this.readVarUint(); if (y_element == null) throw new IllegalArgumentException("reading y");
        visitor.uintValue(y_element);
        }
        visitor.endArray(_y_count.intValue());
        visitor.endField(Field.COMPOUNDARRAYMESSAGE_Y);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseRecursiveMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.RECURSIVEMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of RecursiveMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.RECURSIVEMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.RECURSIVEMESSAGE_X);
{
parseRecursiveMessage(visitor);
        
}
        visitor.endField(Field.RECURSIVEMESSAGE_X);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseNonDeprecatedMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.NONDEPRECATEDMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of NonDeprecatedMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.NONDEPRECATEDMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_A);
{
Long a = this.readVarUint(); if (a == null) throw new IllegalArgumentException("reading a");
        visitor.uintValue(a);
}
        visitor.endField(Field.NONDEPRECATEDMESSAGE_A);
        break;
      }
      case 2: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_B);
{
Long b = this.readVarUint(); if (b == null) throw new IllegalArgumentException("reading b");
        visitor.uintValue(b);
}
        visitor.endField(Field.NONDEPRECATEDMESSAGE_B);
        break;
      }
      case 3: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_C);
        Long _c_count = this.readVarUint();
        if (_c_count == null) throw new IllegalArgumentException("missing array count for c");
        visitor.startArray(_c_count.intValue());
        for (long j = 0; j < _c_count; j++) {
Long c_element = this.readVarUint(); if (c_element == null) throw new IllegalArgumentException("reading c");
        visitor.uintValue(c_element);
        }
        visitor.endArray(_c_count.intValue());
        visitor.endField(Field.NONDEPRECATEDMESSAGE_C);
        break;
      }
      case 4: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_D);
        Long _d_count = this.readVarUint();
        if (_d_count == null) throw new IllegalArgumentException("missing array count for d");
        visitor.startArray(_d_count.intValue());
        for (long j = 0; j < _d_count; j++) {
Long d_element = this.readVarUint(); if (d_element == null) throw new IllegalArgumentException("reading d");
        visitor.uintValue(d_element);
        }
        visitor.endArray(_d_count.intValue());
        visitor.endField(Field.NONDEPRECATEDMESSAGE_D);
        break;
      }
      case 5: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_E);
{
parseByteStruct(visitor);
        
}
        visitor.endField(Field.NONDEPRECATEDMESSAGE_E);
        break;
      }
      case 6: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_F);
{
parseByteStruct(visitor);
        
}
        visitor.endField(Field.NONDEPRECATEDMESSAGE_F);
        break;
      }
      case 7: {
        visitor.startField(Field.NONDEPRECATEDMESSAGE_G);
{
Long g = this.readVarUint(); if (g == null) throw new IllegalArgumentException("reading g");
        visitor.uintValue(g);
}
        visitor.endField(Field.NONDEPRECATEDMESSAGE_G);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseDeprecatedMessage(Visitor visitor) {
  visitor.startMessage(MessageDefinition.DEPRECATEDMESSAGE);
  for(;;) {
    Long _type = this.readVarUint();
    if (_type == null) throw new IllegalArgumentException("premature end of DeprecatedMessage");
    switch (_type.intValue()) {
      case 0: {
        visitor.endMessage(MessageDefinition.DEPRECATEDMESSAGE);
        return;
      }
      case 1: {
        visitor.startField(Field.DEPRECATEDMESSAGE_A);
{
Long a = this.readVarUint(); if (a == null) throw new IllegalArgumentException("reading a");
        visitor.uintValue(a);
}
        visitor.endField(Field.DEPRECATEDMESSAGE_A);
        break;
      }
      case 2: {
        visitor.startField(Field.DEPRECATEDMESSAGE_B);
{
Long b = this.readVarUint(); if (b == null) throw new IllegalArgumentException("reading b");
        visitor.uintValue(b);
}
        visitor.endField(Field.DEPRECATEDMESSAGE_B);
        break;
      }
      case 3: {
        visitor.startField(Field.DEPRECATEDMESSAGE_C);
        Long _c_count = this.readVarUint();
        if (_c_count == null) throw new IllegalArgumentException("missing array count for c");
        visitor.startArray(_c_count.intValue());
        for (long j = 0; j < _c_count; j++) {
Long c_element = this.readVarUint(); if (c_element == null) throw new IllegalArgumentException("reading c");
        visitor.uintValue(c_element);
        }
        visitor.endArray(_c_count.intValue());
        visitor.endField(Field.DEPRECATEDMESSAGE_C);
        break;
      }
      case 4: {
        visitor.startField(Field.DEPRECATEDMESSAGE_D);
        Long _d_count = this.readVarUint();
        if (_d_count == null) throw new IllegalArgumentException("missing array count for d");
        visitor.startArray(_d_count.intValue());
        for (long j = 0; j < _d_count; j++) {
Long d_element = this.readVarUint(); if (d_element == null) throw new IllegalArgumentException("reading d");
        visitor.uintValue(d_element);
        }
        visitor.endArray(_d_count.intValue());
        visitor.endField(Field.DEPRECATEDMESSAGE_D);
        break;
      }
      case 5: {
        visitor.startField(Field.DEPRECATEDMESSAGE_E);
{
parseByteStruct(visitor);
        
}
        visitor.endField(Field.DEPRECATEDMESSAGE_E);
        break;
      }
      case 6: {
        visitor.startField(Field.DEPRECATEDMESSAGE_F);
{
parseByteStruct(visitor);
        
}
        visitor.endField(Field.DEPRECATEDMESSAGE_F);
        break;
      }
      case 7: {
        visitor.startField(Field.DEPRECATEDMESSAGE_G);
{
Long g = this.readVarUint(); if (g == null) throw new IllegalArgumentException("reading g");
        visitor.uintValue(g);
}
        visitor.endField(Field.DEPRECATEDMESSAGE_G);
        break;
      }
      default: throw new IllegalArgumentException("unrecognized field tag: " + _type);
    }
  }
}

public void parseSortedStruct(Visitor visitor) {
  visitor.startStruct(StructDefinition.SORTEDSTRUCT);
  visitor.startField(Field.SORTEDSTRUCT_A1);
{
Boolean a1 = this.readBool(); if (a1 == null) throw new IllegalArgumentException("reading a1");
  visitor.boolValue(a1);
}
  visitor.endField(Field.SORTEDSTRUCT_A1);
  visitor.startField(Field.SORTEDSTRUCT_B1);
{
Byte b1 = this.readByte(); if (b1 == null) throw new IllegalArgumentException("reading b1");
  visitor.byteValue(b1);
}
  visitor.endField(Field.SORTEDSTRUCT_B1);
  visitor.startField(Field.SORTEDSTRUCT_C1);
{
Integer c1 = this.readVarInt(); if (c1 == null) throw new IllegalArgumentException("reading c1");
  visitor.intValue(c1);
}
  visitor.endField(Field.SORTEDSTRUCT_C1);
  visitor.startField(Field.SORTEDSTRUCT_D1);
{
Long d1 = this.readVarUint(); if (d1 == null) throw new IllegalArgumentException("reading d1");
  visitor.uintValue(d1);
}
  visitor.endField(Field.SORTEDSTRUCT_D1);
  visitor.startField(Field.SORTEDSTRUCT_E1);
{
Float e1 = this.readVarFloat(); if (e1 == null) throw new IllegalArgumentException("reading e1");
  visitor.floatValue(e1);
}
  visitor.endField(Field.SORTEDSTRUCT_E1);
  visitor.startField(Field.SORTEDSTRUCT_F1);
{
String f1 = this.readString(); if (f1 == null) throw new IllegalArgumentException("reading f1");
  visitor.stringValue(f1);
}
  visitor.endField(Field.SORTEDSTRUCT_F1);
  visitor.startField(Field.SORTEDSTRUCT_A2);
{
Boolean a2 = this.readBool(); if (a2 == null) throw new IllegalArgumentException("reading a2");
  visitor.boolValue(a2);
}
  visitor.endField(Field.SORTEDSTRUCT_A2);
  visitor.startField(Field.SORTEDSTRUCT_B2);
{
Byte b2 = this.readByte(); if (b2 == null) throw new IllegalArgumentException("reading b2");
  visitor.byteValue(b2);
}
  visitor.endField(Field.SORTEDSTRUCT_B2);
  visitor.startField(Field.SORTEDSTRUCT_C2);
{
Integer c2 = this.readVarInt(); if (c2 == null) throw new IllegalArgumentException("reading c2");
  visitor.intValue(c2);
}
  visitor.endField(Field.SORTEDSTRUCT_C2);
  visitor.startField(Field.SORTEDSTRUCT_D2);
{
Long d2 = this.readVarUint(); if (d2 == null) throw new IllegalArgumentException("reading d2");
  visitor.uintValue(d2);
}
  visitor.endField(Field.SORTEDSTRUCT_D2);
  visitor.startField(Field.SORTEDSTRUCT_E2);
{
Float e2 = this.readVarFloat(); if (e2 == null) throw new IllegalArgumentException("reading e2");
  visitor.floatValue(e2);
}
  visitor.endField(Field.SORTEDSTRUCT_E2);
  visitor.startField(Field.SORTEDSTRUCT_F2);
{
String f2 = this.readString(); if (f2 == null) throw new IllegalArgumentException("reading f2");
  visitor.stringValue(f2);
}
  visitor.endField(Field.SORTEDSTRUCT_F2);
  visitor.startField(Field.SORTEDSTRUCT_A3);
  Long _a3_count = this.readVarUint();
  if (_a3_count == null) throw new IllegalArgumentException("missing array count for a3");
  visitor.startArray(_a3_count.intValue());
  for (long j = 0; j < _a3_count; j++) {
Boolean a3_element = this.readBool(); if (a3_element == null) throw new IllegalArgumentException("reading a3");
  visitor.boolValue(a3_element);
  }
  visitor.endArray(_a3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_A3);
  visitor.startField(Field.SORTEDSTRUCT_B3);
  Long _b3_count = this.readVarUint();
  if (_b3_count == null) throw new IllegalArgumentException("missing array count for b3");
  visitor.startArray(_b3_count.intValue());
  for (long j = 0; j < _b3_count; j++) {
Byte b3_element = this.readByte(); if (b3_element == null) throw new IllegalArgumentException("reading b3");
  visitor.byteValue(b3_element);
  }
  visitor.endArray(_b3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_B3);
  visitor.startField(Field.SORTEDSTRUCT_C3);
  Long _c3_count = this.readVarUint();
  if (_c3_count == null) throw new IllegalArgumentException("missing array count for c3");
  visitor.startArray(_c3_count.intValue());
  for (long j = 0; j < _c3_count; j++) {
Integer c3_element = this.readVarInt(); if (c3_element == null) throw new IllegalArgumentException("reading c3");
  visitor.intValue(c3_element);
  }
  visitor.endArray(_c3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_C3);
  visitor.startField(Field.SORTEDSTRUCT_D3);
  Long _d3_count = this.readVarUint();
  if (_d3_count == null) throw new IllegalArgumentException("missing array count for d3");
  visitor.startArray(_d3_count.intValue());
  for (long j = 0; j < _d3_count; j++) {
Long d3_element = this.readVarUint(); if (d3_element == null) throw new IllegalArgumentException("reading d3");
  visitor.uintValue(d3_element);
  }
  visitor.endArray(_d3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_D3);
  visitor.startField(Field.SORTEDSTRUCT_E3);
  Long _e3_count = this.readVarUint();
  if (_e3_count == null) throw new IllegalArgumentException("missing array count for e3");
  visitor.startArray(_e3_count.intValue());
  for (long j = 0; j < _e3_count; j++) {
Float e3_element = this.readVarFloat(); if (e3_element == null) throw new IllegalArgumentException("reading e3");
  visitor.floatValue(e3_element);
  }
  visitor.endArray(_e3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_E3);
  visitor.startField(Field.SORTEDSTRUCT_F3);
  Long _f3_count = this.readVarUint();
  if (_f3_count == null) throw new IllegalArgumentException("missing array count for f3");
  visitor.startArray(_f3_count.intValue());
  for (long j = 0; j < _f3_count; j++) {
String f3_element = this.readString(); if (f3_element == null) throw new IllegalArgumentException("reading f3");
  visitor.stringValue(f3_element);
  }
  visitor.endArray(_f3_count.intValue());
  visitor.endField(Field.SORTEDSTRUCT_F3);
  visitor.endStruct(StructDefinition.SORTEDSTRUCT);
}

}
}
