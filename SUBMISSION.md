# Running

To run the code go must be installed and setup.

```sh
go install
go run .
```

Example adding a receipt using curl

```sh
curl -X POST \
	-H "Content-Type: application/json" \
	--data @examples/morning-receipt.json \
	localhost:8080/receipts/process
```

Example getting points a few hard coded receipt with curl

```sh
curl localhost:8080/receipts/FC09442B-C532-E7CF-879A-605E837D3709/points
curl localhost:8080/receipts/493AC2FD-22CE-9280-1853-B5C3480C8E92/points
```

## Testing

Run the test suite

```sh
go install .
go test .
```

## Enable debugging

To enable debug logging go to the _init_ function in _main.go_ and comment out
the log disable statements.

## Issues found

In _api.yml_ the pattern for _retailer_ is `^\S+$`. This pattern leaves out
_M&M Corner Market_ as valid retail if you go by extended regular expressions.
As a work around this implementation considers a _retailer_ value of length 1
or more.
