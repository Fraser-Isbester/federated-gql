module github.com/fraser-isbester/federated-gql/services/graphql-gateway

go 1.24.0

require (
	connectrpc.com/connect v1.18.1
	github.com/99designs/gqlgen v0.17.66
	github.com/fraser-isbester/federated-gql/gen/go v0.0.0-20250224025919-47a37e4bc4f6
	github.com/go-chi/chi/v5 v5.2.1
	github.com/go-chi/cors v1.2.1
	github.com/gorilla/websocket v1.5.0
	github.com/vektah/gqlparser/v2 v2.5.22
)

require (
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/urfave/cli/v2 v2.27.5 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/fraser-isbester/federated-gql/gen/go => ../../gen/go

tool github.com/99designs/gqlgen
