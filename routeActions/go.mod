module github.com/Labaster/go-app/routeActions

go 1.23.3

require (
	github.com/Labaster/go-app/dbConf v0.0.0-00010101000000-000000000000
	github.com/Labaster/go-app/structures v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/joho/godotenv v1.5.1
	go.mongodb.org/mongo-driver v1.17.1
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)

replace github.com/Labaster/go-app/dbConf => ../dbConf

replace github.com/Labaster/go-app/structures => ../structures
