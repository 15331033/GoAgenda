package entities

import (
	"fmt"
	"time"
)

// UserInfo .
type UserInfo struct {
    UID        int  `xorm:"pk autoincr 'id'"`//语义标签
    UserName   string
    Password   string
    Email      string
    Created   *time.Time   `xorm:"created"`
}
type MeetingInfo struct {
    MID             int   `orm:"id,auto-inc"` //语义标签
    Title           string
    StartTime       time.Time
    EndTime         time.Time
    Participators   []string
    Created   *time.Time   `xorm:"created"`
}

func init() {
    err := myEngine.Sync2(new(UserInfo),new(MeetingInfo))
    fmt.Println("1231232123")
    checkErr(err)
}
// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
    if len(u.UserName) == 0 {
        panic("UserName shold not null!")
    }
    if u.Created == nil {
        t := time.Now()
        u.Created = &t
    }
    return &u
}
func NewMeetingInfo(m MeetingInfo) *MeetingInfo {
    if len(m.Title) == 0 {
        panic("Metting Title should not null!")
    }
    if m.Created == nil {
        t := time.Now()
        m.Created = &t
    }
    return &m
}