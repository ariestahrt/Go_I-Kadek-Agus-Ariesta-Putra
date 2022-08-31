# Summary

## 1. Time Complexity

Kompleksitas waktu adalah seberapa lama waktu yang diperlukan untuk menjalankan suatu algoritma, perhitungan ini dilakukan dengan pendekatan menghitung banyaknya operasi yang dominan, seperti penjumlahan pengurangan perkalian, assignment. Time complexsity yang dipelajari disini menggunakan notasi big-O

## 2. Perbandingan antar beragam kompleksitas waktu

### 1. Linear time - O(n)
Adalah kompleksitas waktu yang berarti operasi dominan yang terjadi adalah sebanyak N kali. Berikut merupakan contohnya:
```go

a := 0
for i := 0; i<n; i++ {
    a += 1
}

```

### 2. Logarithmic time - 0(log n)
Nilai dari n selalu berkurang setengah untuk setiap iterasi, contohnya:

```go

func logarithmic(n int) int {
    result := 0
    for n > 1 {
        n /= 2
        result += 1
    }
    return result
}

```

### 3. Quadratic time
Biasanya terjadi pada nested-for statement, contohnya:
```go
for i := 0; i<n; i++ {
    for j := 0; j<n; j++ {
        a += 1
    }
}
```

Selain yang disebutkan diatas, ada pula kompleksitas exponensial O(2^n) dan factorial O(n!)

## 3. Space complexity

Space Complexity adalah seberapa besar memori yang kita gunakan untuk menjalankan suatu algoritma.