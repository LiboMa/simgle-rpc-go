## Simple TCP RPC example by GO

### clone project 

git clone https://github.com/LiboMa/simgle.git

### run project

#### Server side

```bash
cd server && go run rpc_simple_server.go
```
    ...
    2021/09/01 16:52:38 line type -> []uint8, msg type -> main.Message
    handle message PING, from client: 3
    2021/09/01 16:52:39 line type -> []uint8, msg type -> main.Message
    handle message PING, from client: 2
    2021/09/01 16:52:39 line type -> []uint8, msg type -> main.Message
    handle message PING, from client: 1
    2021/09/01 16:52:39 line type -> []uint8, msg type -> main.Message
    handle message PING, from client: 3
    ...


#### Client side with client id
```bash
cd client 
go run client.go -id 1
go run client.go -id 2
go run client.go -id 3

```

#### Result

    021/09/01 16:50:43 Sending message <-  {2 PING}
    2021/09/01 16:50:43 Timestame: 2021-09-01 16:50:43.796673 +0800 CST, Original message: {2 PING}, Response message: PONG,
    reponse Type: main.Response
    2021/09/01 16:50:44 Sending message <-  {2 PING}
    2021/09/01 16:50:44 Timestame: 2021-09-01 16:50:44.803715 +0800 CST, Original message: {2 PING}, Response message: PONG,
    reponse Type: main.Response
    2021/09/01 16:50:45 Sending message <-  {2 PING}
    2021/09/01 16:50:45 Timestame: 2021-09-01 16:50:45.805505 +0800 CST, Original message: {2 PING}, Response message: PONG,


### Reference:
* [https://gist.github.com/jordanorelli/2629049](https://gist.github.com/jordanorelli/2629049)
* [https://chai2010.cn/advanced-go-programming-book/ch4-rpc/readme.html](https://chai2010.cn/advanced-go-programming-book/ch4-rpc/readme.html)
* [https://github.com/cirocosta/sample-rpc-go/blob/master/server/server.go](https://github.com/cirocosta/sample-rpc-go/blob/master/server/server.go)
