package mailmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Senders slice.
type Senders []Sender

// Senders data.
type Sender struct {
	ID            string    `rql:"filter,sort,column=sen_id"`
	Name          string    `rql:"filter,sort,column=sen_name"`
	Description   string    `rql:"filter,sort,column=sen_description"`
	Host          string    `rql:"filter,sort,column=sen_host"`
	Port          int       `rql:"filter,sort,column=sen_port"`
	User          string    `rql:"filter,sort,column=sen_user"`
	Password      string    `rql:"filter,sort,column=sen_password"`
	From          string    `rql:"filter,sort,column=sen_from"`
	CreatedAt     time.Time `rql:"filter,sort,column=sen_created_at"`
	UpdatedAt     time.Time `rql:"filter,sort,column=sen_updated_at"`
	Status        base.Status
	RecordActions base.RecordActions
}
