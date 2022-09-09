package main

import(
	"fmt"
	"runtime"
	"sync"
	"strings"
	"bufio"
	"os"
)

func seq(n,p int,mod *int) int {
	return (n/p) + func () int { if *mod > 0 {*mod--; return 1}; return 0 }()
}

func mapChar(str string, mapping *map[rune]int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for _, char := range str {
		if char < 'a' || char > 'z' { continue }
		// fmt.Println(string(char))
		(*mtx).Lock()
		if _, exist := (*mapping)[char]; exist {
			(*mapping)[char]++
		}else{
			(*mapping)[char] = 1
		}
		(*mtx).Unlock()
	}
}

func main() {
	// Testcase here: Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua
	scanner := bufio.NewReader(os.Stdin)
	var input string
	var wg sync.WaitGroup
	var mtx sync.Mutex
	max_process := 1

	// Scan
	fmt.Print("[!] Input string : ")
	input, _ = scanner.ReadString('\n')
	fmt.Print("[!] Max process : ")
	fmt.Scanf("%d", &max_process)
	
	// Setup
	runtime.GOMAXPROCS(max_process)
	n := len(input); p:=max_process
	mod := n%p
	mapping := map[rune]int{}
	input_lower := strings.ToLower(input)

	for i:=0; i<n; {
		prev:=i
		next:=i+seq(n,p,&mod)

		// Tell wg, +1 goroutine bekerja
		wg.Add(1)
		go mapChar(input_lower[prev:next], &mapping, &wg, &mtx)
		fmt.Printf("Counting char from input_lower[%d:%d] = %s\n", prev, next, input_lower[prev:next])
		i=next
	}

	wg.Wait()
	fmt.Println("WG Done~")

	for key,val := range mapping {
		fmt.Printf("%s: %d\n", string(key), val)
	}
}