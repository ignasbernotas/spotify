package main

import (
   "fmt"
   "github.com/89z/parse/protobuf"
   "github.com/philpearl/plenc"
   "google.golang.org/protobuf/proto"
)

func main() {
   user := &UserField{
      Value: []byte("hello world"),
   }
   data, err := proto.Marshal(user)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", data)
   msg := protobuf.Parse(data)
   fmt.Printf("%+v\n", msg)
   u2 := userField{
      Value: []byte("hello world"),
   }
   d2, err := plenc.Marshal(nil, u2)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", d2)
   msg = protobuf.Parse(d2)
   fmt.Printf("%+v\n", msg)
}

type userField struct {
   Key *string `plenc:"1"`
   Value []byte  `plenc:"2"`
}
