// Copyright 2023 Alan Jhu (朱倫君) <alan61109@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/Malarkey-Jhu/miniblog/internal/miniblog"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
