-- a. Delete data pada tabel product dengan id 1.
-- Hapus dulu yang di transaction details dan yang di product_details
DELETE FROM TRANSACTION_DETAILS WHERE PRODUCT_ID=1;
DELETE FROM PRODUCT_DETAILS WHERE ID=1;
-- Baru bisa delete yang di product
DELETE FROM PRODUCTS WHERE ID=1;

-- b. Delete pada pada tabel product dengan product type id 1.
-- Hapus dulu yang di transaction details dan yang di product_details
DELETE FROM TRANSACTION_DETAILS WHERE PRODUCT_ID IN (SELECT ID FROM PRODUCTS WHERE PRODUCT_TYPE_ID = 1)
DELETE FROM PRODUCT_DETAILS WHERE ID IN (SELECT ID FROM PRODUCTS WHERE PRODUCT_TYPE_ID = 1)
-- Baru bisa delete yang di product
DELETE FROM PRODUCTS WHERE PRODUCT_TYPE_ID = 1;