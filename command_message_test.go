package console

import "testing"

func byteEqual(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func stringsEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func msgEqual(m1, m2 CommandMessage) bool {
	if m1.Name != m2.Name {
		return false
	}
	if len(m1.Flags) != len(m2.Flags) {
		return false
	}
	for key, values := range m1.Flags {
		if !stringsEqual(values, m2.Flags[key]) {
			return false
		}
	}
	return true
}

func TestWriteString(t *testing.T) {
	b := make([]byte, 0, 3)
	b = writeString(b, "asd")
	expected := []byte{0, 0, 0, 3, 'a', 's', 'd'}
	if !byteEqual(b, expected) {
		t.Errorf("WriteString(\"asd\") = %v; want %v", b, expected)
	}
}

func TestWriteStringSlice(t *testing.T) {
	b := make([]byte, 0, 3)
	b = writeStringSlice(b, []string{"asd", "asd2"})
	expected := []byte{0, 0, 0, 2, 0, 0, 0, 3, 'a', 's', 'd', 0, 0, 0, 4, 'a', 's', 'd', '2'}
	if !byteEqual(b, expected) {
		t.Errorf("WriteStringSlice([]string{\"asd\", \"asd2\"}) = %v; want %v", b, expected)
	}
}

func TestReadString(t *testing.T) {
	b := []byte{0, 0, 0, 3, 'a', 's', 'd'}
	var s string
	s, b = readString(b)
	if len(b) > 0 {
		t.Errorf("ReadString() []byte = %v; want empty", b)
	}
	expected := "asd"
	if s != expected {
		t.Errorf("ReadString() string = %v; want %v", s, expected)
	}
}

func TestReadStringSlice(t *testing.T) {
	b := []byte{0, 0, 0, 2, 0, 0, 0, 3, 'a', 's', 'd', 0, 0, 0, 4, 'a', 's', 'd', '2'}
	var s []string
	s, b = readStringSlice(b)
	if len(b) > 0 {
		t.Errorf("ReadStringSlice() []byte = %v; want empty", b)
	}
	expected := []string{"asd", "asd2"}
	if !stringsEqual(s, expected) {
		t.Errorf("ReadStringSlice() []string = %v; want %v", s, expected)
	}
}

func TestMessage(t *testing.T) {
	msg := CommandMessage{
		Name: "asd",
		Flags: map[string][]string{
			"asd": {
				"asd", "asd2",
			},
			"asd2": {
				"asd3", "asd4",
			},
		},
	}
	b := msg.ToBytes()
	msg2 := MessageFromBytes(b)
	if !msgEqual(msg, msg2) {
		t.Errorf("MessageFromBytes() = %v; want %v", msg2, msg)
	}
}
