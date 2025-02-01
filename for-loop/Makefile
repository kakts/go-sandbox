.RECIPEPREFIX =  


.PHONY ls:
ls:
    ls -l
    ls -Fal

.PHONY check:
check:
	go build -gcflags=all=-d=loopvar=2 main.go

.PHONY run:
run:
	go run main.go
