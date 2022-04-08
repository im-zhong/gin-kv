# start
`go run main.go`

# get
`curl http://localhost:8080/get`

# get by key
`curl http://localhost:8080/get/somekey`

# put kv
`curl http://localhost:8080/put --include --header "Content-Type: application/json" --request "POST" --data '{"key": "somekey", "value": "somevalue"}'`

# append kv
`curl http://localhost:8080/append --include --header "Content-Type: application/json" --request "POST" --data '{"key": "somekey", "value": "somevalue"}'`