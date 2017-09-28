package camera

import (
	b64 "encoding/base64"
	m "github.com/frzifus/rpic-server/messages"
	"io"
	"io/ioutil"
	//	"sync"
)

type RpicCamera struct {
	addr        string
	isActive    bool
	imageBuffer []byte
	imageIO     io.Reader
	//	mutex       sync.Mutex
	path string
}

func NewRpicCamera(fd string) *RpicCamera {
	c := &RpicCamera{}
	c.imageBuffer = make([]byte, 0)
	c.read()
	c.path = fd
	return c
}

func (c *RpicCamera) read() {
	var err error
	c.imageBuffer, err = ioutil.ReadFile(c.path)
	if err != io.EOF {
		return
	}
}

func (c *RpicCamera) GetPath() string {
	return c.path
}

func (c *RpicCamera) getImageAsBase64() string {
	return b64.StdEncoding.EncodeToString(c.imageBuffer)
}

func (c *RpicCamera) GetImageBuffer() []byte {
	return c.imageBuffer
}

func (c *RpicCamera) GetProtoMsg() ([]byte, error) {
	c.read()
	imgMsg := m.NewImageMessage(c.getImageAsBase64())
	return m.EncodeImgMessage(imgMsg)
}
