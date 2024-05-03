/**
 * Filename: 	go-Emulator-Setun-1958.go
 *
 * Project: Эмулятор троичной машины 'СЕТУНЬ' SETUN-1958 VM на языке программирования Golang
 *          A virtual machine of ternary computer Setun, also
 *          known as "Small Automatic Digital Machine" written on Golang.
 *
 *          Виртуальная машина троичной цифровой машина "Сетунь" также
 *          известная как МЦВМ "Сетунь"), написанная на языке программирования Golang
 *
 * Create date: 28.08.2021
 * Edit date:   03.05.2024
 *
 * Version:		1.03
 *
 */

// Cross Platform Compilation for ARMv7
// env CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build --ldflags '-linkmode external -extldflags "-static"' .
//
// cgo CFLAGS: -Wall
//

package main

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "./ctrslibs/trslib.h"
#include "./ctrslibs/trslib.c"
*/
import "C"

import (
        "fmt"
        "unsafe"
)

// ---------------------------------------------------
// Main
// ---------------------------------------------------
func main() {

	fmt.Printf("Test call function trslib -----------\n")

	testCallC()

	fmt.Printf("Test ternary functions -----------\n")

	fmt.Printf("- calculate trit-1  --------------\n")
	// Троичные переменные
	var a trits
	var b trits
	var c trits
	var carry trits
	// Операции на тритами
	a.SetNil()
	b.SetFalse()
	c.SetTrue()
	carry.SetTrue()

	fmt.Println(" Set trit:")
	var s trits
	s = s.SetTrue()
	fmt.Println(s.SymbChar())
	s = s.SetNil()
	fmt.Println(s.SymbChar())
	s = s.SetFalse()
	fmt.Println(s.SymbChar())

	s = s.Clear()
	fmt.Println(s.SymbChar())

	var t3 trits
	var t2 trits
	t3 = t3.SetFalse()
	t2 = t2.SetNil()
	fmt.Println(cmp_trits(t3, t2))
	t3 = t3.SetNil()
	t2 = t2.SetNil()
	fmt.Println(cmp_trits(t3, t2))

	fmt.Printf("- add_full_t    --------------\n")
	var aa trits
	var bb trits
	var ccarry trits
	aa = aa.SetFalse()
	bb = bb.SetFalse()
	sf, sfc := add_full_slowly_t(aa, bb, ccarry)
	fmt.Println(sf, sfc)

	fmt.Printf("- calculate trits-32 --------------\n")
	// Троичные переменные
	// TODO

	fmt.Printf("--- Operation Setun-1958 ---\n")
	// Троичные переменные
	// TODO

	fmt.Printf("--------------------------------\n")
}

// -------------------------------------------------------
// TRIT Arithmetic  ver. 2.0 for architectures ARM, RISC-V
// -------------------------------------------------------

// Вызов функций из библиотеки на С
func testCallC() {

	fmt.Println("-------------------------------")

        // C Library
	mystr := C.CString("Hello from a C library function")
	C.myPrintFunction(mystr)
	defer C.free(unsafe.Pointer(mystr))

	// Inline C
	C.myPrintFunction2()

	// Inline C math
	var X C.double
	X = 0.5432
	X = C.sin(X)
	fmt.Println(X)

	s2 := C.s2
	fmt.Println(s2)
	s1 := C.s1
	fmt.Println(s1)

	fmt.Println("-------------------------------")
}

// *****************************************************

// Таб.1 Алфавит троичной симметричной системы счисления
// +--------------------------+-------+-------+-------+
// | Числа                    |  -1   |   0   |   1   |
// +--------------------------+-------+-------+-------+
// | Логика                   | false |  nil  | true  |
// +--------------------------+-------+-------+-------+
// | Логика                   |   f   |   n   |   t   |
// +--------------------------+-------+-------+-------+
// | Символы                  |   -   |	  0   |	  +   |
// +--------------------------+-------+-------+-------+
// | Символы                  |   N   |	  Z   |	  P   |
// +--------------------------+-------+-------+-------+
// | Символы                  |   N   |	  O   |	  P   |
// +--------------------------+-------+-------+-------+
// | Символы                  |   0   |	  i   |	  1   |
// +--------------------------+-------+-------+-------+
// | Символы                  |   v   |	  0   |	  ^   |
// +--------------------------+-------+-------+-------+

