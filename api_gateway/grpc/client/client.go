package client

import (
	"api-gateway/config"
	"api-gateway/config/logger"
	pbf "api-gateway/genproto/forum"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Comment  pbf.CommentServiceClient
	Category pbf.CategoryServiceClient
	Post     pbf.PostServiceClient
	PostTag  pbf.PostTagsServiceClient
	Tag      pbf.TagServiceClient
	Logger   logger.Logger
}

func NewGrpcClients(cfg *config.Config, l logger.Logger) (*GrpcClients, error) {
	// connF, err := grpc.NewClient(fmt.Sprintf("localhost%s", "forum_service"+cfg.FORUM_SERVICE_PORT),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }
	connF, err := grpc.NewClient(fmt.Sprintf("forum2%s", cfg.FORUM_SERVICE_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Comment:  pbf.NewCommentServiceClient(connF),
		Category: pbf.NewCategoryServiceClient(connF),
		Post:     pbf.NewPostServiceClient(connF),
		PostTag:  pbf.NewPostTagsServiceClient(connF),
		Tag:      pbf.NewTagServiceClient(connF),
		Logger:   l,
	}, nil
}
