package kv

import (
	"fmt"
	"testing"
	"time"
)

var m = New[int, int]()

// testing Set Get HasKey & delet functions
func Test_runAll(t *testing.T) {
	fmt.Println("================= sync mod ===================")

	intmap := New[int, int]()
	intmap.Set(1, 123)
	intval, _ := intmap.Get(1)
	if intval != 123 {
		t.Errorf("key must be %d", 123)
	}

	if ok := intmap.HasKey(1); ok != true {
		t.Error("'1' key must be exist")
	}

	//
	strmap := New[string, string]()
	strmap.Set("hi", "hello")
	sval, _ := strmap.Get("hi")
	if sval != "hello" {
		t.Errorf("value must be %s", "hello")
	}

	if ok := strmap.HasKey("hi"); ok != true {
		t.Error("'hi' key must be exist")
	}

	fmt.Println("All Functions pass")
	fmt.Println("=============== concurrent mod ==============")
}

// testing Set and HasKey function in parallel
func Test_Set_HasKey(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000000; i++ {
		if ok := m.HasKey(i); ok == true {
			t.Errorf("%d key must be not exist", i)
		}
		m.Set(i, i*2)
		if ok := m.HasKey(i); ok == false {
			t.Errorf("%d key must be exist", i)
		}
	}
	fmt.Println("Set func Pass")
	fmt.Println("HasKey func Pass")
}

// testing Get function in parallel
func Test_Get(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Nanosecond * 1)
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

	fmt.Println("Get func Pass ")
}

// testing Delet function in parallel
func Test_Delete(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond * 150)

	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Nanosecond * 1)
		m.Delete(i)
	}
	if len(m.data) != 0 {
		t.Errorf("len data must be %d not %d\n", 0, len(m.data))
	}

	fmt.Println("Delete functions Pass")
}

/*
// bechmark
func Bench_Set(t *testing.B) {}

func Bench_Get(t *testing.T) {}

func Bench_Delete(t *testing.T) {}

func Bench_HasKey(t *testing.T) {}
*/
