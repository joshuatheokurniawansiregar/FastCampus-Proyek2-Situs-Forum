package posts

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (r *repository) CreateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error{
	var query string = `INSERT INTO user_likes(user_id, post_id, is_liked, created_at, updated_at, created_by, updated_by)
	VALUES(?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, userLikesModel.UserId, userLikesModel.PostId, userLikesModel.IsLiked, userLikesModel.CreatedAt, userLikesModel.UpdatedAt, userLikesModel.CreatedBy, userLikesModel.UpdatedBy)
	if err != nil{
		return err
	}
	return nil
}

func (r *repository) UpdateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error{
	var query string = `UPDATE user_likes SET is_liked=? WHERE user_id=?`
	_, err := r.db.ExecContext(ctx, query, userLikesModel.IsLiked, userLikesModel.UserId)
	if err != nil{
		return err
	}
	return nil
}