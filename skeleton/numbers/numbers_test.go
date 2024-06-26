// An imaginary 3rd-party service which provides lucky numbers.
package numbers

import (
	"fmt"
	"testing"
)

func Test_Client_Get(t *testing.T) {
	client := NewClient()

	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := client.Get(); got < 0 || got > 100 {
				t.Errorf("Client.Get() = %v, which is outside expected range.", got)
			}
		})
	}
}

func Test_mockClient_Get(t *testing.T) {
	client := NewMockClient()
	tests := []struct {
		name string
		want int
	}{
		{"sanity", 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := client.Get(); got != tt.want {
				t.Errorf("mockClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
