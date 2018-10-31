## Installation

`go get -u github.com/tme-dev/api-client-go`

## Examples

Simple request:

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/url"
    "github.com/tme-dev/api-client-go"
)

func main() {
    api := tmeapi.Client("<YOUR_TOKEN>", "<YOUR_SECRET>")
    
    formValues := url.Values{}
    formValues.Add("Country", "gb")
    formValues.Add("Language", "en")
      
    res, err := api.Request("https://api.tme.eu/Utils/Ping.json", formValues)
    
    if err != nil {
        log.Fatal(err)
    }
    result, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", result)
}

```

Array parameters example:

```go
    formValues.Add("SymbolList[0]", "AX-176")
    formValues.Add("SymbolList[1]", "1N4007-DIO")
    formValues.Add("AmountList[0]", "10")
    formValues.Add("AmountList[1]", "10")
```

Key->value parameters example for Products/Search action:

```go
    formValues.Add("SearchParameter[2][0]", "386527")
```

See TME API [documentation][developers-tme] for more details.

[developers-tme]: https://developers.tme.eu/en/