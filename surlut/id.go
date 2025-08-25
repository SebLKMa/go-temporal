package surlutils

import (
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

func GenId() string {
	return GenXid()
}
