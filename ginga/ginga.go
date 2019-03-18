package ginga

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sort"
	"strconv"
	"time"
)

const (
	TcpBufSize = 8192
)

type Data struct {
	Timestamp int `json:"timestamp"`
	Signature string `json:"signature"`
	Nonce string `json:"nonce"`
	Lock bool `json:"lock"`
}

type Client struct {
    Token string
    Endpoint string
    Nonce string
    conn net.Conn
}

func (cls *Client) Lock() error {
    conn, err := net.Dial("tcp", cls.Endpoint)
    if err != nil {
    	return err
	}

    data := &Data{
		Timestamp: int(time.Now().Unix()),
		Nonce: cls.Nonce,
		Lock: true,
	}
    data.Signature = cls.sign(data)
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	count, _ := conn.Write(d)
	if count == 0 {
		return errors.New("conn write count is 0")
	}

	buf := make([]byte, TcpBufSize)
	count, _ = conn.Read(buf)
	if count == 0 {
		return errors.New("conn read count is 0")
	}

	datA := new(Data)
	err = json.Unmarshal(buf[:count], datA)
	if err != nil {
		return err
	}

	if !cls.checkSign(datA) {
		return errors.New("signature error")
	}

	if !datA.Lock {
		return errors.New("this is unlock session")
	}

	cls.conn = conn
	return nil
}

func (cls *Client) Unlock() error {
    if cls.conn == nil {
    	return nil
 	} else {
 		defer cls.conn.Close()
		data := &Data{
			Timestamp: int(time.Now().Unix()),
			Nonce: cls.Nonce,
			Lock: false,
		}
		data.Signature = cls.sign(data)
		d, err := json.Marshal(data)
		if err != nil {
			return err
		}
		count, _ := cls.conn.Write(d)
		if count == 0 {
			return errors.New("conn write count is 0")
		}

		return nil
	}
}

func (cls *Client) checkSign(data *Data) bool {
	if cls.sign(data) == data.Signature {
		return true
	} else {
		return false
	}
}

func (cls *Client) sign(data *Data) string {
	sl := []string{strconv.Itoa(data.Timestamp), data.Nonce, cls.Token}
	sort.Strings(sl)
	h := sha1.New()

	return fmt.Sprintf("%x", h.Sum(nil))
}
