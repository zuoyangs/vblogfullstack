package blog

type Status int

const (
	//草稿
	STATUS_DRAFT Status = 0
	//已发布
	STATUS_PUBLISHED = 1
)

type UpdateMode string

const (
	//全量更新
	UPDATE_MODE_PUT UpdateMode = "put"
	//局部更新
	UPDATE_MODE_PATCH UpdateMode = "patch"
)
