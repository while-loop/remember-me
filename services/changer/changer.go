package changer

import (
	api "github.com/while-loop/remember-me/api/services/v1/changer"
	"github.com/while-loop/remember-me"
	"github.com/while-loop/remember-me/services"
	"log"
	"github.com/while-loop/remember-me/util"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/while-loop/remember-me/managers"
)

//go:generate protoc -I ../../proto/ --go_out=plugins=grpc:../../api/services/v1/changer/ ../../proto/changer.proto

func init() {
	services.Register("changer", NewService)
}

type ChangeService struct {
	app *remme.App
}

func NewService(app *remme.App) services.Service {
	return &ChangeService{app: app}
}

func (c *ChangeService) Register(rpc *grpc.Server) {
	api.RegisterChangerServer(rpc, c)
	log.Println("ChangeService started")
}

func (c *ChangeService) ChangePassword(req *api.ChangeRequest, stream api.Changer_ChangePasswordServer) error {
	man, err := remme.GetManager(req.Manager.String(), req.Email, req.Password)
	if err != nil {
		return err
	}

	pwdgen := util.NewPasswordGenP(req.PasswdConfig)
	statusChan := make(chan api.Status)
	go c.app.ChangePasswords(statusChan, man, pwdgen.Generate)

	for status := range statusChan {
		stream.Send(&status)
	}
	return nil
}

func (c *ChangeService) GetManagers(ctx context.Context, req *api.ManagersRequest) (*api.ManagersReply, error) {
	return &api.ManagersReply{
		Managers: managers.GetManagers(),
	}, nil
}
