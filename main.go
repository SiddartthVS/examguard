package main

import(
	"fmt"
)

func main() {

	go MonitorUSB()
	verifyErr := VerifyChain()
	fmt.Println("Verification result:", verifyErr)
	select{}

	

}
