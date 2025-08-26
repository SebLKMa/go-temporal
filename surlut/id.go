package surlutils

import (
	"math"
	"strings"

	"github.com/inovacc/base62"
	"github.com/rs/xid"
)

// GenXid returns a unique Id
// See:
// https://blog.kowalczyk.info/article/JyRZ/generating-good-random-and-unique-ids-in-go.html
func GenXid() string {
	id := xid.New()
	//fmt.Printf("github.com/rs/xid:              %s\n", id.String())
	return id.String()
}

func Base62Encode(s string) string {
	data := []byte(s)
	encoded := base62.Encode(data)
	return encoded
}

func Base62Decode(s string) (string, error) {
	decoded, err := base62.Decode(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Credit to a reddit post
/*
https://www.reddit.com/r/golang/comments/15ai4sb/what_is_base62_conversion_checking_my_own/

## What is base62 conversion? Checking my own understanding.

Ran into my own limited understanding of base62 conversion while building out a URL shortener in Go and wanted to check my understanding by posting this. :)

Lmk if the comments if I'm misunderstanding something or if this helped you understand it better.

---

Base62 is a common encoding scheme that's perfect for short URL services and compact data storage.

It gets its names from the sixty-two characters used for encoding. Those characters are [0-9] 10 + [a-z] 26 + [A-Z] 26 = 62. To convert a decimal to base62 you use long division to divide the decimal by sixty-two. Then you take the remainder and map it to the character that represents it by locating its position within the base62 digits.

Example: Convert 123 to base62

### Step 1: Divide the decimal by 62 until the quotient is zero and keep track of the remainders.

123 ÷ 62 = 1 quotient 61 remainder

1 ÷ 62 = 0 quotient 1 remainder

### Step 2: Use the remainder to map to the base62 character by its position

61 = Z

The base62 system starts its index at zero, so the last character Z is at position sixty-one

1 = 1

Therefore, 123 converted to base62 is 1Z.

When combining the base62 characters you go from right to left to match the decimal positions.

## Now, let's convert it back into a decimal.

### Step 1: Identify their position value in the base62 system

1 = 1st position

Z = 61st position

### Step 2: Calculate the powers of 62

The calculation of the appropriate power of 62 for each base-62 digit is based on its position.

You find the power of each digit by taking the length of the base62 number minus the digit's position in the number and then subtract 1.

Formula: len(base62) -i -1

1Z = length of 2

Z's power of 62 eq 2 - 1 - 1 = 0

1's power of 62 eq 2 - 0 - 1 = 1

### Step 3: Convert back to decimals

1 * 61^1 = 61

62 * 62^0 = 62

61+62 = 123

Here's some sample code in #golang for reference. :)
https://play.golang.com/p/DmFYZXWdzDU
*/

// Not handling other characters (e.g. -?&=_?)for now
const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.//:"
const shortenerAddress = "http://localhost:4000/v1/"

func DecimalToBase62(n int64) string {
	if n == 0 {
		return "0"
	}

	base62 := make([]byte, 0)
	radix := int64(62)

	for n > 0 {
		remainder := n % radix
		base62 = append([]byte{base62Digits[remainder]}, base62...)
		n /= radix
	}

	return string(base62)
}

func Base62ToDecimal(s string) int64 {
	var decimalNumber int64

	for i, c := range s {
		decimalNumber += int64(strings.IndexByte(base62Digits, byte(c))) * int64(math.Pow(62, float64(len(s)-i-1)))
	}

	return decimalNumber
}

func GenId(s string) string {
	decimalNumber := Base62ToDecimal(s)
	base62Representation := DecimalToBase62(decimalNumber)
	//fmt.Printf("Base62 representation of %d is %s\n", decimalNumber, base62Representation)
	return base62Representation
}
