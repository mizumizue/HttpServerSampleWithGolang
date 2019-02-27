package router

import (
	"github.com/gorilla/mux"
	"go-server-samples/datastore/with_dao_service_pattern/controllers"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"AddTask",
		strings.ToUpper("Post"),
		"/api/v2/tasklists/{listId}/tasks",
		controllers.AddTask,
	},

	Route{
		"DeleteTask",
		strings.ToUpper("Delete"),
		"/api/v2/tasklists/{listId}/tasks/{taskId}",
		controllers.DeleteTask,
	},

	Route{
		"GetTask",
		strings.ToUpper("Get"),
		"/api/v2/tasklists/{listId}/tasks/{taskId}",
		controllers.GetTask,
	},

	Route{
		"GetTasks",
		strings.ToUpper("Get"),
		"/api/v2/tasklists/{listId}/tasks",
		controllers.GetTasks,
	},

	Route{
		"UpdateTask",
		strings.ToUpper("Put"),
		"/api/v2/tasklists/{listId}/tasks",
		controllers.UpdateTask,
	},

	Route{
		"AddTaskList",
		strings.ToUpper("Post"),
		"/api/v2/tasklists",
		controllers.AddTaskList,
	},

	Route{
		"DeleteTaskList",
		strings.ToUpper("Delete"),
		"/api/v2/tasklists/{listId}",
		controllers.DeleteTaskList,
	},

	Route{
		"GetTaskList",
		strings.ToUpper("Get"),
		"/api/v2/tasklists/{listId}",
		controllers.GetTaskList,
	},

	Route{
		"GetTaskLists",
		strings.ToUpper("Get"),
		"/api/v2/tasklists",
		controllers.GetTaskLists,
	},

	Route{
		"UpdateTaskList",
		strings.ToUpper("Put"),
		"/api/v2/tasklists",
		controllers.UpdateTaskList,
	},
}
