package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20181203_144211 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20181203_144211{}
	m.Created = "20181203_144211"

	migration.Register("Post_20181203_144211", m)
}

// Run the migrations
func (m *Post_20181203_144211) Up() {
	m.SQL("CREATE TABLE post (id int NOT NULL PRIMARY KEY AUTO_INCREMENT, title varchar(255), content text);")
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Post_20181203_144211) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
