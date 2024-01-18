package idn_mobile_number_validator

import "regexp"

var (
	regexPhoneNumberE164WithoutPlus = `^628[0-9]\d{7,10}$`
	regexPhoneNumberE164WithPlus    = `^\+628[0-9]\d{7,10}$`
	regexPhoneNumberWithZero        = `^08[0-9]\d{7,10}$`
	regexPhoneNumberWithoutZero     = `^8[0-9]\d{7,10}$`
)

func ValidateE164(msisdn string, withPlus ...bool) error {
	matcher := regexPhoneNumberE164WithoutPlus
	if len(withPlus) > 0 && withPlus[0] {
		matcher = regexPhoneNumberE164WithPlus
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(msisdn, -1)
	if len(matchingNumber) == 0 {
		return InvalidMSISDN
	}

	return nil
}

func ValidateNSN(nsn string, withZero ...bool) error {
	matcher := regexPhoneNumberWithoutZero
	if len(withZero) > 0 && withZero[0] {
		matcher = regexPhoneNumberWithZero
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(nsn, -1)
	if len(matchingNumber) == 0 {
		return InvalidNSN
	}

	return nil
}
