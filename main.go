package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jirayuth289/go-test/basic"
	"github.com/jirayuth289/go-test/jirayuth289"
	"github.com/jirayuth289/go-test/oop"
	basicClass "github.com/jirayuth289/go-test/oop/basic-class"
	solid "github.com/jirayuth289/go-test/oop/solid"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s\n", id)
	jirayuth289.SayHelloJirayuth289()
	jirayuth289.SayTest()

	basicClass.TestCreateBasicClass()
	solid.TestSolid()
	oop.LearnOop()
	basic.Basic()
}

// go run main.go
// go build main.go
// nodemon --exec go run main.go --signal SIGTERM

// จัดการ modules มี 2 files
// 1. go.mod list dependency ทั้งหมด
// 2. go.sum คือ ไฟล์ที่เก็บเวอร์ชันที่ลงเอาไว้

// go mod init {repo_name}
// go list -m all
