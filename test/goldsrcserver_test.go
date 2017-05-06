package test

import (
	"testing"

	servers "github.com/alanfran/SteamCondenserGo"
)

func TestGoldSrc(t *testing.T) {
	goldServer := servers.GoldServer{
		Address: "74.91.113.128:27015",
	}

	_, err := goldServer.GetInfo()
	if err != nil {
		t.Fatalf("Failed to get gold source server info")
	}
}
