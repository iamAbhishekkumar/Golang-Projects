## BookStore

**Design and implement a Bookstore Management System in Golang. The system should allow managing books, customers, and sales transactions.**

## Requirements

### Functional Requirements

1. **Book Management**
   - Add a new book to the inventory.
   - Update details of an existing book.
   - Delete a book from the inventory.
   - Search for a book by title, author, or ISBN.
   - List all available books in the inventory.

2. **Customer Management**
   - Add a new customer.
   - Update customer details.
   - Delete a customer.
   - Search for a customer by name or customer ID.
   - List all registered customers.

3. **Sales Transactions**
   - Create a new sales transaction.
   - Record the sale of one or more books to a customer.
   - Update a sales transaction.
   - Delete a sales transaction.
   - List all sales transactions.
   - Calculate the total sales amount for a given period.

### Non-Functional Requirements

1. **Persistence**
   - Use in-memory data structures (e.g., slices, maps) to store data for books, customers, and transactions.
   - No need for a database; the focus is on learning Golang and OOP principles.

2. **Error Handling**
   - Implement appropriate error handling for all operations (e.g., adding a book with an existing ISBN, updating a non-existent customer).

3. **Modularity**
   - Organize the code into packages/modules to separate concerns (e.g., separate packages for books, customers, transactions).

### Class Design

1. **Book**
   - Attributes: Title, Author, ISBN, Price, Quantity
   - Methods: AddBook, UpdateBook, DeleteBook, SearchBook, ListBooks

2. **Customer**
   - Attributes: Name, CustomerID, Email, Phone
   - Methods: AddCustomer, UpdateCustomer, DeleteCustomer, SearchCustomer, ListCustomers

3. **SalesTransaction**
   - Attributes: TransactionID, CustomerID, BookISBNs, TotalAmount, Date
   - Methods: CreateTransaction, UpdateTransaction, DeleteTransaction, ListTransactions, CalculateTotalSales

### Example Use Case

1. **Adding a Book**
   - Prompt the user to enter book details (title, author, ISBN, price, quantity).
   - Add the book to the inventory and confirm the addition.

2. **Registering a Customer**
   - Prompt the user to enter customer details (name, email, phone).
   - Register the customer and confirm the registration.

3. **Creating a Sales Transaction**
   - Prompt the user to enter the customer ID and the ISBNs of the books being purchased.
   - Record the transaction, update the book quantities, and confirm the transaction.

By implementing this project, you'll get hands-on experience with basic and intermediate Golang concepts, including structs, methods, slices, maps, and basic error handling.
