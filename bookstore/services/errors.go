package services

type BookNotFoundError struct{}

type NoBooksInInventory struct{}

func (e *BookNotFoundError) Error() string {
	return "Book Not Found !"
}

func (e *NoBooksInInventory) Error() string {
	return "No Books in Inventory!"
}
