package main

// grpc services
import (
	_ "github.com/while-loop/remember-me/service/v1/changer"
	_ "github.com/while-loop/remember-me/service/v1/record"
)

// managers
import (
	_ "github.com/while-loop/remember-me/manager/lastpass"
)

// webservices
import (
	_ "github.com/while-loop/remember-me/webservice/facebook"
)

// dbs
import (
	_ "github.com/while-loop/remember-me/storage/dynamodb"
)
