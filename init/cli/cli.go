// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/palantir/pkg/cli"
)

func handleError(ctx cli.Context, err error) int {
	switch theError := err.(type) {
	case nil:
		return 0 // No error
	case *ErrorResponse:
		ctx.Errorf(theError.Error())
		return theError.exitCode
	case *SuccessResponse:
		if theError.msg != "" {
			ctx.Printf(theError.msg)
		}
		return theError.exitCode
	default:
		return 1 // Some other, unknown error
	}
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "go-init"
	app.Usage = "A simple init.sh - style service launcher CLI"
	app.ErrorHandler = handleError

	app.Subcommands = []cli.Command{statusCommand(), startCommand()}
	return app
}
