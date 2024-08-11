package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	configController := NewDatabaseConfigController(service.DatabaseConfigService)

	routerGroup := GetRouterGroup()
	routerGroup.POST("/db/config", configController.CreateDatabaseConfig)
	routerGroup.GET("/db/configs", configController.GetAllDatabaseConfigs)
	routerGroup.GET("/db/config/:id", configController.GetDatabaseConfigByID)
	routerGroup.PUT("/db/config/:id", configController.UpdateDatabaseConfig)
	routerGroup.DELETE("/db/config/:id", configController.DeleteDatabaseConfig)
	routerGroup.POST("/db/config/test", configController.TestDatabaseConfig)

	routerGroup.GET("/db/config/:id/databases", configController.GetDatabases)
	routerGroup.GET("/db/config/:id/database/:db/tables", configController.GetTables)
	routerGroup.GET("/db/config/:id/database/:db/table/:table/info", configController.GetTableInfo)
	routerGroup.GET("/db/config/:id/database/:db/table/:table/data", configController.GetTableData)
	routerGroup.POST("/db/config/:id/database/:db/execute", configController.ExecuteSQL)
}

// DatabaseConfigController handles API requests for database configurations
type DatabaseConfigController struct {
	Service *service.IDatabaseConfigService
}

// NewDatabaseConfigController creates a new instance of DatabaseConfigController
func NewDatabaseConfigController(service *service.IDatabaseConfigService) *DatabaseConfigController {
	return &DatabaseConfigController{Service: service}
}

// CreateDatabaseConfig handles the creation of a new database configuration
func (controller *DatabaseConfigController) CreateDatabaseConfig(c *gin.Context) {
	var config model.DatabaseConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Service.CreateDatabaseConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create database configuration"})
		return
	}

	response.OkWithData(config, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": config})
}

// GetAllDatabaseConfigs retrieves all database configurations
func (controller *DatabaseConfigController) GetAllDatabaseConfigs(c *gin.Context) {
	configs, err := controller.Service.GetAllDatabaseConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve database configurations"})
		return
	}

	response.OkWithData(configs, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": configs})
}

// GetDatabaseConfigByID retrieves a specific database configuration by ID
func (controller *DatabaseConfigController) GetDatabaseConfigByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	config, err := controller.Service.GetDatabaseConfigByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve database configuration"})
		return
	}

	if config == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Database configuration not found"})
		return
	}

	response.OkWithData(config, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": config})
}

// UpdateDatabaseConfig updates a specific database configuration
func (controller *DatabaseConfigController) UpdateDatabaseConfig(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var config model.DatabaseConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.ID = uint(id)

	if err := controller.Service.UpdateDatabaseConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update database configuration"})
		return
	}

	response.OkWithData(config, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": config})
}

// DeleteDatabaseConfig deletes a specific database configuration by ID
func (controller *DatabaseConfigController) DeleteDatabaseConfig(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := controller.Service.DeleteDatabaseConfig(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete database configuration"})
		return
	}

	response.OkWithMessage("Database configuration deleted", c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Database configuration deleted"})
}

func (controller *DatabaseConfigController) TestDatabaseConfig(c *gin.Context) {
	var req model.TestConnectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Service.TestDatabaseConfig(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to test database configuration"})
		return
	}

	response.Ok(c)
}

// =========== table

// GetDatabases handles the retrieval of all databases for a specific config
func (controller *DatabaseConfigController) GetDatabases(c *gin.Context) {
	configID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config ID"})
		return
	}

	databases, err := controller.Service.GetDatabases(uint(configID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve databases"})
		return
	}

	response.OkWithData(databases, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": databases})
}

// GetTables handles the retrieval of all tables in a specific database
func (controller *DatabaseConfigController) GetTables(c *gin.Context) {
	configID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config ID"})
		return
	}

	dbName := c.Param("db")
	tables, err := controller.Service.GetTables(uint(configID), dbName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tables"})
		return
	}

	response.OkWithData(tables, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": tables})
}

// GetTableInfo handles the retrieval of the schema information of a specific table
func (controller *DatabaseConfigController) GetTableInfo(c *gin.Context) {
	configID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config ID"})
		return
	}

	dbName := c.Param("db")
	tableName := c.Param("table")
	columns, err := controller.Service.GetTableInfo(uint(configID), dbName, tableName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve table info"})
		return
	}

	response.OkWithData(columns, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": columns})
}

// GetTableData handles the retrieval of data from a specific table
func (controller *DatabaseConfigController) GetTableData(c *gin.Context) {
	configID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config ID"})
		return
	}

	dbName := c.Param("db")
	tableName := c.Param("table")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	data, err := controller.Service.GetTableData(uint(configID), dbName, tableName, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve table data"})
		return
	}

	response.OkWithData(data, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

// ExecuteSQL handles the execution of a custom SQL statement
func (controller *DatabaseConfigController) ExecuteSQL(c *gin.Context) {
	configID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config ID"})
		return
	}

	dbName := c.Param("db")
	var sqlRequest struct {
		SQL string `json:"sql"`
	}
	if err := c.ShouldBindJSON(&sqlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := controller.Service.ExecuteSQL(uint(configID), dbName, sqlRequest.SQL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute SQL"})
		return
	}

	response.OkWithData(data, c)
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}
