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
	Username  string    `gorm:"primary_key;column:username;type:varchar(12)" comment:"昵称/登陆用户名"`
	Mobile    string    `gorm:"unique;not null;column:mobile;type:varchar(11)" comment:"手机号码"`
	Salt      string    `gorm:"not null;column:salt;type:varchar(255)" comment:"混淆盐"`
	Password  string    `gorm:"not null;column:password;type:varchar(28)" comment:"密码"`
	RoleId    uint      `gorm:"not null;column:role_id;type:int(10) unsigned" comment:"角色ID:0-超级用户,1-普通用户"`
	NickName  string    `gorm:"column:nick_name;type:varchar(12)" comment:"真实姓名"`
	Avatar    string    `gorm:"column:avatar;type:varchar(255)" comment:"用户头像"`
	CreatedAt time.Time `gorm:"not null" comment:"注册时间"`
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
