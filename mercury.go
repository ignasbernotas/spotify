package spotify

import (
   "bytes"
   "encoding/binary"
   "fmt"
   "github.com/golang/protobuf/proto"
   "github.com/89z/spotify/pb"
   "io"
   "sync"
)

const ChunkByteSizeK = chunkSizeK * 4

// Number of bytes to skip at the beginning of the file
const OggSkipBytesK = 167

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func handleHead(reader io.Reader) (seq []byte, flags uint8, count uint16, err error) {
	var seqLength uint16
	err = binary.Read(reader, binary.BigEndian, &seqLength)
	if err != nil {
		return
	}
	seq = make([]byte, seqLength)
	_, err = io.ReadFull(reader, seq)
	if err != nil {
		fmt.Println("read seq")
		return
	}

	err = binary.Read(reader, binary.BigEndian, &flags)
	if err != nil {
		fmt.Println("read flags")
		return
	}
	err = binary.Read(reader, binary.BigEndian, &count)
	if err != nil {
		fmt.Println("read count")
		return
	}

	return
}

func parsePart(reader io.Reader) ([]byte, error) {
	var size uint16
	binary.Read(reader, binary.BigEndian, &size)
	buf := make([]byte, size)
	_, err := io.ReadFull(reader, buf)
	return buf, err
}

func encodeRequest(seq []byte, req Request) ([]byte, error) {
	buf, err := encodeMercuryHead(seq, uint16(1+len(req.Payload)), uint8(1))
	if err != nil {
		return nil, err
	}

	header := &pb.Header{
		Uri:    proto.String(req.Uri),
		Method: proto.String(req.Method),
	}

	if req.ContentType != "" {
		header.ContentType = proto.String(req.ContentType)
	}

	headerData, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint16(len(headerData)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(headerData)
	if err != nil {
		return nil, err
	}

	for _, p := range req.Payload {
		err = binary.Write(buf, binary.BigEndian, uint16(len(p)))
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(p)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
func encodeMercuryHead(seq []byte, partsLength uint16, flags uint8) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint16(len(seq)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(seq)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint8(flags))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, partsLength)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type Callback func(Response)

type Internal struct {
	seqLock sync.Mutex
	nextSeq uint32
	Pending map[string]Pending
	Stream  PacketStream
}

func (m *Internal) NextSeq() (uint32, []byte) {
	m.seqLock.Lock()

	seq := make([]byte, 4)
	seqInt := m.nextSeq
	binary.BigEndian.PutUint32(seq, seqInt)
	m.nextSeq += 1
	m.seqLock.Unlock()

	return seqInt, seq
}

func (m *Internal) Request(req Request) (seqKey string, err error) {
	_, seq := m.NextSeq()
	data, err := encodeRequest(seq, req)
	if err != nil {
		return "", err
	}

	var cmd uint8
	switch {
	case req.Method == "SUB":
		cmd = 0xb3
	case req.Method == "UNSUB":
		cmd = 0xb4
	default:
		cmd = 0xb2
	}

	err = m.Stream.SendPacket(cmd, data)
	if err != nil {
		return "", err
	}

	return string(seq), nil
}

func (m *Internal) ParseResponse(cmd uint8, reader io.Reader) (response *Response, err error) {
	seq, flags, count, err := handleHead(reader)
	if err != nil {
		fmt.Println("error handling response", err)
		return
	}
	seqKey := string(seq)
	pending, ok := m.Pending[seqKey]
	if !ok && cmd == 0xb5 {
		pending = Pending{}
	}
	for i := uint16(0); i < count; i++ {
		part, err := parsePart(reader)
		if err != nil {
			fmt.Println("read part")
			return nil, err
		}

		if pending.partial != nil {
			part = append(pending.partial, part...)
			pending.partial = nil
		}

		if i == count-1 && (flags == 2) {
			pending.partial = part
		} else {
			pending.parts = append(pending.parts, part)
		}
	}

	if flags == 1 {
		delete(m.Pending, seqKey)
		return m.completeRequest(cmd, pending, seqKey)
	} else {
		m.Pending[seqKey] = pending
	}
	return nil, nil
}

func (m *Internal) completeRequest(cmd uint8, pending Pending, seqKey string) (response *Response, err error) {
	headerData := pending.parts[0]
	header := &pb.Header{}
	err = proto.Unmarshal(headerData, header)
	if err != nil {
		return nil, err
	}

	return &Response{
		HeaderData: headerData,
		Uri:        *header.Uri,
		Payload:    pending.parts[1:],
		StatusCode: header.GetStatusCode(),
		SeqKey:     seqKey,
	}, nil

}
type Pending struct {
	parts   [][]byte
	partial []byte
}

type Request struct {
	Method      string
	Uri         string
	ContentType string
	Payload     [][]byte
}

type Response struct {
	HeaderData []byte
	Uri        string
	Payload    [][]byte
	StatusCode int32
	SeqKey     string
}

func (res *Response) CombinePayload() []byte {
	body := make([]byte, 0)
	for _, p := range res.Payload {
		body = append(body, p...)
	}
	return body
}
