package main

import(
	"fmt"
	"examguard/internal/logging"
	"examguard/internal/mointor/usb"
	
)

func main() {

	go usb.MonitorUSB()
	verifyErr := logging.VerifyChain()
	fmt.Println("Verification result:", verifyErr)
	select{}

	

}
