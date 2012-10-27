package awsmeta

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const AwsMetaIP = "169.254.169.254"

func buildUrl(key string) string {
	if key[0] != '/' {
		key = "/" + key
	}

	return "http://" + AwsMetaIP + "/latest" + key 
}

func getData(key string) ([]byte, error) {
	url := buildUrl(key)

	res, er := http.Get(url)
	if er != nil {
		return nil, er
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func getDataAsString(key string) (string, error) {
	bytes, er := getData(key)

	if er != nil {
		return "", er
	}

	return string(bytes), nil
}

func AmiId() (string, error) {
	return getDataAsString("meta-data/ami-id")
}

func PublicIPv4() (string, error) {
	return getDataAsString("meta-data/public-ipv4")
}

func LocalIPv4() (string, error) {
	return getDataAsString("meta-data/local-ipv4")
}

func LocalHostname() (string, error) {
	return getDataAsString("meta-data/local-hostname")
}

func InstanceId() (string, error) {
	return getDataAsString("meta-data/instance-id")
}

func InstanceType() (string, error) {
	return getDataAsString("meta-data/instance-type")
}

func UserData() ([]byte, error) {
	return getData("user-data")
}

func UserDataJson(out interface{}) error {
	bytes, er := UserData()
	if er != nil {
		return er
	}

	return json.Unmarshal(bytes, out)
}
