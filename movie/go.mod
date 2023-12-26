module github.com/kuzminal/movie-app/movie

replace (
	github.com/kuzminal/movie-app/metadata => ../metadata
	github.com/kuzminal/movie-app/rating => ../rating
)

go 1.19

require (
	github.com/kuzminal/movie-app/metadata v0.0.0-00010101000000-000000000000
	github.com/kuzminal/movie-app/rating v0.0.0-00010101000000-000000000000
)
