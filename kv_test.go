package kv

import (
	"testing"
	"time"
)

var m = New[int, int]()

func Test_Set(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000000; i++ {

		m.Set(i, i*2)
	}
}

func Test_Get(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 1000000; i++ {

		val, err := m.Get(i)
		if err != nil {
			t.Error(err)
			return
		}
		if val != i*2 {
			t.Errorf("value must be %d", i*2)
			return
		}

	}

}

func Test_HasKey(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond * 150)

	for i := 0; i < 1000000; i++ {
		yes := m.HasKey(i)
		if yes != true {
			t.Errorf("%d must be exist", i)
			break
			// return
		}
	}
}

func Test_Delete(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond * 200)

	for i := 0; i < 1000000; i++ {
		m.Delete(i)
	}
	if len(m.data) != 0 {
		t.Errorf("len data must be %d not %d\n", 0, len(m.data))
	}

}

/*
// bechmark
func Bench_Set(t *testing.B) {}

func Bench_Get(t *testing.T) {}

func Bench_Delete(t *testing.T) {}

func Bench_HasKey(t *testing.T) {}
*/
