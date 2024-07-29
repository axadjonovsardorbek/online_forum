package postgres

import (
	// "database/sql"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cf "forum-service/config"

	"forum-service/config/logger"
	f "forum-service/genproto/forum"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func TestCategoryRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	em.CheckErr(err)

	repo := NewCategoryRepo(db, logger)

	req := &f.CategoryCreateReq{Name: "Test Category"}
	expectedID := "some-uuid"
	expectedCategory := &f.CategoryRes{Id: expectedID, Name: "Test Category"}

	mock.ExpectQuery(`INSERT INTO categories`).
		WithArgs(sqlmock.AnyArg(), req.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedID, req.Name))

	category, err := repo.Create(req)
	require.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryRepo_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	em.CheckErr(err)

	repo := NewCategoryRepo(db, logger)

	id := &f.GetByIdReq{Id: "10000000-0000-0000-0000-000000000002"}
	expectedCategory := &f.CategoryRes{Id: "10000000-0000-0000-0000-000000000002", Name: "Lifestyle"}

	mock.ExpectQuery(`SELECT id, name FROM categories WHERE deleted_at = 0 AND id = ?`).
		WithArgs(id.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(id.Id, expectedCategory.Name))

	category, err := repo.GetById(id)
	require.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	em.CheckErr(err)
	repo := NewCategoryRepo(db, logger)

	filter := &f.Filter{Limit: 10, Offset: 0}
	expectedCategories := &f.CategoryGetAllRes{
		Category: []*f.CategoryRes{
			{Id: "some-uuid-1", Name: "Test Category 1"},
			{Id: "some-uuid-2", Name: "Test Category 2"},
		},
	}

	mock.ExpectQuery(`SELECT COUNT\(1\) FROM categories WHERE deleted_at=0`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
	mock.ExpectQuery(`SELECT id, name FROM categories WHERE deleted_at = 0 LIMIT 10 OFFSET 0`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow("some-uuid-1", "Test Category 1").
			AddRow("some-uuid-2", "Test Category 2"))

	categories, err := repo.GetAll(filter)
	require.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
}

func TestCategoryRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	em.CheckErr(err)
	repo := NewCategoryRepo(db, logger)

	req := &f.CategoryUpdate{
		Id:       "some-uuid",
		Category: &f.CategoryCreateReq{Name: "Updated Category"},
	}
	expectedCategory := &f.CategoryRes{Id: "some-uuid", Name: "Updated Category"}

	mock.ExpectQuery(`UPDATE categories SET name = \?, updated_at = now\(\) WHERE deleted_at = 0 AND id = \? RETURNING id, name`).
		WithArgs(req.Category.Name, req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedCategory.Id, expectedCategory.Name))

	category, err := repo.Update(req)
	require.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	em.CheckErr(err)
	repo := NewCategoryRepo(db, logger)

	id := &f.GetByIdReq{Id: "some-uuid"}

	mock.ExpectExec(`UPDATE categories SET deleted_at=EXTRACT\(EPOCH FROM NOW\(\)\) WHERE id=\? AND deleted_at = 0`).
		WithArgs(id.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	void, err := repo.Delete(id)
	require.NoError(t, err)
	assert.NotNil(t, void)
}
