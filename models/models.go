package models

import (
	"time"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"

)

// 分类
type Category struct {
	ID              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64  `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

// 文章
type Topic struct {
	ID               int64
	UID              string
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time
	Updated          time.Time
	Views            int64
	Author           string
	ReplayTime       time.Time
	ReplayCount      int64
	ReplayLastUserId int64
}
