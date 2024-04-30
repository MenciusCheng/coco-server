help ()
{
    echo "使用说明"
    echo " 基本语法: dev.sh 命令模块 [option]"
    echo "命令模块："
    echo " help  显示当前帮助内容"
    echo " build 编译代码"
    echo "    例如：sh dev.sh build"
    echo " run   本地运行代码"
    echo "    例如：sh dev.sh run"
    exit 0
}

## 编译代码
build ()
{
  go build -o ./dist/ ./app/server
}

## 本地运行代码
run ()
{
  go run ./app/server
}

## 获取子shell命令
TARGET=$1
shift
case $TARGET in
	help) help ;;
	build) build $*;;
	run) run $*;;
	*) help ;;
esac