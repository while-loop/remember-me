package main

// grpc services
import (
	_ "github.com/while-loop/remember-me/services/changer"
	_ "github.com/while-loop/remember-me/services/record"
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
