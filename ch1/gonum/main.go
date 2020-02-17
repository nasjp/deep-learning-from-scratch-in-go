package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func matPrint(x mat.Matrix) {
	fa := mat.Formatted(x, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	array()
	calculate()
	dimensionArray()
	broadcast()
	accessToElem()
}

func array() {
	// 配列の生成
	x := mat.NewVecDense(3, []float64{1.0, 2.0, 3.0})
	matPrint(x)
	/*
	 ⎡1⎤
	 ⎢2⎥
	 ⎣3⎦
	*/
	fmt.Printf("x type: %T\n", x)
	// x type: *mat.VecDense
}

func calculate() {
	// gonum の算術計算
	x := mat.NewVecDense(3, []float64{1.0, 2.0, 3.0})
	matPrint(x)
	/*
		⎡1⎤
		⎢2⎥
		⎣3⎦
	*/

	y := mat.NewVecDense(3, []float64{2.0, 4.0, 6.0})
	matPrint(y)
	/*
		⎡2⎤
		⎢4⎥
		⎣6⎦
	*/

	// 空行列の生成
	a := mat.NewVecDense(3, nil)
	matPrint(a)
	/*
		⎡0⎤
		⎢0⎥
		⎣0⎦
	*/

	// 和
	a.AddVec(x, y)
	matPrint(a)
	/*
		⎡3⎤
		⎢6⎥
		⎣9⎦
	*/
	a.Reset()

	// 差
	a.SubVec(x, y)
	matPrint(a)
	/*
		⎡-1⎤
		⎢-2⎥
		⎣-3⎦
	*/
	a.Reset()

	// 積
	a.MulElemVec(x, y)
	matPrint(a)
	/*
		⎡ 2⎤
		⎢ 8⎥
		⎣18⎦
	*/
	a.Reset()

	// 割
	a.DivElemVec(x, y)
	matPrint(a)
	/*
		⎡0.5⎤
		⎢0.5⎥
		⎣0.5⎦
	*/
}

func dimensionArray() {
	// gonum のN次元配列
	a := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	matPrint(a)

	// 行列の形を取得
	r, c := a.Dims()
	fmt.Printf("(%d, %d)\n", r, c)
	// (2, 2)
	b := mat.NewDense(2, 2, []float64{3, 0, 0, 6})

	x := &mat.Dense{}

	// 和
	x.Add(a, b)
	matPrint(x)
	/*
	   ⎡4   2⎤
	   ⎣3  10⎦
	*/
	x.Reset()

	// 積(要素ごと)
	x.MulElem(a, b)
	matPrint(x)
	/*
		⎡3   0⎤
		⎣0  24⎦
	*/
	x.Reset()

	// スカラ値
	x.Scale(10, a)
	matPrint(x)
	/*
		⎡10  20⎤
		⎣30  40⎦
	*/
	x.Reset()
}

func broadcast() {
	// ブロードキャストはしてくれない(多分)
	// だって静的型付け言語だから！
	a := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	matPrint(a)
	/*
		⎡1  2⎤
		⎣3  4⎦
	*/
	b := mat.NewDense(1, 2, []float64{10, 20})
	matPrint(b)
	// [10  20]
	// こうするしかない
	bb := &mat.Dense{}
	bb.Stack(b, b)
	matPrint(bb)
	/*
		⎡10  20⎤
		⎣10  20⎦
	*/

	x := &mat.Dense{}
	x.MulElem(a, bb)
	matPrint(x)
	/*
		⎡10  40⎤
		⎣30  80⎦
	*/
}

func accessToElem() {
	x := mat.NewDense(3, 2, []float64{
		51, 55,
		14, 19,
		0, 4,
	})

	matPrint(x)
	/*
		⎡51⎤
		⎣55⎦
	*/

	matPrint(x.RowView(0))
	fmt.Println(x.At(0, 1))
	// 55

	// 1次元配列への変換は実装する必要がある
	v := denseFlatten(x)
	matPrint(v)
	/*
	   ⎡51⎤
	   ⎢55⎥
	   ⎢14⎥
	   ⎢19⎥
	   ⎢ 0⎥
	   ⎣ 4⎦
	*/

	matPrint(mat.NewVecDense(3, []float64{v.AtVec(0), v.AtVec(2), v.AtVec(4)}))
	/*
	  ⎡51⎤
	  ⎢14⎥
	  ⎣ 0⎦
	*/

	// リスト内包表記的なことはできないのでfunction を定義
	matPrint(over(v, 15))
	/*
		⎡51⎤
		⎢55⎥
		⎣19⎦
	*/
}

func denseFlatten(x *mat.Dense) *mat.VecDense {
	r, c := x.Dims()
	a := mat.NewVecDense(r*c, nil)
	for ri := 0; ri < r; ri++ {
		for ci := 0; ci < c; ci++ {
			a.SetVec(ri*c+ci, x.At(ri, ci))
		}
	}
	return a
}

func over(x *mat.VecDense, n float64) *mat.VecDense {
	data := make([]float64, 0)
	for i := 0; i < x.Len(); i++ {
		e := x.AtVec(i)
		if n >= e {
			continue
		}
		data = append(data, e)
	}
	return mat.NewVecDense(len(data), data)
}
