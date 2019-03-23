package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var EP string
var AK string
var SK string

func init()  {
	EP = ""
	AK = ""
	SK = ""
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("oss client error")
		return false
	}

	bucket, err := client.Bucket(bn)

	if err != nil {
		log.Printf("oss bucket error")
		return false
	}

	err = bucket.UploadFile(filename, path, 500 * 1024, oss.Routines(3))

	if err != nil {
		log.Printf("oss upload error")
		return false
	}

	return true
}