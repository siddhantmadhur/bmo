console.log("BMO: Proxy loaded :D")

const websocket = new WebSocket("ws://127.0.0.1:9090/ws/proxy")

websocket.addEventListener("message", async () => {
    setTimeout(function() {
        location.replace(location.href)
    }, 1500)
})
