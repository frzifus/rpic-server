package rpic_messages_test

import (
	m "github.com/frzifus/rpic-server/messages"
	"testing"
)

func TestNewCtlMessage(t *testing.T) {
	msg := m.NewCtlMessage(0, 1)
	if msg.Direction == 0 && msg.CtlType == 1 {
		return
	}
	t.Fatalf("Couldnt create a CtlMessage")
}

func TestEncodeAndDecodeCtlMessage(t *testing.T) {
	msg := m.NewCtlMessage(0, 1)
	dataE, err := m.EncodeCtlMessage(msg)
	if err != nil {
		t.Fatalf("Encoding CtlMessage failed! %s", err)
	}
	dataD, err := m.DecodeCtlMessage(dataE, len(dataE))
	if err != nil {
		t.Fatalf("Decoding CtlMessage failed! %s", err)
	}
	if dataD.CtlType != 1 {
		t.Fatalf("Decoding CtlMessage failed! Data lost %s", err)
	}
}

func TestNewImageMessage(t *testing.T) {
	imgMsg := m.NewImageMessage("test")
	if imgMsg == nil {
		t.Fatalf("Couldnt create ImageMessage!")
	}
}

func TestEncodeAndDecodeImgMessage(t *testing.T) {
	imgMsg := m.NewImageMessage("test64")
	dataE, err := m.EncodeImgMessage(imgMsg)
	if err != nil {
		t.Fatalf("Encoding ImgMessage failed! %s", err)
	}
	dataD, err := m.DecodeImgMessage(dataE, len(dataE))
	if err != nil {
		t.Fatalf("Decodeing ImgMessage failed! %s", err)
	}
	if dataD.Value != "test64" {
		t.Fatalf("Decoding CtlMessage failed! Data lost %s", err)
	}

}
