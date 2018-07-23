package controllers

import (
	"fmt"
	"app"
)

type SiteController struct {
	app.Controller
}

func (p SiteController) Index() {
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
