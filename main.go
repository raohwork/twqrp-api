/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import "net/http"

func main() {
	http.HandleFunc("/api/transfer/create", createTransfer)
	http.HandleFunc("/", home)
	http.ListenAndServe(":9801", nil)
}

func createTransfer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data, err := genContent(r.Form)
	if err != "" {
		errMsg(w, err)
		return
	}

	if r.Form.Get("png") != "" {
		writeQRPNG(w, data)
	} else {
		writeQR(w, data)
	}
}
