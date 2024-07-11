.PHONY run:
run:
	go run src/main.go

.PHONY: get_comments
get_comments:
	curl http://localhost:8080/comments

.PHONY post_comment:
post_comment:
	curl -X POST http://localhost:8080/comment \
		-H "Content-Type: application/json" \
		-d '{"Username":"ユーザー名", "message":"ここにコメント"}'

.PHONY run_k6:
run_k6:
	k6 run k6/script.js
