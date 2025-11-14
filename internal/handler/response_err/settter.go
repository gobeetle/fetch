package response_err

func (r *ResponseErr) WithError(err error) {
	r.Err = err
}

func (r *ResponseErr) WithUseErrorUnwrapper() {
	r.UseErrorUnwrapper = true
}
