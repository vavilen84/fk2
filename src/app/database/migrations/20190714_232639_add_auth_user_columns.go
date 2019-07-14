package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddAuthUserColumns_20190714_232639 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddAuthUserColumns_20190714_232639{}
	m.Created = "20190714_232639"

	migration.Register("AddAuthUserColumns_20190714_232639", m)
}

// Run the migrations
func (m *AddAuthUserColumns_20190714_232639) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN first_name varchar(255) NOT NULL")
	m.SQL("ALTER TABLE user ADD COLUMN last_name varchar(255) NOT NULL")
}

// Reverse the migrations
func (m *AddAuthUserColumns_20190714_232639) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
