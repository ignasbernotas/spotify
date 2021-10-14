package main

import (
   "bytes"
   "fmt"
   "github.com/segmentio/encoding/proto"
)

type Header struct {
   URI string `protobuf:"bytes,1"`
   ContentType string `protobuf:"bytes,2"`
   Method string `protobuf:"bytes,3"`
   StatusCode int32 `protobuf:"zigzag32,4"`
   UserFields []struct {
      Key string `protobuf:"bytes,1"`
      Value string `protobuf:"bytes,2"`
   } `protobuf:"bytes,6"`
}

var in = []byte("\n6hm://metadata/4/track/eef38251727f46c28eed9284b288024e\x12\x1avnd.spotify/metadata-track \x90\x032\x13\n\x04Vary\x12\v_ui.country2\x18\n\nMD-Version\x12\n16340760052\x0f\n\x06MC-TTL\x12\x05846122\x19\n\x0fMC-Cache-Policy\x12\x06public2\x0f\n\aMC-ETag\x12\x04V-G\x1e")

func main() {
   var head Header
   err := proto.Unmarshal(in, &head)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", head)
   out, err := proto.Marshal(head)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", in)
   fmt.Printf("%q\n", out)
   fmt.Println(bytes.Equal(in, out))
}
