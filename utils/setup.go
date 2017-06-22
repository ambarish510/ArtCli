package utils

import(
    "os"
    "log"
    "fmt"
)
//Set the config values. Later take this from config.yaml
func GetConfigValues() {
  LogFileName = "_artLogFile"
  LogFilePath = "logs"
  TimeoutForHttpRequest = 5
  //ArtEndPoint = "http://10.85.59.116" PROD
  ArtEndPoint = "http://10.85.58.239"
}

func SetUpLog() *os.File{

  currentDate := getCurrentDate()
  fileName := currentDate+LogFileName

  //check if the directory already exists ; else create directory
  createDirectory(LogFilePath)

  //create or open the log file
  fileDescriptor, err := os.OpenFile(LogFilePath+"/"+fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    fmt.Printf("error opening file: %v", err)
	}

  //set the log output to logfile
	log.SetOutput(fileDescriptor)
	return fileDescriptor
}
