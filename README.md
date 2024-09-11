# go-template
## Install swag by using
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
## Generate Swagger documents integrated with Apifox
#### Run the command in your project directory
```bash
swag init -g /cmd/main.go -exclude model