// В троичной логике:
//   - 729 коммутативных двухоперандных
//   - 19,683 функций вида F(A,B) = F(B,A)

// ----------------------------------------------------
// TRIT Arithmetic  ver. 1.0
// ----------------------------------------------------

// Объявление троичных типов
type trits struct {
	t uint8 // FALSE,TRUE
	n uint8 // NIL
}

// Метод установить трит в True
func (t trits) SetTrue() (r trits) {
	r.t = 1
	r.n = 1
	return r
}

// Метод установить трит в Nil
func (t trits) SetNil() (r trits) {
	r.t = 0
	r.n = 0
	return r
}

// Метод установить трит в False
func (t trits) SetFalse() (r trits) {
	r.t = 0
	r.n = 1
	return r
}

// Метод трит в False ?
func (t trits) IsFalse() bool {
	if t.n != 0 {
		if t.t == 0 {
			return true
		}
	}
	return false
}

// Метод трит в Nil ?
func (t trits) IsNil() bool {
	if t.n == 0 {
		return true
	}
	return false
}

// Метод трит в Nil ?
func (t trits) IsTrue() bool {
	if t.n != 0 {
		if t.t != 0 {
			return true
		}
	}
	return false
}

// Метод очистить трит
func (t trits) Clear() (r trits) {
	r.t = 0
	r.n = 0
	return r
}

// Метод вернуть символ трита "-1","0","1"
func (t trits) SymbNumb() string {
	if t.n == 0 {
		return "0"
	} else if t.t != 0 {
		return "1"
	} else {
		return "-1"
	}
}

// Метод вернуть символ трита "%false","%nil","%true"
func (t trits) SymbLogic() string {
	if t.n == 0 {
		return "%nil"
	} else if t.t != 0 {
		return "%true"
	} else {
		return "%false"
	}
}

// Метод вернуть символ трита "-","0","+"
func (t trits) SymbTrit() string {
	if t.n == 0 {
		return "0"
	} else if t.t != 0 {
		return "+"
	} else {
		return "-"
	}
}

// Метод вернуть символ трита "M","N","P"
func (t trits) SymbChar() string {
	if t.n == 0 {
		return "N"
	} else if t.t != 0 {
		return "P"
	} else {
		return "M"
	}
}

// Метод вернуть символ трита в int8
func (t trits) ToInt() int8 {
	if t.n == 0 {
		return 0
	} else if t.t != 0 {
		return 1
	} else {
		return -1
	}
}

// Преобразование целое число в трит
func int2trit(i int8) (r trits) {
	if i > 0 {
		return r.SetTrue()
	}
	if i < 0 {
		return r.SetFalse()
	}
	return r.SetNil()
}

// --------------------------
// Функции операций с тритами

// Сравнить триты
func cmp_trits(a trits, b trits) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

// Таб.2 Полусумматор тритов
// .------------------------.
// |     | -1  |  0  |  1   |
// |------------------------|
// | -1  | -+  |  -  |  0   |
// |------------------------|
// |  0  | -1  |  0  |  1   |
// |------------------------|
// |  1  |  0  |  1  |  +-  |
// .------------------------.

// Полусумматор двух тритов с переносом
func add_half_slowly_t(a trits, b trits) (c trits, carry trits) {
	if a.IsFalse() && b.IsFalse() {
		return c.SetTrue(), carry.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return c.SetFalse(), carry.SetTrue()
	} else if a.IsFalse() && b.IsTrue() {
		return c.SetNil(), carry.SetNil()
	} else if a.IsNil() && b.IsFalse() {
		return c.SetFalse(), carry.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return c.SetNil(), carry.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return c.SetTrue(), carry.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return c.SetNil(), carry.SetNil()
	} else if a.IsTrue() && b.IsNil() {
		return c.SetTrue(), carry.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return c.SetFalse(), carry.SetTrue()
	}
	return c.SetNil(), carry.SetNil()
}

