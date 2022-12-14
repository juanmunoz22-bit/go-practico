package main

func main() {
	server := NewServer("localhost:3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
