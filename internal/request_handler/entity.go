package request_handler

import "github.com/gobeetle/fetch/internal/iface"

type RequestHandler[Req any] struct {
	/*
		Location
	*/
	iface.RequestLocationHandler

	/*
		method
	*/
	iface.RequestMethodHandler

	/*
		request
	*/
	iface.RequestRequestHandler

	/*
		client
	*/
	iface.RequestClientHandler

	/*
		header
	*/
	iface.RequestHeaderHandler

	/*
		body
	*/
	iface.RequestBodyHandler[Req]

	/*
		curl
	*/
	iface.RequestCurlHandler
}
