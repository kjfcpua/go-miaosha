/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/31 4:22 下午
 * @Description: 商品表
 */
package models

import (
	"github.com/shopspring/decimal"
	"time"
)

// 定义商品表
type Goods struct {
	GoodsId     string          `gorm:"primary_key;column:goods_id;type:varchar(50)" comment:"商品ID"`
	GoodsName   string          `gorm:"not null;column:goods_name;type:varchar(20)" comment:"商品名称"`
	GoodsTitle  string          `gorm:"not null;column:goods_title;type:varchar(80)" comment:"商品标题"`
	GoodsImg    string          `gorm:"not null;column:goods_img;type:varchar(255)" comment:"商品图片"`
	GoodsDetail string          `gorm:"not null;column:goods_detail;type:varchar(255)" comment:"商品详情"`
	GoodsPrice  decimal.Decimal `gorm:"not null;column:goods_price;type:decimal(20,2)"  comment:"商品价格" `
	GoodsStock  uint            `gorm:"not null;column:goods_stock;type:int(10) unsigned" comment:"商品库存"`
	CreatedAt   time.Time       `gorm:"not null" comment:"创建时间"`
	UpdatedAt   time.Time       `gorm:"not null" comment:"更新时间"`
}

// 设置表名
func (u Goods) TableName() string {
	return "goods"
}

// 创建初始化表
func InitGoods() {
	if !db.HasTable(&Goods{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Goods{}).Error; err != nil {
			panic(err)
		}
	}
}
