install_proto:
	sudo apt install -y protobuf-compiler
	protoc --version
install_plugin:
	export GO111MODULE=on; \
        go get github.com/golang/protobuf/protoc-gen-go
install_python:
	python -m pip install --upgrade pip
	# in case you need a virtualenv:
	# 	
	# python -m pipi install virtualenv
	# virtualenv env
	# source venv/bin/activate
	# python -m pip install --upgrade pip
	python -m pip install grpcio
	# sudo python -m pip install grpcio
	python -m pip install grpcio-tools
