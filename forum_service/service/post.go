package service

import (
	"context"
	f "forum-service/genproto/forum"
	st "forum-service/storage/postgres"
)

type PostService struct {
	storage st.Storage
	f.UnimplementedPostServiceServer
}

func NewPostServiceService(storage *st.Storage) *PostService {
	return &PostService{
		storage: *storage,
	}
}

func (s *PostService) Create(ctx context.Context, req *f.PostCreateReq) (*f.Post, error) {
	res, err := s.storage.PostS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) GetById(ctx context.Context, id *f.GetByIdReq) (*f.PostRes, error) {
	res, err := s.storage.PostS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) GetAll(ctx context.Context, filter *f.Filter) (*f.PostGetAllRes, error) {
	res, err := s.storage.PostS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) Update(ctx context.Context, req *f.PostUpdate) (*f.Post, error) {
	res, err := s.storage.PostS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) Delete(ctx context.Context, id *f.GetByIdReq) (*f.Void, error) {
	_, err := s.storage.PostS.Delete(id)

	return nil, err
}

func (s *PostService) GetPostByUser(ctx context.Context,filter *f.GetFilter) (*f.PostGetAllRes, error){
	res, err := s.storage.PostS.GetPostByUser(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) GetPostByCategory(ctx context.Context,filter *f.GetFilter) (*f.PostGetAllRes, error){
	res, err := s.storage.PostS.GetPostByCategory(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostService) SearchPost(ctx context.Context,filter *f.PostSearch)(*f.PostGetAllRes, error){
	res, err := s.storage.PostS.SearchPost(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}