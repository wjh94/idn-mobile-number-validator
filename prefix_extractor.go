package idn_mobile_number_validator

import "regexp"

func extractPrefixAndSubscriberNumberFromE164(msisdn string, withPlus ...bool) (string, string, error) {
	matcher := regexPhoneNumberE164WithoutPlus
	prefixStartIndex := 3
	if len(withPlus) > 0 && withPlus[0] {
		matcher = regexPhoneNumberE164WithPlus
		prefixStartIndex += 1
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(msisdn, -1)
	if len(matchingNumber) == 0 || len(matchingNumber) != 1 {
		return "", "", InvalidMSISDN
	}

	prefix := matchingNumber[0][prefixStartIndex : prefixStartIndex+prefixLength]
	subscriberNumber := matchingNumber[0][prefixStartIndex+prefixLength:]
	return prefix, subscriberNumber, nil
}

func extractPrefixAndSubscriberNumberFromNSN(nsn string, withZero ...bool) (string, string, error) {
	matcher := regexPhoneNumberWithoutZero
	prefixStartIndex := 1
	if len(withZero) > 0 && withZero[0] {
		matcher = regexPhoneNumberWithZero
		prefixStartIndex += 1
	}

	regex := regexp.MustCompile(matcher)
	matchingNumber := regex.FindAllString(nsn, -1)
	if len(matchingNumber) == 0 || len(matchingNumber) != 1 {
		return "", "", InvalidNSN
	}

	prefix := matchingNumber[0][prefixStartIndex : prefixStartIndex+prefixLength]
	subscriberNumber := matchingNumber[0][prefixStartIndex+prefixLength:]
	return prefix, subscriberNumber, nil
}
