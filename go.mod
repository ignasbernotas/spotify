module github.com/89z/spotify

go 1.17

require github.com/librespot-org/librespot-golang v0.0.0-20200423180623-b19a2f10c856

require (
	github.com/badfortrains/mdns v0.0.0-20160325001438-447166384f51 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/librespot-org/librespot-golang => ./respot
