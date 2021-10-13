package main

import (
   "fmt"
   "github.com/segmentio/encoding/proto"
   "google.golang.org/protobuf/testing/protopack"
)

var BytesType = protopack.BytesType

type (
   Bytes = protopack.Bytes
   LengthPrefix = protopack.LengthPrefix
   Message = protopack.Message
   Tag = protopack.Tag
)

func main() {
   in := Message{
      Tag{6, BytesType}, LengthPrefix{
         Tag{1, BytesType}, Bytes("Vary"),
      },
      Tag{6, BytesType}, LengthPrefix{
         Tag{1, BytesType}, Bytes("MD-Version"),
      },
   }
   data := in.Marshal()
   var out struct {
      Fields []struct {
         Key string `protobuf:"bytes,1"`
      } `protobuf:"bytes,6"`
   }
   err := proto.Unmarshal(data, &out)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", out)
}
