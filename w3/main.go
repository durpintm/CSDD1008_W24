package main

import "fmt"

func main() {
	// printName("Durpin Thapa Magar")
	// echo();
	// secondEcho("This is from second echo");
	// concatStr("Hello", "World!");
	
		fmt.Println(GenerateRandomNumber())

}

func printName(name string) {
	println(name)
}

func secondEcho(str string){
fmt.Println(str)
}

func thirdEcho(str string) string{
fmt.Println(str)
return str
}

func echo() {
	fmt.Println("This string is printed from echo function")
}

func concatStr(str1 string, str2 string){
	fmt.Println("New string is : ", str1 + " " + str2)
	result := str1 + " for the " + str2
	fmt.Println("New second string is : ", result)
}