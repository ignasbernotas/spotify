package crypto

import (
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "log"
   "net"
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

// connectGetInfo stores the information about a Spotify Connect information Request
type connectGetInfo struct {
	Status           int    `json:"status"`
	StatusError      string `json:"statusError"`
	SpotifyError     int    `json:"spotifyError"`
	Version          string `json:"version"`
	DeviceID         string `json:"deviceID"`
	RemoteName       string `json:"remoteName"`
	ActiveUser       string `json:"activeUser"`
	PublicKey        string `json:"publicKey"`
	DeviceType       string `json:"deviceType"`
	LibraryVersion   string `json:"libraryVersion"`
	AccountReq       string `json:"accountReq"`
	BrandDisplayName string `json:"brandDisplayName"`
	ModelDisplayName string `json:"modelDisplayName"`
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

// makeConnectGetInfo builds a connectGetInfo structure with the provided values
func makeConnectGetInfo(deviceId string, deviceName string, publicKey string) connectGetInfo {
	return connectGetInfo{
		Status:           101,
		StatusError:      "ERROR-OK",
		SpotifyError:     0,
		Version:          "1.3.0",
		DeviceID:         deviceId,
		RemoteName:       deviceName,
		ActiveUser:       "",
		PublicKey:        publicKey,
		DeviceType:       "UNKNOWN",
		LibraryVersion:   "0.1.0",
		AccountReq:       "PREMIUM",
		BrandDisplayName: "librespot",
		ModelDisplayName: "librespot",
	}
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

func (d *Discovery) startHttp(done chan int, l net.Listener) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		action := r.FormValue("action")
		fmt.Println("got Request: ", action)
		switch {
		case "connectGetInfo" == action || "resetUsers" == action:
			client64 := base64.StdEncoding.EncodeToString(d.keys.PubKey())
			info := makeConnectGetInfo(d.deviceId, d.deviceName, client64)

			js, err := json.Marshal(info)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case "addUser" == action:
			err := d.handleAddUser(r)
			if err == nil {
				done <- 1
			}
		}
	})

	d.httpServer = &http.Server{}
	err := d.httpServer.Serve(l)
	if err != nil {
		fmt.Println("got an error", err)
	}
}
