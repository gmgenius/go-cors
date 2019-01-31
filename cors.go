// Copyright 2018 George Mulokozi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// email: mulokozi80@gmail.com

package gocors

import (
	"log"
	"net/http"
)

//type to handle our cors
type HandleCors struct {
	AllowOrigin  string       //origin to be allowed. Domains
	AllowMethods string       //methods to be allowed in cors
	AllowHeaders string       //headers to be allowed in cors
	HttpHandler  http.Handler //handler to be called to continue with request processing
}

func (h HandleCors) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//check if user has provided structure variables or set default
	if h.AllowOrigin == "" {
		//allow all domains
		h.AllowOrigin = "*"
	}

	if h.AllowMethods == "" {
		//allow get and post methods by default
		h.AllowMethods = "GET, POST"
	}

	if h.AllowHeaders == "" {
		//allow content type for pre-flight request by default
		h.AllowHeaders = "content-type"
	}

	//set headers for our response writer
	writer.Header().Add("Access-Control-Allow-Origin", h.AllowOrigin)
	writer.Header().Add("Access-Control-Allow-Methods", h.AllowMethods)
	writer.Header().Add("Access-Control-Allow-Headers", h.AllowHeaders)
	//if the request method is options, handle it and return prematurely because it only used for pre-flight the request
	//fmt.Println(h.AllowMethods)
	if request.Method == http.MethodOptions {
		//return from here
		return
	}
	//before trying to execute maker sure our handler is not nil
	if h.HttpHandler == nil {
		log.Fatal("please provider http handler")
	}
	h.HttpHandler.ServeHTTP(writer, request)
}
