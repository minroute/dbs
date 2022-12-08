package example

import "github.com/xiiapp/dbs"

type UserModel struct {
	dbs.HookFunc
	Id int `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
}

func (m *UserModel) UniqueIndex() {

}

func (m *UserModel) Index() {

}

func (m *UserModel) TableName() string {
	return "user"
}

// User 实例化User类
// @todo：参数应该是配置读取，而非传进来，需要修改成，如果有传就用传的，如果没传就用连接池取一个出来。
func User(d *dbs.DB) *dbs.DB {
	return d
}

func (m *UserModel) beforeDel() {
	panic("implement me")
}

func main() {

	// Where("id = ?", 1)
	// Where("id=1")
	// Where("id", &[]int{1, 2, 3})
	// Where(&map["id": 1,"age":&[]int{1, 2, 3}])
	// Where([]slice{})
	// 	User(nil).Select("id ,name as realname").
	// 	Where("id = ?", 1).
	// 	Where("id = ?", 1).
	// 	AndFilterWhere().
	// 	OrFilterWhere().
	// 	GroupBy("id").
	// 	Having("id > ?", 1).
	// 	OrderBy("id desc").
	// 	Limit(1).
	// 	Offset(1).
	// 	Find(&UserModel{})
	// User(nil).FindOne()
	// User(nil).FindAll()
	//
	// User(nil).Select("id ,name as realname").leftJoin("user", "user.id = user.id").Where("id = ?", 1).Find(&UserModel{})

}

func Where(a ...any) string {
	return ""
}
