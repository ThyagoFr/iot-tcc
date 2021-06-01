package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID        primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (e *Entity) BeforeUpdate() {
	e.UpdatedAt = time.Now()
}

func (e *Entity) BeforeSave() {
	e.CreatedAt = time.Now()
}

func (e *Entity) BeforeDelete() {
	e.UpdatedAt = time.Now()
	e.DeletedAt = time.Now()
}
