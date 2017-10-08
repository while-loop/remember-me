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

func main() {
	// Set up a connection to the server.

	manStr, email, password := "", "", ""
	local(manStr, email, password)
	grpcc(email, password)
}

func local(manStr, email, password string) {
	man, err := remme.GetManager(manStr, email, password)
	if err != nil {
		log.Fatalf("%s", err)
	}

	app := remme.NewApp(remme.DefaultDB(), remme.WebServices())
	statusChan := make(chan changer_pb.Status)
	go app.ChangePasswords(statusChan, man, remme.DefaultPasswdFunc)
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
