# Summary

## 1. Perbedaan Sequential, Parallel, and Concurrent

### Sequential Program
Pada program sekuensial sebelum task baru dimulai task sebelumnya harus berada pada kondisi finish.

### Parallel Program
Pada program pararel, banyak task dapat dieksekusi serentak

### Concurrent Program
Pada program concurrent, banyak tugas dapat dijalankan secara independen dan mungkin muncul secara bersamaan

## 2. Go Routines

Goroutine merupakan salah satu bagian penting dalam concurrent programming di Go Lang. Salah satu yang membuat goroutine sangat istimewa adalah eksekusinya dijalankan pada multi core processor. Kita juga dapat menentukan banyak core yang aktif, semakin banyak maka akan semakin cepat eksekusinya.

## 3. Channels, Select dan Race Condition

### Channel
Channel adalah saluran komunkasi yang digunakan goroutines untuk melakukan komunikasi satu dengan yang lainnya.

### Select

Select memudahkan kontrol data komunikasi melalui satu atau banyak channel

### Race Condition

Race condition adalah kondisi dimana dua buah thread mengakses memori yang sama di saat yang bersamaan sebagai contoh satu thread mengakses memori untuk ditulis dan satunya lagi mengakses untuk mendapatkan datanya. Race condition terjadi karena tidak sinkronnya akses kepada shared memory.
