package awsmeta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const AwsMetaIP = "169.254.169.254"

func buildUrl(version, key string) string {
	if key[0] == '/' {
		key = key[1:]
	}

	return fmt.Sprintf("http://%s/%s/%s", AwsMetaIP, version, key)
}

// retrieve a metadata key for a specific version of the instance metadata api.
// if version is empty, the "latest" version is used.
func GetVersion(version, key string) ([]byte, error) {
	if version == "" {
		version = "latest"
	}
	url := buildUrl(version, key)

	res, er := http.Get(url)
	if er != nil {
		return nil, er
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// see GetVersion().
func GetVersionString(version, key string) (string, error) {
	bytes, er := GetVersion(version, key)
	return string(bytes), er
}

// equivalent to GetVersion("latest", key)
func Get(key string) ([]byte, error) {
	return GetVersion("latest", key)
}

// see Get().
func GetString(key string) (string, error) {
	return GetVersionString("latest", key)
}

func AmiId() (string, error) {
	return GetString("meta-data/ami-id")
}

func PublicIPv4() (string, error) {
	return GetString("meta-data/public-ipv4")
}

func LocalIPv4() (string, error) {
	return GetString("meta-data/local-ipv4")
}

func LocalHostname() (string, error) {
	return GetString("meta-data/local-hostname")
}

func InstanceId() (string, error) {
	return GetString("meta-data/instance-id")
}

func InstanceType() (string, error) {
	return GetString("meta-data/instance-type")
}

func UserData() ([]byte, error) {
	return Get("user-data")
}

func UserDataJson(out interface{}) error {
	bytes, er := UserData()
	if er != nil {
		return er
	}

	return json.Unmarshal(bytes, out)
}
