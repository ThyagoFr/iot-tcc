package entity

import "gorm.io/gorm"

type Notification struct {
	gorm.Model `json:"-"`
	ID         string `gorm:"type:uuid;primarykey"`
	DeviceID  string
	UserID    string
	Parameter string
	Condition Condition
	Value     float64
}

type Condition string

const (
	GreaterThan        Condition = "greaterThan"
	GreaterThanOrEqual Condition = "greaterThanOrEqual"
	LessThan           Condition = "lessThan"
	LessThanOrEqual    Condition = "lessThanOrEqual"
	EqualTo            Condition = "equalTo"
)

func (c Condition) String() string {
	return string(c)
}
