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
				Key:  pulumi.String("my-private-key"),
			},
			NodeConfiguration: &k3os.NodeConfigurationArgs{
				SshAuthorizedKeys: pulumi.StringArray{
					pulumi.String("github:spigell"),
				},
				Hostname: pulumi.String("test01"),
				BootCmd: pulumi.StringArray{
					pulumi.String("echo This is boot_cmd"),
				},
				RunCmd: pulumi.StringArray{
					pulumi.String("echo This is run_cmd"),
				},
				InitCmd: pulumi.StringArray{
					pulumi.String("echo This is init_cmd"),
				},
				WriteFiles: &k3os.CloudInitFilesArray{
					&k3os.CloudInitFilesArgs{
						Path:    pulumi.String("/home/rancher/greetings"),
						Content: pulumi.String("Hello World!"),
					},
				},
				K3OS: &k3os.K3OSArgs{
					Password:  pulumi.String("password"),
					Datasources: pulumi.StringArray{
						pulumi.String("cdrom"),
						pulumi.String("digitalocean"),
					},
					ServerUrl: pulumi.String("server2"),
					Token:     pulumi.String("token"),
					Labels: pulumi.StringMap{
						"node.k3os.infraunlimited.tech/local": pulumi.String("true"),
						"node.k3os.infraunlimited.tech/kind":  pulumi.String("local"),
					},
					Modules: pulumi.StringArray{
						pulumi.String("wireguard"),
					},
					Sysctls: pulumi.StringMap{
						"kernel.kptr_restrict": pulumi.String("1"),
					},
					NtpServers: pulumi.StringArray{
						pulumi.String("0.us.pool.ntp.org"),
						pulumi.String("1.us.pool.ntp.org"),
					},
					Environment: pulumi.StringMap{
						"ENV_1": pulumi.String("true"),
					},
					Taints: pulumi.StringArray{
						pulumi.String("node.k3os.infraunlimited.tech/unreachable=Deny:NoExecute"),
					},
				},
			},
		})
		return nil
	})
}
