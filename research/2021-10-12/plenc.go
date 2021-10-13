package main

import (
   "fmt"
   "github.com/89z/parse/protobuf"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "github.com/philpearl/plenc"
)

type header struct {
   URI string `plenc:"1"`
   ContentType string `plenc:"2"`
   Method string `plenc:"3"`
   StatusCode int32 `plenc:"4"`
   UserFields []struct {
      Key string `plenc:"1"`
      Value string `plenc:"2"`
   } `plenc:"6"`
}

var data = []byte("\n6hm://metadata/4/track/eef38251727f46c28eed9284b288024e\x12\x1avnd.spotify/metadata-track \x90\x032\x13\n\x04Vary\x12\v_ui.country2\x18\n\nMD-Version\x12\n16339896062\x0f\n\x06MC-TTL\x12\x05831502\x19\n\x0fMC-Cache-Policy\x12\x06public2\x0f\n\aMC-ETag\x12\x04V-G\x1e")

func marshal(b []byte, v interface{}) ([]byte, error) {
   err := plenc.Unmarshal(b, v)
   if err != nil {
      return nil, err
   }
   return plenc.Marshal(nil, v)
   var head header
   b2, err := marshal(b1, &head)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", b2)
   msg := protobuf.Parse(b2)
   fmt.Println(msg)
}
