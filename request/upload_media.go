package request

import (
	"bytes"
	"io"
	"mime/multipart"
)

func UploadMedia(media []byte) (result MediaResponse, err error) {
	myurl, err := buildURL(NewRequest("media/upload.json?media_category=tweet_image"))
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("media", "image.png")
	io.Copy(part, bytes.NewBuffer(media))
	w.Close()

	result, err = requestPostRaw[MediaResponse](client, myurl, buf, w.FormDataContentType())
	return
}
