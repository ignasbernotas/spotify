package main

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "github.com/philpearl/plenc"
)

type UserField struct {
   Key string `plenc:"1"`
}

type header struct {
   UserFields []UserField `plenc:"6"`
}

func main() {
   key := "key"
   in := &pb.Header{
      UserFields: []*pb.UserField{
         {Key: &key}, {Key: &key},
      },
   }
   data, err := proto.Marshal(in)
   if err != nil {
      panic(err)
   }
   var out header
   if err := plenc.Unmarshal(data, &out); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", out)
}
