package protobuf

import (
	"github.com/golang/protobuf/proto"
)

func Npack(pb proto.Message, Command uint32) (buf []byte, err error) {
	body, err := proto.Marshal(pb)
	if err != nil {
		return buf, err
	}

	head := new(CmdHeader)
	head.Init(uint32(len(body)+HEAD_BYTE_LEN), Command)
	buf = head.PackageHead()
	buf = append(buf, body...)
	return buf, err
}

func Nunpack(buf []byte, pb proto.Message) (err error) {
	return proto.Unmarshal(buf, pb)
}

func Nrepack(head *CmdHeader, body []byte) (buf []byte, err error) {
	buf = head.PackageHead()
	buf = append(buf, body...)
	return buf, err
}
