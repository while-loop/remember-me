package main

// grpc services
import (
	_ "github.com/while-loop/remember-me/remme/serviceervice/v1/changer"
	_ "github.com/while-loop/remember-me/remme/serviceervice/v1/record"
)

// managers
import (
	_ "github.com/while-loop/remember-me/remme/manageranager/lastpass"
)

// webservices
import (
	_ "github.com/while-loop/remember-me/remme/webserviceervice/facebook"
)

// dbs
import (
	_ "github.com/while-loop/remember-me/remme/storagetorage/dynamodb"
)
