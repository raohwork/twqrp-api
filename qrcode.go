/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/base64"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func writeQR(w http.ResponseWriter, data []byte) {
	png, err := qrcode.Encode(string(data), qrcode.Medium, 256)
	if err != nil {
		errMsg(w, EGenQR)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	enc := base64.NewEncoder(base64.StdEncoding, w)
	w.Write([]byte(`{"data":"`))
	enc.Write(png)
	w.Write([]byte(`"`))
}

func writeQRPNG(w http.ResponseWriter, data []byte) {
	png, err := qrcode.Encode(string(data), qrcode.Medium, 256)
	if err != nil {
		errMsg(w, EGenQR)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