// Полусумматор двух тритов с переносом
func add_half_t(a trits, b trits) (c trits, carry trits) {
	switch a.ToInt() + b.ToInt() {
	case -2:
		return c.SetTrue(), carry.SetFalse()
	case -1:
		return c.SetFalse(), carry.SetNil()
	case 0:
		return c.SetNil(), carry.SetNil()
	case 1:
		return c.SetTrue(), carry.SetNil()
	case 2:
		return c.SetFalse(), carry.SetTrue()
	}
	return c.SetNil(), carry.SetNil()
}

// Таб.3 Таблица полного сумматора тритов
// .-----------------------------------------.
// | Перенос из n-1   -1  -1  -1   1   1   1 |
// | 1-е слагаемое    -1  -1  -1   1   1   1 |
// | 2-е слагаемое    -1   0   1  -1   0   1 |
// | Сумма   	       0   1  -1   1  -1   0 |
// | Перенос в n+1    -1  -1   0   0   1   1 |
// .-----------------------------------------.
// Полный сумматор двух тритов с переносом
func add_full_slowly_t(a trits, b trits, incarry trits) (c trits, outcarry trits) {
	s, sc := add_half_slowly_t(a, b)
	d, dc := add_half_slowly_t(s, incarry)
	ss, _ := add_half_slowly_t(sc, dc)
	return d, ss
}

// Полный сумматор двух тритов с переносом
func add_full_t(a trits, b trits, incarry trits) (c trits, outcarry trits) {

	switch a.ToInt() + b.ToInt() + incarry.ToInt() {
	case -3:
		return c.SetNil(), outcarry.SetFalse()
	case -2:
		return c.SetTrue(), outcarry.SetFalse()
	case -1:
		return c.SetFalse(), outcarry.SetNil()
	case 0:
		return c.SetNil(), outcarry.SetNil()
	case 1:
		return c.SetTrue(), outcarry.SetNil()
	case 2:
		return c.SetFalse(), outcarry.SetTrue()
	case 3:
		return c.SetNil(), outcarry.SetTrue()
	}
	return c.SetNil(), outcarry.SetNil()
}

// Таб.4 Троичное умножение
//       MUL
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// Троичное умножение двух тритов с переносом
func mul_slowly_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// Троичное умножение двух тритов с переносом
func mul_t(a trits, b trits) (r trits) {
	switch a.ToInt() * b.ToInt() {
	case 1:
		return r.SetTrue()
	case -1:
		return r.SetFalse()
	case 0:
		return r.SetNil()
	}
	return r.SetNil()
}

// Таб.5 Троичное отрицание
//       NOT
// .-----------.
// |  -  |  +  |
// |-----------|
// |  0  |  0  |
// |-----------|
// |  +  |  -  |
// .-----------.
func not_t(a trits) (r trits) {
	if a.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.6 Троичное умножение
//       AND
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  -  |  -  |
// |-----------------------|
// |  0  |  -  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
//  X AND Y = MIN(X,Y)
// Троичное умножение двух тритов с переносом
func and_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Таб.7 Троичное или
//       OR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  +  |
// |-----------------------|
// |  1  |  +  |  +  |  +  |
// .-----------------------.
//   X OR Y = MAX(X,Y)
func or_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Таб.8 Троичное исключающее или
//       XOR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  +  |  0  |  -  |
// .-----------------------.
func xor_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.9 Троичное
//       EQV
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// EQV(X,Y) = NOT (XOR(X,Y))
func eqv_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Таб.10 Троичное
//       NAND
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  +  |  +  |
// |-----------------------|
// |  0  |  +  |  0  |  0  |
// |-----------------------|
// |  1  |  +  |  0  |  -  |
// .-----------------------.
// NAND(X,Y) = NOT (AND(X,Y))
func nand_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.11 Троичное
//       NOR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  -  |
// |-----------------------|
// |  1  |  -  |  -  |  -  |
// .-----------------------.
// NOR(X,Y) = NOT((OR(X,Y))
func nor_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.12 Троичное
//       IMP
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  +  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
func imp_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Таб.13 Троичное исключающее максимального
//        XMAX
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  -  |  +  |
// |-----------------------|
// |  1  |  +  |  +  |  -  |
// .-----------------------.
// XMAX:
//	F = MAX(A,B), если A != B
//	    0	    , если A == B
//	(имейте в виду - это не XOR).
func xmax_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.14 Троичное Инверсно Исключающий минимального
//        IXMAX
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  -  |  -  |
// |-----------------------|
// |  0  |  -  |  +  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
func ixmax_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Таб.15 Троичное Инверсно Исключающий минимального
//        MEAN
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  +  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  -  |
// .-----------------------.
// Mean:
//	Смотрит насколько "средние" операнды
//	Если ii, то возращает 1
//	Если iX или Xi, то возращает i
//	Иначе 0
func mean_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

// Таб.16 Троичное Инверсно Исключающий минимального
//        Magnitude
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  0  |  -  |  -  |
// |-----------------------|
// |  0  |  +  |  0  |  -  |
// |-----------------------|
// |  1  |  +  |  +  |  0  |
// .-----------------------.
// Magnitude:
//	Сравнение
//	(Функция нессимитричная)
//	Возращает 0, если A < B
//		  i, если A = B
//		  1, если A > B
func magnitude_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetNil()
	}
	return r.SetNil()
}

