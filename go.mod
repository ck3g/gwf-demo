module demoapp

go 1.17

replace github.com/ck3g/gwf => ../gwf

require (
	github.com/ck3g/gwf v0.0.0-20211018084703-71fd9cc80779
	github.com/go-chi/chi/v5 v5.0.4
)

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
