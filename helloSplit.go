package main

import (
	"fmt"
	"os"

	"github.com/splitio/go-client/splitio/client"
	"github.com/splitio/go-client/splitio/conf"
	"github.com/splitio/go-toolkit/logging"
)

func main() {
	//3.Set a ready config time to make sure the SDK is properly loaded and ready before asking it for a treatment.
	cfg := conf.Default()
	cfg.LoggerConfig.LogLevel = logging.LevelError

	//4.Instantiate a Split Factory to pull down the rollout plans from Split.
	factory, err := client.NewSplitFactory(os.Getenv("SDK_KEY"), cfg)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
		return
	}

	//5.If no error occurs, instantiate a client to start calling treatments calls.
	client := factory.Client()
	err = client.BlockUntilReady(25)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
	}

	// which treatment your customer should see
	treatment := client.Treatment("CUSTOMER_ID", "blue-green", nil)

	if treatment == "on" {
		// insert code here to show on treatment
		fmt.Println("Treatment is on, let's do something")
	} else if treatment == "off" {
		// insert code here to show off treatment
		fmt.Println("Treatment is off")
	} else {
		// insert your control treatment code here
		fmt.Println("Treatment is undefined")
	}

	client.Destroy()
}
