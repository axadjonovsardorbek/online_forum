package postgres

import (
	"database/sql"
	"fmt"
	"forum-service/client"
	"forum-service/config/logger"
	f "forum-service/genproto/forum"

	"github.com/google/uuid"
)

type PostTagRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewPostTagRepo(db *sql.DB, logger *logger.Logger) *PostTagRepo {
	return &PostTagRepo{db: db, Logger: logger}
}

func (p *PostTagRepo) Create(req *f.PostTagsCreateReq) (*f.PostTags, error) {

	postTag := f.PostTags{}
	id := uuid.New().String()

	query := `
	INSERT INTO post_tags(
		id,
		post_id,
		tag_id
	) VALUES ($1, $2, $3)
	RETURNING
		id,
		post_id,
		tag_id
	`

	row := p.db.QueryRow(query, id, req.PostId, req.TagId)

	err := row.Scan(
		&postTag.Id,
		&postTag.PostId,
		&postTag.TagId,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while creating post_tag: ", err)
		return nil, err
	}

	p.Logger.INFO.Println("Successfully created post_tag")

	return &postTag, nil
}
func (p *PostTagRepo) GetById(id *f.GetByIdReq) (*f.PostTagsRes, error) {

	res := f.PostTagsRes{
		Post: &f.PostRes{
			User:     &f.UserResp{},
			Category: &f.CategoryRes{},
		},
		Tag: &f.TagRes{},
	}

	selectQuery := `
	SELECT 
		pt.id,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name,
		tg.id as tag_id,
		tg.name
	FROM post_tags as pt
	JOIN posts as p ON pt.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	JOIN tags as tg ON pt.tag_id = tg.id AND tg.deleted_at = 0
	WHERE pt.id = $1 AND pt.deleted_at = 0
	`

	row := p.db.QueryRow(selectQuery, id.Id)

	err := row.Scan(
		&res.Id,
		&res.Post.Id,
		&res.Post.Title,
		&res.Post.Body,
		&res.Post.Category.Id,
		&res.Post.Category.Name,
		&res.Tag.Id,
		&res.Tag.Name,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post_tag details:", err)
		return nil, err
	}

	userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`

	row = p.db.QueryRow(userPostQuery, res.Post.Id)

	var userId string
	err = row.Scan(&userId)
	if err != nil {
		p.Logger.ERROR.Println("Error while getting user id:", err)
		return nil, err
	}

	usPost, err := client.GetUser(userId)

	if err != nil {
		p.Logger.ERROR.Println("Error while getting user:", err)
		return nil, err
	}

	res.Post.User.Id = usPost.ID
	res.Post.User.Email = usPost.Email
	res.Post.User.Username = usPost.Username

	return &res, nil
}

func (p *PostTagRepo) GetAll(filter *f.Filter) (*f.PostTagsGetAllRes, error) {

	res := f.PostTagsGetAllRes{
		Posttag: []*f.PostTagsRes{},
	}

	selectQuery := `
	SELECT 
		pt.id,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name,
		tg.id as tag_id,
		tg.name
	FROM post_tags as pt
	JOIN posts as p ON pt.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	JOIN tags as tg ON pt.tag_id = tg.id AND tg.deleted_at = 0
	WHERE pt.deleted_at = 0 LIMIT $1 OFFSET $2
	`

	rows, err := p.db.Query(selectQuery, filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post_tags:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		resPos := f.PostTagsRes{
			Post: &f.PostRes{
				User:     &f.UserResp{},
				Category: &f.CategoryRes{},
			},
			Tag: &f.TagRes{},
		}

		err := rows.Scan(
			&resPos.Id,
			&resPos.Post.Id,
			&resPos.Post.Title,
			&resPos.Post.Body,
			&resPos.Post.Category.Id,
			&resPos.Post.Category.Name,
			&resPos.Tag.Id,
			&resPos.Tag.Name,
		)
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving post_tag details:", err)
			return nil, err
		}
	
		userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row := p.db.QueryRow(userPostQuery, resPos.Post.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usPost, err := client.GetUser(userId)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPos.Post.User.Id = usPost.ID
		resPos.Post.User.Email = usPost.Email
		resPos.Post.User.Username = usPost.Username

		res.Posttag = append(res.Posttag, &resPos)
	}

	return &res, nil
}

func (p *PostTagRepo) Update(req *f.PostTagsUpdate) (*f.PostTags, error) {

	postTag := f.PostTags{}

	query := `
	UPDATE post_tags
	SET
		post_id = $1,
		tag_id = $2,
		updated_at = now()
	WHERE
		id = $3
	AND 
		deleted_at = 0
	`

	_, err := p.db.Exec(query, req.Posttag.PostId, req.Posttag.TagId, req.Id)

	if err != nil {
		p.Logger.ERROR.Println("Error while updating post_tag:", err)
		return nil, err
	}

	selectQuery := `
	SELECT 
		id,
		post_id,
		tag_id
	FROM post_tags
	WHERE id = $1 AND deleted_at = 0
	`

	row := p.db.QueryRow(selectQuery, req.Id)

	err = row.Scan(
		&postTag.Id,
		&postTag.PostId,
		&postTag.TagId,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post_tag details:", err)
		return nil, err
	}

	return &postTag, nil
}

func (p *PostTagRepo) Delete(id *f.GetByIdReq) (*f.Void, error) {
	void := f.Void{}

	query := `UPDATE post_tags SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	res, err := p.db.Exec(query, id.Id)
	if err != nil {
		p.Logger.ERROR.Println("Error while deleting post tag : ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("post tag with id %s not found", id.Id)
	}

	p.Logger.INFO.Println("Successfully deleted post tag")

	return &void, nil
}

func (p *PostTagRepo) GetPostByTag(filter *f.GetFilter) (*f.PostTagsGetAllRes, error){
	
	res := f.PostTagsGetAllRes{
		Posttag: []*f.PostTagsRes{},
	}

	selectQuery := `
	SELECT 
		pt.id,
		p.id as post_id,
		p.title,
		p.body,
		ct.id as category_id,
		ct.name,
		tg.id as tag_id,
		tg.name
	FROM post_tags as pt
	JOIN posts as p ON pt.post_id = p.id AND p.deleted_at = 0
	JOIN categories as ct ON p.category_id = ct.id AND ct.deleted_at = 0
	JOIN tags as tg ON pt.tag_id = tg.id AND tg.deleted_at = 0
	WHERE pt.tag_id = $1 AND pt.deleted_at = 0 LIMIT $2 OFFSET $3
	`

	rows, err := p.db.Query(selectQuery,filter.Id, filter.Limit, filter.Offset)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving post_tags:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		resPos := f.PostTagsRes{
			Post: &f.PostRes{
				User:     &f.UserResp{},
				Category: &f.CategoryRes{},
			},
			Tag: &f.TagRes{},
		}

		err := rows.Scan(
			&resPos.Id,
			&resPos.Post.Id,
			&resPos.Post.Title,
			&resPos.Post.Body,
			&resPos.Post.Category.Id,
			&resPos.Post.Category.Name,
			&resPos.Tag.Id,
			&resPos.Tag.Name,
		)
		if err != nil {
			p.Logger.ERROR.Println("Error while retrieving post_tag details:", err)
			return nil, err
		}
	
		userPostQuery := `SELECT user_id FROM posts WHERE id = $1 AND deleted_at = 0`
	
		row := p.db.QueryRow(userPostQuery, resPos.Post.Id)
	
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user id:", err)
			return nil, err
		}
	
		usPost, err := client.GetUser(userId)
	
		if err != nil {
			p.Logger.ERROR.Println("Error while getting user:", err)
			return nil, err
		}
	
		resPos.Post.User.Id = usPost.ID
		resPos.Post.User.Email = usPost.Email
		resPos.Post.User.Username = usPost.Username

		res.Posttag = append(res.Posttag, &resPos)
	}

	return &res, nil
}

func (p *PostTagRepo) GetPopularTag(nth *f.Void)(*f.PopularTag, error){

	res := f.PopularTag{}

	selectQuery := `
	SELECT 
		COUNT(t.id) as count,
		t.id,
		t.name
	FROM post_tags as pt
	JOIN tags as t ON t.id = pt.tag_id AND t.deleted_at = 0
	WHERE pt.deleted_at = 0
	GROUP by t.id
	ORDER by count DESC
	LIMIT 1	
	`
	row := p.db.QueryRow(selectQuery)

	err := row.Scan(
		&res.Count,
		&res.Id,
		&res.Name,
	)

	if err != nil {
		p.Logger.ERROR.Println("Error while retrieving tag details:", err)
		return nil, err
	}

	return &res, nil
}