package handler

import (
	"net/http"

	"github.com/devdhanadiya/rss-aggregator/helper"
)

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	helper.ResError(w, 400, "Something went wrong")
}
