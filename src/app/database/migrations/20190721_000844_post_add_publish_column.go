package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PostAddPublishColumn_20190721_000844 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PostAddPublishColumn_20190721_000844{}
	m.Created = "20190721_000844"

	migration.Register("PostAddPublishColumn_20190721_000844", m)
}

// Run the migrations
func (m *PostAddPublishColumn_20190721_000844) Up() {
	m.SQL("ALTER TABLE user ADD COLUMN publish int")
}

// Reverse the migrations
func (m *PostAddPublishColumn_20190721_000844) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
