


package main

import (
	"sync"
)

//defines a node and holds metadata for it
type node struct{
	id	string
	ip string
	job string
	cpu uint64
	mem uint64
	disk uint64
	mu	sync.RWMutex
}



//defines a node and holds metadata for it
type PeerNode struct{
	id	string
	ip string
	job string
	cpu uint64
	mem uint64
	disk uint64
	//mu	sync.RWMutex
}


type ResourceMSG struct{
	Func string `json:"func,omitempty`
	Id	string `json:"id,omitempty`
	Ip string `json:"ip,omitempty`
	Job string `json:"job,omitempty`
	CPU uint64 `json:"cpu,omitempty`
	Mem uint64 `json:"mem,omitempty`
	Disk uint64 `json:"disk,omitempty`
	//mu	sync.RWMutex
}