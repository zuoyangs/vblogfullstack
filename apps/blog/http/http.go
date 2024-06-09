package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zuoyangs/vblog_backend/blog"
)

// 构造函数，接收一个 Service 的实例，并返回一个 Handler 的实例
func NewHandler(svr blog.Service) *Handler {
	return &Handler{
		service: svr,
	}
}

// Handler 结构体，包含一个 Service 字段
type Handler struct {
	service blog.Service
}

// 将业务逻辑和路由相关联注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/vblog/api/v1/blogs", h.CreateBlog)
}

// 处理函数，接收一个 gin.Context 参数。调用 Service 的业务处理方法进行业务逻辑处理
func (h *Handler) CreateBlog(ctx *gin.Context) {
	h.service.CreateBlog(nil, nil)
}
