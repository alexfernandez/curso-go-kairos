package memcached

import "testing"
import "net/http"

func TestStart(t *testing.T) {
	key := "key"
	//value := "value"
	go Start()
	resp, err := http.Get("http://127.0.0.1:8080/get/" + key)
	if err != nil {
		t.Errorf("Could not get first: %v", err)
	}
	if resp == nil {
		t.Errorf("No response")
	}
	Stop()
}
