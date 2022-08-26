# wscat
一个 `WebSocket` 的客户端，可以把 `WebSocket` 协议中的 `BinaryMessage` 转化标准输入输出或者TCP请求。
结合 `WebSocket` 后端实现可以完成 `TCP Over WebSocket` 的能力！

# 使用方法
```
NAME:
   wscat - a websocket tool.

USAGE:
   wscat [global options]  

GLOBAL OPTIONS:
   --endpoint value, -e value, -t value  WebSocket endpoint, ws:// or wss://
   --listen value, -l value, -p value    Listen port to replace stdin, 1337, 127.0.0.1:1337
```
## 标准输入输出
```
wscat --endpoint ws://example.com/ws
wscat --endpoint wss://example.com/ws
```

## 监听端口并转发
```
wscat --endpoint ws://example.com/ws -p 1337
wscat --endpoint ws://example.com/ws -l :1337
wscat --endpoint ws://example.com/ws -l 127.0.0.1:1337
wscat --endpoint ws://example.com/ws -l 0.0.0.0:1337
```
