package service

import (
	"coco-server/util/bookmark/dao"
	"coco-server/util/bookmark/model"
)

type FolderService struct {
	DAO *dao.FolderDAO
}

func NewFolderService(dao *dao.FolderDAO) *FolderService {
	return &FolderService{DAO: dao}
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

func (s *FolderService) GetFolderTree() ([]model.FolderTree, error) {
	// 获取所有根文件夹（假设根文件夹的 ParentID 为 0）
	rootFolders, err := s.DAO.GetFoldersByParentID(0)
	if err != nil {
		return nil, err
	}

	var folderTrees []model.FolderTree
	for _, folder := range rootFolders {
		folderTree, err := s.buildFolderTree(folder)
		if err != nil {
			return nil, err
		}
		folderTrees = append(folderTrees, folderTree)
	}

	return folderTrees, nil
}

func (s *FolderService) buildFolderTree(folder model.Folder) (model.FolderTree, error) {
	bookmarks, err := s.DAO.GetBookmarksByFolderID(folder.ID)
	if err != nil {
		return model.FolderTree{}, err
	}

	childFolders, err := s.DAO.GetFoldersByParentID(folder.ID)
	if err != nil {
		return model.FolderTree{}, err
	}

	var subFolders []model.FolderTree
	for _, childFolder := range childFolders {
		subFolderTree, err := s.buildFolderTree(childFolder)
		if err != nil {
			return model.FolderTree{}, err
		}
		subFolders = append(subFolders, subFolderTree)
	}

	return model.FolderTree{
		ID:         folder.ID,
		Name:       folder.Name,
		Bookmarks:  bookmarks,
		SubFolders: subFolders,
	}, nil
}
