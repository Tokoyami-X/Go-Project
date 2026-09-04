module example/m

go 1.27.1

replace example/hello => ./hello

replace example/study => ./study

require (
	example/hello v0.0.0-00010101000000-000000000000
	example/study v0.0.0-00010101000000-000000000000
)
