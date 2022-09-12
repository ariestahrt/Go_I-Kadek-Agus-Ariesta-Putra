# Summary

## 1. Database

> Database adalah sekumpulan data yang terorganisir.

> Database adalah sekumpulan data atau informasi yang tersimpan secara sistematis. Basis data ini dibutuhkan ketika mengakses perangkat lunak.

W. Edwards Deming pernah bersabda:
> "Without data, you're just another person with an opinion"

Contohnya adalah:

| Nama | Desa |
| -----|---------|
| Naruto|Konoha|
| Sasuke|Konoha|
| Bee | Kumogakure|
| Gaara | Sunagakure|

## 2. Database Relationship

Terdapat 3 hubungan antar tabel pada database:

1. One to One

Untuk lebih mudah dipahami kita akan gunakan analogi akun twitter.
Analoginya seperti 1 user hanya memiliki 1 photo profil.

2. One to Many

Analoginya seperti 1 user dapat memiliki lebih dari 1 tweet

3. Many to Many

Analoginya seperti:
- 1 user dapat memiliki lebih dari 1 follower, dan 1 user dapat melakukan follow lebih ke 1 user.
- 1 Mahasiswa dapat memiliki banyak mata kuliah, 1 mata kuliah bisa diambil oleh banyak mahasiswa

## 3. SQL Statement

Terdapat 3 jenis perintah sql

1. Data Definition Language (DDL)

Dari namanya sudah jelas fungsinya untuk mendefinisikan. Adapun statement DDL contohnya sebagai berikut:

```sql
# Membuat database
CREATE DATABASE db_name;
# Membuat tabel
CREATE TABLE tb_name (val TYPE, ...)
```

TYPE Data yang digunakan pada database berbeda-beda tergantung pada DBMS yang digunakan.

2. Data Manipulation language (DML).
Perintah SQL yang berhubungan dengan manipulasi data yang ada dalam database. Adapun statement DML contohnya sebagai berikut:
```sql
INSERT ...
SELECT ...
UPDATE ...
DELETE ...
```

3. Data Control Language (DCL)
DCL termasuk perintah seperti GRANT dan REVOKE yang terutama berhubungan dengan hak, izin, dan kontrol lain dari sistem database

Perintah DCL : 

- GRANT: This command gives users access privileges to the database.
- REVOKE: This command withdraws the user’s access privileges given by using the GRANT command.