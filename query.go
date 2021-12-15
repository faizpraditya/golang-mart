package main

const (
	INSERT_PRODUCT = `INSERT INTO m_product
	(id, name, stock, price)
	VALUES
	(:id, :name, :stock, :price)`

	UPDATE_STATUS = `UPDATE m_product SET status= :status WHERE id = :id`

	DETAIL_PRODUCT = `SELECT * FROM m_product WHERE status=1`
)
