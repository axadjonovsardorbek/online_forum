package service

import (
	"context"
	f "forum-service/genproto/forum"
	st "forum-service/storage/postgres"
)

type CommentService struct {
	storage st.Storage
	f.UnimplementedCommentServiceServer
}

func NewCommentService(storage *st.Storage) *CommentService {
	return &CommentService{
		storage: *storage,
	}
}

func (s *CommentService) Create(ctx context.Context, req *f.CommentCreateReq) (*f.Comment, error) {
	res, err := s.storage.CommentS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CommentService) GetById(ctx context.Context, id *f.GetByIdReq) (*f.CommentRes, error) {
	res, err := s.storage.CommentS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CommentService) GetAll(ctx context.Context, filter *f.Filter) (*f.CommentGetAllRes, error) {
	res, err := s.storage.CommentS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CommentService) Update(ctx context.Context, req *f.CommentUpdate) (*f.Comment, error) {
	res, err := s.storage.CommentS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CommentService) Delete(ctx context.Context, id *f.GetByIdReq) (*f.Void, error) {
	_, err := s.storage.CommentS.Delete(id)

	return nil, err
}

func (s *CommentService) GetCommentByPost(ctx context.Context, filter *f.GetFilter)(*f.CommentGetAllRes, error){
	res, err := s.storage.CommentS.GetCommentByPost(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}