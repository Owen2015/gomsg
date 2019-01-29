 package main

 import (
	 "net/http"
	 "github.com/gorilla/websocket"
	"html/template"
//	"fmt"
)

var upgrader=websocket.Upgrader{}


func main(){
	var message []byte
	msg_ch:=make(chan []byte)
	msg_chs:=[] chan []byte(nil)

  http.HandleFunc("/",indexHandler) 

   http.HandleFunc("/sender",func(w http.ResponseWriter, r *http.Request){
	   var conn, _ = upgrader.Upgrade(w,r,nil) 
	   go func(conn *websocket.Conn){
		   ch_i:=make(chan []byte)
		   msg_chs=append(msg_chs,ch_i)
		  first:=true	
		   for{
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
		   }
	   }(conn)

   })

go func () {
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
	t,_:=template.ParseFiles("/go/src/github.com/owen2015/gomsg/web/index.html")
	t.Execute(w,nil)
}

