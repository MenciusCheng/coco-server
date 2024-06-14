package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"github.com/gin-gonic/gin"
	"regexp"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/regexp/find", RegexpApi.Find)
	routerGroup.POST("/regexp/replace", RegexpApi.Replace)
}

type regexpApi struct{}

var RegexpApi = new(regexpApi)

func (a *regexpApi) Find(c *gin.Context) {
	//ctx := c.Request.Context()
	req := &model.RegexpFindReq{}
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	expr, err := regexp.Compile(req.Expr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	s := req.S

	res := &model.RegexpFindRes{
		MatchString:                expr.MatchString(s),
		FindString:                 expr.FindString(s),
		FindStringIndex:            expr.FindStringIndex(s),
		FindStringSubmatch:         expr.FindStringSubmatch(s),
		FindStringSubmatchIndex:    expr.FindStringSubmatchIndex(s),
		FindAllString:              expr.FindAllString(s, -1),
		FindAllStringIndex:         expr.FindAllStringIndex(s, -1),
		FindAllStringSubmatch:      expr.FindAllStringSubmatch(s, -1),
		FindAllStringSubmatchIndex: expr.FindAllStringSubmatchIndex(s, -1),
	}
	response.OkWithData(res, c)
	return
}

func (a *regexpApi) Replace(c *gin.Context) {
	//ctx := c.Request.Context()
	req := &model.RegexpReplaceReq{}
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	expr, err := regexp.Compile(req.Expr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	s := req.S

	res := &model.RegexpReplaceRes{
		MatchString: expr.MatchString(s),
		Des:         expr.ReplaceAllString(s, req.Repl),
	}
	response.OkWithData(res, c)
	return
}
