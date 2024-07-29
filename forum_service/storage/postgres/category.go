package postgres

import (
	"database/sql"
	"fmt"
	"forum-service/config/logger"

	f "forum-service/genproto/forum"

	"github.com/google/uuid"
)

type CategoryRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewCategoryRepo(db *sql.DB, logger *logger.Logger) *CategoryRepo {
	return &CategoryRepo{db: db, Logger: logger}
}

func (c *CategoryRepo) Create(req *f.CategoryCreateReq) (*f.CategoryRes, error) {

	id := uuid.New().String()
	category := f.CategoryRes{}

	query := `
	INSERT INTO categories(
		id,
		name
	) VALUES ($1, $2)
	RETURNING
		id, 
		name
	`

	row := c.db.QueryRow(query, id, req.Name)

	err := row.Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while creating category: ", err)
		return nil, err
	}

	c.Logger.INFO.Println("Successfully created category")
	return &category, nil
}

func (c *CategoryRepo) GetById(id *f.GetByIdReq) (*f.CategoryRes, error) {

	category := f.CategoryRes{}

	query := `
	SELECT 
		id, 
		name
	FROM 
		categories
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := c.db.QueryRow(query, id.Id)

	err := row.Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while getting category by id : ", err)
		return nil, err
	}

	c.Logger.INFO.Println("Successfully got category")

	return &category, nil
}
func (c *CategoryRepo) GetAll(filter *f.Filter) (*f.CategoryGetAllRes, error) {

	categories := f.CategoryGetAllRes{}

	query := `
	SELECT 
		id,
		name
	FROM 
		categories
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := c.db.QueryRow("SELECT COUNT(1) FROM categories WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		c.Logger.ERROR.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := c.db.Query(query, args...)

	if err != nil {
		c.Logger.ERROR.Println("Error while getting all categories : ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		category := f.CategoryRes{}

		err := rows.Scan(
			&category.Id,
			&category.Name,
		)

		if err != nil {
			c.Logger.ERROR.Println("Error while scanning all categories : ", err)
			return nil, err
		}

		categories.Category = append(categories.Category, &category)
	}

	c.Logger.INFO.Println("Successfully fetched all categories")

	return &categories, nil
}

func (c *CategoryRepo) Update(req *f.CategoryUpdate) (*f.CategoryRes, error) { 

	category := f.CategoryRes{}
	
	query := `
	UPDATE categories
	SET 
		name = $1,
		updated_at = now()
	WHERE
		deleted_at = 0
	AND 
		id = $2	
	RETURNING
		id,
		name
	`

	row := c.db.QueryRow(query, req.Category.Name, req.Id)

	err := row.Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while updating category : ", err)
		return nil, err
	}

	c.Logger.INFO.Println("Successfully updated category")

	return &category, nil
}
func (c *CategoryRepo) Delete(id *f.GetByIdReq) (*f.Void, error)             { 
	
	void := f.Void{}

	query := `UPDATE categories SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	res, err := c.db.Exec(query, id.Id)
	if err != nil {
		c.Logger.ERROR.Println("Error while deleting category : ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("category with id %s not found", id.Id)
	}

	c.Logger.INFO.Println("Successfully deleted category")

	return &void, nil
}
