package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var p = fmt.Println

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()

	id := getUserByName("John")
	p(id)

	ch := make(chan *Message, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))

}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {

	time.Sleep(time.Second * 1)

	ch <- &Message{
		friends: []string{
			"John",
			"Jane",
			"Joe",
			"Tiago",
			"James",
		},
	}
	wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	ch <- &Message{
		chats: []string{
			"John",
			"Jane",
			"Joe",
		},
	}
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

// what not to do
func leaky() {
	ch := make(chan int)

	go func() {
		msg := <-ch
		log.Println(msg)
	}()
}
