package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	//"os"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

type UserRait struct {
	user_name  string
	score      int
	money_sum  int
	people_all int
}

// search ищет всех сотрудников, email которых начинается с prefix.
// Из функции возвращается список EmailSearchHint, отсортированный по Email.
// Размер возвращаемого списка ограничен значением limit.
func search(ctx context.Context, dbpool *pgxpool.Pool, limit int) ([]UserRait, error) {
	const sql = `
select
    user_name,
    score,
	money_sum,
	people_all
from users
ORDER BY score DESC
limit $1;
`

	rows, err := dbpool.Query(ctx, sql, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()

	// В слайс hints будут собраны все строки, полученные из базы
	var hints []UserRait

	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint UserRait

		// Scan записывает значения столбцов в свойства структуры hint
		err = rows.Scan(&hint.user_name, &hint.score, &hint.money_sum, &hint.people_all)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		hints = append(hints, hint)
	}

	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}

	return hints, nil
}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64

	attacker := func(stopAt time.Time) {
		for {
			_, err := search(ctx, dbpool, 10)
			if err != nil {
				log.Fatal(err)
			}

			atomic.AddUint64(&queries, 1)

			if time.Now().After(stopAt) {
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}

	wg.Wait()

	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func main() {
	ctx := context.Background()

	url := "postgres://testuser:12345@localhost:5433/mydbfordz"

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	cfg.MaxConns = 8
	cfg.MinConns = 8

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	duration := time.Duration(10 * time.Second)
	threads := 1000
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)

	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)
}

// func main() {
// 	ctx := context.Background()

// 	url := "postgres://testuser:12345@localhost:5433/mydbfordz"

// 	cfg, err := pgxpool.ParseConfig(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer dbpool.Close()

// 	limit := 10
// 	hints, err := search(ctx, dbpool, limit)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, hint := range hints {
// 		fmt.Println(hint.user_name, hint.score, hint.money_sum, hint.people_all)
// 	}
// }

// func main() {
// 	ctx := context.Background()

// 	// В этом уроке используется сервер PostgreSQL, поднятый с использованием Docker.
// 	// Так как контейнер с PostgreSQL запускается с параметром -p 5432:5432,
// 	// то соединение с сервером базы данных должно успешно устанавливаться по пути localhost:5432.
// 	// Пользователь myuser и база данных mydb были созданы в прошлых уроках.
// 	url := "postgres://testuser:12345@localhost:5433/mydbfordz"

// 	conn, err := pgx.Connect(ctx, url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(ctx)

// 	// В метод Scan передаётся ссылка на переменную greeting
// 	// туда будет записан результат работы запроса.
// 	// Если выборка идет из нескольких столбцов,
// 	// то для каждого столбца в Scan передаётся по одной ссылке
// 	var greeting string
// 	err = conn.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(greeting)
// }
