from socket import *
import sys

serverSocket = socket(AF_INET, SOCK_STREAM)

# Prepare a server socket
serverPort = 6789
serverSocket.bind(('10.218.8.133', serverPort))
serverSocket.listen(1)

while True:
    print('Ready to serve...')
    connectionSocket, addr = serverSocket.accept()

    try:
        message = connectionSocket.recv(1024).decode()
        filename = message.split()[1]
        f = open(filename[1:])  # Remove leading '/'
        outputdata = f.read()
        f.close()

        # Send HTTP header
        header = 'HTTP/1.1 200 OK\r\n\r\n'
        connectionSocket.send(header.encode())

        # Send HTML content
        connectionSocket.send(outputdata.encode())
        connectionSocket.send("\r\n".encode())

    except IOError:
        # File not found
        errorMessage = "HTTP/1.1 404 Not Found\r\n\r\n"
        connectionSocket.send(errorMessage.encode())
        connectionSocket.send("<html><body><h1>404 Not Found</h1></body></html>\r\n".encode())

    connectionSocket.close()

serverSocket.close()
sys.exit()
