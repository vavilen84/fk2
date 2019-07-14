package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddTestDb_20181208_131338 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddTestDb_20181208_131338{}
	m.Created = "20181208_131338"

	migration.Register("AddTestDb_20181208_131338", m)
}

// Run the migrations
func (m *AddTestDb_20181208_131338) Up() {
	m.SQL("CREATE DATABASE IF NOT EXISTS godb_test;")
}

// Reverse the migrations
func (m *AddTestDb_20181208_131338) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
