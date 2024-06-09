package blog

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"
)

// Service 接口的定义
type Service interface {

	// 创建博客文章
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	// 查询文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

	// 更新文章
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	// 删除文章
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

	// 文章详情页
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

	// 修改文章状态
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)

	//文章审核
	AuditBlog(context.Context, *AuditBlogRequest) (*Blog, error)
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: make(map[string]string, 0),
	}
}

func NewBlog(blogCreationReq *CreateBlogRequest) *Blog {
	return &Blog{
		CreatedAt:    time.Now().Unix(),
		Status:       STATUS_DRAFT,
		CreateParams: blogCreationReq,
	}
}

func (b *Blog) TableName() string {
	return "blogs"
}

func (b *Blog) string() string {
	jsonData, err := json.Marshal(b)
	if err != nil {
		log.Printf("Error marshaling Blog failed: %v\n", err)
		return "Error marshaling Blog"
	}
	return string(jsonData)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Blogs: make([]*Blog, 0),
	}
}

func (b *BlogSet) string() string {
	jsonData, err := json.Marshal(b)
	if err != nil {
		log.Printf("Error marshaling BlogSet failed: %v\n", err)
		return "Error marshaling BlogSet"
	}
	return string(jsonData)
}

func (b *BlogSet) Add(items ...*Blog) {
	b.Blogs = append(b.Blogs, items...)
}

// 查询文章列表
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   10,
		PageNumber: 1,
		Usernames:  make([]string, 0), //用户列表
	}
}

// 分页
func (q *QueryBlogRequest) Offset() int {
	return int((q.PageNumber - 1) * q.PageSize)
}

// 添加用户
func (q *QueryBlogRequest) Addusername(username ...string) {
	q.Usernames = append(q.Usernames, username...)
}

// 解析页面大小
func (q *QueryBlogRequest) ParsePageSize(pagesize string) {
	parsedPageSize, err := strconv.ParseInt(pagesize, 10, 64)
	if err != nil {
		log.Printf("解析 pagesize 时候出错: %v", err)
		return
	}

	if parsedPageSize <= 0 {
		log.Printf("pagesize 大小必须大于0,但得到的是: %d", parsedPageSize)
		return
	}
	q.PageSize = int(parsedPageSize)

}

// 解析页码
func (q *QueryBlogRequest) ParsePageNumber(pagenumber string) {
	parsedPageNumber, err := strconv.ParseInt(pagenumber, 10, 64)
	if err != nil {
		log.Printf("解析 pagenumber 时候出错: %v", err)
		return
	}

	if parsedPageNumber <= 0 {
		log.Printf("pagenumber 大小必须大于0,但得到的是: %d", parsedPageNumber)
		return
	}

	q.PageNumber = int(parsedPageNumber)
}

// 设置状态
func (q *QueryBlogRequest) SetStatus(s Status) {
	q.Status = &s
}

func NewPutUpdateBlogStatusRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		ID: id,
	}
}

func NewAuditBlogRequest(id int) *AuditBlogRequest {
	return &AuditBlogRequest{
		BlogId: id,
	}
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}
