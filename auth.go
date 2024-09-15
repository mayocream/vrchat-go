package vrchat

import (
	"fmt"
)

// Authenticate authenticates the client with the VRChat API
func (c *Client) Authenticate(username, password, totp string) error {
	c.client.SetHeaders(map[string]string{
		"User-Agent":   "vrchat-go/0.0.0 mayo@linux.com",
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	resp, err := c.client.R().
		SetBasicAuth(username, password).
		SetBody(map[string]string{
			"code": totp,
		}).
		Post("/auth/twofactorauth/totp/verify")
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("failed to authenticate: %s", resp.String())
	}

	cookies := resp.Cookies()
	c.client.SetCookies(cookies)

	return nil
}
