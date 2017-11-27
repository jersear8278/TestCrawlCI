package main

import (
  "net/http"
  "testing"
)

func Test_API(t *testing.T) {
  _, err := http.Get("http://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20171125&stockNo=2330&_=1511567214171")
  if err != nil {
      t.Error("link failed")
  }
}
