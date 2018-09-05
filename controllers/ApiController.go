package controllers

import (
	"github.com/kataras/iris"
	"demo/models"
	"github.com/kataras/iris/context"
)

type ApiController struct {
	Ctx    iris.Context
	RetMsg models.RetMsg
}

func (c *ApiController) Success(data interface{}) {
	msg := c.RetMsg
	msg.Data = data

	opts := context.DefaultJSONOptions
	opts.UnescapeHTML=true

	c.Ctx.JSON(&msg,opts)
}

func (c *ApiController) Fail(errMsg error) {
	msg := c.RetMsg
	msg.ErrNo = 1
	msg.ErrMsg = errMsg
	c.Ctx.JSON(&msg)
}

func init()  {

}