# Blog Backend (Gin + GORM + SQLite + JWT)

## 运行环境
- Go 1.21+
- SQLite（内置，无需安装服务）

## 依赖安装
go mod tidy

## 启动
go run main.go
默认监听 :8080，数据库文件 blog.db 自动创建。

## 接口
见上文接口清单表格。
认证方式：Authorization: Bearer <JWT>

## 关键设计
- 密码使用 bcrypt 加密存储
- JWT 载荷包含用户 ID、用户名、过期时间（24h）
- 文章/评论通过 Preload 关联用户信息
- 统一响应格式 {code, message, data}
- 中间件记录请求日志与耗时