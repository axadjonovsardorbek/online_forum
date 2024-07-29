package postgres

import (
	"database/sql"
	"fmt"
	"forum-service/config/logger"
	"forum-service/client"
	f "forum-service/genproto/forum"

	"github.com/google/uuid"
)

type PostRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewPostRepo(db *sql.DB, logger *logger.Logger) *PostRepo {
	return &PostRepo{db: db, Logger: logger}
}

func (p *PostRepo) Create(req *f.PostCreateReq) (*f.Post, error) {

	id := uuid.New().String()
	post := f.Post{}

	query := `
	INSERT INTO posts(
		id,
		user_id,
		title,
		body,
		category_id
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING
		id,
		user_id,
		title,
		body,
		category_id
	`

	row := p.db.QueryRow(query, id, req.UserId, req.Title, req.Body, req.CategoryId)

	err := row.Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
		&post.CategoryId,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while creating post: ", err)
		return nil, err
	}

	p.Logger.INFO.Println("Successfully created post")

	return &post, nil
}

func (p *PostRepo) GetById(id *f.GetByIdReq) (*f.PostRes, error) {

	res := f.PostRes{
		User: &f.UserResp{},
		Category: &f.CategoryRes{},
	}

	selectQuery := `
	SELECT 
		p.id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM posts as p
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE p.id = $1 AND p.deleted_at = 0			
	`

	row := p.db.QueryRow(selectQuery, id.Id)

	err := row.Scan(
		&res.Id,
		&res.Title,
		&res.Body,
		&res.Category.Id,
		&res.Category.Name,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post details:", err)
		return nil, err
	}

	userQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`

	row = p.db.QueryRow(userQuery, res.Id)

	var userId string
	err = row.Scan(&userId)
	if err != nil {
		p.Logger.ERROR.Println("Error while getting user id:", err)
		return nil, err
	}

	us, err := client.GetUser(userId)

	if err != nil {
		p.Logger.ERROR.Println("Error while getting user:", err)
		return nil, err
	}

	res.User.Id = us.ID
	res.User.Email = us.Email
	res.User.Username = us.Username

	return &res, nil
}

func (p *PostRepo) GetAll(filter *f.Filter) (*f.PostGetAllRes, error) {

	res := f.PostGetAllRes{
		Post: []*f.PostRes{},
	}

	selectQuery := `
	SELECT 
		p.id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM posts as p
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE p.deleted_at = 0 LIMIT $1 OFFSET $2	
	`

	rows, err := p.db.Query(selectQuery, filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving posts:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resPost := f.PostRes{
			User: &f.UserResp{},
			Category: &f.CategoryRes{},
		}

		err := rows.Scan(
			&resPost.Id,
			&resPost.Title,
			&resPost.Body,
			&resPost.Category.Id,
			&resPost.Category.Name,
		)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving posts details:", err)
			return nil, err
		}
	
		userQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row := p.db.QueryRow(userQuery, resPost.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		us, err := client.GetUser(userId)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPost.User.Id = us.ID
		resPost.User.Email = us.Email
		resPost.User.Username = us.Username

		res.Post = append(res.Post, &resPost)
	}

	return &res, nil
}

func (p *PostRepo) Update(req *f.PostUpdate) (*f.Post, error) { 

	post := f.Post{}

	query := `
	UPDATE posts
	SET
		user_id = $1, 
		title = $2,
		body = $3,
		category_id = $4,
		updated_at = now()
	WHERE 
		id = $5
	AND
		deleted_at = 0
	`

	_, err := p.db.Exec(query, req.UpdatePost.UserId, req.UpdatePost.Title, req.UpdatePost.Body, req.UpdatePost.CategoryId, req.Id)

	if err != nil {
		p.Logger.ERROR.Println("Error while updating post:", err)
		return nil, err
	}

	selectQuery := `
	SELECT 
		id,
		user_id,
		title,
		body,
		category_id
	FROM posts
	WHERE id = $1 AND deleted_at = 0
	`

	row := p.db.QueryRow(selectQuery, req.Id)

	err = row.Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
		&post.CategoryId,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post details:", err)
		return nil, err
	}
			
	return &post, nil 
}

func (p *PostRepo) Delete(id *f.GetByIdReq) (*f.Void, error) {
	
	void := f.Void{}

	query := `UPDATE posts SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	res, err := p.db.Exec(query, id.Id)
	if err != nil {
		p.Logger.ERROR.Println("Error while deleting post : ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("post with id %s not found", id.Id)
	}

	p.Logger.INFO.Println("Successfully deleted post")

	return &void, nil
}

func (p *PostRepo) GetPostByUser(filter *f.GetFilter)(*f.PostGetAllRes, error){
	res := f.PostGetAllRes{
		Post: []*f.PostRes{},
	}

	selectQuery := `
	SELECT 
		p.id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM posts as p
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE p.user_id = $1 AND p.deleted_at = 0 LIMIT $2 OFFSET $3	
	`

	rows, err := p.db.Query(selectQuery,filter.Id, filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving posts:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resPost := f.PostRes{
			User: &f.UserResp{},
			Category: &f.CategoryRes{},
		}

		err := rows.Scan(
			&resPost.Id,
			&resPost.Title,
			&resPost.Body,
			&resPost.Category.Id,
			&resPost.Category.Name,
		)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving posts details:", err)
			return nil, err
		}
	
		us, err := client.GetUser(filter.Id)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPost.User.Id = us.ID
		resPost.User.Email = us.Email
		resPost.User.Username = us.Username

		res.Post = append(res.Post, &resPost)
	}

	return &res, nil
}

