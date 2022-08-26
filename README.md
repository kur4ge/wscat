# wscat
A `WebSocket` client can convert `BinaryMessage` in the `WebSocket` protocol to standard input and output or TCP requests.
Combined with the `WebSocket` backend implementation, the ability of `TCP Over WebSocket` can be completed!

# USAGE
```
NAME:
   wscat - a websocket tool.

USAGE:
   wscat [global options]  

GLOBAL OPTIONS:
   --endpoint value, -e value, -t value  WebSocket endpoint, ws:// or wss://
   --listen value, -l value, -p value    Listen port to replace stdin, 1337, 127.0.0.1:1337
```
## Standard input and output
```
wscat --endpoint ws://example.com/ws
wscat --endpoint wss://example.com/ws
```

## Listen port and forward
```
wscat --endpoint ws://example.com/ws -p 1337
wscat --endpoint ws://example.com/ws -l :1337
wscat --endpoint ws://example.com/ws -l 127.0.0.1:1337
wscat --endpoint ws://example.com/ws -l 0.0.0.0:1337
```
