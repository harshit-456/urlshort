package utility
type MyUrl struct{
	ID string `json:"id"`
	//json mapping to know how it looks in server whatver we define in json will be reflected or seen in database
	LongUrl string `json:"longurl"`
	ShortUrl string `json:"shorturl"`
}