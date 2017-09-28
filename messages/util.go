package rpic_messages

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

type RpicMessage struct {
	Wrapped interface{}
}

func NewCtlMessage(direction CtlMessage_Direction, ctltype CtlMessage_CtlType) *CtlMessage {
	return &CtlMessage{
		// value == speed
		Value: 2000,
		// 0 - X
		// 1 - Y
		Direction: direction,
		// 0 - Stop
		// 1 - VEHICLE
		// 2 - CAMERA
		CtlType: ctltype,
	}
}

func EncodeCtlMessage(msg *CtlMessage) ([]byte, error) {
	return proto.Marshal(msg)
}

func DecodeCtlMessage(data []byte, len int) (*CtlMessage, error) {
	protodata := &CtlMessage{}
	err := proto.Unmarshal(data[0:len], protodata)
	if err != nil {
		return protodata, err
	}
	return protodata, nil
}

func NewImageMessage(image64 string) *ImgMessage {
	return &ImgMessage{
		Value: image64,
		Time:  ptypes.TimestampNow(),
	}
}

func EncodeImgMessage(msg *ImgMessage) ([]byte, error) {
	return proto.Marshal(msg)
}

// func DecodeMessage(data []byte, len int) (*message, error) {
//	protoData, err := DecodeCtlMessage(data, len)
//	if err == nil {
//		return protoData, nil
//	}
// }

func DecodeImgMessage(data []byte, len int) (*ImgMessage, error) {
	protodata := &ImgMessage{}
	err := proto.Unmarshal(data[0:len], protodata) //
	return protodata, err
}