// Таб.17 Троичное дополнительный код
//        NEG
// .-----------.
// |  -  |  -  |
// |-----------|
// |  0  |  1  |
// |-----------|
// |  +  |  0  |
// .-----------.
func neg_t(a trits) (r trits) {
	if a.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Расмотрим еще некоторые функции
// { -, 0, + }
//
//  Сложение по модулю
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  0
//	0 | -  0  +
//	+ | 0  +  -
//	--+----------
func add_mod_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

//  Перенос в сложении по модулю
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  0  0
//	0 | 0  0  0
//	+ | 0  0  +
//	--+----------
func carry_add_mod_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Сложение с насышением
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  -  0
//	0 | -  0  +
//	+ | 0  +  +
//	--+----------
func add_satiation_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//
//  Функция Webb
//	--+----------
//	  | -  0  +
//	--+----------
//	- | 0  +  -
//	0 | +  +  -
//	+ | -  -  -
//	--+----------
func webb_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetFalse()
	}
	return r.SetNil()
}

//   Тождество (строгоe)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  -
//	0 | -  +  -
//	+ | -  -  +
//	--+----------
func identity_strict_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//   Тождество (weak)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | 0  0  0      в общем как умножение
//	+ | -  0  +
//	--+----------
func weak_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Коньюнкция Лукашевича (сильная)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  -  -
//	0 | -  -  0
//	+ | -  0  +
//	--+----------
func conjunction_lukashevich_strong_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Импликация Лукашевича
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | +  +  0
//	+ | +  +  +
//	--+----------
func lukashevich_implication_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Коньюнкция Клини
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  0  -
//	0 | 0  0  0
//	+ | -  0  +
//	--+----------
func klini_conjunction_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Импликация Клини
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  +  +
//	0 | 0  0  0
//	+ | -  0  +
//	--+----------
func implication_clinic_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetFalse()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Интуиционистская импликация Геделя
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  -
//	0 | +  +  0
//	+ | +  +  +
//	--+----------
func goedel_intuitionistic_implication_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetFalse()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Материальная импликация
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | +  0  0
//	+ | +  +  +
//	--+----------
func material_implication_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetTrue()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

//  Функция следования Брусенцова
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | 0  0  0
//	+ | 0  0  +
//	--+----------
func following_brusentsov_t(a trits, b trits) (r trits) {
	if a.IsFalse() && b.IsFalse() {
		return r.SetTrue()
	} else if a.IsFalse() && b.IsNil() {
		return r.SetNil()
	} else if a.IsFalse() && b.IsTrue() {
		return r.SetFalse()
	} else if a.IsNil() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsNil() && b.IsNil() {
		return r.SetNil()
	} else if a.IsNil() && b.IsTrue() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsFalse() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsNil() {
		return r.SetNil()
	} else if a.IsTrue() && b.IsTrue() {
		return r.SetTrue()
	}
	return r.SetNil()
}

