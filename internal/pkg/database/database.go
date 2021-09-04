package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/diwise/api-notify/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Db interface {
	CreateSubscription(ctx context.Context, subscription *models.Subscription) error
	DeleteSubscription(ctx context.Context, subscriptionId string) error
	GetSubscriptionById(subscriptionId string) (*models.Subscription, error)
	GetSubscriptionsByIdOrType(ctx context.Context, id string, entityType string) ([]models.Subscription, error)
	ListSubscriptions(ctx context.Context, limit int) ([]models.Subscription, error)
	UpdateSubscription(ctx context.Context, subscription *models.Subscription) error
}

type myDB struct {
	pool *pgxpool.Pool
}

func NewDatabase(dbUrl string) (Db, error) {
	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %s", err.Error())
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pgxpool: %s", err.Error())
	}

	_, err = db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS Subscriptions (subscriptionId varchar(255) NOT NULL PRIMARY KEY, subscriptionData json NOT NULL)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscriptions table: %s", err.Error())
	}

	return myDB{
		pool: db,
	}, nil

}

func (db myDB) GetSubscriptionsByIdOrType(ctx context.Context, id string, entityType string) ([]models.Subscription, error) {
	rows, err := db.pool.Query(ctx, `
	
	select subscriptiondata from (
		select subscriptionid, subscriptiondata, json_array_elements(subscriptiondata -> 'entities') as entities
		from subscriptions
	) as subs
	where entities ->> 'type' = $1
	   or entities ->> 'id' = $2
		   
	`, entityType, id)

	if err != nil {
		return nil, err
	}

	return rowsToSubscriptions(rows), nil
}

func rowsToSubscriptions(rows pgx.Rows) []models.Subscription {
	subscriptions := make([]models.Subscription, 0)

	for rows.Next() {
		var str string
		var s models.Subscription

		if err := rows.Scan(&str); err == nil {
			if err := json.Unmarshal([]byte(str), &s); err == nil {
				subscriptions = append(subscriptions, s)
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

	return subscriptions
}

func (db myDB) ListSubscriptions(ctx context.Context, limit int) ([]models.Subscription, error) {

	rows, err := db.pool.Query(ctx, `select subscriptiondata from subscriptions`)
	if err != nil {
		return nil, err
	}

	return rowsToSubscriptions(rows), nil
}

func (db myDB) GetSubscriptionById(subscriptionId string) (*models.Subscription, error) {

	var data string
	err := db.pool.QueryRow(context.Background(), "select subscriptionData from subscriptions where subscriptionId=$1", subscriptionId).Scan(&data)
	if err != nil {
		return nil, err
	}

	sub := &models.Subscription{}

	err = json.Unmarshal([]byte(data), sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (db myDB) CreateSubscription(ctx context.Context, subscription *models.Subscription) error {
	data, err := json.Marshal(subscription)
	if err != nil {
		return err
	}

	_, err = db.pool.Exec(
		ctx,
		`insert into subscriptions (subscriptionid, subscriptiondata) values($1, $2) 
 		 on conflict (subscriptionid) do update set subscriptiondata=excluded.subscriptiondata`,
		subscription.Id,
		data)

	return err
}

func (db myDB) UpdateSubscription(ctx context.Context, subscription *models.Subscription) error {
	data, err := json.Marshal(subscription)
	if err != nil {
		return err
	}

	_, err = db.pool.Exec(
		ctx,
		`update subscriptions set subscriptiondata = $2 where subscriptionid = $1`,
		subscription.Id,
		data)

	return err
}

func (db myDB) DeleteSubscription(ctx context.Context, subscriptionId string) error {
	_, err := db.pool.Exec(ctx, `delete from subscriptions subscriptionid = $1`, subscriptionId)
	return err
}
