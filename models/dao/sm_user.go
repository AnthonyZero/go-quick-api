package dao

import (
	"time"

	"github.com/anthonyzero/go-quick-api/pkg/db"
)

// SmUser 用户表
type SmUser struct {
	UserID     int64     `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"userId"`   // 用户ID
	UserName   string    `gorm:"column:user_name;type:varchar(128);not null" json:"userName"`        // 用户名称
	UserCode   string    `gorm:"unique;column:user_code;type:varchar(128);not null" json:"userCode"` // 登录账号
	Password   string    `gorm:"column:password;type:varchar(128);not null" json:"password"`         // 用户密码
	DepartID   int64     `gorm:"column:depart_id;type:bigint(20);not null" json:"departId"`          // 部门ID
	Salt       string    `gorm:"column:salt;type:varchar(64)" json:"salt"`                           // 盐
	Avatar     string    `gorm:"column:avatar;type:varchar(128)" json:"avatar"`                      // 头像图片
	UserType   int       `gorm:"column:user_type;type:int(11);default:1" json:"userType"`            // 用户类型 1系统用户
	Email      string    `gorm:"column:email;type:varchar(128)" json:"email"`                        // 邮箱
	Mobile     string    `gorm:"column:mobile;type:varchar(32)" json:"mobile"`                       // 手机号码
	UserDesc   string    `gorm:"column:user_desc;type:varchar(500)" json:"userDesc"`
	State      int       `gorm:"column:state;type:int(11);not null" json:"state"`             // 状态 0 无效1 正常
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"createTime"` // 创建时间
}

func (s *SmUser) CreateUser() error {
	if err := db.GetDefault().Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *SmUser) DeleteUser() error {
	return db.GetDefault().Where("user_id = ?", s.UserID).Delete(&SmUser{}).Error
}
