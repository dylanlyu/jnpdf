package api

import (
	"jnpdf/container"

	"github.com/astaxie/beego"
)

// PrintController is print to pdf controller
type PrintController struct {
	beego.Controller
}

// Print is sed printer data
// @Title 資料帶入印表氣
// @Description 資料帶入印表氣API
// @Success 200	{object} schema.Reply
// @Failure 403 body is empty
// @router /print/?:id [get]
func (controller *PrintController) Print() {
	id := ""

	if controller.Ctx.Input.Query("id") != "" {
		id = controller.Ctx.Input.Query("id")
	} else if controller.GetString(":id") != "" {
		id = controller.GetString(":id")
	}

	reply := container.Print(id)

	controller.Data["json"] = reply
	controller.ServeJSON()
}
