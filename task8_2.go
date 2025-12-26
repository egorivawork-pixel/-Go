package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) List(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, "select id, name, age from users order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UserRepo) Get(ctx context.Context, id int) (User, bool, error) {
	var u User
	err := r.db.QueryRowContext(ctx, "select id, name, age from users where id=$1", id).
		Scan(&u.ID, &u.Name, &u.Age)
	if err == sql.ErrNoRows {
		return User{}, false, nil
	}
	if err != nil {
		return User{}, false, err
	}
	return u, true, nil
}

func (r *UserRepo) Create(ctx context.Context, u User) (User, error) {
	err := r.db.QueryRowContext(ctx,
		"insert into users(name, age) values($1,$2) returning id",
		u.Name, u.Age).Scan(&u.ID)
	return u, err
}

func (r *UserRepo) Update(ctx context.Context, id int, u User) (bool, error) {
	res, err := r.db.ExecContext(ctx, "update users set name=$1, age=$2 where id=$3", u.Name, u.Age, id)
	if err != nil {
		return false, err
	}
	aff, _ := res.RowsAffected()
	return aff > 0, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int) (bool, error) {
	res, err := r.db.ExecContext(ctx, "delete from users where id=$1", id)
	if err != nil {
		return false, err
	}
	aff, _ := res.RowsAffected()
	return aff > 0, nil
}
