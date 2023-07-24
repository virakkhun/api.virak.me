package shared_services

import "testing"

func TestApiPrefixWithBlogPath(t *testing.T) {
	path := "blog"
	toBePath := "/api/v1/blog"
	pathWithPrefix := ApiPrefix(path)

	if pathWithPrefix != toBePath {
		t.Failed()
	}
}
