package posts

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (r *repository) CreateComment(ctx context.Context,  model posts.CommentModel) error{
	query := "INSERT INTO comments(user_id, post_id, comments, created_at, updated_at, created_by, updated_by) VALUES(?, ?, ?, ?, ? , ?, ?)"
	_, err := r.db.ExecContext(ctx, query, model.UserId, model.PostId, model.Comments, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil{
		return err
	}
	return nil
}

func (r *repository) GetCommentsByPostIdID(ctx context.Context, postId int64)([]posts.Comment, error){
	query := `SELECT c.id, c.user_id, c.comments, u.user_name
	FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id=?`

	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	response := make([]posts.Comment, 0)
	for rows.Next(){
		var comment posts.Comment
		err = rows.Scan(&comment.ID, &comment.UserId, &comment.CommentContent, &comment.Username)
		if err != nil{
			return []posts.Comment{}, err
		}
		response = append(response, posts.Comment{
			ID: comment.ID,
			UserId: comment.UserId,
			CommentContent: comment.CommentContent,
			Username: comment.Username,
		})
	}
	return response, nil
}


