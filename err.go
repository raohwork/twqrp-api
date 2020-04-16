/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"io"
)

const (
	ENoQuota  = "E1000"
	EGenTWQRP = "E1002"
	EGenQR    = "E1003"
)

func errMsg(w io.Writer, code string) {
	fmt.Fprintf(w, `{"errors":[{"code":"%s"}]}`, code)
}
