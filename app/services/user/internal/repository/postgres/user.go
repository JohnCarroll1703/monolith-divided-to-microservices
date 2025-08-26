package postgres

import (
	"context"
	"errors"
	"fmt"
	"monolith-divided-to-microservices/app/services/user/internal/model"
	"monolith-divided-to-microservices/app/services/user/internal/schema"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, email, username, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err = rows.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (id, username, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.Password, u.CreatedAt)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, filter schema.UserFilters) (
	[]model.User, error) {
	query := `SELECT id, email, username, created_at FROM users WHERE 1=1`
	args := []interface{}{}
	i := 1

	if filter.ID != "" {
		query += fmt.Sprintf(" AND id = $%d", i)
		args = append(args, filter.ID)
		i++
	}
	if filter.Username != "" {
		query += fmt.Sprintf(" AND username ILIKE $%d", i)
		args = append(args, "%"+filter.Username+"%")
		i++
	}
	if filter.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", i)
		args = append(args, "%"+filter.Email+"%")
		i++
	}
	if !filter.CreatedBefore.IsZero() {
		query += fmt.Sprintf(" AND created_at < $%d", i)
		args = append(args, filter.CreatedBefore)
		i++
	}
	if !filter.CreatedAfter.IsZero() {
		query += fmt.Sprintf(" AND created_at > $%d", i)
		args = append(args, filter.CreatedAfter)
		i++
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE email = $1`

	var user model.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &user, fmt.Errorf("user not found")
		}
		return &user, fmt.Errorf("failed to fetch user by email: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &user, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user by id: %w", err)
	}

	return &user, nil
}
