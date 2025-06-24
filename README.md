# Todo List API
<<此專案僅提供練習，切勿使用在線上環境>>

一個基於 Go 的簡易 RESTful API，用於管理待辦事項。採用 DDD (Domain-Driven Design) 架構和 Standard Go Project Layout。

## 技術

- **語言**: Go 1.24.4
- **框架**: Gin
- **資料庫**: SQLite
- **ORM**: GORM
- **環境配置**: dotenv

## 功能需求

- ✅ 新增待辦事項
- ✅ 查詢待辦事項 (單個/全部)
- ✅ 更新待辦事項
- ✅ 刪除待辦事項
- ✅ 切換待辦事項完成狀態
- ✅ 使用者管理 (新增/查詢/更新/刪除)
- ✅ 查詢特定使用者的所有待辦事項

## 專案架構

本專案採用 **Standard Go Project Layout** 和 **DDD (Domain-Driven Design)** 架構：

```
todo-list/
├── cmd/                        # 入口點
│   └── main.go
├── internal/                   # 內部程式碼
│   ├── domain/                 # 領域層
│   │   ├── todo.go             # Todo 實例和介面定義
│   │   └── todo_test.go
│   ├── application/            # 應用層
│   │   └── todo_service.go
│   ├── infrastructure/         # 基礎設施層
│   │   ├── database.go
│   │   └── todo_repository.go
│   └── interfaces/             # 介面層
│       ├── handlers.go
│       └── routes.go
├── pkg/                        # 共用庫
├── configs/                    # 配置
├── docs/                       # 文件
│   └── api.md                  # API 文件
├── data/                       # DB
├── Dockerfile                  # Docker 配置
├── docker-compose.yml          # Docker Compose 配置
├── Makefile
└── README.md
```

## 快速開始

### 方式一：直接執行

1. **下載專案**:
```bash
git clone <repository-url>
cd todo-list
```

2. **安裝依賴**:
```bash
go mod tidy
```

3. **建立環境配置**:
```bash
cp configs/env.example .env
```

4. **執行程式**:
```bash
go run cmd/main.go
```

### 方式二：使用 Makefile

```bash
# 安裝依賴並建置
make all

# 開發模式執行
make dev

# 查看所有可用指令
make help
```

### 方式三：使用 Docker

```bash
# 使用 Docker Compose
docker-compose up -d

# 或使用 Docker 直接建置
docker build -t todo-list .
docker run -p 8080:8080 todo-list
```

## API 端點

| 方法 | 端點 | 描述 |
|------|------|------|
| GET | `/health` | 健康檢查 |
| GET | `/api/v1/users` | 取得所有使用者 |
| GET | `/api/v1/users/:id` | 取得特定使用者 |
| POST | `/api/v1/users` | 建立新使用者 |
| PUT | `/api/v1/users/:id` | 更新使用者 |
| DELETE | `/api/v1/users/:id` | 刪除使用者 |
| GET | `/api/v1/todos/:id` | 取得特定待辦事項 |
| POST | `/api/v1/todos` | 建立新待辦事項 |
| PUT | `/api/v1/todos/:id` | 更新待辦事項 |
| DELETE | `/api/v1/todos/:id` | 刪除待辦事項 |
| PATCH | `/api/v1/todos/:id/toggle` | 切換完成狀態 |
| GET | `/api/v1/user-todos/:user_id` | 取得指定使用者的所有待辦事項 |
| GET | `/api/v1/users/:id/todos` | 取得特定使用者及其所有待辦事項 |

詳細的 API 文件請參考 [docs/api.md](docs/api.md)

## 測試

```bash
# 執行所有測試
go test ./...

# 執行特定套件的測試
go test ./internal/domain

# 執行測試並顯示覆蓋率
go test -cover ./...
```

## 配置

### 環境變數配置

專案使用 `.env` 檔案來管理環境變數。複製 `env.example` 到 `.env` 並根據需要修改：

```bash
cp env.example .env
```

主要配置項：

```env
# 伺服器配置
PORT=8080
MODE=debug

# 資料庫配置
DB_PATH=./data/todo.db

# log 配置
LOG_LEVEL=info
```

### 載入環境變數

1. 複製 `env.example` 到 `.env` 檔案
2. 修改 `.env` 的設定
3. 如果都不存在，使用預設值

### 環境變數說明

| 變數名 | 預設值 | 描述 |
|--------|--------|------|
| PORT | 8080 | 伺服器連接埠 |
| MODE | debug | 模式 (debug/release) |
| DB_PATH | ./data/todo.db | DB 檔案路徑 |
| LOG_LEVEL | info | 日誌級別 |

## 部署

### 生產環境部署

1. **使用 Docker 部署**:
```bash
docker-compose -f docker-compose.prod.yml up -d
```
