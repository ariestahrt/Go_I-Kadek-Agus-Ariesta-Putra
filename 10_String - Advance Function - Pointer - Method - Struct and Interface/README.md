# Summary

## 1. String & Function


### Strings
Fungsi string terdapat dalam package "strings", adapun fungsi-fungsi yang sering digunakan adalah:

1. len()
Untuk menghitung panjang string
2. strings.Contains(a, b string)
Untuk menentukan apakah string b terdapat pada string a
3. strings.Replace()
Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu.

### Function

Terdapat variadic function berfungsi untuk membuat slice sementara untuk memasukkannya ke dalam fungsi, digunakan ketika total parameter masukan tidak diketahui banyaknya, dan mempermudah pembacaan kode.

Terdapat pula anonymous function yaitu fungsi yang tidak punya nama. Contohnya:

```go
funct() { 
    //insert code here
}()
```

Terdapat pula closure yang berfungsi agar suatu variabel tidak dapat diakses dari luar fungsi, contoh penerapannya adalah penggunaan increment otomatis ketika fungsi dipanggil, namun increment tersebut tidak dapat kita modifikasi di luar fungsi.

Dan terdapat defer function, yang hanya dieksekusi ketika parrent function memberikan nilai kembalian dan berjalan sebagai `stack`.

## 2. Pointer, Struct & Interface

### Pointer

Pointer adalah variabel yang menyimpan alamat memori dari variabel lainnya. Digunakan untuk melakukan modifikasi nilai data pada variabel tertarget. Terdapat dua operator penting pada pointer:

1. Deferencing Operator : `*`
Digunakan nutuk mendeklarasi variabel pointer atau mengakses nilai yang disimpan pada alamat
2. Referencing Operator : `&`
Mengakses alamat dari variabel atau mengakses alamat dari variabel ke pointer

### Struct

Adalah user-defined type yang mengandung koleksi dari variabel atau fungsi(metode). Deklarasinya sebagai berikut:

```go
type struct_name struct {
    field <data_type>
    field2 <data_type>
}
```

Dalam struct dapat terdapat method yang merupakan fungsi yang ditambahkan pada struct. Berikut cara deklarasinya:

```go
func (receiver StructType) MethodName(...params) (returnTypes) {
    // code here
}
```
### Interface

Interface adalah kolehsi dari method signatures yang dapat diimplementasikan objek. Interface mendefinisikan perilaku dari suatu objek. Deklarasinya:
```go
type interface_name interface {
    method1 <return_type>
    method2 <return_type>
}
```

## 3. Package & Error Handling

### Package 
Adalah kumpulan fungsi dan data, contohnya "fmt"

### Error handling
Dalam golang error handling dilakukan dengan panic dan recover

Panic adalah kondisi ketika go runctime mendeteksi kesalahan. Recover adalah kemampuan untuk menangani panic error, bisa menambahkan fungsi anonim atau custom fungsion yang dipanggil menggunakan `defer` di dalam sebuah method.