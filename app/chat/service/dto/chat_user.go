package dto

import (
	"time"

	"go-admin/app/chat/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ChatUserGetPageReq struct {
	dto.Pagination `search:"-"`
	ChatUserOrder
}

type ChatUserOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:chat_user"`
	Nickname      string `form:"nicknameOrder"  search:"type:order;column:nickname;table:chat_user"`
	Username      string `form:"usernameOrder"  search:"type:order;column:username;table:chat_user"`
	Initial       string `form:"initialOrder"  search:"type:order;column:initial;table:chat_user"`
	Password      string `form:"passwordOrder"  search:"type:order;column:password;table:chat_user"`
	Avatar        string `form:"avatarOrder"  search:"type:order;column:avatar;table:chat_user"`
	Phone         string `form:"phoneOrder"  search:"type:order;column:phone;table:chat_user"`
	Sex           string `form:"sexOrder"  search:"type:order;column:sex;table:chat_user"`
	Birthday      string `form:"birthdayOrder"  search:"type:order;column:birthday;table:chat_user"`
	School        string `form:"schoolOrder"  search:"type:order;column:school;table:chat_user"`
	Signature     string `form:"signatureOrder"  search:"type:order;column:signature;table:chat_user"`
	Education     string `form:"educationOrder"  search:"type:order;column:education;table:chat_user"`
	Major         string `form:"majorOrder"  search:"type:order;column:major;table:chat_user"`
	LastLoginTime string `form:"lastLoginTimeOrder"  search:"type:order;column:last_login_time;table:chat_user"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:chat_user"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:chat_user"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:chat_user"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:chat_user"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:chat_user"`
}

func (m *ChatUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ChatUserInsertReq struct {
	Id            int       `json:"-" comment:""` //
	Nickname      string    `json:"nickname" comment:"昵称"`
	Username      string    `json:"username" comment:"登录账号"`
	Initial       string    `json:"initial" comment:"首字母"`
	Password      string    `json:"password" comment:"密码"`
	Avatar        string    `json:"avatar" comment:"头像"`
	Phone         string    `json:"phone" comment:"手机号"`
	Sex           string    `json:"sex" comment:"性别：1男2女3未知"`
	Birthday      string    `json:"birthday" comment:"生日"`
	School        string    `json:"school" comment:"学校"`
	Signature     string    `json:"signature" comment:"个性签名"`
	Education     string    `json:"education" comment:"学历：1无2小学3初中4高中5专科6本科7硕士8博士9博士后"`
	Major         string    `json:"major" comment:"专业"`
	LastLoginTime time.Time `json:"lastLoginTime" comment:""`
	common.ControlBy
}

func (s *ChatUserInsertReq) Generate(model *models.ChatUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Nickname = s.Nickname
	model.Username = s.Username
	model.Initial = s.Initial
	model.Password = s.Password
	model.Avatar = s.Avatar
	model.Phone = s.Phone
	model.Sex = s.Sex
	model.Birthday = s.Birthday
	model.School = s.School
	model.Signature = s.Signature
	model.Education = s.Education
	model.Major = s.Major
	model.LastLoginTime = s.LastLoginTime
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ChatUserInsertReq) GetId() interface{} {
	return s.Id
}

type ChatUserUpdateReq struct {
	Id            int       `uri:"id" comment:""` //
	Nickname      string    `json:"nickname" comment:"昵称"`
	Username      string    `json:"username" comment:"登录账号"`
	Initial       string    `json:"initial" comment:"首字母"`
	Password      string    `json:"password" comment:"密码"`
	Avatar        string    `json:"avatar" comment:"头像"`
	Phone         string    `json:"phone" comment:"手机号"`
	Sex           string    `json:"sex" comment:"性别：1男2女3未知"`
	Birthday      string    `json:"birthday" comment:"生日"`
	School        string    `json:"school" comment:"学校"`
	Signature     string    `json:"signature" comment:"个性签名"`
	Education     string    `json:"education" comment:"学历：1无2小学3初中4高中5专科6本科7硕士8博士9博士后"`
	Major         string    `json:"major" comment:"专业"`
	LastLoginTime time.Time `json:"lastLoginTime" comment:""`
	common.ControlBy
}

func (s *ChatUserUpdateReq) Generate(model *models.ChatUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Nickname = s.Nickname
	model.Username = s.Username
	model.Initial = s.Initial
	model.Password = s.Password
	model.Avatar = s.Avatar
	model.Phone = s.Phone
	model.Sex = s.Sex
	model.Birthday = s.Birthday
	model.School = s.School
	model.Signature = s.Signature
	model.Education = s.Education
	model.Major = s.Major
	model.LastLoginTime = s.LastLoginTime
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ChatUserUpdateReq) GetId() interface{} {
	return s.Id
}

// ChatUserGetReq 功能获取请求参数
type ChatUserGetReq struct {
	Id int `uri:"id"`
}

func (s *ChatUserGetReq) GetId() interface{} {
	return s.Id
}

// ChatUserDeleteReq 功能删除请求参数
type ChatUserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ChatUserDeleteReq) GetId() interface{} {
	return s.Ids
}
