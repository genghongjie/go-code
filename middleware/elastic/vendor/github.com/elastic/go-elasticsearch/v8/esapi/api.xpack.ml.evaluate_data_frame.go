// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 8.3.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newMLEvaluateDataFrameFunc(t Transport) MLEvaluateDataFrame {
	return func(body io.Reader, o ...func(*MLEvaluateDataFrameRequest)) (*Response, error) {
		var r = MLEvaluateDataFrameRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// MLEvaluateDataFrame - Evaluates the data frame analytics for an annotated index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html.
//
type MLEvaluateDataFrame func(body io.Reader, o ...func(*MLEvaluateDataFrameRequest)) (*Response, error)

// MLEvaluateDataFrameRequest configures the ML Evaluate Data Frame API request.
//
type MLEvaluateDataFrameRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLEvaluateDataFrameRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + len("/_ml/data_frame/_evaluate"))
	path.WriteString("http://")
	path.WriteString("/_ml/data_frame/_evaluate")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f MLEvaluateDataFrame) WithContext(v context.Context) func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLEvaluateDataFrame) WithPretty() func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLEvaluateDataFrame) WithHuman() func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLEvaluateDataFrame) WithErrorTrace() func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLEvaluateDataFrame) WithFilterPath(v ...string) func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f MLEvaluateDataFrame) WithHeader(h map[string]string) func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f MLEvaluateDataFrame) WithOpaqueID(s string) func(*MLEvaluateDataFrameRequest) {
	return func(r *MLEvaluateDataFrameRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
