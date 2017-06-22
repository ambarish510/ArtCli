package utils

import(
    "os"
    "log"
    "time"
		"runtime"
    "net/http"
    "bytes"
)

func PrintStackTraceToLogFile(){
		trace := make([]byte, 1024)
		count := runtime.Stack(trace, true)
		log.Printf("Stack of %d bytes: %s\n", count, trace)
}

func getCurrentDate() string{
   currentTime := time.Now().Local()
   return currentTime.Format("2006-01-02")
}

func ExistsFile(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

//  creates  directory if the directory does not exist already
func createDirectory(filePath string) {
  val, err := ExistsFile(filePath)
  if err!=nil {
    log.Fatalf("Error in check for file exists")
  }
  if val==false && err==nil {
    //create log directory
    err = os.Mkdir(filePath,0777)
    if err!=nil {
      log.Fatalf("error creating log directory: %v", err)
    }
  }
}

func Contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item]
    return ok
}

//make the get request
func MakeGetCall(getURL string) (*http.Response, error) {
  client := &http.Client{
    Timeout:  time.Duration(TimeoutForHttpRequest) * time.Second,
  }
  req, err := http.NewRequest("GET", getURL, nil)
  req.Header.Add("Accept", "application/json")
  resp,err := client.Do(req)

  return resp,err
}

func ValidateHTTPResponse(resp *http.Response, err error, actionName string)(int, string){
  if err != nil {
    log.Println("Unable to access the host")
        return 1,"Failed accessing the host"
    }
    log.Println("Http response code: ",resp.StatusCode)
    log.Println("Http response header: ",resp.Header)

    if resp.StatusCode == 401{
      return 1,"Authorization failed"
    }else if resp.StatusCode >= 500{
      return 1,"Unable to " + actionName
    }else if resp.StatusCode >= 400{
      reponseHeaderErr := resp.Header.Get("X-Error")
      if reponseHeaderErr == "" {
        reponseHeaderErr = "Unable to "+ actionName
      }
      return 1,reponseHeaderErr
    }else if resp.StatusCode >= 200 {
      defer resp.Body.Close()
      return 0,"Success"
    }else{
      return 1,"Unable to "+ actionName
    }
}

//make the post request
func MakePostRequest(postURL string, payloadForPost string, contentType string) (*http.Response, error) {
  client := &http.Client{
    Timeout:  time.Duration(TimeoutForHttpRequest) * time.Second,
  }
  var (
    req *http.Request
    err error
    resp *http.Response
  )
  if payloadForPost != "" {
    req, err = http.NewRequest("POST", postURL, bytes.NewBuffer([]byte(payloadForPost)))
  } else {
    req, err = http.NewRequest("POST", postURL, nil)
  }
  req.Header.Set("Content-Type", contentType)
  // if(username=="")&&(password==""){
  //   req.SetBasicAuth(username, password)
  // }
  resp,err = client.Do(req)
  return resp,err
}
