package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runbashcache/myredis"
	"time"
)

func main() {
	fmt.Println("Hello world!")
	var filename = flag.String("f", "invalid", "filename to analyse")
	flag.Parse()
	fmt.Println("file name - ", *filename)
	bash_command0 := "tshark"
	bash_command1 := "-nr"
	bash_command2 := "-qz"
	bash_command3 := "rtp,streams"
	cmd := exec.Command(bash_command0, bash_command1, *filename, bash_command2, bash_command3)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Some error in tshark command !!")
		//log.Fatal(err)
	} else {
		fmt.Println("Successfully ran tshark command !!")
	}

	//fmt.Println("Command output - ", string(out))
	value := string(out)
	hash := "hash"
	key := "bley"

	myredis.MyHset(hash, key, value)
	myredis.MySetExp(hash)
	mySleep(2)
	val := myredis.MyHget(hash, key)
	fmt.Println(val)

	key = "hohoh"
	value = "nononno"
	myredis.MyHset(hash, key, value)
	myredis.MySetExp(hash)
	mySleep(2)
	val = myredis.MyHget(hash, key)
	fmt.Println(val)
}

func mySleep(sec int) {
	fmt.Println("Sleeping ", sec, " secs....")
	time.Sleep(time.Duration(sec*1000) * time.Millisecond)
}
