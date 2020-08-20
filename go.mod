module github.com/rmc-code/rmc-client

go 1.13

replace github.com/rmc-code/rmc-client/ethereumRMC => ./ethereumRMC

require (
	github.com/rmc-code/rmc-client/ethereumRMC v1.9.15
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)
