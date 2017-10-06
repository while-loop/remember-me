package main

// grpc services
import (
	_ "github.com/while-loop/remember-me/services/v1/changer"
	_ "github.com/while-loop/remember-me/services/v1/record"
)

// managers
import (
	_ "github.com/while-loop/remember-me/managers/lastpass"
)

// webservices
import (
	_ "github.com/while-loop/remember-me/webservices/facebook"
)

// dbs
import (
	_ "github.com/while-loop/remember-me/db/dynamodb"
)
