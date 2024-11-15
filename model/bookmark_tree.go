package model

const (
	TypeBookmarkTreeFolder   = 0 // 文件夹
	TypeBookmarkTreeBookmark = 1 // 书签
)

type BookmarkTreeReq struct {
}

type BookmarkTreeRes struct {
	List []*BookmarkTreeData `json:"list"` // 列表
}

type BookmarkTreeData struct {
	BookmarkId int64               `json:"bookmarkId"` // 书签id
	FolderId   int64               `json:"folderId"`   // 文件夹id
	Type       int64               `json:"type"`       // 类型,0文件夹,1书签
	Name       string              `json:"name"`       // 名称
	Url        string              `json:"url"`        // 书签URL
	Icon       string              `json:"icon"`       // 书签图标
	Remark     string              `json:"remark"`     // 备注
	List       []*BookmarkTreeData `json:"list"`       // 子列表
}
