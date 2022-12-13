package main

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a DigitalOcean resource (Domain)

		cluster, err := digitalocean.NewKubernetesCluster(ctx, "my-cluster", &digitalocean.KubernetesClusterArgs{
			Name:    pulumi.String("my-cluster"),
			Region:  pulumi.String("fra1"),
			Version: pulumi.String("1.25.4-do.0"),
			NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
				Name:      pulumi.String("default"),
				NodeCount: pulumi.Int(1),
				AutoScale: pulumi.Bool(false),
				Size:      pulumi.String("s-2vcpu-4gb"),
			},
		})
		if err != nil {
			return err
		}

		_, err = digitalocean.NewKubernetesNodePool(ctx, "my-cluster-pool", &digitalocean.KubernetesNodePoolArgs{
			ClusterId: cluster.ID(),
			Name:      pulumi.String("my-cluster-pool"),
			NodeCount: pulumi.Int(1),
			AutoScale: pulumi.Bool(false),
			Size:      pulumi.String("s-2vcpu-4gb"),
		})
		if err != nil {
			return err
		}

		// Export the name of the domain
		ctx.Export("clusterName", cluster.Name)
		return nil
	})
}
