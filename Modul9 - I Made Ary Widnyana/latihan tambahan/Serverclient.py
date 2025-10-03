from socket import *
import threading

def handle_client(connectionSocket):
    try:
        message = connectionSocket.recv(1024).decode()
        filename = message.split()[1]
        f = open(filename[1:])
        outputdata = f.read()
        f.close()

        connectionSocket.send('HTTP/1.1 200 OK\r\n\r\n'.encode())
        connectionSocket.send(outputdata.encode())
        connectionSocket.send("\r\n".encode())
    except IOError:
        connectionSocket.send('HTTP/1.1 404 Not Found\r\n\r\n'.encode())
        connectionSocket.send("<html><body><h1>404 Not Found</h1></body></html>\r\n".encode())
    connectionSocket.close()

serverSocket = socket(AF_INET, SOCK_STREAM)
serverSocket.bind(('10.218.8.133', 6789))
serverSocket.listen(5)

print('Web Server ready...')

while True:
    connectionSocket, addr = serverSocket.accept()
    threading.Thread(target=handle_client, args=(connectionSocket,)).start()
