package record

import (
	"log"

	"github.com/while-loop/remember-me/remme"
	api "github.com/while-loop/remember-me/remme/api/services/v1/record"
	"github.com/while-loop/remember-me/remme/service"
	"github.com/while-loop/remember-me/remme/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//go:generate protoc -I ../../../proto/ --go_out=plugins=grpc:../../../api/services/v1/record/ ../../../proto/record.proto

var _ api.RecordServer = &recordService{}

type recordService struct {
	db storage.DataStore
}

func init() {
	service.Register("record", New)
}

func New(app *remme.App) service.Service {
	return &recordService{db: app.Datastore}
}

func (r *recordService) Register(rpc *grpc.Server) {
	api.RegisterRecordServer(rpc, r)
	log.Println("RecordService started")
}

func (r *recordService) GetRecord(ctx context.Context, req *api.RecordRequest) (*api.LogRecord, error) {
	lr, err := r.db.GetLog(req.JobId)
	if err != nil {
		return nil, err
	}

	return lr2Proto(lr), nil
}

func (r *recordService) TailEvents(*api.RecordRequest, api.Record_TailEventsServer) error {
	return nil
}