func (p *PostRepo) GetPostByCategory(filter *f.GetFilter)(*f.PostGetAllRes, error){
	res := f.PostGetAllRes{
		Post: []*f.PostRes{},
	}

	selectQuery := `
	SELECT 
		p.id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM posts as p
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE ct.id = $1 AND p.deleted_at = 0 LIMIT $2 OFFSET $3	
	`

	rows, err := p.db.Query(selectQuery,filter.Id, filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving posts:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resPost := f.PostRes{
			User: &f.UserResp{},
			Category: &f.CategoryRes{},
		}

		err := rows.Scan(
			&resPost.Id,
			&resPost.Title,
			&resPost.Body,
			&resPost.Category.Id,
			&resPost.Category.Name,
		)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving posts details:", err)
			return nil, err
		}
	
		userQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row := p.db.QueryRow(userQuery, resPost.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}

		us, err := client.GetUser(userId)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPost.User.Id = us.ID
		resPost.User.Email = us.Email
		resPost.User.Username = us.Username

		res.Post = append(res.Post, &resPost)
	}

	return &res, nil
}

func (p *PostRepo) SearchPost(filter *f.PostSearch)(*f.PostGetAllRes, error){
	res := f.PostGetAllRes{
		Post: []*f.PostRes{},
	}

	selectQuery := `
	SELECT 
		p.id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM posts as p
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE 
		LOWER(p.title) LIKE LOWER($1)
	AND 
		LOWER(p.body) LIKE LOWER($2)
	AND 
		p.deleted_at = 0 LIMIT $3 OFFSET $4	
	`

	rows, err := p.db.Query(selectQuery,"%" + filter.Title + "%", "%" + filter.Body + "%", filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving posts:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resPost := f.PostRes{
			User: &f.UserResp{},
			Category: &f.CategoryRes{},
		}

		err := rows.Scan(
			&resPost.Id,
			&resPost.Title,
			&resPost.Body,
			&resPost.Category.Id,
			&resPost.Category.Name,
		)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving posts details:", err)
			return nil, err
		}
	
		userQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row := p.db.QueryRow(userQuery, resPost.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}

		us, err := client.GetUser(userId)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPost.User.Id = us.ID
		resPost.User.Email = us.Email
		resPost.User.Username = us.Username

		res.Post = append(res.Post, &resPost)
	}

	return &res, nil
}