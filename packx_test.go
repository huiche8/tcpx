package tcpx

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"testing"
)

var pack = Packx{Marshaller: JsonMarshaller{}}

func TestTCPx_Pack_UnPack(t *testing.T) {
	type Request struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	var clientRequest = Request{
		Username: "packx",
		Age:      24,
	}
	buf, e := pack.Pack(1, clientRequest, map[string]interface{}{
		"note": "this is a map note",
	})
	if e != nil {
		panic(e)
	}
	fmt.Println("客户端发送请求:", clientRequest)
	fmt.Println("内容:", buf)

	var serverRequest Request
	message, e := pack.Unpack(buf, &serverRequest)
	if e != nil {
		panic(e)
	}
	fmt.Println("收到客户端请求:", serverRequest)
	fmt.Println("客户端信息:", message)
}

func TestTCPx_Packx_Property(t *testing.T) {
	type Request struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	var clientRequest = Request{
		Username: "packx",
		Age:      24,
	}
	buf, e := pack.Pack(1, clientRequest, map[string]interface{}{
		"note": "this is a map note",
	})
	if e != nil {
		panic(e)
	}
	fmt.Println("客户端发送请求:", clientRequest)
	fmt.Println("内容:", buf)
	fmt.Println(pack.MessageIDOf(buf))
	fmt.Println(pack.HeaderLengthOf(buf))
	fmt.Println(pack.BodyLengthOf(buf))
	fmt.Println(pack.HeaderBytesOf(buf))
	fmt.Println(pack.BodyBytesOf(buf))
	fmt.Println(packx.HeaderOf(buf))

	header, _ := pack.HeaderBytesOf(buf)

	body, _ := pack.BodyBytesOf(buf)

	var result Request
	e = json.Unmarshal(body, &result)
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(result)
	var resultHeader map[string]interface{}
	e = json.Unmarshal(header, &resultHeader)
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(resultHeader)
}

func TestPackx_PackWithBody(t *testing.T) {
	packx := NewPackx(JsonMarshaller{})
	buf ,e :=packx.PackWithBody(1, newBytes(1, 2, 3, 4, 5))
	if e!=nil {
		fmt.Println(errorx.Wrap(e))
		t.Fail()
		return
	}

	body, e:=packx.BodyBytesOf(buf)
	if e!=nil {
		fmt.Println(errorx.Wrap(e))
		t.Fail()
		return
	}
	fmt.Println(body)
	if len(body) != 5 {
		fmt.Println(fmt.Sprintf("body unpack want length 5 but got %d, %v",len(body), body))
		t.Fail()
		return
	}
}
func TestPackWithMarshallerName_UnPackWithUnmarshalName(t *testing.T) {
	packx := NewPackx(JsonMarshaller{})
	buf ,e :=packx.Pack(1, "hello")
	if e!=nil {
		fmt.Println(errorx.Wrap(e))
		t.Fail()
		return
	}
	var receive string
	_,e =UnpackWithMarshallerName(buf ,&receive, "json")
	if e!=nil {
		fmt.Println(errorx.Wrap(e))
		t.Fail()
		return
	}
	if receive != "hello" {
		fmt.Println(fmt.Sprintf("received want %s but got %s", "hello", receive))
		t.Fail()
		return
	}
}
func newBytes(a ...byte) []byte {
	return a
}
