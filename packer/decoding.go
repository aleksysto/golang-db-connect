package packer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

//	"github.com/aleksysto/golang-db-connect.git/markers"
)
type Decoder struct {
    reader *bytes.Reader
    buffer *bytes.Buffer
}
func NewDecoder(r *bytes.Reader, b *bytes.Buffer) Decoder {
    return Decoder{
        r,
        b,
    }
}
func Unmarshal(m []byte) {
    d := NewDecoder(bytes.NewReader(m), &bytes.Buffer{})
    err := d.read()
    if err != nil {
        fmt.Println("error")
    }
}

func ReadLength(n []byte) (uint16, error) {
    if len(n) != 2 {
        return 0, errors.New("wrong size")
    }
    length := binary.BigEndian.Uint16(n)
    return length, nil
}

func (d *Decoder) read() error {
    lenData := make([]byte, 2)
    n, err := d.reader.Read(lenData)
    if n != 2 || err != nil {
        return errors.New("Couldnt read length bytes from message")
    }

    chunkLength := binary.BigEndian.Uint16(lenData)
    err = d.readData(chunkLength)
    if err != nil {
        return nil
    }
    return nil
}

func (d *Decoder) readData(length uint16) (error) {
    if length == 0 {
        return nil
    }
    var readData uint16
    output := make([]byte, length)
    data := make([]byte, length - readData)
    
    n, err := d.reader.Read(data)
    if err != nil {
        return err 
    }
    if n == 0 {
        return err
    }
    copy(output, data)
    (*d.buffer).Write(output)
    return nil
}

func (d *Decoder) decode() (interface{}, error) {
    return 21, errors.New("fd")
}
