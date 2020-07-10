# purpope

create a grpc between different languages

# installing python related stuff

(I haven't used this yet, that is next in the list.)

pip install grpcio grpcio-tools

# golang


# creating a proto

check the makefile in src/grpc directory

# running the example

```bash
cd src/go/

go run ./main.go
go test
```

The first one will appear hanging, but it is running as a server, expecting a request.

The second one mocks that where an empty request is sent and the default
response is obtained.


