package idn_mobile_number_validator

type Operator int

const (
	TELKOMSEL Operator = iota + 1
	IOH
	XL
	SMARTFREN
)

var operatorByPrefix = map[string]Operator{
	"11": TELKOMSEL,
	"12": TELKOMSEL,
	"13": TELKOMSEL,
	"14": IOH,
	"15": IOH,
	"16": IOH,
	"17": XL,
	"18": XL,
	"19": XL,
	"21": TELKOMSEL,
	"22": TELKOMSEL,
	"23": TELKOMSEL,
	"31": XL,
	"32": XL,
	"33": XL,
	"38": XL,
	"51": TELKOMSEL,
	"52": TELKOMSEL,
	"53": TELKOMSEL,
	"55": IOH,
	"56": IOH,
	"57": IOH,
	"58": IOH,
	"59": XL,
	"77": XL,
	"78": XL,
	"79": XL,
	"81": SMARTFREN,
	"82": SMARTFREN,
	"83": SMARTFREN,
	"84": SMARTFREN,
	"85": SMARTFREN,
	"86": SMARTFREN,
	"87": SMARTFREN,
	"88": SMARTFREN,
	"89": SMARTFREN,
	"95": IOH,
	"96": IOH,
	"97": IOH,
	"98": IOH,
	"99": IOH,
}
