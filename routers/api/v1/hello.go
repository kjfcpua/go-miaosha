/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/3 11:24 下午
 * @Description: 路由层|控制层 hello
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/togettoyou/go-miaosha/pkg/app"
	"github.com/togettoyou/go-miaosha/service/hello_service"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary 测试输出Hello
// @Produce  json
// @Param pageSize path int true "pageSize"
// @Param page path int true "page"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/hello/{pageSize}/{page} [get]
func Hello(c *gin.Context) {
	appG := app.Gin{C: c}
	helloService := hello_service.Hello{
		PageSize: com.StrTo(c.Param("pageSize")).MustInt(),
		Page:     com.StrTo(c.Param("page")).MustInt(),
	}
	hello, err := helloService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "查询出错", err)
		return
	}
	appG.Response(http.StatusOK, "成功", hello)
}
