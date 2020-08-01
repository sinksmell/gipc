package sem

import (
	"math/rand"
	"testing"
	"time"
)

func TestSem(t *testing.T) {
	s, err := GetSysVSem(8848)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = s.SetValue(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		if err = s.P(); err != nil {
			t.Fatal(err)
			return
		}

		t.Logf("Go")
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		if err = s.V(); err != nil {
			t.Fatal(err)
			return
		}

		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	}

	time.Sleep(time.Second * 10)

	if err := s.Destroy(); err != nil {
		t.Fatal(err)
	}
}
