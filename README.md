1. The Application Backend is developed in Golang. For HTTP request gin framework is used and for database ORM Gorm(mysql) is used.
2. Code contains 4 routes :
 1) Register :
    Curl : curl--location 'http://localhost:8080/register' \--header 'Content-Type: application/json' \--data-raw '{"Name": "aditya", "Phone_Number": "9876543210", "Email":
 "abc@xyz.com", "password": "12345678"}'
 2) SpamMarking:
    Curl : curl--location 'http://localhost:8080/mark-as-spam' \--header 'Content-Type: application/json' \--header 'Authorization:
 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU1OTk5MDcsInVzZXJfaWQi
 OjB9.xqNiT1LaXKPskpml6KSV-AbWEEK4drYuxgKoBFci_T4' \--data '{"Number": "98765432183"}'
3) Search using Phone Number:
    Curl : curl--location 'http://localhost:8080/search-by-phone' \--header 'Content-Type: application/json' \--header 'Authorization:
 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MDMyMjQsInVzZXJfaWQ
 iOjB9.Dm0H5nitfAnXLOZ1bL9pd9oMmlTLmJKivgfEhnFNhFY' \--data '{"Phone_Number": "9876543291"}'
 4) Search using name:
    Curl :  curl--location 'http://localhost:8080/search-by-name' \--header 'Content-Type: application/json' \--header 'Authorization:
 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MDMyMjQsInVzZXJfaWQ
 iOjB9.Dm0H5nitfAnXLOZ1bL9pd9oMmlTLmJKivgfEhnFNhFY' \--data '{"Name": "af"}'
3. Have created 3 table for registered users, global data and spam numbers.
