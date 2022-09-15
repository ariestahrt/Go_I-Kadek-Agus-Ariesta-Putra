-- 1. Insert 5 operators pada table operators.
INSERT INTO OPERATORS (NAME) VALUES
	('Telkomsel'),
	('Indosat'),
	('XL'),
	('Smartfren'),
	('ByU');

-- 2. Insert 3 product type.
INSERT INTO PRODUCT_TYPES (NAME) VALUES
	('Internet'),
	('Roaming'),
	('SMS');

-- 3. Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO PRODUCTS ( PRODUCT_TYPE_ID, OPERATOR_ID, CODE, NAME, PRICE, STATUS)
	VALUES
	(1,3, 'INET001', 'Kuota 5 GB', 10000, 1),
	(1,3, 'INET002', 'Kuota 15 GB', 20000, 1),
	(1,3, 'INET002', 'Kuota 20 GB', 30000, 1);
	
-- 4. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO PRODUCTS ( PRODUCT_TYPE_ID, OPERATOR_ID, CODE, NAME, PRICE, STATUS)
	VALUES
	(2,1, 'ROAMING001', 'Sinyal di Arab', 100000, 1),
	(2,1, 'ROAMING002', 'Sinyal di Turki', 50000, 1),
	(2,1, 'ROAMING003', 'Sinyal di Malaysia', 40000, 1);

-- 5. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO PRODUCTS ( PRODUCT_TYPE_ID, OPERATOR_ID, CODE, NAME, PRICE, STATUS)
	VALUES
	(3,4, 'SMS001', '300 SMS ke semua operator', 10000, 1),
	(3,4, 'SMS002', 'Gratis SMS ke sesama operator', 11000, 1),
	(3,4, 'SMS003', '100 SMS ke semua operator', 12000, 1);

-- 6. Insert product description pada setiap product.
INSERT INTO PRODUCT_DESCRIPTIONS ( ID, DESCRIPTION)
	VALUES
	(1, "Dapat kuota 5GB"),
	(2, "Dapat kuota 15GB"),
	(3, "Dapat kuota 20GB"),
	(4, "Dapat sinyal buat smsan, internan di Arab"),
	(5, "Dapat sinyal buat smsan, internan di Turki"),
	(6, "Dapat sinyal buat smsan, internan di Malaysia"),
	(7, "Gratis 300 SMS ke semua operator"),
	(8, "Gratis SMS sepuasnya ke sesama operator"),
	(9, "Gratis 100 SMS ke semua operator");

-- 7. Insert 3 payment methods.
INSERT INTO PAYMENT_METHODS ( NAME, STATUS )
	VALUES
	("SHOPEEPAY", 1),
	("GO-PAY", 1),
	("DANA", 1);

-- 8. Insert 5 user pada tabel user.
INSERT INTO USERS ( NAME, STATUS, DOB, GENDER)
	VALUES
	("Kadek", 1, "2000-11-12", 'M'),
	("Agus", 1, "2000-08-12", 'M'),
	("Ariesta", 1, "2000-04-21", 'M'),
	("Putri", 1, "2001-07-16", 'F'),
	("Sabine", 1, "2002-04-12", 'F');

-- 9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO TRANSACTIONS ( USER_ID, PAYMENT_METHOD_ID, STATUS )
	VALUES
	(1, 1, 1),
	(1, 2, 1),
	(1, 3, 1),
	(2, 1, 1),
	(2, 2, 1),
	(2, 3, 1),
	(3, 1, 1),
	(3, 2, 1),
	(3, 3, 1),
	(4, 1, 1),
	(4, 2, 1),
	(4, 3, 1),
	(5, 1, 1),
	(5, 2, 1),
	(5, 3, 1);

-- 10. Insert 3 product di masing-masing transaksi.
INSERT INTO TRANSACTION_DETAILS (TRANSACTION_ID, PRODUCT_ID, STATUS, QTY)
    VALUES
    (1, 1, 1, 4),
    (1, 2, 1, 2),
    (1, 3, 1, 1),
    (2, 4, 1, 5),
    (2, 5, 1, 1),
    (2, 6, 1, 3),
    (3, 7, 1, 8),
    (3, 8, 1, 5),
    (3, 9, 1, 2),
    (4, 1, 1, 4),
    (4, 2, 1, 8),
    (4, 3, 1, 1),
    (5, 4, 1, 3),
    (5, 5, 1, 5),
    (5, 6, 1, 6),
    (6, 7, 1, 2),
    (6, 8, 1, 3),
    (6, 9, 1, 5),
    (7, 1, 1, 6),
    (7, 2, 1, 7),
    (7, 3, 1, 8),
    (8, 4, 1, 11),
    (8, 5, 1, 1),
    (8, 6, 1, 2),
    (9, 7, 1, 3),
    (9, 8, 1, 4),
    (9, 9, 1, 5),
    (10, 1, 1, 7),
    (10, 2, 1, 1),
    (10, 3, 1, 2),
    (11, 4, 1, 3),
    (11, 5, 1, 2),
    (11, 6, 1, 3),
    (12, 7, 1, 5),
    (12, 8, 1, 3),
    (12, 9, 1, 2),
    (13, 1, 1, 5),
    (13, 3, 1, 3),
    (13, 5, 1, 2),
    (14, 1, 1, 4),
    (14, 3, 1, 12),
    (14, 5, 1, 2),
    (15, 1, 1, 4),
    (15, 3, 1, 31),
    (15, 5, 1, 2);