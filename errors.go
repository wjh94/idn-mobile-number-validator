package idn_mobile_number_validator

import "errors"

var (
	InvalidMSISDN         = errors.New("invalid MSISDN")
	InvalidNSN            = errors.New("invalid NSN")
	InvalidOperatorPrefix = errors.New("invalid operator prefix")
)
