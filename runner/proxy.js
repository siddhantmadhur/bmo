console.log("Hello world");


const ws = new WebSocket("ws://localhost:9090/ws/proxy")

ws.addEventListener("message", async () => {
    console.log("Hello")
    var finish = false
    while (!finish) {
        const res = await fetch("http://localhost:9090");
        if (res.ok) {
            finish = true
        }
    }
    window.location.reload()
})
