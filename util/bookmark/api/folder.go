package api

import (
	"coco-server/util/bookmark/model"
	"coco-server/util/bookmark/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FolderAPI struct {
	Service *service.FolderService
}

func NewFolderAPI(service *service.FolderService) *FolderAPI {
	return &FolderAPI{Service: service}
}

func (api *FolderAPI) Create(c *gin.Context) {
	var folder model.Folder
	if err := c.ShouldBindJSON(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.Service.Create(&folder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, folder)
}

func (api *FolderAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	folder, err := api.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}
	c.JSON(http.StatusOK, folder)
}

func (api *FolderAPI) Update(c *gin.Context) {
	var folder model.Folder
	if err := c.ShouldBindJSON(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.Service.Update(&folder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, folder)
}

func (api *FolderAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := api.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Folder deleted"})
}

func (api *FolderAPI) GetFolderTree(c *gin.Context) {
	folderTree, err := api.Service.GetFolderTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, folderTree)
}
