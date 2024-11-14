package model

type FolderTree struct {
	ID         uint         `json:"id"`
	Name       string       `json:"name"`
	Bookmarks  []Bookmark   `json:"bookmarks"`
	SubFolders []FolderTree `json:"sub_folders"`
}
