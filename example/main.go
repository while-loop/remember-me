package main

// managers
import (
	_ "github.com/while-loop/remember-me/remme/manager/lastpass"
)

// webservices
import (
	_ "github.com/while-loop/remember-me/remme/webservice/facebook"
)

import (
	"context"
	"fmt"

	"time"

	"github.com/sirupsen/logrus"
	"github.com/while-loop/remember-me/remme"
	"github.com/while-loop/remember-me/remme/api/services/v1/changer"
	"github.com/while-loop/remember-me/remme/log"
	"github.com/while-loop/remember-me/remme/manager"
	"github.com/while-loop/remember-me/remme/storage/stub"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme/webservice"
	"google.golang.org/grpc"
)

func main() {
	log.Logger.SetLevel(logrus.DebugLevel)
	// Set up a connection to the server.
	manStr, email, password := "lastpass", "", ""
	local(manStr, email, password)
	//grpcMode(email, password)
}

func local(manStr, email, password string) {
	man, err := manager.GetManager(manStr, email, password)
	if err != nil {
		log.Fatalf("%s", err)
	}

	app := remme.NewApp(stub.New(), webservice.Services())
	log.Info(man.GetSites())
	jobId := app.ChangePasswords(man, util.DefaultPasswdFunc)
	fmt.Println(jobId)
	time.Sleep(5 * time.Second)
}

func grpcMode(email, password string) {
	fmt.Println("setting client")
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := changer.NewChangerClient(conn)

	req := &changer.ChangeRequest{
		Password: password,
		Email:    email,
		Manager:  changer.ChangeRequest_LASTPASS,
		PasswdConfig: &changer.PasswdConfig{
			Length: 8,
		},
	}

	reply, err := c.ChangePassword(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(reply)

}
