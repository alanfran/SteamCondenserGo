package test

import (
	"testing"

	servers "github.com/alanfran/SteamCondenserGo"
)

func TestMinecraft(t *testing.T) {
	minecraftServer := servers.MinecraftServer{
		Address: "178.32.48.244:25565",
	}

	_, err := minecraftServer.GetInfo()
	if err != nil {
		t.Error(err)
	}
}
