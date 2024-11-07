# go-test-with-db-ci

## フォルダ構成

```
project-root/
├── app/                              # アプリケーションロジック
│   ├── cmd/                          # サーバーエントリーポイント
│   │   └── server/
│   │       └── main.go               # サーバー起動コード
│   ├── api/                          # APIのエンドポイント (Handler層)
│   │   └── handler/
│   │       └── todo_handler.go       # TODO用のハンドラ
│   ├── internal/                     # 内部ロジック (ドメイン, ユースケース, リポジトリなど)
│   │   ├── domain/                   # ドメインモデル (例: Todo)
│   │   │   └── todo.go               # Todoモデルの定義
│   │   ├── usecase/                  # ユースケース
│   │   │   └── todo_usecase.go       # Todoのユースケース実装
│   │   ├── repository/               # リポジトリ
│   │   │   └── todo_repository.go    # Todoのリポジトリ実装
│   │   └── infrastructure/           # インフラ層
│   │       └── db/
│   │           ├── mysql.go          # DB接続設定
│   │           └── db_config_test.go # テスト用DB設定
│   └── go.mod                        # Goモジュール
│
└── iac/                              # Infrastructure as Code (IaC)
    ├── docker-compose.yml            # docker-compose定義
    ├── web/                          # アプリケーション用Docker設定
    │   └── Dockerfile                # Webサーバー用のDockerfile

```
