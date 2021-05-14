package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dongsubkim/go-web-programming/Chapter_8_Testing_Web_Applications/8_24_test_ginkgo"
)

var _ = Describe("824TestGinkgo", func() {
	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		mux.HandleFunc("/post/", HandleRequest(post))
		writer = httptest.NewRecorder()
	})

	Context("Get a post using an id", func() {
		It("Should get a post", func() {
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.Id).To(Equal(1))
		})
	})

	Context("Get an eror if post id is not an integer", func() {
		It("Should get a HTTP 500 response", func() {
			request, _ := http.NewRequest("GET", "/post/hello", nil)
			mux.ServeHTTP(writer, request)
			Expect(writer.Code).To(Equal(500))
		})
	})
})