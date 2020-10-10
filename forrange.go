package main

import "fmt"

func learnForRange() {
	arr := [3]int{0, 1, 2}

	slice := []int{0, 1, 2, 3, 4}

	fmt.Printf("原数组：%T, %v \n", arr, arr)
	fmt.Println(`for range 是复制对象进行操作，数组是值类型，在range内的i，v一定是数组修改前的值
当i==0时修改后面数组的值，并不影响i==1、i==2时的v
所以最后原数组又被修改`)
	for i, v := range arr {
		if i == 0 {
			arr[1] = 99
			arr[2] = 999
			fmt.Printf("当i==0时修改数组值：%v \n", arr)
		}

		arr[i] = v + 100
	}

	fmt.Printf("结果数组：%v \n", arr)

	fmt.Printf("\n\n原切片：%T, %v \n", slice, slice)
	fmt.Println(`for range 是复制对象进行操作，切片是引用类型，
当i==0时修改切片后面的值，直接影响i==1、i==2时的v`)
	for i, v := range slice {
		if i == 0 {
			slice = slice[:3]
			slice[1] = 99
			slice[2] = 999
			fmt.Printf("当i==0时修改切片的值，并将切片切剩三个值：%v \n", slice)
		}

		if i < len(slice) {
			slice[i] = v + 100
		} else {
			fmt.Printf("当i==%d时，由于切片已经被切了，所以再用slice[%d]就会报错，但依然可以输出i、v：%d %d\n", i, i, i, v)
		}
	}

	fmt.Printf("结果切片：%v \n", slice)

}
