package sopsearch

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetKubeClient creates a Kubernetes client which can be used to access
//the ConfigMap stored on the oc environment.
func GetKubeClient() (client.Client, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// creating a client for reading the AccountID
	cli, err := client.New(config, client.Options{})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// getConfigMap uses the Kubernetes client and context to get the ConfigMap that
//stores various different values used throughout the program.
func getConfigMap(ctx context.Context, cli client.Client) (v1.ConfigMap, error) {
	var cmap v1.ConfigMap
	if err := cli.Get(ctx, types.NamespacedName{
		Name:      "configmap",
		Namespace: "sop-search",
	}, &cmap); err != nil {
		return v1.ConfigMap{}, err
	}
	return cmap, nil
}
