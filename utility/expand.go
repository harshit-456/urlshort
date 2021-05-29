package utility

import (
	"context"	
	
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"

	
)


func ExpandEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	url := req.URL.Query() //to read it from url
	shorturl := url.Get("ShortUrl")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectid)
	if err != nil {
		log.Fatalf("failed to create a firestore  client:%v", err)
		return //\v format specifier for error
	}
	defer client.Close()
	const conn string = "urlshort"
	iterator := client.Collection(conn).Documents(ctx)
	var res MyUrl
	for {
		doc, err := iterator.Next()
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		if doc.Data()["ShortUrl"].(string) == shorturl {
			res = MyUrl{
				ID:       doc.Data()["ID"].(string),
				LongUrl:  doc.Data()["LongUrl"].(string),
				ShortUrl: doc.Data()["ShortUrl"].(string),
			}
			break
		}
	}
	json.NewEncoder(w).Encode(res)
}
//to run it type http://localhost:1235/expand/?ShortUrl=your id 