package create

import (
	"github.com/urfave/cli"
	"fmt"
  "github.com/Flipkart/artcli/utils"
  "log"
  "strings"
)
import "io/ioutil"
import "os/user"
// import "net/url"
import "os"
// import "bytes"
// import "time"
// import "net/http"
var githubURL string

func GetCreateFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
    	Name:  "store",
			Usage: "The store to which the artifact belongs to like Maven,gradle,python,npm,ruby etc",
		},cli.StringFlag{
			Name:  "packagename",
			Usage: "Name of the package in the format groupid:artifactid",
		},
    cli.StringFlag{
			Name:  "version",
			Usage: "Version of the package",
		},cli.StringFlag{
			Name:  "classifier",
			Usage: "Classifier for package (like javadocs,null,...)",
		},cli.StringFlag{
			Name:  "external_download_url",
			Usage: "External download URL for the package",
		},cli.StringFlag{
			Name:  "sources_url",
			Usage: "Sources URL can be null for external packages, for internal use github repo url",
		},

	}
	return flags
}

func AddArtifact(c *cli.Context) error{
	ArtStore := strings.ToUpper(c.String("store"))
	ArtPackageName := c.String("packagename")
  Version := c.String("version")
  Classifier := c.String("classifier")
  ExtDwnldURL := c.String("external_download_url")
  SourceURL := c.String("sources_url")
  usr,_ := user.Current()
  User := usr.Username

  log.Println("Store: ",ArtStore,"\nPackage Name: ",ArtPackageName,"\nVersion: ",Version,"\nClassifier: ",Classifier)
  log.Println("external_download_url: ",ExtDwnldURL,"\nsources_url: ",SourceURL,"\nUser: ",User)

  returnStatus,returnMessage := validateCreateInputs(ArtStore,ArtPackageName)
  if returnStatus==1{
    log.Println("Input validation Failed : ",returnMessage)
    fmt.Println("Input validation Failed : ",returnMessage,"\nCheck the format below")
    cli.ShowCommandHelp(c, "create")
    os.Exit(1)
  }

  //make create api call
  createURL :=utils.ArtEndPoint+"/artifactory/v1.0/artifacts/create"
  payloadForCreate := "[{\"artifact\": {\"package_name\": \""+ ArtPackageName +"\",\"version\": \"" +Version+ "\" },\"classifier\":\""+Classifier +"\",\"store\": \""+ ArtStore +"\",\"requested_by\": \""+ User +"\",\"external_download_url\": \""+ ExtDwnldURL +"\",\"sources_url\": \""+ SourceURL+ "\"}]"
  contentType := "application/json"

  log.Println("payloadForCreate\n",payloadForCreate)
  resp,err := utils.MakePostRequest(createURL,payloadForCreate,contentType)
  //defer resp.Body.Close()
  //fmt.Println("resp.StatusCode",resp.StatusCode)
  //fmt.Println("resp.Header",resp.Header)
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  bodyString := string(bodyBytes)
  fmt.Println(bodyString)

  status, message := utils.ValidateHTTPResponse(resp, err, "Add artifact")
  log.Println("status  ",status,"\n message ", message)
	if status != 0 {
      fmt.Println(message)
      os.Exit(1)
  }
  //fmt.Println("Request made for adding artifact to Art repository\n")
	return nil
}

func validateCreateInputs(ArtStore string,ArtPackageName string) (int,string) {
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
