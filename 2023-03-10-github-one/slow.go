package main

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

var possibles = []username{
"asdfasdfasdfadsasfdfads",
"i0y",
"i16",
"i1i",
"i24",
"i2z",
"i32",
"i3f",
"i3j",
"i4f",
"i4i",
"i4w",
"i5b",
"i5p",
"i5r",
"i5y",
"i68",
"i6f",
"i6j",
"i6p",
"i6w",
"i6z",
"i79",
"i7w",
"i88",
"i8b",
"i8d",
"i8e",
"i8l",
"i8p",
"i8w",
"i96",
"i9o",
"i9q",
"i9v",
"ib5",
"idv",
"ie3",
"ie5",
"ie7",
"ie8",
"iei",
"ifr",
"ih4",
"ih6",
"ihv",
"ihy",
"ii6",
"ij9",
"ilt",
"iml",
"iq9",
"iro",
"is1",
"is5",
"iuu",
"iv5",
"iw7",
"iy0",
"iy3",
"iy5",
"iy6",
"izg",
"izl",
}

type username string

func (u username) body() string {
   var b strings.Builder
   b.WriteString("-----------------------------23788153391227195155267468389\r\n")
   b.WriteString("Content-Disposition: form-data; name=\"authenticity_token\"\r\n")
   b.WriteString("\r\n")
   b.WriteString("VCGr4tRCiI8FmJlGUIQTxN8O6FowDkDl2GjY5fqeSnfveH8KxoCK/Km2fvtKItL5FigJQ/XExp9TKd1R8OzWBw==\r\n")
   b.WriteString("-----------------------------23788153391227195155267468389\r\n")
   b.WriteString("Content-Disposition: form-data; name=\"value\"\r\n")
   b.WriteString("\r\n")
   b.WriteString(string(u))
   b.WriteString("\r\n")
   b.WriteString("-----------------------------23788153391227195155267468389--\r\n")
   return b.String()
}

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Content-Type"] = []string{"multipart/form-data; boundary=---------------------------23788153391227195155267468389"}
   req.Header["Cookie"] = []string{"_gh_sess=EuASVLi7BSRIISQHuC1wgQjXoJGMAYZbvWgUWBGXJzHgN6LSsxO038tCJGc4eVDRAoFfaxYUdAs6XIIU%2BWtvqpG7OYAo%2FycxXUo7cicDlTclDHdjU%2FZDSXn8MvBAt2bCtVzAgIZui2miHOVKYqkAn7Pj%2FDsCM88CuHLFtVhW3zqz4o9holu5WcBfdgSmakkIt2b3LYjNRjXr8fKlkp43fAJQ3lUwDgPqoebdkMUCih0ixopuADJ%2B2a1mXzy39YGQd8NcIf%2BT416LcmBo8VSaa9D7ip2uD72S3Hd2aGbMxZXIFpy2--k4%2BAjsYWyCJhZ%2FZE--lRe503VvZlkXdRTcVcPbKQ%3D%3D"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "github.com"
   req.URL.Path = "/signup_check/username"
   req.URL.Scheme = "https"
   sleep := time.Second
   fmt.Println(sleep)
   time.Sleep(sleep)
   for _, possible := range possibles {
      req.Body = io.NopCloser(strings.NewReader(possible.body()))
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(res.Status, possible)
      time.Sleep(2699 * time.Millisecond)
   }
}
