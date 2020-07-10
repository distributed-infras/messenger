# purpope

create a grpc between different languages

# installing go related stuff

[install](https://grpc.io/docs/languages/go/quickstart/) has most of the
information that you will need.

Since I already had go in my system, the remaining steps:

```bash
apt install -y protobuf-compiler
protoc --version                                # Ensure compiler version is 3+
export GO111MODULE=on                           # Enable module mode
go get github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"
```

For the latter part, I created a env.sh file you can just source

```bash
. ./env.sh
```

## test run

You can follow the example and clone grpc-go repo and follow the instructions
in the quickstart.

```bash
git clone -b v1.30.0 https://github.com/grpc/grpc-go
cd grpc-go/examples/helloworld
go run greeter_server/main.go
go run greeter_client/main.go
```

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

# editor

[vim-go](http://github.com/fatih/vim-go)

## notes

Simple commands you can use right away:

- :GoRun
- :GoTest
- [tutorial](https://github.com/fatih/vim-go-tutorial)

## how to deal with the quickfix window
- :cc [number]           - go to a particular error
- :cn                    - go to next error
- :cp                    - go to previous error
- :cclose                - close quickfix window
- :help quickfix
- [more help](http://vimdoc.sourceforge.net/htmldoc/quickfix.html#quickfix.txt)

## indent

I don't like the default indent (8 spaces! -yuck-)

```vimrc
autocmd BufNewFile,BufRead *.go setlocal noexpandtab tabstop=4 shiftwidth=4 
```

or you can add the following file

```bash
echo setlocal noexpandtab tabstop=2 shiftwidth=2 > ~/.vim/after/ftplugin/go.vim
```

# Discord webhook

1. created a #github channel
1. settings/webhooks/create webhook
1. copy the link
1. github repo/settings/webhooks/add webhook
1. payload url: past the discord link and add this suffix /github
1. content type: application/json
1. send me everything
1. create webhook
1. now #github should be receiving notification when repo is changed
1. for pictorial [gist](https://gist.github.com/jagrosh/5b1761213e33fc5b54ec7f6379034a22)

# installing python related stuff

(I haven't used this yet, that is next in the list.)

pip install grpcio grpcio-tools
