package domain

import (
	"time"
)

// 硬件信息
type Device struct {
	ID         int64        `json:"-" gorm:"primary_key;column:tid;unique_index:devices_pkey"`
	SN         string       `json:"sn,omitempty"`
	IMEI       string       `json:"imei,omitempty" gorm:"column:imei"`
	SEQ        string       `json:"seq,omitempty" gorm:"column:seq"`
	IOSVersion string       `json:"ios_version,omitempty" gorm:"column:ios_version"`
	MAC        string       `json:"mac,omitempty" gorm:"column:mac"`
	WIFI       string       `json:"wifi,omitempty" gorm:"column:wifi"`
	Model      string       `json:"model,omitempty"`
	IDFA       string       `json:"idfa,omitempty" gorm:"column:idfa"`
	IDFV       string       `json:"idfv,omitempty" gorm:"column:idfv"`
	Region     string       `json:"region,omitempty"`
	ModelNum   string       `json:"model_num,omitempty"`
	DeviceName string       `json:"device_name,omitempty"`
	Used       int          `json:"-"`
	Status     DeviceStatus `json:"-"`
	CreatedAt  *time.Time   `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt  *time.Time   `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*Device) TableName() string {
	return "devices"
}

type DeviceStatus int

const (
	_ DeviceStatus = iota
	DeviceEnable
	DeviceDisabled
)

// GPS信息
type GPSLocation struct {
	ID        int64   `json:"-" gorm:"primary_key;column:tid;unique_index:gps_locations_pkey"`
	GPSID     string  `json:"-" gorm:"column:gps_id"`
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Province  string  `json:"province,omitempty"`
	City      string  `json:"city,omitempty"`
	Type      GPSType `json:"type,omitempty"`
}

func (*GPSLocation) TableName() string {
	return "gpss"
}

type GPSType int

const (
	_ GPSType = iota
	GPSTypeNormal
	GPSTypeCentral
)

type PhotoGroup struct {
	ID       int64        `json:"-" gorm:"primary_key;column:tid;unique_index:photo_groups_pkey"`
	PhotosID string       `json:"photos_id" gorm:"unique_index:photos_id_idx"`
	Chat     ChatType     `json:"chat"`
	Random   int          `json:"-" gorm:"index:random_idx"`
	Status   PhotosStatus `json:"photos_status"`
}

func (*PhotoGroup) TableName() string {
	return "photo_groups"
}

type PhotosStatus int

const (
	_                   PhotosStatus = iota
	PhotosStatusFree                 // 可用
	PhotosStatusUsed                 // 已用
	PhotosStatusDisable              // 禁用
)

// 套图信息
type Photo struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	PhotosID string `json:"photos_id,omitempty"`
	Seq      int    `json:"seq,omitempty"`
	URL      string `json:"url,omitempty"`
}

func (*Photo) TableName() string {
	return "photos"
}

type NickName struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:nick_names_pkey"`
	NickName string `json:"-"`
}

func (*NickName) TableName() string {
	return "nick_names"
}

// Reply
type Reply struct {
	ID        int64      `json:"-" gorm:"primary_key;column:tid;unique_index:replys_pkey"`
	ReplyID   string     `json:"reply_id"`
	Chat      ChatType   `json:"chat"`
	ReplyType int        `json:"reply_type"`
	Content   string     `json:"content"`
	Used      int        `json:"-"`
	Free      int        `json:"-"`
	Priority  int        `json:"-"`
	Group     string     `json:"-"`
	Status    int        `json:"-"`
	CreatedAt *time.Time `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt *time.Time `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*Reply) TableName() string {
	return "replys"
}

type ReplyType int

const (
	_ ReplyType = iota
	ReplyTypeWord
	ReplyTypePhoto
)

type ReplyStatus int

const (
	_ ReplyStatus = iota
	ReplyStatusEnable
	ReplyStatusDisable
)

type AccountType int

const (
	_     AccountType = iota
	QQ                // 1 QQ账号
	Phone             // 2 手机账号
)

type GenderType int

const (
	_      GenderType = iota
	Man               // 1 男
	Female            // 2 女
)

type AccountStatus int

const (
	_                        AccountStatus = iota
	AccountStatusUnRegister                // 1 未注册
	AccountStatusFree                      // 2 已注册(可用)
	AccountStatusRegistering               // 3 正在注册
	AccountStatusDisabled                  // 4 被禁用
	AccountStatusOnline                    // 5 在线中
	AccountStatusLocked                    // 6 锁定中(正在注册)
)

// AccountReply
type ChatReply struct {
	ID          int64       `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	Account     string      `json:"account"`
	AccountType AccountType `json:"account_type"`
	Chat        ChatType    `json:"chat"`
	ReplyID     string      `json:"reply_id"`
	CreatedAt   *time.Time  `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt   *time.Time  `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*ChatReply) TableName() string {
	return "chat_replys"
}

type ChatType int

const (
	_ ChatType = iota
	ChatTypeCommon
	ChatTypeMomo
	ChatTypeTantan
)
