.PHONY run:
run:
	go run src/main.go

.PHONY: get_comments
get_comments:
	curl http://localhost:8080/comments

.PHONY post_comment:
post_comment:
	curl -X POST http://localhost:8080/comments \
		-H "Content-Type: application/json" \
		-d '{"Username":"ユーザー名", "message":"ここにコメント"}'