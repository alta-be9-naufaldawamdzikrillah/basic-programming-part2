package main

import "fmt"

func Palindrome(input string) bool {
	// your code here
	var mengecekBoolean bool = true
	var membalikanKalimatInput string = ""

	for i := len(input) - 1; i >= 0; i-- {			//fungsi untuk membolak-balikan kalimat yang diinput
		membalikanKalimatInput += string(input[i])
	}

	for i := range input {							//jika setelah dibolak-balikan kalimat tersebut hasil sama maka true jika tidak maka false
		if input[i] != membalikanKalimatInput[i] {	//hasil output ditampilkan oleh variable mengecekBoolean
			mengecekBoolean = false
		}
	}
	return mengecekBoolean

}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
