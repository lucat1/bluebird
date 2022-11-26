package request

import (
	"bytes"
	"mime/multipart"
)

func UploadMedia(media []byte) (result MediaResponse, err error) {
	myurl, err := buildURL(NewRequest("media/upload.json"))
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("media", "test")
	part.Write(media)
	w.Close()

	result, err = requestPostRaw[MediaResponse](client, myurl, buf, "application/x-www-form-urlencoded", true)
	return
}
