package handler

import (
	"net/http"

	//Local Modules
	"github.com/devdhanadiya/rss-aggregator/helper"
)

func ResHandler(w http.ResponseWriter, r *http.Request) {
	helper.ResJson(w, 200, struct{}{})
}
