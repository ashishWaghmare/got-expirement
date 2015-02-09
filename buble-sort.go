package main

import "fmt"

func main(){
	list :=[]int{2,31,2,4,2,4,5}
  fmt.Println("unsorted",list)
  optsort(list)
  fmt.Println("sorted",list)
}

func bublesort(a[]int) {
	for i :=0;i <= len(a)-1;i++{
		for j :=i; j <len(a) -1;j++{
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}
		}
	}
}

func optsort(a[]int){
	hasChanged := true 
	itemCount :=len(a)-1
	for hasChanged {
		hasChanged = false 
   	for index := 0;index < itemCount; index++ {
			if a[index] > a[index+1] {
				a[index],a[index+1] = a[index+1],a[index]
				hasChanged = true
			}
		}
  }
}

