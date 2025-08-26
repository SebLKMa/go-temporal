package surldb

import (
	"database/sql"
	"fmt"
	"time"

	dm "src/github.com/seblkma/go-temporal/surldm"
)

func CreateSurlId(db *sql.DB, data dm.SurlId) error {

	t := time.Now()
	s := `INSERT INTO surl_id(
		unique_id,
		long_url,
		short_url,
		expires_on,
		created_on,
		modified_on
		) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(s,
		data.UniqueID,
		data.LongUrl,
		data.ShortUrl,
		data.ExpiresOn,
		t,
		t,
	)

	if err != nil {
		return err
	}
	return nil
}

func UpdateSurlId(db *sql.DB, data dm.SurlId) error {

	t := time.Now()
	s := `UPDATE surl_id
		SET 
		long_url = $1,
		short_url = $2,
		expires_on = $3,
		modified_on = $4
		WHERE unique_id = $5
	`

	_, err := db.Exec(s,
		data.LongUrl,
		data.ShortUrl,
		data.ExpiresOn,
		t,
		data.UniqueID,
	)

	if err != nil {
		return err
	}
	return nil
}

func CreateOrUpdateSurlId(db *sql.DB, data dm.SurlId) error {

	_, err := GetSurlIdByLongUrl(db, data.LongUrl)
	if err != nil {
		//fmt.Printf("%+v\n", err)
		return CreateSurlId(db, data)
	} else {
		return UpdateSurlId(db, data)
	}
}

func GetSurlIdByLongUrl(db *sql.DB, longUrl string) (dm.SurlId, error) {
	fmt.Println(longUrl)
	s := `
	SELECT 
    *
	FROM surl_id
	WHERE long_url = $1 
	`
	var data dm.SurlId
	r := db.QueryRow(s, longUrl)
	err := r.Scan(
		&data.IndexID,
		&data.UniqueID,
		&data.LongUrl,
		&data.ShortUrl,
		&data.ExpiresOn,
		&data.CreatedOn,
		&data.ModifiedOn,
	)
	if err != nil {
		return data, err
	}
	fmt.Println(data)
	/* not in use
	var results []mdl.AlarmConfigThresholds
	rows, err := db.Query(s, pq.Array(paths))
	if err != nil {
		return results, err
	}
	defer rows.Close()

	resultsMap := make(map[string]mdl.AlarmConfigThresholds, len(paths))
	for rows.Next() {
		c := mdl.AlarmConfigThresholds{}
		err = rows.Scan(
			&c.IndexID,
			&c.InternalID,
			&c.SensorQmaxTh1,
			&c.CreatedOn,
			&c.ModifiedOn,
		)
		if err != nil {
			return results, err
		}

		resultsMap[c.InternalID] = c
	}
	*/
	return data, nil
}

func GetSurlIdByShortUrl(db *sql.DB, shortUrl string) (dm.SurlId, error) {
	fmt.Println(shortUrl)
	s := `
	SELECT 
    *
	FROM surl_id
	WHERE short_url = $1 
	`
	var data dm.SurlId
	r := db.QueryRow(s, shortUrl)
	err := r.Scan(
		&data.IndexID,
		&data.UniqueID,
		&data.LongUrl,
		&data.ShortUrl,
		&data.ExpiresOn,
		&data.CreatedOn,
		&data.ModifiedOn,
	)
	if err != nil {
		return data, err
	}
	fmt.Println(data)

	return data, nil
}
