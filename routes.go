package main

import (
	"fmt"
	"html"
	"os"
	"net/http"
	_"github.com/gorilla/mux"
	"encoding/json"

)


// type ResourceMSG struct{
// 	Func string `json:"func,omitempty`
// 	Id	string `json:"id,omitempty`
// 	Ip string `json:"ip,omitempty`
// 	Job string `json:"job,omitempty`
// 	CPU uint64 `json:"cpu,omitempty`
// 	Mem uint64 `json:"mem,omitempty`
// 	Disk uint64 `json:"disk,omitempty`
// 	//mu	sync.RWMutex
// }

func PeerResponse()(ResourceMSG){
	var res ResourceMSG
	res = ResourceMSG{
		Func: "peer_response" ,
		Id: n.Get_id(),
		Ip: n.Get_ip(),
		Job: n.Get_job(),
		CPU : n.Get_cpu(),
		Mem : n.Get_mem(), 
		Disk: n.Get_disk(),
		}

	return res
}

func Register(w http.ResponseWriter, r *http.Request){
	var rm ResourceMSG
	err := json.NewDecoder(r.Body).Decode(&rm)
	if err != nil{
		fmt.Printf("ERROR: ")
		fmt.Printf(err.Error())
		fmt.Printf("\n")
		os.Exit(1)
	}

	fmt.Printf("Received: %s\n",rm.Id)
	fmt.Printf("Received: %s\n",rm.Ip)
	fmt.Printf("Received: %s\n",rm.Job)
	fmt.Printf("Received: %d\n",rm.CPU)
	fmt.Printf("Received: %d\n",rm.Mem)
	fmt.Printf("Received: %d\n",rm.Disk)


	//fmt.Printf("Sending: %s\n",n.id)
	json.NewEncoder(w).Encode(PeerResponse())
}


func Init(w http.ResponseWriter, r *http.Request){
	//params := mux.Vars(r)
	
		
	n.Set_job("worker")
	printStats()
	fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))

}


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))
}