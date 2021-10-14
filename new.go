package spotify

import (
   "github.com/segmentio/encoding/proto"
   "io"
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

func (m *internal) parseResponse(cmd uint8, reader io.Reader) (*response, error) {
   seq, flags, count, err := handleHead(reader)
   if err != nil {
      return nil, err
   }
   seqKey := string(seq)
   pend := m.Pending[seqKey]
   for i := uint16(0); i < count; i++ {
      part, err := parsePart(reader)
      if err != nil {
         return nil, err
      }
      if pend.partial != nil {
         part = append(pend.partial, part...)
         pend.partial = nil
      }
      if i == count-1 && (flags == 2) {
         pend.partial = part
      } else {
         pend.parts = append(pend.parts, part)
      }
   }
   if flags == 1 {
      delete(m.Pending, seqKey)
      hData := pend.parts[0]
      var head Header
      err := proto.Unmarshal(hData, &head)
      if err != nil {
         return nil, err
      }
      return &response{
         headerData: hData,
         payload: pend.parts[1:],
         seqKey: seqKey,
         statusCode: head.StatusCode,
         uri: head.URI,
      }, nil
   } else {
      m.Pending[seqKey] = pend
   }
   return nil, nil
}
