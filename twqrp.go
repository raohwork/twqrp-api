/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"net/url"
	"strconv"

	"github.com/raohwork/twqrp-go"
)

func genContent(data url.Values) (ret []byte, err string) {
	g, e := twqrp.NewTransfer(data.Get("bank"), data.Get("account"))
	if e != nil {
		err = EGenTWQRP
		return
	}

	g.Name = data.Get("service")
	if g.Name == "" {
		g.Name = "轉帳"
	}

	if astr := data.Get("amount"); astr != "" {
		amount, e := strconv.ParseInt(astr, 10, 32)
		if e != nil {
			err = EGenTWQRP
			return
		}

		if e = g.TryAmount(int(amount)); e != nil {
			err = EGenTWQRP
			return
		}
	}

	if note := data.Get("note"); note != "" {
		if e = g.TryNote(note); e != nil {
			err = EGenTWQRP
			return

		}
	}

	ret = []byte(g.String())
	return
}
