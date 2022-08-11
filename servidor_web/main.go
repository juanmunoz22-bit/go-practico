package main

func main() {
	server := NewServer("localhost:3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
