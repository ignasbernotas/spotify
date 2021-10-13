package main

import (
   "fmt"
   "github.com/segmentio/encoding/proto"
   "google.golang.org/protobuf/testing/protopack"
)

type header struct {
   Fields []field `protobuf:"bytes,6"`
}

type field struct {
   Key string `protobuf:"bytes,1"`
}

func main() {
   in := header{
      Fields: []field{
         {Key: "Vary"}, {Key: "MD-Version"},
      },
   }
   data, err := proto.Marshal(in)
   if err != nil {
      panic(err)
   }
   var out protopack.Message
   out.UnmarshalAbductive(data, nil)
   fmt.Printf("%+v\n", out)
}
