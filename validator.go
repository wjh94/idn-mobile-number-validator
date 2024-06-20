package idn_mobile_number_validator

const (
	prefixLength = 2
)

var (
	regexPhoneNumberE164WithoutPlus = `^628[0-9]\d{7,10}$`
	regexPhoneNumberE164WithPlus    = `^\+628[0-9]\d{7,10}$`
	regexPhoneNumberWithZero        = `^08[0-9]\d{7,10}$`
	regexPhoneNumberWithoutZero     = `^8[0-9]\d{7,10}$`
)

// ValidateE164 validates if the given msisdn represents a valid format of Indonesian mobile phone number
// in ITU E.164 format. (Reference: https://www.itu.int/rec/T-REC-E.164-201011-I/en)
//
// An optional true boolean can be added if the validation process involves the international calling prefix "+" (plus).
// Thus, the received msisdn input should start by either "628", or "+628" if the true boolean parameter is added.
//
// A nil error will be expected if validation is successful. Otherwise, below error code will be returned:
//
//  - InvalidMSISDN: returned if the msisdn does not satisfy the matcher criteria.
//  - InvalidOperatorPrefix: returned if the msisdn does not contain the valid mobile number prefixes.
//  - InvalidSubscriberNumberLength: returned if the msisdn does not satisfy valid subscriber number
//    (digits following prefix) length.
func ValidateE164(msisdn string, withPlus ...bool) error {
	prefix, subscriberNumber, err := extractPrefixAndSubscriberNumberFromE164(msisdn, withPlus...)
	if err != nil {
		return err
	}

	if !validPrefixes[prefix] {
		return InvalidOperatorPrefix
	}
	if len(subscriberNumber) < minLengthPerPrefix[prefix] || len(subscriberNumber) > maxLengthPerPrefix[prefix] {
		return InvalidSubscriberNumberLength
	}

	return nil
}

// ValidateNSN validates if the given msisdn represents a valid format of Indonesian mobile phone number
// as an NSN (national (significant) number). (Reference: https://www.itu.int/rec/T-REC-E.164-201011-I/en)
//
// An optional true boolean can be added if the validation process involves the national trunk prefix "0".
// Thus, the received msisdn input should start by either "8", or "08" if the true boolean parameter is added.
//
// A nil error will be expected if validation is successful. Otherwise, below error code will be returned:
//
//  - InvalidMSISDN: returned if the msisdn does not satisfy the matcher criteria.
//  - InvalidOperatorPrefix: returned if the msisdn does not contain the valid mobile number prefixes.
//  - InvalidSubscriberNumberLength: returned if the msisdn does not satisfy valid subscriber number
//    (digits following prefix) length.
func ValidateNSN(nsn string, withZero ...bool) error {
	prefix, subscriberNumber, err := extractPrefixAndSubscriberNumberFromNSN(nsn, withZero...)
	if err != nil {
		return err
	}

	if !validPrefixes[prefix] {
		return InvalidOperatorPrefix
	}
	if len(subscriberNumber) < minLengthPerPrefix[prefix] || len(subscriberNumber) > maxLengthPerPrefix[prefix] {
		return InvalidSubscriberNumberLength
	}

	return nil
}
