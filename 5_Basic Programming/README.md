# Summary

## Tipe Variabel

Variabel adalah tempat untuk menyimpan data, data ini biasanya memiliki identitas nama yang kita gunakan untuk memanggil variabel tersebut. Di bahasa golang variabel hanya bisa menampung tipe data yang sama. Adapun tipe data pada golang dapat dibagi sebagai berikut:

### Boolean
Tipe data boolean hanya bernilai True dan False yang merepresentasikan nilai benar/salah. Lebih lanjut pada : https://go.dev/ref/spec#Boolean_types

### Numeric
Tipe data numeric dibagi menjadi 3 yaitu

1. Integer
Integer hanya menyimpan nilai bilangan bulat
2. Float
Float menyimpan bilangan desimal
3. Complex
Menyimpan bilangan kompleks beserta bagian imaginer

Lebih lanjut pada https://go.dev/ref/spec#Numeric_types

### String

String adalah kumpulan karakter dalam sebuah set. Lebih lanjut pada https://go.dev/ref/spec#String_types

Untuk mendeklarasikan sebuah variabel pada golang contohnya adalah sebagai berikut: `var nilaiA int` yaitu sebuah variabel bernama `nilaiA` dengan type `int`. Lebih detailnya:

Dengan cara yang panjang

1. `var <nama_vairabel> <type_data>` : `var name string`
2. `var <nama_vairabel> <type_data> = value` : `var name string = "Ariesta"`
3. `var <daftar_nama_vairabel> <type_data>` : `var name,title string`
4. `var <daftar_nama_vairabel> <type_data>` : `var name,title = "Ariesta","Mr."`

Dengan cara singkat:

1. `<nama_variabel> := <value>` : `name := "Ariesta"`

Selain itu ada pula variabel constant, variabel konstant adalah variabel yang memiliki nilai yang tetap alias tidak berubah-ubah sehingga tidak dapat kita assign kembali. variabel constant dideklarasikan dengan `const <nama_variabel> <type_data> = <value>` atau `const ( [daftar variabel konstant] )`

## Operator

Operasi aritmatika yang ada pada bahasa golang adalah sebagai beriktu:

1. Pertambahan `+`
2. Pengurangan `-`
3. Pembagian `/`
4. Perkalian `*`
5. Modulus `%`
6. Increment `++`
7. Decrement `--`

Adapun terdapat operator pembanding sebagai berikut yang maknanya sudah sangat jelas `==, !=, >, <, >=, <=` . Dan operator logical, bitwise, assignment tidak ada yang spesial di golang terkait operator, kecuali terdapat penggunaan pointer (`&, *`).

## Control Structures ( Branching & Looping )

Berikut merupakan contoh penggunakan statemen `if` dan `if else`

```go
var myAge = 177

if myAge == 5 {
    // ...
}

if (myAge == 5) {
    // ...
}

if myAge >= 17 && myAge <=30 {
    // ...
}

if dadAge := 9; dadAge < myAge {
    // ...
}

if myAge > 5 {
    // ...
} else if myAge > 0 {
    // ...
} else {
    // ...
}

```

Untuk `while` dan `for` tidak ada yang spesial di golang, dari contoh di atas yang membedakan adalah penulisan kondisi dapat dilakukan dengan kurung atau tanpa kurung dan dapat melakukan assign variabel yang memilki lifescope pada statement tersebut. Selanjutnya pada `switch` hal yang istimewa di golang adalah setiap `case` secara default sudah memiliki `break` statement. Berikut contoh dari `switch`

```go
var today int = 2

switch today {
    case 1:
        fmt.Printf("Today is Monday")
    case 2:
        // ...
    default:
        // ...
}
```

Untuk memasuki case di bawahnya, pada golang dapat menambahkan `fallthrough`. Berikut merupakan contohnya, pada saat eksekusi program dijamin memasuki `case 2` apabila `case 1` bernilai `true`

```go
var today int = 1

switch today {
    case 1:
        fmt.Printf("Today is Monday")
        fallthrough
    case 2:
        // ...
    default:
        // ...
}
```