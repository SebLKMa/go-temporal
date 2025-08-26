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

// GOFLAGS="-count=1" go test -v surlutils_test.go -run TestGenId
func TestGenId(t *testing.T) {
	shortId := id.GenId("https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/README.md")
	fmt.Printf("GenId: %s\n", shortId)

	shortId = id.GenId("https://www.ardanlabs.com/blog/2018/12/schedulingingopart3.html?mc_cid=66407ad02b&mc_eid=41cad80de5")
	fmt.Printf("GenId: %s\n", shortId)
}
