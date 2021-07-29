1. Реализовать приложение, которое реализует основной use-case вашей системы, т.е. поддерживает выполнение типовых запросов (из файла queries.sql урока 3, достаточно покрыть один-два запроса). Необходимо реализовать только Storage Layer вашего приложения, т.е. только часть взаимодействия с базой данных.
2. Реализовать интеграционное тестирование функциональности по выборке данных из базы.
3. Реализовать автоматизацию миграции структуры базы данных (файл schema.sql из предыдущих уроков). В файле README.md в корне проекта описать, как запускать миграцию структуры базы данных.

Установка миграции через го:
PostGSQL\DZ5\migration> go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Инициализация начальных схем миграции:
PostGSQL\DZ5\migration> migrate create -seq -ext sql -dir migrations init_schema

А что бы запустить миграцию:
migrate -path DZ5/migration -database "postgres://testuser:12345@localhost:5433/mydbfordz" -verbose up
