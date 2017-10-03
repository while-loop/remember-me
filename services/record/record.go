package record

import (
	api "github.com/while-loop/remember-me/api/services/v1/record"
	"github.com/while-loop/remember-me/services"
	"github.com/while-loop/remember-me"
	"google.golang.org/grpc"
	"github.com/while-loop/remember-me/db"
	"golang.org/x/net/context"
	"log"
)

//go:generate protoc -I ../../api/services/v1/record/ --go_out=plugins=grpc:../../api/services/v1/record/ ../../api/services/v1/record/record.proto

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

