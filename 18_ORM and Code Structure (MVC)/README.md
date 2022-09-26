# SUMMARY

## ORM
ORM (Object Relation Mapping) merupakan teknik yang merubah suatu table menjadi sebuah object yang nantinya mudah untuk digunakan. Object yang dibuat memiliki property yang sama dengan field — field yang ada pada table tersebut. ORM memungkinkan kita melakukan query dan memanipulasi data di database menggunakan object oriented.

Biasanya developer melakukan kesalahan ketika menuliskan query database. Karena memang terdapat aturan tersendiri dalam penulisan syntaxnya. Hal ini membuat developer kesulitan dan bahkan memerlukan banyak waktu untuk hanya melakukan query. Maka dari itu diperlukan ORM untuk mempermudah mengakses database tanpa melakukan query sama sekali.

Keuntungan ORM adalah:

- Less repetitive query
- Secara otomatis melakukan fetch data ke object siap pakai
- Cara mudah untuk screening data sebelum disimpan ke database
- Beberapa memiliki fitur cache query

Kekurangan ORM

- Memuat relationship data yang tidak diperlukan
- Query yang kompleks dan beberapa query spesifik belum dapat didukung

## Database Migration

Cara memperbarui versi basis data untuk mematuhi perubahan versi aplikasi. Perubahan dapat di upgrade ke versi yang baru atau di rollback ke versio sebelumnya.

Kenapa menggunakan DB Migration?

- Update database simple
- Rollback database simple
- Terdapat track changes pada struktur database
- Riwayat struktur basis data ditulis pada kode
- Selalu compatible dengan perubahan versi aplikasi

## MVC ( Model View Controller )

MVC atau Model View Controller adalah sebuah pola desain arsitektur dalam sistem pengembangan website yang terdiri dari tiga bagian, yaitu: 

- **Model**, bagian yang mengelola dan berhubungan langsung dengan database;
- **View**, bagian yang akan menyajikan tampilan informasi kepada pengguna;
- **Controller**, bagian yang menghubungkan model dan view dalam setiap proses request dari user dan berisi proses bisnis dari aplikasi kita.