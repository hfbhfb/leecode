## 事务处理

当使用事务处理时，需要创建 Session 对象。在进行事务处理时，可以混用 ORM 方法和 RAW 方法，如下代码所示：

```Go
func MyTransactionOps() error {
    session := engine.NewSession()
    defer session.Close()

    // add Begin() before any action
    if err := session.Begin(); err != nil {
        return err
    }

    user1 := Userinfo{Username: "xiaoxiao", Departname: "dev", Alias: "lunny", Created: time.Now()}
    if _, err := session.Insert(&user1); err != nil {
        return err
    }
    user2 := Userinfo{Username: "yyy"}
    if _, err = session.Where("id = ?", 2).Update(&user2); err != nil {
        return err
    }

    if _, err = session.Exec("delete from userinfo where username = ?", user2.Username); err != nil {
        return err
    }

    // add Commit() after all actions
    return session.Commit()
}
```

* 注意如果您使用的是 mysql，数据库引擎为 innodb 事务才有效，myisam 引擎是不支持事务的。
