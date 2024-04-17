package domain

import (
	"github.com/gin-gonic/gin"
)

type Endpoint func(ctx *gin.Context, requestBody interface{}) (response interface{}, err error)

type User struct {
	ID      uint   `json:"id" gorm:"column:id;primaryKey;check:id>0;<-:false"`
	Name    string `json:"name" gorm:"column:name;size:255;not null"`
	Email   string `json:"email" gorm:"column:email;size:255;not null;unique_index"`
	Age     int    `json:"age" gorm:"column:age;check:age >= 18"`
	Created int64  `json:"created,omitempty" gorm:"column:created;autoCreateTime"`
	Updated int64  `json:"updated,omitempty" gorm:"column:updated;autoUpdateTime"`
	Deleted int64  `json:"-" gorm:"column:deleted"`
}
