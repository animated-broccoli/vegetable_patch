/*
* Author: Lucas Chaufournier
* Created: April 2nd, 2018
* This defines a node type to store data about a peer machine.
*/

package main

// import (
// 	"sync"
// 	_"fmt"
// 	"github.com/shirou/gopsutil/mem"
// 	"github.com/shirou/gopsutil/cpu"
// 	"github.com/shirou/gopsutil/disk"
// 	"github.com/satori/go.uuid"
// )


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



