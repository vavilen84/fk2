package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CollectionTableInitial_20190717_213057 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CollectionTableInitial_20190717_213057{}
	m.Created = "20190717_213057"

	migration.Register("CollectionTableInitial_20190717_213057", m)
}

// Run the migrations
func (m *CollectionTableInitial_20190717_213057) Up() {
	m.SQL("CREATE TABLE collection (" +
		"id int NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"name varchar(255), " +
		"status int(10)" +
		");")
	m.SQL("CREATE TABLE image_to_collection (" +
		"image_uuid varchar(255), " +
		"collection_id int(10) " +
		");")
	m.SQL("ALTER TABLE image_to_collection ADD UNIQUE INDEX image_uuid_collection_id_idx(image_uuid,collection_id);")
}

// Reverse the migrations
func (m *CollectionTableInitial_20190717_213057) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
