package posts

import (
	"context"

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

// func (r *repository) GetAllPosts(ctx context.Context, limit, offset int)(posts.GetAllPostsResponse, error){
// 	query:= `SELECT p.id, p.user_id, u.user_name, p.post_title, p.post_content, p.hashtag, p.created_at, p.updated_at, p.created_by, p.updated_by 
// 	FROM posts p JOIN users u ON p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

// 	var(
// 		response posts.GetAllPostsResponse
// 		rows *sql.Rows
// 		err error
// 	)
// 	response = posts.GetAllPostsResponse{}
// 	rows, err = r.db.QueryContext(ctx, query, )
// 	if err != nil{
// 		return response, err
// 	}
// 	defer rows.Close()
// 	var data []posts.Post = make([]posts.Post, 0)
// 	if rows.Next(){
// 		var postModel posts.PostModel
// 		var userName string
// 		err = rows.Scan(&postModel.ID, &postModel.UserId, userName, &postModel.PostTitle, &postModel.PostContent, &postModel.PostHashtags, &postModel.CreatedAt, &postModel.UpdatedAt, &postModel.CreatedBy, &postModel.UpdatedBy)
// 		data = append(data, posts.Post{
// 			ID: postModel.ID,
// 			UserId: postModel.UserId,
// 			Username: userName,
// 			PostTitle: postModel.PostTitle,
// 			PostContent: postModel.PostContent,
// 			PostHashTags: strings.Split(postModel.PostHashtags,","),
// 		})
// 	}
// 	response = posts.GetAllPostsResponse{
// 		Data: data,
// 		Pagination: ,
// 	}
// }