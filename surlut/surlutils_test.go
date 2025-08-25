package surlutils

import (
	"fmt"
	id "src/github.com/seblkma/go-temporal/surlut"
	"testing"
)

// GOFLAGS="-count=1" go test -v surlutils_test.go -run TestGenXid
func TestGenXid(t *testing.T) {
	xid := id.GenXid()
	fmt.Printf("GenXid: %s\n", xid)
}

// GOFLAGS="-count=1" go test -v surlutils_test.go -run TestGenXidBase62
func TestGenXidBase62(t *testing.T) {
	xid := id.GenXid()
	fmt.Printf("GenXid: %s\n", xid)

	b62Id := id.Base62Encode(xid)
	fmt.Printf("Base62Encode: %s\n", b62Id)

	xid2, err := id.Base62Decode(b62Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("Base62Decode: %s\n", xid2)
}
