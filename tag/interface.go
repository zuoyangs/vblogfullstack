package tag

import "context"

type Service interface {
	//给文章添加 tag
	AddTag(context.Context, *AddTagRequest) (TagSet, error)
	//给文章移除 tag
	RemoveTag(context.Context, *RemoveTagRequest) (TagSet, error)
}
