package test

import (
	"app/models/post"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOneById(t *testing.T) {
	or := beginTransaction()

	Convey("Subject: Test OneById func\n", t, func() {
		Convey("Post is found", func() {
			post, _ := post.OneById(1, or)
			So(post.Id, ShouldEqual, 1)
		})
		Convey("Post is not found", func() {
			post, _ := post.OneById(111111, or)
			So(post.Id, ShouldEqual, 0)
		})
		Convey("Create new Post", func() {
			_, err := post.Create("title", "content", or)
			So(err, ShouldBeNil)
		})
		Convey("Find created Post", func() {
			post, _ := post.OneById(1, or)
			So(post, ShouldNotBeNil)
		})
	})

	rollbackTransaction(or)
}

func TestDel(t *testing.T) {
	or := beginTransaction()

	Convey("Subject: Test Del func\n", t, func() {
		Convey("Find created Post", func() {
			post, _ := post.OneById(1, or)
			So(post, ShouldNotBeNil)
		})
		Convey("Delete post", func() {
			m, _ := post.OneById(1, or)
			post.Del(m.Id, or)
		})
		Convey("Post is not found", func() {
			post, _ := post.OneById(1, or)
			So(post.Id, ShouldEqual, 0)
		})
	})

	rollbackTransaction(or)
}

func TestUpdate(t *testing.T) {
	or := beginTransaction()

	Convey("Subject: Test Update func\n", t, func() {
		updateTitle := "updated-title"
		updatedContent := "updated-content"
		Convey("Update new Post", func() {
			post.Update(1, updateTitle, updatedContent, or)
		})
		Convey("See updated fields", func() {
			post, _ := post.OneById(1, or)
			So(post.Title, ShouldEqual, updateTitle)
			So(post.Content, ShouldEqual, updatedContent)
		})
	})

	rollbackTransaction(or)
}
