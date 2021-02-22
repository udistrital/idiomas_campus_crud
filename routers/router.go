// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/idiomas_campus_crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/soporte_conocimiento_idioma",
			beego.NSInclude(
				&controllers.SoporteConocimientoIdiomaController{},
			),
		),

		beego.NSNamespace("/idioma",
			beego.NSInclude(
				&controllers.IdiomaController{},
			),
		),

		beego.NSNamespace("/clasificacion_nivel_idioma",
			beego.NSInclude(
				&controllers.ClasificacionNivelIdiomaController{},
			),
		),

		beego.NSNamespace("/valor_nivel_idioma",
			beego.NSInclude(
				&controllers.ValorNivelIdiomaController{},
			),
		),

		beego.NSNamespace("/conocimiento_idioma",
			beego.NSInclude(
				&controllers.ConocimientoIdiomaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
