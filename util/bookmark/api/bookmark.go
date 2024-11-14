package api

import (
	"coco-server/util/bookmark/model"
	"coco-server/util/bookmark/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookmarkAPI struct {
	Service *service.BookmarkService
}

func NewBookmarkAPI(service *service.BookmarkService) *BookmarkAPI {
	return &BookmarkAPI{Service: service}
}

func (api *BookmarkAPI) Create(c *gin.Context) {
	var bookmark model.Bookmark
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.Service.Create(&bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookmark)
}

func (api *BookmarkAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	bookmark, err := api.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}
	c.JSON(http.StatusOK, bookmark)
}

func (api *BookmarkAPI) Update(c *gin.Context) {
	var bookmark model.Bookmark
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.Service.Update(&bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookmark)
}

func (api *BookmarkAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := api.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark deleted"})
}
