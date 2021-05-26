package app

import (
	"encoding/json"
	"net/http"
)

//JSendJSON ...
type JSendJSON struct {
	data []byte
	code int
}

func (js *JSendJSON) Send(w http.ResponseWriter) {
	w.WriteHeader(js.code)
	w.Write(js.data)
}

//IJSendJSONBuilder ...
type IJSendJSONBuilder interface {
	Code(int) IJSendJSONBuilder
	Data(interface{}) IJSendJSONBuilder
	Message(string) IJSendJSONBuilder
	Build() *JSendJSON
}

//JSendJSONBuilder ...
type JSendJSONBuilder struct {
	code    int
	data    interface{}
	message string
}

//Code ...
func (builder *JSendJSONBuilder) Code(code int) IJSendJSONBuilder {
	builder.code = code
	return builder
}

//Data ...
func (builder *JSendJSONBuilder) Data(data interface{}) IJSendJSONBuilder {
	builder.data = data
	return builder
}

//Message ...
func (builder *JSendJSONBuilder) Message(message string) IJSendJSONBuilder {
	builder.message = message
	return builder
}

//Build ...
func (builder *JSendJSONBuilder) Build() *JSendJSON {
	data := make(map[string]interface{})
	code := builder.code
	res := &JSendJSON{}
	if builder.data != nil {
		data["data"] = builder.data
	}

	if code >= 200 && code < 300 {
		data["status"] = "success"
	}
	if code >= 400 && code < 500 {
		data["status"] = "fail"
	}
	if code >= 500 {
		data["status"] = "error"
	}
	if builder.message != "" {
		data["message"] = builder.message
	}
	sentBytes, err := json.Marshal(data)
	if err != nil {
		code = 500
	} else {
		res.data = sentBytes
	}
	res.code = code

	return res
}

//NewJSendJSONBuilder ...
func NewJSendJSONBuilder() IJSendJSONBuilder {
	return &JSendJSONBuilder{}
}
