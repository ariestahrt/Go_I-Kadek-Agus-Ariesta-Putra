# Summary

## 1. Recursive

Rekursive atau rekursi adalah kondisi dimana sebuah fungsi memiliki alur penyelesaian dengan memanggil dirinya sendiri. Rekursi diperlukan untuk memudahkan problem solving dan mempersingkat penulisan kode. Terdapat dua hal utama pada pendekatan dengan metode rekursi

1. Base case
Adalah kasus termudah dalam masalah
2. Recurrence relations
Bagaimana hubungan kondisi saat ini apakah dapat dipecah menjadi masalah yang lebih kecil?

Salah satu contoh dari penggunaan rekursi adalah fungsi faktorial sebagai berikut:

```go

func factorial(n int) int {
    if n == 1 { return 1 }
    return n * factorial(n-1)
}

```

## 2. Number Theory

Teori bilangan adalah cabang matematika yang mempelajari bilangan bulat (menurut ppt). Contohnya adalah teori bilangan prima, FPB dan KPK, faktorial, faktor prima, dan lain-lain.

Contohnya adalah penyelesaian FPB (Faktor Persekutuan Besar) dengan menggunakan Euclid's Algorithm.
```go

func gcd(a,b int) int {
    if b == 0 { return a }
    return gcd(b, a%b)
}

```

## 3. Searching & Sorting

### Searching

Searching adalah proses mencari nilai dari suatu container. Terdapat beberapa algoritma search yang populer

1. Linear Search - O(n)

```go

func linierSearch(elements []int, search int) int {
	for idx, val := range elements {
		if val == search { return idx } 
	}
    return -1
}

```

2. Binary Search - O(log n)

Hanya bisa dilakukan pada data yang telah urut. Ref : https://www.geeksforgeeks.org/binary-search/

### Sorting

Sorting adalah proses mengurutkan data (besar ke kecil / kecil ke besar). Terdapat beberapa algoritma sorting yang populer

1. Selection sort - O(n^2)
Ref: https://www.geeksforgeeks.org/selection-sort/
2. Counting sort - O(n+k)
Ref: https://www.geeksforgeeks.org/counting-sort/
3. Merge sort O(log n)
Ref: https://www.geeksforgeeks.org/merge-sort/
