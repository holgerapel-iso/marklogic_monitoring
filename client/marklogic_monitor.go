package main

import (
	"encoding/json"
	"fmt"
	"github.com/pteich/http-digest-auth-client"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ClusterHealthReport struct {
	State        string
	ResourceType string `json:"resource-type"`
	ResourceId   string `json:"resource-id"`
	ResourceName string `json:"resource-name"`
	Code         string
	Message      string
}

type ClusterHealthReports struct {
	ClusterId string                `json:"cluster-id"`
	Reports   []ClusterHealthReport `json:"cluster-health-report"`
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == "*" || a == e {
			return true
		}
	}
	return false
}

func selectReports(reports []ClusterHealthReport, resourceTypes []string, ignoreStates []string) []ClusterHealthReport {
	result := make([]ClusterHealthReport, 0)
	for _, report := range reports {
		if contains(resourceTypes, report.ResourceType) && !contains(ignoreStates, report.State) {
			result = append(result, report)
		}
	}
	return result
}

func main() {
	var url string
	username := ""
	password := ""
	resourceTypes := []string{"*"}
	ignoreStates := []string{"info"}

	if len(os.Args) < 4 {
		log.Fatal("Usage: marklogic_monitor url username password [resourceTypes] [ignoreStates]")
	}

	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	if len(os.Args) > 2 {
		username = os.Args[2]
	}
	if len(os.Args) > 3 {
		password = os.Args[3]
	}
	if len(os.Args) > 4 {
		resourceTypes = strings.Split(os.Args[4], ",")
	}
	if len(os.Args) > 5 {
		ignoreStates = strings.Split(os.Args[5], ",")
	}

	monitorClient := http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(username) > 0 {
		var DigestAuth *httpDigestAuth.DigestHeaders
		DigestAuth = &httpDigestAuth.DigestHeaders{}
		DigestAuth, err = DigestAuth.Auth(username, password, url)
		DigestAuth.ApplyAuth(req)
	}

	res, err := monitorClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	var content ClusterHealthReports
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}
	result := selectReports(content.Reports, resourceTypes, ignoreStates)
	if len(result) == 0 {
		fmt.Println("OK")
	} else {
		for _, item := range result {
			fmt.Println(item)
		}
	}
}
