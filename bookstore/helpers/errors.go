package helpers

type BookNotFoundError struct{}

type NoBooksInInventory struct{}

type NotEnoughBooksQty struct{}

func (e *BookNotFoundError) Error() string {
	return "Book Not Found !"
}

func (e *NoBooksInInventory) Error() string {
	return "No Books in Inventory!"
}

func (e *NotEnoughBooksQty) Error() string {
	return "Not enough books in inventory!"
}
