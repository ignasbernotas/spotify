package main

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
)

func main() {
   key := "key"
   head := &pb.Header{
      UserFields: []*pb.UserField{
         {Key: &key}, {Key: &key},
      },
   }
   out, err := proto.Marshal(head)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", out)
}
