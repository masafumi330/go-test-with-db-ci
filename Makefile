test:
	# migration
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:password@tcp(localhost:3310)/test_db" goose -dir ./migrate up

	# test
	cd app && go test -v ./...
