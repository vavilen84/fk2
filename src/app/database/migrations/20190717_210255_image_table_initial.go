package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ImageTableInitial_20190717_210255 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ImageTableInitial_20190717_210255{}
	m.Created = "20190717_210255"

	migration.Register("ImageTableInitial_20190717_210255", m)
}

// Run the migrations
func (m *ImageTableInitial_20190717_210255) Up() {
	m.SQL("CREATE TABLE image (" +
		"uuid varchar(255) NOT NULL PRIMARY KEY, " +
		"original_filename varchar(255), " +
		"ext varchar(255), " +
		"filepath text, " +
		"status int(10)" +
		");")
}

// Reverse the migrations
func (m *ImageTableInitial_20190717_210255) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