// Троичное сложение двух тритов с переносом
// (Для измерения производительности операций int)
func sum_t(a int8, b int8, p0 int8) (int8, int8) {

	if a > 0 {
		a = 1
	}
	if a < 0 {
		a = -1
	}
	s := a + b + p0
	switch s {
	case -3:
		return 0, 1
		break
	case -2:
		return 1, -1
		break
	case -1:
		return -1, 0
		break
	case 0:
		return 0, 0
		break
	case 1:
		return 1, 0
		break
	case 2:
		return -1, 1
		break
	case 3:
		return 0, 1
		break
	default:
		return 0, 0
		break
	}
	return 0, 0
}

//
// Реализация чтение трита
//
func get_trit(t1 uint32, t0 uint32, p uint8) int {
	var trit int
	if (t0 & (1 << p)) == 0 {
		trit = 0
	} else {
		if (t1 & (1 << p)) == 0 {
			trit = -1
		} else {
			trit = 1
		}
	}
	return trit
}

// ----------------------------------------------------
// 32-TRITS
// ----------------------------------------------------
const TRITSMAX = 32

// Троичный тип данных
type trs struct {
	l  uint8  // длина троичного числа в тритах
	t1 uint32 // двоичное битовое поле троичного числа {FALSE,TRUE}
	t0 uint32 // двоичное битовое поле троичного числа {NIL}
	// if t0[i] == 0 then trit[i] = NIL emdif
	// if (t0[i] == 1) && (t1[i] == 0)  then trit[i] = FALSE emdif //
	// if (t0[i] == 1) && (t1[i] == 1)  then trit[i] = TRUE emdif //
}

// Метод получить трит в позиции троичного числа
func (ts trs) GetTrit(p uint8) int8 {
	if p > TRITSMAX {
		return 0
	}

	if (ts.t0 & (1 << p)) == 0 {
		return 0
	} else if (ts.t1 & (1 << p)) != 0 {
		return 1
	} else {
		return 0
	}
	return 0
}

// Метод получить трит в позиции троичного числа
func (ts trs) SetTrit(p uint8, t int8) (r trs) {
	if p > TRITSMAX {
		return ts
	}

	if t > 0 {
		ts.t1 |= (1 << p)
		ts.t0 |= (1 << p)
	}
	if t < 0 {
		ts.t1 &^= (1 << p)
		ts.t0 |= (1 << p)
	}
	ts.t1 &^= (1 << p)
	ts.t0 &^= (1 << p)

	return ts
}

// Метод установить трит в True
func (ts trs) SetTrue(p uint8) trs {
	if p > TRITSMAX {
		p = TRITSMAX
	}
	ts.t1 |= (1 << p)
	ts.t0 |= (1 << p)
	return ts
}

// Метод установить трит в Nil
func (ts trs) SetNil(p uint8) trs {
	if p > TRITSMAX {
		p = TRITSMAX
	}
	ts.t1 &^= (1 << p)
	ts.t0 &^= (1 << p)
	return ts
}

// Метод установить трит в False
func (ts trs) SetFalse(p uint8) trs {
	if p > TRITSMAX {
		p = TRITSMAX
	}
	ts.t1 &^= (1 << p)
	ts.t0 |= (1 << p)
	return ts
}

// Метод трит в False ?
func (ts trs) IsFalse(p uint8) bool {

	if (ts.t0 & (1 << p)) == 0 {
		return false
	}
	if (ts.t1 & (1 << p)) == 0 {
		return true
	} else {
		return false
	}
}

// Метод трит в Nil ?
func (ts trs) IsNil(p uint8) bool {
	if (ts.t0 & (1 << p)) == 0 {
		return true
	}
	if (ts.t1 & (1 << p)) == 0 {
		return false
	} else {
		return false
	}
}

// Метод трит в Nil ?
func (ts trs) IsTrue(p uint8) bool {
	if (ts.t0 & (1 << p)) == 0 {
		return false
	}
	if (ts.t1 & (1 << p)) == 0 {
		return false
	} else {
		return true
	}
}

