-- a. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE PRODUCTS SET NAME="product dummy" WHERE ID=1;

-- b. Update qty = 3 pada transaction detail dengan product id 1.
UPDATE TRANSACTION_DETAILS SET QTY=3 WHERE PRODUCT_ID=1;
