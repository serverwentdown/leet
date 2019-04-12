
# leet

Environmental lighting over gRPC (WTF?!)

## WTF?!

At home, I have a strip of 287 WS2812B LEDs connected to a Raspberry Pi, that needs to be used as a light, mood lighting, and for notifications. I don't want to be a typical person and write a monolithic application, so `leet` is a simple gRPC server that renders basic opacity composition and buffered animations onto the LED strip. This also enables me to add aditional sources of data in the future from anywhere on my home network, including servers, phones and your IoT Blockchain AI appliances.

## Architecture

```
 -------------------         ---------------
 | Lighting Web UI |  -------| IMAP Client |
 -------------------  |      ---------------
           |          |
 -------------------- | --------------------------
 | Lighting Backend | | | Server Load Monitoring |
 -------------------- | --------------------------
           |          |              |
        ----------------------------------  -----------------
        |    Leet on the Raspberry Pi    |--| WS2812B Strip |
        ----------------------------------  -----------------
```

## Protocol

See [`framebuffer.proto`](framebuffer/framebuffer.proto) for the actual gRPC protocol

## Developing

1. [Install](https://grpc.io/docs/quickstart/go.html#install-protocol-buffers-v3) `protoc` and `protoc-gen-go`
2. Install the `rpi_ws281x` library: `./get-ws2812x.sh` (WARNING: this installs four files into /usr/local)
3. Get dependencies: `go get`
4. Start building with `go build`

