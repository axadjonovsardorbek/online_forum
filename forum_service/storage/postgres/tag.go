package postgres

import (
	"database/sql"
	"fmt"
	"forum-service/config/logger"
	f "forum-service/genproto/forum"

	"github.com/google/uuid"
)

type TagRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewTagRepo(db *sql.DB, logger *logger.Logger) *TagRepo {
	return &TagRepo{db: db, Logger: logger}
}

func (t *TagRepo) Create(req *f.TagCreateReq) (*f.TagRes, error) {

	id := uuid.New().String()
	tag := f.TagRes{}

	query := `
	INSERT INTO tags(
		id,
		name
	) VALUES ($1, $2)
	RETURNING
		id, 
		name
	`

	row := t.db.QueryRow(query, id, req.Name)

	err := row.Scan(
		&tag.Id,
		&tag.Name,
	)

	if err != nil {
		t.Logger.ERROR.Println("Error while creating tag: ", err)
		return nil, err
	}

	t.Logger.INFO.Println("Successfully created tag")

	return &tag, nil
}
func (t *TagRepo) GetById(id *f.GetByIdReq) (*f.TagRes, error) {

	tag := f.TagRes{}

	query := `
	SELECT 
		id, 
		name
	FROM 
		tags
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := t.db.QueryRow(query, id.Id)

	err := row.Scan(
		&tag.Id,
		&tag.Name,
	)

	if err != nil {
		t.Logger.ERROR.Println("Error while getting tag by id : ", err)
		return nil, err
	}

	t.Logger.INFO.Println("Successfully got tag")

	return &tag, nil
}
func (t *TagRepo) GetAll(filter *f.Filter) (*f.TagGetAllRes, error) {

	tags := f.TagGetAllRes{}

	query := `
	SELECT 
		id,
		name
	FROM 
		tags
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := t.db.QueryRow("SELECT COUNT(1) FROM tags WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		t.Logger.ERROR.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := t.db.Query(query, args...)

	if err != nil {
		t.Logger.ERROR.Println("Error while getting all tags : ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tag := f.TagRes{}

		err := rows.Scan(
			&tag.Id,
			&tag.Name,
		)

		if err != nil {
			t.Logger.ERROR.Println("Error while scanning all tags : ", err)
			return nil, err
		}

		tags.Tag = append(tags.Tag, &tag)
	}

	t.Logger.INFO.Println("Successfully fetched all tags")

	return &tags, nil
}

func (t *TagRepo) Update(req *f.TagUpdate) (*f.TagRes, error) { 
	
	tag := f.TagRes{}
	
	query := `
	UPDATE tags
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

	row := t.db.QueryRow(query, req.Tag.Name, req.Id)

	err := row.Scan(
		&tag.Id,
		&tag.Name,
	)

	if err != nil {
		t.Logger.ERROR.Println("Error while updating tag : ", err)
		return nil, err
	}

	t.Logger.INFO.Println("Successfully updated tag")

	return &tag, nil
}
func (t *TagRepo) Delete(id *f.GetByIdReq) (*f.Void, error)   { 
	
	void := f.Void{}

	query := `UPDATE tags SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	res, err := t.db.Exec(query, id.Id)
	if err != nil {
		t.Logger.ERROR.Println("Error while deleting tag : ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("tag with id %s not found", id.Id)
	}

	t.Logger.INFO.Println("Successfully deleted tag")

	return &void, nil
}
