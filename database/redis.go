package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
)

const KEY_NAME = "f2pool_queries"

type DB struct {
	client *redis.Client
}

func NewDB(db int) *DB {
	// TODO: secure redis
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       db,
	})

	return &DB{
		client,
	}
}

func (db *DB) SaveQuery(query *models.ModelQuery) error {
	jsonStr, err := query.MarshalBinary()
	if err != nil {
		return err
	}

	_, err = db.client.LPush(KEY_NAME, jsonStr).Result()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) FetchQueries(params *history.QueriesHistoryParams) ([]*models.ModelQuery, error) {
	val, err := db.client.LLen(KEY_NAME).Result()
	if err != nil {
		return nil, err
	}

	fmt.Printf("len %d\n", val)
	results, err := db.client.LRange(KEY_NAME, val-20, val).Result()
	if err != nil {
		return nil, err
	}

	var queries []*models.ModelQuery
	for _, result := range results {
		var query = &models.ModelQuery{}
		err := query.UnmarshalBinary([]byte(result))
		if err != nil {
			return nil, err
		}
		queries = append(queries, query)
	}

	return queries, nil
}
