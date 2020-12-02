package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/qihach/go-week-02/storage"
)

// GetBlogs ...
func GetBlogs(w http.ResponseWriter, req *http.Request) {

	keys, ok := req.URL.Query()["user_id"]
	if !ok || len(keys) == 0 {
		http.Error(w, "user_id not provided", http.StatusBadRequest)
		return
	}
	userID, err := strconv.ParseInt(keys[0], 10, 64)
	if err != nil {
		http.Error(w, "user_id is not an integer", http.StatusBadRequest)
		return
	}
	// mock storage object
	storage := storage.NewBlogStorage(&storage.BlogDAO{})
	blogs, err := storage.GetBlogsForUserID(userID)
	if err != nil {
		log.Printf("%+v\n", errors.Wrapf(err, "failed to get blogs for user id %d", userID))
		http.Error(w, "blogs not found", http.StatusNotFound)
		return
	}
	var res string
	for _, blog := range blogs {
		res += blog.Text
	}
	fmt.Fprintf(w, res)
}
