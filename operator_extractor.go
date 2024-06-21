package idn_mobile_number_validator

func GetProviderFromE164(msisdn string, withPlus ...bool) (Operator, error) {
	if err := ValidateE164(msisdn, withPlus...); err != nil {
		return 0, err
	}

	prefix, _, _ := extractPrefixAndSubscriberNumberFromE164(msisdn, withPlus...)
	operator, isFound := operatorByPrefix[prefix]
	if !isFound {
		return 0, InvalidOperatorPrefix
	}

	return operator, nil
}

func GetProviderFromNSN(nsn string, withZero ...bool) (Operator, error) {
	if err := ValidateNSN(nsn, withZero...); err != nil {
		return 0, err
	}

	prefix, _, _ := extractPrefixAndSubscriberNumberFromNSN(nsn, withZero...)
	operator, isFound := operatorByPrefix[prefix]
	if !isFound {
		return 0, InvalidOperatorPrefix
	}

	return operator, nil
}
