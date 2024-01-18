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
			name: "given without plus, when invalid, then should return an error",
			args: args{
				msisdn:   "628120102010201",
				withPlus: nil,
			},
			wantErr: InvalidMSISDN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantErr, ValidateE164(tt.args.msisdn, tt.args.withPlus...))
		})
	}
}
