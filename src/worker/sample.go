package main

import (
	"encoding/json"
	"fmt"
	"github.com/benmanns/goworker"
	_ "time"
)

func addWorker(queue string, args ...interface{}) error {

	fmt.Println("Into addworker")
	var sum int64 = 0
	for _, arg := range args {
		fmt.Println("Into args")
		//sum, _ += strconv.Atoi(arg.(string))
		num, ok := arg.(json.Number)
		if !ok {
			fmt.Printf("Invalid param ! %v\n", arg)
			return fmt.Errorf("Invalid param %v", num)
		}

		numAsInt, err := num.Int64()
		if err != nil {
			fmt.Printf("Invalid param while converting to INt64 ! %v\n", num)
			return fmt.Errorf("Invalid param %v", num)
		}

		sum += numAsInt
	}

	fmt.Printf("From %s, %v | Answer : %v\n", queue, args, sum)

	return nil
}

/*
	To build : go build worker
	To run : ./worker -queues="sampleadd" -namespace="goworker"  -use-number
*/

func main() {
	goworker.Register("SampleAddJobClass", addWorker)

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}

}
