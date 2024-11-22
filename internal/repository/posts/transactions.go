package posts

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func execCreateUserLikesTransaction(r *repository, ctx context.Context, query string, userLikesModel posts.UserLikesModel)error{
	tx, err := r.db.BeginTx(ctx,nil)
	if err != nil{
		return nil
	}
	_, err = tx.ExecContext(ctx, query, userLikesModel.UserId, userLikesModel.PostId, userLikesModel.IsLiked, userLikesModel.CreatedAt, userLikesModel.UpdatedAt, userLikesModel.CreatedBy, userLikesModel.UpdatedBy)
	if err != nil{
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil{
		return err
	}
	return nil
}

func execUpdateUserLikesTransaction(r *repository, ctx context.Context, query string, userLikesModel posts.UserLikesModel)error{
	tx, err := r.db.BeginTx(ctx,nil)
	if err != nil{
		return nil
	}
	_, err = tx.ExecContext(ctx, query, userLikesModel.IsLiked, userLikesModel.UpdatedAt, userLikesModel.UpdatedBy, userLikesModel.UserId, userLikesModel.PostId)
	if err != nil{
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil{
		return err
	}
	return nil
}