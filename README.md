# gonati
_A simple interface for Luminati's proxies!_

## Usage
```go
package main

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "github.com/qAMP/gonati"
)

func main() {
  client := gonati.CreateProxy("lum-customer-XX-zone-YY", "pass", 1337) // Substitute LEET for a port (22225)

  request, err := http.NewRequest("GET", "https://api.ipify.org?format=json", nil)

  resp, err := client.Do(request)
  if err != nil {
    log.Println(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
  }

  fmt.Println(string(body))
}
```