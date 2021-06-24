package database

import (
	"context"
	"encoding/json"

	"github.com/diwise/api-notify/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

type Db struct {
	pool *pgxpool.Pool
}

func NewDatabase(dbUrl string) Db {
	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		panic(err.Error())
	}

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err.Error())
	}

	return Db{
		pool: db,
	}
}

func (db *Db) GetSubscriptionsByIdOrType(ctx context.Context, id string, entityType string) []models.Subscription {
	if rows, err := db.pool.Query(ctx, `
	
	select subscriptiondata from (
		select subscriptionid, subscriptiondata, json_array_elements(subscriptiondata -> 'entities') as entities
		from subscriptions
	) as subs
	where entities ->> 'type' = $1
	   or entities ->> 'id' = $2
		   
	`, entityType, id); err == nil {
		return rowsToSubscriptions(rows)
	} else {
		return nil
	}
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

func (db *Db) ListSubscriptions(ctx context.Context, limit int) ([]models.Subscription, error) {

	if rows, err := db.pool.Query(ctx, `select subscriptiondata from subscriptions`); err == nil {

		s := rowsToSubscriptions(rows)

		return s, nil
	} else {
		return nil, rows.Err()
	}
}

func (db *Db) GetSubscriptionById(subscriptionId string) (models.Subscription, error) {
	var data string
	var s *models.Subscription

	err := db.pool.QueryRow(context.Background(), "select subscriptionData from subscriptions where subscriptionId=$1", subscriptionId).Scan(&data)

	if err == nil && len(data) > 0 {
		if err := json.Unmarshal([]byte(data), &s); err == nil {
			return *s, nil
		} else {
			return models.Subscription{}, err
		}
	} else {
		return models.Subscription{}, err
	}

}

func (db *Db) CreateSubscription(ctx context.Context, subscription *models.Subscription) error {
	if data, err := json.Marshal(subscription); err == nil {
		id := subscription.Id
		if _, err := db.pool.Exec(ctx, `insert into subscriptions (subscriptionid, subscriptiondata) values($1, $2) 
								        on conflict (subscriptionid) do update set subscriptiondata=excluded.subscriptiondata`, id, data); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (db *Db) UpdateSubscription(ctx context.Context, subscription *models.Subscription) error {
	if data, err := json.Marshal(subscription); err == nil {
		id := subscription.Id
		if _, err := db.pool.Exec(ctx, `update subscriptions set subscriptiondata = $2 where subscriptionid = $1`, id, data); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (db *Db) DeleteSubscription(ctx context.Context, subscriptionId string) error {
	if _, err := db.pool.Exec(ctx, `delete from subscriptions subscriptionid = $1`, subscriptionId); err == nil {
		return nil
	} else {
		return err
	}
}

func (db *Db) CreateDatabase() error {
	_, err := db.pool.Exec(context.Background(), `CREATE TABLE Subscriptions (subscriptionId varchar(255) NOT NULL PRIMARY KEY, subscriptionData json NOT NULL)`)
	return err
}
