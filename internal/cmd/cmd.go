package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gview"

	"gov2panel/internal/controller/admin"
	"gov2panel/internal/controller/public"
	"gov2panel/internal/controller/server_api"
	user_c "gov2panel/internal/controller/user"
	"gov2panel/internal/logic/user"
	"gov2panel/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {

				d, err := service.Setting().GetSettingAllMap()
				if err != nil {
					print(err.Error())
				}
				group.Middleware(func(r *ghttp.Request) { //设置参数
					// 中间件处理逻辑
					d, err := service.Setting().GetSettingAllMap()
					if err != nil {
						print(err.Error())
					}

					r.Assigns(gview.Params{
						"setting": d,
					})

					r.Middleware.Next()
				})

				group.Middleware(ghttp.MiddlewareHandlerResponse) //处理处理程序响应对象及其错误的默认中间件
				group.Bind(
					public.NewV1(),
				)
				group.Middleware(user.Middleware().CORS) //跨域处理

				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(func(r *ghttp.Request) { //设置参数
						// 中间件处理逻辑
						d, err := service.Setting().GetSettingAllMap()
						if err != nil {
							print(err.Error())
						}

						if r.Get("token").String() != d["api_key"].String() {
							r.Response.WriteExit(`{"message": "token is error"}`)
						}

						r.Middleware.Next()
					}) //token
					group.Bind(
						server_api.NewV1(),
					)

				})

				group.Group("/"+d["admin_path"].String(), func(group *ghttp.RouterGroup) {
					group.Middleware(user.Middleware().Auth) //权限处理
					group.Bind(
						admin.NewV1(),
					)

				})

				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Middleware(user.Middleware().Auth) //权限处理
					group.Bind(
						user_c.NewV1(),
					)

				})

			})
			s.Run()
			return nil
		},
	}
)