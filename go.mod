module example/m

go 1.27.1

replace example/hello => ./hello

replace example/study => ./study

replace example/contact => ./contact

require example/contact v0.0.0
