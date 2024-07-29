package storage

import f "forum-service/genproto/forum"

type CategoryI interface{
	Create(*f.CategoryCreateReq)(*f.CategoryRes, error)
    GetById(*f.GetByIdReq)(*f.CategoryRes, error)
    GetAll(*f.Filter)(*f.CategoryGetAllRes, error)
    Update(*f.CategoryUpdate)(*f.CategoryRes, error)
    Delete(*f.GetByIdReq)(*f.Void, error)
}

type CommentI interface{
	Create(*f.CommentCreateReq)(*f.Comment, error)
    GetById(*f.GetByIdReq)(*f.CommentRes, error)
    GetAll(*f.Filter)(*f.CommentGetAllRes, error)
    Update(*f.CommentUpdate)(*f.Comment, error)
    Delete(*f.GetByIdReq)(*f.Void, error)
	GetCommentByPost(*f.GetFilter)(*f.CommentGetAllRes, error)
}

type PostTagI interface{
	Create(*f.PostTagsCreateReq)(*f.PostTags, error)
    GetById(*f.GetByIdReq)(*f.PostTagsRes, error)
    GetAll(*f.Filter)(*f.PostTagsGetAllRes, error)
    Update(*f.PostTagsUpdate)(*f.PostTags, error)
    Delete(*f.GetByIdReq)(*f.Void, error)
	GetPostByTag(*f.GetFilter) (*f.PostTagsGetAllRes, error)
	GetPopularTag(*f.Void)(*f.PopularTag, error)
}

type PostI interface{
	Create(*f.PostCreateReq)(*f.Post, error)
    GetById(*f.GetByIdReq)(*f.PostRes, error)
    GetAll(*f.Filter)(*f.PostGetAllRes, error)
    Update(*f.PostUpdate)(*f.Post, error)
    Delete(*f.GetByIdReq)(*f.Void, error)
	GetPostByUser(*f.GetFilter)(*f.PostGetAllRes, error)
    GetPostByCategory(*f.GetFilter)(*f.PostGetAllRes, error)
	SearchPost(*f.PostSearch)(*f.PostGetAllRes, error)
}

type TagI interface{
	Create(*f.TagCreateReq)(*f.TagRes, error)
    GetById(*f.GetByIdReq)(*f.TagRes, error)
    GetAll(*f.Filter)(*f.TagGetAllRes, error)
    Update(*f.TagUpdate)(*f.TagRes, error)
    Delete(*f.GetByIdReq)(*f.Void, error)
}