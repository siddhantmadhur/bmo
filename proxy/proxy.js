console.log("BMO: Proxy loaded :D")

const websocket = new WebSocket("ws://127.0.0.1:9090/ws/proxy")

websocket.addEventListener("message", (event) => {
    console.log("BMO: " + event)
})
