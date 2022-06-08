package models

import (
	"time"
)

type Stock struct {
	//gorm.Model
	Id int `gorm:"AUTO_INCREMENT"` //自增
	Market string `gorm:"size:2"` // string默认长度2
	Code string `gorm:"size:6"`
	Name string `gorm:"size:8"`
	Date time.Time
	Open float64
	Close float64
	High float64
	Low float64
	Volume float64
	OutstandingShare float64
	Turnover float64
}

func (Stock) TableName() string {
	return "stock"
}


type User struct {
	Id int `gorm:"AUTO_INCREMENT"`
	// 默认解析出来的列名为user_name
	UserName string `gorm:"column:username; size:150"`
	Password string `gorm:"column:password; size:128"`
	Email string `gorm:"size254"`
	IsActive int8 `gorm:"default:1"` // 没有默认值自动会给类型零值 int类型为0
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "user"
}

type Role struct {
	Id int `gorm:"AUTO_INCREMENT"`
	// 默认解析出来的列名为user_name
	Name string `gorm:"column:name; size:150"`
	IsActive int8 `gorm:"default:1"` // 没有默认值自动会给类型零值 int类型为0
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (Role) TableName() string {
	return "role"
}

// Department Has many
type Department struct {
	Id int `gorm:"AUTO_INCREMENT"`
	// 默认解析出来的列名为user_name
	Code string `gorm:"column:code; size:150"`
	Name string `gorm:"column:name; size:150"`
	ParentId int32 `gorm:"column:parent_id"`
	IsDeleted int8 `gorm:"default:1"` // 没有默认值自动会给类型零值 int类型为0
	CreatedBy int16 `gorm:"column:created_by"`
	UpdatedBy int16 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	//Doctor []Doctor
}

func (Department) TableName() string {
	return "department"
}

// Doctor example Belongs to Department
type Doctor struct {
	Id int `gorm:"AUTO_INCREMENT"`
	// 默认解析出来的列名为user_name
	Code string `gorm:"column:code; size:150"`
	Name string `gorm:"column:name; size:150"`
	DeptId int32 `gorm:"column:dept_id"`
	IsDeleted int8 `gorm:"default:1"` // 没有默认值自动会给类型零值 int类型为0

	// todo 数据库设置的默认值不起作用吗
	CreatedBy *int16 `gorm:"column:created_by"` // 指针类型时，没有设置默认值，数据库为not null，当不传入该值时，报错，数据库为null, 不传则为null
	UpdatedBy int16 `gorm:"column:updated_by"` // 值类型时，没有设置默认值，不管数据库有没有默认值，当不传入该值时，都会存入对应类型零值
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	Department Department `gorm:"foreignKey:DeptId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 关联表Department未初始化，创建Doctor时不会创建Department
}

func (Doctor) TableName() string {
	return "doctor"
}

type Group struct {
	Id int `gorm:"AUTO_INCREMENT"`
	// 默认解析出来的列名为user_name
	Name string `gorm:"column:name; size:150"`
	ParentId int32 `gorm:"column:parent_id"`
	Content string `gorm:"column:content"` // 没有默认值自动会给类型零值 int类型为0
	CreatedBy int16 `gorm:"column:created_by"`
	UpdatedBy int16 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (Group) TableName() string {
	return "group"
}


