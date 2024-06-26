// An imaginary 3rd-party service which provides lucky numbers.
package numbers

import "math/rand"

type Client interface {
	Get() int
}

// The real client
type client struct{}

func NewClient() *client {
	return &client{}
}

// In reality, this would make a call to the external service.
func (c *client) Get() int {
	return rand.Intn(100)
}

// The mock client
type mockClient struct {
}

func NewMockClient() *mockClient {
	return &mockClient{}
}

func (c *mockClient) Get() int {
	return 42
}
