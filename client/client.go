package main

import (
	"fmt"
	"net/rpc"
)

// same structs from server/server.go
// copied so this package doesnt depends
// on server pkg. Yeah, pretty obvious.
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

// wrapper function to rpc call
// ensures the data types are correct
func NewInstance(client *rpc.Client, instReq *InstanceRequest, instResp *InstanceResponse) error {
	return client.Call("IaaS.NewInstance", instReq, instResp)
}

func main() {
	client, e := rpc.DialHTTP("tcp", "localhost:1234")
	if e != nil {
		panic(e)
	}
	instReq := new(InstanceRequest)
	instResp := new(InstanceResponse)
	e = NewInstance(client, instReq, instResp)
	fmt.Println(instResp)
}
