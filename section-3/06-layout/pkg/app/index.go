package app

import (
	"net/http"

	"github.com/acoshift/golang-for-beginner/section-3/06-layout/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
