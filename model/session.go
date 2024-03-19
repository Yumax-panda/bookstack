package model

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/gofrs/uuid"
)

type SessionRecord struct {
	Token       string    `gorm:"type:varchar(50);primaryKey"`
	ReferenceID uuid.UUID `gorm:"type:char(36);unique"`
	UserID      uuid.UUID `gorm:"type:varchar(36);index"`
	Data        []byte    `gorm:"type:longblob"`
	Created     time.Time `gorm:"precision:6"`
}

func (*SessionRecord) TableName() string {
	return "r_sessions"
}

func (sr *SessionRecord) SetData(data map[string]interface{}) {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(data); err != nil {
		panic(err)
	}
	sr.Data = b.Bytes()
}

func (sr *SessionRecord) GetData() (data map[string]interface{}, err error) {
	return data, gob.NewDecoder(bytes.NewReader(sr.Data)).Decode(&data)
}
