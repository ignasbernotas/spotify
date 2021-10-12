package main

import (
   "encoding/json"
   "github.com/89z/parse/protobuf"
   "os"
)

var data = []byte("\n6hm://metadata/4/track/eef38251727f46c28eed9284b288024e\x12\x1avnd.spotify/metadata-track \x90\x032\x13\n\x04Vary\x12\v_ui.country2\x18\n\nMD-Version\x12\n16339896062\x0f\n\x06MC-TTL\x12\x05831502\x19\n\x0fMC-Cache-Policy\x12\x06public2\x0f\n\aMC-ETag\x12\x04V-G\x1e")

func main() {
   fields := protobuf.Parse(data)
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(fields)
}
