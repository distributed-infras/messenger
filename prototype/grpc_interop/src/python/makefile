.PHONY: server_start client_start

venv=${HOME}/anaconda3/envs/py38/bin
server_start:
	go run ../go/main.go
	
client_start:
	$(venv)/python client.py
