package param

type CreateLoadBalancerParam struct {
	BaseParam
	Params CreateLoadBalancerDetailParam `json:"params"`
}

type CreateLoadBalancerDetailParam struct {
	Name    string `json:"name"`    // Resource name
	VipUUID string `json:"vipUuid"` // VIP UUId
}

type UpdateLoadBalancerParam struct {
	BaseParam
	UUID               string                        `json:"uuid"` // Resource UUID, uniquely identifies the resource
	UpdateLoadBalancer UpdateLoadBalancerDetailParam `json:"updateLoadBalancer"`
}

type UpdateLoadBalancerDetailParam struct {
	Name        string `json:"name"`                  // Resource name
	Description string `json:"description,omitempty"` // Detailed description
}

type CreateLoadBalancerListenerParam struct {
	BaseParam
	Params CreateLoadBalancerListenerDetailParam `json:"params"`
}

type CreateLoadBalancerListenerDetailParam struct {
	Name             string `json:"name"`
	InstancePort     int    `json:"instancePort"`
	LoadBalancerPort int    `json:"loadBalancerPort"`
	Protocol         string `json:"protocol"`
	AclStatus        string `json:"aclStatus"`
	AclType          string `json:"aclType"`
}

type AddServerGroupToLoadBalancerListenerParam struct {
	BaseParam
	Params AddServerGroupToLoadBalancerListenerDetailParam `json:"params"`
}

type AddServerGroupToLoadBalancerListenerDetailParam struct {
	ServerGroupUUID string `json:"serverGroupUuid"`
}

type CreateLoadBalancerServerGroupParam struct {
	BaseParam
	Params CreateLoadBalancerServerGroupDetailParam `json:"params"`
}

type CreateLoadBalancerServerGroupDetailParam struct {
	Name string `json:"name"` // Resource name
}

type AddBackendServerToServerGroupParam struct {
	BaseParam
	Params AddBackendServerToServerGroupDetailParam `json:"params"`
}

type AddBackendServerToServerGroupDetailParam struct {
	VmNics []AddBackendServerToServerGroupVmNics `json:"vmNics"`
}

type AddBackendServerToServerGroupVmNics struct {
	UUID string `json:"uuid"`
}

type RemoveBackendServerFromServerGroupParam struct {
	BaseParam
	Params RemoveBackendServerFromServerGroupDetailParam `json:"removeBackendServerFromServerGroup"`
}

type RemoveBackendServerFromServerGroupDetailParam struct {
	VmNics     []string `json:"vmNicUuids"`
	ServersIps []string `json:"serverIps"`
}

type ChangeLoadBalancerBackendServerParam struct {
	BaseParam
	Params ChangeLoadBalancerBackendServerDetailParam `json:"changeLoadBalancerBackendServer"`
}

type ChangeLoadBalancerBackendServerDetailParam struct {
	VmNics     []string `json:"vmNics"`
	ServersIps []string `json:"servers"`
}

type GetCandidateVmNicsForLoadBalancerServerGroupParam struct {
	ServerGroupUuid  string
	LoadBalancerUuid string
}
