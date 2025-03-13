package response

import (
	"github.com/nelsonfalves/boot-help/util"
)

func (mr *Response) Content(content any) *Response {
	if mr != nil {
		mr.contentData = content
	}
	return mr
}

func (mr *Response) Status(status int) *Response {
	if mr != nil && status > 99 && status < 600 {
		mr.statusCode = status
	}
	return mr
}

func (mr *Response) Type(contentType string) *Response {
	if mr != nil && !util.EmptyString(contentType) {
		mr.contentType = contentType
	}
	return mr
}

func (mr *Response) Header(key, value string) *Response {
	if mr != nil && !util.EmptyString(key) {
		mr.headers[key] = value
	}
	return mr
}

func (mr *Response) Headers(headers map[string]string) *Response {
	if mr == nil {
		return mr
	}
	for k, v := range headers {
		if util.EmptyString(k) {
			continue
		}
		mr.headers[k] = v
	}
	return mr
}
