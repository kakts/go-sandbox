.PHONY run:
run:
	go run src/main.go

.PHONY check_pprof:
check_pprof:
	go tool pprof -raw fd.pprof

.PHONY check_pprof_web:
check_pprof_web:
	go tool pprof -http=:8080 fd.pprof
