package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	dbpool *pgxpool.Pool
}

func NewPG(dbpool *pgxpool.Pool) *PG {
	return &PG{dbpool}
}

// В рамках каждого слоя желательно работать только со структурами,
// принадлежащими этому слою.
// Это многословно, но код получается явным и легко изменяемым.

type UserRait struct {
	user_name  string
	score      int
	money_sum  int
	people_all int
}

func (s *PG) Search(ctx context.Context, limit int) ([]UserRait, error) {
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

	rows, err := s.dbpool.Query(ctx, sql, limit)
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
