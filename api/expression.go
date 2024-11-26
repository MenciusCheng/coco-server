package api

import (
	"coco-server/model/common/response"
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/expression/run", ExpressionApi.Run)
}

type expressionApi struct{}

var ExpressionApi = new(expressionApi)

func (a *expressionApi) Run(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(ExpressionCompileReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_ = ctx

	cEnv := map[string]interface{}{
		"param":   &ExpressionEventParam{},
		"print":   fmt.Println,
		"sprintf": fmt.Sprintf,
		"pattern": req.Pattern,
	}

	program, err := expr.Compile(req.Code, expr.Env(cEnv))
	if err != nil {
		panic(err)
	}

	eventParam := &ExpressionEventParam{
		RelId:     5,
		PlayerId:  123456,
		Num:       1,
		TriggerId: 111,
		TargetId:  222,
		RoomId:    88,
	}

	rEnv := map[string]interface{}{
		"param":   eventParam,
		"print":   fmt.Println,
		"sprintf": fmt.Sprintf,
		"pattern": req.Pattern,
	}

	output, err := expr.Run(program, rEnv)
	if err != nil {
		panic(err)
	}
	res := map[string]interface{}{
		"output": output,
	}
	response.OkWithData(res, c)
}

type ExpressionCompileReq struct {
	Code    string `json:"code"`
	Pattern string `json:"pattern"`
}

type ExpressionEventParam struct {
	RelId     int32 `json:"rel_id"`     // 子活动id
	PlayerId  int64 `json:"player_id"`  // 用户id
	Num       int64 `json:"num"`        // 数量
	TriggerId int64 `json:"trigger_id"` // 触发id 礼物id
	TargetId  int64 `json:"target_id"`  // 目标对象id，送礼事件时为收礼人id
	RoomId    int64 `json:"room_id"`    // 房间id
}
