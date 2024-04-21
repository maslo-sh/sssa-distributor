package server

func RunServer() {
	r := NewRouter()
	r.Run(":8080")
}
