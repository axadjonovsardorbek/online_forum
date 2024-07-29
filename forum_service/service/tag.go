package service

import (
	"context"
	f "forum-service/genproto/forum"
	st "forum-service/storage/postgres"
)

type TagService struct {
	storage st.Storage
	f.UnimplementedTagServiceServer
}

func NewTagService(storage *st.Storage) *TagService {
	return &TagService{
		storage: *storage,
	}
}

func (s *TagService) Create(ctx context.Context, req *f.TagCreateReq) (*f.TagRes, error) {
	res, err := s.storage.TagS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TagService) GetById(ctx context.Context, id *f.GetByIdReq) (*f.TagRes, error) {
	res, err := s.storage.TagS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TagService) GetAll(ctx context.Context, filter *f.Filter) (*f.TagGetAllRes, error) {
	res, err := s.storage.TagS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TagService) Update(ctx context.Context, req *f.TagUpdate) (*f.TagRes, error) {
	res, err := s.storage.TagS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TagService) Delete(ctx context.Context, id *f.GetByIdReq) (*f.Void, error) {
	_, err := s.storage.TagS.Delete(id)

	return nil, err
}
