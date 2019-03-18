package main

import (
	"Ginga-Client-Go/ginga"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i ++ {
		go do(i)
	}

	var forever chan struct{}
	<- forever
}


func do(i int) {
	c := ginga.Client{
		Token: "test_token_1",
		Endpoint: "0.0.0.0:1903",
		Nonce: "SDFSDFSS",
	}

    for {
		err := c.Lock()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(time.Now(), i)

		time.Sleep(time.Second)
		err = c.Unlock()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}