// Метод очистить трит
func (ts trs) Clear(p uint8) trs {
	ts.t1 &^= (1 << p)
	ts.t0 &^= (1 << p)
	return ts
}

// Метод вернуть символ трита "-1","0","1"
//func (t trits) SymbNumb(p int8) string {
//	if t.n == 0 {
//		return "0"
//	} else if t.t != 0 {
//		return "1"
//	} else {
//		return "-1"
//	}
//}

// Метод вернуть символ трита "%false","%nil","%true"
//func (t trits) SymbLogic(p int8) string {
//	if t.n == 0 {
//		return "%nil"
//	} else if t.t != 0 {
//		return "%true"
//	} else {
//		return "%false"
//	}
//}

// Метод вернуть символ трита "-","0","+"
//func (t trits) SymbTrit(p int8) string {
//	if t.n == 0 {
//		return "0"
//	} else if t.t != 0 {
//		return "+"
//	} else {
//		return "-"
//	}
//}

// Метод вернуть символ трита "M","N","P"
//func (t trits) SymbChar(p int8) string {
//	if t.n == 0 {
//		return "N"
//	} else if t.t != 0 {
//		return "P"
//	} else {
//		return "M"
//	}
//}

// Метод вернуть символ трита в int8
//func (t trits) ToInt(p int8) int8 {
//	if t.n == 0 {
//		return 0
//	} else if t.t != 0 {
//		return 1
//	} else {
//		return -1
//	}
//}

//
// Операция сдвига тритов
// Версия 1
// Параметр:
// if(d > 0) then "Вправо"
// if(d == 0) then "Нет сдвига"
// if(d < 0) then "Влево"
// Возврат: Троичное число
//
func shift_ts(tr trs, d int8) trs {
	if d > 0 {
		tr.t1 >>= d
		tr.t0 >>= d
	} else if d < 0 {
		tr.t1 <<= -d
		tr.t0 <<= -d
	}
	return tr
}

// Изменить порядок тритов в трайте
//func reverseTryte(input []interface{}) []interface{} {
//	if len(input) == 0 {
//		return input
//	}
//	return append(reverseTryte(input[1:]), input[0])
//}

// Изменить порядок тритов в трайте
//func printTryteInt(input []interface{}) []interface{} {
//	if len(input) == 0 {
//		return input
//	}
//	return append(printTryteInt(input[1:]), trit2int(input[0]))
//}

// Изменить порядок тритов в трайте
//func printTryteSymb(input []interface{}) []interface{} {
//	if len(input) == 0 {
//		return input
//	}
//	return append(printTryteSymb(input[1:]), trit2symb(input[0]))
//}

// ***************************************************************************
// Виртуальный процессор: TRISC-32
// Автор: @oberon87
//
// Links:
// 1) https://habr.com/ru/users/oberon87/
// 2) https://people.inf.ethz.ch/wirth/FPGA-relatedWork/RISC-Arch.pdf
// 3) https://people.inf.ethz.ch/wirth/ProjectOberon/PO.Computer.pdf
// 4) https://habr.com/ru/post/258727/
// ---------------------------------------------------------------------------

// Основные регистры в порядке пульта управления
var (
	K  trs // K(1:9)  код команды (адрес ячейки оперативной памяти)
	F  trs // F(1:5)  индекс регистр
	CR trs // C(1:5)  программный счетчик
	W  trs // W(1:1)  знак троичного числа
	//
	ph1 trs // ph1(1:1) 1 разряд переполнения
	ph2 trs // ph2(1:1) 2 разряд переполнения
	S   trs // S(1:18) аккумулятор
	R   trs // R(1:18) регистр множителя
	MB  trs // MB(1:4) троичное число зоны магнитного барабана
	// Дополнительный
	MR trs // временный регистр для обмена троичным числом
)

