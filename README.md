# D-Chat
## Development
### Required environment
- docker
- docker-compose
### Start
1. start up container
    ```
    docker-compose up -d
    ```
1. start up server
    ```
    docker-compose exec server go run main.go
    ```