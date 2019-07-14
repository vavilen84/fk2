package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddNewUserColumns_20190714_232109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddNewUserColumns_20190714_232109{}
	m.Created = "20190714_232109"

	migration.Register("AddNewUserColumns_20190714_232109", m)
}

// Run the migrations
func (m *AddNewUserColumns_20190714_232109) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN role integer NOT NULL")
}

// Reverse the migrations
func (m *AddNewUserColumns_20190714_232109) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
