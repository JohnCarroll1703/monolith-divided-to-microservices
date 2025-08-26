package postgres

import (
	"context"
	"errors"
	"monolith-divided-to-microservices/app/services/item/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository struct {
	db *pgxpool.Pool
}

func NewItem(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]model.Item, error) {
	query := `SELECT id, name, description, category_id, country_id, stock, price, added_at FROM items`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.CategoryID, &item.CountryID,
			&item.Stock, &item.Price, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *model.Item) error {
	query := `INSERT INTO items (id, name, description, category_id, country_id, stock, price, added_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(ctx, query, item.ID, item.Name, item.Description, item.CategoryID,
		item.CountryID, item.Stock, item.Price, item.AddedAt)
	return err
}

func (r *ItemRepository) GetItemByID(ctx context.Context, id string) (*model.Item, error) {
	query := `SELECT id, name, description, category_id, country_id, 
	stock, price, added_at FROM items WHERE id=$1`
	row := r.db.QueryRow(ctx, query, id)
	var item model.Item
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CategoryID,
		&item.CountryID, &item.Stock, &item.Price, &item.AddedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}
