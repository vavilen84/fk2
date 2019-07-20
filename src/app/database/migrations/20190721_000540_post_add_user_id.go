package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PostAddUserId_20190721_000540 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PostAddUserId_20190721_000540{}
	m.Created = "20190721_000540"

	migration.Register("PostAddUserId_20190721_000540", m)
}

// Run the migrations
func (m *PostAddUserId_20190721_000540) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN user_id int")
}

// Reverse the migrations
func (m *PostAddUserId_20190721_000540) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
