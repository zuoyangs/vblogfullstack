package blog

type CreateBlogRequest struct {
	// 文章标题，用于展示在博客页面
	Title string `json:"title"`
	// 博客文章的作者名称
	Author string `json:"author"`
	// 创建博客的用户身份，通过Token识别
	Creator string `json:"creator"`
	// 博客文章的具体内容
	Content string `json:"content"`
	// 博客文章的简短概要，用于预览或列表展示
	Summary string `json:"summary"`
	// 博客文章的标签集合，用于分类和搜索，例如：语言:Golang, 分类:后端
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
}

type Blog struct {
	// 博客文章的唯一标识符
	ID int64 `json:"id"`
	// 博客文章的简短概要，从文章内容中提取并生成
	Summary string `json:"summary"`
	// 博客文章的创建时间，以Unix时间戳形式存储
	CreatedAt int64 `json:"created_at"`
	// 博客文章的最后更新时间，以Unix时间戳形式存储
	UpdatedAt int64 `json:"updated_at"`
	// 博客文章的发布时间，以Unix时间戳形式存储
	PublishedAt int64 `json:"published_at"`
	// 博客文章的审核时间，以Unix时间戳形式存储
	AuditedAt int64 `json:"audited_at"`
	// 博客文章的审核状态，例如：未审核、审核通过、审核拒绝等
	AuditStatus int `json:"audit_status"`
	// 博客文章的状态，例如：草稿、已发布、已下架等
	Status Status `json:"status"`
	// 创建博客时所使用的参数详情，存储为CreateBlogRequest类型的指针，用于内部逻辑处理，不直接暴露给JSON序列化
	CreateParams *CreateBlogRequest `json:"-"`
}

type BlogSet struct {
	// 存储博客列表的切片，每个元素都是一个Blog类型的指针
	Blogs []*Blog `json:"blogs"`
	// 博客列表的总数，用于分页等功能的计算
	TotalCount int64 `json:"total_count"`
}

// 查询博客
type QueryBlogRequest struct {
	PageSize   int      `json:"page_size"`
	PageNumber int      `json:"page_number"`
	Usernames  []string `json:"usernames"`
	Status     *Status  `json:"status"`
	Keywords   string   `json:"keywords"`
}

type UpdateBlogStatusRequest struct {
	ID     int64  `json:"id"`
	Status Status `json:"status"`
}

// 审核博客
type AuditBlogRequest struct {
	// 博客ID
	BlogId int `json:"blog_id"`
	// 审核状态
	BlogAuditStatus int `json:"blog_audit_status"`
}

// 获取博客详情页
type DescribeBlogRequest struct {
	BlogId int `json:"blog_id"`
}

type UpdateBlogRequest struct {
	ID                 string       `json:"id"`
	Scope              *common.Scop `json:"scope"`
	UpdateMode         UpdateMode   `json:"update_mode"`
	*CreateBlogRequest `json:"-"`
}
