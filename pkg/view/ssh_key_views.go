package view

type SshKeyInventoryView struct {
	BaseInfoView
	BaseTimeView

	PublicKey string `json:"publicKey"`
}
