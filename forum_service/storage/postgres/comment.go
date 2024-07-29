package postgres

import (
	"database/sql"
	"fmt"
	"forum-service/client"
	"forum-service/config/logger"
	f "forum-service/genproto/forum"

	"github.com/google/uuid"
)

type CommentRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewCommentRepo(db *sql.DB, logger *logger.Logger) *CommentRepo {
	return &CommentRepo{db: db, Logger: logger}
}

func (c *CommentRepo) Create(req *f.CommentCreateReq) (*f.Comment, error) {

	comment := f.Comment{}
	id := uuid.New().String()

	query := `
	INSERT INTO comments(
		id,
		user_id,
		post_id,
		body
	) VALUES ($1, $2, $3, $4)
	RETURNING
		id,
		user_id,
		post_id,
		body	
	`

	row := c.db.QueryRow(query, id, req.UserId, req.PostId, req.Body)

	err := row.Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&comment.Body,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while creating coment: ", err)
		return nil, err
	}

	c.Logger.INFO.Println("Successfully created comment")

	return &comment, nil
}
func (c *CommentRepo) GetById(id *f.GetByIdReq) (*f.CommentRes, error) {

	res := f.CommentRes{
		User : &f.UserResp{},
		Post: &f.PostRes{
			User: &f.UserResp{},
			Category: &f.CategoryRes{},
		},
	}

	selectQuery := `
	SELECT
		c.id,
		c.body,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM comments as c
	JOIN posts as p ON c.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE c.id = $1 AND c.deleted_at = 0
	`

	row := c.db.QueryRow(selectQuery, id.Id)

	err := row.Scan(
		&res.Id,
		&res.Body,
		&res.Post.Id,
		&res.Post.Title,
		&res.Post.Body,
		&res.Post.Category.Id,
		&res.Post.Category.Name,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while retrieving comment details:", err)
		return nil, err
	}

	userComQuery := `SELECT user_id FROM comments WHERE id = $1 AND deleted_at = 0`

	row = c.db.QueryRow(userComQuery, res.Id)

	var userId string
	err = row.Scan(&userId)
	if err != nil {
		c.Logger.ERROR.Println("Error while getting user id:", err)
		return nil, err
	}

	usCom, err := client.GetUser(userId)

	if err != nil {
		c.Logger.ERROR.Println("Error while getting user:", err)
		return nil, err
	}

	res.User.Id = usCom.ID
	res.User.Email = usCom.Email
	res.User.Username = usCom.Username

	userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`

	row = c.db.QueryRow(userPostQuery, res.Post.Id)

	err = row.Scan(&userId)
	if err != nil {
		c.Logger.ERROR.Println("Error while getting user id:", err)
		return nil, err
	}

	usPost, err := client.GetUser(userId)

	if err != nil {
		c.Logger.ERROR.Println("Error while getting user:", err)
		return nil, err
	}

	res.Post.User.Id = usPost.ID
	res.Post.User.Email = usPost.Email
	res.Post.User.Username = usPost.Username

	return &res, nil
}
func (c *CommentRepo) GetAll(filter *f.Filter) (*f.CommentGetAllRes, error) {

	res := f.CommentGetAllRes{
		Comment: []*f.CommentRes{},
	}

	selectQuery := `
	SELECT
		c.id,
		c.body,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM comments as c
	JOIN posts as p ON c.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE c.deleted_at = 0 LIMIT $1 OFFSET $2
	`

	rows, err := c.db.Query(selectQuery, filter.Limit, filter.Offset)

	if err != nil {
		c.Logger.ERROR.Println("Error while retrieving comments:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resCom := f.CommentRes{
			User : &f.UserResp{},
			Post: &f.PostRes{
				User: &f.UserResp{},
				Category: &f.CategoryRes{},
			},
		}
		err := rows.Scan(
			&resCom.Id,
			&resCom.Body,
			&resCom.Post.Id,
			&resCom.Post.Title,
			&resCom.Post.Body,
			&resCom.Post.Category.Id,
			&resCom.Post.Category.Name,
		)

		if err != nil {
			c.Logger.ERROR.Println("Error while retrieving comment details:", err)
			return nil, err
		}
	
		userComQuery := `SELECT user_id FROM comments WHERE id = $1 AND deleted_at = 0`
	
		row := c.db.QueryRow(userComQuery, resCom.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usCom, err := client.GetUser(userId)
	
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resCom.User.Id = usCom.ID
		resCom.User.Email = usCom.Email
		resCom.User.Username = usCom.Username
	
		userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row = c.db.QueryRow(userPostQuery, resCom.Post.Id)
	
		err = row.Scan(&userId)
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usPost, err := client.GetUser(userId)
	
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resCom.Post.User.Id = usPost.ID
		resCom.Post.User.Email = usPost.Email
		resCom.Post.User.Username = usPost.Username
	
		res.Comment = append(res.Comment, &resCom)
	}

	return &res, nil
}
func (c *CommentRepo) Update(req *f.CommentUpdate) (*f.Comment, error) {

	comment := f.Comment{}

	query := `
	UPDATE comments
	SET
		user_id = $1,
		post_id = $2,
		body = $3,
		updated_at = now()
	WHERE
		id = $4
	AND 
		deleted_at = 0	
	`

	_, err := c.db.Exec(query, req.UpdateComment.UserId, req.UpdateComment.PostId, req.UpdateComment.Body, req.Id)

	if err != nil {
		c.Logger.ERROR.Println("Error while updating comment:", err)
		return nil, err
	}

	selectQuery := `
	SELECT 
		id,
		user_id,
		post_id,
		body
	FROM comments
	WHERE id = $1 AND deleted_at = 0
	`

	row := c.db.QueryRow(selectQuery, req.Id)

	err = row.Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&comment.Body,
	)

	if err != nil {
		c.Logger.ERROR.Println("Error while retrieving comment details:", err)
		return nil, err
	}

	return &comment, nil
}
func (c *CommentRepo) Delete(id *f.GetByIdReq) (*f.Void, error) {

	void := f.Void{}

	query := `UPDATE comments SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	res, err := c.db.Exec(query, id.Id)
	if err != nil {
		c.Logger.ERROR.Println("Error while deleting comment : ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("comment with id %s not found", id.Id)
	}

	c.Logger.INFO.Println("Successfully deleted comment")

	return &void, nil
}

func (c *CommentRepo) GetCommentByPost(filter *f.GetFilter)(*f.CommentGetAllRes, error){
	res := f.CommentGetAllRes{
		Comment: []*f.CommentRes{},
	}

	selectQuery := `
	SELECT
		c.id,
		c.body,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name
	FROM comments as c
	JOIN posts as p ON c.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	WHERE p.id = $1 AND c.deleted_at = 0 LIMIT $2 OFFSET $3
	`

	rows, err := c.db.Query(selectQuery,filter.Id, filter.Limit, filter.Offset)

	if err != nil {
		c.Logger.ERROR.Println("Error while retrieving comments:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		resCom := f.CommentRes{
			User : &f.UserResp{},
			Post: &f.PostRes{
				User: &f.UserResp{},
				Category: &f.CategoryRes{},
			},
		}
		err := rows.Scan(
			&resCom.Id,
			&resCom.Body,
			&resCom.Post.Id,
			&resCom.Post.Title,
			&resCom.Post.Body,
			&resCom.Post.Category.Id,
			&resCom.Post.Category.Name,
		)

		if err != nil {
			c.Logger.ERROR.Println("Error while retrieving comment details:", err)
			return nil, err
		}
	
		userComQuery := `SELECT user_id FROM comments WHERE id = $1 AND deleted_at = 0`
	
		row := c.db.QueryRow(userComQuery, resCom.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usCom, err := client.GetUser(userId)
	
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resCom.User.Id = usCom.ID
		resCom.User.Email = usCom.Email
		resCom.User.Username = usCom.Username
	
		userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row = c.db.QueryRow(userPostQuery, resCom.Post.Id)
	
		err = row.Scan(&userId)
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usPost, err := client.GetUser(userId)
	
		if err != nil {
			c.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resCom.Post.User.Id = usPost.ID
		resCom.Post.User.Email = usPost.Email
		resCom.Post.User.Username = usPost.Username
	
		res.Comment = append(res.Comment, &resCom)
	}

	return &res, nil
}