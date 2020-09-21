package middleman

import (
	"github.com/astaxie/beego/context"
)

func FilterWebUser(ctx *context.Context) {
	//beego.Debug(ctx.Request.Header)
	//if ctx.Request.Header.Get("Authorization") != "" {
	//	token := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)
	//
	//	if len(token) != 2 || token[1] == "" {
	//		beego.Info("1URL:User=error1")
	//		ctx.Redirect(401, "/token")
	//	} else if !util.CheckToken(token[1]) {
	//		beego.Info("1URL:User=error2")
	//		ctx.Redirect(401, "/token")
	//	}
	//} else {
	//	beego.Info("2URL=error3")
	//	//ctx.Abort(401,"")
	//	ctx.Redirect(401, "/")
	//}
}
