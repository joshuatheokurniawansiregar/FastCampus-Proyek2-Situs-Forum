package posts

import (
	"context"
	"log"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func(r *repository) CreateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error{
	var query string = `INSERT INTO user_likes(user_id, post_id, is_liked, created_at, updated_at, created_by, updated_by)
	VALUES(?,?,?,?,?,?,?)`
	err := execCreateUserLikesTransaction(r, ctx, query, userLikesModel)
	if err != nil{
		return err
	}
	return nil
}

func(r *repository) UpdateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error{
	var query string = `UPDATE user_likes SET is_liked=?, updated_at=?, updated_by=? WHERE user_id=? AND post_id=?`
	err := execUpdateUserLikesTransaction(r, ctx, query, userLikesModel)
	if err != nil{
		log.Fatal(err.Error())
		return err
	}
	return nil
}