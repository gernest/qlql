build:
	gopherjs build github.com/gernest/qlql/app -o app.js

build-go:
	go build -o ql-api-server ./api

