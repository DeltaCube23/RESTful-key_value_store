# RESTful-key_value_store

Use curl to access the key value store and perform any of the CRUD functions<br>

POST request : ```curl -d "name=<key>&address=<value>" -X POST http://localhost:8080/create```<br>
GET request : ```curl -X GET  http://localhost:8080/read/?name=<key>```<br>
PUT request : ```curl -d "name=<key>&address=<value>" -X PUT http://localhost:8080/update```<br>
DELETE request : ```curl -X DELETE  http://localhost:8080/delete/<key>```<br>
