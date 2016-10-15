package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type thingTime struct {
	atLeast int
	atMost  int
}

func doThing(people []string, op string, t thingTime) {
	var getReadyWg sync.WaitGroup
	r := rand.New(rand.NewSource(99))
	for _, name := range people {
		fmt.Printf("%s started %s\n", name, op)
		getReadyWg.Add(1)
		var randWait = r.Intn(t.atMost-t.atLeast) + t.atLeast
		var sleepTime = time.Duration(randWait) * time.Millisecond
		go func(person string) {
			time.Sleep(sleepTime)
			fmt.Printf("%s spent %d seconds %s\n", person, randWait, op)
			defer getReadyWg.Done()
		}(name)
	}
	getReadyWg.Wait()
}

func armAlarm(alarm chan bool) {
	var countDown = 60 * time.Millisecond
	fmt.Println("Arming alarm.")
	go func() {
		fmt.Println("Alarm is counting down.")
		time.Sleep(countDown)
		fmt.Println("Alarm is armed.")
		alarm <- true
	}()
}

func main() {
	fmt.Println("Let's go for a walk!")

	doThing([]string{"Alice", "Bob"}, "getting ready", thingTime{60, 90})

	alarm := make(chan bool, 1)
	armAlarm(alarm)

	doThing([]string{"Alice", "Bob"}, "putting on shoes", thingTime{35, 45})

	fmt.Println("Exiting and locking the door")

	_ = <-alarm
}
