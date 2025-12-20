# TODO API - ドメイン駆動設計

Golang + Echo + MySQL で構築したTODO管理APIです。ドメイン駆動設計（DDD）のクリーンアーキテクチャで実装しています。

## プロジェクトの特徴

- **ドメイン駆動設計（DDD）**: 4層のレイヤードアーキテクチャ
- **依存性逆転の原則**: Domain層にRepositoryインターフェースを配置
- **DIコンテナ**: uber/digを使用した依存性注入
- **バリデーション**: go-playground/validatorによる宣言的バリデーション
- **ホットリロード**: Airによる開発体験の向上
- **マイグレーション管理**: golang-migrateによるDB管理
- **テスト**: Unit Test（gomock）とFeature Test（統合テスト）

## ディレクトリ構成

```
.
├── cmd/
│   └── api/
│       └── main.go                    # エントリーポイント
├── internal/
│   ├── presentation/                  # プレゼンテーション層
│   │   ├── controller/todo/          # エンドポイントごとのコントローラー
│   │   ├── request/                  # リクエストDTO
│   │   │   ├── todo/                 # TODOリクエスト
│   │   │   └── validator.go          # バリデーター
│   │   └── response/todo/            # レスポンスDTO
│   ├── application/                   # アプリケーション層
│   │   └── usecase/todo/             # ユースケース
│   ├── domain/                        # ドメイン層
│   │   ├── entity/todo/              # エンティティ
│   │   │   ├── todo.go
│   │   │   └── valueobject/          # 値オブジェクト
│   │   └── repository/               # リポジトリインターフェース
│   ├── infrastructure/                # インフラストラクチャ層
│   │   ├── persistence/todo/         # リポジトリ実装
│   │   ├── database/                 # DB接続
│   │   └── di/                       # DIコンテナ
│   └── router/                        # ルーティング
├── test/
│   ├── feature/                       # 統合テスト
│   │   ├── helper/                   # テスト共通処理
│   │   └── todo/                     # TODO機能のテスト
│   ├── unit/usecase/todo/            # ユニットテスト
│   └── mock/                          # モック
├── db/migrations/                     # マイグレーションファイル
├── docs/
│   └── swagger.yaml                   # OpenAPI定義
├── docker/                            # Dockerファイル
├── docker-compose.yml                 # 開発環境
├── docker-compose.test.yml            # テスト環境
├── Makefile                           # タスク管理
└── README.md
```

## 技術スタック

| カテゴリ | 技術 |
|---------|------|
| 言語 | Go 1.22+ |
| フレームワーク | Echo v4 |
| ORM | GORM |
| DB | MySQL 8.0 |
| DIコンテナ | uber/dig |
| バリデーション | go-playground/validator |
| マイグレーション | golang-migrate |
| テスト | gomock, httptest |
| ホットリロード | Air |
| コンテナ | Docker, Docker Compose |

## セットアップ

### 前提条件

- Go 1.22以上
- Docker & Docker Compose
- Make

### 1. リポジトリのクローン

```bash
git clone https://github.com/jobpay/todo.git
cd todo
```

### 2. 必要なツールのインストール

```bash
make setup
```

### 3. Dockerコンテナの起動

```bash
make up
```

### 4. マイグレーション実行

```bash
# コンテナが起動するまで少し待ってから実行
make migrate-up
```

### 5. 動作確認

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

## 使い方

### API エンドポイント

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/api/todos` | TODO一覧取得 |
| GET | `/api/todos/:id` | TODO詳細取得 |
| POST | `/api/todos` | TODO作成 |
| PUT | `/api/todos/:id` | TODO更新 |
| DELETE | `/api/todos/:id` | TODO削除 |

### 使用例

#### TODO作成

```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "プロジェクトの実装",
    "description": "ドメイン駆動設計でTODOアプリを実装する",
    "due_date": "2025-12-31T23:59:59Z"
  }'
```

#### TODO一覧取得

```bash
curl http://localhost:8080/api/todos
```

#### TODO詳細取得

```bash
curl http://localhost:8080/api/todos/1
```

#### TODO更新

```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "プロジェクトの実装完了",
    "description": "実装が完了しました",
    "completed": true,
    "due_date": "2025-12-31T23:59:59Z"
  }'
```

#### TODO削除

```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

## テスト

### ユニットテスト実行

```bash
make test
```

### Feature Test（統合テスト）実行

```bash
# テスト用DBを起動してテスト実行
make test-all
```

### モック生成

```bash
make mock
```


## アーキテクチャ

### レイヤー構成

```
┌─────────────────────────────────────────┐
│     Presentation Layer (Controller)     │  HTTPリクエスト/レスポンス
├─────────────────────────────────────────┤
│     Application Layer (UseCase)         │  ビジネスロジック
├─────────────────────────────────────────┤
│     Domain Layer (Entity/Repository)    │  ドメインモデル
├─────────────────────────────────────────┤
│  Infrastructure Layer (Persistence/DB)  │  DB接続、外部API
└─────────────────────────────────────────┘
```

### 依存関係の方向

```
Presentation → Application → Domain ← Infrastructure
```

- **依存性逆転の原則**: Infrastructure層がDomain層のインターフェース（Repository）を実装
- **単一責任の原則**: 各コントローラーは1つのエンドポイントのみ担当

## 開発用コマンド

```bash
# ヘルプ表示
make help

# コンテナ起動
make up

# コンテナ停止
make down

# ログ表示
make logs

# ローカルで実行（Air使用）
make run

# ビルド
make build

# マイグレーション実行
make migrate-up

# マイグレーションロールバック
make migrate-down

# 新しいマイグレーション作成
make migrate-create NAME=add_users_table

# テスト
make test              # ユニットテスト
make test-feature      # 統合テスト
make test-all          # 全テスト

# モック生成
make mock

# クリーンアップ
make clean
```