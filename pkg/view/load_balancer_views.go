package view

type LoadBalancerInventoryView struct {
	BaseInfoView
	BaseTimeView

	VipUUID      string `json:"vipUuid"`      // VIP UUID
	ResourceUUID string `json:"resourceUuid"` // Resource UUID
}

type LoadBalancerListenerInventoryView struct {
	BaseInfoView
	BaseTimeView

	LoadBalancerUUID string                    `json:"loadBalancerUuid"`
	InstancePort     int                       `json:"instancePort"`
	LoadBalancerPort int                       `json:"loadBalancerPort"`
	Protocol         string                    `json:"protocol"`
	VmNicRefs        []string                  `json:"vmNicRefs"`
	AclRefs          []string                  `json:"aclRefs"`
	CertificateRefs  []string                  `json:"certificateRefs"`
	ServerGroupRefs  []ListenerServerGroupRefs `json:"serverGroupRefs"`
}

type LoadBalancerServerGroupInventoryView struct {
	BaseInfoView
	BaseTimeView

	LoadBalancerUUID string                 `json:"loadBalancerUuid"`
	VmNicRefs        []ServerGroupVmnicRefs `json:"vmNicRefs"`
}

type ListenerServerGroupRefs struct {
	ID              string `json:"id"`
	ListenerUUID    string `json:"listenerUuid"`
	ServerGroupUUID string `json:"serverGroupUuid"`
}

type ServerGroupVmnicRefs struct {
	ID              string `json:"id"`
	ServerGroupUuid string `json:"serverGroupUuid"`
	VmNicUuid       string `json:"vmNicUuid"`
	Weight          string `json:"weight"`
	Status          string `json:"status"`
}
