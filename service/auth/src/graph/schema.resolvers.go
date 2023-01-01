package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"auth/constant"
	autherror "auth/graph/error"
	"auth/graph/generated"
	"auth/graph/model"
	"auth/sql"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jaevor/go-nanoid"
)

// SignUpByUserName is the resolver for the signUpByUserName field.
func (r *mutationResolver) SignUpByUserName(ctx context.Context, userName string, pwd string) (*model.Result, error) {
	log.Printf("signup by user name: %s\n", userName)
	canonicID, err := nanoid.Standard(14)
	if err != nil {
		log.Printf("nanoid.Standard err: %v\n", err)
		return nil, err
	}
	userId := canonicID()
	accessToken, err := r.AuthUtil.SignToken(userId, time.Duration(r.Config.Access_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign access token err: %v\n", err)
		return nil, err
	}
	refreshToken, err := r.AuthUtil.SignToken(userId, time.Duration(r.Config.Refresh_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign refresh token err: %v\n", err)
		return nil, err
	}
	byteMd5Pwd := md5.Sum([]byte(pwd))
	md5Pwd := hex.EncodeToString(byteMd5Pwd[:])
	_, err = r.CrdbClient.Pool.Exec(ctx, sql.InsertAuth,
		userName,
		constant.AuthIdTypeUserName,
		md5Pwd,
		userId)
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		return nil, errors.New(autherror.UserExist)
	} else if err != nil {
		log.Printf("insert auth err: %v\n", err)
		return nil, err
	}
	result := &model.Result{
		Auth: &model.Auth{
			AuthID:     userName,
			AuthIDType: constant.AuthIdTypeUserName,
			UserID:     userId,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}

// SignInByUserName is the resolver for the signInByUserName field.
func (r *queryResolver) SignInByUserName(ctx context.Context, userName string, pwd string) (*model.Result, error) {
	log.Printf("signin by user name: %s\n", userName)
	row := r.CrdbClient.Pool.QueryRow(ctx, sql.QueryAuth, userName)
	var auth_id, auth_id_type, user_id, auth_id_type_username_pwd string
	err := row.Scan(&auth_id, &auth_id_type, &user_id, &auth_id_type_username_pwd)
	if err == pgx.ErrNoRows {
		log.Printf("query row, but no rows")
		return nil, errors.New(autherror.UserNotExist)
	} else if err != nil {
		log.Printf("query row err: %v\n", err)
		return nil, err
	}
	byteMd5Pwd := md5.Sum([]byte(pwd))
	md5Pwd := hex.EncodeToString(byteMd5Pwd[:])
	if md5Pwd != auth_id_type_username_pwd {
		log.Println("pwd wrong")
		return nil, errors.New(autherror.PwdWrong)
	}
	accessToken, err := r.AuthUtil.SignToken(user_id, time.Duration(r.Config.Access_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign access token err: %v\n", err)
		return nil, err
	}
	refreshToken, err := r.AuthUtil.SignToken(user_id, time.Duration(r.Config.Refresh_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign refresh token err: %v\n", err)
		return nil, err
	}
	result := &model.Result{
		Auth: &model.Auth{
			AuthID:     auth_id,
			AuthIDType: auth_id_type,
			UserID:     user_id,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, err
}

// RefreshToken is the resolver for the RefreshToken field.
func (r *queryResolver) RefreshToken(ctx context.Context, refreshToken string) (*model.Result, error) {
	//get user id from refresh token and check token not expired, then sign new access and refresh token
	user, err := r.AuthUtil.UserFromJwt(refreshToken)
	if err != nil {
		return nil, err
	}
	log.Printf("user(%v) refresh token\n", user)
	accessToken, err := r.AuthUtil.SignToken(user.Id, time.Duration(r.Config.Access_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign access token err: %v\n", err)
		return nil, err
	}
	newRefreshToken, err := r.AuthUtil.SignToken(user.Id, time.Duration(r.Config.Refresh_token_expire_seconds)*time.Second)
	if err != nil {
		log.Printf("sign refresh token err: %v\n", err)
		return nil, err
	}
	result := &model.Result{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}
	return result, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
