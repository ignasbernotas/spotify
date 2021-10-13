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
   in := header{
      UserFields: []UserField{
         {Key: key}, {Key: key},
      },
   }
   data, err := plenc.Marshal(nil, in)
   if err != nil {
      panic(err)
   }
   var out pb.Header
   if err := proto.Unmarshal(data, &out); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", out)
}
