/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/31 11:58 上午
 * @Description: 数据层 hello
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// 定义表模型
type Hello struct {
	Name string
}

// 设置表名
func (h Hello) TableName() string {
	return "hello"
}

// 创建初始化表
func InitHello() {
	//如果数据库中查询不到hello表，则创建
	if !db.HasTable(&Hello{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Hello{}).Error; err != nil {
			panic(err)
		}
		insertHelloData()
	}
}

func insertHelloData() {
	data := []string{"One", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten"}
	for _, v := range data {
		hello := Hello{Name: v}
		if err := db.Create(&hello).Error; err != nil {
			continue
		}
	}
}

// 根据分页获取
// pageSize 页面大小
// page 页码
func GetHello(pageSize int, page int) ([]Hello, error) {
	var (
		hello []Hello
		err   error
	)
	if pageSize > 0 && page > 0 {
		err = db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&hello).Error
	} else {
		err = db.Find(&hello).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return hello, nil
}
