package test

import (
	"testing"

	servers "github.com/alanfran/SteamCondenserGo"
)

func TestGoldSrc(t *testing.T) {
	_, err := servers.QueryGoldServer("74.91.113.128:27015")
	if err != nil {
		t.Fatalf("Failed to get gold source server info")
	}
}
