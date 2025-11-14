package request_handler

import "github.com/gobeetle/fetch/internal/iface"

type RequestHandler struct {
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
	iface.RequestBodyHandler
}
