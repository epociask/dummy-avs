package main

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
)

func main() {
	logger, err := logging.NewZapLogger("development")
	if err != nil {
		panic(err)
	}

	nodeApi := nodeapi.NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)
	nodeApi.RegisterNewService("testServiceId", "testServiceName", "testServiceDescription", nodeapi.ServiceStatusInitializing)

	listener := nodeApi.Start()

	nodeApi.UpdateHealth(nodeapi.PartiallyHealthy)
	_ = nodeApi.UpdateServiceStatus("testServiceId", nodeapi.ServiceStatusDown)
	_ = nodeApi.DeregisterService("testServiceId")

	nodeApi.RegisterNewService("testServiceId", "testServiceName", "testServiceDescription", nodeapi.ServiceStatusInitializing)

	<-listener
}
