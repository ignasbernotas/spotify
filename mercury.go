package spotify

import (
   "bytes"
   "encoding/binary"
   "fmt"
   "io"
   "sync"
)

type request struct {
   contentType string
   method      string
   payload     [][]byte
   uri         string
}

type response struct {
   StatusCode int32
   Uri        string
   headerData []byte
   payload    [][]byte
   seqKey     string
}

const chunkByteSizeK = chunkSizeK * 4

// Number of bytes to skip at the beginning of the file
const oggSkipBytesK = 167

func encodeMercuryHead(seq []byte, partsLength uint16, flags uint8) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   err := binary.Write(buf, binary.BigEndian, uint16(len(seq)))
   if err != nil {
      return nil, err
   }
   if _, err := buf.Write(seq); err != nil {
      return nil, err
   }
   if err := binary.Write(buf, binary.BigEndian, uint8(flags)); err != nil {
      return nil, err
   }
   if err := binary.Write(buf, binary.BigEndian, partsLength); err != nil {
      return nil, err
   }
   return buf, nil
}

func handleHead(reader io.Reader) ([]byte, uint8, uint16, error) {
   var seqLength uint16
   err := binary.Read(reader, binary.BigEndian, &seqLength)
   if err != nil {
      return nil, 0, 0, err
   }
   seq := make([]byte, seqLength)
   if _, err := io.ReadFull(reader, seq); err != nil {
      return nil, 0, 0, fmt.Errorf("read seq %v", err)
   }
   var flags uint8
   if err := binary.Read(reader, binary.BigEndian, &flags); err != nil {
      return nil, 0, 0, fmt.Errorf("read flags %v", err)
   }
   var count uint16
   if err := binary.Read(reader, binary.BigEndian, &count); err != nil {
      return nil, 0, 0, fmt.Errorf("read count %v", err)
   }
   return seq, flags, count, nil
}

func min(a, b int) int {
   if a < b {
      return a
   }
   return b
}

func parsePart(reader io.Reader) ([]byte, error) {
   var size uint16
   binary.Read(reader, binary.BigEndian, &size)
   buf := make([]byte, size)
   _, err := io.ReadFull(reader, buf)
   if err != nil {
      return nil, err
   }
   return buf, nil
}

type callback func(response)

type internal struct {
   Pending map[string]pending
   Stream  packetStream
   nextSequence uint32
   seqLock sync.Mutex
}

func (m *internal) nextSeq() (uint32, []byte) {
   m.seqLock.Lock()
   seq := make([]byte, 4)
   seqInt := m.nextSequence
   binary.BigEndian.PutUint32(seq, seqInt)
   m.nextSequence += 1
   m.seqLock.Unlock()
   return seqInt, seq
}

func (m *internal) request(req request) (seqKey string, err error) {
   _, seq := m.nextSeq()
   data, err := encodeRequest(seq, req)
   if err != nil {
      return "", err
   }
   var cmd uint8
   switch {
   case req.method == "SUB":
      cmd = 0xb3
   case req.method == "UNSUB":
      cmd = 0xb4
   default:
      cmd = 0xb2
   }
   err = m.Stream.sendPacket(cmd, data)
   if err != nil {
      return "", err
   }
   return string(seq), nil
}

func (m *internal) parseResponse(cmd uint8, reader io.Reader) (*response, error) {
   seq, flags, count, err := handleHead(reader)
   if err != nil {
      return nil, err
   }
   seqKey := string(seq)
   pend, ok := m.Pending[seqKey]
   if !ok && cmd == 0xb5 {
      pend = pending{}
   }
   for i := uint16(0); i < count; i++ {
      part, err := parsePart(reader)
      if err != nil {
         fmt.Println("read part")
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
      return m.completeRequest(cmd, pend, seqKey)
   } else {
      m.Pending[seqKey] = pend
   }
   return nil, nil
}

type pending struct {
	parts   [][]byte
	partial []byte
}


func (res *response) combinePayload() []byte {
	body := make([]byte, 0)
	for _, p := range res.payload {
		body = append(body, p...)
	}
	return body
}
