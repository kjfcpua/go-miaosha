/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/31 2:38 下午
 * @Description: 用户表
 */
package models

import "time"

// 定义用户表
type User struct {
	Username  string `gorm:"primary_key;column:username" form:"username" json:"username" comment:"昵称/登陆用户名" columnType:"varchar(50)" dataType:"varchar" columnKey:"UNI"`
	Mobile    string `gorm:"column:mobile" form:"mobile" json:"mobile" comment:"手机号码" columnType:"varchar(11)" dataType:"varchar" columnKey:"UNI"`
	Salt      string `gorm:"column:salt" form:"salt" json:"salt" comment:"混淆盐" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	Password  string `gorm:"column:password" form:"password" json:"password" comment:"密码" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	RoleId    uint   `gorm:"column:role_id" form:"role_id" json:"role_id" comment:"角色ID:0-超级用户,1-普通用户" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	NickName  string `gorm:"column:nick_name" form:"nick_name" json:"nick_name" comment:"真实姓名"`
	Avatar    string `gorm:"column:avatar" form:"avatar" json:"avatar" comment:"用户头像" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	CreatedAt time.Time
}

// 设置表名
func (u User) TableName() string {
	return "user"
}

// 创建初始化表
func InitUser() {
	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}
