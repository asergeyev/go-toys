package reuse

import "testing"

func TestServer(t *testing.T) {
	var err error

	port := 5555
	workers := make([]*Service, 3)

	for i := 0; i < len(workers); i++ {
		workers[i], err = NewService(port)
		if err != nil {
			t.Fatalf("Failed to create clone #%d %s", i, err)
		}
	}
	
	// success
}
