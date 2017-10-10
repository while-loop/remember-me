package main

import (
	"context"
	"fmt"
	"github.com/while-loop/remember-me"
	changer_pb "github.com/while-loop/remember-me/api/services/v1/changer"
	"google.golang.org/grpc"
	"log"
)

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
	"github.com/while-loop/remember-me/manager"
	"github.com/while-loop/remember-me/storage/stub"
	"github.com/while-loop/remember-me/webservice"
	"github.com/while-loop/remember-me/util"
)

func main() {
	// Set up a connection to the server.

	manStr, email, password := "", "", ""
	local(manStr, email, password)
	grpcc(email, password)
}

func local(manStr, email, password string) {
	man, err := manager.GetManager(manStr, email, password)
	if err != nil {
		log.Fatalf("%s", err)
	}

	app := remme.NewApp(stub.New(), webservice.Services())
	statusChan := make(chan changer_pb.Status)
	go app.ChangePasswords(statusChan, man, util.DefaultPasswdFunc)
	for status := range statusChan {
		log.Print(status)
	}
}

func grpcc(email, password string) {
	fmt.Println("setting client")
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := changer_pb.NewChangerClient(conn)

	req := &changer_pb.ChangeRequest{
		Password: password,
		Email:    email,
		Manager:  changer_pb.ChangeRequest_LASTPASS,
		PasswdConfig: &changer_pb.PasswdConfig{
			Length: 5,
		},
	}

	stream, err := c.ChangePassword(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		status, err := stream.Recv()
		if err != nil {
			fmt.Println("err", err)
			break
		}
		fmt.Printf("%v\n", status)
	}
}
