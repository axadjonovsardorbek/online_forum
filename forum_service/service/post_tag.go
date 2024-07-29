package service

import (
	"context"
	f "forum-service/genproto/forum"
	st "forum-service/storage/postgres"
)

type PostTagsService struct {
	storage st.Storage
	f.UnimplementedPostTagsServiceServer
}

func NewPostTagsService(storage *st.Storage) *PostTagsService {
	return &PostTagsService{
		storage: *storage,
	}
}

func (s *PostTagsService) Create(ctx context.Context, req *f.PostTagsCreateReq) (*f.PostTags, error) {
	res, err := s.storage.PostTagS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostTagsService) GetById(ctx context.Context, id *f.GetByIdReq) (*f.PostTagsRes, error) {
	res, err := s.storage.PostTagS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostTagsService) GetAll(ctx context.Context, filter *f.Filter) (*f.PostTagsGetAllRes, error) {
	res, err := s.storage.PostTagS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostTagsService) Update(ctx context.Context, req *f.PostTagsUpdate) (*f.PostTags, error) {
	res, err := s.storage.PostTagS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostTagsService) Delete(ctx context.Context, id *f.GetByIdReq) (*f.Void, error) {
	_, err := s.storage.PostTagS.Delete(id)

	return nil, err
}

func (s *PostTagsService) GetPostByTag(ctx context.Context, filter *f.GetFilter)(*f.PostTagsGetAllRes, error){
	res, err := s.storage.PostTagS.GetPostByTag(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostTagsService) GetPopularTag(ctx context.Context, nothing *f.Void)(*f.PopularTag, error){
	res, err := s.storage.PostTagS.GetPopularTag(nothing)

	if err != nil {
		return nil, err
	}

	return res, nil
}