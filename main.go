package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

var urlListFileName = flag.String("list", "./urls.yaml", "YAML file with list of URLS")

type Address struct {
	Description string
	URL         string
}

func main() {
	flag.Parse()

	urlList, err := readUrlList(urlListFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	numURLs := len(*urlList)
	errCount := performConnectivityChecks(urlList)
	if errCount == 0 {
		fmt.Printf("All connections to [%d] urls check out. All good\n", numURLs)
	} else {
		fmt.Printf("there were errors connecting to %d/%d total urls\n", errCount, numURLs)
	}

}

// try connecting to URLs and return number of errors
func performConnectivityChecks(allAddresses *[]Address) int {
	errCount := 0
	timeout := time.Duration(1 * time.Second)
	for i, oneAddress := range *allAddresses {
		fmt.Printf("Checking connectivity to [%s]\n", oneAddress.Description)
		client := http.Client{
			Timeout: timeout,
		}
		_, err := client.Get(oneAddress.URL)
		if err != nil {
			fmt.Printf("unable to connect to item [%d]: [%s] with error [%v]", i, oneAddress.Description, err)
			errCount++
		}
	}

	return errCount
}

func readUrlList(urlListFileName *string) (*[]Address, error) {
	yamlFile, err := ioutil.ReadFile(*urlListFileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to load yaml file [%s]: #%v\n", *urlListFileName, err)
	}

	var urls []Address
	err = yaml.Unmarshal(yamlFile, &urls)
	if err != nil {
		log.Fatalf("Error unmarshalling the list of urls: %v", err)
	}

	fmt.Printf("loaded %d urls to check\n", len(urls))
	return &urls, nil
}
