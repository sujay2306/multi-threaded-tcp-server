### multi-threaded-tcp-server

This Go TCP server accepts incoming connections, processes basic HTTP requests, and sends back responses. The server demonstrates key concepts like socket programming and system calls. It handles incoming connections, reads requests, processes them, and sends responses. However, the initial version of the server could only handle one connection at a time due to its single-threaded nature.

### Key Concepts

- **Raw Socket Programming:**  
  The server uses Go’s `net` package to create a TCP socket and listen for incoming connections.

- **System Calls:**  
  The server interacts with the operating system to accept connections, read requests, and send responses.

- **Basic HTTP Handling:**  
  The server reads simple HTTP requests and sends back basic HTTP responses.

- **Single-threaded Limitation:**  
  The server processes requests one by one and can’t handle multiple connections simultaneously.

- **Concurrency with Goroutines:**  
  To improve performance, the server uses **goroutines** (Go's lightweight threads) to handle multiple requests at the same time.


#### Sequencial Processing 


https://github.com/user-attachments/assets/ef232bbe-d83e-4023-972f-33c390885653




#### Parallel Processing 



https://github.com/user-attachments/assets/45558c2a-7e5d-4aac-a2ea-f9f447958ba2






