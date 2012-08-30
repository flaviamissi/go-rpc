// this package implements a basic api structure to
// allow pluggable backends for differents IaaS software
// apis implementations.
// Two common examples are the EC2 api and the OpenStack
// api (the only api supported by rackspace)
package server

type InstanceRequest struct {
	name      string //openstack only
	ImageRef  string // id for amz, id or url for openstack
	MinCount  int    //amz only
	MaxCount  int    //amz only
	KeyName   string
	AvailZone string
	UserData  []byte
	SecGroup  string
}

type InstanceResponse struct {
	Id          string // Instance id
	TenantId    string // owner id for amz
	PrivateAddr string
	PublicAddr  string
	Status      string
	Progress    int //openstack only, goes from 0% to 100%
}

type IaaS struct{}

// NewInstance sample implementation
// the filled fields in instResp should
// be filled with values from the request made
// to the IaaS api
func (s *IaaS) NewInstance(instReq *InstanceRequest, instResp *InstanceResponse) error {
	instResp.Id = "i-001"
	instResp.TenantId = "uuid123"
	instResp.PrivateAddr = "192.0.4.5"
	instResp.PublicAddr = "192.0.6.9"
	instResp.Status = "BUILD"
	instResp.Progress = 0
	return nil
}

func main() {
	println("foo bar")
}
