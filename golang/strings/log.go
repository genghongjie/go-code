package main

import (
	"log"
)

func main() {
	//var buf bytes.Buffer
	//logger := log.New(&buf, "", log.Lshortfile)
	//LstdFlags
	//logger.Print("this is logger")
	//logger.Print("122223323")
	//fmt.Print(&buf)
	//log.SetFlags(log.Lshortfile)
	//fmt.Println(log.Flags())

	log.Println("this is log 1")
	log.Fatalln("this is log 2")
	log.Printf("%s\n", "this is log 3")
	log.Panicln("this is log 4")

}
