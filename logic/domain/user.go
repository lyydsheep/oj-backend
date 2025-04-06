package domain

type User struct {
	ID           int64   `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:id" json:"id"`                              // id
	UID          string  `gorm:"column:uid;type:varchar(32);not null;comment:uid" json:"uid"`                                           // uid
	UserAccount  string  `gorm:"column:user_account;type:varchar(256);not null;comment:账号" json:"user_account"`                         // 账号
	UserPassword string  `gorm:"column:user_password;type:varchar(512);not null;comment:密码" json:"user_password"`                       // 密码
	UnionID      *string `gorm:"column:union_id;type:varchar(256);index:idx_unionId,priority:1;comment:微信开放平台id" json:"union_id"`       // 微信开放平台id
	OpenID       *string `gorm:"column:open_id;type:varchar(256);comment:公众号openId" json:"open_id"`                                     // 公众号openId
	Username     *string `gorm:"column:username;type:varchar(256);comment:用户昵称" json:"username"`                                        // 用户昵称
	UserAvatar   *string `gorm:"column:user_avatar;type:varchar(1024);comment:用户头像" json:"user_avatar"`                                 // 用户头像
	UserProfile  *string `gorm:"column:user_profile;type:varchar(512);comment:用户简介" json:"user_profile"`                                // 用户简介
	UserRole     string  `gorm:"column:user_role;type:varchar(256);not null;default:user;comment:用户角色：user/admin/ban" json:"user_role"` // 用户角色：user/admin/ban
}
