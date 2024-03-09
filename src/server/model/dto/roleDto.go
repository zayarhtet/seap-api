package dto

type RoleDto struct {
	RoleId uint   `gorm:"primary_key" json:"roleId"`
	Name   string `json:"name"`
}

func (RoleDto) TableName() string {
	return "role"
}
