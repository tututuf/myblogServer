package apiRouters

import (
	"fmt"
	"myblog/modules"

	"github.com/gin-gonic/gin"
)

type Admininfo struct {
	Id        int    `json:"id"`
	AdminUuid int    `json:"admin_name"`
	Username  string `json:"admin_age"`
	NickName  string `json:"nick_name"`
}

func SetAdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/get_admin_info", func(c *gin.Context) {
			adminUuid := c.Query("admin_uuid")
			fmt.Println(adminUuid)
			var adminInfo *Admininfo
			modules.DB.First(&adminInfo, "admin_uuid = ?", adminUuid)
			// admin := &AdminInfo{
			// 	Id:       1,
			// 	NickName: "张三",
			// }
			c.JSON(200, adminInfo)
		})
	}
}
