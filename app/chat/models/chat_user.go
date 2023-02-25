package models

import (
	"time"

	"go-admin/common/models"
)

type ChatUser struct {
	models.Model

	Nickname      string    `json:"nickname" gorm:"type:varchar(255);comment:昵称"`
	Username      string    `json:"username" gorm:"type:varchar(20);comment:登录账号"`
	Initial       string    `json:"initial" gorm:"type:char(1);comment:首字母"`
	Password      string    `json:"password" gorm:"type:varchar(255);comment:密码"`
	Avatar        string    `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Phone         string    `json:"phone" gorm:"type:char(11);comment:手机号"`
	Sex           string    `json:"sex" gorm:"type:tinyint(2);comment:性别：1男2女3未知"`
	Birthday      string    `json:"birthday" gorm:"type:date;comment:生日"`
	School        string    `json:"school" gorm:"type:varchar(255);comment:学校"`
	Signature     string    `json:"signature" gorm:"type:varchar(255);comment:个性签名"`
	Education     string    `json:"education" gorm:"type:tinyint(2);comment:学历：1无2小学3初中4高中5专科6本科7硕士8博士9博士后"`
	Major         string    `json:"major" gorm:"type:varchar(255);comment:专业"`
	LastLoginTime time.Time `json:"lastLoginTime" gorm:"type:datetime;comment:LastLoginTime"`
	models.ModelTime
	models.ControlBy
}

func (ChatUser) TableName() string {
	return "chat_user"
}

func (e *ChatUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ChatUser) GetId() interface{} {
	return e.Id
}
