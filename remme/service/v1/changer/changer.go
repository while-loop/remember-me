package changer

import (
	"log"

	"time"

	"github.com/while-loop/remember-me/remme"
	api "github.com/while-loop/remember-me/remme/api/services/v1/changer"
	"github.com/while-loop/remember-me/remme/manager"
	"github.com/while-loop/remember-me/remme/service"
	"github.com/while-loop/remember-me/remme/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//go:generate protoc -I ../../../proto/ --go_out=plugins=grpc:../../../api/services/v1/changer/ ../../../proto/changer.proto

func init() {
	service.Register("changer", New)
}

type ChangeService struct {
	app *remme.App
}

func New(app *remme.App) service.Service {
	return &ChangeService{app: app}
}

func (c *ChangeService) ChangePassword(ctx context.Context, req *api.ChangeRequest) (*api.ChangeReply, error) {
	man, err := manager.GetManager(req.Manager.String(), req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	pwdgen := util.NewPasswordGen(uint(req.PasswdConfig.Length), req.PasswdConfig.SpecialChars, req.PasswdConfig.Numbers)
	jobId := c.app.ChangePasswords(man, pwdgen.Generate)

	reply := &api.ChangeReply{
		JobId:     jobId,
		StartTime: uint64(time.Now().Unix()),
	}
	return reply, nil
}

func (c *ChangeService) Register(rpc *grpc.Server) {
	api.RegisterChangerServer(rpc, c)
	log.Println("ChangeService started")
}

func (c *ChangeService) GetManagers(ctx context.Context, req *api.ManagersRequest) (*api.ManagersReply, error) {
	return &api.ManagersReply{
		Managers: manager.GetManagers(),
	}, nil
}
