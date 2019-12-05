# 随机生成密码

## 替代方案

echo "alias rpwd='openssl rand -base64 16'" >> ~/.zshrc

## 编译命令

### 正式版编译方式

GOOS=linux GOARCH=amd64 go build -a -ldflags "-X main.buildStamp=`date '+%Y-%m-%d_%I:%M:%S'`"

## 参数

### 数字字母下划线

默认

### 长度

-e 10
--length=10

### 大写字母

-u --Upper

### 小写字母

-l --Lower

### 数字

-n --Number

### 特殊字符

-c --Character
