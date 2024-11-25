package memberships

import (
	"context"
	"database/sql"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

func (r *repository) GetUser(ctx context.Context, email string, userName string, userID int64)(*memberships.UsersModel, error){
	var query string= "SELECT id, email, password, user_name, created_at, updated_at, created_by, updated_by FROM users WHERE email=? OR user_name=? OR id=?"
	var row = r.db.QueryRowContext(ctx,query, email, userName, userID)
	var user memberships.UsersModel
	
	err:= row.Scan(&user.ID, &user.Email, &user.Password, &user.UserName, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserForSignin(ctx context.Context, email, userName string)(*memberships.UsersModel, *memberships.UsersModel, error){
	queryByEmail:= "SELECT id, email, password, user_name, created_at, updated_at, created_by, updated_by FROM users WHERE email=?"
	rowsByEmail := r.db.QueryRowContext(ctx, queryByEmail, email)
	var responseByEmail memberships.UsersModel
	err := rowsByEmail.Scan(&responseByEmail.ID, &responseByEmail.Email, &responseByEmail.Password, &responseByEmail.UserName, &responseByEmail.CreatedAt, &responseByEmail.UpdatedAt, &responseByEmail.CreatedBy, &responseByEmail.UpdatedBy)
	
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil, nil
		}
		return nil, nil, err
	}

	queryByUserName := "SELECT id, email, password, user_name, created_at, updated_at, created_by, updated_by FROM users WHERE user_name=?"
	var rowByUserName *sql.Row = r.db.QueryRowContext(ctx, queryByUserName, userName)
	var responseByUserName memberships.UsersModel
	err = rowByUserName.Scan(&responseByUserName.ID,&responseByUserName.Email, &responseByUserName.Password, &responseByUserName.UserName, &responseByUserName.CreatedAt, &responseByUserName.UpdatedAt, &responseByUserName.CreatedBy, &responseByUserName.UpdatedBy)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil, nil
		}
		return nil,nil, err
	}
	
	return &responseByEmail, &responseByUserName, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UsersModel)error{
	query:= `INSERT INTO users(email, password, user_name, created_at, updated_at, created_by, updated_by)
	VALUES(?, ?, ?, ?, ?, ?, ?)`

	_, err:=r.db.ExecContext(ctx, query, model.Email, model.Password, model.UserName, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.CreatedBy)
	if err != nil{
		return err
	}

	return nil
}