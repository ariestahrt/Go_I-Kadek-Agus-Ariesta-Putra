package main
import "fmt"

func caesar(offset int, input string) string {
	str := []rune(input)
	for i, _ := range str {
		str[i] = 'a' + (str[i] % 'a' + rune(offset % 26)) % 26
	}
	return string(str)
}

type Student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	str := []rune((*s).name)
	for i, _ := range str {
		str[i] = 'a' + (25 - str[i] % 'a')
	}
	(*s).nameEncode = string(str)
	return string(str)
}

func (s *Student) Decode() string {
	str := []rune((*s).nameEncode)
	for i, _ := range str {
		str[i] = 'a' + (25 - str[i] % 'a')
	}
	return string(str)
}

func main() {
	var menu int
	var a = Student{}

	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
       fmt.Print("\nInput Student's Name : ")
       fmt.Scan(&a.name)
       fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
   } else if menu == 2 {
       fmt.Print("\nInput Student's Encode Name : ")
       fmt.Scan(&a.nameEncode)
       fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())
   } else {
       fmt.Println("Wrong input name menu")
   }
}
