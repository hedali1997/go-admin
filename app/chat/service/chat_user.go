package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/chat/models"
	"go-admin/app/chat/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ChatUser struct {
	service.Service
}

// GetPage 获取ChatUser列表
func (e *ChatUser) GetPage(c *dto.ChatUserGetPageReq, p *actions.DataPermission, list *[]models.ChatUser, count *int64) error {
	var err error
	var data models.ChatUser

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ChatUserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ChatUser对象
func (e *ChatUser) Get(d *dto.ChatUserGetReq, p *actions.DataPermission, model *models.ChatUser) error {
	var data models.ChatUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetChatUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ChatUser对象
func (e *ChatUser) Insert(c *dto.ChatUserInsertReq) error {
	var err error
	var data models.ChatUser
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ChatUserService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ChatUser对象
func (e *ChatUser) Update(c *dto.ChatUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ChatUser{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ChatUserService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ChatUser
func (e *ChatUser) Remove(d *dto.ChatUserDeleteReq, p *actions.DataPermission) error {
	var data models.ChatUser

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveChatUser error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
