package main

// grpc services
import (
	_ "github.com/while-loop/remember-me/remme/service/v1/changer"
	_ "github.com/while-loop/remember-me/remme/service/v1/record"
)

// managers
import (
	_ "github.com/while-loop/remember-me/remme/manager/lastpass"
)

// webservices
import (
	_ "github.com/while-loop/remember-me/remme/webservice/facebook"
)

// dbs
import (
	_ "github.com/while-loop/remember-me/remme/storage/dynamodb"
)
