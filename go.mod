module vebsrv

go 1.19

replace handlers => /home/danil/go/vebserv/finsimple/advent-2017/handlers

replace version => /home/danil/go/vebserv/finsimple/advent-2017/version

require (
	handlers v0.0.0-00010101000000-000000000000
	version v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/mux v1.8.0 // indirect
