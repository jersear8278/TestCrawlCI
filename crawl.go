package main

import (
    //"fmt"
    //"html"
    "net/http"
    "io/ioutil"
    "fmt"
)

func main() {
    resp, err := http.Get("http://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20171125&stockNo=2330&_=1511567214171")
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
    err = ioutil.WriteFile("output.json", body, 0644)
}
