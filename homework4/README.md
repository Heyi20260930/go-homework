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
| 方法 | 路径 | 认证 | 说明 |
|------|------|------|------|
| POST | /api/register | 否 | 用户注册 |
| POST | /api/login | 否 | 用户登录，返回 JWT Token |
| GET | /api/posts | 否 | 获取所有文章列表 |
| GET | /api/posts/:id | 否 | 获取单篇文章详情 |
| POST | /api/posts | 是 | 创建文章（需登录） |
| PUT | /api/posts/:id | 是 | 更新文章（仅限作者） |
| DELETE | /api/posts/:id | 是 | 删除文章（仅限作者） |
| GET | /api/posts/:id/comments | 否 | 获取某篇文章的评论列表 |
| POST | /api/posts/:id/comments | 是 | 发表评论（需登录） |

**认证方式**：请求头中携带 `Authorization: Bearer <JWT Token>`。

## 关键设计
- 密码使用 bcrypt 加密存储
- JWT 载荷包含用户 ID、用户名、过期时间（24h）
- 文章/评论通过 Preload 关联用户信息
- 统一响应格式 {code, message, data}
- 中间件记录请求日志与耗时
