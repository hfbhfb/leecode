# 常见问题

* 如何使用Like？

答：
```Go
engine.Where("column like ?", "%"+char+"%").Find
```

* 怎么同时使用xorm的tag和json的tag？

答：使用空格

```Go
type User struct {
    Name string `json:"name" xorm:"name"`
}
```

* 我的struct里面包含bool类型，为什么它不能作为条件也没法用Update更新？

答：默认bool类型因为无法判断是否为空，所以不会自动作为条件也不会作为Update的内容。可以使用UseBool函数，也可以使用Cols函数

```Go
engine.Cols("bool_field").Update(&Struct{BoolField:true})
// UPDATE struct SET bool_field = true
```

* 我的struct里面包含float64和float32类型，为什么用他们作为查询条件总是不正确？

答：默认float32和float64映射到数据库中为float,real,double这几种类型，这几种数据库类型数据库的实现一般都是非精确的。因此作为相等条件查询有可能不会返回正确的结果。如果一定要作为查询条件，请将数据库中的类型定义为Numeric或者Decimal。

```Go
type Account struct {
    Money float64 `xorm:"Numeric"`
}
```

* 为什么Update时Sqlite3返回的affected和其它数据库不一样？

答：Sqlite3默认Update时返回的是update的查询条件的记录数条数，不管记录是否真的有更新。而Mysql和Postgres默认情况下都是只返回记录中有字段改变的记录数。

* xorm有几种命名映射规则？

答：目前支持SnakeMapper, SameMapper和GonicMapper三种。SnakeMapper支持结构体和成员以驼峰式命名而数据库表和字段以下划线连接命名；SameMapper支持结构体和数据库的命名保持一致的映射。GonicMapper在SnakeMapper的基础上对一些特定名词，比如ID的映射会映射为id，而不是像SnakeMapper那样为i_d。

* xorm支持复合主键吗？

答：支持。在定义时，如果有多个字段标记了pk，则这些字段自动成为复合主键，顺序为在struct中出现的顺序。在使用Id方法时，可以用`ID(xorm.PK{1, 2})`的方式来用。

* xorm如何使用Join？

答：一般我们配合Join()和extends标记来进行，比如我们要对两个表进行Join操作，我们可以这样：

```go
type Userinfo struct {
    Id int64
    Name string
    DetailId int64
}

type Userdetail struct {
    Id int64
    Gender int
}

type User struct {
    Userinfo `xorm:"extends"`
    Userdetail `xorm:"extends"`
}

var users = make([]User, 0)
err := engine.Table(&Userinfo{}).Join("LEFT", "userdetail", "userinfo.detail_id = userdetail.id").Find(&users)
```

请注意这里的Userinfo在User中的位置必须在Userdetail的前面，因为他在join语句中的顺序在userdetail前面。如果顺序不对，那么对于同名的列，有可能会赋值出错。

当然，如果Join语句比较复杂，我们也可以直接用Sql函数

```go
err := engine.Sql("select * from userinfo, userdetail where userinfo.detail_id = userdetail.id").Find(&users)
```

* 如果有自动增长的字段，Insert如何写？
答：Insert时，如果需要自增字段填充为自动增长的数值，请保持自增字段为0；如果自增字段为非0，自增字段将会被作为普通字段插入。

* 如果设置数据库时区？
答：

```Go
location, err = time.LoadLocation("Asia/Shanghai")
engine.TZLocation = location
```







1. 任意门海外业务搬迁项目容器支持，完成现网业务割接前容器CCE自检，降低现网割接风险
2. 竹芒客户现场支撑，将客户当前k8swan系统对接的SDK，由腾讯云切换到华为云，在新环境能够部署调通
3. 作为客户中青报容器测试支持和能链小象业务迁移支持接口人，积极主动解答客户的提问20+，回复内容详实
4. CICD文档可执行性易读性验证（gitlab，jerkins）执行测试用例完成测试任务
5. 参与了《贵阳一局点开局验证》《CCI自动化测试》《华南广州灰度友好 330版本验收》等测试任务完成学习任务
6. 通过岗位必要的考试，通过了 《SRE运维可信2.0》《远程交付上岗证考试》《SD流程考试》 《IRT流程考试》





| **工作职责** | 1。任意门海外业务搬迁项目容器支持，完成现网业务割接前容器CCE自检，降低现网割接风险竹芒客户现场支撑，将客户当前k8swan系统对接的SDK，由腾讯云切换到华为云，在新环境能够部署调通作为客户中青报容器测试支持和能链小象业务迁移支持接口人，积极主动解答客户的提问20+，回复内容详实二、测试支持，文档和证书：方案文档内部验证，CICD文档可执行性易读性验证（gitlab，jerkins）执行测试用例完成测试任务，参与了《贵阳一局点开局验证》《CCI自动化测试》《华南广州灰度友好 330版本验收》等测试任务完成学习任务，通过岗位必要的考试，通过了 《SRE运维可信2.0》《远程交付上岗证考试》《SD流程考试》 《IRT流程考试》 |
| ------------ | ------------------------------------------------------------ |
|              |                                                              |
|              |                                                              |













