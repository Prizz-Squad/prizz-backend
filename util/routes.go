package util

// Routes these routes will be skipped by the middleware
var Routes = []string{
	"/prizz/api/v1/login",
	"/prizz/api/v1/register",
}

// AdminRoutes MemberRoutes CustomerRoutes are the routes for each role that they aren't allowed to access
var AdminRoutes = []string{}

var ManagerRoutes = []string{
	"/prizz/api/v1/project",
}

var MemberRoutes = []string{
	"/prizz/api/v1/users",
	"/prizz/api/v1/project",
	"/prizz/api/v1/taskHistory",
	"/prizz/api/v1/taskHistory/create",
	"/prizz/api/v1/taskHistoryByDateAndUser",
	"/prizz/api/v1/totalHours",
}

var CustomerRoutes = []string{
	"/prizz/api/v1/users",
	"/prizz/api/v1/project",
	"/prizz/api/v1/taskHistory/create",
}
