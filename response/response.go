package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gawbsouza/boot-help/httperr"
	"github.com/gawbsouza/boot-help/util"
)

type Response struct {
	writer      http.ResponseWriter
	err         *httperr.HttpError
	headers     map[string]string
	statusCode  int
	contentType string
	contentData any
}

func To(w http.ResponseWriter) *Response {
	return &Response{
		writer:     w,
		statusCode: http.StatusOK,
	}
}

func (mr *Response) Send() error {
	if mr.writer == nil {
		return nil
	}
	if util.EmptyString(mr.contentType) {
		return mr.SendText()
	}
	mr.writeContentTypeAndHeaders(mr.contentType)
	return mr.responseAsText()
}

func (mr *Response) SendText() error {
	if mr.writer == nil {
		return nil
	}
	mr.writeContentTypeAndHeaders("text/plain; charset=utf-8")
	return mr.responseAsText()
}

func (mr *Response) responseAsText() error {
	if mr.err != nil {
		_, err := mr.writer.Write([]byte(mr.err.String()))
		return err
	}
	if s, ok := mr.contentData.(fmt.Stringer); ok {
		_, err := mr.writer.Write([]byte(s.String()))
		return err
	}
	s := fmt.Sprintf("%+v", mr.contentData)
	_, err := mr.writer.Write([]byte(s))
	return err
}

func (mr *Response) SendJSON() {
	if mr.writer == nil {
		return
	}
	mr.writeContentTypeAndHeaders("application/json; charset=utf-8")
	if mr.err != nil {
		json.NewEncoder(mr.writer).Encode(mr.err)
	} else {
		json.NewEncoder(mr.writer).Encode(mr.contentData)
	}
}

func (mr *Response) writeContentTypeAndHeaders(contentType string) {
	mr.writer.Header().Add("Content-Type", contentType)
	mr.writeHeaders()
	mr.writeStatusCode()
}

func (mr *Response) writeHeaders() {
	for k, v := range mr.headers {
		mr.writer.Header().Add(k, v)
	}
}

func (mr *Response) writeStatusCode() {
	if mr.err != nil {
		mr.writer.WriteHeader(mr.err.StatusCode)
	} else {
		mr.writer.WriteHeader(mr.statusCode)
	}
}
