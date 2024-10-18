package utils

// DB queries
const (
	DB_CREATE_PRODUCT = "INSERT INTO products (name, title, description, price, quantity) VALUES (?,?,?,?)"
	DB_GET_PRODUCT    = "SELECT * FROM products WHERE id = ?"
	DB_DELETE_PRODUCT = "DELETE FROM products WHERE id = ?"
)
