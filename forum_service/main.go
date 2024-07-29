package main

import (
	"forum-service/config/logger"
	"forum-service/storage/postgres"
	"log"
	"net"
	"path/filepath"
	"runtime"

	"google.golang.org/grpc"
	cf "forum-service/config"
	pb "forum-service/genproto/forum"
	service "forum-service/service"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	db, err := postgres.NewPostgresStorage(config, logger)
	em.CheckErr(err)
	defer db.Db.Close()

	listener, err := net.Listen("tcp", config.FORUM_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCategoryServiceServer(s, service.NewCategoryService(db))
	pb.RegisterTagServiceServer(s, service.NewTagService(db))
	pb.RegisterCommentServiceServer(s, service.NewCommentService(db))
	pb.RegisterPostServiceServer(s, service.NewPostServiceService(db))
	pb.RegisterPostTagsServiceServer(s, service.NewPostTagsService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
