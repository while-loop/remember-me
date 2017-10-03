package main

// remme daemon server
// currently only serves gRPC

import (
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/while-loop/remember-me"
	"github.com/while-loop/remember-me/services"
	"github.com/while-loop/remember-me/db/dynamodb"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpc := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	db := dynamodb.NewDynamoDB()
	app := remme.NewApp(db, remme.WebServices())

	services.StartServices(app, rpc)

	// Register reflection service on gRPC server.
	reflection.Register(rpc)

	metrics(rpc)

	log.Println("Running gRPC Server", lis.Addr())
	if err := rpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func metrics(rpc *grpc.Server) {
	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Println("no metrics for you")
		return
	}

	grpc_prometheus.Register(rpc)
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	go http.Serve(l, m)
	log.Println("running metrics", l.Addr())

}
