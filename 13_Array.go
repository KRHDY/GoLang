package main

import "fmt"

func main()  {

	// ******************************
	// ************ 排列 ************
	// ******************************

	array_01();
		// array_01 呼出し

	array_02();
		// array_02(); 呼び出し

}

func array_01()  {
	var x[5]int;
		// ｘという排列（５つ）を作る

	x[4] = 100
		// x排列の５番目に１００を入れる

	fmt.Println(x)
		// ｘ排列を出力

}

func array_02()  {
	var x [5]float64
		// ｘという排列（５つ）を作る

	x[0] = 98
		// xという排列の一番目に93を入れる

	x[1] = 93
		// xという排列の二番目に93を入れる

	x[2] = 77
		// xという排列の三番目に77を入れる

	x[3] = 82
		// xという排列の四番目に82を入れる

	x[4] = 83
		// xという排列の五番目に83を入れる

	var total float64 = 0
		// トタルという変数を０で作る

	for i := 0; i < 5; i++ {
		total += x[i]
			// トタルはトタルプラスｘという排列のI番目の値をいれる

	}
	fmt.Println(total / 5)
		// トタルに５を割る
	
	
	for i := 0; i < len(x); i++ {
 		total += x[i]
			// 変数トタルにｘ排列の「Ｉ」に入れる
		
	}
	fmt.Println(total / len(x))
		// トタルをx排列の数に割る

}
