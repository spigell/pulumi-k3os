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
				Hostname: pulumi.String("test01"),
				BootCmd: pulumi.StringArray{
					pulumi.String("echo This is boot_cmd"),
				},
				K3OS: &k3os.K3OSArgs{
					Password: pulumi.String("password"),
					ServerUrl: pulumi.String("server2"),
					Token: pulumi.String("token"),
					Labels: pulumi.StringMap{
						"k3os.infraunlimited.tech/local": pulumi.String("true"),
						"k3os.infraunlimited.tech/kind": pulumi.String("local"),
					},
					Environment: pulumi.StringMap{
						"ENV_1": pulumi.String("true"),
					},
					Taints: pulumi.StringArray{
						pulumi.String("k3os.infraunlimited.tech/unreachable=Deny:NoExecute"),
					},
				},
			},
		})
		return nil
	})
}
