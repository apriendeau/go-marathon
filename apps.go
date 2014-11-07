package marathon

import "fmt"

// CreateApp builds an app in marathon
func (c Client) CreateApp() {
	c.HTTP.Method = "POST"
	fmt.Println("test")
}
