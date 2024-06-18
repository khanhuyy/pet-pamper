gql-gen:
	go run github.com/99designs/gqlgen generate


media-server:
	filebrowser -d database/media/file-server.db --noauth -p 3005

run-all:
	go run main.go | filebrowser -d database/media/file-server.db --noauth -p 3005
