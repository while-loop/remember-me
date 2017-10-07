package record

import (
	"github.com/while-loop/remember-me"
	api "github.com/while-loop/remember-me/api/services/v1/record"
	"github.com/while-loop/remember-me/db"
	"github.com/while-loop/remember-me/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

//go:generate protoc -I ../../../proto/ --go_out=plugins=grpc:../../../api/services/v1/record/ ../../../proto/record.proto

var _ api.RecordServer = &RecordService{}

type RecordService struct {
	db db.DataStore
}

func init() {
	services.Register("record", NewService)
}

func NewService(app *remme.App) services.Service {
	return &RecordService{db: app.Datastore}
}

func (r *RecordService) Register(rpc *grpc.Server) {
	api.RegisterRecordServer(rpc, r)
	log.Println("RecordService started")
}

func (r *RecordService) GetRecord(ctx context.Context, req *api.RecordRequest) (*api.LogRecord, error) {
	lr, err := r.db.GetLog(req.JobId)
	if err != nil {
		return nil, err
	}

	return lr2Proto(lr), nil
}
