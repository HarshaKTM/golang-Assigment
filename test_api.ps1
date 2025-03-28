$baseUrl = "http://localhost:5001"

Write-Host "Testing Book API endpoints..." -ForegroundColor Green
Write-Host ""

# 1. List all books
Write-Host "1. GET /books (list all books)" -ForegroundColor Yellow
$books = Invoke-RestMethod -Uri "$baseUrl/books" -Method Get
$books | ConvertTo-Json -Depth 3
Write-Host ""

# 2. Search for fantasy books
Write-Host "2. GET /books/search?q=fantasy (search for fantasy books)" -ForegroundColor Yellow
$searchResults = Invoke-RestMethod -Uri "$baseUrl/books/search?q=fantasy" -Method Get
$searchResults | ConvertTo-Json -Depth 3
Write-Host ""

# 3. Create a new book
Write-Host "3. POST /books (create a new book)" -ForegroundColor Yellow
$newBook = @{
    authorId = "test-author-001"
    publisherId = "test-publisher-001"
    title = "Test Book Created via API"
    publicationDate = "2023-03-15"
    isbn = "1234567890123"
    pages = 150
    genre = "Test"
    description = "This is a test book created via the API for testing purposes."
    price = 12.99
    quantity = 5
}

$newBookJson = $newBook | ConvertTo-Json
$createdBook = Invoke-RestMethod -Uri "$baseUrl/books" -Method Post -Body $newBookJson -ContentType "application/json"
$createdBook | ConvertTo-Json -Depth 3
$bookId = $createdBook.bookId
Write-Host "Created book with ID: $bookId"
Write-Host ""

# 4. Get book by ID
Write-Host "4. GET /books/$bookId (get book by ID)" -ForegroundColor Yellow
$book = Invoke-RestMethod -Uri "$baseUrl/books/$bookId" -Method Get
$book | ConvertTo-Json -Depth 3
Write-Host ""

# 5. Update book
Write-Host "5. PUT /books/$bookId (update book)" -ForegroundColor Yellow
$book.title = "Updated Book Title"
$book.description = "This book has been updated via the API testing script."
$book.price = 19.99
$book.quantity = 10

$updatedBookJson = $book | ConvertTo-Json
$updatedBook = Invoke-RestMethod -Uri "$baseUrl/books/$bookId" -Method Put -Body $updatedBookJson -ContentType "application/json"
$updatedBook | ConvertTo-Json -Depth 3
Write-Host ""

# 6. Search for updated book
Write-Host "6. GET /books/search?q=updated (search for updated book)" -ForegroundColor Yellow
$searchResults = Invoke-RestMethod -Uri "$baseUrl/books/search?q=updated" -Method Get
$searchResults | ConvertTo-Json -Depth 3
Write-Host ""

# 7. Delete book
Write-Host "7. DELETE /books/$bookId (delete book)" -ForegroundColor Yellow
Invoke-RestMethod -Uri "$baseUrl/books/$bookId" -Method Delete
Write-Host "Book deleted successfully."
Write-Host ""

# 8. Verify deletion
Write-Host "8. GET /books/$bookId (verify deletion)" -ForegroundColor Yellow
try {
    $deletedBook = Invoke-RestMethod -Uri "$baseUrl/books/$bookId" -Method Get
    Write-Host "Book still exists!" -ForegroundColor Red
} catch {
    Write-Host "Book was successfully deleted (404 Not Found)" -ForegroundColor Green
}
Write-Host ""

Write-Host "API testing completed!" -ForegroundColor Green 