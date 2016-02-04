package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/QuentinPerez/go-stuff/rpc/rpcdefinition"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %v", err)
	}
	var result rpcdefinition.Result
	//call remote procedure with args
	err = client.Call("Call.FastCall", rpcdefinition.Args{time.Now()}, &result)
	end := time.Now()
	if err != nil {
		log.Fatalf("error in Call %v", err)
	}
	fmt.Println("FastCall:")
	fmt.Println("Start   at:", result.Start)
	fmt.Println("Receive at:", result.Middle, "diff S/E", result.Middle.Sub(result.Start))
	fmt.Println("End     at:", end, "diff M/E", end.Sub(result.Middle))
	fmt.Println("Total", end.Sub(result.Start), result.NB)

	err = client.Call("Call.SlowCall", rpcdefinition.Args{time.Now()}, &result)
	end = time.Now()
	if err != nil {
		log.Fatalf("error in Call %v", err)
	}
	fmt.Println()
	fmt.Println("SlowCall:")
	fmt.Println("Start   at:", result.Start)
	fmt.Println("Receive at:", result.Middle, "diff S/E", result.Middle.Sub(result.Start))
	fmt.Println("End     at:", end, "diff M/E", end.Sub(result.Middle))
	fmt.Println("Total", end.Sub(result.Start), result.NB)

	asyncCall := client.Go("Call.SlowCall", rpcdefinition.Args{time.Now()}, &result, nil)
	fmt.Println()
	fmt.Println("SlowCall Async:")
	id := 0
	for {
		select {
		case ret := <-asyncCall.Done:
			end = time.Now()
			if ret.Error != nil {
				log.Fatalf("error in Call %v", ret.Error)
			}
			fmt.Println()
			fmt.Println("Start   at:", result.Start)
			fmt.Println("Receive at:", result.Middle, "diff S/E", result.Middle.Sub(result.Start))
			fmt.Println("End     at:", end, "diff M/E", end.Sub(result.Middle))
			fmt.Println("Total", end.Sub(result.Start), result.NB)
			return
		default:
			fmt.Printf("\rI'm waiting.... %v", id)
			id++
		}
	}
}
