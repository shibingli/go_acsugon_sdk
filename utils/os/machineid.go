package os

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"os"

	"github.com/panta/machineid"
)

func MachineID(appID string) (id string, err error) {
	var mid string
	mid, err = machineid.ID()
	if err != nil {
		return
	}

	var hostname string
	hostname, err = os.Hostname()
	if err != nil {
		return
	}

	mac := hmac.New(md5.New, []byte(mid))
	mac.Write([]byte(appID + ":" + hostname))

	id = fmt.Sprintf("%x", mac.Sum(nil))
	return
}
