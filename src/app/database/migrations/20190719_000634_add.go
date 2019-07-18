package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Add_20190719_000634 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Add_20190719_000634{}
	m.Created = "20190719_000634"

	migration.Register("Add_20190719_000634", m)
}

// Run the migrations
func (m *Add_20190719_000634) Up() {
	m.SQL("CREATE TABLE image_to_user (" +
		"image_uuid varchar(255), " +
		"user_id int(10) " +
		");")
	m.SQL("ALTER TABLE image_to_user ADD UNIQUE INDEX image_uuid_user_id_idx(image_uuid,user_id);")
}

// Reverse the migrations
func (m *Add_20190719_000634) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
