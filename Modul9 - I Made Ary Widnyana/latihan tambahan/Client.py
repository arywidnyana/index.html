from socket import *
import sys

# Ambil parameter dari command line
server_host = sys.argv[1]
server_port = int(sys.argv[2])
filename = sys.argv[3]

# Buat socket TCP dan hubungkan ke server
clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect((server_host, server_port))

# Kirim permintaan GET ke server
request = f"GET /{filename} HTTP/1.1\r\nHost: {server_host}\r\n\r\n"
clientSocket.send(request.encode())

# Terima respon dari server (hanya 1 kali)
response = b""
while True:
    potongan = clientSocket.recv(4096)
    if not potongan:
        break
    response += potongan

print(response.decode())


# Tutup koneksi
clientSocket.close()
