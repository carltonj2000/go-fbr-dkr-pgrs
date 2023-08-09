# GoLang, Fiber, Postgres, Docker Example

## Code History

The code in this repository is based on the following videos:
- [Build a REST API from scratch with Go, Docker & Postgres](https://youtu.be/p08c0-99SyU)

## Creation History

After installing docker desktop verify in setting->resources->'file sharing' 
is enable for the present directory.

```bash
docker compose up
docker compose run --service-ports web bash
```

In sided the docker go container started by the above command do the following.

```bash
go mod init github.com/carltonj2000/go-fbr-dkr-pgrs
go get github.com/gofiber/fiber/v2
mkdir cmd
touch touch cmd/main.go
```
