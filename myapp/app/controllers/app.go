package controllers

import (
	data "github.com/jayatwork/team27live/myapp/app/data"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("The name here is required")
	c.Validation.MinSize(myName, 3).Message("The name given is not long enough")
	c.Validation.MaxSize(myName, 12).Message("this is toooo long ")
	data.initdb()

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
