package service

import (
	"coco-server/util/bookmark/dao"
	"coco-server/util/bookmark/model"
)

type BookmarkService struct {
	DAO *dao.BookmarkDAO
}

func NewBookmarkService(dao *dao.BookmarkDAO) *BookmarkService {
	return &BookmarkService{DAO: dao}
}

func (s *BookmarkService) Create(bookmark *model.Bookmark) error {
	return s.DAO.Create(bookmark)
}

func (s *BookmarkService) GetByID(id uint) (*model.Bookmark, error) {
	return s.DAO.GetByID(id)
}

func (s *BookmarkService) Update(bookmark *model.Bookmark) error {
	return s.DAO.Update(bookmark)
}

func (s *BookmarkService) Delete(id uint) error {
	return s.DAO.Delete(id)
}
