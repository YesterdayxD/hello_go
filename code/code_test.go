package code

import "testing"

func TestMyAdd(t *testing.T) {

	for _, value := range []struct {
		x      int
		y      int
		result int
	}{
		{1, 9, 10},
		{1, 3, 4},
		{1, 2, 3},
		{1, 2, 3},
	} {
		if value.result != myAdd(value.x, value.y) {
			t.Error("result is wrong")
		} else {
			t.Log("result is right")
		}

	}
}
func BenchmarkMyAdd(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		myAdd(4, 5)
	}
}
