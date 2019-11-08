package ssl

import (
	. "home.rivage.tk/gitea/shunf4/filestash/server/common"
	"path/filepath"
	"os"
)

var keyPEMPath  string = filepath.Join(GetCurrentDir(), CERT_PATH, "key.pem")
var certPEMPath string = filepath.Join(GetCurrentDir(), CERT_PATH, "cert.pem")

func init() {
	os.MkdirAll(filepath.Join(GetCurrentDir(), CERT_PATH), os.ModePerm)
}

func Clear() {
	clearPrivateKey()
	clearCert()
}
