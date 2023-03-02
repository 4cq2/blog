package main

import (
   "fmt"
   "net/http"
   "net/url"
   "os"
   "strings"
   "time"
)

type username struct {
   one byte
   two byte
   three byte
}

func (u username) String() string {
   var b strings.Builder
   b.WriteByte(u.one)
   b.WriteByte(u.two)
   b.WriteByte(u.three)
   return b.String()
}

var alpha = []byte{
   '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e',
   'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
   'u', 'v', 'w', 'x', 'y', 'z',
}

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "HEAD"
   req.URL = new(url.URL)
   req.URL.Host = "github.com"
   req.URL.Scheme = "https"
   u := username{one: 'i'}
   // u := username{one: 'j'}
   // u := username{one: 'k'}
   // u := username{one: 'l'}
   // u := username{one: 'm'}
   // u := username{one: 'n'}
   // u := username{one: 'o'}
   // u := username{one: 'p'}
   // u := username{one: 'q'}
   // u := username{one: 'r'}
   // u := username{one: 's'}
   // u := username{one: 't'}
   // u := username{one: 'u'}
   // u := username{one: 'v'}
   // u := username{one: 'w'}
   // u := username{one: 'x'}
   // u := username{one: 'y'}
   // u := username{one: 'z'}
   for _, u.two = range alpha {
      for _, u.three = range alpha {
         if u.String() <= "elt" {
            continue
         }
         req.URL.Path = "/" + u.String()
         res, err := new(http.Transport).RoundTrip(req)
         if err != nil {
            panic(err)
         }
         if err := res.Body.Close(); err != nil {
            panic(err)
         }
         fmt.Println(res.Status, u)
         if res.StatusCode == http.StatusNotFound {
            file, err := os.Create(u.String())
            if err != nil {
               panic(err)
            }
            if err := file.Close(); err != nil {
               panic(err)
            }
         } else if res.StatusCode != http.StatusOK {
            panic(res.Status)
         }
         time.Sleep(399 * time.Millisecond)
      }
   }
}
