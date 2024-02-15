# IDN Mobile Number Validator

---

### Description

This library is intended to validate Indonesian mobile phone numbers (MSISDN).

Reference: [JDIH Kominfo](https://jdih.kominfo.go.id/produk_hukum/view/id/620/t/peraturan+menteri+komunikasi+dan+informatika+nomor+14+tahun+2018+tanggal+2+oktober+2018)

### Getting Started

Assuming the project is using Go Mod as dependency manager,
this library can be included by running this command in the project root:

```shell
go get github.com/wjh94/idn-mobile-number-validator@latest
```

### Example

#### ValidateE164

##### Code

```go
package main

import (
	"fmt"

	mobileNrValidator "github.com/wjh94/idn-mobile-number-validator"
)

func main() {
	phoneNumber := "6281200000000"
	if err := mobileNrValidator.ValidateE164(phoneNumber); err != nil {
		fmt.Println("Phone number failed on validation:", err)
		return
        }
	
	fmt.Println("Phone number passed validation")
}
```

##### Result

```
Phone number passed validation
```

#### ValidateNSN

##### Code

```go
package main

import (
	"fmt"

	mobileNrValidator "github.com/wjh94/idn-mobile-number-validator"
)

func main() {
	phoneNumber := "81200000000"
	if err := mobileNrValidator.ValidateNSN(phoneNumber); err != nil {
		fmt.Println("Phone number failed on validation:", err)
		return
        }
	
	fmt.Println("Phone number passed validation")
}
```

##### Result

```
Phone number passed validation
```

### Have recommendations?

You are welcome to open an issue for this repository.
Discussions will be made on the comment section.