@echo off
setlocal enabledelayedexpansion
set BASE_URL=http://localhost:5001

echo Testing API endpoints...
echo.

echo 1. GET /books (list all books)
curl -s %BASE_URL%/books
echo.
echo.

echo 2. Testing search for "fantasy" books
curl -s "%BASE_URL%/books/search?q=fantasy"
echo.
echo.

:: For simplicity, we'll use the first book returned by the API
curl -s %BASE_URL%/books > books_temp.json
for /f "tokens=2 delims=:," %%a in ('findstr "bookId" books_temp.json') do (
    set BOOK_ID=%%a
    set BOOK_ID=!BOOK_ID:"=!
    set BOOK_ID=!BOOK_ID:}=!
    set BOOK_ID=!BOOK_ID: =!
    goto :found_id
)
:found_id

echo 3. GET /books/{id} (get book by ID)
echo Using ID: !BOOK_ID!
curl -s %BASE_URL%/books/!BOOK_ID!
echo.
echo.

echo 4. Testing PUT /books/{id} (update book)
curl -s -X PUT %BASE_URL%/books/!BOOK_ID! -H "Content-Type: application/json" -d "{\"title\":\"Updated Book Title\",\"authorId\":\"e0d91f68-a183-477d-8aa4-1f44ccc78a70\",\"publisherId\":\"2f7b19e9-b268-4440-a15b-bed8177ed607\",\"publicationDate\":\"2023-01-01\",\"isbn\":\"1234567890\",\"pages\":100,\"genre\":\"Test\",\"description\":\"This book has been updated for testing purposes\",\"price\":19.99,\"quantity\":10}"
echo.
echo.

echo 5. Testing search for the updated book
curl -s "%BASE_URL%/books/search?q=updated"
echo.
echo.

echo 6. Create a new book for deletion test
curl -s -X POST %BASE_URL%/books -H "Content-Type: application/json" -d "{\"title\":\"Book To Delete\",\"authorId\":\"e0d91f68-a183-477d-8aa4-1f44ccc78a70\",\"publisherId\":\"2f7b19e9-b268-4440-a15b-bed8177ed607\",\"publicationDate\":\"2023-01-01\",\"isbn\":\"9999999999\",\"pages\":100,\"genre\":\"Test\",\"description\":\"This book will be deleted\",\"price\":9.99,\"quantity\":1}" > new_book.json

for /f "tokens=2 delims=:," %%a in ('findstr "bookId" new_book.json') do (
    set DELETE_ID=%%a
    set DELETE_ID=!DELETE_ID:"=!
    set DELETE_ID=!DELETE_ID:}=!
    set DELETE_ID=!DELETE_ID: =!
    goto :found_delete_id
)
:found_delete_id

echo 7. Testing DELETE /books/{id}
echo Deleting book with ID: !DELETE_ID!
curl -s -X DELETE %BASE_URL%/books/!DELETE_ID!
echo.
echo.

echo 8. Verifying deletion by trying to retrieve the deleted book
curl -s %BASE_URL%/books/!DELETE_ID!
echo.
echo.

:: Clean up temp files
del books_temp.json
del new_book.json

echo API testing completed!
endlocal 