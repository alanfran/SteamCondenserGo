package servers

type server struct {
	Address string
}

type Response struct {
	Name       string
	Map        string
	Game       string
	Players    int
	MaxPlayers int
	Bots       int
	ServerType int
	Secured    bool
}
