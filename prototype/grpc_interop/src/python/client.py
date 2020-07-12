from __future__ import print_function
import logging

import grpc

from gpython import message_pb2, message_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:3000') as channel:
        stub = message_pb2_grpc.MessageStub(channel)
        response = stub.Search(message_pb2.Request(payload='hello'))
    print("Response: " + response.payload)


if __name__ == '__main__':
    logging.basicConfig()
    run()
