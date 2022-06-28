## How this example client was generated

> Follow the instructions for the tool of choice. Here we provide an example in golang.

1. Add a [`go.mod`](./go.mod) file to set a module name `github.com/ConquestSolutions/conquest_api.examples/go-swagger`

2. Install go-swagger from https://github.com/go-swagger/go-swagger

> Hint: Download the github release and put it in this folder.

3. Generate the client using `./swagger generate client -f https://raw.githubusercontent.com/ConquestSolutions/conquest_api.openapi/master/4.1.5.json`

4. Then fetch the dependencies for the new client (as suggested in the output).

    `go get -u -f ./...`
