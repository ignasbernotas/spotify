package main

import (
   "fmt"
   "github.com/philpearl/plenc"
)

type UserField struct {
   Key string `plenc:"1"`
}

type header struct {
   UserFields []UserField `plenc:"6"`
}

func main() {
   head := header{
      UserFields: []UserField{
         {Key: "key"}, {Key: "key"},
      },
   }
   out, err := plenc.Marshal(nil, head)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", out)
}
