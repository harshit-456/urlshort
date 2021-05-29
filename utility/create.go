package utility



import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/speps/go-hashids"
	
)
const	projectid string="restapi-604cf"
//const	collectionName string ="posts"

func CreateEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var url MyUrl
	err := json.NewDecoder(req.Body).Decode(&url)
	fmt.Println(url)
	//array declared without size is called slice,slice helps to make array varaiable length and also provides other utuilities
	if err != nil {
		fmt.Println("error in reading decoded body")
		return
	}

	ctx:=context.Background()//context needs to be passed when creating firestore client
	client,err:=firestore.NewClient(ctx,projectid)

	if err != nil {
		log.Fatalf("failed to create a firestore  client:%v", err)
		return //\v format specifier for error
	}
	defer client.Close()
	const conn string = "urlshort"
	iterator := client.Collection(conn).Documents(ctx)

	var longurl MyUrl = MyUrl{}

	for {
		doc, err := iterator.Next()
		if doc == nil {
			break
		}

		if err != nil {
			break

		}

		fmt.Println("1")
		if doc.Data() != nil {
			longurl = MyUrl{
				ID:       doc.Data()["ID"].(string),
				LongUrl:  doc.Data()["LongUrl"].(string),
				ShortUrl: doc.Data()["ShortUrl"].(string),
			}
		}
		if longurl.LongUrl == url.LongUrl {
			break
		}
		longurl = MyUrl{}
	}

	if (longurl == MyUrl{}) {
		fmt.Println("5")
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)

		now := time.Now()
		url.ID, _ = h.Encode([]int{int(now.Unix())}) //we will be hashing the timestamp,becoz thats unique
		url.ShortUrl = "http://localhost:1235/" + url.ID
		//bucket.Insert(url.ID,url,0)
		_, _, err = client.Collection(conn).Add(ctx, map[string]interface{}{
			"ID":       url.ID,
			"LongUrl":  url.LongUrl,
			"ShortUrl": url.ShortUrl,
		})
		if(err!=nil){
			log.Fatalf("failed to create a shorturl:%v", err)
		return //\v format specifier for error
		}
	} else {
		url = longurl
	}

	
	json.NewEncoder(w).Encode(url)
}
// to run it type http://localhost:1235/create/
//   make a body

/*
eg {
	"longurl":"www.google.com"
}*/