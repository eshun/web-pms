package controllers

import (
	"web-pms/models"

	"fmt"
)

type EventController struct {
	ApiController
}

func (c *EventController) Post() {
	var model models.Event
	err := c.Ctx.ReadJSON(&model)
	if err != nil {
		c.Fail(err)
	} else {
		_, err = models.InsertEvent(&model)
		fmt.Println(model)
		if (model.EventType == "performance") {
			model.PageTimes.EventId = model.Id
			fmt.Println(model.PageTimes)
			_, err = models.InsertPageTimes(&model.PageTimes)

			//resources,_:=model.Resources.([]models.Resource)//通过断言实现类型转换
			fmt.Println(model.Resources)
			_, err = models.InsertResources(&model.Resources, model.Id)
		} else if (model.EventType == "httpError") {
			model.Req.EventId = model.Id
			fmt.Println(model.Req)
			_, err = models.InsertReq(&model.Req)

			model.Res.EventId = model.Id
			fmt.Println(model.Res)
			_, err = models.InsertRes(&model.Res)
		} else if (model.EventType == "resourceError" || model.EventType == "click") {
			model.Target.EventId = model.Id
			fmt.Println(model.Target)
			_, err = models.InsertTarget(&model.Target)
		}

		fmt.Println("err", err)
		if err != nil {
			c.Fail(err)
		} else {
			c.Success("")
		}
	}
}