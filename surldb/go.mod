module src/github.com/seblkma/go-temporal/surldb

go 1.24

toolchain go1.24.6

replace (
	src/github.com/seblkma/go-temporal/surldm => ../surldm
	src/github.com/seblkma/go-temporal/surlut => ../surlut
)

require (
	github.com/lib/pq v1.10.9
	src/github.com/seblkma/go-temporal/surldm v0.0.0-00010101000000-000000000000
	src/github.com/seblkma/go-temporal/surlut v0.0.0-00010101000000-000000000000
)

require (
	github.com/inovacc/base62 v1.0.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
)
