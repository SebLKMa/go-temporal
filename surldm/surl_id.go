package surldm

import "time"

type SurlId struct {
	IndexID    int64     `json:"IndexID"`
	UniqueID   string    `json:"UniqueID"`
	LongUrl    string    `json:"LongUrl"`
	ShortUrl   string    `json:"ShortUrl"`
	ExpiresOn  time.Time `json:"ExpiresOn"`
	CreatedOn  time.Time `json:"CreatedOn"`
	ModifiedOn time.Time `json:"ModifiedOn"`
}

type InputString struct {
	Input string `json:"Input"`
}
