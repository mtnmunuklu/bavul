#!/bin/bash

#go clean --cache && go test -v -cover ./security/... ./authentication/...
go build -o ../authentication/authsvc ../authentication/main.go
go build -o ../vulnerability/vulnsvc ../vulnerability/main.go
go build -o ../api/apisvc ../api/main.go
go build -o ../web/websvc ../web/main.go