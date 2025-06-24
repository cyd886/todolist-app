# Todo List API 文件

## 概述

Todo List API 是一個基於 RESTful 設計的待辦事項管理 API，提供完整的 CRUD 操作，支援使用者管理和使用者特定的待辦事項。

## 基礎資訊

- **Base URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **編碼**: UTF-8

## 端點清單

### 使用者管理端點

#### 1. 取得所有使用者

**GET** `/users`

**回應範例**:
```json
[
  {
    "id": 1,
    "username": "jimmy",
    "email": "jimmy@example.com",
    "name": "Jimmy Chen",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
]
```

#### 2. 取得特定使用者

**GET** `/users/{id}`

**參數**:
- `id` (path): 使用者 ID

**回應範例**:
```json
{
  "id": 1,
  "username": "jimmy",
  "email": "jimmy@example.com",
  "name": "Jimmy Chen",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

#### 3. 建立新使用者

**POST** `/users`

**請求內容**:
```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "123456",
  "name": "New User"
}
```

**回應範例**:
```json
{
  "id": 2,
  "username": "newuser",
  "email": "newuser@example.com",
  "name": "New User",
  "created_at": "2024-01-01T11:00:00Z",
  "updated_at": "2024-01-01T11:00:00Z"
}
```

#### 4. 更新使用者

**PUT** `/users/{id}`

**參數**:
- `id` (path): 使用者 ID

**請求內容**:
```json
{
  "username": "updateduser",
  "email": "updated@example.com",
  "name": "Updated User"
}
```

#### 5. 刪除使用者

**DELETE** `/users/{id}`

**參數**:
- `id` (path): 使用者 ID

**回應範例**:
```json
{
  "message": "User deleted successfully"
}
```

### 待辦事項端點

#### 1. 取得特定待辦事項

**GET** `/todos/{id}`

**參數**:
- `id` (path): 待辦事項 ID

#### 2. 建立新待辦事項

**POST** `/todos`

**請求內容**:
```json
{
  "title": "新待辦事項",
  "description": "待辦事項描述",
  "user_id": 1
}
```

**回應範例**:
```json
{
  "id": 2,
  "title": "新待辦事項",
  "description": "待辦事項描述",
  "completed": false,
  "created_at": "2024-01-01T11:00:00Z",
  "updated_at": "2024-01-01T11:00:00Z",
  "user_id": 1
}
```

#### 3. 更新待辦事項

**PUT** `/todos/{id}`

**參數**:
- `id` (path): 待辦事項 ID

**請求內容**:
```json
{
  "title": "更新的標題",
  "description": "更新的描述",
  "completed": true
}
```

#### 4. 刪除待辦事項

**DELETE** `/todos/{id}`

**參數**:
- `id` (path): 待辦事項 ID

#### 5. 切換待辦事項狀態

**PATCH** `/todos/{id}/toggle`

**參數**:
- `id` (path): 待辦事項 ID

### 使用者特定待辦事項端點

#### 1. 取得指定使用者的所有待辦事項

**GET** `/user-todos/{user_id}`

**參數**:
- `user_id` (path): 使用者 ID

**回應範例**:
```json
[
  {
    "id": 1,
    "title": "使用者的待辦事項",
    "description": "描述",
    "completed": false,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z",
    "user_id": 1
  }
]
```

### 取得特定使用者及其所有待辦事項

**GET** `/users/{id}/todos`

**參數**:
- `id` (path): 使用者 ID

**回應範例**:
```json
{
  "id": 1,
  "username": "jimmy",
  "email": "jimmy@example.com",
  "name": "Jimmy Chen",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z",
  "todos": [
    {
      "id": 1,
      "title": "學習 Go 語言",
      "description": "學習 Go 語言基礎和 DDD 架構",
      "completed": false,
      "created_at": "2024-01-01T10:05:00Z",
      "updated_at": "2024-01-01T10:05:00Z",
      "user_id": 1
    }
  ]
}
```

### 健康檢查

**GET** `/health`

**回應範例**:
```json
{
  "status": "ok",
  "message": "Todo List API is running"
}
```

## 錯誤回應

### 400 Bad Request
```json
{
  "error": "Invalid ID"
}
```

### 404 Not Found
```json
{
  "error": "Todo not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Database connection failed"
}
```

## 使用範例

### 使用 curl 測試 API

1. **建立使用者**:
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username": "jimmy", "email": "jimmy@example.com", "password": "123456", "name": "Jimmy Chen"}'
```

2. **建立待辦事項**:
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "學習 Go", "description": "學習 Go 語言基礎", "user_id": 1}'
```

3. **取得使用者的所有待辦事項**:
```bash
curl -X GET http://localhost:8080/api/v1/user-todos/1
```

4. **更新待辦事項**:
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "學習 Go 進階", "description": "學習 Go 語言進階特性", "completed": true}'
```

5. **切換待辦事項狀態**:
```bash
curl -X PATCH http://localhost:8080/api/v1/todos/1/toggle
```

6. **刪除待辦事項**:
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```