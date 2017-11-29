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
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    }

    fmt.Println(string(body))
    err = ioutil.WriteFile("output.json", body, 0644)
}
