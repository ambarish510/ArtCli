package prehook

import (
        "fmt"
        "os"
        "strings"
        "github.com/Flipkart/asb/util/utilFunctions"
)

func AutoUpdate()  {
  downloadFromGithub()
}

func downloadFromGithub(){
  //curl -LJO# -H 'Accept: application/octet-stream' "https://api.github.com/repos/Flipkart/artifactory-service/releases/assets/4167867?access_token=3dbc5aaf29eecdfed24acc76496b8433a839805c"
  //make a get call to this url
}
