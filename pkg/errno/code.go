package errno

var (
	OK = &Errno{HTTP: 200, Code: "OK", Message: "OK"}

	InternalServerError = &Errno{HTTP: 500, Code: "InternalServerError", Message: "Internal server error."}

	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.ErrPageNotFound", Message: "Page not found."}
)
