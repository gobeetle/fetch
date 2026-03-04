/*
Originally taken from https://github.com/moul/http2curl

The MIT License (MIT)

Copyright (c) 2015 Manfred Touron

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package http2curl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

// CurlCommand contains exec.Command compatible slice + helpers
type CurlCommand struct {
	slice []string
}

// append appends a string to the CurlCommand
func (c *CurlCommand) append(newSlice ...string) {
	c.slice = append(c.slice, newSlice...)
}

// String returns a ready to copy/paste command
func (c *CurlCommand) String() string {
	return strings.Join(c.slice, " ")
}

// nopCloser is used to create a new io.ReadCloser for req.Body
type nopCloser struct {
	io.Reader
}

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

func (nopCloser) Close() error { return nil }

// GetCurlCommand returns a CurlCommand corresponding to an http.Request
func GetCurlCommand(req *http.Request) (*CurlCommand, error) {
	command := CurlCommand{}
	command.append("curl")
	command.append("-X", bashEscape(req.Method))
	if req.Body != nil || req.GetBody != nil {
		var bodyReader io.ReadCloser
		if req.GetBody != nil {
			r, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			defer r.Close()
			bodyReader = r
		} else {
			bodyReader = req.Body
		}
		body, err := io.ReadAll(bodyReader)
		if err != nil {
			return nil, err
		}
		if len(body) > 0 {
			bodyEscaped := bashEscape(string(body))
			command.append("-d", bodyEscaped)
		}
		if req.GetBody != nil {
			newBody, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = newBody
		} else {
			req.Body = nopCloser{bytes.NewBuffer(body)}
		}
	}
	var keys []string
	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}
	command.append(bashEscape(req.URL.String()))
	return &command, nil
}
