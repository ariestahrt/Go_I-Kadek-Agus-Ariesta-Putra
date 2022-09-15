-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT * FROM USERS WHERE gender='M';

-- b. Tampilkan product dengan id = 3.
SELECT * FROM PRODUCTS WHERE ID=3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM USERS WHERE CREATED_AT > DATE_ADD(CURRENT_DATE, INTERVAL -7 DAY);

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) FROM USERS WHERE gender='F';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM USERS ORDER BY NAME ASC;

-- f. Tampilkan 5 data pada data product
SELECT * FROM PRODUCTS LIMIT 5;