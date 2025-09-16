package client

import (
	"fmt"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

func (cli *ZSClient) CreateLoadBalancer(params param.CreateLoadBalancerParam) (view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	return resp, cli.Post("v1/load-balancers", params, &resp)
}

func (cli *ZSClient) DeleteLoadBalancer(uuid string, mode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers", uuid, string(mode))
}

func (cli *ZSClient) QueryLoadBalancer(params param.QueryParam) ([]view.LoadBalancerInventoryView, error) {
	var resp []view.LoadBalancerInventoryView
	return resp, cli.List("v1/load-balancers", &params, &resp)
}

func (cli *ZSClient) GetLoadBalancer(uuid string) (view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	return resp, cli.Get("v1/load-balancers", uuid, nil, &resp)
}

func (cli *ZSClient) UpdateLoadBalancer(params param.UpdateLoadBalancerParam) (view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	return resp, cli.Put("v1/load-balancers", params.UUID, params, &resp)
}

func (cli *ZSClient) CreateLoadBalancerListener(loadBalancerUuid string, params param.CreateLoadBalancerListenerParam) (view.LoadBalancerListenerInventoryView, error) {
	var resp view.LoadBalancerListenerInventoryView
	requestURL := fmt.Sprintf("v1/load-balancers/%s/listeners", loadBalancerUuid)
	return resp, cli.Post(requestURL, params, &resp)
}

func (cli *ZSClient) QueryLoadBalancerListener(params param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List("v1/load-balancers/listeners", &params, &resp)
}

func (cli *ZSClient) GetLoadBalancerListener(uuid string) (view.LoadBalancerListenerInventoryView, error) {
	var resp view.LoadBalancerListenerInventoryView
	return resp, cli.Get("v1/load-balancers/listeners", uuid, nil, &resp)
}

func (cli *ZSClient) DeleteLoadBalancerListener(uuid string, mode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners", uuid, string(mode))
}

func (cli *ZSClient) CreateLoadBalancerServerGroup(loadBalancerUuid string, params param.CreateLoadBalancerServerGroupParam) (view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.LoadBalancerServerGroupInventoryView
	requestURL := fmt.Sprintf("v1/load-balancers/%s/servergroups", loadBalancerUuid)
	return resp, cli.Post(requestURL, params, &resp)
}

func (cli *ZSClient) QueryLoadBalancerServerGroup(params param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, error) {
	var resp []view.LoadBalancerServerGroupInventoryView
	return resp, cli.List("v1/load-balancers/servergroups", &params, &resp)
}

func (cli *ZSClient) GetLoadBalancerServerGroup(uuid string) (view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.LoadBalancerServerGroupInventoryView
	return resp, cli.Get("v1/load-balancers/servergroups", uuid, nil, &resp)
}

func (cli *ZSClient) DeleteLoadBalancerServerGroup(uuid string, mode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups", uuid, string(mode))
}

func (cli *ZSClient) AddServerGroupToLoadBalancerListener(listenerUuid string, params param.AddServerGroupToLoadBalancerListenerParam) error {
	requestURL := fmt.Sprintf("v1/load-balancers/listeners/%s/servergroups", listenerUuid)
	return cli.Post(requestURL, params, nil)
}

func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(params *param.GetCandidateVmNicsForLoadBalancerServerGroupParam) error {
	return cli.Get("v1/load-balancers/servergroups/candidate-nics", "", params, nil)
}

func (cli *ZSClient) AddBackendServerToServerGroup(serverGroupUuid string, params param.AddBackendServerToServerGroupParam) error {
	requestURL := fmt.Sprintf("v1/load-balancers/servergroups/%s/backendservers", serverGroupUuid)
	return cli.Post(requestURL, params, nil)
	// return cli.PostWithRespKey("v1/load-balancers/servergroups", serverGroupUuid, "backendservers/actions", "", params, nil)
}

func (cli *ZSClient) RemoveBackendServerToServerGroup(serverGroupUuid string, params param.RemoveBackendServerFromServerGroupParam) error {
	return cli.PutWithSpec("v1/load-balancers/servergroups", serverGroupUuid, "backendservers/actions", "", params, nil)
}

func (cli *ZSClient) ChangeLoadBalancerBackendServer(serverGroupUuid string, params param.ChangeLoadBalancerBackendServerParam) error {
	return cli.PutWithSpec("v1/load-balancers/servergroups", serverGroupUuid, "backendservers/actions", "", params, nil)
}
