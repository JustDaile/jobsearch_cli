package internal

// These should be set via `go build -ldflags`

var (
	AppName       string
	SiteName      string = "default"
	Environment   string = "dev"
	Version       string
	Timestamp     string
	RedisHost     string = "localhost"
	RedisPort     string = "6543"
	RedisPassword string = ""
)
