/*Дано натуральное число n.
а) Выяснить, входить ли цифра 3 в запись числа 2
n .
б) Поменять порядок цифр числа n на обратный.*/

package main

import (
	"fmt"
)

func reverseOrder(number uint32) string{
	str := fmt.Sprint(number)
	var reverse string
	for i := len(str)-1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}
func searchThreeInN2(number uint32) bool{
		n := number*number
		for n>=0 {
			if n%10 == 3 {
				return true
			} else {
				n/=10
			}
		}
		return false
}

func main(){
	var n uint32
	fmt.Println("Enter n: ")
	fmt.Scanf("%d \n", &n)
	fmt.Println(searchThreeInN2(n))
	fmt.Println("n in reverse order is: ")
	fmt.Println(reverseOrder(n))
}
