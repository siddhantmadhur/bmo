console.log("Hello world");


const ws = new WebSocket("ws://localhost:9090/ws/proxy")

ws.addEventListener("message", async () => {
    let finish = true;
    while (finish) {
        const res = await fetch("http://localhost:9090");
        finish = !res.ok 
        if (!finish) {
            window.location.reload();
        }
    }
     
})
