package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserTableInitial_20181205_143944 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserTableInitial_20181205_143944{}
	m.Created = "20181205_143944"

	migration.Register("UserTableInitial_20181205_143944", m)
}

// Run the migrations
func (m *UserTableInitial_20181205_143944) Up() {
	m.SQL("CREATE TABLE user (id int NOT NULL PRIMARY KEY AUTO_INCREMENT, email varchar(255), password text, salt text);")
	m.SQL("ALTER TABLE user ADD UNIQUE INDEX email_idx(email);")
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *UserTableInitial_20181205_143944) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
