package service

import (
	"coco-server/middleware/db"
	"coco-server/model"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// IDatabaseConfigService handles the operations for database configurations
type IDatabaseConfigService struct {
}

var DatabaseConfigService = NewDatabaseConfigService()

// NewDatabaseConfigService creates a new instance of IDatabaseConfigService
func NewDatabaseConfigService() *IDatabaseConfigService {
	return &IDatabaseConfigService{}
}

// =========== CRUD

// CreateDatabaseConfig adds a new database configuration
func (service *IDatabaseConfigService) CreateDatabaseConfig(config *model.DatabaseConfig) error {
	return db.MySQLCon.Create(config).Error
}

// GetAllDatabaseConfigs retrieves all database configurations
func (service *IDatabaseConfigService) GetAllDatabaseConfigs() ([]model.DatabaseConfig, error) {
	var configs []model.DatabaseConfig
	err := db.MySQLCon.Find(&configs).Error
	return configs, err
}

// GetDatabaseConfigByID retrieves a specific database configuration by ID
func (service *IDatabaseConfigService) GetDatabaseConfigByID(id uint) (*model.DatabaseConfig, error) {
	var config model.DatabaseConfig
	err := db.MySQLCon.First(&config, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &config, err
}

// UpdateDatabaseConfig updates a specific database configuration
func (service *IDatabaseConfigService) UpdateDatabaseConfig(config *model.DatabaseConfig) error {
	return db.MySQLCon.Save(config).Error
}

// DeleteDatabaseConfig deletes a database configuration by ID
func (service *IDatabaseConfigService) DeleteDatabaseConfig(id uint) error {
	return db.MySQLCon.Delete(&model.DatabaseConfig{}, id).Error
}

func (service *IDatabaseConfigService) TestDatabaseConfig(req *model.TestConnectionRequest) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		req.Username, req.Password, req.Host, req.Port, req.Database)

	// 测试数据库连接
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := mdb.DB()
	if err != nil {
		return err
	}

	// Ping数据库以验证连接
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}

// =========== table

// GetDatabases retrieves all databases in a given MySQL server configuration
func (service *IDatabaseConfigService) GetDatabases(configID uint) ([]string, error) {
	var config model.DatabaseConfig
	if err := db.MySQLCon.First(&config, configID).Error; err != nil {
		return nil, err
	}

	// 建立一个新的数据库连接以获取数据库列表
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询所有数据库
	var databases []string
	rows, err := mdb.Raw("SHOW DATABASES").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var database string
		if err := rows.Scan(&database); err != nil {
			return nil, err
		}
		databases = append(databases, database)
	}

	return databases, nil
}

// GetTables retrieves all tables in a specific database under a given MySQL server configuration
func (service *IDatabaseConfigService) GetTables(configID uint, dbName string) ([]string, error) {
	var config model.DatabaseConfig
	if err := db.MySQLCon.First(&config, configID).Error; err != nil {
		return nil, err
	}

	// 建立连接到指定数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, dbName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询所有表
	var tables []string
	rows, err := mdb.Raw("SHOW TABLES").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

// GetTableInfo retrieves the schema information of a specific table
func (service *IDatabaseConfigService) GetTableInfo(configID uint, dbName, tableName string) ([]map[string]interface{}, error) {
	var config model.DatabaseConfig
	if err := db.MySQLCon.First(&config, configID).Error; err != nil {
		return nil, err
	}

	// 建立连接到指定数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, dbName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询表信息
	rows, err := mdb.Raw(fmt.Sprintf("DESCRIBE %s", tableName)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []map[string]interface{}
	for rows.Next() {
		var field, typeStr, null, key, extra string
		var defaultVal sql.NullString
		if err := rows.Scan(&field, &typeStr, &null, &key, &defaultVal, &extra); err != nil {
			return nil, err
		}

		// 如果 defaultVal.Valid 为 true，使用默认值，否则设为nil
		column := map[string]interface{}{
			"Field":   field,
			"Type":    typeStr,
			"Null":    null,
			"Key":     key,
			"Default": nil,
			"Extra":   extra,
		}
		if defaultVal.Valid {
			column["Default"] = defaultVal.String
		}

		columns = append(columns, column)
	}

	return columns, nil
}

// GetTableData retrieves the data of a specific table
func (service *IDatabaseConfigService) GetTableData(configID uint, dbName, tableName string, limit, offset int) ([]map[string]interface{}, error) {
	var config model.DatabaseConfig
	if err := db.MySQLCon.First(&config, configID).Error; err != nil {
		return nil, err
	}

	// 建立连接到指定数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, dbName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询表数据
	rows, err := mdb.Table(tableName).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列信息
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 获取行数据
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make(map[string]interface{})
		cols := make([]interface{}, len(columns))
		colsPtr := make([]interface{}, len(columns))
		for i := range cols {
			colsPtr[i] = &cols[i]
		}

		if err := rows.Scan(colsPtr...); err != nil {
			return nil, err
		}

		for i, col := range columns {
			row[col] = cols[i]
		}
		results = append(results, row)
	}

	// UTF-8 解码
	decodedRows := make([]map[string]interface{}, len(results))
	for i, row := range results {
		decodedRow := make(map[string]interface{})
		for key, value := range row {
			strValue, ok := value.([]uint8)
			if ok {
				decodedRow[key] = decodeUTF8String(strValue)
			} else {
				decodedRow[key] = value
			}
		}
		decodedRows[i] = decodedRow
	}

	return decodedRows, nil
}

// ExecuteSQL executes a custom SQL statement on a specific database
func (service *IDatabaseConfigService) ExecuteSQL(configID uint, dbName, sqlStatement string) ([]map[string]interface{}, error) {
	var config model.DatabaseConfig
	if err := db.MySQLCon.First(&config, configID).Error; err != nil {
		return nil, err
	}

	// 建立连接到指定数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, dbName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 执行自定义SQL
	rows, err := mdb.Raw(sqlStatement).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列信息
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 获取行数据
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make(map[string]interface{})
		cols := make([]interface{}, len(columns))
		colsPtr := make([]interface{}, len(columns))
		for i := range cols {
			colsPtr[i] = &cols[i]
		}

		if err := rows.Scan(colsPtr...); err != nil {
			return nil, err
		}

		for i, col := range columns {
			row[col] = cols[i]
		}
		results = append(results, row)
	}

	// UTF-8 解码
	decodedRows := make([]map[string]interface{}, len(results))
	for i, row := range results {
		decodedRow := make(map[string]interface{})
		for key, value := range row {
			strValue, ok := value.([]uint8)
			if ok {
				decodedRow[key] = decodeUTF8String(strValue)
			} else {
				decodedRow[key] = value
			}
		}
		decodedRows[i] = decodedRow
	}

	return decodedRows, nil
}

func decodeUTF8String(input []uint8) string {
	// 解码为 UTF-8 字符串
	return string(input)
}
