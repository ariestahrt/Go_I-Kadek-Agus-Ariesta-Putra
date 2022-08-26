# Praktikum

Diberikan tugas praktikum sebagai berikut:
https://docs.google.com/document/d/13pHvT0VmRZf4FqZwltrJJpTLZay-U_iW0eU-PCQR3YE/edit#

## Problem 1 - Bilangan Prima (max 50)

Solusi:

Untuk menentukan bilangan N adalah prima atau tidak kita perlu melakukan looping dari 2 s/d sqrt(N) untuk kemudian dicek apakah dari looping tersebut ada bilangan yang dapat membagi N? apabila iya maka dipastikan bilangan tersebut bukan bilangan prima, begitu sebaliknya. Namun perlu kita tentukan beberapa pengecualian yaitu bilangan dibawah 2 bukan merupakan bilangan prima. Sebenarnya cara lebih efektif apabila telah diberikan batasan permasalah dan terdapat banyak bilangan prima yang ingin di cek lebih baik menggunakan algoritma [Sieve of Eratosthenes](https://www.geeksforgeeks.org/sieve-of-eratosthenes/).

Berikut merupakan flowchart dari Problem 1:
![Flowchart Problem 1](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-1.png)

## Problem 2 - Lampu dan Tombol! (max : 50)

Solusi:

Yang mempengaruhi kondisi akhir lampu tersebut adalah banyaknya faktor dari input N tersebut apakah berjumlah ganjil atau genap. Untuk menentukan faktor dari input N tersebut berjumlah ganjil atau genap dapat dilakukan dengan memastikan apakah N memiliki akar kuadrat sempurna? faktor akan berjumlah ganjil jika N memiliki akar kuadrat sempurna sehingga kondisi akhir adalah dalam keadaan menyala, berlaku sebaliknya. sehingga solusi flowchart adalah sebagai berikut:

![Flowchart Problem 2](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-2.png)

## Problem 3 - Instalasi Go

Saya telah menginstall go, dan berikut adalah screenshot dari `go --version`

![Go Version](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-3.png)

## Problem 4 - Hello World!

Kode program:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

Hasil implementasi command

### go run
`go run problem-4.go`
![Go Run](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-4-go-run.png)

### go build
`go build problem-4.go`
![Go Run](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-4-go-build.png)

### go install
`go install problem-4.go`
![Go Run](/4_Introduction%20to%20Algorithm%20and%20Golang/screenshot/Problem-4-go-install.png)
