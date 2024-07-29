package service

import (
	"context"
	f "forum-service/genproto/forum"
	st "forum-service/storage/postgres"
)

type CategoryService struct {
	storage st.Storage
	f.UnimplementedCategoryServiceServer
}

func NewCategoryService(storage *st.Storage) *CategoryService {
	return &CategoryService{
		storage: *storage,
	}
}

func (s *CategoryService) Create(ctx context.Context, req *f.CategoryCreateReq) (*f.CategoryRes, error) {
	res, err := s.storage.CategoryS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) GetById(ctx context.Context, id *f.GetByIdReq) (*f.CategoryRes, error) {
	res, err := s.storage.CategoryS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) GetAll(ctx context.Context, filter *f.Filter) (*f.CategoryGetAllRes, error) {
	res, err := s.storage.CategoryS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) Update(ctx context.Context, req *f.CategoryUpdate) (*f.CategoryRes, error) {
	res, err := s.storage.CategoryS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryService) Delete(ctx context.Context, id *f.GetByIdReq) (*f.Void, error) {
	_, err := s.storage.CategoryS.Delete(id)

	return nil, err
}
