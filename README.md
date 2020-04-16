簡單的 PoC api server，驗證 TWQRP 產生程式 [PoC server](http://twqrp.lao.gy/)

# By Docker

先修改 docker-compose.yml 裡 port 的設定，然後執行

```sh
docker-compose up
```

# 自己編譯

先安裝 Go 的編譯環境，然後在源碼目錄中執行

```sh
go build
```

**此服務寫死在 9801 port**

# License

MPL 2.0, see LICENSE.txt
