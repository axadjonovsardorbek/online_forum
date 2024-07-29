package postgres

import (
	"database/sql"
	"fmt"

	"forum-service/config"
	"forum-service/config/logger"
	"forum-service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db        *sql.DB
	Logger    *logger.Logger
	CategoryS storage.CategoryI
	CommentS  storage.CommentI
	PostTagS  storage.PostTagI
	PostS     storage.PostI
	TagS      storage.TagI
}

func NewPostgresStorage(config config.Config, logger *logger.Logger) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	category := NewCategoryRepo(db, logger)
	comment := NewCommentRepo(db, logger)
	postTag := NewPostTagRepo(db, logger)
	post := NewPostRepo(db, logger)
	tag := NewTagRepo(db, logger)

	return &Storage{
		Db:        db,
		CategoryS: category,
		CommentS:  comment,
		PostTagS:  postTag,
		PostS:     post,
		TagS:      tag,
	}, nil
}
