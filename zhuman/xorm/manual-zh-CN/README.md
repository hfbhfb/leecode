# xorm

xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。xorm的目标并不是让你完全不去学习SQL，我们认为SQL并不会为ORM所替代，但是ORM将可以解决绝大部分的简单SQL需求。xorm支持两种风格的混用。

## 特性

* 支持 Struct 和数据库表之间的灵活映射，并支持自动同步
* 事务支持
* 同时支持原始SQL语句和 ORM 操作的混合执行
* 使用连写来简化调用
* 支持使用ID, In, Where, Limit, Join, Having, Table, SQL, Cols等函数和结构体等方式作为条件
* 支持级联加载 Struct
* Schema支持（仅Postgres）
* 支持缓存
* 通过 [xorm.io/reverse](https://xorm.io/reverse) 支持根据数据库自动生成 xorm 结构体
* 支持记录版本（即乐观锁）
* 通过 [xorm.io/builder](https://xorm.io/builder) 内置 SQL Builder 支持
* 上下文缓存支持
* 支持日志上下文

## 驱动支持

xorm 当前支持的驱动和数据库如下：

* [Mysql5.*](https://github.com/mysql/mysql-server/tree/5.7) / [Mysql8.*](https://github.com/mysql/mysql-server) / [Mariadb](https://github.com/MariaDB/server) / [Tidb](https://github.com/pingcap/tidb)
  - [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
  - [github.com/ziutek/mymysql/godrv](https://github.com/ziutek/mymysql/godrv)

* [Postgres](https://github.com/postgres/postgres) / [Cockroach](https://github.com/cockroachdb/cockroach)
  - [github.com/lib/pq](https://github.com/lib/pq)

* [SQLite](https://sqlite.org)
  - [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

* MsSql
  - [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)

* Oracle
  - [github.com/mattn/go-oci8](https://github.com/mattn/go-oci8) (试验性支持)

## 安装

```
go get xorm.io/xorm
```

## 文档

* [操作指南](https://xorm.io/docs)

* [Godoc代码文档](http://pkg.go.dev/xorm.io/xorm)

## 讨论

请加入QQ群：280360085 进行讨论。

## 贡献

如果您也想为Xorm贡献您的力量，请查看 [CONTRIBUTING](https://gitea.com/xorm/xorm/blob/master/CONTRIBUTING.md)

## LICENSE

BSD License
[http://creativecommons.org/licenses/BSD/](http://creativecommons.org/licenses/BSD/)