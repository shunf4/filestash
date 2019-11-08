package plugin

import (
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_starter_tunnel"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_starter_tor"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_image_light"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_backblaze"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_dav"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_mysql"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_ftps"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_ldap"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_backend_dropbox"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_security_scanner"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_security_svg"
	_ "home.rivage.tk/gitea/shunf4/filestash/server/plugin/plg_handler_console"
	. "home.rivage.tk/gitea/shunf4/filestash/server/common"
)

func init() {
	Log.Debug("Plugin loader")
}
