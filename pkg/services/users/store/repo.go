package store

import (
	"context"
	"database/sql"
	"github.com/kott/go-service-example/pkg/services/users"
	"github.com/kott/go-service-example/pkg/utils/log"
)

const (
	selectUser = `select * from users where id = $1`
	selectManyUser = `select * from users order by created_at desc limit $1 offset $2`
	insertUser = `insert into users (id, username, mobile, created_at, updated_at) values ($1, $2, $3, now(), now())`
	updateUser = `update users set username = $1, mobile = $2, updated_at = now() where id = $3`
)

type UserRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) UserRepo {
	return UserRepo{DB: db}
}

func (r *UserRepo) Get(ctx context.Context, id string) (users.User, error) {
	var u users.User

	err := r.DB.QueryRow(selectUser, id).
		Scan(&u.ID, &u.Username, &u.Mobile, &u.CreatedAt, &u.UpdatedAt, &u.DisabledAt)
	if err != nil {
		log.Info(ctx, "select u error: %s", err.Error())
		return u, users.ErrUserNotFound
	}
	return u, nil
}

func (r *UserRepo) GetAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	ul := make([]users.User, 0)

	rows, err := r.DB.Query(selectManyUser, limit, offset)
	if err != nil {
		log.Warn(ctx, "query users error: %s", err.Error())
		return ul, users.ErrUserQuery
	}
	defer rows.Close()

	for rows.Next() {
		var u users.User
		err := rows.Scan(&u.ID, &u.Username, &u.Mobile, &u.CreatedAt, &u.UpdatedAt, &u.DisabledAt)
		if err != nil {
			log.Info(ctx, "scan users error: %s", err.Error())
			return ul, users.ErrUserQuery
		}
		ul = append(ul, u)
	}

	return ul, nil
}

func (r *UserRepo) Create(ctx context.Context, user users.UserCreateUpdate) (string, error) {
	return "", nil
}

func (r *UserRepo) Update(ctx context.Context, user users.UserCreateUpdate) error {

	return nil
}
