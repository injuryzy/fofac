# fofac

fofac 是对fofa api的一个封住，使用golang 编码,方便不同系统调用,通过查询的接口数据生成一个excel表格

## 用法

- -k, --key string fofa 用户的key (**必填**)
- -e, --email string fofa 用户的邮箱(**必填**)
- -q, --query stringArray 查询参数 （**排除after 和 before**）
- -a, --after string 在此之后，不包括,时间格式 yyyy-MM-dd
- -b, --before string 在此之前，不包括,时间格式 yyyy-MM-dd
- -s, --size int 每页条数 (default 10000)
- -f, --full 开启一年查询
- -h, --help help for search
- -p, --page int8 页码 (default 1)
- -t, --timeInterval int 查询时间间隔，单位为天，默认为1天 (default 1)

## 安装

```text
 1.可以下载release包
 2.自己下载源码打包
```

## 示例

```shell

search
-e
XXXXXXXX
-k
XXXXXXXX
# -q 多参数查询
-q  
app="ZTE-路由器"
-q
country="CN"
-s
100
-b
2024-05-08
```

## 参考

- 使用了cobar 来构建命令系统
- 使用 excelize 来生成excel表格
- https://github.com/inspiringz/fofa

**如果工具对你们有用，麻烦给我点一个start，谢谢**


