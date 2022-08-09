#How to run
1) Run main/producer-api.go
2) Run main/main-consumer.go
3) Send curl with as below example to the producer-api
- Example of curl:
`curl --location --request POST '0.0.0.0:3000/api/v1/comment' \
--header 'Content-Type: application/json' \
--data-raw '{
"text": "testing"
}'`