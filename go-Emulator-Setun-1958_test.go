package main

import (
	"fmt"
	"testing"
)

func Test_pow3(t *testing.T) {
	var i int8
	for i = 0; i < 18; i++ {
		fmt.Printf("pow3(%d)=%v\n", i, pow3(i))
	}
}

func Benchmark_pow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3(31)
	}
}

func Benchmark_shift_ts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x trs
		shift_ts(x, -31)
	}
}

func Benchmark_add_half_slowly_t(b *testing.B) {
	var aa trits
	var bb trits
	aa = aa.SetFalse()
	bb = bb.SetFalse()
	for i := 0; i < b.N; i++ {
		add_half_slowly_t(aa, bb)
	}
}

func Benchmark_add_full_t(b *testing.B) {
	var aa trits
	var bb trits
	var cc trits
	aa = aa.SetFalse()
	bb = bb.SetFalse()
	cc = cc.SetFalse()
	for i := 0; i < b.N; i++ {
		add_full_t(aa, bb, cc)
	}
}

func Benchmark_mul_t(b *testing.B) {
	var aa trits
	var bb trits
	for i := 0; i < b.N; i++ {
		mul_t(aa.SetFalse(), bb.SetTrue())
	}
}

func Benchmark_sum_t(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum_t(1, 1, 1)
	}
}

func Benchmark_get_trit(b *testing.B) {

	var t1 uint32
	var t0 uint32
	var p uint8

	p = 0
	t1 = t1 &^ (1 << p)
	t0 |= (1 << p)

	for i := 0; i < b.N; i++ {
		get_trit(t1, t0, p)
	}
}

func Benchmark_sgn_trs(b *testing.B) {
	var x trs
	x.l = 32
	for i := 0; i < b.N; i++ {
		sgn_trs(x)
	}
}

func Benchmark_shift_trs(b *testing.B) {
	var x trs
	x.l = 32
	for i := 0; i < b.N; i++ {
		x = shift_trs(x, -31)
	}
}

func Benchmark_add_trs(b *testing.B) {
	var x trs
	var y trs
	x.l = 32
	y.l = 32
	for i := 0; i < b.N; i++ {
		add_trs(x, y)
	}
}

func Benchmark_sub_trs(b *testing.B) {
	var x trs
	var y trs
	x.l = 32
	y.l = 32
	for i := 0; i < b.N; i++ {
		sub_trs(x, y)
	}
}

func Benchmark_int2trs(b *testing.B) {
	var x trs
	x.l = 32
	for i := 0; i < b.N; i++ {
		int2trs(x, 31, -1)
	}
}

func Benchmark_trs2int(b *testing.B) {
	var x trs
	x.l = 32
	for i := 0; i < b.N; i++ {
		trs2int(x, 31)
	}
}

//func BenchmarkCalculate(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//	}
//}

///func BenchmarkSample(b *testing.B) {
//    b.SetBytes(2)
//    for i := 0; i < b.N; i++ {
//        if x := fmt.Sprintf("%d", 42); x != "42" {
//            b.Fatalf("Unexpected string: %s", x)
//        }
//    }
//}
