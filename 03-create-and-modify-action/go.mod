module github.com/ConquestSolutions/apiv2.examples/03-create-and-modify-action

go 1.14

replace github.com/ConquestSolutions/apiv2.examples/go-swagger v0.0.0 => ../go-swagger

replace github.com/ConquestSolutions/apiv2.examples/conquest_api v0.0.0 => ../conquest_api

require github.com/ConquestSolutions/apiv2.examples/conquest_api v0.0.0

require (
	github.com/ConquestSolutions/apiv2.examples/go-swagger v0.0.0
	github.com/go-openapi/runtime v0.19.14
	github.com/go-openapi/strfmt v0.19.5
	github.com/stretchr/testify v1.4.0
)
