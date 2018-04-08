package main

import(
	"fmt"
	"os"
	"os/exec"
	"math/rand"
	"strconv"
	_"net"
	_"strings"
)


//check that wget is present
//check that python library is present
func check_depends(){

}


//function that randomly generates a width and height and then fetches the image.
// saves images as image.png
func fetch_image(){
	fmt.Println("Fetching image")
	rand.Seed(17)

	x := strconv.Itoa(rand.Intn(400))
	y := strconv.Itoa(rand.Intn(400))
	fmt.Printf("X: %s, Y: %s\n",x,y)
	//cmd := []string{}
	cmd := "wget"
	args := []string{"https://picsum.photos/"+x+"/"+y,"-O", "raw.png"}
	if err := exec.Command(cmd,args...).Run(); err != nil{
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
	fmt.Println("Successfully fetched image")
	//process := exec.Command(cmd[0])

}

//function that takes the downloaded image and encode our address
func encode_address(){
	ip_address := getIP()
	fmt.Println("Encoding image w/ "+ip_address)

	cmd := "/usr/local/bin/stegano-lsb"
	//args := []string{"-h"}
	args := []string{"hide", "-i", "raw.png", "-m", "\""+ip_address+"\"", "-o" ,"post.png"}
	// for _,a := range args{
	// 	fmt.Println(a)
	// }
	//fmt.Println(args)
	if err := exec.Command(cmd,args...).Run(); err != nil{
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
	fmt.Println("Successfully encoded image")
}


//function that takes an image and fetches an address.
func decode_address(image string){
	var (
		cmdOut []byte
		err    error
	)


	fmt.Println("Decoding image: "+image)

	cmd := "/usr/local/bin/stegano-lsb"

	args := []string{"reveal", "-i",image}
	if cmdOut,err = exec.Command(cmd,args...).Output(); err != nil{
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
	// if err := exec.Command(cmd,args...).Run(); err != nil{
	// 	fmt.Fprintln(os.Stderr,err)
	// 	os.Exit(1)
	// }
	peer_ip := string(cmdOut)
	//fmt.Fprintln(os.Stdin)
	fmt.Println("Successfully decoded image: "+image+" found: "+peer_ip)
	//cmd := "stegano-lsb"
}






// func getIP()(string){
//     conn, err := net.Dial("udp", "8.8.8.8:80")
//     // handle err...
// 	if err != nil{
// 		fmt.Println("Unable to get IP. No Network")
// 		os.Exit(1)
// 	}
//      defer conn.Close()
// 	 localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	 localAddr_string := strings.Split(localAddr.String(),":")
// 	 return localAddr_string[0]
// }