package main;

import org.junit.jupiter.api.Test;

import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicReference;

import static org.junit.jupiter.api.Assertions.*;

class KiwiTest {

    @Test
    void parseStruct() {
        byte[] buf = new byte[] { (byte)150, 4, (byte)185, 96, (byte)177, (byte)168, 3, (byte)193, 2, 0 };
        Schema.Parser p = new Schema.Parser(buf);
        final AtomicInteger calledStartStruct = new AtomicInteger(0);
        final AtomicInteger calledA = new AtomicInteger(0);
        final AtomicInteger calledB = new AtomicInteger(0);
        final AtomicInteger calledC = new AtomicInteger(0);
        final AtomicInteger calledD = new AtomicInteger(0);
        final AtomicInteger calledMapStart = new AtomicInteger(0);
        final AtomicInteger calledMapEnd = new AtomicInteger(0);
        final AtomicInteger setAwithUint = new AtomicInteger(0);

        p.parseNestedStruct(new Schema.VisitorBase() {
            private final AtomicReference<Schema.Field> field = new AtomicReference<>();

            @Override
            public void startStruct(Schema.StructDefinition def) {
                calledStartStruct.incrementAndGet();
            }

            @Override
            public void startMap(int len) {
                calledMapStart.incrementAndGet();
            }

            @Override
            public void endMap(int len) {
                calledMapEnd.incrementAndGet();
            }

            @Override
            public void startField(Schema.Field f) {
                field.set(f);
                switch(f) {
                    case NESTEDSTRUCT_A:
                        calledA.incrementAndGet();
                        break;
                    case NESTEDSTRUCT_B:
                        calledB.incrementAndGet();
                        break;
                    case NESTEDSTRUCT_C:
                        calledC.incrementAndGet();
                        break;
                    case NESTEDSTRUCT_D:
                        calledD.incrementAndGet();
                        break;
                }
            }

            @Override
            public void endField(Schema.Field f) {
                field.set(null);
            }

            @Override
            public void uintValue(long v) {
                if (Schema.Field.NESTEDSTRUCT_A.equals(field.get())) {
                    setAwithUint.set((int) v);
                }
            }
        });
        assertTrue(p.isEmpty());
        assertEquals(1, calledA.get());
        assertEquals(1, calledB.get());
        assertEquals(1, calledC.get());
        assertEquals(1, calledD.get());
        assertEquals(1, calledMapStart.get());
        assertEquals(1, calledMapEnd.get());
        assertEquals(534, setAwithUint.get());
    };

