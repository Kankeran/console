package console

import "encoding/binary"

type CommandMessage struct {
	Name  string
	Flags map[string][]string
}

func writeString(b []byte, s string) []byte {
	b = binary.BigEndian.AppendUint32(b, uint32(len(s)))
	return append(b, []byte(s)...)
}

func writeStringSlice(b []byte, s []string) []byte {
	b = binary.BigEndian.AppendUint32(b, uint32(len(s)))
	for _, v := range s {
		b = writeString(b, v)
	}
	return b
}

func readString(b []byte) (string, []byte) {
	n := binary.BigEndian.Uint32(b) + 4
	return string(b[4:n]), b[n:]
}

func readStringSlice(b []byte) ([]string, []byte) {
	var n uint32
	n, b = binary.BigEndian.Uint32(b), b[4:]
	s := make([]string, n)
	for i := uint32(0); i < n; i++ {
		s[i], b = readString(b)
	}
	return s, b
}

func MessageFromBytes(data []byte) CommandMessage {
	c := CommandMessage{}
	c.Name, data = readString(data)
	c.Flags = make(map[string][]string)
	var name string
	for len(data) > 0 {
		name, data = readString(data)
		c.Flags[name], data = readStringSlice(data)
	}

	return c
}

func (c CommandMessage) ToBytes() []byte {
	b := make([]byte, 0, 1024)
	b = writeString(b, c.Name)
	for key, values := range c.Flags {
		b = writeString(b, key)
		b = writeStringSlice(b, values)
	}

	return b
}
