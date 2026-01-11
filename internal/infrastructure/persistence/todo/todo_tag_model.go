package todo

import "time"

type TodoTagModel struct {
	TodoID    int       `gorm:"primaryKey;column:todo_id"`
	TagID     int       `gorm:"primaryKey;column:tag_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (TodoTagModel) TableName() string {
	return "todo_tags"
}

