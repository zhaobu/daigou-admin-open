package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var (
	// ErrNotFound error when record not found
	ErrNotFound = fmt.Errorf("record Not Found")

	// ErrUnableToMarshalJSON error when json payload corrupt
	ErrUnableToMarshalJSON = fmt.Errorf("json payload corrupt")

	// ErrUpdateFailed error when update fails
	ErrUpdateFailed = fmt.Errorf("db update error")

	// ErrInsertFailed error when insert fails
	ErrInsertFailed = fmt.Errorf("db insert error")

	// ErrDeleteFailed error when delete fails
	ErrDeleteFailed = fmt.Errorf("db delete error")

	// ErrBadParams error when bad params passed in
	ErrBadParams = fmt.Errorf("bad params error")
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`                 //创建时间
	UpdatedAt time.Time  `json:"updatedAt"`                 //更新时间
	DeletedAt *time.Time `json:"deletedAt"`                 //删除时间
	CreateBy  string     `json:"createBy" gorm:"size:128;"` //创建人
	UpdateBy  string     `json:"updateBy" gorm:"size:128;"` //更新者
}

type Time struct {
	time.Time
}

func Now() *Time {
	return &Time{time.Now()}
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	t.Time = now
	return
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type JSTime struct {
	time.Time
}

const (
	TimeApiFormart  = "2006-01-02 15:04:05"
	TimeDataFormat  = "2006-01-02"
	TimeJsISOFormat = "2006-01-02T15:04:05.999Z07:00"
)

func (t *JSTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeJsISOFormat+`"`, string(data), time.Local)
	t.Time = now
	return
}

func (t *JSTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeJsISOFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, TimeJsISOFormat)
	b = append(b, '"')
	return b, nil
}

func (t *JSTime) String() string {
	return t.Format(TimeJsISOFormat)
}
