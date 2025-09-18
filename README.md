Run this code in separate terminal for POST method:

curl -i -X POST http://localhost:3000/books/ \
  -H "Content-Type: application/json" \
  -d '{"title":"The Pragmatic Programmer","author":"Andy Hunt & Dave Thomas","year":1999}'

.