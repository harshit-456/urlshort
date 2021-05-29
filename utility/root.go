package utility
import (
	"context"	
	

	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"cloud.google.com/go/firestore"

	

)

func RootEndPoint(w http.ResponseWriter,req *http.Request){
	//w.Header().Set("Content-type","application/json")

	params:=mux.Vars(req)//url parameter
	var url MyUrl
	id:=params["id"]//params is mapof string to string
	ctx:=context.Background()
	client,err:=firestore.NewClient(ctx,projectid)
	if(err !=nil){
		log.Fatalf("failed to create a firestore  client:%v",err)
		return //\v format specifier for error
	}
	defer client.Close()
	const conn string ="urlshort"
iterator:=client.Collection(conn).Documents(ctx)

for{
	doc,err:=iterator.Next()
	if(err!=nil){
		fmt.Println("error",err)
		return
	}
	if(doc.Data()["ID"].(string)==id){
		url=MyUrl{
			ID: doc.Data()["ID"].(string),
		LongUrl: doc.Data()["LongUrl"].(string),
		ShortUrl:doc.Data()["ShortUrl"].(string),
		  }
		break
	}
}
http.Redirect(w,req,url.LongUrl,301)

 }
 // to try it simple copy shorturl from any document eg http://localhost:1235/kwpplgx and paste on google chrome