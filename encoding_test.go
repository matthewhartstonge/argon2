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
	"errors"
	"testing"
)

func Test_Decode_Error(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "valid: argon2i hash",
			input:   "$argon2i$v=19$m=16,t=2,p=1$MmpQV3BIRTVnOVZHOFBISQ$VDE34WMQOVG4X6Sqje0gkg",
			wantErr: nil,
		},
		{
			name:    "valid: argon2d hash",
			input:   "$argon2d$v=19$m=16,t=2,p=1$MmpQV3BIRTVnOVZHOFBISQ$NyuqklQf+S+vzz3AtQCL8w",
			wantErr: nil,
		},
		{
			name:    "valid: argon2id hash",
			input:   "$argon2id$v=19$m=16,t=2,p=1$MmpQV3BIRTVnOVZHOFBISQ$zirDUv1ZjLw0/layHCmWmQ",
			wantErr: nil,
		},
		{
			name:    "invalid hash",
			input:   "$2y$10$B93GqMy3DNkIvyLbsxgtFOG2jwqvatQNUTeh3bPYvcCv9jiQgCO9S",
			wantErr: ErrIncorrectType,
		},
		{
			name:    "invalid hash mode",
			input:   "$argon2e$v=19$m=65536,t=3,p=4$MmpQV3BIRTVnOVZHOFBISQ$zirDUv1ZjLw0/layHCmWmQ",
			wantErr: ErrIncorrectType,
		},
		{
			name:    "empty salt",
			input:   "$argon2id$v=19$m=65536,t=3,p=4$$RdescudvJCsgt3ub+b+dWRWJTmaaJObG",
			wantErr: ErrDecodingFail,
		},
		{
			name:    "empty hash",
			input:   "$argon2id$v=19$m=65536,t=3,p=4$c29tZXNhbHQ$",
			wantErr: ErrDecodingFail,
		},
		{
			name:    "invalid salt",
			input:   "$argon2id$v=19$m=65536,t=3,p=4$!!invalid!!$RdescudvJCsgt3ub+b+dWRWJTmaaJObG",
			wantErr: ErrDecodingFail,
		},
		{
			name:    "invalid hash",
			input:   "$argon2id$v=19$m=65536,t=3,p=4$c29tZXNhbHQ$!!invalid!!",
			wantErr: ErrDecodingFail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decode([]byte(tt.input))
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}

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

func Test_checkMode(t *testing.T) {
	type args struct {
		pa *parser
	}
	tests := []struct {
		name     string
		args     args
		wantMode Mode
		wantErr  bool
	}{
		{
			name: "should error parsing",
			args: args{
				pa: &parser{
					buf: []byte("0$"),
				},
			},
			wantMode: modeArgon2d,
			wantErr:  true,
		},
		{
			name: "should error parsing word",
			args: args{
				pa: &parser{
					buf: []byte("argon2d$"),
				},
			},
			wantMode: modeArgon2d,
			wantErr:  true,
		},
		{
			name: "should parse argon2d mode",
			args: args{
				pa: &parser{
					buf: []byte("d$"),
				},
			},
			wantMode: modeArgon2d,
			wantErr:  false,
		},
		{
			name: "should parse argon2i mode",
			args: args{
				pa: &parser{
					buf: []byte("i$"),
				},
			},
			wantMode: ModeArgon2i,
			wantErr:  false,
		},
		{
			name: "should parse argon2id mode",
			args: args{
				pa: &parser{
					buf: []byte("id$"),
				},
			},
			wantMode: ModeArgon2id,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMode, err := checkMode(tt.args.pa)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkMode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMode != tt.wantMode {
				t.Errorf("checkMode() gotMode = %v, want %v", gotMode, tt.wantMode)
			}
		})
	}
}
