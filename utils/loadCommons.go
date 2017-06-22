package utils

import(
    "os"
)

var (
   Home string
   Pwd string
)

func LoadGlobalVariables() {
    Home = os.Getenv("HOME")
    Pwd = os.Getenv("PWD")
}
