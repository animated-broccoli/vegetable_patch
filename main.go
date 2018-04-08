package main


import (
	"fmt"
	"html"
	"log"
	"net/http"
	//"github.com/animated-broccoli/announce"
	"github.com/gorilla/mux"
)





// Function to serve raw http calls
func serve(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080",nil))

}


// Function to serve rest calls
func rest(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",Index)
	log.Fatal(http.ListenAndServe(":8080",router))
}


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))
}


func printStats(){
	fmt.Printf("ID: %s\n",n.Get_id())
	fmt.Printf("IP: %s\n",n.Get_ip())
	fmt.Printf("Job: %s\n",n.Get_job())
	fmt.Printf("Memory: %d\n",n.Get_mem())
	fmt.Printf("Disk: %d\n",n.Get_disk())
	fmt.Printf("CPU: %d\n",n.Get_cpu())
}

func main(){
	//node{Name: "Write presentation"}
	n = Node()
	n.Set_job("Worker")
	printStats()
	//fmt.Printf("Gardens Galore!\n")
	//serve()
	
	//rest()
	//fetch_image()
	//encode_address()
	//decode_address("post.png")
}