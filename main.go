 package main

 import (
	 "net/http"
	 "github.com/gorilla/websocket"
//	 "time"
	"html/template"
	"fmt"
)

var upgrader=websocket.Upgrader{}

func fanout(in chan []byte, out [] chan []byte) {
//	temp:=make([] byte,10)

	for{
		temp:=<-in
		println("fanout")
		println(len(out))
		for i:=0;i<len(out);i++ {
			println(i)
			println(string(temp))
			out[i]<-temp
			select {
			case out[i]<-temp:
				fmt.Printf("every thing is fine.")
			default:
				fmt.Printf("one of client down.")
			}		
		}
	}
}

func main(){
	 //message:="null"
	//message:=make([]byte)
	var message []byte
	msg_ch:=make(chan []byte)
//	msg_chs:=make([] chan []byte,0)
	msg_chs:=[] chan []byte(nil)

  http.HandleFunc("/",indexHandler) 

   http.HandleFunc("/sender",func(w http.ResponseWriter, r *http.Request){
	   var conn, _ = upgrader.Upgrade(w,r,nil) 
	   go func(conn *websocket.Conn){
		   ch_i:=make(chan []byte)
		   msg_chs=append(msg_chs,ch_i)
		  first:=true	
		   for{
			 // println("handle Func /sender")
			 if first {
			  conn.WriteMessage(websocket.TextMessage,message)
			  first=false
			 } else {
			  msg_test:=<-ch_i
			  conn.WriteMessage(websocket.TextMessage,msg_test)
		 	 }
		  }
	   }(conn)
   })

 http.HandleFunc("/receiver",func(w http.ResponseWriter, r *http.Request){
	   var conn, _ = upgrader.Upgrade(w,r,nil)
	   go func(conn *websocket.Conn){
		   for{
			   _,msg,err :=conn.ReadMessage()
			   if err!= nil{
				   conn.Close()
				   return 
			   }
			   msg_ch<-msg
			  // println(string(msg))
		   }
	   }(conn)

   })

//  go fanout(msg_ch,msg_chs)
go func () {
//	temp:=make([] byte,10)
	for{
		message=<-msg_ch
		for i:=0;i<len(msg_chs);i++ {
			msg_chs[i]<-message
		}
	}
}()

   http.ListenAndServe(":3000",nil)  
}   

func indexHandler(w http.ResponseWriter, r *http.Request){
	t,_:=template.ParseFiles("src/github.com/owen2015/gomsg/web/index.html")
	t.Execute(w,nil)
}

