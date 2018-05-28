package argon2

import (
	"testing"
)

func Test_parser_parseUint8(t *testing.T) {
	type fields struct {
		buf []byte
		off int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			name: "should parse uint8 zero",
			fields: fields{
				buf: []byte("0"),
				off: 0,
			},
			want: 0,
		},
		{
			name: "should parse uint8 max val",
			fields: fields{
				buf: []byte("255"),
				off: 0,
			},
			want: 255,
		},
		{
			name: "should parse uint8 overflow as zero",
			fields: fields{
				buf: []byte("256"),
				off: 0,
			},
			want: 0,
		},
		{
			name: "should parse uint8 128",
			fields: fields{
				buf: []byte("128"),
				off: 0,
			},
			want: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				buf: tt.fields.buf,
				off: tt.fields.off,
			}
			if got := p.parseUint8(); got != tt.want {
				t.Errorf("parser.parseUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
