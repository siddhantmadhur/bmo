package main



func main() {
    var config Config
    config.Read()

    var runner Runner
    runner.Start(&config)
}
