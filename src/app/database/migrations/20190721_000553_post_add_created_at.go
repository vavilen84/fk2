package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PostAddCreatedAt_20190721_000553 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PostAddCreatedAt_20190721_000553{}
	m.Created = "20190721_000553"

	migration.Register("PostAddCreatedAt_20190721_000553", m)
}

// Run the migrations
func (m *PostAddCreatedAt_20190721_000553) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN created_at int")
}

// Reverse the migrations
func (m *PostAddCreatedAt_20190721_000553) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
