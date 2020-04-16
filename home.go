/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import "net/http"

const homeHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>轉帳 QR code 產生器</title>
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
<script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous" defer></script>
</head>
<body class="py-5 bg-light"><div class="container"><div class="row"><div class="col-12 my=3">
<div class="card">
  <div class="card-header"><h5>使用規範</h5></div>
  <div class="card-body">
  <ul>
    <li>此 API 僅供測試用，本人不提供任何形式及任何目的的保證。</li>
    <li>頻寬及 VPS 資源昂貴，請不要 DoS 我</li>
    <li>如需大量產生，請自行到 <a href="https://github.com/raohwork/twqrp-api">GitHub</a> 下載本服務的源碼自行修改使用</li>
  <ul>
  </div>
</div></div><div class="col-12 my-3">
<div class="card">
  <div class="card-header"><h5>手動產生</h5></div>
  <div class="card-body">
    <form action="/api/transfer/create" method="POST">
      <div class="form-group">
        <label for="bank">銀行轉帳代碼 (3 碼數字)</label>
        <input type="number" name="bank" id="bank"/>
      </div>
      <div class="form-group">
        <label for="account">帳號 (最多 16 碼數字)</label>
        <input type="number" name="account" id="account"/>
      </div>
      <div class="form-group">
        <label for="amount">金額 (最多 5 碼數字)</label>
        <input type="number" name="amount" id="amount"/>
      </div>
      <div class="form-group">
        <label for="note">附言 (最多 19 字)</label>
        <input type="number" name="note" id="note"/>
      </div>
      <div class="form-group">
        <label for="service">服務名稱</label>
        <input type="text" name="service" id="service" value="轉帳"/>
      </div>
      <button type="submit" class="btn btn-primary">產生 PNG</button>
      <button type="button" onclick="gen()" class="btn text-primary">產生 API URL</button>
    </form>
  </div>
</div>
</div><div class="col-12 my-3">
<div class="card">
  <div class="card-header"><h5>API URL</h5></div>
  <div class="card-body">
    <div class="card-text" id="urlresult">
    </div>
  </div>
</div>
</div></div></div>
<script>
function gen() {
var str = "http://twqrp.lao.gy/api/transfer/create?";
var arr = ["png=1"];
var keys = ["service","amount","note","bank","account"];
var tmp;

for (var i = 0; i < keys.length; i++) {
  tmp = $('#'+keys[i]).val();
  if (tmp !== '') {
    arr.push(keys[i]+"="+tmp);
  }
}

str += arr.join("&");
$('#urlresult').text(str);

}</script>
</body>
</html>`

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(homeHTML))
}
