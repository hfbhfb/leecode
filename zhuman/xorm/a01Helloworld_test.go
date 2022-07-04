package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	logxorm "xorm.io/xorm/log"
)

var engine *xorm.Engine

type User struct {
	Id        int64
	Name      string    `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
	Email     string    // 新增加一个字段,看sql有什么差别
	GroupId   int64     `xorm:"index"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

type TAutoCreated struct {
	Id        int64
	CreatedAt time.Time `xorm:"created"`
}

type Order struct {
	Id        int64
	Name      string
	Time      time.Time
	CreatedAt time.Time `xorm:"created"`
}

type Group struct {
	Id   int64
	Name string
}

type Question struct {
	Id      int64
	Content string
}

type Userinfo struct {
	Id       int64
	Name     string
	DetailId int64
}
type UserinfoTrans struct {
	Id         int64
	Username   string
	Departname string
	Alias      string
	Created    time.Time `xorm:"created"`
}

type Userdetail struct {
	Id     int64
	Gender int
}

type UserLock struct {
	Id        int64
	Name      string
	Version   int       `xorm:"version"`
	UpdatedAt time.Time `xorm:"updated"`
}

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	var err error
	UserTmp := "root"
	Password := "yourpass123"
	Host := "render.tpddns.cn"
	Port := "16015"
	Database := "mydb1"
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", UserTmp, Password, Host, Port, Database)
	// mysql -u root -pyourpass123 -P 16015 -hrender.tpddns.cn --database=mydb1

	engine, err = xorm.NewEngine("mysql", source)
	if err != nil {
		panic(err.Error())
	}

	// 打开sql调试,只有prepare的信息?用不了
	engine.ShowSQL(true)
	engine.Logger().SetLevel(logxorm.LOG_DEBUG)

	// err = engine.Sync(new(User), new(Group), new(Question))
	err = engine.Sync(new(User), new(Group), Question{}, TAutoCreated{}, Order{}, Userinfo{}, UserLock{}, UserinfoTrans{})
	if err != nil {
		panic(err.Error())
	}

	log.Println("完成数据库的初始化和同步")
}

func TestHelloWorldXorm(t *testing.T) {

}

// 获取user表所有的数据
func TestGetUserAll(t *testing.T) {
	users := make([]User, 0)
	log.Println("order start len : ", len(users))
	err := engine.Find(&users)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("")
		for k, v := range users {
			fmt.Println(k+1, v)
		}
		fmt.Println("")
		log.Println("order info : ", users)
	}
}

/*
022-07-03T13:36:50.104593Z       141 Prepare   SELECT `id`, `usr_name`, `email`, `group_id`, `updated_at`, `deleted_at` FROM `user` WHERE (`deleted_at`=? OR `deleted_at` IS NULL)
2022-07-03T13:36:50.258049Z       141 Execute   SELECT `id`, `usr_name`, `email`, `group_id`, `updated_at`, `deleted_at` FROM `user` WHERE (`deleted_at`='0001-01-01 00:00:00' OR `deleted_at` IS NULL)
*/
