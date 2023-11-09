package gopop

import (
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	s := NewPop3Server(110, "", false)
	go s.Start()

	time.Sleep(2 * time.Second)

	s.Stop()

	time.Sleep(10 * time.Second)

}
