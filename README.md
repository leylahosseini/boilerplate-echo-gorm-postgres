# boilerplate-echo-gorm-postgres

curl -X POST localhost:8080/Tasks    -d '{"name":"task1" , "description":"task1" , "completed":true}'   --header "Content-Type: application/json"



curl -X PUT localhost:8080/Tasks/:id -d '{"name":"task1" , "description":"task1" , "completed":true}'   --header "Content-Type: application/json"


curl -X DELETE localhost:8080/Tasks/:id


curl -X GET localhost:8080/Tasks
