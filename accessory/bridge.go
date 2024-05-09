package accessory

import "github.com/LUJUNQUAN/hap/service"

type Bridge struct {
	*A
	ServiceLabel *service.ServiceLabel
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info) *Bridge {
	a := Bridge{}
	a.A = New(info, TypeBridge)

	a.ServiceLabel = service.NewServiceLabel()
	a.AddS(a.ServiceLabel.S)

	return &a
}
