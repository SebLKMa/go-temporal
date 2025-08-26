package surldb

import (
	"fmt"
	db "src/github.com/seblkma/go-temporal/surldb"
	dm "src/github.com/seblkma/go-temporal/surldm"
	ut "src/github.com/seblkma/go-temporal/surlut"
	"testing"
	"time"
)

// go clean -testcache

// GOFLAGS="-count=1" go test -v surl_id_test.go -run TestCreateOrUpdateSurlId
func TestCreateOrUpdateSurlId(t *testing.T) {
	dbconn := db.NewConnection("localhost", 5432, "pdtester", "pdtester", "pddbtest")
	surl_test_db := dbconn.Connect()
	defer surl_test_db.Close()

	longUrl := "https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/README.md"
	var data dm.SurlId
	data.UniqueID = ut.GenId(longUrl)
	if data.UniqueID == "" {
		t.Error("Error in processing URL")
		t.FailNow()
	}
	data.LongUrl = longUrl
	data.ShortUrl = "https://go/" + data.UniqueID
	data.ExpiresOn = time.Now().AddDate(2, 0, 0) // expires in 2 years

	err := db.CreateOrUpdateSurlId(surl_test_db, data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	surldata, err := db.GetSurlIdByLongUrl(surl_test_db, longUrl)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("Record : %+v\n", surldata)
}

// GOFLAGS="-count=1" go test -v surl_id_test.go -run TestGetSurlId
func TestGetSurlId(t *testing.T) {
	dbconn := db.NewConnection("localhost", 5432, "pdtester", "pdtester", "pddbtest")
	surl_test_db := dbconn.Connect()
	defer surl_test_db.Close()

	longUrl := "https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/README.md"

	surldata, err := db.GetSurlIdByLongUrl(surl_test_db, longUrl)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("Record : %+v\n", surldata)
}
