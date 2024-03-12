console.log("Hello world");


const ws = new WebSocket("ws://localhost:9090/ws/proxy")

ws.addEventListener("message", async () => {

    const res = await fetch("http://localhost:9090");
    document.body.innerHTML = await res.text()
    finish = true
})
