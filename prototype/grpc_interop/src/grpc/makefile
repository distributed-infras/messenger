# make sure you source ../../env.sh

# one shell
# https://stackoverflow.com/questions/24736146/how-to-use-virtualenv-in-makefile
# use this so that all commands will run from the same shell

SHELL=/bin/bash
.ONESHELL:

.PHONY: top

proto=./message.proto
pyprotoc=python -m grpc_tools.protoc
pyenv=py38
pyoutdir=./gpython
pyout1=$(pyoutdir)/message_pb2_grpc.py
pyout2=$(pyoutdir)/message_pb2.py
pyout=$(pyout1) $(pyout2)

# my version of make doesn't support this feature
# .NOTPARALLEL: $(pyout)

godir=message
goout=$(godir)/my.org/message/message.pb.go

top: $(goout) $(pyout)

$(goout): $(proto)
	mkdir -p $(godir)
	protoc --go_out=plugins=grpc:message $(proto)

$(pyout2): $(proto)
	source $$(conda info --base)/etc/profile.d/conda.sh
	conda activate $(pyenv)
	$(pyprotoc) -I. --python_out=$(pyoutdir) --grpc_python_out=$(pyoutdir) $(proto)

$(pyout1): $(pyout2)
	# do nothing
	
