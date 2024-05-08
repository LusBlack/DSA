package pkg

import (
	"fmt"
	"log"
	"time"
)

var p = fmt.Println

func main() {

	id := getUserByName("John")
	p(id)

	chats := getUserChats(id)
	friends := getUserFriends(id)

	log.Println(chats)
	log.Println(friends)

}

func getUserFriends(id string) []string {

	return []string{
		"John",
		"Jane",
		"Joe",
		"Tiago",
		"James",
	}
}

func getUserChats(id string) []string {
	time.Sleep(time.Second * 2)

	return []string{
		"John",
		"Jane",
		"Joe",
	}

}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

// p("main execution started")
// p("No. of CPUs:", runtime.NumCPU())
// p("No. of Goroutines:", runtime.NumGoroutine())

// p("OS:", runtime.GOOS)
// p("Arch:", runtime.GOARCH)

// p("GOMAXPROCS", runtime.GOMAXPROCS(0))
