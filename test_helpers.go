package dedupe

import (
	"context"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func withRedis(t *testing.T, f func()) {
	ctx := context.Background()
	containerRequest := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})

	if err != nil {
		t.Fatalf("Failed to start Redis container: %v", err)
	}

	host, _ := redisContainer.Host(ctx)
	port, _ := redisContainer.MappedPort(ctx, "6379")

	os.Setenv("REDIS_HOST", host)
	os.Setenv("REDIS_PORT", port.Port())

	defer testcontainers.CleanupContainer(t, redisContainer)
	f()
}
