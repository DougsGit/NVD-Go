//Example strings can be found: https://nvd.nist.gov/developers/vulnerabilities
package main

import (
  "fmt"
  "net/http"
  //"io/ioutil"
  "io"
)

func main() {

  url := "https://services.nvd.nist.gov/rest/json/cves/2.0?cveId=CVE-2023-35078"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))     //-----This line prints the output to stdout
  //Below are my initial attempts to spit the output out in nicely formated JSON
  
}
