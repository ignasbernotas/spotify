package crypto

import (
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "log"
   "net/http"
   "net/url"
   "strings"
   "sync"
)

// connectInfo stores the information about Spotify Connect connection
type connectInfo struct {
	DeviceID  string `json:"deviceID"`
	PublicKey string `json:"publicKey"`
}

type connectDeviceMdns struct {
	Path string
	Name string
}

// Discovery stores the information about Spotify Connect Discovery Request
type Discovery struct {
   keys       PrivateKeys
   cachePath  string
   loginBlob  BlobInfo
   deviceId   string
   deviceName string
   httpServer  *http.Server
   devices     []connectDeviceMdns
   devicesLock sync.RWMutex
}

func (d *Discovery) DeviceId() string {
	return d.deviceId
}

func (d *Discovery) DeviceName() string {
	return d.deviceName
}

func (d *Discovery) LoginBlob() BlobInfo {
	return d.loginBlob
}

func (d *Discovery) Devices() []connectDeviceMdns {
	res := make([]connectDeviceMdns, 0, len(d.devices))
	return append(res, d.devices...)
}

func (d *Discovery) ConnectToDevice(address string) {
	resp, err := http.Get(address + "?action=connectGetInfo")
	resp, err = http.Get(address + "?action=resetUsers")
	resp, err = http.Get(address + "?action=connectGetInfo")

	fmt.Println("start get")
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	info := connectInfo{}
	err = decoder.Decode(&info)
	if err != nil {
		panic("bad json")
	}
	fmt.Println("resposne", resp)

	client64 := base64.StdEncoding.EncodeToString(d.keys.PubKey())
	blob, err := d.loginBlob.MakeAuthBlob(info.DeviceID,
		info.PublicKey, d.keys)
	if err != nil {
		panic("bad blob")
	}

	body := makeAddUserRequest(d.loginBlob.Username, blob, client64, d.deviceId, d.deviceName)
	resp, err = http.PostForm(address, body)
	defer resp.Body.Close()
	decoder = json.NewDecoder(resp.Body)
	var f interface{}
	err = decoder.Decode(&f)

	fmt.Println("got", f, resp, err)
}

func makeAddUserRequest(username string, blob string, key string, deviceId string, deviceName string) url.Values {
	v := url.Values{}
	v.Set("action", "addUser")
	v.Add("userName", username)
	v.Add("blob", blob)
	v.Add("clientKey", key)
	v.Add("deviceId", deviceId)
	v.Add("deviceName", deviceName)
	return v
}

func findCpath(info []string) string {
	for _, i := range info {
		if strings.Contains(i, "CPath") {
			return strings.Split(i, "=")[1]
		}
	}
	return ""
}

func (d *Discovery) handleAddUser(r *http.Request) error {
	//already have login info, ignore
	if d.loginBlob.Username != "" {
		return nil
	}

	username := r.FormValue("userName")
	client64 := r.FormValue("clientKey")
	blob64 := r.FormValue("blob")

	if username == "" || client64 == "" || blob64 == "" {
		log.Println("Bad Request, addUser")
		return errors.New("bad username Request")
	}

	blob, err := NewBlobInfo(blob64, client64, d.keys,
		d.deviceId, username)
	if err != nil {
		return errors.New("failed to decode blob")
	}

	err = blob.SaveToFile(d.cachePath)
	if err != nil {
		log.Println("failed to cache login info")
	}

	d.loginBlob = blob
	return nil
}
