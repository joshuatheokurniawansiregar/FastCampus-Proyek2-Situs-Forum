package posts

import (
	"context"
	"database/sql"
	"strings"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel)error{
	var query string = "INSERT INTO posts(user_id, post_title, post_content, hashtag, created_at, updated_at, created_by, updated_by) VALUES(?,?,?,?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query, &model.UserId, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.CreatedBy)
	if err != nil{
		return err
	}
	return nil
}

func (r *repository) GetAllPosts(ctx context.Context, limit, offset int64)(posts.GetAllPostsResponse, error){
	query:= `SELECT p.id, p.user_id, u.user_name, p.post_title, p.post_content, p.hashtag, p.created_at, p.updated_at, p.created_by, p.updated_by FROM posts p INNER JOIN users u ON p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`
	var response posts.GetAllPostsResponse
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil{
		return response, err
	}
	defer rows.Close()
	var data []posts.Post = make([]posts.Post, 0)
	for rows.Next(){
		var postModel posts.PostModel
		var userName string
		err = rows.Scan(&postModel.ID, &postModel.UserId, &userName, &postModel.PostTitle, &postModel.PostContent, &postModel.PostHashtags, &postModel.CreatedAt, &postModel.UpdatedAt, &postModel.CreatedBy, &postModel.UpdatedBy)
		if err != nil{
			return response, err
		}
		data = append(data, posts.Post{
			ID: postModel.ID,
			UserId: postModel.UserId,
			Username: userName,
			PostTitle: postModel.PostTitle,
			PostContent: postModel.PostContent,
			PostHashTags: strings.Split(postModel.PostHashtags,","),
		})
	}
	response.Data = data
	response.Pagination = posts.Pagination{
		Limit: limit,
		Offset: offset,
	}
	return response, nil
}

func(r *repository) GetPostById(ctx context.Context, postId int64)(posts.Post, error){
	var(
		query string
		userName string
		post posts.Post
		postModel posts.PostModel
		row *sql.Row
		err error
	)
	query = `SELECT p.id, p.user_id, u.user_name, p.post_title, p.post_content, p.hashtag, p.created_at FROM posts p JOIN users u WHERE p.id=?`
	row = r.db.QueryRowContext(ctx, query, postId)
	err = row.Scan(&postModel.ID, &postModel.UserId, &userName, &postModel.PostTitle, &postModel.PostContent, &postModel.PostHashtags, &post.CreatedAt)
	if err != nil{
		return post, err
	}
	post.ID = postModel.ID
	post.UserId = postModel.UserId
	post.Username = userName
	post.PostTitle = postModel.PostTitle
	post.PostContent = postModel.PostContent
	post.PostHashTags = strings.Split(postModel.PostHashtags, ",")
	post.CreatedAt = postModel.CreatedAt
	return post, nil
}