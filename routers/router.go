// @APIVersion 1.0.0
// @Title API doc
// @Description http heard包含token(Authorization:Bearer token)
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jnpdf/api"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/exec",
			beego.NSInclude(
				&api.PrintController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
