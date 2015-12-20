package database

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Operation int

const (
	Added Operation = iota
	Exited
	MarkedAsCurrent
	MarkedAsFinished
)

type Queue struct {
	ID              int
	Title           string `sql:"type:varchar(255)"`
	Description     string `sql:"type:varchar(255)"`
	MaxPeoples      int
	Creator         User   // One2One
	CreatorID       int    `json:"-"`
	CurrentMember   Member // One2One
	CurrentMemberID int    `json:"-"`
	Members         []Member
	Created         time.Time
	IsActive        bool
	IsDeleted       bool
}

type Member struct {
	ID               int
	SubscriptionTime time.Time
	QueueID          int  `sql:"unique_index:queue_user"` // Foreign Key
	User             User // One2One
	UserID           int  `sql:"unique_index:queue_user"`
}

type User struct {
	ID           int
	Email        string `sql:"type:varchar(255);unique_index"`
	FirstName    string `sql:"type:varchar(65)"`
	LastName     string `sql:"type:varchar(65)"`
	Password     string `sql:"type:varchar(255)" json:"-"`
	IsSuperAdmin bool
}

type History struct {
	ID        int
	Time      time.Time
	Queue     Queue
	User      User
	Operation Operation
}
