# 自己署名証明書の作成

## keyの作成
```
make key
```
server.keyが作成される

## csrの作成
```
make csr
```
CSRファイル(server.csr)が作成される

## 自己署名証明書の作成
```
make crt
```
自己署名証明書(server.crt)が作成される


TODO chromeで問題なく使えるようにする
