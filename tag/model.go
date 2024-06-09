package tag

type TagSet struct {
	Items []*Tag
}

type Tag struct {
	Id       int
	CreateAt int64
	*CreateTagRequest
}

type CreateTagRequest struct {
	//关联博客的ID
	BlogId int
	//标签的名称
	Key string
	//标签的值
	Value string
	//标签的颜色
	color string
}

type AddTagRequest struct {
	//一次可以添加多个 Tag
	Tags []*CreateTagRequest
}

type RemoveTagRequest struct {
	//一次可以移除多个 Tag
	TagIds []int
}
