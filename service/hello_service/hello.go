/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/31 11:55 上午
 * @Description: 业务层 hello
 */
package hello_service

import (
	"github.com/togettoyou/go-miaosha/models"
)

type Hello struct {
	PageSize int
	Page     int
}

func (h *Hello) Get() ([]models.Hello, error) {
	hello, err := models.GetHello(h.PageSize, h.Page)
	if err != nil {
		return nil, err
	}
	return hello, nil
}
