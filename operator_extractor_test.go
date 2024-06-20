package idn_mobile_number_validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProviderFromE164(t *testing.T) {
	type args struct {
		msisdn   string
		withPlus []bool
	}

	tests := []struct {
		name    string
		args    args
		want    Operator
		wantErr error
	}{
		{
			name: "given withPlus is not provided, when the number is invalid, then should return an error",
			args: args{
				msisdn: "123",
			},
			want:    0,
			wantErr: InvalidMSISDN,
		},
		{
			name: "given withPlus is not provided, when the prefix is not found, then should return an error",
			args: args{
				msisdn: "6282820202020",
			},
			want:    0,
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given withPlus is not provided, when the prefix is found, then should return an Operator and a nil error",
			args: args{
				msisdn: "6281820202020",
			},
			want:    XL,
			wantErr: nil,
		},
		{
			name: "given withPlus is provided, when the number is invalid, then should return an error",
			args: args{
				msisdn:   "123",
				withPlus: []bool{true},
			},
			want:    0,
			wantErr: InvalidMSISDN,
		},
		{
			name: "given withPlus is provided, when the prefix is not found, then should return an error",
			args: args{
				msisdn:   "+6282820202020",
				withPlus: []bool{true},
			},
			want:    0,
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given withPlus is provided, when the prefix is found, then should return an Operator and a nil error",
			args: args{
				msisdn:   "+6281820202020",
				withPlus: []bool{true},
			},
			want:    XL,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetProviderFromE164(test.args.msisdn, test.args.withPlus...)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestGetProviderFromNSN(t *testing.T) {
	type args struct {
		msisdn   string
		withZero []bool
	}

	tests := []struct {
		name    string
		args    args
		want    Operator
		wantErr error
	}{
		{
			name: "given withZero is not provided, when the number is invalid, then should return an error",
			args: args{
				msisdn: "123",
			},
			want:    0,
			wantErr: InvalidNSN,
		},
		{
			name: "given withZero is not provided, when the prefix is not found, then should return an error",
			args: args{
				msisdn: "82820202020",
			},
			want:    0,
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given withZero is not provided, when the prefix is found, then should return an Operator and a nil error",
			args: args{
				msisdn: "81820202020",
			},
			want:    XL,
			wantErr: nil,
		},
		{
			name: "given withZero is provided, when the number is invalid, then should return an error",
			args: args{
				msisdn:   "0123",
				withZero: []bool{true},
			},
			want:    0,
			wantErr: InvalidNSN,
		},
		{
			name: "given withZero is provided, when the prefix is not found, then should return an error",
			args: args{
				msisdn:   "082820202020",
				withZero: []bool{true},
			},
			want:    0,
			wantErr: InvalidOperatorPrefix,
		},
		{
			name: "given withZero is provided, when the prefix is found, then should return an Operator and a nil error",
			args: args{
				msisdn:   "081820202020",
				withZero: []bool{true},
			},
			want:    XL,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetProviderFromNSN(test.args.msisdn, test.args.withZero...)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
