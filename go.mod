module example/m

go 1.27.1

replace example/hello => ./hello

replace example/study => ./study

replace example/contact => ./contact

require (
	example/contact v0.0.0
	github.com/eiannone/keyboard v0.0.0-20220611211555-0d226195f203
)

require golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
