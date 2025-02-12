## 更新数据
    
更新数据使用`Update`方法，Update方法的第一个参数为需要更新的内容，可以为一个结构体指针或者一个Map[string]interface{}类型。当传入的为结构体指针时，只有非空和0的field才会被作为更新的字段。当传入的为Map类型时，key为数据库Column的名字，value为要更新的内容。

`Update`方法将返回两个参数，第一个为 更新的记录数，需要注意的是 `SQLITE` 数据库返回的是根据更新条件查询的记录数而不是真正受更新的记录数。

```Go
user := new(User)
user.Name = "myname"
affected, err := engine.ID(id).Update(user)
```

这里需要注意，Update会自动从user结构体中提取非0和非nil得值作为需要更新的内容，因此，如果需要更新一个值为0，则此种方法将无法实现，因此有两种选择：

* 1.通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。

```Go
affected, err := engine.ID(id).Cols("age").Update(&user)
```

* 2.通过传入map[string]interface{}来进行更新，但这时需要额外指定更新到哪个表，因为通过map是无法自动检测更新哪个表的。

```Go
affected, err := engine.Table(new(User)).ID(id).Update(map[string]interface{}{"age":0})
```

有时候希望能够指定必须更新某些字段，而其它字段根据值的情况自动判断，可以使用 `MustCols` 来组合 `Update` 使用。

```Go
affected, err := engine.ID(id).MustCols("age").Update(&user)
```

另外，如果需要更新所有的字段，可以使用 `AllCols()`。

```Go
affected, err := engine.ID(id).AllCols(").Update(&user)
```
