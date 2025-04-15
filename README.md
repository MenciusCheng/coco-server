# Coco Server

## 启动方式

### 本地启动

#### 编译

```
./dev.sh build
```

构建完成后，可执行文件会生成在 `./dist/server`。

#### 运行项目

```
./dist/server
```

或

```
./dist/server -c conf/config/dev/config.json
```

### Docker 启动

#### 构建 Docker 镜像

   运行以下命令构建 Docker 镜像：
```
docker-compose build
```

#### 启动服务

运行以下命令启动服务：

```
docker-compose up
```

## 访问服务

服务启动后，默认监听 9390 端口，可以通过以下地址访问：http://localhost:9390

## 备份数据库

进入备份目录，如：`../coco-db-backup`

### 导出数据库

```
mysqldump -uroot -p coco > coco_20250415.sql
```

### 导入数据库

```
mysql -uroot -p coco > coco_20250415.sql
```
