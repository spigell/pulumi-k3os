package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spigell/pulumi-k3os/sdk/go/k3os"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, _ = k3os.NewNode(ctx, "test", &k3os.NodeArgs{
			Connection: &k3os.ConnectionArgs{
				Addr: pulumi.String("myserver:22"),
				User: pulumi.String("rancher"),
				Key: pulumi.String("my-private-key"),
			},
			NodeConfiguration: &k3os.NodeConfigurationArgs{
				Hostname: pulumi.String("node01"),
				K3OS: &k3os.K3OSArgs{
					Password: pulumi.String("password"),
				},
			},
		})
		return nil
	})
}
