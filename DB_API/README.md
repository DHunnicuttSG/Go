## DB_API 
initial test was successful  
Use the following code to test CRUD functionality  

# Create
curl -sS -X POST http://localhost:8080/contacts \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Ada",
    "lastName": "Lovelace",
    "company": "Analytical Engines Ltd",
    "email": "ada@example.com",
    "phone": "+1-555-0100"
  }'

# List
curl -sS "http://localhost:8080/contacts?page=1&pageSize=50"

# Get
curl -sS http://localhost:8080/contacts/1

# Update
curl -sS -X PUT http://localhost:8080/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Ada",
    "lastName": "Byron",
    "company": "Analytical Engines Ltd",
    "email": "ada.byron@example.com",
    "phone": "+1-555-0101"
  }'

# Patch
curl -sS -X PATCH http://localhost:8080/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{"company":"AE Labs"}'

# Delete
curl -sS -X DELETE http://localhost:8080/contacts/1 -i
