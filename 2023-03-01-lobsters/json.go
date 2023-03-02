package main

import (
   "errors"
   "encoding/json"
   "fmt"
   "net/http"
   "strings"
   "time"
)

func main() {
   err := JSON(
      //fail
      //2 * time.Minute,
      // pass
      3 * time.Minute,
      
      //fail
      //8 * time.Second,
      
      9 * time.Second,
   )
   if err != nil {
      panic(err)
   }
}

type user struct {
   About string
}

func JSON(one, two time.Duration) error {
   fmt.Println(one)
   time.Sleep(one)
   for i, name := range names {
      res, err := http.Get(name + ".json")
      if err != nil {
         return err
      }
      if res.StatusCode != http.StatusOK {
         return errors.New(res.Status)
      }
      var u user
      if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
         return err
      }
      if strings.Contains(u.About, "github.com/") {
         fmt.Println("pass", len(names)-i, name)
      } else {
         fmt.Println("fail", len(names)-i, name)
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
      time.Sleep(two)
   }
   return nil
}
