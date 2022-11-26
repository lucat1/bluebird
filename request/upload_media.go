package request

import (
	"bytes"
	"fmt"
	"mime/multipart"
)

func UploadMedia(media []byte) (err error) {
	myurl, err := buildURL(NewRequest("media/upload.json"))
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("media", "test")
	part.Write(media)
	w.Close()

	requestPostRaw[mediaResponse](client, myurl, buf, "application/x-www-form-urlencoded")
	fmt.Println("uploaded")
	return
}
