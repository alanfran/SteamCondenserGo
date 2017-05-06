package servers

import (
	"testing"
)

func TestMinecraft(t *testing.T) {
	minecraftServer := MinecraftServer{
		Address: "178.32.48.244:25565",
	}

	_, err := minecraftServer.GetInfo()
	if err != nil {
		t.Error(err)
	}
}
