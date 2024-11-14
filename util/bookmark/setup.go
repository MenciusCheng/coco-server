package bookmark

import (
	"coco-server/util/bookmark/api"
	"coco-server/util/bookmark/dao"
	"coco-server/util/bookmark/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.RouterGroup, db *gorm.DB) {
	// 初始化 DAO
	bookmarkDAO := dao.NewBookmarkDAO(db)
	folderDAO := dao.NewFolderDAO(db)

	// 初始化服务层
	bookmarkService := service.NewBookmarkService(bookmarkDAO)
	folderService := service.NewFolderService(folderDAO)

	// 初始化 API 层
	bookmarkAPI := api.NewBookmarkAPI(bookmarkService)
	folderAPI := api.NewFolderAPI(folderService)

	// 设置路由
	r.POST("/bookmark", bookmarkAPI.Create)
	r.GET("/bookmark/:id", bookmarkAPI.GetByID)
	r.PUT("/bookmark", bookmarkAPI.Update)
	r.DELETE("/bookmark/:id", bookmarkAPI.Delete)

	r.POST("/folder", folderAPI.Create)
	r.GET("/folder/:id", folderAPI.GetByID)
	r.PUT("/folder", folderAPI.Update)
	r.DELETE("/folder/:id", folderAPI.Delete)
	r.GET("/folder/tree", folderAPI.GetFolderTree)
}
