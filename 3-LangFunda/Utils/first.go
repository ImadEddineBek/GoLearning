package Utils

import "fmt"

func First(){
	fmt.Println("hello world")
}
func Second(name string){
	fmt.Println("hello "+name)
}
func Third(){
	var name int
	n , err := fmt.Scanf("%d",&name)
	println(n,err)
	fmt.Println("hello ",name)
}