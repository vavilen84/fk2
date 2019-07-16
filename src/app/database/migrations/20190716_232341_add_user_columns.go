package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddUserColumns_20190716_232341 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUserColumns_20190716_232341{}
	m.Created = "20190716_232341"

	migration.Register("AddUserColumns_20190716_232341", m)
}

// Run the migrations
func (m *AddUserColumns_20190716_232341) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN type int(10)")
	m.SQL("ALTER TABLE user ADD COLUMN about text")
	m.SQL("ALTER TABLE user ADD COLUMN pinterest_link varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN instagram_link varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN facebook_link varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN pnone varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN skype varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN telegram varchar(255)")
	m.SQL("ALTER TABLE user ADD COLUMN avatar text")
}

// Reverse the migrations
func (m *AddUserColumns_20190716_232341) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
