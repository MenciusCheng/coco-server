package service

import (
	"coco-server/util/bookmark/dao"
	"coco-server/util/bookmark/model"
)

type FolderService struct {
	DAO         *dao.FolderDAO
	BookmarkDAO *dao.BookmarkDAO
}

func NewFolderService(dao *dao.FolderDAO, BookmarkDAO *dao.BookmarkDAO) *FolderService {
	return &FolderService{DAO: dao, BookmarkDAO: BookmarkDAO}
}

func (s *FolderService) Create(folder *model.Folder) error {
	return s.DAO.Create(folder)
}

func (s *FolderService) GetByID(id uint) (*model.Folder, error) {
	return s.DAO.GetByID(id)
}

func (s *FolderService) Update(folder *model.Folder) error {
	return s.DAO.Update(folder)
}

func (s *FolderService) Delete(id uint) error {
	return s.DAO.Delete(id)
}

func (s *FolderService) GetFolderTree() ([]*model.FolderTree, error) {
	folders, err := s.DAO.GetAllFolders()
	if err != nil {
		return nil, err
	}

	bookmarks, err := s.BookmarkDAO.GetAllBookmarks()
	if err != nil {
		return nil, err
	}

	folderMap := make(map[uint]*model.FolderTree)
	for _, folder := range folders {
		folderMap[folder.ID] = &model.FolderTree{
			ID:         folder.ID,
			Name:       folder.Name,
			Bookmarks:  []*model.Bookmark{},
			SubFolders: []*model.FolderTree{},
		}
	}

	for _, bookmark := range bookmarks {
		if folder, exists := folderMap[bookmark.FolderID]; exists {
			folder.Bookmarks = append(folder.Bookmarks, bookmark)
		}
	}

	var rootFolders []*model.FolderTree
	for _, folder := range folders {
		if folder.ParentFolderID == 0 {
			rootFolders = append(rootFolders, folderMap[folder.ID])
		} else {
			if parentFolder, exists := folderMap[folder.ParentFolderID]; exists {
				parentFolder.SubFolders = append(parentFolder.SubFolders, folderMap[folder.ID])
			}
		}
	}

	return rootFolders, nil
}
