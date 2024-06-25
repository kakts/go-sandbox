.PHONY: key
key:
	openssl genrsa 4096 > server.key

# csrファイルの生成
.PHONY: csr
csr:
	openssl req -new -key server.key -out server.csr

# crtファイルの生成
.PHONY: crt
crt:
	openssl x509 -req -days 10 -in server.csr -signkey server.key -out server.crt

# crtファイルの確認
.PHONY: check
check:
	openssl x509 -text -in server.crt

.PHONY: run
run:
	go run src/main.go