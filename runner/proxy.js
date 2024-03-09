console.log("BMO: Proxy loaded :D")

const websocket = new WebSocket("ws://127.0.0.1:9090/ws/proxy")

websocket.addEventListener("message", async () => {
    var finish = false;
    while (!finish) {
        const res = await fetch(location.href)
        finish = res.ok
        if (res.ok) {
            document.body.innerHTML = await res.text()
        }
    }
})
