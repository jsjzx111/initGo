package main

import "fmt"

// & 取地址 * 取地址对应的值
func main() {

	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	var num int = 100

	ip = &num
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)

	arr := [3]int{1, 2, 3}

	var arrp [3]*int

	//这边不能用range，否则拿到的地址是一样的
	for i := 0; i < len(arr); i++ {
		arrp[i] = &arr[i]
	}

	fmt.Printf("arrp: %v\n", arrp)

	for _, v := range arrp {
		fmt.Printf("%v \n", *v)
	}

}
