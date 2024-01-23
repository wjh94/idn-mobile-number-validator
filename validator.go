package idn_mobile_number_validator

import (
	"regexp"
)

const (
	prefixLength = 2
)

var (
	regexPhoneNumberE164WithoutPlus = `^628[0-9]\d{7,10}$`
	regexPhoneNumberE164WithPlus    = `^\+628[0-9]\d{7,10}$`
	regexPhoneNumberWithZero        = `^08[0-9]\d{7,10}$`
	regexPhoneNumberWithoutZero     = `^8[0-9]\d{7,10}$`
)

func ValidateE164(msisdn string, withPlus ...bool) error {
	matcher := regexPhoneNumberE164WithoutPlus
	prefixStartIndex := 3
	if len(withPlus) > 0 && withPlus[0] {
		matcher = regexPhoneNumberE164WithPlus
		prefixStartIndex += 1
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(msisdn, -1)
	if len(matchingNumber) == 0 || len(matchingNumber) != 1 {
		return InvalidMSISDN
	}

	prefix := matchingNumber[0][prefixStartIndex : prefixStartIndex+prefixLength]
	subscriberNumber := matchingNumber[0][prefixStartIndex+prefixLength:]
	if !validPrefixes[prefix] {
		return InvalidOperatorPrefix
	}
	if len(subscriberNumber) < minLengthPerPrefix[prefix] || len(subscriberNumber) > maxLengthPerPrefix[prefix] {
		return InvalidSubscriberNumberLength
	}

	return nil
}

func ValidateNSN(nsn string, withZero ...bool) error {
	matcher := regexPhoneNumberWithoutZero
	prefixStartIndex := 1
	if len(withZero) > 0 && withZero[0] {
		matcher = regexPhoneNumberWithZero
		prefixStartIndex += 1
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(nsn, -1)
	if len(matchingNumber) == 0 || len(matchingNumber) != 1 {
		return InvalidNSN
	}

	prefix := matchingNumber[0][prefixStartIndex : prefixStartIndex+prefixLength]
	subscriberNumber := matchingNumber[0][prefixStartIndex+prefixLength:]
	if !validPrefixes[prefix] {
		return InvalidOperatorPrefix
	}
	if len(subscriberNumber) < minLengthPerPrefix[prefix] || len(subscriberNumber) > maxLengthPerPrefix[prefix] {
		return InvalidSubscriberNumberLength
	}

	return nil
}
