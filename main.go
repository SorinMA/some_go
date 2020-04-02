package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(kill_8000()) // will kill all procs on port 8000
	time.Sleep(time.Second)  // wait 1 sec till the kill is done
	setup_server()           // lets server
}
