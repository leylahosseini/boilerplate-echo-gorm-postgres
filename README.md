# Echo Gorm Postgres

# This project is basically made for Golang, which uses Echo Gorm and Postgres ,which includes the CRUD operation.




curl -X POST localhost:8080/tasks    -d '{"name":"task1" , "description":"task1" , "completed":true}'   --header "Content-Type: application/json"



curl -X PUT localhost:8080/tasks/:id -d '{"name":"task1" , "description":"task1" , "completed":true}'   --header "Content-Type: application/json"


curl -X DELETE localhost:8080/tasks/:id


curl -X GET localhost:8080/tasks
