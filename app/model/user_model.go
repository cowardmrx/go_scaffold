package model

type UserModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name;type:char(30);not null"`
	Age  int    `gorm:"column:age;type:int(10);not null"`
	Timestamp
}

func (UserModel) TableName() string {
	return "user"
}
