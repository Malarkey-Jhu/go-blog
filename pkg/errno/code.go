// Copyright 2023 Alan Jhu (朱倫君) <alan61109@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errno

var (
	OK = &Errno{HTTP: 200, Code: "OK", Message: "OK"}

	InternalServerError = &Errno{HTTP: 500, Code: "InternalServerError", Message: "Internal server error."}

	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.ErrPageNotFound", Message: "Page not found."}
)
