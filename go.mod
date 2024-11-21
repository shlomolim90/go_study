module github.com/go_study

go 1.23.3

replace github.com/go_study/simple_module => ./simple_module

replace github.com/go_study/simple_server => ./simple_server

require (
	github.com/go_study/simple_module v0.0.0-00010101000000-000000000000
	github.com/go_study/simple_server v0.0.0-00010101000000-000000000000
	github.com/godbus/dbus/v5 v5.1.0
)
