# Login_JWT

使用 Gin 搭配 MySQL 的 JWT 身份驗證範例。

## 📖 簡介

這是一個以 Go、Gin 和 GORM 架構的身份驗證範例，實作了：

- 註冊（Register Users）
- 登入（Authenticate & issue JWT）
- 使用 JWT 中介軟體保護 API 路由

## ⚙️ 技術棧

- Go
- Gin
- GORM
- MySQL
- bcrypt（密碼雜湊）
- github.com/golang-jwt/jwt/v4
- godotenv（環境變數管理）

## 🛠 快速開始

### 1. ✨ 前置需求

- 安裝 [Go](https://golang.org/)
- 一個 MySQL 實例
- （可選）Docker 與 docker-compose

### 2. 建立專案

```bash
git clone https://github.com/JWCodeWrote/Login_JWT.git
cd GoGin-JWT-Auth

3. 配置環境變數

建立 .env，內容類似：

DB_DSN=root:password@tcp(localhost:3306)/dbname?parseTime=true
JWT_SECRET=your_secret_key
PORT=8080

4. 安裝相依套件

go mod tidy

5. 建立資料庫

使用 MySQL CLI 或管理工具：

CREATE DATABASE dbname;

6. 執行專案

go run main.go
# 或使用 Docker Compose（如有設定）
docker-compose up --build

應用程式將在 http://localhost:8080 運行。

🚦 API 使用方式

1. 註冊（Register）
	•	URL: POST /register
	•	請求 body:

{
  "email": "user@example.com",
  "password": "yourpassword"
}


	•	回應 body:

{ "id": 1 }



2. 登入（Login）
	•	URL: POST /login
	•	請求 body:

{
  "email": "user@example.com",
  "password": "yourpassword"
}


	•	回應 body:

{ "token": "<JWT token>" }



3. 受保護路由（Protected）
	•	URL: GET /api/profile
	•	Header: Authorization: Bearer <token>
	•	回應 body:

{ "userId": "1" }



路由由 JWT 中介保護，無 token 或 token 無效皆回傳 401。

🔐 安全性與功能建議
	•	使用 bcrypt 雜湊密碼，保護使用者資訊  ￼ ￼ ￼。
	•	可擴展內容：
	•	Refresh Token
	•	使用者角色與權限控管
	•	密碼重設、Email 驗證
	•	Token 黑名單／Redis 支援

📂 專案結構

.
├── main.go
├── .env
├── controllers/
│   └── auth.go
├── database/
│   └── db.go
├── middlewares/
│   └── jwt.go
└── models/
    └── user.go

🎯 範例

註冊後可使用下列 curl 作測試：

curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"foo@bar.com","password":"123456"}'

curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"foo@bar.com","password":"123456"}'

curl http://localhost:8080/api/profile \
  -H "Authorization: Bearer <your_token>"


⸻

📚 參考資源
	•	Gin + JWT 學習範例  ￼ ￼ ￼
	•	JWT 基本理論與實作教學 ()