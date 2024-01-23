package idn_mobile_number_validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateE164(t *testing.T) {
	type args struct {
		msisdn   string
		withPlus []bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "given without plus, when longer than 14 in length, then should return an error",
			args: args{
				msisdn:   "628120102010201",
				withPlus: nil,
			},
			wantErr: InvalidMSISDN,
		},
		{
			name: "given without plus, when shorter than 11 in length, then should return an error",
			args: args{
				msisdn:   "6281200000",
				withPlus: nil,
			},
			wantErr: InvalidMSISDN,
		},
		{
			name: "given without plus, when prefix is invalid, then should return an error",
			args: args{
				msisdn:   "6282510102020",
				withPlus: nil,
			},
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given without plus, when subscriber number length invalid, then should return an error",
			args: args{
				msisdn:   "62812000000",
				withPlus: nil,
			},
			wantErr: InvalidSubscriberNumberLength,
		},
		{
			name: "given without plus, when is 11 in length, then should return a nil error",
			args: args{
				msisdn:   "62819100100",
				withPlus: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without plus, when is 12 in length, then should return a nil error",
			args: args{
				msisdn:   "628120000012",
				withPlus: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without plus, when is 13 in length, then should return a nil error",
			args: args{
				msisdn:   "6281200000123",
				withPlus: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without plus, when is 14 in length, then should return a nil error",
			args: args{
				msisdn:   "62895000001234",
				withPlus: nil,
			},
			wantErr: nil,
		},
		{
			name: "given with plus, when does not have plus, then should return an error",
			args: args{
				msisdn:   "62812000001",
				withPlus: []bool{true},
			},
			wantErr: InvalidMSISDN,
		},
		{
			name: "given with plus, when longer than 15 in length, then should return an error",
			args: args{
				msisdn:   "+628120102010201",
				withPlus: []bool{true},
			},
			wantErr: InvalidMSISDN,
		},
		{
			name: "given with plus, when shorter than 12 in length, then should return an error",
			args: args{
				msisdn:   "+6281200000",
				withPlus: []bool{true},
			},
			wantErr: InvalidMSISDN,
		},
		{
			name: "given without plus, when prefix is invalid, then should return an error",
			args: args{
				msisdn:   "+6282510102020",
				withPlus: []bool{true},
			},
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given with plus, when subscriber number length invalid, then should return an error",
			args: args{
				msisdn:   "+62812000000",
				withPlus: []bool{true},
			},
			wantErr: InvalidSubscriberNumberLength,
		},
		{
			name: "given with plus, when is 12 in length, then should return a nil error",
			args: args{
				msisdn:   "+62819000001",
				withPlus: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with plus, when is 13 in length, then should return a nil error",
			args: args{
				msisdn:   "+628120000012",
				withPlus: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with plus, when is 14 in length, then should return a nil error",
			args: args{
				msisdn:   "+6281200000123",
				withPlus: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with plus, when is 15 in length, then should return a nil error",
			args: args{
				msisdn:   "+62895000001234",
				withPlus: []bool{true},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantErr, ValidateE164(tt.args.msisdn, tt.args.withPlus...))
		})
	}
}

func TestValidateNSN(t *testing.T) {
	type args struct {
		nsn      string
		withZero []bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "given without zero, when longer than 9 in length, then should return an error",
			args: args{
				nsn:      "8120102010201",
				withZero: nil,
			},
			wantErr: InvalidNSN,
		},
		{
			name: "given without zero, when shorter than 12 in length, then should return an error",
			args: args{
				nsn:      "81200000",
				withZero: nil,
			},
			wantErr: InvalidNSN,
		},
		{
			name: "given without zero, when has invalid prefix, then should return an error",
			args: args{
				nsn:      "825000001",
				withZero: nil,
			},
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given without zero, when has invalid subscriber number length, then should return an error",
			args: args{
				nsn:      "89990009900",
				withZero: nil,
			},
			wantErr: InvalidSubscriberNumberLength,
		},
		{
			name: "given without zero, when is 9 in length, then should return a nil error",
			args: args{
				nsn:      "819000001",
				withZero: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without zero, when is 10 in length, then should return a nil error",
			args: args{
				nsn:      "8120000012",
				withZero: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without zero, when is 11 in length, then should return a nil error",
			args: args{
				nsn:      "81200000123",
				withZero: nil,
			},
			wantErr: nil,
		},
		{
			name: "given without zero, when is 12 in length, then should return a nil error",
			args: args{
				nsn:      "895000001234",
				withZero: nil,
			},
			wantErr: nil,
		},
		{
			name: "given with zero, when does not start from zero, then should return an error",
			args: args{
				nsn:      "812000001",
				withZero: []bool{true},
			},
			wantErr: InvalidNSN,
		},
		{
			name: "given with zero, when longer than 13 in length, then should return an error",
			args: args{
				nsn:      "08120102010201",
				withZero: []bool{true},
			},
			wantErr: InvalidNSN,
		},
		{
			name: "given with zero, when shorter than 10 in length, then should return an error",
			args: args{
				nsn:      "081200000",
				withZero: []bool{true},
			},
			wantErr: InvalidNSN,
		},
		{
			name: "given without zero, when has invalid prefix, then should return an error",
			args: args{
				nsn:      "0825000001",
				withZero: []bool{true},
			},
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given without zero, when has invalid subscriber number length, then should return an error",
			args: args{
				nsn:      "089990009900",
				withZero: []bool{true},
			},
			wantErr: InvalidSubscriberNumberLength,
		},
		{
			name: "given with zero, when is 10 in length, then should return a nil error",
			args: args{
				nsn:      "0819000001",
				withZero: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with zero, when is 11 in length, then should return a nil error",
			args: args{
				nsn:      "08120000012",
				withZero: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with zero, when is 12 in length, then should return a nil error",
			args: args{
				nsn:      "081200000123",
				withZero: []bool{true},
			},
			wantErr: nil,
		},
		{
			name: "given with zero, when is 13 in length, then should return a nil error",
			args: args{
				nsn:      "0895000001234",
				withZero: []bool{true},
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.wantErr, ValidateNSN(test.args.nsn, test.args.withZero...))
		})
	}
}
