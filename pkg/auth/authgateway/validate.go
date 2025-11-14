package authgateway

import "fmt"

func (a *Authn) Validate() error {
	if a == nil {
		return fmt.Errorf("auth gateway authn cannot be nil")
	}
	if len(a.login_url_format) == 0 {
		return fmt.Errorf("login_url_format cannot be empty")
	}
	if len(a.access_key) == 0 {
		return fmt.Errorf("access_key cannot be empty")
	}
	return nil
}
