package main

import (
	"context"
	"fmt"
	//"sync"
	//"sync/atomic"
	//"time"
	"F:/PostGSQL/DZ5/pkg/score/storage/pg.go"

	//"os"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()

	url := "postgres://testuser:12345@localhost:5433/mydbfordz"

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	limit := 10
	hints, err := Search(ctx, dbpool, limit)
	if err != nil {
		log.Fatal(err)
	}

	for _, hint := range hints {
		fmt.Println(hint.user_name, hint.score, hint.money_sum, hint.people_all)
	}
}
