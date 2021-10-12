package resolve

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "math/big"
   "net/http"
   "strings"
)

func hexEncode(id string) string {
   b62 := new(big.Int)
   b62.SetString(id, 62)
   return hex.EncodeToString(b62.Bytes())
}

type resolve struct {
   AccessPoint []string
}

func newResolve() (*resolve, error) {
   res, err := http.Get("http://apresolve.spotify.com?type=accesspoint")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sol := new(resolve)
   if err := json.NewDecoder(res.Body).Decode(sol); err != nil {
      return nil, err
   }
   return sol, nil
}

func (r resolve) https() string {
   for _, ap := range r.AccessPoint {
      if strings.HasSuffix(ap, ":443") {
         return ap
      }
   }
   return ""
}

func (r resolve) track(id string) error {
   addr := "https://" + r.https() + "/metadata/4/track/" + id
   fmt.Println("GET", addr)
   _, err := http.Get(addr)
   return err
}
