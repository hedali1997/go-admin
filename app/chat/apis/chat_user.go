package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/chat/models"
	"go-admin/app/chat/service"
	"go-admin/app/chat/service/dto"
	"go-admin/common/actions"
)

type ChatUser struct {
	api.Api
}

// GetPage 获取前台用户列表
// @Summary 获取前台用户列表
// @Description 获取前台用户列表
// @Tags 前台用户
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ChatUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/chat-user [get]
// @Security Bearer
func (e ChatUser) GetPage(c *gin.Context) {
	req := dto.ChatUserGetPageReq{}
	s := service.ChatUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.ChatUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取前台用户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取前台用户
// @Summary 获取前台用户
// @Description 获取前台用户
// @Tags 前台用户
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ChatUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/chat-user/{id} [get]
// @Security Bearer
func (e ChatUser) Get(c *gin.Context) {
	req := dto.ChatUserGetReq{}
	s := service.ChatUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.ChatUser

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取前台用户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建前台用户
// @Summary 创建前台用户
// @Description 创建前台用户
// @Tags 前台用户
// @Accept application/json
// @Product application/json
// @Param data body dto.ChatUserInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/chat-user [post]
// @Security Bearer
func (e ChatUser) Insert(c *gin.Context) {
	req := dto.ChatUserInsertReq{}
	s := service.ChatUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建前台用户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改前台用户
// @Summary 修改前台用户
// @Description 修改前台用户
// @Tags 前台用户
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ChatUserUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/chat-user/{id} [put]
// @Security Bearer
func (e ChatUser) Update(c *gin.Context) {
	req := dto.ChatUserUpdateReq{}
	s := service.ChatUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改前台用户失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除前台用户
// @Summary 删除前台用户
// @Description 删除前台用户
// @Tags 前台用户
// @Param data body dto.ChatUserDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/chat-user [delete]
// @Security Bearer
func (e ChatUser) Delete(c *gin.Context) {
	s := service.ChatUser{}
	req := dto.ChatUserDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除前台用户失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