// Троичное умножение двух тритов с переносом
func and_trit(a trits, b trits) int8 {
	if a.IsFalse() && b.IsFalse() {
		return -1
	} else if a.IsFalse() && b.IsNil() {
		return -1
	} else if a.IsFalse() && b.IsTrue() {
		return -1
	} else if a.IsNil() && b.IsFalse() {
		return -1
	} else if a.IsNil() && b.IsNil() {
		return 0
	} else if a.IsNil() && b.IsTrue() {
		return -1
	} else if a.IsTrue() && b.IsFalse() {
		return -1
	} else if a.IsTrue() && b.IsNil() {
		return 0
	} else if a.IsTrue() && b.IsTrue() {
		return 1
	}
	return 0
}

func or_trit(a trits, b trits) int8 {
	if a.IsFalse() && b.IsFalse() {
		return -1
	} else if a.IsFalse() && b.IsNil() {
		return 0
	} else if a.IsFalse() && b.IsTrue() {
		return 1
	} else if a.IsNil() && b.IsFalse() {
		return 0
	} else if a.IsNil() && b.IsNil() {
		return 0
	} else if a.IsNil() && b.IsTrue() {
		return 1
	} else if a.IsTrue() && b.IsFalse() {
		return 1
	} else if a.IsTrue() && b.IsNil() {
		return 1
	} else if a.IsTrue() && b.IsTrue() {
		return 1
	}
	return 0
}

func xor_trit(a trits, b trits) int8 {
	if a.IsFalse() && b.IsFalse() {
		return -1
	} else if a.IsFalse() && b.IsNil() {
		return 0
	} else if a.IsFalse() && b.IsTrue() {
		return 1
	} else if a.IsNil() && b.IsFalse() {
		return 0
	} else if a.IsNil() && b.IsNil() {
		return 0
	} else if a.IsNil() && b.IsTrue() {
		return 0
	} else if a.IsTrue() && b.IsFalse() {
		return 1
	} else if a.IsTrue() && b.IsNil() {
		return 0
	} else if a.IsTrue() && b.IsTrue() {
		return -1
	}
	return 0
}

// Очистить троичное число и длину
func clear_full_trs(tr *trs) {
	tr.l = 0
	tr.t0 ^= tr.t0
}

// Очистить троичное число
func clear_trs(tr *trs) {
	tr.t0 ^= tr.t0
}

// Возведение в степень по модулю 3
func pow3(x int8) int32 {
	var i int8
	var r int32 = 1
	for i = 0; i < x; i++ {
		r *= 3
	}
	return r
}

// Преобразование трита в целое число
func trs2int(tr trs, p uint8) int8 {
	if p > TRITSMAX-1 {
		p = TRITSMAX - 1
	}
	if (tr.t0 & (1 << p)) == 0 {
		return 0
	} else {
		if (tr.t1 & (1 << p)) == 0 {
			return -1
		} else {
			return 1
		}
	}
}

// Преобразование целое число в трит
func int2trs(tr trs, p uint8, i int8) trs {
	if p > TRITSMAX-1 {
		p = TRITSMAX - 1
	}
	if i > 0 {
		tr.t1 |= (1 << p)
		tr.t0 |= (1 << p)
		return tr
	}
	if i < 0 {
		tr.t1 &^= (1 << p)
		tr.t0 |= (1 << p)
		return tr
	}
	tr.t1 &^= (1 << p)
	tr.t0 &^= (1 << p)
	return tr
}

//  Преобразовать трит в символ '-','0','+'
func trs2symb(tr trs, p uint8, i int8) string {
	if p > TRITSMAX-1 {
		p = TRITSMAX - 1
	}
	if tr.t0 == 0 {
		return "0"
	} else {
		if tr.t1 == 0 {
			return "-"
		} else {
			return "+"
		}
	}
}

// Операция знак SGN троичного числа
func sgn_trs(x trs) int8 {
	var i int8
	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	for i = int8(x.l) - 1; i >= 0; i -= 1 {
		if ((x.t0 & (1 << i)) > 0) && ((x.t1 & (1 << i)) == 0) {
			return -1
		} else if ((x.t0 & (1 << i)) > 0) && ((x.t1 & (1 << i)) > 0) {
			return 1
		}
	}
	return 0
}

