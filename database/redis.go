package database

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

const KEY_NAME = "f2pool_queries"

type DB struct {
	client *redis.Client
}

func NewDB() *DB {
	// TODO: pass configuration from env vars
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"), //"localhost:6379",
		Password: "",                     // no password set
		DB:       0,                      // use default DB
	})

	return &DB{
		client,
	}
}

func (db *DB) SaveQuery(params *tools.LookupDomainParams) (*models.ModelQuery, error) {
	ctx := context.Background()
	addrs, err := net.DefaultResolver.LookupIP(ctx, "ip4", params.Domain) //LookupIP()
	if err != nil {
		return nil, err
	}

	query := &models.ModelQuery{
		Addresses: models.NewModelAddresses(addrs),
		// TODO: consider req.Header.Get("X-Forwarded-For")
		ClientIP:  params.HTTPRequest.RemoteAddr,
		CreatedAt: time.Now().Unix(),
	}

	jsonStr, err := query.MarshalBinary()
	if err != nil {
		return nil, err
	}

	_, err = db.client.LPush(KEY_NAME, jsonStr).Result()
	if err != nil {
		return nil, err
	}

	return query, nil
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
