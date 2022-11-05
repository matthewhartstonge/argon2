/*
 * Copyright 2022. Matthew Hartstonge <matt@mykro.co.nz>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
