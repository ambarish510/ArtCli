package search

import (
	"github.com/urfave/cli"
	"fmt"
  "github.com/Flipkart/artcli/utils"
  "log"
  "strings"
	"encoding/json"
)
import "io/ioutil"
// import "os/user"
// import "net/url"
 import "os"
// import "bytes"
// import "time"
// import "net/http"
var githubURL string

func GetSearchFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
    	Name:  "store",
			Usage: "The store to which the artifact belongs to like Maven,gradle,python,npm,ruby etc",
		},cli.StringFlag{
			Name:  "packagename",
			Usage: "Name of the package in the format groupid:artifactid",
		},
	}
	return flags
}

func SearchArtifactory(c *cli.Context) error{
	ArtStore := strings.ToUpper(c.String("store"))
	ArtPackageName := c.String("packagename")
  log.Println("Store: ",ArtStore,"Package Name: ",ArtPackageName)

  returnStatus,returnMessage := validateSearchInputs(ArtStore,ArtPackageName)
  if returnStatus==1{
    log.Println("Input validation Failed : ",returnMessage)
    fmt.Println("Input validation Failed : ",returnMessage,"\nCheck the format below")
    cli.ShowCommandHelp(c, "search")
    os.Exit(1)
  }

  //make search api call
	getURL := utils.ArtEndPoint+"/artifactory/v1.0/artifacts/search?packageName="+ArtPackageName+"&store="+ArtStore+""
	log.Println("getURL: ",getURL)
  resp,err := utils.MakeGetCall(getURL)

	//fmt.Println("resp.StatusCode",resp.StatusCode)
  //fmt.Println("resp.Header",resp.Header)
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  bodyString := string(bodyBytes)
  fmt.Println(bodyString)

  status, message := utils.ValidateHTTPResponse(resp, err, "Search artifact")
  log.Println("status  ",status,"\n message ", message)

  //Parse the json response
  //SearchResponseParse()

	return nil
}
//this function validates inputs
//expects ArtStore in lowercase gradle/python/npm/ruby
func validateSearchInputs(ArtStore string,ArtPackageName string) (int,string) {
  listOfStores := []string{"MAVEN", "GRADLE","RUBY","PYTHON","NPM"}

  if ArtStore == "" {
    return 1,"Empty Store name"
  }else if ArtPackageName == ""{
    //include conditions to check the github URL format
    return 1,"Empty Package Name"
  }else if ! utils.Contains(listOfStores, ArtStore) {
    return 1,"Invalid store"
  }
  return 0, "Search Inputs validation Success"
}

type SearchResponse struct {
  ArtifactDetails struct {
    package_name string `json:"package_name"`
    version string `json:"version"`
  }
  Store string `json:"store"`
	Status string `json:"status"`
}


func SearchResponseParse(keysBody []byte) {
	//keysBody := []byte(`[{"artifact":{"package_name":"com.github.tomakehurst:wiremock-standalone","version":"2.6.0"},"store": "1","status": "-"}]`)
	keys := make([]SearchResponse,0)
	json.Unmarshal(keysBody, &keys)
	fmt.Println(keys)
  fmt.Println(keys[0].Store)
  fmt.Println(keys[0].Status)
  fmt.Println(keys[0].ArtifactDetails.package_name)
}