// Операция AND trs
func and_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	if y.l > TRITSMAX-1 {
		y.l = TRITSMAX - 1
	}

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = and_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

// Операция OR trs
func or_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	if y.l > TRITSMAX-1 {
		y.l = TRITSMAX - 1
	}

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = or_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

// Операция XOR trs
func xor_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	if y.l > TRITSMAX-1 {
		y.l = TRITSMAX - 1
	}

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = xor_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

/**
 * Операция сдвига тритов
 * Параметр:
 * if(d > 0) then "Вправо"
 * if(d == 0) then "Нет сдвига"
 * if(d < 0) then "Влево"
 * Возврат: Троичное число
 */
func shift_trs(tr trs, d int8) trs {
	if tr.l > TRITSMAX-1 {
		tr.l = TRITSMAX - 1
	}
	if d > 0 {
		tr.t1 >>= d
		tr.t0 >>= d
	} else if d < 0 {
		tr.t1 <<= -d
		tr.t0 <<= -d
	}
	return tr
}

// Троичное сложение троичных чисел
func add_trs(x trs, y trs) trs {
	var i, j uint8
	var a, b, s, p0, p1 int8
	var r trs

	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	if y.l > TRITSMAX-1 {
		y.l = TRITSMAX - 1
	}
	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	r.l = j
	r.t0 = 0

	p0 = 0
	p1 = 0

	for i = 0; i < j; i++ {

		a = trs2int(x, i)
		b = trs2int(y, i)
		s, p1 = sum_t(a, b, p0)
		r = int2trs(r, i, s)
		p0 = p1

		x.t1 >>= 1
		x.t0 >>= 1
		y.t1 >>= 1
		y.t0 >>= 1
	}
	return r
}

// Троичное сложение троичных чисел
func sub_trs(x trs, y trs) trs {
	var i, j uint8
	var a, b, s, p0, p1 int8
	var r trs

	if x.l > TRITSMAX-1 {
		x.l = TRITSMAX - 1
	}
	if y.l > TRITSMAX-1 {
		y.l = TRITSMAX - 1
	}
	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	r.l = j
	r.t0 = 0

	p0 = 0
	p1 = 0

	for i = 0; i < j; i++ {

		a = trs2int(x, i)
		b = 0 - trs2int(y, i)
		s, p1 = sum_t(a, b, p0)
		r = int2trs(r, i, s)
		p0 = p1

		x.t1 >>= 1
		x.t0 >>= 1
		y.t1 >>= 1
		y.t0 >>= 1
	}
	return r
}

// Сложение  в  S (S)+(A*)=>(S)
// Вычитание в  S (S)-(A*)=>(S)
// Умножение 0 (S)=>(R); (A*)(R)=>(S)
// Умножение + (S)+(A*)(R)=>(S)
// Умножение - (A*)+(S)(R)=>(S)
// Поразрядное умножение (A*)[x](S)=>(S)
// Посылка в R (A*)=>(R)
// Нормализация	Норм.(S)=>(A*); (N)=>(S)
// Аппаратный сброс.

// Очистить память и регистры
// виртуальной машины "Сетунь-1958"
func reset_setun_1958() {
	//
	// clean_fram();	/* Очистить  FRAM */
	// clean_drum();	/* Очистить  DRUM */
	//
	clear_full_trs(&K) /* K(1:9) */
	K.l = 9
	clear_full_trs(&F) /* F(1:5) */
	F.l = 5
	clear_full_trs(&CR) /* K(1:5) */
	CR.l = 5
	clear_full_trs(&W) /* W(1:1) */
	W.l = 1
	//
	clear_full_trs(&ph1) /* ph1(1:1) */
	ph1.l = 1
	clear_full_trs(&ph2) /* ph2(1:1) */
	ph2.l = 1
	clear_full_trs(&S) /* S(1:18) */
	S.l = 18
	clear_full_trs(&R) /* R(1:18) */
	R.l = 18
	clear_full_trs(&MB) /* MB(1:4) */
	MB.l = 4
	//
	clear_full_trs(&MR) /* Временный регистр данных MR(1:9) */
	MR.l = 9
}


