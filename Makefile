.PHONY: help
help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## 初期セットアップ
	go mod download
	go install github.com/air-verse/air@v1.52.3
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golang/mock/mockgen@latest

.PHONY: up
up: ## Dockerコンテナを起動
	docker-compose up -d

.PHONY: down
down: ## Dockerコンテナを停止
	docker-compose down

.PHONY: logs
logs: ## ログを表示
	docker-compose logs -f

.PHONY: migrate-up
migrate-up: ## マイグレーションを実行
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/todo_app" up

.PHONY: migrate-down
migrate-down: ## マイグレーションをロールバック
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/todo_app" down

.PHONY: migrate-create
migrate-create: ## 新しいマイグレーションファイルを作成 (例: make migrate-create NAME=add_users_table)
	migrate create -ext sql -dir db/migrations -seq $(NAME)

.PHONY: test-up
test-up: ## テスト用DBコンテナを起動
	docker-compose -f docker-compose.test.yml up -d

.PHONY: test-down
test-down: ## テスト用DBコンテナを停止
	docker-compose -f docker-compose.test.yml down

.PHONY: test
test: ## ユニットテストを実行
	go test -v ./internal/...

.PHONY: test-feature
test-feature: ## Feature Testを実行
	go test -v ./test/feature/...

.PHONY: test-all
test-all: test-up ## 全てのテストを実行
	@echo "Waiting for test database..."
	@sleep 5
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3307)/todo_app_test" up
	go test -v ./test/unit/... ./test/feature/...
	$(MAKE) test-down

.PHONY: build
build: ## ビルド
	go build -o bin/api ./cmd/api

.PHONY: run
run: ## ローカルで実行（Air使用）
	air

.PHONY: clean
clean: ## クリーンアップ
	rm -rf bin/ tmp/
	docker-compose down -v

.PHONY: mock
mock: ## モックを生成
	@echo "Generating mocks..."
	@for file in internal/domain/repository/*_repository.go; do \
		filename=$$(basename $$file); \
		mockgen -source=$$file -destination=test/mock/mock_$$filename -package=mock; \
		echo "Generated test/mock/mock_$$filename"; \
	done

.PHONY: lint
lint: ## Lintを実行
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

