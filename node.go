/*
* Author: Lucas Chaufournier
* Created: April 2nd, 2018
* This defines a node type to store data about a machine.
*/

package main

import (
	"sync"
	"fmt"
	"os"
	"net"
	"strings"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/satori/go.uuid"
)


func getIP()(string){
    conn, err := net.Dial("udp", "8.8.8.8:80")
    // handle err...
	if err != nil{
		fmt.Println("Unable to get IP. No Network")
		os.Exit(1)
	}
     defer conn.Close()
	 localAddr := conn.LocalAddr().(*net.UDPAddr)
	 localAddr_string := strings.Split(localAddr.String(),":")
	 return localAddr_string[0]
}

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


func (n *node)Get_ip()(string){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.ip
}

func (n *node)Set_job(data string){
	n.mu.Lock()
	defer n.mu.Unlock()
	n.job = data
}

func (n *node)Get_job()(string){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.job
}

func (n *node)Set_id(data string){
	n.mu.Lock()
	defer n.mu.Unlock()
	n.id = data
}

func (n *node)Get_id()(string){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.id
}


func (n *node)Set_cpu(data uint64){
	n.mu.Lock()
	defer n.mu.Unlock()
	n.cpu = data
}

func (n *node)Get_cpu()(uint64){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.cpu
}

func (n *node)Set_mem(data uint64){
	n.mu.Lock()
	defer n.mu.Unlock()
	n.mem = data
}

func (n *node)Get_mem()(uint64){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.mem
}

func (n *node)Set_disk(data uint64){
	n.mu.Lock()
	defer n.mu.Unlock()
	n.disk = data
}

func (n *node)Get_disk()(uint64){
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.disk
}


var ( 
	n *node
	once sync.Once
)

func fetch_cores()(uint64){
	c,_ := cpu.Info()
	total_cores := uint64(0)
	for _,core := range c{
		total_cores += uint64(core.Cores)
		//fmt.Printf("Cores: %d\n",core.Cores)
	}

	return total_cores

}

func Node()*node{
	once.Do(func(){

		//query for stats
		v, _ := mem.VirtualMemory()
		d, _ := disk.Usage("/")
		c := fetch_cores()
		id := uuid.Must(uuid.NewV4())

		if n == nil{
			n = &node{id: id.String(),
				ip: getIP(),
				job: "",
				cpu : c,
				mem : v.Total, 
				disk: d.Total,
				}
		}
	})
	
	return n
}