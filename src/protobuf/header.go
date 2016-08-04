package protobuf

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	HEAD_BYTE_LEN = 4 * 8
)

type CmdHeader struct {
	Version   uint32
	Length    uint32
	Command   uint32
	Vendor_id uint32
	Market    uint32
	Is_cksum  uint32
	Check_Sum uint32
	Extend    uint32
}

func GetBodyLen(buf []byte) (bodylen uint32, err error) {
	if len(buf) < HEAD_BYTE_LEN {
		return 0, errors.New("not enough info to get head")
	}
	//网络序采用的是大端，调用大端函数转换为整数（与自身大小端无关）
	bodylen = binary.BigEndian.Uint32(buf[4:8])
	return bodylen, nil
}

func NewHeader(buf []byte) (head *CmdHeader, err error) {
	if len(buf) < HEAD_BYTE_LEN {
		return nil, errors.New("not enough info to get head")
	}

	head = new(CmdHeader)
	headbuf := bytes.NewBuffer(buf[:HEAD_BYTE_LEN])

	//网络序采用的是大端，调用大端函数转换为整数（与自身大小端无关）
	err = binary.Read(headbuf, binary.BigEndian, head)
	return head, err
}

func (head *CmdHeader) Init(Length uint32, Command uint32) {
	head.Length = Length
	head.Command = Command
}

func (head *CmdHeader) PackageHead() (headBuf []byte) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, head)
	headBuf = buf.Bytes()
	return headBuf
}
