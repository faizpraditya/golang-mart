package main

const (
	INSERT_PRODUCT = `INSERT INTO m_product
	(id, name, stock, price)
	VALUES
	(:id, :name, :stock, :price)`

	UPDATE_STATUS = `UPDATE m_product SET status= :status WHERE id = :id`

	GET_PRODUCT = `SELECT * FROM m_product WHERE id = $1`

	INSERT_DETAIL_TRANSACTION = `INSERT INTO purchase_detail
	(id, purchase_id, product_id, price, qty)
	VALUES
	(:id, :purchase_id, :product_id, :price, :qty)`

	INSERT_TRANSACTION = `INSERT INTO purchase
	(id, customer_id, purchase_date)
	VALUES
	(:id, :customer_id, CURRENT_TIMESTAMP)`

	UPDATE_STOCK_DECREASE = `UPDATE m_product
	SET stock = (stock - :qty)
	WHERE id = :id`

	GET_DETAIL_TRANSACTION = `SELECT pd.*, pur.customer_id, pur.purchase_date
	FROM purchase_detail pd
	JOIN purchase pur
	ON pd.purchase_id = pur.id
	WHERE pd.id = $1`
)
