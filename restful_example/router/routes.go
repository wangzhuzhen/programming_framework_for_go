package router

import (
	"net/http"
	"github.com/wangzhuzhen/programming_framework_for_go/restful_example/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handler.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handler.TodoShow,
	},
}