    @Test
    void parseMessage() {
        byte[] buf = new byte[] { 1, (byte) 234, 1, 2, 1, 5, 2, 6, 0, 3, 123, 4, 1, 97, 0, 1, 1, 2, 2, 0, 0 };
        final AtomicInteger calledStartMessage = new AtomicInteger(0);
        final AtomicInteger calledA = new AtomicInteger(0);
        final AtomicInteger calledB = new AtomicInteger(0);
        final AtomicInteger calledC = new AtomicInteger(0);
        final AtomicInteger calledD = new AtomicInteger(0);
        final AtomicInteger calledMapStart = new AtomicInteger(0);
        final AtomicInteger calledMapEnd = new AtomicInteger(0);
        final AtomicInteger calledMapEntryStart = new AtomicInteger(0);
        final AtomicInteger calledMapEntryEnd = new AtomicInteger(0);
        final AtomicInteger calledMapKeyStart = new AtomicInteger(0);
        final AtomicInteger calledMapKeyEnd = new AtomicInteger(0);
        final AtomicInteger calledMapValueStart = new AtomicInteger(0);
        final AtomicInteger calledMapValueEnd = new AtomicInteger(0);
        final AtomicReference<String> calledMapValueString = new AtomicReference<>();
        final AtomicInteger calledCompoundX = new AtomicInteger(0);
        final AtomicInteger calledCompoundY = new AtomicInteger(0);
        final AtomicInteger setAwithUint = new AtomicInteger(0);

        new Schema.Parser(buf).parseNestedMessage(new Schema.Visitor() {
            public void startStruct(Schema.StructDefinition def) {
                System.out.println("start struct " + def);
            }

            public void endStruct(Schema.StructDefinition def) {
                System.out.println("end struct " + def);
            }

            public void startMessage(Schema.MessageDefinition def) {
                System.out.println("start message " + def);
            }

            public void endMessage(Schema.MessageDefinition def) {
                System.out.println("end message " + def);
            }

            public void startArray(int len) {
                System.out.println("start array with len " + len);
            }

            public void endArray(int len) {
                System.out.println("end array with len " + len);
            }

            public void startMap(int len) {
                System.out.println("start map with len " + len);
            }

            public void endMap(int len) {
                System.out.println("end map with len " + len);
            }

            public void startMapEntry() {
                System.out.println("start map entry");
            }

            public void endMapEntry() {
                System.out.println("end map entry");
            }

            public void startMapKey() {
                System.out.println("start map key");
            }

            public void endMapKey() {
                System.out.println("end map key");
            }

            public void startMapValue() {
                System.out.println("start map value");
            }

            public void endMapValue() {
                System.out.println("end map value");
            }

            public void startField(Schema.Field f) {
                System.out.println("start field " + f);
            }

            public void endField(Schema.Field f) {
                System.out.println("end field " + f);
            }

            public void boolValue(boolean v) {
                System.out.println("bool value " + v);
            }

            public void byteValue(byte v) {
                System.out.println("byte value " + v);
            }

            public void intValue(int v) {
                System.out.println("int value " + v);
            }

            public void uintValue(long v) {
                System.out.println("uint value " + v);
            }

            public void floatValue(float v) {
                System.out.println("float value " + v);
            }

            public void float32Value(float v) {
                System.out.println("float32 value " + v);
            }

            public void stringValue(String v) {
                System.out.println("string value " + v);
            }
        });
        new Schema.Parser(buf).parseNestedMessage(new Schema.VisitorBase() {
            private final AtomicReference<Schema.Field> field = new AtomicReference<>();

            @Override
            public void startMessage(Schema.MessageDefinition def) {
                calledStartMessage.incrementAndGet();
            }

            @Override
            public void startMap(int len) {
                calledMapStart.incrementAndGet();
            }

            @Override
            public void endMap(int len) {
                calledMapEnd.incrementAndGet();
            }

            @Override
            public void startMapEntry() {
                calledMapEntryStart.incrementAndGet();
            }

            @Override
            public void endMapEntry() {
                calledMapEntryEnd.incrementAndGet();
            }

            @Override
            public void startField(Schema.Field f) {
                field.set(f);
                switch(f) {
                    case NESTEDMESSAGE_A:
                        calledA.incrementAndGet();
                        break;
                    case NESTEDMESSAGE_B:
                        calledB.incrementAndGet();
                        break;
                    case NESTEDMESSAGE_C:
                        calledC.incrementAndGet();
                        break;
                    case NESTEDMESSAGE_D:
                        calledD.incrementAndGet();
                        break;

                    case COMPOUNDMESSAGE_X:
                        calledCompoundX.incrementAndGet();
                        break;
                    case COMPOUNDMESSAGE_Y:
                        calledCompoundY.incrementAndGet();
                        break;

                    case NESTEDMESSAGE_D_MAP_KEY:
                        calledMapKeyStart.incrementAndGet();
                        break;
                    case NESTEDMESSAGE_D_MAP_VALUE:
                        calledMapValueStart.incrementAndGet();
                        break;
                }
            }

            @Override
            public void endField(Schema.Field f) {
                field.set(null);
                switch (f) {
                    case NESTEDMESSAGE_D_MAP_KEY:
                        calledMapKeyEnd.incrementAndGet();
                        break;
                    case NESTEDMESSAGE_D_MAP_VALUE:
                        calledMapValueEnd.incrementAndGet();
                        break;
                }
            }

            @Override
            public void stringValue(String v) {
                if (field.get() == Schema.Field.NESTEDMESSAGE_D_MAP_KEY) {
                    calledMapValueString.set(v);
                }
            }

            @Override
            public void uintValue(long v) {
                if (field.get() == Schema.Field.NESTEDMESSAGE_A) {
                    setAwithUint.set((int) v);
                }
            }
        });
        assertEquals(1, calledA.get());
        assertEquals(1, calledB.get());
        assertEquals(1, calledC.get());
        assertEquals(1, calledD.get());
        assertEquals(1, calledMapStart.get());
        assertEquals(1, calledMapEnd.get());

        assertEquals(1, calledMapEntryStart.get());
        assertEquals(1, calledMapEntryEnd.get());

        assertEquals(1, calledMapKeyStart.get());
        assertEquals(1, calledMapKeyEnd.get());

        assertEquals(1, calledMapValueStart.get());
        assertEquals(1, calledMapValueEnd.get());

        assertEquals(234, setAwithUint.get());
        assertEquals("a", calledMapValueString.get());
        assertEquals(2, calledCompoundX.get());
        assertEquals(2, calledCompoundY.get());
    };
}
