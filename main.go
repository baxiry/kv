package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/kardianos/osext"
	"github.com/shirou/gopsutil/mem"
	"github.com/tidwall/gjson"

	"bline/hashmap"

	"github.com/panjf2000/ants"

	talkservice "bline/linethrift"
	"bline/oop"
)

type (
	mentions struct {
		MENTIONEES []struct {
			Start string `json:"S"`
			End   string `json:"E"`
			Mid   string `json:"M"`
		} `json:"MENTIONEES"`
	}
	kickop struct {
		Kick  []string
		Inv   []string
		Opinv []int64
	}
	Stickers struct {
		Id  string
		Pid string
	}
	clustering struct {
		mem string
		tm  int64
		fr  []string
	}
	DATA struct {
		Authoken      []string             `json:"Authoken"`
		CreatorBack   []string             `json:"CreatorBack"`
		BuyerBack     []string             `json:"BuyerBack"`
		OwnerBack     []string             `json:"OwnerBack"`
		MasterBack    []string             `json:"MasterBack"`
		AdminBack     []string             `json:"AdminBack"`
		ResponBack    string               `json:"ResponBack"`
		RnameBack     string               `json:"RnameBack"`
		SnameBack     string               `json:"SnameBack"`
		BotBack       []string             `json:"BotBack"`
		Dalltime      string               `json:"Dalltime"`
		Logobot       string               `json:"Logobot"`
		SellerBack    []string             `json:"SellerBack"`
		BanBack       []string             `json:"BanBack"`
		FuckBack      []string             `json:"FuckBack"`
		Limit         string               `json:"MLimit"`
		Fresh         string               `json:"MFfresh"`
		MuteBack      []string             `json:"MuteBack"`
		AnnunceBack   []string             `json:"AnnunceBack"`
		ProQrBack     []string             `json:"ProQrBack"`
		ProjoinBack   []string             `json:"ProjoinBack"`
		ProInviteBack []string             `json:"ProInviteBack"`
		ProCancelBack []string             `json:"ProCancelBack"`
		ProkickBack   []string             `json:"ProkickBack"`
		GbanBack      map[string][]string  `json:"GbanBack"`
		GadminBack    map[string][]string  `json:"GadminBack"`
		GownerBack    map[string][]string  `json:"GownerBack"`
		TimeBanBack   map[string]time.Time `json:"TimeBanBack"`
	}
)

func Resprem() {
	rngcmd := GetComs(3, "joinqrkick")
	rngcmd = GetComs(3, "/squadbot")
	rngcmd = GetComs(4, "joinqr")
	rngcmd = GetComs(3, "cancelall")
	rngcmd = GetComs(3, "kickall")
	rngcmd = GetComs(7, "none")
	rngcmd = GetComs(7, "max")
	rngcmd = GetComs(6, "allowall")
	rngcmd = GetComs(6, "denyall")
	rngcmd = GetComs(5, "hostage")
	rngcmd = GetComs(3, "backup")
	rngcmd = GetComs(5, "upgname")
	rngcmd = GetComs(5, "welcome")
	rngcmd = GetComs(4, "sendimage")
	rngcmd = GetComs(4, "leave")
	rngcmd = GetComs(5, "announce")
	rngcmd = GetComs(5, "unban")
	rngcmd = GetComs(6, "bio")
	rngcmd = GetComs(6, "tag")
	rngcmd = GetComs(6, "image")
	rngcmd = GetComs(5, "contact")
	rngcmd = GetComs(5, "ban")
	rngcmd = GetComs(6, "kick")
	rngcmd = GetComs(6, "vkick")
	rngcmd = GetComs(6, "invite")
	rngcmd = GetComs(6, "cancel")
	rngcmd = GetComs(7, "ungban")
	rngcmd = GetComs(3, "unbot")
	rngcmd = GetComs(8, "tagall")
	rngcmd = GetComs(4, "statusall")
	rngcmd = GetComs(5, "status")
	rngcmd = GetComs(5, "whois")
	rngcmd = GetComs(5, "mute")
	rngcmd = GetComs(4, "fuck")
	rngcmd = GetComs(3, "setlimiter")
	rngcmd = GetComs(3, "setcancel")
	rngcmd = GetComs(3, "setkick")
	rngcmd = GetComs(3, "msglimit")
	rngcmd = GetComs(3, "msgstatus")
	rngcmd = GetComs(3, "msglurk")
	rngcmd = GetComs(4, "msgleave")
	rngcmd = GetComs(4, "speed")
	rngcmd = GetComs(6, "lurk")
	rngcmd = GetComs(4, "msgwelcome")
	rngcmd = GetComs(3, "msgrespon")
	rngcmd = GetComs(3, "setrname")
	rngcmd = GetComs(3, "setsname")
	rngcmd = GetComs(3, "logmode")
	rngcmd = GetComs(4, "killmode")
	rngcmd = GetComs(3, "unowner")
	rngcmd = GetComs(6, "name")
	rngcmd = GetComs(3, "Stats")
	rngcmd = GetComs(3, "buyers")
	rngcmd = GetComs(3, "upname")
	rngcmd = GetComs(3, "upstatus")
	rngcmd = GetComs(3, "acceptall")
	rngcmd = GetComs(3, "declineall")
	rngcmd = GetComs(6, "abort")
	rngcmd = GetComs(3, "accept")
	rngcmd = GetComs(3, "decline")
	rngcmd = GetComs(3, "invme")
	rngcmd = GetComs(3, "gleave")
	rngcmd = GetComs(4, "Purgeallbans")
	rngcmd = GetComs(4, "purgeall")
	rngcmd = GetComs(5, "unsend")
	rngcmd = GetComs(1, "creators")
	rngcmd = GetComs(3, "upvcover")
	rngcmd = GetComs(1, "unseller")
	rngcmd = GetComs(1, "clearseller")
	rngcmd = GetComs(2, "sellers")
	rngcmd = GetComs(1, "seller")
	rngcmd = GetComs(0, "uncreator")
	rngcmd = GetComs(0, "clearcreator")
	rngcmd = GetComs(3, "upvimage")
	rngcmd = GetComs(3, "upcover")
	rngcmd = GetComs(3, "upimage")
	rngcmd = GetComs(2, "clearbuyer")
	rngcmd = GetComs(2, "unbuyer")
	rngcmd = GetComs(2, "buyer")
	rngcmd = GetComs(7, "gaccess")
	rngcmd = GetComs(4, "allbanlist")
	rngcmd = GetComs(3, "access")
	rngcmd = GetComs(7, "expel")
	rngcmd = GetComs(4, "listcmd")
	rngcmd = GetComs(3, "owner")
	rngcmd = GetComs(3, "hide")
	rngcmd = GetComs(3, "unhide")
	rngcmd = GetComs(3, "hidelist")
	rngcmd = GetComs(3, "clearhide")
	rngcmd = GetComs(6, "mid")
	rngcmd = GetComs(4, "cleargowner")
	rngcmd = GetComs(3, "logmode")
	rngcmd = GetComs(3, "clearowner")
	rngcmd = GetComs(4, "unmaster")
	rngcmd = GetComs(5, "unmute")
	rngcmd = GetComs(4, "setcmd")
	rngcmd = GetComs(6, "gowner")
	rngcmd = GetComs(4, "master")
	rngcmd = GetComs(4, "gojoin")
	rngcmd = GetComs(5, "ungowner")
	rngcmd = GetComs(5, "settings")
	rngcmd = GetComs(5, "set")
	rngcmd = GetComs(4, "runtime")
	rngcmd = GetComs(4, "timenow")
	rngcmd = GetComs(3, "timeleft")
	rngcmd = GetComs(8, "say")
	rngcmd = GetComs(5, "curl")
	rngcmd = GetComs(5, "ourl")
	rngcmd = GetComs(8, "here")
	rngcmd = GetComs(7, "gbanlist")
	rngcmd = GetComs(4, "clearcache")
	rngcmd = GetComs(4, "clearchat")
	rngcmd = GetComs(4, "cleargban")
	rngcmd = GetComs(3, "clearbot")
	rngcmd = GetComs(3, "botlist")
	rngcmd = GetComs(5, "bans")
	rngcmd = GetComs(4, "fixed")
	rngcmd = GetComs(7, "gban")
	rngcmd = GetComs(3, "bot")
	rngcmd = GetComs(4, "stay")
	rngcmd = GetComs(3, "leaveall")
	rngcmd = GetComs(4, "go")
	rngcmd = GetComs(5, "stayall")
	rngcmd = GetComs(4, "bringall")
	rngcmd = GetComs(3, "listprotect")
	rngcmd = GetComs(6, "cleargadmin")
	rngcmd = GetComs(3, "clearban")
	rngcmd = GetComs(3, "clearadmin")
	rngcmd = GetComs(2, "upallname")
	rngcmd = GetComs(3, "upallstatus")
	rngcmd = GetComs(5, "limitout")
	rngcmd = GetComs(5, "sayall")
	rngcmd = GetComs(6, "count")
	rngcmd = GetComs(8, "ping")
	rngcmd = GetComs(4, "leave")
	rngcmd = GetComs(1, "addallsquads")
	rngcmd = GetComs(2, "addallbots")
	rngcmd = GetComs(4, "limits")
	rngcmd = GetComs(3, "adds")
	rngcmd = GetComs(3, "friends")
	rngcmd = GetComs(2, "upvallcover")
	rngcmd = GetComs(2, "upvallimage")
	rngcmd = GetComs(5, "unsend")
	rngcmd = GetComs(2, "upallcover")
	rngcmd = GetComs(2, "upallimage")
	rngcmd = GetComs(5, "rollcall")
	rngcmd = GetComs(6, "respon")
	rngcmd = GetComs(5, "banlist")
	rngcmd = GetComs(4, "antitag")
	rngcmd = GetComs(6, "admins")
	rngcmd = GetComs(8, "gadmin")
	rngcmd = GetComs(4, "squadmid")
	rngcmd = GetComs(7, "ungadmin")
	rngcmd = GetComs(5, "unadmin")
	rngcmd = GetComs(5, "masters")
	rngcmd = GetComs(5, "gowners")
	rngcmd = GetComs(5, "admin")
	rngcmd = GetComs(4, "unfuck")
	rngcmd = GetComs(3, "remote")
	rngcmd = GetComs(3, "groupinfo")
	rngcmd = GetComs(3, "banpurge")
	rngcmd = GetComs(3, "autoban")
	rngcmd = GetComs(3, "autopurge")
	rngcmd = GetComs(3, "canceljoin")
	rngcmd = GetComs(3, "nukejoin")
	rngcmd = GetComs(3, "lagjoin")
	rngcmd = GetComs(3, "kicktime")
	rngcmd = GetComs(3, "groups")
	rngcmd = GetComs(3, "gourl")
	rngcmd = GetComs(3, "groupcast")
	rngcmd = GetComs(4, "fucklist")
	rngcmd = GetComs(5, "mutelist")
	rngcmd = GetComs(3, "autojoin")
	rngcmd = GetComs(3, "perm")
	rngcmd = GetComs(3, "permlist")
	rngcmd = GetComs(3, "clearallprotect")
	rngcmd = GetComs(3, "clearmute")
	rngcmd = GetComs(3, "clearfuck")
	rngcmd = GetComs(3, "clearmaster")
	fmt.Println(rngcmd)
}

func abort() {
	remotegrupidto = ""
	StartChangeImg = false
	StartChangevImg = false
	StartChangevImg2 = false
	Sinderremote = []string{}
	remotegrupid = ""
	changepic = []*oop.Account{}
	ChangName = false
	ChangCover = false
	ChangPict = false
	ChangeBio = false
	ChangVpict = false
	ChangVcover = false
	AllCheng = false
	MsgBio = ""
	MsgName = ""
	timeabort = time.Now()
}

var (
	Whitelist        = []string{}
	kicktime         = false
	SetHelper        = &oop.Helper{Rngcmd: make(map[string]int)}
	DB               *DATA
	ClientBot        []*oop.Account
	Midlist          []string
	aclear           = time.Now()
	Grupas           []*talkservice.Group
	Poll             *oop.Account
	Self             *oop.Account
	cpu              int
	botleave         = &hashmap.HashMap{}
	changepic        []*oop.Account
	timeabort        = time.Now()
	TimeSave         = time.Now()
	Laststicker      = &hashmap.HashMap{}
	ChangCover       = false
	MsgRespon        = "𝐈'𝐦 𝐇𝐞𝐫𝐞 𝐁𝐨𝐬𝐬..."
	MsgBan           = "Cleared %v blacklist."
	MsFresh          = " ✓"
	MsLimit          = " ✘"
	MsSname          = "."
	MsRname          = "bot"
	AllCheng         = false
	Lastleave        = &hashmap.HashMap{}
	ChangPict        = false
	ChangName        = false
	AutokickBan      = false
	ChangVpict       = false
	ChangVcover      = false
	ChangeBio        = false
	CmdHelper        = &hashmap.HashMap{}
	cewel            = &hashmap.HashMap{}
	cleave           = &hashmap.HashMap{}
	NukeJoin         = false
	lagjoin          = false
	AutoBan          = true
	canceljoin       = false
	Autojoin         = "off"
	backlist         = &hashmap.HashMap{}
	cekoptime        = []int64{}
	Ceknuke          = &hashmap.HashMap{}
	Cekstaybot       = &hashmap.HashMap{}
	Commands         = &oop.Command{}
	Waitlistin       = map[string][]string{}
	AutoproN         = false
	LogMode          = false
	LogGroup         = ""
	delayed          = 10 * time.Second
	MsgBio           = ""
	MsgName          = ""
	StartChangeImg   = false
	StartChangevImg  = false
	StartChangevImg2 = false
	AutoPro          = true
	Command          = &hashmap.HashMap{}
	tempginv         = []string{}
	remotegrupidto   = ""
	ModeBackup       = "invite"
	checkHaid        = []string{}
	botStart         = time.Now()
	TimeBackup       = time.Now()
	oplist           = []int64{}
	oplistinvite     = []int64{}
	PurgeOP          = []int64{}
	oplistjoin       = []int64{}
	AutoPurge        = true
	ClientMid        = map[string]*oop.Account{}
	Squadlist        = []string{}
	argsRaw          = os.Args
	Sinderremote     = []string{}
	GO               = argsRaw[1]
	StartChangeVideo = false
	tempgroup        = []string{}
	Lastinvite       = &hashmap.HashMap{}
	Lastkick         = &hashmap.HashMap{}
	Lastjoin         = &hashmap.HashMap{}
	Lastcancel       = &hashmap.HashMap{}
	Nkick            = &hashmap.HashMap{}
	Lastupdate       = &hashmap.HashMap{}
	Lastmid          = &hashmap.HashMap{}
	filterop         = &hashmap.HashMap{}
	Lasttag          = &hashmap.HashMap{}
	Lastcon          = &hashmap.HashMap{}
	Commandss        = &hashmap.HashMap{}
	Detectjoin       = &oop.SaveJoin{}
	Banned           = &oop.BanUser{Banlist: []string{}, Fucklist: []string{}, Mutelist: []string{}}
	UserBot          = &oop.Access{Creator: []string{}, Seller: []string{}, Buyer: []string{}, Owner: []string{}, Master: []string{}, Admin: []string{}, Bot: []string{}}
	timeSend         = []int64{}
	opjoin           = []string{}
	Cekpurge         = []int64{}
	MaxCancel        = 50
	MaxKick          = 50
	cekGo            = []int64{}
	UpdatePicture    = map[string]bool{}
	UpdateCover      = map[string]bool{}
	UpdateVProfile   = map[string]bool{}
	UpdateVCover     = map[string]bool{}
	Qrwar            = false
	filterWar        = &kickop{Kick: []string{}, Inv: []string{}, Opinv: []int64{}}
	ColorCyan        = "\033[36m"
	ColorReset       = "\033[0m"
	Data             DATA
	remotegrupid     = ""
	LastActive       = &hashmap.HashMap{}
	used             = ""
	IPServer         string
	Killmode         = "kill"

	filtermsg   = &hashmap.HashMap{}
	Opinvite    = []int64{}
	stringToInt = []rune("01")
	DATABASE    = "db/" + GO + ".json"
	MAKERS      = []string{"u27603a2c021c18746b7aa34e3d2b2220", "u9ef25059bb2e71a0834fb3b2cadd6297", "u7b566e01279ac3dcf0108e8248b67e41"}
	HostName    = []string{
		"legy-jp-long",
	}
	carierMap = map[string]string{
		"IOSIPAD":     "51089, 1-0",
		"IOS":         "51089, 1-0",
		"ANDROIDLITE": "51000",
		"ANDROID":     "51010",
		"CHROMEOS":    "",
		"DESKTOPMAC":  "",
		"DESKTOPWIN":  "",
		"CHANNELCP":   "51010",
	}
	helppro = []string{
		"Allowall",
		"Allow invite",
		"Allow kick",
		"Allow cancel",
		"Allow join",
		"Allow link",
		"Allow name",
		"Denyall",
		"Deny invite",
		"Deny kick",
		"Deny cancel",
		"Deny join",
		"Deny link",
		"Deny name",
		"Protect max",
		"Protect none",
	}
	ListIp = []string{
		"172.105.226.94",
		"",
	}
	helpmaker = []string{
		"Addallbots",
		"Addallsquads",
		"Addday",
		"Addmonth",
		"Addyear",
		"Appname",
		"Buyer",
		"Creator",
		"Creators",
		"Clearbuyer",
		"Clearcreator",
		"Clearfriend",
		"Clearseller",
		"Expel",
		"Friends",
		"Hostname",
		"Reboot",
		"Unbuyer",
		"Uncreator",
		"Unfriend",
		"Upallcover",
		"Upallimage",
		"Upallname",
		"Upallstatus",
		"Upvallcover",
		"Upvallimage",
		"Useragent",
		"Setdate",
	}
	helpbuyer = []string{
		"About",
		"Access",
		"Accept",
		"Acceptall",
		"Adds",
		"Autoban",
		"Autojoin",
		"Autopro",
		"Autopurge",
		"Banpurge",
		"Backup",
		"Bot",
		"Botlist",
		"Buyers",
		"Cancelall",
		"Canceljoin",
		"Clearallprotect",
		"Clearban",
		"Clearbot",
		"Clearbuyer",
		"Clearowner",
		"Clearmaster",
		"Clearadmin",
		"Clearfuck",
		"Clearhide",
		"Clearmute",
		"Decline",
		"Declineall",
		"Expel",
		"Friends",
		"Gleave",
		"Gourl",
		"Groupcast",
		"Groups",
		"Groupinfo:",
		"Hide",
		"Hidelist",
		"Invme",
		"Joinqrkick:",
		"Kickall",
		"Leaveall",
		"List protect",
		"Msglimit",
		"Msglurk",
		"Msgrespon",
		"Msgstatus",
		"Nukejoin",
		"lagjoin",
		"kicktime",
		"Owner",
		"Perm",
		"Permlist:",
		"Remote",
		"Stats",
		"Setcancel",
		"setlogo",
		"Setkick",
		"Setlimiter",
		"Setrname",
		"Setsname",
		"Statusall",
		"/squadbot",
		"Timeleft",
		"Unbot",
		"Unhide",
		"Unowner",
		"Upcover",
		"Upimage",
		"upgname",
		"Upname",
		"Upstatus",
		"Upvcover",
		"Upvimage",
	}
	helpowner = []string{
		"Allbanlist",
		"Antitag",
		"Bringall",
		"Clearcache",
		"Clearchat",
		"Cleargban",
		"Cleargowner",
		"Squadcontact",
		"Expel",
		"Fixed",
		"Fuck",
		"Fucklist",
		"Joinqr:",
		"Killmode:",
		"Leave",
		"Limits",
		"Limitout",
		"Listcmd",
		"Master",
		"Owners",
		"Purgeall",
		"Purgeallbans:",
		"Runtime",
		"Set",
		"Setcmd",
		"Sendimage",
		"Squadmid",
		"Status",
		"Stay",
		"Timenow",
		"Unfuck",
		"Unmaster",
	}
	helpmaster = []string{
		"Admin",
		"Announce",
		"Ban",
		"Banlist",
		"Bans",
		"Bring",
		"Cleargadmin",
		"Contact",
		"Count",
		"Curl",
		"Expel",
		"Go",
		"Gojoin",
		"Gowner",
		"Hostage",
		"Leave on/off",
		"Masters",
		"Msgleave",
		"Msgwelcome",
		"Mute",
		"Mutelist",
		"Ourl",
		"Rollcall",
		"Sayall",
		"Settings",
		"Speed",
		"Unadmin",
		"Unban",
		"Ungowner",
		"Unmute",
		"Unsend",
		"Welcome on/off",
		"Whois",
	}
	helpadmin = []string{
		"Abort",
		"Admins",
		"Allow all",
		"Allow invite",
		"Allow kick",
		"Allow cancel",
		"Allow join",
		"Allow link",
		"Allow name",
		"Bio",
		"Cancel",
		"Cover",
		"Deny all",
		"Deny invite",
		"Deny kick",
		"Deny cancel",
		"Deny link",
		"Deny name",
		"Expel",
		"Gaccess",
		"Gadmin",
		"Gban",
		"Gbanlist",
		"Gadmins",
		"Here",
		"Image",
		"Invite",
		"Kick",
		"vKick",
		"Lcon",
		"Linvite",
		"Lkick",
		"Lcancel",
		"Lleave",
		"Ljoin",
		"Lmid",
		"Ltag",
		"Lurk name",
		"Lurk mention",
		"Lurk hide",
		"Lurk on/off",
		"Lurks",
		"Mid",
		"Max",
		"None",
		"Name",
		"Ping",
		"Prefix",
		"Protect max",
		"Protect none",
		"Respon",
		"Rname",
		"Sname",
		"Say",
		"Tag",
		"Tagall",
		"Ungadmin",
		"Ungban",
	}
	details = map[string]string{
		"acceptall":       "'%s%s'\n\n.acceptall- accept all group invitations",
		"clearhide":       "'%s%s'\n\n.clearhide- clear hidelist",
		"declineall":      "'%s%s'\n\n.declineall- decline all group invitations",
		"groupinfo":       "'%s%s'\n\n.groupinfo “number”- gives information about that group",
		"list protect":    "'%s%s'\n\n.list protect- shows all protection of all groups",
		"msglurk":         "'%s%s'\n\n.msglurk “message”- The message that appears during lurk command",
		"permlist":        "'%s%s'\n\n.permlist- shows perms of all commands",
		"setlogo":         "'%s%s'\n\n.setlogo “logo”- sets logo for help menu",
		"unbot":           "'%s%s'\n\n.unbot @tag, lcon, lkick, etc- removes from botlist",
		"upgname":         "'%s%s'\n\n.upgname “name”- update group name",
		"limits":          "'%s%s'\n\n.limits- shows bots ban status",
		"status":          "'%s%s'\n\n.status- shows bot status",
		"unfuck":          "'%s%s'\n\n.unfuck @tag, lcon, lkick, etc- removes from fucklist",
		"msgleave":        "'%s%s'\n\n.msgleave “message”- sets leave message",
		"msgwelcome":      "'%s%s'\n\n.msgwelcome “message”- sets welcome message",
		"bio":             "'%s%s'\n\n.bio @tag, lcon, lkick, etc- shows status/bio",
		"vkick":           "'%s%s'\n\n.vkick @tag, lcon, lkick, etc- kicks, invites, then cancels target. Clears chat history for target",
		"prefix":          "'%s%s'\n\nPrefix- shows bots rname and sname",
		"about":           "'%s%s'\n\nbout-shows bot account status.",
		"squadbot":        "'%s%s'\n\n./squadbot- Sends contacts of all bots and says which are on limit",
		"decline":         "'%s%s'\n\ndecline “number” -decline invitation to group.",
		"unhide":          "'%s%s'\n\nunhide @tag, lcon, lkick, etc or “number” -unhides from unhide list.",
		"restartcmd":      "'%s%s'\n\nrestartcmd-clears changed commands from listcmd list.",
		"hidelist":        "'%s%s'\n\nhidelist-shows hidden list.",
		"msglimit":        "'%s%s'\n\n“symbol”- sets symbol shown when checking limit bots",
		"cancelall":       "'%s%s'\n\ncancelall-cancels all pending .",
		"clearallprotect": "'%s%s'\n\nclearallprotect-turns off all protection in all groups.",
		"setkick":         "'%s%s'\n\nsetkick-sets kick limiter .",
		"setcancel":       "'%s%s'\n\nsetcancel-sets cancel limiter.",
		"setlimiter":      "'%s%s'\n\nsetlimiter 'number'- sets limit for kicks and cancels.",
		"listcmd":         "'%s%s'\n\nlistcmd-shows lists of changed commads.",
		"msglimits":       "'%s%s'\n\n msglimits 'message'-when use status command bots have this message next to them    when banned.",
		"msgstatus":       "'%s%s'\n\nmsgstatus 'message'- when use status, this message appears next to fresh/no limit    bots.",
		"remote":          "'%s%s'\n\nremote 'number'- Use grouplist first and find group number. Then use command.    When in remote type the commands you want done remotely. Use abort to stop remote.",
		"joinqrkick":      "'%s%s'\n\njoinqrkick: 'link'- joins qr and kicks on joining.",
		"unfriend":        "'%s%s'\n\nunfriend lkick, ltag, @tag, etc- bots unfriend person.",
		"canceljoin":      "'%s%s'\n\ncanceljoin on/off- bots cancelall immediately when joining.",
		"allbanlist":      "'%s%s'\n\nallbanlist-shows banlist, fucklist, and gbanlist.",
		"bot":             "'%s%s'\n\nbot @tag, lkick, lcon, etc- adds to botlist.",
		"backup":          "'%s%s'\n\nbackup on/off- reinvites people illegally kicked or cancelled by non staff.",
		"statusall":       "'%s%s'\n\nstatusall-shows all status of bots.",
		"clearcache":      "'%s%s'\n\nclearcache-clears cache of the bots.",
		"gleave":          "'%s%s'\n\ngleave 'number'-bots leave from specific group.",
		"fuck":            "'%s%s'\n\nfuck @tag, lkick, lcon, etc- add to fucklist, similar to banlist.",
		"ban":             "'%s%s'\n\nban @tag, lkick, lcon, etc-adds to banlist.",
		"sname":           "'%s%s'\n\nsname- use without prefix and bots respond with sname.",
		"rname":           "'%s%s'\n\nrname-use without prefix and bots respond with rname.",
		"unban":           "'%s%s'\n\nunban ltag, @tag, lkick, etc- removes from banlist.",
		"mute":            "'%s%s'\n\nmute @tag, lkick, lcon-mutes person so when they spak they get kicked.",
		"unmute":          "'%s%s'\n\nunmute @tag, lkick, lcon-removes from mutelist.",
		"gowner":          "'%s%s'\n\ngowner @tag, lkick, lcon- makes group owner only.",
		"gourl":           "'%s%s'\n\ngourl “number”- Opens and gets link from group remotely.",
		"gcurl":           "'%s%s'\n\ngcurl “number”- Closes group link remotely.",
		"timenow":         "'%s%s'\n\ntimenow-shows time and date.",
		"gban":            "'%s%s'\n\ngban @tag, lkick, lcon-adds to group ban list.",
		"ungban":          "'%s%s'\n\nungban @tag, lkick, lcon-removes from group banlist.",
		"gadmin":          "'%s%s'\n\ngadmin @tag, lkick, lcon-adds to group admin.",
		"lurk mention":    "'%s%s'\n\nlurk mention-enables lurk with tag.",
		"lurk name":       "'%s%s'\n\nlurk name-enables lurk with name.",
		"lurk off":        "'%s%s'\n\nlurk off-disables lurk.",
		"lurks":           "'%s%s'\n\nlurks-shows lurkers when lurk is on.",
		"banpurge":        "'%s%s'\n\nbanpurge on/off- auto purges ban list from every group.",
		"squadcontact":    "'%s%s'\n\nsquadcontact- sends all the squads contacts.",
		"mybots":          "'%s%s'\n\nmybots- sends all the squads contacts.",
		"mybot":           "'%s%s'\n\nmybot- sends all the squads contacts.",
		"lurk hide":       "'%s%s'\n\nlurk hide- Activates lurk but does not show lurkers until lurk is stopped.",
		"@me":             "'%s%s'\n\nContact @me, .mid @me, etc- Uses yourself as a parameter for some commands.",
		"@pending":        "'%s%s'\n\nContact @pending, mid @pending, etc- Uses pending as parameter for commands.",
		"@all":            "'%s%s'\n\nall Contact @all, .mid @all, etc- Uses everyone in group as parameter for some     commands.",
		"shutdown":        "'%s%s'\n\nShutting down the bot's.",
		"perm":            "'%s%s .<grade>.<command>'\n\nAvailable grade buyer/owner/master/admin",
		"nukejoin":        "'%s%s' on/off\nkickall member's while bot has invited..",
		"lagjoin":         "'%s%s' on/off\nWhen this command is used, it enters the group and sends messages that disrupt the movement of the group",
		"kicktime":        "'%s%s' on/off\nkickall Bot's while bot has if you usd kick @..",
		"announce":        "'%s%s on/off'\n\nEnable detect announce.",
		"hostage":         "'%s%s on/off'\n\nEnable auto invite leave member.",
		"accept":          "'%s%s <number>'\n\nAccept group invitation by number.",
		"reject":          "'%s%s <number>'\n\nReject group invitation by number.",
		"welcome":         "'%s%s on/off'\n\nEnable welcome message.",
		"leave":           "'%s%s on/off'\n\nEnable leave message.",
		"setcmd":          "'%s%s <state> <command>'\n\nUsed to enabling/disabling command\nAvailable state lock/unlock/disable/enable.",
		"fixed":           "'%s%s'\n\nIf bot's error, please use this command to autofix.",
		"logmode":         "'%s%s <state>'\n\nUsed to see bot's activity.\nAvailable state on/off",
		"go":              "'%s%s <number>'\n\nSet bot to stay on group invitation.\nDefault is 2 bot.",
		"unseller":        "'%s%s <range/lcon/lkick/etc>'Used to expel seller.\nAvailable range '<', '>', '-', ',' with number.",
		"unbuyer":         "'%s%s <range/lcon/lkick/etc>'Used to expel buyer.\nAvailable range '<', '>', '-', ',' with number.",
		"unowner":         "'%s%s <range/lcon/lkick/etc>'Used to expel owner.\nAvailable range '<', '>', '-', ',' with number.",
		"unadmin":         "'%s%s <range/lcon/lkick/etc>'Used to expel admin.\nAvailable range '<', '>', '-', ',' with number.",
		"unmaster":        "'%s%s <range/lcon/lkick/etc>'Used to expel master.\nAvailable range '<', '>', '-', ',' with number.",
		"ungowner":        "'%s%s <range/lcon/lkick/etc>'Used to expel gowner.\nAvailable range '<', '>', '-', ',' with number.",
		"ungadmin":        "'%s%s <range/lcon/lkick/etc>'Used to expel gadmin.\nAvailable range '<', '>', '-', ',' with number.",
		"clearseller":     "'%s%s'\n\nClearing all sellers.",
		"clearbuyer":      "'%s%s'\n\nClearing all buyer list.",
		"clearowner":      "'%s%s'\n\nClearing all owner list.",
		"clearmaster":     "'%s%s'\n\nClearing all master list.",
		"clearadmin":      "'%s%s'\n\nClearing all admin list.",
		"cleargadmin":     "'%s%s'\n\nClearing all gadmin list.",
		"cleargowner":     "'%s%s'\n\nClearing all gowner list.",
		"clearbot":        "'%s%s'\n\nClearing all bot list.",
		"clearban":        "'%s%s'\n\nClearing all ban list.",
		"clearfuck":       "'%s%s'\n\nClearing all fuck list.",
		"clearmute":       "'%s%s'\n\nClearing all mute list.",
		"cleargban":       "'%s%s'\n\nClearing all gban list.",
		"clearchat":       "'%s%s'\n\nClearing all squad messages.",
		"upvallimage":     "'%s%s'\n\nUpdating all bot's video profile.",
		"upvimage":        "'%s%s'\n\nUpdating all bot's video profile.",
		"upallimage":      "'%s%s'\n\nUpdating all bot's picture profile.",
		"upimage":         "'%s%s'\n\nUpdating bot's profile picture.",
		"upvallcover":     "'%s%s'\n\nUpdating all bot's video cover.",
		"upvcover":        "'%s%s @tag bot'\n\nUpdating bot's video cover.",
		"upcover":         "'%s%s' @tag\n\nUpdating bot's cover picture.",
		"upallcover":      "'%s%s'\n\nUpdating all bot's cover picture.",
		"upname":          "'%s%s newname'\n\nUpdating bot's displayname.",
		"upallname":       "'%s%s newname'\n\nUpdating all bot's displayname.",
		"leaveall":        "'%s%s'\n\nleave all bot's from all group's.",
		"groups":          "'%s%s'\n\nsee bot group's.",
		"stayall":         "'%s%s'\n\naccepting all group invitation.",
		"setcom":          "'%s%s .key .value'\n\nChange command.",
		"upstatus":        "'%s%s <status message>'\n\nUpdating bot's profile bio.",
		"upallstatus":     "'%s%s <status message>'\n\nUpdating all bot's profile bio.",
		"kick":            "'%s%s @tag/lcon/lkick/etc'\n\nKick member's.",
		"invme":           "'%s%s gnumber'\n\nInvite user to the destination group.",
		"autojoin":        "'%s%s qr/invite/off'\n\nForcing bot's to joinall while invited.",
		"autoban":         "'%s%s on/off'\n\nAuto banned user.",
		"sellers":         "'%s%s'\n\nShow seller list.",
		"buyers":          "'%s%s'\n\nShow buyer list.",
		"owners":          "'%s%s'\n\nShow owner list.",
		"masters":         "'%s%s'\n\nShow master list.",
		"admins":          "'%s%s'\n\nShow admin list.",
		"gowners":         "'%s%s'\n\nShow gowner list.",
		"gadmins":         "'%s%s'\n\nShow gadmin list.",
		"botlist":         "'%s%s'\n\nShow bot list.",
		"banlist":         "'%s%s'\n\nShow ban list.",
		"fucklist":        "'%s%s'\n\nShow fuck list.",
		"mutelist":        "'%s%s'\n\nShow mutelist list.",
		"gbanlist":        "'%s%s'\n\nShow gban list.",
		"hides":           "'%s%s'\n\nShow Invisible user.",
		"hide":            "'%s%s @tag/lcon/lkick/etc'\n\nAdded user to invisible list.",
		"kickall":         "'%s%s'\n\nKick all group member's.",
		"group info":      "'%s%s'\n\nShow all group member's./pendings/access",
		"autopurge":       "'%s%s on/off'\n\nEnable autopurge.",
		"lurk":            "'%s%s on/off'\n\nEnable lurking mode.",
		"lurkmsg":         "'%s%s <message>'\n\nSet lurk message.\nUse @! for placing user tagging.",
		"antitag":         "'%s%s on/off'\n\nEnable antitag.",
		"killmode":        "'%s%s kill/purge/on/off/range'\n\nKiller mode to kick all banlist/squad.",
		"autopro":         "'%s%s on/off'\n\nAuto protect max while bot's join.",
		"setlimit":        "'%s%s number'\n\nSet max kick in killmode /bot.",
		"stay":            "'%s%s number'\n\nSet amount of bot's in group invite via link invitation.",
		"bringall":        "'%s%s'\n\nBring all bot's by invitation.",
		"bring":           "'%s%s number'\n\nSet amount of bot's in group via invitation.",
		"here":            "'%s%s'\n\nShow amount of bot's in group.",
		"friends":         "'%s%s'\n\nTag bot and use, shows all friends of tagged bot",
		"msgrespon":       "'%s%s respon'\n\nSet bot's response.",
		"setrname":        "'%s%s newrname'\n\nChange the rname prefix.",
		"setsname":        "'%s%s newsname'\n\nChange the sname prefix.",
		"invite":          "'%s%s @tag/lcon/lkick/etc'\n\nInvite target to the group's.",
		"gaccess":         "'%s%s'\n\nSee all group access list.",
		"limitout":        "'%s%s'\n\nLeave the kicbanned bot's.",
		"say":             "'%s%s word'\n\nThe bot's would said the word.",
		"sayall":          "'%s%s word'\n\nAll bot's would said the word.",
		"expel":           "'%s%s @tag/lcon/lkick/etc'\n\nUsed to expel user access.",
		"respon":          "'%s%s'\n\nBot response.",
		"ping":            "'%s%s'\n\nBot response.",
		"settings":        "'%s%s'\n\nShow the group preset status in group.",
		"set":             "'%s%s'\n\nShow the bot's set.",
		"help":            "'%s%s'\n\nShow the help command.",
		"deny":            "'%s%s invite/kick/qr/join/cancel/off/all/max'\n\nEnable the protection.",
		"allow":           "'%s%s invite/kick/qr/join/cancel/all'\n\nDisable the protection.",
		"ourl":            "'%s%s'\n\nOpen group links.",
		"curl":            "'%s%s'\n\nClose group links.",
		"count":           "'%s%s'\n\nShow bot's number.",
		"speed":           "'%s%s'\n\nShow bot response speed.",
		"unsend":          "'%s%s count'\n\nUnsend recent bot's message.\nIf count not definde, it would unsend all recent message.",
		"tagall":          "'%s%s'\n\nTagging all member's.",
		"ftagall":         "'%s%s'\n\nTagging all member's with sticker.",
		"access":          "'%s%s'\n\nShow all bot access.",
		"bans":            "'%s%s'\n\nShow the bot's status.",
		"runtime":         "'%s%s'\n\nShiw the bot's time alive.",
		"timeleft":        "'%s%s'\n\nShow the bot's timeleft.",
		"linvite":         "'%s%s'\n\nShow the last invited in group.",
		"lkick":           "'%s%s'\n\nShow the last kicked in group.",
		"lmid":            "'%s%s'\n\nShow the last mid in group.",
		"lcon":            "'%s%s'\n\nShow the last contact in group.",
		"ltag":            "'%s%s'\n\nShow the last tag in group.",
		"lban":            "'%s%s'\n\nShow the last banned in group.",
		"lcancel":         "'%s%s'\n\nShow the last cancel in group.",
		"lqr":             "'%s%s'\n\nShow the last upded qr in group.",
		"ljoin":           "'%s%s'\n\nShow the last join in group.",
		"lleave":          "'%s%s'\n\nShow the last leave in group.",
		"abort":           "'%s%s'\n\nAborting command.",
		"groupcast":       "'%s%s <your word>'\n\nBroadcasting message to all groups.",
		"contact":         "'%s%s @tag/lcon/lkick/etc'\n\nUsed to get contact's.",
		"rollcall":        "'%s%s'\n\nShow bot's name.",
		"gojoin":          "'%s%s'\n\nJoining bot's from invitation list.",
		"mid":             "'%s%s @tag/lcon/lkick/etc'\n\nGet midlist.",
		"name":            "'%s%s @tag/lcon/lkick/etc'\n\nGet namelist.",
		"purgeall":        "'%s%s'\n\nPurge all banlist in all group.",
		"squadmid":        "'%s%s'\n\nShow all bots mid.",
		"whois":           "'%s%s @tag/lcon/lkick/etc'\n\nSee user info.",
		"cancel":          "'%s%s @tag/lcon/lkick/etc'\n\nCancel group invitation.",
	}
)

func AllBanList(self *oop.Account) string {
	listadm := "✠ 𝗔𝗹𝗹 𝗯𝗮𝗻𝗹𝗶𝘀𝘁𝘀 ✠"
	if len(Banned.Banlist) != 0 {
		listadm += "\n\n ☠️ 𝗕𝗮𝗻𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Banlist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(Banned.Fucklist) != 0 {
		listadm += "\n\n ☠️ 𝗙𝘂𝗰𝗸𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Fucklist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(Banned.Mutelist) != 0 {
		listadm += "\n\n ☠️ 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Mutelist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	return listadm
}

func Checkserver(ip string) bool {
	if InArray2(ListIp, ip) {
		return true
	}
	return false
}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func MemBan(to, user string) bool {
	defer oop.PanicOnly()
	if Banned.GetBan(user) {
		return true
	} else if Banned.GetFuck(user) {
		return true
	} else {
		Room := oop.GetRoom(to)
		if InArray2(Room.Gban, user) {
			return true
		}
		return false
	}
	return false
}
func MemBan2(to, user string) bool {
	defer oop.PanicOnly()
	if Banned.GetBan(user) {
		return true
	} else if Banned.GetFuck(user) {
		return true
	} else if Banned.GetMute(user) {
		return true
	} else {
		Room := oop.GetRoom(to)
		if InArray2(Room.Gban, user) {
			return true
		}
		return false
	}
	return false
}
func Cekbanwhois(client *oop.Account, to string, targets []string) {
	room := oop.GetRoom(to)
	list := ""
	if len(targets) > 1 {
		ban := []string{}
		fuck := []string{}
		mute := []string{}
		Gban := []string{}
		for _, from := range targets {
			if Banned.GetFuck(from) {
				fuck = append(fuck, from)
			} else if Banned.GetBan(from) {
				ban = append(ban, from)
			} else if Banned.GetMute(from) {
				mute = append(mute, from)
			} else if InArray2(room.Gban, from) {
				Gban = append(Gban, from)
			}
		}
		if len(ban) != 0 {
			list += "𝗘𝘅𝗶𝘀𝘁.𝗶𝗻 𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range ban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(fuck) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗳𝘂𝗰𝗸𝗹𝗶𝘀𝘁:\n"
			for n, xx := range fuck {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Gban) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range Gban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(mute) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁:\n"
			for n, xx := range mute {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
	} else {
		for _, from := range targets {
			if Banned.GetFuck(from) {
				list += "User have access exist in fuck list."
			} else if Banned.GetBan(from) {
				list += "User have access exist in ban list."
			} else if InArray2(room.Gban, from) {
				list += "User have access exist in gban list."
			} else if Banned.GetMute(from) {
				list += "User have access exist in mute list."
			}

		}
	}
	if list != "" {
		client.SendMessage(to, list)
	}
}
func autokickban(client *oop.Account, to string, target string) {
	if AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, target)
			go client.CancelChatInvitations(aa, target)
		}
	}
}
func KIckbansPurges(client *oop.Account, group string) {
	defer oop.PanicOnly()
	gr, _ := client.GetGroupIdsJoined()
	nus := []string{}
	list := ""
	list += fmt.Sprintf("Purged %v groups: \n", len(gr))
	for num, aa := range gr {
		num++
		//list += fmt.Sprintf("%v- %v :", num, name)
		for _, v := range Banned.Banlist {
			if oop.IsMembers(client, aa, v) == true {
				if Banned.GetBan(v) {
					go func(v string) { client.DeleteOtherFromChats(aa, v) }(v)
					if oop.IsPending(client, aa, v) == true {
						client.CancelChatInvitations(group, v)
					}
					new := client.Getcontactuser(v)
					rengs := strconv.Itoa(num)
					if new != nil {
						list += "\n " + rengs + ". Closed Account"
					} else {
						x, _ := client.GetContact(v)
						nus = append(nus, v)
						list += "\n " + rengs + ". " + x.DisplayName
					}
				}
			}
		}
	}
	list += fmt.Sprintf("\n\nTotal kicks: %v.", len(nus))
	client.SendMessage(group, list)
}
func Ungban(group string, asu string) {
	room := oop.GetRoom(group)
	if InArray2(room.Gban, asu) {
		room.Gban = Remove(room.Gban, asu)
	}
}

func Addgban(asu string, group string) {
	room := oop.GetRoom(group)
	if !InArray2(room.Gban, asu) && asu != "" {
		room.Gban = append(room.Gban, asu)
	}
}

func IndexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
			break
		}
	}
	return -1
}
func Joinsave(Pelaku string, Optime int64) {
	defer oop.PanicOnly()
	ix := IndexOf(Detectjoin.User, Pelaku)
	if ix == -1 {
		Detectjoin.User = append(Detectjoin.User, Pelaku)
		Detectjoin.Time = append(Detectjoin.Time, Optime)
	} else {
		Detectjoin.Time[ix] = Optime
	}
}

func Checkmulti(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if InArray2(list2, v) {
			return true
		}
	}
	return false
}
func AddbanOp3(mid []string) {
	for _, m := range mid {
		Banned.AddBan(m)
	}
}
func SaveData() {
	defer oop.PanicOnly()
	file, _ := json.MarshalIndent(Data, "", "  ")
	_ = ioutil.WriteFile(DATABASE, file, 0644)
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Sutting down application.")
		os.Exit(0)
	}()
}
func randomToString(count int) string {
	numb := make([]rune, count)
	for i := range numb {
		numb[i] = stringToInt[rand.Intn(len(stringToInt))]
	}
	return string(numb)
}
func panicHandle(s string) {
	if r := recover(); r != nil {
		Ides := fmt.Sprintf("\nEror \nFunc: %v", s)
		println(Ides)
	}
}
func GetMentionData(data string) []string {
	var midmen []string
	var midbefore []string
	res := mentions{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArray2(midbefore, v.Mid) == false {
			midbefore = append(midbefore, v.Mid)
			midmen = append(midmen, v.Mid)
		}
	}

	return midmen
}
func MentionList(op *talkservice.Operation) []string {
	msg := op.Message
	str := fmt.Sprintf("%v", msg.ContentMetadata["MENTION"])
	taglist := GetMentionData(str)

	return taglist
}
func GetIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
func main() {
	defer ants.Release()
	defer oop.PanicOnly()
	debug.SetGCPercent(500)
	cpu = runtime.NumCPU()
	jsonFile, err := os.Open(DATABASE)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Data)
	IPServer = fmt.Sprintf("%v", GetIP())
	fmt.Println("\n_Started Login_:")
	go gracefulShutdown()
	for no, tok := range Data.Authoken {
		time.Sleep(250 * time.Millisecond)
		sort := rand.Intn(9999-1000) + 1000
		app := fmt.Sprintf("ANDROID\t11.6.1\tAndroid OS\t7.1.%v", sort)
		mids := strings.Split(tok, ":")
		mid := mids[0]
		var ua = "Line/11.6.1"
		cl, err := oop.CreateNewLogin(tok, no, mid, app, ua, HostName[0])
		if err == nil {
			fmt.Println("\n\n  ↳ DisplayName : " + cl.Namebot + "\n  ↳ Mid : " + cl.MID + "\n  ↳ AppName : " + cl.AppName + "\n  ↳ UserAgent : " + cl.UserAgent + "\n  ↳ Bots No: " + fmt.Sprintf("%v", no+1))
			ClientBot = append(ClientBot, cl)
			ClientMid[cl.MID] = cl
			Squadlist = append(Squadlist, cl.MID)
			cl.RemoveLeterSelling()
		} else {
			rs := err.Error()
			if strings.Contains(rs, "INTERNAL_ERROR") || strings.Contains(rs, "AUTHENTICATION_FAILED") {
				fmt.Println("\n  ↳ Status : Freez" + "\n  ↳ Mid : " + mid + "\n  ↳ Bots No: " + fmt.Sprintf("%v", no+1))
				cl.MID = mid
				cl.Limited = true
			} else {
				logs := fmt.Sprintf("\n\n▪︎ No: %v ERROR: %s", no+1, err)
				fmt.Println(logs)
			}
		}
	}
	ab := `			  ________
	         _._      | D7OM |
	       .'   '.    | BOTS |
	      / //\\\ \   |______|
	     ( ( -\- ) )     |
	      '-\_=_/-'      /)
	     .-'\   /'-.    (|/
	    /    '-'    \  / /
	    | \__   __/_/\/ /|
	    | |\     / \   /
	    \  \     \  '-'
	     "\/\     ;
	      |/|\    |
	      |   {}  |
	      |       |
	      |       |
	      |_______|
	       |  |  |
	        \ | /
	        /=|=\
	       (_/T\_)
 `
	abc := "\n\n** GO D7OM Bots  **"
	fmt.Println(string(ColorCyan), ab, string(ColorReset))
	fmt.Println(abc)
	for m, _ := range oop.HashToMap(oop.GetBlock) {
		if !InArray2(Squadlist, m) {
			oop.GetBlock.Del(m)
		}
	}
	ch := make(chan int, len(ClientBot))
	if len(ClientBot) != 0 {
		acl := len(ClientBot)
		for x := 0; x < acl; x++ {
			cc := x
			cla := ClientBot[cc]
			runtime.Gosched()
			go RunBot(cla, ch)
		}
		list := append([]*oop.Account{}, ClientBot...)
		sort.Slice(list, func(i, j int) bool {
			return list[i].KickCount < list[j].KickCount
		})
		for i, cl := range list {
			kk := i * 30
			cl.KickPoint = kk
			ko := i * 10
			cl.CustomPoint = ko
		}
		if oop.IsFriends(ClientBot[0], MAKERS[0]) == false {
			ClientBot[0].FindAndAddContactsByMid(MAKERS[0])
		}
		ClientBot[0].SendMessage(MAKERS[0], "Im fetcher.")
		Resprem()
		for i := range ClientBot {
			for _, x := range Squadlist {
				if !InArray2(ClientBot[i].Squads, x) && x != ClientBot[i].MID {
					ClientBot[i].Squads = append(ClientBot[i].Squads, x)
				}
			}
		}
		go func() {
			for {
				autoset()
				time.Sleep(3 * time.Second)
			}
		}()
		for v := range ch {
			if v == 69 {
				break
			}
		}
		fmt.Println("__GOOD_LUCK__")
	}
}
func BackSeave() {
	fmt.Println("start Backup Data *__*")
	TimeBackup = time.Time{}
	MsSname = Data.SnameBack
	MsRname = Data.RnameBack
	MsgRespon = Data.ResponBack
	if len(Data.TimeBanBack) != 0 {
		now := time.Now()
		for a := range Data.TimeBanBack {
			if InArray2(Squadlist, a) {
				tims := Data.TimeBanBack[a]
				if now.Sub(tims) < 24*time.Hour {
					self := GetKorban(a)
					if !oop.InArrayCl(oop.KickBans, self) {
						oop.KickBans = append(oop.KickBans, self)
						self.TimeBan = tims
					}
					self.Limited = true
					if _, ok := oop.GetBlock.Get(self.MID); !ok {
						oop.GetBlock.Set(self.MID, tims)
					}
				}

			}
		}
	}
	if len(Data.CreatorBack) != 0 {
		for _, i := range Data.CreatorBack {
			UserBot.AddCreator(i)
		}
	}
	if len(Data.SellerBack) != 0 {
		for _, i := range Data.SellerBack {
			UserBot.AddSeller(i)
		}
	}
	if len(Data.BuyerBack) != 0 {
		for _, i := range Data.BuyerBack {
			UserBot.AddBuyer(i)
		}
	}
	if len(Data.OwnerBack) != 0 {
		for _, i := range Data.OwnerBack {
			UserBot.AddOwner(i)
		}
	}
	if len(Data.MasterBack) != 0 {
		for _, i := range Data.MasterBack {
			UserBot.AddMaster(i)
		}
	}
	if len(Data.AdminBack) != 0 {
		for _, i := range Data.AdminBack {
			UserBot.AddAdmin(i)
		}
	}
	if len(Data.BotBack) != 0 {
		for _, i := range Data.BotBack {
			UserBot.AddBot(i)
		}
	}
	if len(Data.ProkickBack) != 0 {
		for _, to := range Data.ProkickBack {
			Room := oop.GetRoom(to)
			Room.ProKick = true
		}
	}
	if len(Data.ProCancelBack) != 0 {
		for _, to := range Data.ProCancelBack {
			Room := oop.GetRoom(to)
			Room.ProCancel = true
		}
	}
	if len(Data.ProInviteBack) != 0 {
		for _, to := range Data.ProInviteBack {
			Room := oop.GetRoom(to)
			Room.ProInvite = true
		}
	}
	if len(Data.ProQrBack) != 0 {
		for _, to := range Data.ProQrBack {
			Room := oop.GetRoom(to)
			Room.ProQr = true
		}
	}
	if len(Data.ProjoinBack) != 0 {
		for _, to := range Data.ProjoinBack {
			Room := oop.GetRoom(to)
			Room.ProJoin = true
		}
	}
	if len(Data.AnnunceBack) != 0 {
		for _, to := range Data.AnnunceBack {
			Room := oop.GetRoom(to)
			Room.Announce = true
		}
	}
	if len(Data.GadminBack) != 0 {
		for to := range Data.GadminBack {
			Room := oop.GetRoom(to)
			if len(Data.GadminBack[to]) != 0 {
				for _, user := range Data.GadminBack[to] {
					if !InArray2(Room.Gadmin, user) {
						Room.Gadmin = append(Room.Gadmin, user)
					}
				}
			}
		}
	}
	if len(Data.GownerBack) != 0 {
		for to := range Data.GownerBack {
			Room := oop.GetRoom(to)
			if len(Data.GownerBack[to]) != 0 {
				for _, user := range Data.GownerBack[to] {
					if !InArray2(Room.Gowner, user) {
						Room.Gowner = append(Room.Gowner, user)
					}
				}
			}
		}
	}
	if len(Data.GbanBack) != 0 {
		for to := range Data.GbanBack {
			Room := oop.GetRoom(to)
			if len(Data.GbanBack[to]) != 0 {
				for _, user := range Data.GbanBack[to] {
					if MemUser(to, user) {
						if !InArray2(Room.Gban, user) {
							Room.Gban = append(Room.Gban, user)
						}
					}
				}
			}
		}
	}
	if len(Data.BanBack) != 0 {
		for _, user := range Data.BanBack {
			Banned.AddBan(user)
		}
	}
	if len(Data.FuckBack) != 0 {
		for _, user := range Data.FuckBack {
			Banned.AddFuck(user)
		}
	}
	if len(Data.MuteBack) != 0 {
		for _, user := range Data.MuteBack {
			Banned.AddMute(user)
		}
	}
	TimeSave = time.Now()
	fmt.Println("done Backup Data *__*")
}

func SaveBackup() {
	fmt.Println("start Save Data *__*")
	Data.GbanBack = map[string][]string{}
	Data.GownerBack = map[string][]string{}
	Data.GadminBack = map[string][]string{}
	Data.BanBack = []string{}
	Data.SnameBack = MsSname
	Data.RnameBack = MsRname
	Data.ResponBack = MsgRespon
	Data.FuckBack = []string{}
	Data.MuteBack = []string{}
	Data.AnnunceBack = []string{}
	Data.ProQrBack = []string{}
	Data.ProjoinBack = []string{}
	Data.ProInviteBack = []string{}
	Data.ProCancelBack = []string{}
	Data.ProkickBack = []string{}
	Data.CreatorBack = []string{}
	Data.SellerBack = []string{}
	Data.BuyerBack = []string{}
	Data.OwnerBack = []string{}
	Data.MasterBack = []string{}
	Data.AdminBack = []string{}
	Data.BotBack = []string{}
	Data.TimeBanBack = map[string]time.Time{}
	if len(oop.KickBans) != 0 {
		for _, cl := range oop.KickBans {
			if _, ok := oop.GetBlock.Get(cl.MID); ok {
				Data.TimeBanBack[cl.MID] = cl.TimeBan
			}
		}
	}
	for _, room := range oop.SquadRoom {
		Data.GbanBack[room.Id] = []string{}
		Data.GownerBack[room.Id] = []string{}
		Data.GadminBack[room.Id] = []string{}
		if room.ProKick {
			Data.ProkickBack = append(Data.ProkickBack, room.Id)
		}
		if room.ProCancel {
			Data.ProCancelBack = append(Data.ProCancelBack, room.Id)
		}
		if room.ProInvite {
			Data.ProInviteBack = append(Data.ProInviteBack, room.Id)
		}
		if room.ProQr {
			Data.ProQrBack = append(Data.ProQrBack, room.Id)
		}
		if room.ProJoin {
			Data.ProjoinBack = append(Data.ProjoinBack, room.Id)
		}
		if room.Announce {
			Data.AnnunceBack = append(Data.AnnunceBack, room.Id)
		}
	}
	if len(UserBot.Creator) != 0 {
		for _, i := range UserBot.Creator {
			if !InArray2(Data.CreatorBack, i) {
				Data.CreatorBack = append(Data.CreatorBack, i)
			}
		}
	}
	if len(UserBot.Seller) != 0 {
		for _, i := range UserBot.Seller {
			if !InArray2(Data.SellerBack, i) {
				Data.SellerBack = append(Data.SellerBack, i)
			}
		}
	}
	if len(UserBot.Buyer) != 0 {
		for _, i := range UserBot.Buyer {
			if !InArray2(Data.BuyerBack, i) {
				Data.BuyerBack = append(Data.BuyerBack, i)
			}
		}
	}
	if len(UserBot.Owner) != 0 {
		for _, i := range UserBot.Owner {
			if !InArray2(Data.OwnerBack, i) {
				Data.OwnerBack = append(Data.OwnerBack, i)
			}
		}
	}
	if len(UserBot.Master) != 0 {
		for _, i := range UserBot.Master {
			if !InArray2(Data.MasterBack, i) {
				Data.MasterBack = append(Data.MasterBack, i)
			}
		}
	}
	if len(UserBot.Admin) != 0 {
		for _, i := range UserBot.Admin {
			if !InArray2(Data.AdminBack, i) {
				Data.AdminBack = append(Data.AdminBack, i)
			}
		}
	}
	if len(UserBot.Bot) != 0 {
		for _, i := range UserBot.Bot {
			if !InArray2(Data.BotBack, i) {
				Data.BotBack = append(Data.BotBack, i)
			}
		}
	}
	if len(Data.GbanBack) != 0 {
		for to := range Data.GbanBack {
			Room := oop.GetRoom(to)
			if len(Room.Gban) != 0 {
				for _, i := range Room.Gban {
					if MemUser(to, i) {
						if !InArray2(Data.GbanBack[to], i) {
							Data.GbanBack[to] = append(Data.GbanBack[to], i)
						}
					}
				}
			}
		}
	}
	if len(Data.GownerBack) != 0 {
		for to := range Data.GownerBack {
			Room := oop.GetRoom(to)
			if len(Room.Gowner) != 0 {
				for _, i := range Room.Gowner {
					if !InArray2(Data.GownerBack[to], i) {
						Data.GownerBack[to] = append(Data.GownerBack[to], i)
					}
				}
			}
		}
	}
	if len(Data.GadminBack) != 0 {
		for to := range Data.GadminBack {
			Room := oop.GetRoom(to)
			if len(Room.Gadmin) != 0 {
				for _, i := range Room.Gadmin {
					if !InArray2(Data.GadminBack[to], i) {
						Data.GadminBack[to] = append(Data.GadminBack[to], i)
					}
				}
			}
		}
	}
	if len(Banned.Banlist) != 0 {
		for _, i := range Banned.Banlist {
			if MemAccsess(i) {
				if !InArray2(Data.BanBack, i) {
					Data.BanBack = append(Data.BanBack, i)
				}
			}
		}
	}
	if len(Banned.Fucklist) != 0 {
		for _, i := range Banned.Fucklist {
			if MemAccsess(i) {
				if !InArray2(Data.FuckBack, i) {
					Data.FuckBack = append(Data.FuckBack, i)
				}
			}
		}
	}
	if len(Banned.Mutelist) != 0 {
		for _, i := range Banned.Mutelist {
			if MemAccsess(i) {
				if !InArray2(Data.MuteBack, i) {
					Data.MuteBack = append(Data.MuteBack, i)
				}
			}
		}
	}
	fmt.Println("done save Data *__*")
	SaveData()
}
func Checkkickuser(group string, user string, invited string) bool {
	Room := oop.GetRoom(group)
	if InArray2(MAKERS, invited) {
		if !InArray2(MAKERS, user) {
			return true
		}
	} else if UserBot.GetCreator(invited) {
		if !SendMycreator(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetSeller(invited) {
		if !SendMyseller(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetBuyer(invited) {
		if !SendMybuyer(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetOwner(invited) {
		if !SendMyowner(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetMaster(invited) {
		if !SendMymaster(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetAdmin(invited) {
		if !SendMyadmin(user) && !Allbotlist(user) {
			return true
		}
	} else if InArray2(Room.Gowner, invited) {
		if !SendMygowner(group, user) && !Allbotlist(user) {
			return true
		}
	} else if InArray2(Room.Gadmin, invited) {
		if MemUser(group, user) {
			return true
		}
	} else if UserBot.GetBot(invited) {
		if MemUser(group, user) {
			return true
		}
	}
	return false
}
func MemUser(group string, from string) bool {
	Room := oop.GetRoom(group)
	if InArray2(Squadlist, from) {
		return false
	} else if UserBot.GetBot(from) {
		return false
	} else if InArray2(MAKERS, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	} else if InArray2(Room.Gowner, from) {
		return false
	} else if InArray2(Room.Gadmin, from) {
		return false
	}
	return true
}
func MemAccsess(from string) bool {
	if InArray2(Squadlist, from) {
		return false
	} else if UserBot.GetBot(from) {
		return false
	} else if InArray2(MAKERS, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	}
	return true
}

func MemUserN(group string, from string) bool {
	Room := oop.GetRoom(group)
	if UserBot.GetBot(from) {
		return false
	} else if InArray2(MAKERS, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	} else if InArray2(Room.Gowner, from) {
		return false
	} else if InArray2(Room.Gadmin, from) {
		return false
	}
	return true
}
func Allbotlist(user string) bool {
	if InArray2(Squadlist, user) {
		return true
	} else if UserBot.GetBot(user) {
		return true
	}
	return false
}

func SendMycreator(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	}
	return false
}
func SendMyseller(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	}
	return false
}
func SendMybuyer(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	}
	return false
}
func SendMyowner(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	}
	return false
}
func SendMymaster(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	}
	return false
}
func SendMyadmin(from string) bool {
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	}
	return false
}

func SendMygowner(group string, from string) bool {
	Room := oop.GetRoom(group)
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	} else if InArray2(Room.Gowner, from) == true {
		return true
	}
	return false
}
func SendMygadmin(group string, from string) bool {
	Room := oop.GetRoom(group)
	if InArray2(MAKERS, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	} else if InArray2(Room.Gowner, from) {
		return true
	} else if InArray2(Room.Gadmin, from) {
		return true
	}
	return false
}
func InArrayInt64(arr []int64, str int64) bool {
	for _, tar := range arr {
		if tar == str {
			return true
		}
	}
	return false
}
func GetCodeprem(num int, sinder string, group string) bool {
	if num == 0 {
		if InArray2(MAKERS, sinder) {
			return true
		}
	} else if num == 1 {
		if SendMycreator(sinder) {
			return true
		}
	} else if num == 2 {
		if SendMyseller(sinder) {
			return true
		}
	} else if num == 3 {
		if SendMybuyer(sinder) {
			return true
		}
	} else if num == 4 {
		if SendMyowner(sinder) {
			return true
		}
	} else if num == 5 {
		if SendMymaster(sinder) {
			return true
		}
		return false
	} else if num == 6 {
		if SendMyadmin(sinder) {
			return true
		}
	} else if num == 7 {
		if SendMygowner(group, sinder) {
			return true
		}
	} else if num == 8 {
		if SendMygadmin(group, sinder) {
			return true
		}
	}
	return false
}
func PerCheckList() string {
	list := ""
	var test1 string
	if SetHelper.Rngcmd != nil {
		list += "✠ 𝗟𝗶𝘀𝘁 𝗽𝗲𝗿𝗺 :\n\n"
		for i := range SetHelper.Rngcmd {
			if SetHelper.Rngcmd[i] == 0 {
				test1 = "Maker"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 1 {
				test1 = "Creator"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 2 {
				test1 = "Seller"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 3 {
				test1 = "Buyer"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 4 {
				test1 = "Owner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 5 {
				test1 = "Master"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 6 {
				test1 = "Admin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 7 {
				test1 = "Gowner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 8 {
				test1 = "Gadmin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			}
		}
	}
	return list
}
func Addpermcmd(client *oop.Account, to string, test1 string, test2 string) {
	x := 0
	numr := 0
	list := ""
	if test1 == "maker" {
		x = 0
	} else if test1 == "creator" {
		x = 1
	} else if test1 == "seller" {
		x = 2
	} else if test1 == "buyer" {
		x = 3
	} else if test1 == "owner" {
		x = 4
	} else if test1 == "master" {
		x = 5
	} else if test1 == "admin" {
		x = 6
	} else if test1 == "gowner" {
		x = 7
	} else if test1 == "gadmin" {
		x = 8
	} else {
		list += "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin."
	}
	if list != "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin." {
		cmd2 := test2
		_, value := SetHelper.Rngcmd[cmd2]
		if value == true {
			if SetHelper.Rngcmd[cmd2] != x {
				SetHelper.Rngcmd[cmd2] = x
				numr = 5
			}
		} else {
			list += "Command not found.\nUse ths Command First."
		}
	}
	if list != "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin." {
		if list != "Command not found.\nUse ths Command First." {
			if numr != 5 {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("%v is already a %v command.\n", cmd2, cmd1)
			} else {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("Changed permission to %v for: %v \n", cmd1, cmd2)
			}
			client.SendMessage(to, list)
		} else {
			client.SendMessage(to, list+"\n")
		}
	} else {
		client.SendMessage(to, list+"\n")
	}

}
func GetComs(gr int, data string) int {
	defer oop.PanicOnly()
	_, value := SetHelper.Rngcmd[data]
	if value == false {
		SetHelper.Rngcmd[data] = gr
	}
	xx := SetHelper.Rngcmd[data]
	return xx
}
func CheckAccount(user string) *oop.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}
func Checkuser(client *oop.Account, group string) ([]*oop.Account, []string) {
	list := []string{}
	err, _, memlist := client.GetGroupMembers(group)
	if err != nil {
		return nil, list
	}
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := CheckAccount(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemUser(group, mid) {
			list = append(list, mid)
		}
	}
	if len(exe) != 0 && len(list) != 0 {
		return exe, list
	}
	return nil, list
}
func CheckBan(client *oop.Account, group string) []string {
	list := []string{}
	err, _, memlist := client.GetGroupMembers(group)
	if err != nil {
		return list
	}
	for mid, _ := range memlist {
		if MemUser(group, mid) {
			if MemBan(group, mid) {
				list = append(list, mid)
			}
		}
	}
	return list
}
func GETgrade(num int) string {
	if num == 0 {
		return "Maker"
	} else if num == 1 {
		return "Creator"
	} else if num == 2 {
		return "seller"
	} else if num == 3 {
		return "Buyer"
	} else if num == 4 {
		return "Owner"
	} else if num == 5 {
		return "Master"
	} else if num == 6 {
		return "Admin"
	} else if num == 7 {
		return "Gowner"
	} else if num == 8 {
		return "Gadmin"
	}
	return "None"
}
func Canceljoin(Client *oop.Account, Group string) {
	defer panicHandle("canceljoin")
	_, _, pind := Client.GetChatList(Group)
	for _, i := range pind {
		if oop.IsPending(Client, Group, i) == true {
			Client.CancelChatInvitations(Group, i)
		}
	}
}
func Lagejoin(Client *oop.Account, Group string) {
	defer panicHandle("Lagejoin")
	file1, _ := os.Open("ios.txt")
	file2, _ := os.Open("andr.txt")
	b1, _ := ioutil.ReadAll(file1)
	b2, _ := ioutil.ReadAll(file2)
	Client.SendMessage(Group, string(b1))
	Client.SendMessage(Group, string(b2))
	Client.SendMessage(Group, "The group is deleted")
}
func Nukjoin(Client *oop.Account, Optime int64, Group string) {
	defer panicHandle("Nukejoin")
	_, ok := Ceknuke.Get(Optime)
	if !ok {
		Ceknuke.Set(Optime, 1)
	} else {
		return
	}
	exe, list := Checkuser(Client, Group)
	if exe != nil {
		no := 0
		i := 0
		lm := len(list)
		acts := []*oop.Account{}
		var cl *oop.Account
		for ; i < lm; i++ {
			if no >= len(exe) {
				no = 0
			}
			acts = append(acts, exe[no])
			no += 1
		}
		for n, target := range list {
			go func(n int, target string) {
				cl = acts[n]
				cl.DeleteOtherFromChats(Group, target)
			}(n, target)
		}
	}
}
func RemoveSticker(items []*Stickers, item *Stickers) []*Stickers {
	defer oop.PanicOnly()
	newitems := []*Stickers{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
func AutopurgeEnd(client *oop.Account, Group string, mem []string) {
	defer panicHandle("AutopurgeEnd")
	for _, target := range mem {
		client.DeleteOtherFromChats(Group, target)
	}
}

func AppendLastSticker(s []*Stickers, e *Stickers) []*Stickers {
	defer oop.PanicOnly()
	s = RemoveSticker(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}
func CheckMessage(waktu int64, typ int8) bool {
	if typ == 1 {
		for _, wkt := range timeSend {
			if wkt == waktu {
				return false
				break
			}
		}
		timeSend = append(timeSend, waktu)
		return true
	}
	return false
}
func Remove(s []string, r string) []string {
	new := make([]string, len(s))
	copy(new, s)
	for i, v := range new {
		if v == r {
			return append(new[:i], new[i+1:]...)
		}
	}
	return s
}
func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	x := d
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	if x < 60*time.Second {
		return fmt.Sprintf("%v", x)
	} else if x < 3600*time.Second {
		return fmt.Sprintf("%02dMinutes's", m)
	} else if x < 86400*time.Second {
		return fmt.Sprintf("%02dHour's %02dMinute's", h%24, m)
	} else {
		return fmt.Sprintf("%02dDay's %02dHour's %02dMinute's", h/24, h%24, m)
	}
}

func CancelEnd(client *oop.Account, Group string, mem []string) {
	defer panicHandle("CancelEnd")
	for _, target := range mem {
		client.CancelChatInvitations(Group, target)
	}
}

func Setpurgealln(client *oop.Account, to string, invits []string) {
	for _, cc := range invits {
		if oop.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, cc)
		} else if oop.IsPending(client, to, cc) == true {
			client.CancelChatInvitations(to, cc)
		}
	}

}
func SelectBot(client *oop.Account, to string) (*oop.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], err
	}
	return nil, err
}
func CheckBot(client *oop.Account, to string) (*oop.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], err
	}
	return nil, err
}
func GetKorban(user string) *oop.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}
func squadMention(mlist []string) (m *oop.Account, b bool) {
	for _, l := range mlist {
		if InArray2(Squadlist, l) {
			cl := GetKorban(l)
			return cl, true
		}
	}
	return nil, false
}
func StripOut(kata string) string {
	kata = strings.TrimSpace(kata)
	return kata
}

func Setinviteto(client *oop.Account, to string, invits []string) {
	news := []string{}
	for _, cc := range invits {
		if oop.IsMembers(client, to, cc) == false && oop.IsPending(client, to, cc) == false {
			news = append(news, cc)
		}
	}
	if len(news) != 0 {
		client.InviteIntoChatPollVer(to, news)
	}
}
func Setinvitetomsg(client *oop.Account, to string, invits []string) []string {
	defer panicHandle("Setinvitetomsg")
	bans := []string{}
	news := []string{}
	room := oop.GetRoom(to)
	exe := room.HaveClient
	for _, cc := range invits {
		if oop.IsMembers(client, to, cc) == false && oop.IsPending(client, to, cc) == false {
			if !MemBan(to, cc) {
				if oop.IsFriends(client, cc) == false {
					client.FindAndAddContactsByMidnew(cc)
					time.Sleep(250 * time.Millisecond)
				}
				news = append(news, cc)
			} else {
				bans = append(bans, cc)

			}
		}
	}
	if len(news) != 0 && len(exe) != 0 {
		celek := len(news)
		no := 0
		bat := 5
		ClAct := len(exe)
		if ClAct != 0 {
			if celek < bat {
				for _, cl := range exe {
					cl.GetRecommendationIds()
					for _, mid := range news {
						oop.AddContact3(cl, mid)
					}
					fl, _ := cl.GetAllContactIds()
					bb := []string{}
					for _, mid := range news {
						if InArray2(fl, mid) {
							bb = append(bb, mid)
							news = Remove(news, mid)
						}
					}
					if len(bb) != 0 {
						cl.InviteIntoGroupNormal(to, bb)
					}
					if len(news) == 0 {
						break
					}
				}
			} else {
				hajar := []string{}
				z := celek / bat
				y := z + 1
				for i := 0; i < y; i++ {
					if no >= ClAct {
						no = 0
					}
					client := exe[no]
					if i == z {
						hajar = news[i*bat:]
					} else {
						hajar = news[i*bat : (i+1)*bat]
					}
					if len(hajar) != 0 {
						client.GetRecommendationIds()
						for _, mid := range hajar {
							oop.AddContact3(client, mid)
						}
						fl, _ := client.GetAllContactIds()
						bb := []string{}
						for _, mid := range hajar {
							if InArray2(fl, mid) {
								bb = append(bb, mid)
							}
						}
						if len(bb) != 0 {
							client.InviteIntoGroupNormal(to, bb)
						}
					}
					no += 1
				}
			}
		}
	}
	return bans
}
func gettxt(from string, client *oop.Account, pesan string, rname string, sname string, Mid string, MentionMsg []string, group string) string {
	var txt string
	ca, ok := squadMention(MentionMsg)
	if ok {
		pr, _ := ca.GetContact(ca.MID)
		name := pr.DisplayName
		Vs := fmt.Sprintf("@%v", name)
		Vs = strings.ToLower(Vs)
		Vs = strings.TrimSuffix(Vs, " ")
		txt = strings.Replace(pesan, Vs, "", 1)
		txt = strings.TrimPrefix(txt, " ")
		for _, men := range MentionMsg {
			prs, _ := ca.GetContact(men)
			names := prs.DisplayName
			jj := fmt.Sprintf("@%v", names)
			jj = strings.ToLower(jj)
			jj = strings.TrimSuffix(jj, " ")
			txt = strings.Replace(txt, jj, "", 1)
			txt = StripOut(txt)
		}
		used = rname
	}
	if strings.HasPrefix(pesan, rname) {
		txt = strings.Replace(pesan, rname, "", 1)
		used = rname
	} else if strings.HasPrefix(pesan, sname) {
		txt = strings.Replace(pesan, sname, "", 1)
		used = sname
	}
	txt = StripOut(txt)
	return txt
}
func LogLast(op *talkservice.Operation, midds string) {
	defer oop.PanicOnly()
	if op.Type == 26 {
		if op.Message.ContentType == 18 {
			return
		}
	}
	LastActive.Set(midds, op)
}
func LogOp(op *talkservice.Operation, client *oop.Account) {
	defer oop.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	if tipe == 124 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 133 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 130 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 122 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 55 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 128 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 26 {
		msg := op.Message
		if InArray2(Squadlist, msg.From_) {
			return
		}
		LogLast(op, msg.From_)
	}
}
func Setkickto(client *oop.Account, to string, invits []string) {
	defer panicHandle("Setkickto")
	for _, cc := range invits {
		if oop.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, cc)
		}
	}

}
func addCon(cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !InArray2(fl, con) && con != cl.MID && !cl.Limitadd {
				cl.FindAndAddContactsByMidnew2(con)
				time.Sleep(3 * time.Second)
			}
		}
		n += 1
	}
}
func ReloginProgram() error {
	file, err := osext.Executable()
	if err != nil {
		return err
	}
	err = syscall.Exec(file, os.Args, os.Environ())
	if err != nil {
		return err
	}
	return nil
}
func GenerateTimeLog(client *oop.Account, to string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a := time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0" + strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0" + strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0" + strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := "↳Date : " + dd + "-" + MM + "-" + yyyy + "\n↳Time : " + hhconv + ":" + mmconv + ":" + ssconv
	client.SendMessage(to, times)
}
func SelectallBot(client *oop.Account, to string) ([]*oop.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			exe = append(exe, cl)
		}
	}
	if len(exe) != 0 {
		return exe, err
	}
	return nil, err
}

func Setcancelto(client *oop.Account, to string, invits []string) {
	defer panicHandle("Setcancelto")
	for _, x := range invits {
		if oop.IsPending(client, to, x) == true {
			client.CancelChatInvitations(to, x)
		}
	}
}

func cekOp(optime int64) bool {
	for _, tar := range oplist {
		if tar == optime {
			return false
		}
	}
	oplist = append(oplist, optime)
	return true
}

func cekOpinvite(optime int64) bool {
	for _, tar := range oplistinvite {
		if tar == optime {
			return false
		}
	}
	oplistinvite = append(oplistinvite, optime)
	return true
}

func AppendLastD(s [][]string, e []string) [][]string {
	defer oop.PanicOnly()
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}
func AppendLast(s []string, e string) []string {
	defer oop.PanicOnly()
	s = Remove(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}
func LlistCheck(client *oop.Account, to string, typec string, nCount int, sender string, rplay string, mentionlist []string) (ss []string) {
	saodd := []string{}
	if len(mentionlist) != 0 {
		for a := range mentionlist {
			if !InArray2(saodd, mentionlist[a]) && !InArray2(Squadlist, mentionlist[a]) {
				saodd = append(saodd, mentionlist[a])
			}

		}
		return saodd
	} else if rplay != "" {
		if !InArray2(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "lmid" {
		g, ok := Lastmid.Get(to)
		if !ok {
			g = [][]string{}
			Lastmid.Set(to, g)
		} else {
			num := nCount
			c := g.([][]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i]...)
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "linvite" {
		g, ok := Lastinvite.Get(to)
		if !ok {
			g = []string{}
			Lastinvite.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lkick" {
		g, ok := Lastkick.Get(to)
		if !ok {
			g = []string{}
			Lastkick.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcancel" {
		g, ok := Lastcancel.Get(to)
		if !ok {
			g = []string{}
			Lastcancel.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lqr" {
		g, ok := Lastupdate.Get(to)
		if !ok {
			g = []string{}
			Lastupdate.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ljoin" {
		g, ok := Lastjoin.Get(to)
		if !ok {
			g = []string{}
			Lastjoin.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ltag" {
		g, ok := Lasttag.Get(to)
		if !ok {
			g = []string{}
			Lasttag.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcon" {
		g, ok := Lastcon.Get(to)
		if !ok {
			g = []string{}
			Lastcon.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lleave" {
		g, ok := Lastleave.Get(to)
		if !ok {
			g = []string{}
			Lastleave.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "@me" {
		if !InArray2(saodd, sender) {
			saodd = append(saodd, sender)
		}
		return saodd
	} else if typec == "@all" {
		Member := client.GetChatListMem(to)
		for i := 0; i < len(Member); i++ {
			if !InArray2(saodd, Member[i]) {
				saodd = append(saodd, Member[i])
			}
		}
		return saodd
	} else if typec == "@pending" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !InArray2(saodd, i) {
				saodd = append(saodd, i)
			}
		}
		return saodd
	}
	return saodd
}
func savejoin(Pelaku string, Optime int64) {
	defer oop.PanicOnly()
	ix := IndexOf(Detectjoin.User, Pelaku)
	if ix == -1 {
		Detectjoin.User = append(Detectjoin.User, Pelaku)
		Detectjoin.Time = append(Detectjoin.Time, Optime)
	} else {
		Detectjoin.Time[ix] = Optime
	}
}
func LogGet(op *talkservice.Operation) {
	defer oop.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	korban := op.Param3
	if tipe == 124 || tipe == 123 {
		var invites []string
		if tipe == 124 {
			invites = strings.Split(korban, "\x1e")
		} else {
			invites = strings.Split(pelaku, "\x1e")
		}
		ll := len(invites)
		if ll != 0 {
			g, ok := Lastinvite.Get(op.Param1)
			if !ok {
				Lastinvite.Set(op.Param1, invites)
			} else {
				c := g.([]string)
				for _, can := range invites {
					c = AppendLast(c, can)
				}
				Lastinvite.Set(op.Param1, c)
			}
		}

	} else if tipe == 133 {
		g, ok := Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 132 {
		g, ok := Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 130 {
		g, ok := Lastjoin.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastjoin.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastjoin.Set(op.Param1, c)
		}
	} else if tipe == 125 {
		g, ok := Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 126 {
		g, ok := Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 122 {
		g, ok := Lastupdate.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastupdate.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastupdate.Set(op.Param1, c)
		}

	} else if tipe == 128 {
		g, ok := Lastleave.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastleave.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastleave.Set(op.Param1, c)
		}

	} else if tipe == 26 {
		var MentionMsg = MentionList(op)
		msg := op.Message
		if InArray2(Squadlist, msg.From_) {
			return
		}
		if len(MentionMsg) != 0 {
			g, ok := Lasttag.Get(msg.To)
			if !ok {
				g = MentionMsg
				Lasttag.Set(msg.To, g)
			} else {
				c := g.([]string)
				for _, men := range MentionMsg {
					c = AppendLast(c, men)
				}
				Lasttag.Set(msg.To, c)
			}
		} else if msg.ContentType == 13 {
			mids := msg.ContentMetadata["mid"]
			g, ok := Lastcon.Get(msg.To)
			if !ok {
				g = []string{mids}
				Lastcon.Set(msg.To, g)
			} else {
				c := g.([]string)
				c = AppendLast(c, mids)
				Lastcon.Set(msg.To, c)
			}

		} else if msg.ContentType == 7 {
			var ids []string
			var pids []string
			zx := msg.ContentMetadata
			vok, cook := zx["REPLACE"]
			if cook {
				ress := gjson.Get(vok, "sticon")
				mp := ress.Map()
				yo := mp["resources"]
				vls := yo.Array()
				for _, vl := range vls {
					mm := vl.Map()
					pids = append(pids, mm["productId"].String())
					ids = append(ids, mm["sticonId"].String())
				}
			} else {
				ids = []string{zx["STKID"]}
				pids = []string{zx["STKPKGID"]}
			}

			g, ok := Laststicker.Get(msg.To)
			if !ok {
				g = []*Stickers{&Stickers{Id: ids[0], Pid: pids[0]}}
				Laststicker.Set(msg.To, g)
			} else {
				c := g.([]*Stickers)
				c = AppendLastSticker(c, &Stickers{Id: ids[0], Pid: pids[0]})
				Laststicker.Set(msg.To, c)
			}

		} else if msg.ContentType == 0 {
			if strings.Contains(msg.Text, "u") {
				regex, _ := regexp.Compile(`u\w{32}`)
				links := regex.FindAllString(msg.Text, -1)
				mmd := []string{}
				for _, a := range links {
					if len(a) == 33 {
						mmd = append(mmd, a)
					}
				}
				if len(mmd) != 0 {
					g, ok := Lastmid.Get(msg.To)
					if !ok {
						g = [][]string{mmd}
						Lastmid.Set(msg.To, g)
					} else {
						c := g.([][]string)
						c = AppendLastD(c, mmd)
						Lastmid.Set(msg.To, c)
					}
				}
			}
		}
	}
}
func BanAll(memlist []string) {
	ilen := len(memlist)
	for i := 0; i < ilen; i++ {
		Banned.AddBan(memlist[i])
	}
}
func botDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	return fmt.Sprintf("%2d Days, %2d Hours, %2d Mins.", h/24, h%24, m)
}
func (self *kickop) ceko(pelaku int64) bool {
	if !InArrayInt64(self.Opinv, pelaku) {
		self.Opinv = append(self.Opinv, pelaku)
		return true
	}
	return false
}

func (self *kickop) cek(pelaku string) bool {
	if !InArray2(self.Kick, pelaku) {
		self.Kick = append(self.Kick, pelaku)
		return true
	}
	return false
}

func (self *kickop) del(pelaku string) {
	self.Kick = Remove(self.Kick, pelaku)
}

func (self *kickop) ceki(pelaku string) bool {
	defer oop.PanicOnly()
	if !InArray2(self.Inv, pelaku) {
		self.Inv = append(self.Inv, pelaku)
		return true
	}
	return false
}

func (self *kickop) deli(pelaku string) {
	self.Inv = Remove(self.Inv, pelaku)
}

func (self *kickop) clear() {
	self.Inv = []string{}
	self.Kick = []string{}
	self.Opinv = []int64{}
}

func Checklistexpel(client *oop.Account, to string, targets []string, pl int, sinder string) {
	Room := oop.GetRoom(to)
	if len(targets) > 1 {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if InArray2(MAKERS, from) {
				if !InArray2(MAKERS, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetCreator(from) {
				if InArray2(MAKERS, sinder) {
					target = append(target, from)
					UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetSeller(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = Remove(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = Remove(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Expeled from Buyer\n"
			} else if pl == 2 {
				list += "Expeled from Owner\n"
			} else if pl == 3 {
				list += "Expeled from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				logAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				logAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				logAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				logAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				logAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				logAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				logAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, list)
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, list)
		}
	} else {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if InArray2(MAKERS, from) {
				if !InArray2(MAKERS, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetCreator(from) {
				if InArray2(MAKERS, sinder) {
					target = append(target, from)
					UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetSeller(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = Remove(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = Remove(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Removed from Buyer\n"
			} else if pl == 2 {
				list += "Removed from Owner\n"
			} else if pl == 3 {
				list += "Removed from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				logAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				logAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				logAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				logAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				logAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				logAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				logAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, list)
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, list)
		}
	}
}
func Checkqr() {
	Qrwar = true
	time.Sleep(1 * time.Second)
	Qrwar = false
}
func Cmdlistcheck() string {
	list2 := "✠ 𝗟𝗶𝘀𝘁 𝗖𝗺𝗱 :\n\n"
	list := ""
	if Commands.Botname != "" {
		list += fmt.Sprintf(" - Botname: %s\n", Commands.Botname)
	}
	if Commands.Upallimage != "" {
		list += fmt.Sprintf(" - Upallimage: %s\n", Commands.Upallimage)
	}
	if Commands.Upallcover != "" {
		list += fmt.Sprintf(" - Upallcover: %s\n", Commands.Upallcover)
	}
	if Commands.Unsend != "" {
		list += fmt.Sprintf(" - Unsend: %s\n", Commands.Unsend)
	}
	if Commands.Upvallimage != "" {
		list += fmt.Sprintf(" - Upvallimage: %s\n", Commands.Upvallimage)
	}
	if Commands.Upvallcover != "" {
		list += fmt.Sprintf(" - Upvallcover: %s\n", Commands.Upvallcover)
	}
	if Commands.Appname != "" {
		list += fmt.Sprintf(" - Appname: %s\n", Commands.Appname)
	}
	if Commands.Useragent != "" {
		list += fmt.Sprintf(" - Useragent: %s\n", Commands.Useragent)
	}
	if Commands.Hostname != "" {
		list += fmt.Sprintf(" - Hostname: %s\n", Commands.Hostname)
	}
	if Commands.Friends != "" {
		list += fmt.Sprintf(" - Friends: %s\n", Commands.Friends)
	}
	if Commands.Adds != "" {
		list += fmt.Sprintf(" - Adds: %s\n", Commands.Adds)
	}
	if Commands.Limits != "" {
		list += fmt.Sprintf(" - Limits: %s\n", Commands.Limits)
	}
	if Commands.Addallbots != "" {
		list += fmt.Sprintf(" - Addallbots: %s\n", Commands.Addallbots)
	}
	if Commands.Addallsquads != "" {
		list += fmt.Sprintf(" - Addallsquads: %s\n", Commands.Addallsquads)
	}
	if Commands.Leave != "" {
		list += fmt.Sprintf(" - Leave: %s\n", Commands.Leave)
	}
	if Commands.Respon != "" {
		list += fmt.Sprintf(" - Respon: %s\n", Commands.Respon)
	}
	if Commands.Ping != "" {
		list += fmt.Sprintf(" - Ping: %s\n", Commands.Ping)
	}
	if Commands.Count != "" {
		list += fmt.Sprintf(" - Count: %s\n", Commands.Count)
	}
	if Commands.Limitout != "" {
		list += fmt.Sprintf(" - 1111111: %s\n", Commands.Limitout)
	}
	if Commands.Access != "" {
		list += fmt.Sprintf(" - Access: %s\n", Commands.Access)
	}
	if Commands.Allbanlist != "" {
		list += fmt.Sprintf(" - Allbanlist: %s\n", Commands.Allbanlist)
	}
	if Commands.Gaccess != "" {
		list += fmt.Sprintf(" - Gaccess: %s\n", Commands.Gaccess)
	}
	if Commands.Checkram != "" {
		list += fmt.Sprintf(" - Checkram: %s\n", Commands.Checkram)
	}
	if Commands.Backups != "" {
		list += fmt.Sprintf(" - Backups: %s\n", Commands.Backups)
	}
	if Commands.Upimage != "" {
		list += fmt.Sprintf(" - Upimage: %s\n", Commands.Upimage)
	}
	if Commands.Upcover != "" {
		list += fmt.Sprintf(" - Upcover: %s\n", Commands.Upcover)
	}
	if Commands.Upvimage != "" {
		list += fmt.Sprintf(" - Upvimage: %s\n", Commands.Upvimage)
	}
	if Commands.Upvcover != "" {
		list += fmt.Sprintf(" - Upvcover: %s\n", Commands.Upvcover)
	}
	if Commands.Bringall != "" {
		list += fmt.Sprintf(" - Bringall: %s\n", Commands.Bringall)
	}
	if Commands.Purgeall != "" {
		list += fmt.Sprintf(" - Purgeall: %s\n", Commands.Purgeall)
	}
	if Commands.Banlist != "" {
		list += fmt.Sprintf(" - Banlist: %s\n", Commands.Banlist)
	}
	if Commands.Clearban != "" {
		list += fmt.Sprintf(" - Clearban: %s\n", Commands.Clearban)
	}
	if Commands.Stayall != "" {
		list += fmt.Sprintf(" - Stayall: %s\n", Commands.Stayall)
	}
	if Commands.Clearchat != "" {
		list += fmt.Sprintf(" - Clearchat: %s\n", Commands.Clearchat)
	}
	if Commands.Here != "" {
		list += fmt.Sprintf(" - Here: %s\n", Commands.Here)
	}
	if Commands.Speed != "" {
		list += fmt.Sprintf(" - Speed: %s\n", Commands.Speed)
	}
	if Commands.Status != "" {
		list += fmt.Sprintf(" - Status: %s\n", Commands.Status)
	}
	if Commands.Tagall != "" {
		list += fmt.Sprintf(" - Tagall: %s\n", Commands.Tagall)
	}
	if Commands.Kick != "" {
		list += fmt.Sprintf(" - Kick: %s\n", Commands.Kick)
	}
	if Commands.Max != "" {
		list += fmt.Sprintf(" - Protect Max: %s\n", Commands.Max)
	}
	if Commands.None != "" {
		list += fmt.Sprintf(" - Protect None: %s\n", Commands.None)
	}
	if Commands.Kickall != "" {
		list += fmt.Sprintf(" - Kickall: %s\n", Commands.Kickall)
	}
	if Commands.Cancelall != "" {
		list += fmt.Sprintf(" - Cancelall: %s\n", Commands.Cancelall)
	}
	if list != "" {
		return list2 + list

	} else {
		return "Not found set Cmd.\n"
	}
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func TimeDown(Fucking int) bool {
	switch Fucking {
	case 0:
		time.Sleep(200 * time.Millisecond)
		return true
	case 1:
		time.Sleep(400 * time.Millisecond)
		return true
	case 2:
		time.Sleep(600 * time.Millisecond)
		return true
	case 3:
		time.Sleep(800 * time.Millisecond)
		return true
	case 4:
		time.Sleep(1000 * time.Millisecond)
		return true
	case 5:
		time.Sleep(1200 * time.Millisecond)
		return true
	case 6:
		time.Sleep(1400 * time.Millisecond)
		return true
	case 7:
		time.Sleep(1600 * time.Millisecond)
		return true
	case 8:
		time.Sleep(1800 * time.Millisecond)
		return true
	case 9:
		time.Sleep(2000 * time.Millisecond)
		return true
	case 10:
		time.Sleep(2200 * time.Millisecond)
		return true
	case 11:
		time.Sleep(2400 * time.Millisecond)
		return true
	case 12:
		time.Sleep(2600 * time.Millisecond)
		return true
	case 13:
		time.Sleep(2800 * time.Millisecond)
		return true
	case 14:
		time.Sleep(3000 * time.Millisecond)
		return true
	case 15:
		time.Sleep(3200 * time.Millisecond)
		return true
	case 16:
		time.Sleep(3400 * time.Millisecond)
		return true
	case 17:
		time.Sleep(3600 * time.Millisecond)
		return true
	case 18:
		time.Sleep(3800 * time.Millisecond)
		return true
	case 19:
		time.Sleep(4000 * time.Millisecond)
		return true
	case 20:
		time.Sleep(4200 * time.Millisecond)
		return true
	case 21:
		time.Sleep(4400 * time.Millisecond)
		return true
	case 22:
		time.Sleep(4600 * time.Millisecond)
		return true
	case 23:
		time.Sleep(4800 * time.Millisecond)
		return true
	default:
		return false
	}
}
func LeaveallGroups(client *oop.Account, to string) []string {
	allg := []string{}
	for i := range ClientBot {
		groups, _ := ClientBot[i].GetGroupIdsJoined()
		grup, _ := ClientBot[i].GetGroups(groups)
		for _, gi := range grup {
			if gi.ChatMid != to {
				ClientBot[i].LeaveGroup(gi.ChatMid)
				time.Sleep(1 * time.Second)
				if !InArray2(allg, gi.ChatMid) {
					allg = append(allg, gi.ChatMid)
				}
			}
		}
	}
	return allg
}
func logAccess(client *oop.Account, group, from, tipe string, targets []string, tempat int64) {
	defer panicHandle("logAccess")
	if !LogMode || SendMyseller(from) {
		return
	}
	nm, _, _ := client.GetChatList(group)
	var ts = ""
	if tipe == "ban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! banned %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! banned %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! unbaned %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! unbaned %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "owner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! promoted owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! promoted owner %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unowner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! expeled owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! expeled owner %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "bot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! promoted bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! promoted bot %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unbot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! expeled bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! expeled bot %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "mute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! muted %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! muted %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unmute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! unmuted %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! unmuted %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "fuck" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  fuck %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  fuck %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "master" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  promoted Master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  promoted Master %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unmaster" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled Master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled Master %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "admin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  promoted admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  promoted admin %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unadmin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled admin %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "gowner" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  promoted gowner %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "ungowner" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  expeled gowner %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "gadmin" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  promoted gadmin %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "ungadmin" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  expeled gadmin %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "expel" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled access %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled access %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearowner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all owner %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearmaster" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all master %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearadmin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all admin %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all banlist %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all banlist %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearbot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all bot %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearmute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all mutelist %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all mutelist %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "kick" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  kick %v user's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "cancel" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  cancel %v invitation's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "invite" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  invite %v user's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "Kickall" {
		ts += fmt.Sprintf("@!  Nukeall %v user's from\n%s", len(targets), nm)
	} else if tipe == "purgeall" {
		ts += fmt.Sprintf("@!  purgeall %v user's from\n%s", len(targets), nm)
	} else if tipe == "purge" {
		ts += fmt.Sprintf("@!  purge %v user's in:\n%s", len(targets), nm)
	} else if tipe == "cancelall" {
		ts += fmt.Sprintf("@!  cancelall %v invitation's from\n%s", len(targets), nm)
	} else if tipe == "leave" {
		ts += fmt.Sprintf("@!  bot's leave from\n%s", nm)
	} else if tipe == "bringbot" {
		ts += fmt.Sprintf("@!  Invite bot's\n%s", nm)
	} else if tipe == "addfrind" {
		ts += fmt.Sprintf("@! he added %v as friend\nMid : \n%s", group, from)
	}
	room := oop.GetRoom(LogGroup)
	if len(room.Client) != 0 {
		exe, err := SelectBot(room.Client[0], LogGroup)
		if err == nil {
			if exe != nil {
				exe.SendMention(LogGroup, ts, []string{from})
			}
		} else {
			LogMode = false
			LogGroup = ""
		}
	}
}
func Upsetcmd(text string, text2 string) string {
	count := 0
	if text == "rollcall" {
		Commands.Botname = text2
		count = count + 1
	} else if text == "upallimage" {
		Commands.Upallimage = text2
		count = count + 1
	} else if text == "upallcover" {
		Commands.Upallcover = text2
		count = count + 1
	} else if text == "unsend" {
		Commands.Unsend = text2
		count = count + 1
	} else if text == "upvallimage" {
		Commands.Upvallimage = text2
		count = count + 1
	} else if text == "upvallcover" {
		Commands.Upvallcover = text2
		count = count + 1
	} else if text == "appname" {
		Commands.Appname = text2
		count = count + 1
	} else if text == "useragent" {
		Commands.Useragent = text2
		count = count + 1
	} else if text == "hostname" {
		Commands.Hostname = text2
		count = count + 1
	} else if text == "friends" {
		Commands.Friends = text2
		count = count + 1
	} else if text == "adds" {
		Commands.Adds = text2
		count = count + 1
	} else if text == "limits" {
		Commands.Limits = text2
		count = count + 1
	} else if text == "addallbots" {
		Commands.Addallbots = text2
		count = count + 1
	} else if text == "addallsquads" {
		Commands.Addallsquads = text2
		count = count + 1
	} else if text == "leave" {
		Commands.Leave = text2
		count = count + 1
	} else if text == "respon" {
		Commands.Respon = text2
		count = count + 1
	} else if text == "ping" {
		Commands.Ping = text2
		count = count + 1
	} else if text == "count" {
		Commands.Count = text2
		count = count + 1
	} else if text == "limitout" {
		Commands.Limitout = text2
		count = count + 1
	} else if text == "access" {
		Commands.Access = text2
		count = count + 1
	} else if text == "allbanlist" {
		Commands.Allbanlist = text2
		count = count + 1
	} else if text == "gaccess" {
		Commands.Gaccess = text2
		count = count + 1
	} else if text == "checkram" {
		Commands.Checkram = text2
		count = count + 1
	} else if text == "upimage" {
		Commands.Upimage = text2
		count = count + 1
	} else if text == "upcover" {
		Commands.Upcover = text2
		count = count + 1
	} else if text == "upvimage" {
		Commands.Upvimage = text2
		count = count + 1
	} else if text == "upvcover" {
		Commands.Upvcover = text2
		count = count + 1
	} else if text == "Purgeall" {
		Commands.Purgeall = text2
		count = count + 1
	} else if text == "banlist" {
		Commands.Banlist = text2
		count = count + 1
	} else if text == "clearban" {
		Commands.Clearban = text2
		count = count + 1
	} else if text == "bringall" {
		Commands.Bringall = text2
		count = count + 1
	} else if text == "stayall" {
		Commands.Stayall = text2
		count = count + 1
	} else if text == "clearchat" {
		Commands.Clearchat = text2
		count = count + 1
	} else if text == "here" {
		Commands.Here = text2
		count = count + 1
	} else if text == "speed" {
		Commands.Speed = text2
		count = count + 1
	} else if text == "status" {
		Commands.Status = text2
		count = count + 1
	} else if text == "tagall" {
		Commands.Tagall = text2
		count = count + 1
	} else if text == "kick" {
		Commands.Kick = text2
		count = count + 1
	} else if text == "max" {
		Commands.Max = text2
		count = count + 1
	} else if text == "none" {
		Commands.None = text2
		count = count + 1
	} else if text == "kickall" {
		Commands.Kickall = text2
		count = count + 1
	} else if text == "cancelall" {
		Commands.Cancelall = text2
		count = count + 1
	}
	if count != 0 {
		kowe := text
		jancuk := text2
		newsend := "Changed cmd: " + kowe + " to " + jancuk + "\n"
		return newsend
	}
	return ""
}
func checkunbanbots(client *oop.Account, to string, targets []string, pl int, sinder string) {
	room := oop.GetRoom(to)
	target := []string{}
	for _, from := range targets {
		if Banned.GetFuck(from) {
			target = append(target, from)
			Banned.DelFuck(from)
		} else if Banned.GetBan(from) {
			target = append(target, from)
			Banned.DelBan(from)
		} else if InArray2(room.Gban, from) {
			target = append(target, from)
			Ungban(to, from)
		} else if Banned.GetMute(from) {
			target = append(target, from)
			Banned.DelMute(from)
		}
	}
	if len(target) != 0 {
		list := ""
		if pl == 1 {
			list += "Removed from banlist:\n"
		} else if pl == 2 {
			list += "Removed from fucklist:\n"
		} else if pl == 3 {
			list += "Removed from gbanlist:\n"
		} else if pl == 4 {
			list += "Removed from mutelist:\n"
		}
		for i := range target {
			list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
		}
		client.SendPollMention(to, list, target)
		if pl == 1 {
			logAccess(client, to, sinder, "unban", target, 2)
		} else if pl == 2 {
			logAccess(client, to, sinder, "unfuck", target, 2)
		} else if pl == 3 {
			logAccess(client, to, sinder, "ungban", target, 2)
		} else if pl == 4 {
			logAccess(client, to, sinder, "unmute", target, 2)
		}
	} else {
		list := ""
		if pl == 1 {
			list += "User(s) not in banlist.\n"
		} else if pl == 2 {
			list += "User(s) not in fucklist.\n"
		} else if pl == 3 {
			list += "User(s) not in gbanlist.\n"
		} else if pl == 4 {
			list += "User(s) not in mutelist.\n"
		}
		client.SendMessage(to, list)
	}
}
func CheckExprd(s *oop.Account, to string, sender string) bool {
	base := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
	d := fmt.Sprintf("%v", Data.Dalltime)
	has := strings.Split(d, "-")
	has2 := strings.Split(has[2], "T")
	yy, _ := strconv.Atoi(has[0])
	mm, _ := strconv.Atoi(has[1])
	timeup, _ := strconv.Atoi(has2[0])
	batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
	if batas.Before(base) {
		if !SendMycreator(sender) {
			s.SendMessage(to, "Sorry your bots is expired, Please Contact with our Creator to renew your squad. ;-)")
			return false
		}
		return true
	}
	return true
}
func fmtDurations(d time.Duration) string {
	d = d.Round(time.Second)
	x := d
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	if x < 60*time.Second {
		return fmt.Sprintf("%v", x)
	} else if x < 3600*time.Second {
		return fmt.Sprintf("%02dMin", m)
	} else if x < 86400*time.Second {
		return fmt.Sprintf("%02dH %02dMin", h%24, m)
	} else {
		return fmt.Sprintf("%02dD %02dH %02dMin", h/24, h%24, m)
	}
}
func CekDuedate() time.Time {
	bod := string(Data.Dalltime)
	date, _ := time.Parse(time.RFC3339, bod)
	return date
}
func CheckLastActive(client *oop.Account, targets string) string {
	list := ""
	mek, tu := LastActive.Get(targets)
	if tu {
		asu := mek.(*talkservice.Operation)
		if asu.Type == 55 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			list += "- LastActive: " + sp + "\n- Type: Read Message\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 124 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			invites := strings.Split(asu.Param3, "\x1e")
			nos := 0
			her := ""
			for _, ampemng := range invites {
				nos += 1
				pr, _ := client.GetContact(ampemng)
				her += fmt.Sprintf("\n  %v. %v", nos, pr.DisplayName)
			}
			list += "- LastActive: " + sp + "\n- Type: Invited member\n- Group: " + names1 + "\n- Target: " + her + "\n\n"
		} else if asu.Type == 133 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- LastActive: " + sp + "\n- Type : Kick member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 126 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- LastActive: " + sp + "\n- Type: Cancel member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 26 {
			msg := asu.Message
			if msg.ToType == 2 {
				names1, _ := client.GetGroupMember(msg.To)
				cok := asu.CreatedTime / 1000
				i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
				tm := time.Unix(i, 0)
				ss := time.Since(tm)
				sp := fmtDuration(ss)
				tx := ""
				if msg.ContentType == 0 {
					tx = msg.Text
				} else {
					tx = "Non Text Message"
				}
				list += "- LastActive: " + sp + "\n- Type: Send Message\n- Group: " + names1 + "\n- Message: " + tx + "\n\n"
			}
		} else if asu.Type == 130 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			list += "- LastActive: " + sp + "\n- Type: Join Group\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 122 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			var ti string
			if asu.Param3 == "4" {
				g, _ := client.GetGroup3(asu.Param1)
				if g.Extra.GroupExtra.PreventedJoinByTicket == false {
					ti = "Open qr"
				} else {
					ti = "Close qr"
				}
			} else if asu.Param3 == "1" {
				ti = "Change Group Name"
			}
			list += "- LastActive: " + sp + "\n- Type: Update Group\n- Group: " + names1 + "\n- Type: " + ti + "\n\n"
		}
	}
	return list
}
func InArray2(ArrList []string, rstr string) bool {
	for _, x := range ArrList {
		if x == rstr {
			return true
		}
	}
	return false
}

func Checklistaccess(client *oop.Account, group string, targets []string, pl int, sinder string) {
	Room := oop.GetRoom(group)
	if pl == 12 {
		countr := 0
		countr1 := 0
		list := "Account Info: \n\n"
		for n, xx := range targets {
			new := client.Getcontactuser(xx)
			if new != nil {
				list += "Name: Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("Name: %v \n", x.DisplayName)
				status := "status: None\n\n"
				if InArray2(MAKERS, targets[n]) {
					status = "status: Makers\n\n"
				} else if UserBot.GetCreator(targets[n]) {
					status = "status: Creator\n\n"
				} else if UserBot.GetBuyer(targets[n]) {
					status = "status: Buyer\n\n"
				} else if UserBot.GetOwner(targets[n]) {
					status = "status: Owner\n\n"
				} else if UserBot.GetMaster(targets[n]) {
					status = "status: Master\n\n"
				} else if UserBot.GetAdmin(targets[n]) {
					status = "status: Admin\n\n"
				} else if InArray2(Room.Gowner, targets[n]) {
					status = "status: GroupOwnar\n\n"
				} else if InArray2(Room.Gadmin, targets[n]) {
					status = "status: GroupAdmin\n\n"
				} else if UserBot.GetBot(targets[n]) {
					status = "status: Bot\n\n"
				} else if Banned.GetFuck(targets[n]) {
					status = "status: Fuck\n\n"
				} else if Banned.GetBan(targets[n]) {
					status = "status: Ban\n\n"
				} else if Banned.GetMute(targets[n]) {
					status = "status: Mute\n\n"
				} else if InArray2(Room.Gban, targets[n]) {
					status = "status: Groupban\n\n"
				} else if InArray2(Squadlist, targets[n]) {
					status = "status: My team\n\n"
				} else if UserBot.GetSeller(targets[n]) {
					status = "status: My Seller\n\n"
				}
				list += status
				if !InArray2(checkHaid, targets[n]) {
					new := CheckLastActive(client, targets[n])
					list += new
				}
				listGroup := "\nMember of:\n"
				listPinde := "\nPending of:\n"
				grs, _ := client.GetGroupIdsJoined()
				groups, _ := client.GetGroups(grs)
				for _, x := range groups {
					if oop.IsMembers(client, x.ChatMid, targets[n]) == true {
						countr = countr + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listGroup += nm + "\n"
					}
					if oop.IsPending(client, x.ChatMid, targets[n]) == true {
						countr1 = countr1 + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listPinde += nm + "\n"
					}
				}
				if countr != 0 {
					list += fmt.Sprintf("Groups: %v\n", countr)

				} else {
					list += "Groups: 0\n"
				}
				if countr1 != 0 {
					list += fmt.Sprintf("Pendings: %v\n", countr1)
				} else {
					list += "Pendings: 0\n"
				}
				if countr != 0 {
					if !InArray2(checkHaid, targets[n]) {
						list += listGroup
					}
				}
				if countr1 != 0 {
					if !InArray2(checkHaid, targets[n]) {
						list += listPinde
					}
				}

			}
		}
		client.SendMessage(group, list)
	} else if pl == 16 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			new := client.Getcontactuser(xx)
			if new != nil {
				list += rengs + ". Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

			}
		}
		client.SendMessage(group, list)
	} else if pl == 14 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			new := client.Getcontactuser(xx)
			if new != nil {
				list += rengs + ". Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("%v. %v\n_%v\n", n+1, x.DisplayName, targets[n])

			}
		}
		client.SendMessage(group, list)
	} else {
		if len(targets) > 1 {
			creator := []string{}
			buyer := []string{}
			owner := []string{}
			master := []string{}
			admin := []string{}
			gowner := []string{}
			gadmin := []string{}
			squad := []string{}
			bot := []string{}
			ban := []string{}
			fuck := []string{}
			mute := []string{}
			Gban := []string{}
			Glist := []string{}
			Maker := []string{}
			Seller := []string{}
			for _, from := range targets {
				if MemUser(group, from) && !MemBan2(group, from) {
					if !InArray2(Glist, from) {
						Glist = append(Glist, from)
					}
				} else if UserBot.GetCreator(from) {
					creator = append(creator, from)
				} else if UserBot.GetSeller(from) {
					Seller = append(Seller, from)
				} else if InArray2(MAKERS, from) {
					Maker = append(Maker, from)
				} else if UserBot.GetBuyer(from) {
					buyer = append(buyer, from)
				} else if UserBot.GetOwner(from) {
					owner = append(owner, from)
				} else if UserBot.GetMaster(from) {
					master = append(master, from)
				} else if UserBot.GetAdmin(from) {
					admin = append(admin, from)
				} else if InArray2(Room.Gowner, from) {
					gowner = append(gowner, from)
				} else if InArray2(Room.Gadmin, from) {
					gadmin = append(gadmin, from)
				} else if UserBot.GetBot(from) {
					bot = append(bot, from)
				} else if Banned.GetFuck(from) {
					fuck = append(fuck, from)
				} else if Banned.GetBan(from) {
					ban = append(ban, from)
				} else if Banned.GetMute(from) {
					mute = append(mute, from)
				} else if InArray2(Room.Gban, from) {
					Gban = append(Gban, from)
				} else if InArray2(Squadlist, from) {
					squad = append(squad, from)
				}
			}
			list2 := ""
			if len(Glist) != 0 {
				if pl == 1 {
					list2 += "Promoted as Buyer:\n\n"
				} else if pl == 2 {
					list2 += "Promoted as Owner:\n\n"
				} else if pl == 3 {
					list2 += "Promoted as Master:\n\n"
				} else if pl == 4 {
					list2 += "Promoted as Admin:\n\n"
				} else if pl == 5 {
					list2 += "Promoted as Bot:\n\n"
				} else if pl == 6 {
					list2 += "Promoted as Gowner:\n\n"
				} else if pl == 7 {
					list2 += "Promoted as Gadmin\n\n"
				} else if pl == 8 {
					list2 += "Added to banlist:\n\n"
				} else if pl == 9 {
					list2 += "Added to fucklist:\n\n"
				} else if pl == 10 {
					list2 += "Added to gbanlist:\n\n"
				} else if pl == 11 {
					list2 += "Added to mutelist:\n\n"
				} else if pl == 13 {
					list2 += "Added to Creatorlist:\n\n"
				} else if pl == 17 {
					list2 += "Added to Sellerlist:\n\n"
				}
				for n, xx := range Glist {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list2 += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list2 += rengs + ". " + x.DisplayName + "\n"
						if pl == 1 {
							UserBot.AddBuyer(xx)
						} else if pl == 2 {
							UserBot.AddOwner(xx)
						} else if pl == 13 {
							UserBot.AddCreator(xx)
						} else if pl == 3 {
							UserBot.AddMaster(xx)
						} else if pl == 4 {
							UserBot.AddAdmin(xx)
						} else if pl == 5 {
							UserBot.AddBot(xx)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, xx)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, xx)
						} else if pl == 8 {
							Banned.AddBan(xx)
						} else if pl == 9 {
							Banned.AddBan(xx)
						} else if pl == 10 {
							Addgban(xx, group)
						} else if pl == 11 {
							Banned.AddBan(xx)
						} else if pl == 17 {
							UserBot.AddSeller(xx)
						}
					}
				}
				if pl == 2 {
					logAccess(client, group, sinder, "owner", Glist, 2)
				} else if pl == 3 {
					logAccess(client, group, sinder, "master", Glist, 2)
				} else if pl == 4 {
					logAccess(client, group, sinder, "admin", Glist, 2)
				} else if pl == 5 {
					logAccess(client, group, sinder, "bot", Glist, 2)
				} else if pl == 6 {
					logAccess(client, group, sinder, "gowner", Glist, 2)
				} else if pl == 7 {
					logAccess(client, group, sinder, "gadmin", Glist, 2)
				} else if pl == 8 {
					logAccess(client, group, sinder, "ban", Glist, 2)
				} else if pl == 9 {
					logAccess(client, group, sinder, "fuck", Glist, 2)
				} else if pl == 10 {
					logAccess(client, group, sinder, "gban", Glist, 2)
				} else if pl == 11 {
					logAccess(client, group, sinder, "mute", Glist, 2)
				}
			}
			list := "Users have access:\n"
			if len(creator) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘁𝗲𝗮𝗺:\n"
				for n, xx := range creator {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Seller) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗦𝗲𝗹𝗹𝗲𝗿:\n"
				for n, xx := range Seller {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Maker) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘁𝗲𝗮𝗺:\n"
				for n, xx := range Maker {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(buyer) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗯𝘂𝘆𝗲𝗿𝘀:\n"
				for n, xx := range buyer {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(owner) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗼𝘄𝗻𝗲𝗿𝘀:\n"
				for n, xx := range owner {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(master) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗺𝗮𝘀𝘁𝗲𝗿𝘀:\n"
				for n, xx := range master {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(admin) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗮𝗱𝗺𝗶𝗻𝘀:\n"
				for n, xx := range admin {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gowner) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗼𝘄𝗻𝗲𝗿𝘀:\n"
				for n, xx := range gowner {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gadmin) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗮𝗱𝗺𝗶𝗻𝘀:\n"
				for n, xx := range gadmin {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(bot) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗯𝗼𝘁𝗹𝗶𝘀𝘁\n"
				for n, xx := range bot {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(squad) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘀𝗾𝘂𝗮𝗱:\n"
				for n, xx := range squad {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(ban) != 0 {
				list += "𝗘𝘅𝗶𝘀𝘁.𝗶𝗻 𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
				for n, xx := range ban {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(fuck) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗳𝘂𝗰𝗸𝗹𝗶𝘀𝘁:\n"
				for n, xx := range fuck {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Gban) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n\n"
				for n, xx := range Gban {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(mute) != 0 {
				list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁:\n\n"
				for n, xx := range mute {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if list != "Users have access:\n" {
				if list2 != "" {
					list2 += "\n"
				}
				client.SendMessage(group, list2+list)
			} else {
				client.SendMessage(group, list2)
			}
		} else {
			list := ""
			for n, from := range targets {
				if InArray2(MAKERS, from) {
					list += "User have access exist in taem list."
				} else if UserBot.GetCreator(from) {
					list += "User have access exist in Creator list."
				} else if UserBot.GetSeller(from) {
					list += "User have access exist in seller list."
				} else if UserBot.GetBuyer(from) {
					list += "User have access exist in buyer list."
				} else if UserBot.GetOwner(from) {
					list += "User have access exist in owner list."
				} else if UserBot.GetMaster(from) {
					list += "User have access exist in master list."
				} else if UserBot.GetAdmin(from) {
					list += "User have access exist in admin list."
				} else if InArray2(Room.Gowner, from) {
					list += "User have access exist in gowner list."
				} else if InArray2(Room.Gadmin, from) {
					list += "User have access exist in gadmin list."
				} else if UserBot.GetBot(from) {
					list += "User have access exist in bot list."
				} else if Banned.GetFuck(from) {
					list += "User have access exist in fuck list."
				} else if Banned.GetBan(from) {
					list += "User have access exist in ban list."
				} else if InArray2(Room.Gban, from) {
					list += "User have access exist in gban list."
				} else if InArray2(Squadlist, from) {
					list += "User have access exist in squad list."
				} else if Banned.GetMute(from) {
					list += "User have access exist in mute list."
				} else if MemUser(group, from) && !MemBan2(group, from) {
					if pl == 1 {
						list += "Promoted as Buyer:\n"
					} else if pl == 2 {
						list += "Promoted as Owner:\n"
					} else if pl == 3 {
						list += "Promoted as Master:\n"
					} else if pl == 4 {
						list += "Promoted as Admin:\n"
					} else if pl == 5 {
						list += "Promoted as Bot:\n"
					} else if pl == 6 {
						list += "Promoted as Gowner:\n"
					} else if pl == 7 {
						list += "Promoted as Gadmin:\n"
					} else if pl == 8 {
						list += "Added to banlist:\n"
					} else if pl == 9 {
						list += "Added to fucklist:\n"
					} else if pl == 10 {
						list += "Added to gbanlist:\n"
					} else if pl == 11 {
						list += "Added to mutelist:\n"
					} else if pl == 13 {
						list += "Added to Creatorlist:\n"
					} else if pl == 17 {
						list += "Added to Sellerlist:\n"
					}
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(from)
					if new != nil {
						list += "\n   " + rengs + ". Closed Account"
					} else {
						x, _ := client.GetContact(from)
						list += "\n   " + rengs + ". " + x.DisplayName
						if pl == 1 {
							UserBot.AddBuyer(from)
						} else if pl == 2 {
							UserBot.AddOwner(from)
						} else if pl == 3 {
							UserBot.AddMaster(from)
						} else if pl == 4 {
							UserBot.AddAdmin(from)
						} else if pl == 5 {
							UserBot.AddBot(from)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, from)
						} else if pl == 13 {
							UserBot.AddCreator(from)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, from)
						} else if pl == 8 {
							autokickban(client, group, from)
							Banned.AddBan(from)
						} else if pl == 9 {
							Banned.AddFuck(from)
						} else if pl == 10 {
							Addgban(from, group)
						} else if pl == 11 {
							Banned.AddMute(from)
						} else if pl == 17 {
							UserBot.AddSeller(from)
						}
					}
					if pl == 2 {
						logAccess(client, group, sinder, "owner", []string{from}, 2)
					} else if pl == 3 {
						logAccess(client, group, sinder, "master", []string{from}, 2)
					} else if pl == 4 {
						logAccess(client, group, sinder, "admin", []string{from}, 2)
					} else if pl == 5 {
						logAccess(client, group, sinder, "bot", []string{from}, 2)
					} else if pl == 6 {
						logAccess(client, group, sinder, "gowner", []string{from}, 2)
					} else if pl == 7 {
						logAccess(client, group, sinder, "gadmin", []string{from}, 2)
					} else if pl == 8 {
						logAccess(client, group, sinder, "ban", []string{from}, 2)
					} else if pl == 9 {
						logAccess(client, group, sinder, "fuck", []string{from}, 2)
					} else if pl == 10 {
						logAccess(client, group, sinder, "gban", []string{from}, 2)
					} else if pl == 11 {
						logAccess(client, group, sinder, "mute", []string{from}, 2)
					}
				}

			}
			client.SendMessage(group, list)
		}
	}
}

func GetSquad(tok *oop.Account, to string) []*oop.Account {
	defer panicHandle("GetSquad")
	nm, memlist, invitee := tok.GetChatList(to)
	Bots := []*oop.Account{}
	MIdbot := []string{}
	GoClint := []*oop.Account{}
	Gomid := []string{}
	for _, ym := range memlist {
		if InArray2(Squadlist, ym) {
			idx := GetKorban(ym)
			MIdbot = append(MIdbot, ym)
			Bots = append(Bots, idx)
		}
	}
	room := oop.GetRoom(to)
	room.Name = nm
	for _, ym := range invitee {
		if InArray2(Squadlist, ym) {
			Gomid = append(Gomid, ym)
			idx := GetKorban(ym)
			GoClint = append(GoClint, idx)
		}
	}
	room.AddSquad(MIdbot, Bots, GoClint, Gomid)
	return Bots
}
func AutojoinQr(client *oop.Account, to string) {
	defer panicHandle("AutojoinQr")
	ti, err := client.ReissueChatTicket(to)
	if err == nil {
		go client.UpdateChatQrV2(to, false)
		all := []*oop.Account{}
		room := oop.GetRoom(to)
		cuk := room.Client
		for _, x := range ClientBot {
			if !oop.InArrayCl(cuk, x) && !oop.InArrayCl(oop.KickBans, x) && !oop.InArrayCl(room.GoClient, x) {
				all = append(all, x)
			}
		}
		sort.Slice(all, func(i, j int) bool {
			return all[i].KickPoint < all[j].KickPoint
		})
		var wg sync.WaitGroup
		wi := GetSquad(client, to)
		for i := 0; i < len(all); i++ {
			l := all[i]
			if l != client && !oop.InArrayCl(wi, l) {
				wg.Add(1)
				go func() {
					l.AcceptTicket(to, ti)
					wg.Done()
				}()
			}
		}
		wg.Wait()
		client.UpdateChatQrV2(to, true)
		GetSquad(client, to)
	}
}
func qrGo(cl *oop.Account, cans []*oop.Account, to string) {
	defer panicHandle("QR_go")
	Room := oop.GetRoom(to)
	mes := make(chan bool)
	go func() {
		err := cl.UpdateChatQrV2(to, false)
		if err != nil {
			mes <- false
		} else {
			mes <- true
		}
	}()
	Room.Qr = false
	var ticket string
	link, err := cl.ReissueChatTicket(to)
	if err == nil {
		ticket = link
	} else {
		ticket = "error"
	}
	var wg sync.WaitGroup
	if ticket != "error" && ticket != "" {
		ok := <-mes
		if !ok {
			return
		}
		for _, cc := range cans {
			wg.Add(1)
			go func(c *oop.Account) {
				err := c.AcceptTicket(to, ticket)
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}(cc)
		}
		wg.Wait()
		Room.Qr = true
	}
	if Room.Qr {
		go func() {
			err := cl.UpdateChatQrV2(to, true)
			if err != nil {
				mes <- true
			} else {
				mes <- false
			}
		}()
	}
}
func hstg(to, u string) {
	room := oop.GetRoom(to)
	if !InArray2(room.LeaveBack, u) {
		room.LeaveBack = append(room.LeaveBack, u)
	}
}
func RunBot(client *oop.Account, ch chan int) {
	defer panicHandle("RunBot")
	runtime.GOMAXPROCS(cpu)
	client.Revision = -1
	for {
		multiFunc, err := client.FetchOps(25)
		if err != nil || len(multiFunc) == 0 {
			continue
		}
		go func(fetch []*talkservice.Operation) {
			for _, op := range multiFunc {
				if op.Type == 124 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					rngcmd := GetComs(4, "invitebot")
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := oop.GetRoom(Group)
					if InArray2(invited, client.MID) {
						if oop.IoGOBot(Group, client) {
							if InArray2(client.Squads, user) {
								client.AcceptGroupInvitationNormal(Group)
								if client.Limited == false {
									if !InArrayInt64(cekGo, Optime) {
										cekGo = append(cekGo, Optime)
										AcceptJoin(client, Group)
									}
								}
							} else if UserBot.GetBot(user) {
								client.AcceptGroupInvitationNormal(Group)
								if client.Limited == false {
									if !InArrayInt64(cekGo, Optime) {
										cekGo = append(cekGo, Optime)
										AcceptJoin(client, Group)
									}
								}
							} else if GetCodeprem(rngcmd, user, Group) {
								client.AcceptGroupInvitationNormal(Group)
								if client.Limited == false {
									if !InArrayInt64(cekGo, Optime) {
										cekGo = append(cekGo, Optime)
										AcceptJoin(client, Group)
									}
								}
							} else {
								grs, _ := client.GetGroupIdsJoined()
								if InArray2(grs, Group) {
									client.LeaveGroup(Group)
									fl, _ := client.GetAllContactIds()
									if InArray2(fl, user) {
										client.UnFriend(user)
									}
								}
							}
						}
					} else {
						Optime := op.CreatedTime
						if Room.ProInvite {
							if MemUser(Group, user) {
								go func() {
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										Banned.AddBan(user)
									}
								}()
								if filterWar.ceko(Optime) {
									Room.ListInvited = invited
									CancelProtect(client, invited, Group)
								}
							} else {
								if filterWar.ceko(Optime) {
									cancelallcek(client, invited, Group)
								}
							}
						} else {
							if MemBan(Group, user) {
								go func() {
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
									}
								}()
								if filterWar.ceko(Optime) {
									BanAll(invited)
									cancelall(client, invited, Group)
								}
							} else {
								if MemUser(Group, user) {
									go func() {
										if filterWar.ceki(user) {
											for _, vo := range invited {
												if MemBan(Group, vo) {
													kickPelaku(client, Group, user)
													Banned.AddBan(user)
													break
												}
											}
											filterWar.deli(user)
										}
									}()
									if filterWar.ceko(Optime) {
										cancelallcek(client, invited, Group)
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 133 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := oop.GetRoom(Group)
					if client.MID == Invited {
						oop.Gones(Group, client)
						if MemUser(Group, user) {
							Banned.AddBan(user)
						}
					} else if !InArray2(Room.GoMid, client.MID) {
						if InArray2(client.Squads, Invited) {
							if MemUser(Group, user) {
								if oop.IoGOBot(Group, client) {
									Banned.AddBan(user)
									go func() {
										if filterWar.cek(user) {
											groupBackupKick(client, Group, user, true)
											filterWar.del(user)
										}
									}()
									if filterWar.cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										filterWar.del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										if MemUser(Group, user) {
											Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProKick {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := Nkick.Get(user); !ok {
											Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											Banned.AddBan(user)
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								Banned.AddBan(user)
								_, memlist := client.GetGroupMember(Group)
								oke := []string{}
								for mid, _ := range memlist {
									if InArray2(Squadlist, mid) {
										oke = append(oke, mid)
									}
								}
								if len(oke) == 0 {
									if !InArrayInt64(cekGo, Optime) {
										cekGo = append(cekGo, Optime)
										cls := []*oop.Account{}
										Bot2 := Room.Bot
										bots := Room.HaveClient
										for n, cl := range Room.GoClient {
											if n < 2 {
												go cl.AcceptGroupInvitationNormal(Group)
												cls = append(cls, cl)
											}
										}
										cc := len(cls)
										if cc != 0 {
											Purgesip(Group, cls[0])
											if Autojoin == "qr" {
												qrGo(cls[0], bots, Group)
											} else if Autojoin == "invite" {
												cls[0].InviteIntoChatPollVer(Group, Bot2)
											}
											for _, cl := range cls {
												Room.ConvertGo(cl)
											}
										}
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 130 {
					runtime.GOMAXPROCS(cpu)
					Group, user := op.Param1, op.Param2
					Room := oop.GetRoom(Group)
					if oop.IoGOBot(Group, client) {
						if Room.ProJoin {
							if MemUser(Group, user) {
								if filterWar.ceki(user) {
									kickPelaku(client, Group, user)
									filterWar.deli(user)
									Banned.AddBan(user)
								}
							}
						} else {
							if MemBan(Group, user) {
								if MemUser(Group, user) {
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										Banned.AddBan(user)
									}
								}
							} else {
								if InArray2(Room.ListInvited, user) {
									if MemUser(Group, user) {
										if cekjoin(user) {
											kickPelaku(client, Group, user)
											deljoin(user)
											Room.ListInvited = Remove(Room.ListInvited, user)
										}
									} else {
										Room.ListInvited = Remove(Room.ListInvited, user)
									}
								} else {
									if Room.Welcome {
										if _, ok := cewel.Get(user); !ok {
											cewel.Set(user, 1)
											if cekjoin(user) {
												if !InArray2(Squadlist, user) {
													Room.WelsomeSet(client, Group, user)
												}
											}
										}
									}
								}

							}
						}
					}
					Optime := op.CreatedTime
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 122 {
					runtime.GOMAXPROCS(cpu)
					Group, user, invited := op.Param1, op.Param2, op.Param3
					Optime := op.CreatedTime
					Room := oop.GetRoom(Group)
					if client.Limited == false && oop.IoGOBot(Group, client) {
						if MemUser(Group, user) {
							if Room.ProQr || AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										go func() {
											cans := oop.Actor(Group)
											for _, cl := range cans {
												err := cl.UpdateChatQrV2(Group, true)
												if err == nil {
													break
												}
											}
										}()
										if filterWar.ceki(user) {
											kickPelaku(client, Group, user)
											filterWar.deli(user)
											Banned.AddBan(user)
										}
									}
								}
							} else if Room.ProName || AutoproN == true {
								if invited == "1" {
									if cekOp2(Optime) {
										go func() {
											cans := oop.Actor(Group)
											for _, cl := range cans {
												err := cl.UpdateChatName(Group, Room.Name)
												if err == nil {
													break
												}
											}
										}()
										if filterWar.ceki(user) {
											kickPelaku(client, Group, user)
											filterWar.deli(user)
										}
									}
								}
							} else {
								if MemBan(Group, user) {
									if invited == "4" {
										if cekOp2(Optime) {
											go func() {
												cans := oop.Actor(Group)
												for _, cl := range cans {
													err := cl.UpdateChatQrV2(Group, true)
													if err == nil {
														break
													}
												}
											}()
											if filterWar.ceki(user) {
												kickPelaku(client, Group, user)
												filterWar.deli(user)
												Banned.AddBan(user)
											}
										}
									} else if invited == "1" {
										if cekOp2(Optime) {
											go func() {
												cans := oop.Actor(Group)
												for _, cl := range cans {
													err := cl.UpdateChatName(Group, Room.Name)
													if err == nil {
														break
													}
												}
											}()
											if filterWar.ceki(user) {
												kickPelaku(client, Group, user)
												filterWar.deli(user)
											}
										}
									}
								}
							}

						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 126 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := oop.GetRoom(Group)
					if client.MID == Invited {
						oop.Gones(Group, client)
						if MemUser(Group, user) {
							Banned.AddBan(user)
						}
					} else if !InArray2(Room.GoMid, client.MID) {
						if InArray2(client.Squads, Invited) {
							if MemUser(Group, user) {
								if oop.IoGOBot(Group, client) {
									Banned.AddBan(user)
									go func() {
										if filterWar.cek(user) {
											groupBackupKick(client, Group, user, true)
											filterWar.del(user)
										}
									}()
									if filterWar.cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										filterWar.del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										if MemUser(Group, user) {
											Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProCancel {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := Nkick.Get(user); !ok {
											Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											Banned.AddBan(user)
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								Banned.AddBan(user)
								_, memlist := client.GetGroupMember(Group)
								oke := []string{}
								for mid, _ := range memlist {
									if InArray2(Squadlist, mid) {
										oke = append(oke, mid)
									}
								}
								if len(oke) == 0 {
									if !InArrayInt64(cekGo, Optime) {
										cekGo = append(cekGo, Optime)
										cls := []*oop.Account{}
										Bot := Room.Bot
										bots := Room.HaveClient
										for n, cl := range Room.GoClient {
											if n < 2 {
												go cl.AcceptGroupInvitationNormal(Group)
												cls = append(cls, cl)
											}
										}
										cc := len(cls)
										if cc != 0 {
											Purgesip(Group, cls[0])
											if Autojoin == "qr" {
												qrGo(cls[0], bots, Group)
											} else if Autojoin == "invite" {
												cls[0].InviteIntoChatPollVer(Group, Bot)
											}
											for _, cl := range cls {
												Room.ConvertGo(cl)
											}
										}
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 5 {
					Group := op.Param1
					if LogMode && !InArray2(client.Squads, Group) {
						logAccess(client, client.Namebot, Group, "addfrind", []string{}, 2)
					}
				} else if op.Type == 55 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if client.Limited == false && oop.IoGOBot(Group, client) {
						if cekOp(Optime) {
							if MemBan(Group, user) {
								kickPelaku(client, Group, user)
							} else {
								Room := oop.GetRoom(Group)
								if Room.Lurk && !InArray2(checkHaid, user) {
									Room.CheckLurk(client, Group, user)
								}
							}
						}

					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 26 {
					msg := op.Message
					Optime := op.CreatedTime
					if msg.ContentType != 18 {
						if _, ok := Command.Get(Optime); !ok {
							Command.Set(Optime, client)
							if _, ok := filterop.Get(Optime); !ok {
								filterop.Set(Optime, 1)
								Bot(op, client, ch)
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 128 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if !InArray2(Squadlist, user) {
						Room := oop.GetRoom(Group)
						if Room.Backleave {
							jangan := true
							tm, ok := botleave.Get(user)
							if ok {
								if time.Now().Sub(tm.(time.Time)) < 5*time.Second {
									jangan = false
								}
							}
							if jangan {
								if filterWar.ceki(user) {
									if !MemBan(Group, user) && !InArray2(Squadlist, user) && !UserBot.GetBot(user) && !InArray2(Room.GoMid, user) {
										hstg(Group, user)
										Room.Leave = time.Now()
									}
								}
							}
						} else {
							if Room.Leavebool {
								if _, ok := cleave.Get(user); !ok {
									cleave.Set(user, 1)
									if !MemBan(Group, user) && !InArray2(Squadlist, user) && !UserBot.GetBot(user) && !InArray2(Room.GoMid, user) {
										Room.LeaveSet(client, Group, user)
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 30 {
					Group := op.Param1
					Room := oop.GetRoom(Group)
					if Room.Announce && oop.IoGOBot(Group, client) {
						Optime := op.CreatedTime
						if cekOp(Optime) {
							Room.CheckAnnounce(client, Group)
						}
					}
				} else if op.Type == 123 {
					client.CInvite()
				} else if op.Type == 132 {
					client.CountKick()
				} else if op.Type == 125 {
					client.CCancel()
				}
			}
		}(multiFunc)
		for _, ops := range multiFunc {
			if ops.Type == 0 {
				client.CorrectRevision(ops, false, true, true)
			} else {
				client.CorrectRevision(ops, true, false, false)
			}
		}
	}
}

////NEW
func kickPelaku(client *oop.Account, to, pelaku string) {
	defer panicHandle("kickPelaku")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*oop.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
			oke = append(oke, mid)
		}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		Room.HaveClient = exe
		if _, ok := memlist[pelaku]; ok {
			exe[0].DeleteOtherFromChats(to, pelaku)
		}
	}
	oop.SetAva(to, oke)
}
func addwl(g string, w []string) {
	for _, mid := range w {
		if !MemBan(g, mid) {
			if !InArray2(Whitelist, mid) && MemUser(g, mid) {
				Whitelist = append(Whitelist, mid)
			}
		}
	}
}
func CancelProtect(client *oop.Account, mem []string, to string) {
	defer panicHandle("cancelall")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, target)
		}(target, no)
		if ah >= MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}
func cancelall(client *oop.Account, mem []string, to string) {
	defer panicHandle("cancelall")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, target)
		}(target, no)
		if ah >= MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}

func getfuck(cl *oop.Account, vo string, Group string) {
	defer panicHandle("getfuck")
	runtime.GOMAXPROCS(cpu)
	if MemBan(Group, vo) {
		cl.CancelChatInvitations(Group, vo)
	}
}
func cancelallcek(client *oop.Account, mem []string, to string) {
	defer panicHandle("cancelallcek")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	Cans := []*oop.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				Cans = append(Cans, cl)
			}
			oke = append(oke, mid)
		}
	}
	if len(Cans) != 0 {
		sort.Slice(Cans, func(i, j int) bool {
			return Cans[i].KickPoint < Cans[j].KickPoint
		})
		Room.HaveClient = Cans
		no := 0
		ah := 0
		for _, target := range mem {
			go getfuck(Cans[no], target, to)
			if ah >= MaxCancel {
				no++
				if no >= len(Cans) {
					no = 0
				}
				ah = 0
			}
			ah++
		}
	}
}

func groupBackupKick(client *oop.Account, to, pelaku string, cek bool) {
	defer panicHandle("groupBackup")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	memlist, _ := client.GetChatListMap(to)
	ban := []string{}
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
		} else if MemBan(to, mid) {
			ban = append(ban, mid)
		}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		Room.HaveClient = exe
		if Killmode != "none" && AutoBan && cek {
			if Killmode == "kill" {
				kill := exe[0].GetSameJoiningTime(to, pelaku)
				for _, i := range kill {
					if MemUser(to, i) && !InArray2(ban, i) {
						Banned.AddBan(i)
						ban = append(ban, i)
					}
				}
			}
			no := 0
			ah := 0
			for _, target := range ban {
				go func(target string, no int) {
					exe[no].DeleteOtherFromChats(to, target)
				}(target, no)
				if ah >= MaxKick {
					no++
					if no >= len(exe) {
						no = 0
					}
					ah = 0
				}
				ah++
			}
		} else {
			if _, ok := memlist[pelaku]; ok {
				exe[0].DeleteOtherFromChats(to, pelaku)
			}
		}
	}
}
func deljoin(user string) {
	for _, us := range opjoin {
		if us == user {
			opjoin = Remove(opjoin, user)
		}
	}
}
func groupBackupInv(client *oop.Account, to string, optime int64, korban string) {
	runtime.GOMAXPROCS(cpu)
	defer panicHandle("groupBackupInv")
	memlist, _ := client.GetChatListMap(to)
	exe := []*oop.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
			oke = append(oke, mid)
		}
	}
	ClAct := len(exe)
	if ClAct != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		if ModeBackup == "invite" {
			exe[0].InviteIntoGroupNormal(to, []string{korban})
		}
	}
	oop.SetAva(to, oke)
}
func getBot(client *oop.Account, to string) []*oop.Account {
	_, memlist := client.GetGroupMember(to)
	exe := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
		}
	}
	sort.Slice(exe, func(i, j int) bool {
		return exe[i].KickPoint < exe[j].KickPoint
	})
	oop.GetRoom(to).HaveClient = exe
	return exe
}
func AddContact2(cl *oop.Account, con string) int {
	fl, _ := cl.GetAllContactIds()
	if !InArray2(fl, con) {
		if con != cl.MID && !cl.Limitadd {
			_, err := cl.FindAndAddContactsByMidnew(con)
			if err != nil {
				println(fmt.Sprintf("%v", err.Error()))
				return 0
			}
			return 1
		} else {
			return 0
		}
	}
	return 1
}
func kickDirt(client *oop.Account, to, pelaku string) {
	runtime.GOMAXPROCS(cpu)
	cans := oop.Actor(to)
	for _, cl := range cans {
		if oop.GetRoom(to).Act(cl) {
			in := cl.DeleteOtherFromChat(to, pelaku)
			if in == 35 || in == 10 {
				continue
			} else {
				break
			}
		}
	}
}
func CekPurge(optime int64) bool {
	defer oop.PanicOnly()
	for _, tar := range PurgeOP {
		if tar == optime {
			return false
		}
	}
	PurgeOP = append(PurgeOP, optime)
	return true
}
func cekjoin(optime string) bool {
	defer oop.PanicOnly()
	for _, tar := range opjoin {
		if tar == optime {
			return false
		}
	}
	opjoin = append(opjoin, optime)
	return true
}
func cekOp2(optime int64) bool {
	for _, tar := range cekoptime {
		if tar == optime {
			return false
		}
	}
	cekoptime = append(cekoptime, optime)
	return true
}
func getKey(cmd string) string {
	mp := oop.HashToMap(CmdHelper)
	for k, v := range mp {
		if v.(string) == cmd {
			return k
		}
	}
	return cmd
}
func LogFight(room *oop.LineRoom) {
	defer panicHandle("logfight")
	if LogMode {
		var tx = ""
		for i := 0; i < len(ClientBot); i++ {
			exe := ClientBot[i]
			if !exe.Frez {
				g, err := exe.GetGroupMember(room.Id)
				if err != nil {
					continue
				} else {
					room.Name = g
					break
				}
			}
		}

		tx += fmt.Sprintf("Squad action's in Group:\n%s\n", room.Name)
		if room.Kick != 0 {
			tx += fmt.Sprintf("\nKick's: %v", room.Kick)
		}
		if room.Invite != 0 {
			tx += fmt.Sprintf("\nInvite's: %v", room.Invite)
		}
		if room.Cancel != 0 {
			tx += fmt.Sprintf("\nCancel's: %v", room.Cancel)
		}
		if room.Kick == 0 && room.Invite == 0 && room.Cancel == 0 {
			room.Kick = 0
			room.Invite = 0
			room.Cancel = 0
			return
		}
		room := oop.GetRoom(LogGroup)
		if len(room.Client) != 0 {
			exe, err := SelectBot(room.Client[0], LogGroup)
			if err == nil {
				if exe != nil {
					exe.SendMessage(LogGroup, tx)
				}
			} else {
				LogMode = false
				LogGroup = ""
			}
		}
	}
	room.Kick = 0
	room.Invite = 0
	room.Cancel = 0
}

func autoset() {
	defer panicHandle("autoset")
	now := time.Now()
	for _, room := range oop.SquadRoom {
		if !room.Fight.IsZero() {
			if now.Sub(room.Fight) >= 3*time.Second {
				if AutoPro {
					room.AutoBro()
				}
				room.Fight = time.Time{}
				var cll *oop.Account
				if len(room.Client) != 0 {
					cll = room.Client[0]
					name, mem, pending := cll.GetChatList(room.Id)
					room.Name = name
					room.Reset()
					sort.Slice(room.Ava, func(i, j int) bool {
						return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
					})
					sort.Slice(room.HaveClient, func(i, j int) bool {
						return room.HaveClient[i].KickPoint < room.HaveClient[j].KickPoint
					})
					exe := []*oop.Account{}
					for _, cls := range room.Client {
						if InArray2(mem, cls.MID) && !cls.Frez && !cls.Limited {
							exe = append(exe, cls)
						}
					}
					room.HaveClient = exe
					if len(exe) != 0 {
						Backup := []string{}
						li, ok := backlist.Get(room.Id)
						if ok {
							mems := li.([]string)
							for _, l := range mems {
								if !InArray2(mem, l) && !InArray2(Backup, l) && !InArray2(pending, l) {
									Backup = append(Backup, l)
								}
							}
						}
						backlist.Set(room.Id, []string{})
						if len(Backup) != 0 {
							celek := len(Backup)
							no := 0
							bat := 5
							ClAct := len(exe)
							if ClAct != 0 {
								if celek < bat {
									for _, cl := range exe {
										cl.GetRecommendationIds()
										for _, mid := range Backup {
											oop.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range Backup {
											if InArray2(fl, mid) {
												bb = append(bb, mid)
												Backup = Remove(Backup, mid)
											}
										}
										if len(bb) != 0 {
											cl.InviteIntoGroupNormal(room.Id, bb)
										}
										if len(Backup) == 0 {
											break
										}
									}
								} else {
									hajar := []string{}
									z := celek / bat
									y := z + 1
									for i := 0; i < y; i++ {
										if no >= ClAct {
											no = 0
										}
										client := exe[no]
										if i == z {
											hajar = Backup[i*bat:]
										} else {
											hajar = Backup[i*bat : (i+1)*bat]
										}
										if len(hajar) != 0 {
											client.GetRecommendationIds()
											for _, mid := range hajar {
												oop.AddContact3(client, mid)
											}
											fl, _ := client.GetAllContactIds()
											bb := []string{}
											for _, mid := range hajar {
												if InArray2(fl, mid) {
													bb = append(bb, mid)
												}
											}
											if len(bb) != 0 {
												client.InviteIntoGroupNormal(room.Id, bb)
											}
										}
										no += 1
									}
								}
							}
						}
					}
				} else {
					oop.SquadRoom = oop.RemoveRoom(oop.SquadRoom, room)
				}
				filterWar.clear()
				Nkick = &hashmap.HashMap{}
				filterop = &hashmap.HashMap{}
				oplist = []int64{}
				Ceknuke = &hashmap.HashMap{}
				cekoptime = []int64{}
				PurgeOP = []int64{}
				filtermsg = &hashmap.HashMap{}
				opjoin = []string{}
				room.ListInvited = []string{}
				Cekpurge = []int64{}
				AutoproN = false
				cekGo = []int64{}
			}
		}
		if !room.Leave.IsZero() {
			if now.Sub(room.Leave) >= 3*time.Second {
				room.Leave = time.Time{}
				if len(room.LeaveBack) != 0 {
					var cll *oop.Account
					if len(room.Client) != 0 {
						cll = room.Client[0]
						botleave = &hashmap.HashMap{}
						name, mem, invs := cll.GetChatList(room.Id)
						room.Name = name
						exe := []*oop.Account{}
						for _, cls := range room.Client {
							if InArray2(mem, cls.MID) && !InArray2(room.GoMid, cls.MID) {
								exe = append(exe, cls)
							}
						}
						inv := []string{}
						asu := room.LeaveBack
						room.LeaveBack = []string{}
						if len(exe) != 0 {
							for _, l := range asu {
								if !MemBan(room.Id, l) && !InArray2(inv, l) && !InArray2(mem, l) && !InArray2(invs, l) {
									inv = append(inv, l)
								}
							}
							if len(inv) != 0 {
								cls := exe
								for _, cl := range cls {
									if !cl.Limited {
										cl.GetRecommendationIds()
										for _, mid := range inv {
											oop.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range inv {
											if InArray2(fl, mid) {
												bb = append(bb, mid)
											}
										}
										cl.InviteIntoGroupNormal(room.Id, bb)
										for _, mid := range bb {
											if MemUser(room.Id, mid) {
												cl.UnFriend(mid)
											}
										}
										break
									}
								}
							}
						}
					} else {
						oop.SquadRoom = oop.RemoveRoom(oop.SquadRoom, room)
					}
				}
			}
		}
	}
	for _, cl := range oop.Waitadd {
		v, ok := oop.BlockAdd.Get(cl.MID)
		if !ok {
			if now.Sub(cl.TimeBan) >= 1*time.Hour {
				cl.Limitadd = false
				cl.Add = 0
				cl.Lastadd = now
				oop.Waitadd = oop.RemoveCl(oop.Waitadd, cl)
				oop.BlockAdd.Del(cl.MID)
			}
		} else {
			if now.Sub(v.(time.Time)) >= 24*time.Hour {
				oop.BlockAdd.Del(cl.MID)
				cl.Limitadd = false
				cl.Add = 0
				cl.Lastadd = now
				oop.Waitadd = oop.RemoveCl(oop.Waitadd, cl)
				oop.BlockAdd.Del(cl.MID)
			}
		}
	}
	for _, cl := range ClientBot {
		if now.Sub(cl.Lastadd) >= 1*time.Hour {
			cl.Add = 0
			cl.Lastadd = now
		}
		if now.Sub(cl.Lastkick) >= 1*time.Hour {
			cl.TempKick = 0
			cl.TempInv = 0
		}
	}
	for _, cl := range oop.KickBans {
		v, ok := oop.GetBlock.Get(cl.MID)
		if !ok {
			if now.Sub(cl.TimeBan) >= 1*time.Hour {
				oop.KickBans = oop.RemoveCl(oop.KickBans, cl)
				cl.Limited = false
				cl.TempKick = 0
				cl.TempInv = 0
				cl.Frez = false
				oop.GetBlock.Del(cl.MID)
			}
		} else {
			if now.Sub(v.(time.Time)) >= 24*time.Hour {
				oop.GetBlock.Del(cl.MID)
				oop.KickBans = oop.RemoveCl(oop.KickBans, cl)
				cl.Limited = false
				cl.Frez = false
				cl.TempKick = 0
				cl.TempInv = 0
				cl.KickCount = 0
				cl.KickPoint = 0
				cl.InvCount = 0
				cl.CountDay = 0
			}
		}
	}
	for m, v := range oop.HashToMap(oop.GetBlockAdd) {
		cl := GetKorban(m)
		if cl.Limited {
			if now.Sub(v.(time.Time)) >= 1*time.Hour {
				cl.Limitadd = false
				oop.GetBlockAdd.Del(cl.MID)
			}
		}
	}
	if now.Sub(aclear) >= 30*time.Second {
		filterop = &hashmap.HashMap{}
		Nkick = &hashmap.HashMap{}
		filterWar.clear()
		oplist = []int64{}
		timeSend = []int64{}
		Ceknuke = &hashmap.HashMap{}
		cekoptime = []int64{}
		filtermsg = &hashmap.HashMap{}
		aclear = now
		PurgeOP = []int64{}
		Cekpurge = []int64{}
		opjoin = []string{}
		cekGo = []int64{}
		AutoproN = false
	}
	if now.Sub(TimeSave) >= 3*time.Hour {
		SaveBackup()
		TimeBackup = now
	}
	if !TimeBackup.IsZero() {
		BackSeave()
	}
}
func back(to, u string) {
	li, ok := backlist.Get(to)
	if ok {
		list := li.([]string)
		if !InArray2(list, u) {
			list = append(list, u)
		}
		backlist.Set(to, list)
	} else {
		list := []string{u}
		backlist.Set(to, list)
	}
}
func Purgesip(Group string, cl *oop.Account) {
	defer panicHandle("purgesip")
	mem := make(chan []string)
	go func(m chan []string) {
		memlistss := []string{}
		_, memlists := cl.GetGroupMember(Group)
		for target, _ := range memlists {
			if MemBan(Group, target) {
				memlistss = append(memlistss, target)
			}
		}
		m <- memlistss
	}(mem)
	Cans := oop.Actor(Group)
	ClAct := len(Cans)
	hajar := []string{}
	var client *oop.Account
	memlist := <-mem
	celek := len(memlist)
	if celek > MaxKick {
		if ClAct != 0 {
			z := celek / MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= ClAct {
					no = 0
				}
				if i != 0 {
					client = Cans[no]
				} else {
					client = cl
				}
				if i == z {
					hajar = memlist[i*MaxKick:]
				} else {
					hajar = memlist[i*MaxKick : (i+1)*MaxKick]
				}
				if len(hajar) != 0 {
					for _, target := range hajar {
						go client.DeleteOtherFromChats(Group, target)
					}
				}
				no += 1
			}
		} else if !cl.Limited {
			for _, target := range memlist {
				go cl.DeleteOtherFromChats(Group, target)
			}
		}
	} else if !cl.Limited {
		for _, target := range memlist {
			go cl.DeleteOtherFromChats(Group, target)
		}
	}
}
func InArrayChat(arr []*talkservice.Chat, str *talkservice.Chat) bool {
	for _, tar := range arr {
		if tar.ChatMid == str.ChatMid {
			return true
		}
	}
	return false
}
func InfoGroup(client *oop.Account, gid string) string {
	list := ""
	GetSquad(client, gid)
	Room := oop.GetRoom(gid)
	_, mem, pending := client.GetChatList(gid)
	creator := []string{}
	buyer := []string{}
	owner := []string{}
	master := []string{}
	admin := []string{}
	gowner := []string{}
	gadmin := []string{}
	squad := []string{}
	bot := []string{}
	ban := []string{}
	fuck := []string{}
	mute := []string{}
	Gban := []string{}
	Glist := []string{}
	Maker := []string{}
	Seller := []string{}
	mGlist := []string{}
	for _, from := range mem {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(Glist, from) {
				Glist = append(Glist, from)
			}
		} else if UserBot.GetCreator(from) {
			creator = append(creator, from)
		} else if UserBot.GetSeller(from) {
			Seller = append(Seller, from)
		} else if InArray2(MAKERS, from) {
			Maker = append(Maker, from)
		} else if UserBot.GetBuyer(from) {
			buyer = append(buyer, from)
		} else if UserBot.GetOwner(from) {
			owner = append(owner, from)
		} else if UserBot.GetMaster(from) {
			master = append(master, from)
		} else if UserBot.GetAdmin(from) {
			admin = append(admin, from)
		} else if InArray2(Room.Gowner, from) {
			gowner = append(gowner, from)
		} else if InArray2(Room.Gadmin, from) {
			gadmin = append(gadmin, from)
		} else if UserBot.GetBot(from) {
			bot = append(bot, from)
		} else if Banned.GetFuck(from) {
			fuck = append(fuck, from)
		} else if Banned.GetBan(from) {
			ban = append(ban, from)
		} else if Banned.GetMute(from) {
			mute = append(mute, from)
		} else if InArray2(Room.Gban, from) {
			Gban = append(Gban, from)
		} else if InArray2(Squadlist, from) {
			squad = append(squad, from)
		}
	}
	for _, from := range pending {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(mGlist, from) {
				mGlist = append(mGlist, from)
			}
		} else if UserBot.GetCreator(from) {
			creator = append(creator, from)
		} else if UserBot.GetSeller(from) {
			Seller = append(Seller, from)
		} else if InArray2(MAKERS, from) {
			Maker = append(Maker, from)
		} else if UserBot.GetBuyer(from) {
			buyer = append(buyer, from)
		} else if UserBot.GetOwner(from) {
			owner = append(owner, from)
		} else if UserBot.GetMaster(from) {
			master = append(master, from)
		} else if UserBot.GetAdmin(from) {
			admin = append(admin, from)
		} else if InArray2(Room.Gowner, from) {
			gowner = append(gowner, from)
		} else if InArray2(Room.Gadmin, from) {
			gadmin = append(gadmin, from)
		} else if UserBot.GetBot(from) {
			bot = append(bot, from)
		} else if Banned.GetFuck(from) {
			fuck = append(fuck, from)
		} else if Banned.GetBan(from) {
			ban = append(ban, from)
		} else if Banned.GetMute(from) {
			mute = append(mute, from)
		} else if InArray2(Room.Gban, from) {
			Gban = append(Gban, from)
		} else if InArray2(Squadlist, from) {
			squad = append(squad, from)
		}
	}
	list += fmt.Sprintf("Group Info: %s", Room.Name)
	if len(Glist) != 0 {
		list += "\n\nMember: \n"
		cuh, _ := client.GetContacts(Glist)
		for _, prs := range cuh {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}
	if len(mGlist) != 0 {
		chp, _ := client.GetContacts(mGlist)
		list += "\n\n Pending: \n"
		for _, prs := range chp {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}
	if len(Glist)+len(mGlist) != len(pending)+len(mem) {
		list += "\n\nUsers have access:\n"
		if len(creator) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘁𝗲𝗮𝗺:\n"
			for n, xx := range creator {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Seller) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗦𝗲𝗹𝗹𝗲𝗿:\n"
			for n, xx := range Seller {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Maker) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘁𝗲𝗮𝗺:\n"
			for n, xx := range Maker {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(buyer) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗯𝘂𝘆𝗲𝗿𝘀:\n"
			for n, xx := range buyer {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(owner) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗼𝘄𝗻𝗲𝗿𝘀:\n"
			for n, xx := range owner {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(master) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗺𝗮𝘀𝘁𝗲𝗿𝘀:\n"
			for n, xx := range master {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(admin) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗮𝗱𝗺𝗶𝗻𝘀:\n"
			for n, xx := range admin {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(gowner) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗼𝘄𝗻𝗲𝗿𝘀:\n"
			for n, xx := range gowner {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(gadmin) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗮𝗱𝗺𝗶𝗻𝘀:\n"
			for n, xx := range gadmin {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(bot) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗯𝗼𝘁𝗹𝗶𝘀𝘁\n"
			for n, xx := range bot {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(squad) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝘀𝗾𝘂𝗮𝗱:\n"
			for n, xx := range squad {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(ban) != 0 {
			list += "𝗘𝘅𝗶𝘀𝘁.𝗶𝗻 𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range ban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(fuck) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗳𝘂𝗰𝗸𝗹𝗶𝘀𝘁:\n"
			for n, xx := range fuck {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Gban) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n\n"
			for n, xx := range Gban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(mute) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁:\n\n"
			for n, xx := range mute {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
	}
	return list
}
func nukeAll(Client *oop.Account, Group string) {
	defer oop.PanicOnly()
	memlist := []string{}
	_, memlists := Client.GetGroupMember(Group)
	act := []*oop.Account{}
	for mid, _ := range memlists {
		if MemUser(Group, mid) {
			memlist = append(memlist, mid)
		} else if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				act = append(act, cl)
			}
		}
	}
	lact := len(act)
	if lact == 0 {
		return
	} else {
		sort.Slice(act, func(i, j int) bool {
			return act[i].KickPoint < act[j].KickPoint
		})
		celek := len(memlist)
		if celek < MaxKick || lact == 1 {
			cl := act[0]
			for _, mem := range memlist {
				go cl.DeleteOtherFromChat(Group, mem)
			}
		} else {
			hajar := []string{}
			z := celek / MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= lact {
					no = 0
				}
				go func(Group string, no int, i int, z int, memlist []string, act []*oop.Account) {
					Client = act[no]
					if i == z {
						hajar = memlist[i*MaxKick:]
					} else {
						hajar = memlist[i*MaxKick : (i+1)*MaxKick]
					}
					if len(hajar) != 0 {
						for _, target := range hajar {
							go Client.DeleteOtherFromChat(Group, target)
						}
					}
				}(Group, no, i, z, memlist, act)
				no += 1
			}
		}
		oop.GetRoom(Group).HaveClient = act
	}
}
func AcceptJoin(client *oop.Account, Group string) {
	defer panicHandle("AcceptJoin")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(Group)
	if AutoPro {
		Room.AutoBro()
	}
	_, memlist := client.GetGroupMember(Group)
	oke := []string{}
	ban := []string{}
	exe := []*oop.Account{}
	Botss := []*oop.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			oke = append(oke, mid)
			cl := GetKorban(mid)
			Botss = append(Botss, cl)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemBan(Group, mid) {
			ban = append(ban, mid)
		}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		Room.HaveClient = exe
		Room.Client = Botss
		Room.Bot = oke
		oop.SetAva(Group, oke)
		if canceljoin {
			Canceljoin(client, Group)
		} else if NukeJoin {
			nukeAll(client, Group)
		} else if lagjoin {
			Lagejoin(client, Group)
		} else {
			if AutoPurge {
				if len(ban) != 0 {
					no := 0
					ah := 0
					for _, target := range ban {
						go func(target string, no int) {
							exe[no].DeleteOtherFromChats(Group, target)
						}(target, no)
						if ah >= MaxKick {
							no++
							if no >= len(exe) {
								no = 0
							}
							ah = 0
						}
						ah++
					}
				}
			}
		}
		if Autojoin == "qr" {
			AutojoinQr(exe[0], Group)
		} else {
			if Autojoin == "invite" {
				Setinviteto(exe[0], Group, exe[0].Squads)
			}
		}
	}
}

func Bot(op *talkservice.Operation, client *oop.Account, ch chan int) {
	defer panicHandle("Bot")
	msg := op.Message
	if msg.ToType != 2 {
		return
	}
	if _, ok := Commandss.Get(op.CreatedTime); ok {
		return
	} else {
		Commandss.Set(op.CreatedTime, client)
	}
	if time.Now().Sub(timeabort) >= 60*time.Second {
		abort()
	}
	Rname := MsRname
	Sname := MsSname
	sender := op.Message.From_
	text := op.Message.Text
	receiver := op.Message.To
	var pesan = strings.ToLower(text)
	var to string
	mentions := mentions{}
	if op.Message.ToType == 0 {
		to = sender
	} else {
		to = receiver
	}
	if len(Sinderremote) != 0 {
		if InArray2(Sinderremote, sender) {
			if remotegrupid != "" {
				remotegrupidto = to
				to = remotegrupid
			}
		}
	}
	mentionlist := []string{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["MENTION"]), &mentions)
	for _, mention := range mentions.MENTIONEES {
		if !InArray2(mentionlist, mention.Mid) {
			mentionlist = append(mentionlist, mention.Mid)
		}
	}
	var Rplay = ""
	var room *oop.LineRoom
	var bks = []*oop.Account{}
	room = oop.GetRoom(to)
	bks = room.Client
	if len(bks) == 0 {
		GetSquad(client, to)
		room = oop.GetRoom(to)
		bks = room.Client
	}
	sort.Slice(room.Ava, func(i, j int) bool {
		return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
	})
	bk := []*oop.Account{}
	bk2 := []*oop.Account{}
	for _, n := range bks {
		bk = append(bk, n)
		if !n.Limited {
			bk2 = append(bk2, n)
		}
	}
	clen := len(bk2)
	if clen != 0 {
		client = bk2[0]
		room.Exe = bk2[0]
		room.Limit = false
	} else {
		room.Limit = true
	}
	if room.AntiTag && MemUser(to, msg.From_) && len(mentionlist) != 0 && !room.Automute {
		if room.Limit {
			client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot Can't kick in AntiTag")
			return
		}
		if client.Limited == false {
			client.DeleteOtherFromChats(to, msg.From_)
		} else {
			for _, bot := range bk {
				if bot.Limited == false {
					bot.DeleteOtherFromChats(to, msg.From_)
					break
				}
			}
		}
	}
	if op.Message.RelatedMessageId != "" && len(mentionlist) == 0 {
		asu, _ := client.GetRecentMessagesV2(op.Message.To)
		for _, xx := range asu {
			if xx.ID == op.Message.RelatedMessageId {
				Rplay = xx.From_
				break
			}
		}
	}
	if ChangPict && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, "Send your image.")
				StartChangeImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangCover && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, "Send your image.")
				StartChangeImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangVpict && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, "Send your video.")
				StartChangevImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangVcover && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, "Send your video.")
				StartChangevImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangName {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				if MsgName != "" {
					for i := range changepic {
						if TimeDown(i) {
							star := MsgName
							changepic[i].UpdateProfileName(star)
							changepic[i].SendMessage(to, "Profile name updated.")
						}
					}
				} else {
					client.SendMessage(to, "Add name first.")
				}
				ChangName = false
				MsgName = ""
			}
			timeabort = time.Now()
		}
	} else if ChangeBio {
		if len(mentionlist) != 0 {
			changepic = []*oop.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !oop.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				if MsgBio != "" {
					for i := range changepic {
						if TimeDown(i) {
							star := MsgBio
							changepic[i].UpdateProfileBio(star)
							changepic[i].SendMessage(to, "Profile status updated.")
						}
					}
				} else {
					client.SendMessage(to, "Add Bio first.")
				}
				ChangeBio = false
				MsgBio = ""
			}
			timeabort = time.Now()
		}
	}
	if op.Message.ContentType == 1 {
		if StartChangeImg && len(changepic) != 0 {
			if !MemUser(to, sender) {
				if ChangPict {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for n, p := range changepic {
							if TimeDown(n) {
								go func(p *oop.Account) {
									if StartChangevImg2 {
										err := p.UpdatePictureProfile(path, "v")
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, "Update dual profile failure.")
										} else {
											p.SendMessage(to, "Update video picture done.")
										}
									} else {
										err := p.UpdatePictureProfile(path, "p")
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, "Update picture profile failure.")
										} else {
											p.SendMessage(to, "Update Image picture done.")
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.Remove(path)
					} else {
						fmt.Println(err)
						if StartChangevImg2 {
							client.SendMessage(to, "Download video picture Failure.")
						} else {
							client.SendMessage(to, "Download Image picture Failure.")
						}
					}
					StartChangevImg2 = false
					StartChangeImg = false
					ChangPict = false
				} else if ChangCover {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for n, p := range changepic {
							if TimeDown(n) {
								go func(p *oop.Account) {
									if StartChangevImg2 {
										err := p.UpdateCoverWithVideo(path)
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, "Update video cover failure.")
										} else {
											p.SendMessage(to, "Update video cover done.")
											time.Sleep(2 * time.Second)
										}
									} else {
										err := p.UpdateCover(path)
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, "Update picture cover failure.")
										} else {
											p.SendMessage(to, "Update Image cover done.")
											time.Sleep(2 * time.Second)
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.Remove(path)
					} else {
						fmt.Println(err)
						if StartChangevImg2 {
							client.SendMessage(to, "Download video cover Failure.")
						} else {
							client.SendMessage(to, "Download Image cover Failure.")
						}
					}
					StartChangevImg2 = false
					StartChangeImg = false
					ChangCover = false
				}
				timeabort = time.Now()
			}
		}
	} else if op.Message.ContentType == 2 {
		if StartChangevImg && len(changepic) != 0 {
			if !MemUser(to, sender) {
				if ChangVpict {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for _, p := range changepic {
							go func(p *oop.Account) {
								err := p.UpdateVideoProfile(path)
								if err != nil {
									fmt.Println(err)
									p.SendMessage(to, "Update video profile failure.")
								}
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, "Upload video done, now send your image.")
						os.Remove(path)
						StartChangevImg2 = true
						ChangPict = true
						StartChangeImg = true
						ChangVpict = false
						StartChangevImg = false
					} else {
						fmt.Println(err)
						client.SendMessage(to, "Download Image Failure.")
					}
				} else if ChangVcover {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for _, p := range changepic {
							go func(p *oop.Account) {
								p.UpdateCoverVideo(path)
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, "Upload video done, now send your image.")
						os.Remove(path)
						StartChangevImg2 = true
						StartChangeImg = true
						ChangCover = true
						StartChangevImg = false
						ChangVcover = false
					} else {
						fmt.Println(err)
						client.SendMessage(to, "Download Image Failure.")
					}
				}
				timeabort = time.Now()
			}
		}
	} else if msg.ContentType == 0 && msg.Text != "" {
		if room.Automute && MemUser(to, msg.From_) {
			if client.Limited == false {
				client.DeleteOtherFromChats(to, msg.From_)
			} else {
				for _, bot := range bk {
					if bot.Limited == false {
						bot.DeleteOtherFromChats(to, msg.From_)
						break
					}
				}
			}
		} else {
			if MemBan2(to, msg.From_) && MemUser(to, msg.From_) {
				if client.Limited == false {
					client.DeleteOtherFromChats(to, msg.From_)
				} else {
					for _, bot := range bk {
						if bot.Limited == false {
							bot.DeleteOtherFromChats(to, msg.From_)
							break
						}
					}
				}
			}
		}
		cmds := gettxt(sender, client, pesan, Rname, Sname, client.MID, mentionlist, to)
		text := op.Message.Text
		newsend := ""
		var pesan = strings.ToLower(text)
		for _, cmd := range strings.Split(cmds, ",") {
			if strings.HasPrefix(cmd, "creator") && cmd != "creators" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 13
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "creators" {
				rngcmd := GetComs(1, "creators")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Creator) != 0 {
							list := " ✠ 𝗖𝗿𝗲𝗮𝘁𝗼𝗿𝘀 ✠ \n"
							for num, xd := range UserBot.Creator {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if cmd == "clearcreator" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Creator) != 0 {
							newsend += fmt.Sprintf("Cleared %v Creatorlist\n", len(UserBot.Creator))
							UserBot.ClearCreator()
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "uncreator") {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "uncreator"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Creator)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "seller") && cmd != "sellers" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "sellers" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Seller) != 0 {
							list := " ✠ 𝗦𝗲𝗹𝗹𝗲𝗿𝘀 ✠ \n"
							for num, xd := range UserBot.Seller {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if cmd == "clearseller" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Seller) != 0 {
							newsend += fmt.Sprintf("Cleared %v sellerlist\n", len(UserBot.Seller))
							UserBot.ClearSeller()
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unseller") {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unseller"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Seller)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "listcmd" {
				rngcmd := GetComs(4, "listcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						list := Cmdlistcheck()
						client.SendMessage(to, list)
					}
				}
			} else if strings.HasPrefix(cmd, "expel") {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "access" || cmd == Commands.Access && Commands.Access != "" {
				rngcmd := GetComs(4, "access")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "✠ 𝗔𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁 ✠"
						if len(UserBot.Owner) != 0 {
							listadm += "\n\n 👑 𝗼𝘄𝗻𝗲𝗿𝘀 👑"
							for num, xd := range UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Master) != 0 {
							listadm += "\n\n 🎩 𝗺𝗮𝘀𝘁𝗲𝗿𝘀 🎩"
							for num, xd := range UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Admin) != 0 {
							listadm += "\n\n 🎓 𝗮𝗱𝗺𝗶𝗻𝘀 🎓"
							for num, xd := range UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "𝗔ccess is empty.\n"
						}
					}
				}
			} else if cmd == "allbanlist" || cmd == Commands.Allbanlist && Commands.Allbanlist != "" {
				rngcmd := GetComs(4, "allbanlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listadm := AllBanList(client)
						if listadm != "✠ 𝗔𝗹𝗹 𝗯𝗮𝗻𝗹𝗶𝘀𝘁𝘀 ✠" {
							newsend += listadm + "\n"
						} else {
							newsend += "𝗔ccess is empty.\n"
						}
					}
				}
			} else if cmd == "gaccess" || cmd == Commands.Gaccess && Commands.Gaccess != "" {
				rngcmd := GetComs(7, "gaccess")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "✠ 𝗚𝗮𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁 ✠"
						if len(room.Gowner) != 0 {
							listadm += "\n\n👑 𝗴𝗼𝘄𝗻𝗲𝗿𝘀 👑"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(room.Gadmin) != 0 {
							listadm += "\n\n 🎓 𝗴𝗮𝗱𝗺𝗶𝗻𝘀 🎓"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "Gaccess is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "buyer") && cmd != "buyers" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setdate ") {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						ha := strings.Split((cmd), "setdate ")
						haj := ha[1]
						haj = StripOut(haj)
						has := strings.Split(haj, "-")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(has[2])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							times := time2.Format(time.RFC3339)
							Data.Dalltime = times
							str := fmt.Sprintf("⚙️ Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\n⚙️ Remaining:\n  %v", botDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "addweek" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 7 * mont
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addday" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addmonth" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 30 * mont
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "reboot" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						client.SendMessage(to, "Waiting Rebooting...")
						ReloginProgram()
					}
				}
			} else if strings.HasPrefix(cmd, "unbuyer") {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unbuyer"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Buyer)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "checkram" || cmd == Commands.Checkram && Commands.Checkram != "" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						v, _ := mem.VirtualMemory()
						r := fmt.Sprintf("  ↳Cpu : %v core\n  ↳Ram : %v mb\n  ↳Free : %v mb\n  ↳Cache : %v mb\n  ↳UsedPercent : %f %%", cpu, bToMb(v.Used+v.Free+v.Buffers+v.Cached), bToMb(v.Free), bToMb(v.Buffers+v.Cached), v.UsedPercent)
						newsend += r + "\n"
					}
				}
			} else if cmd == "clearbuyer" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Buyer) != 0 {
							newsend += fmt.Sprintf("Cleared %v buyerlist\n", len(UserBot.Buyer))
							UserBot.ClearBuyer()
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "upimage" || cmd == Commands.Upimage && Commands.Upimage != "" {
				rngcmd := GetComs(3, "upimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangPict = true
						newsend += "Which bot's you want to update Pict.\n"
					}
				}
			} else if cmd == "upcover" || cmd == Commands.Upcover && Commands.Upcover != "" {
				rngcmd := GetComs(3, "upcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangCover = true
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if cmd == "upvimage" || cmd == Commands.Upvimage && Commands.Upvimage != "" {
				rngcmd := GetComs(3, "upvimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangVpict = true
						newsend += "Which bot's you want to update Pict ?.\n"
					}
				}
			} else if cmd == "upvcover" || cmd == Commands.Upvcover && Commands.Upvcover != "" {
				rngcmd := GetComs(3, "upvcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangVcover = true
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "unsend ") {
				rngcmd := GetComs(4, "unsend")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									Nganu, _ := client.GetRecentMessagesV2(op.Message.To)
									Mid := []string{}
									unsed := []string{}
									for _, chat := range Nganu {
										if InArray2(Squadlist, chat.From_) {
											Mid = append(Mid, chat.ID)
										}
									}
									for i := 0; i < len(Mid); i++ {
										if i < result2 {
											unsed = append(unsed, Mid[i])
										}
									}
									if len(unsed) != 0 {
										exess, _ := SelectallBot(client, to)
										if exess != nil {
											for i := range exess {
												Nganu2, _ := exess[i].GetRecentMessagesV2(op.Message.To)
												for _, chat := range Nganu2 {
													if chat.From_ == exess[i].MID {
														if InArray2(unsed, chat.ID) {
															exess[i].UnsendChatnume(to, chat.ID)
														}
													}
												}
											}
										}
									}
								} else {
									client.SendMessage(to, "out of range.")
								}
							}
						} else {
							client.SendMessage(to, "Msg not fund number")
						}
					}
				}
			} else if cmd == "purgeall" || cmd == Commands.Purgeall && Commands.Purgeall != "" {
				rngcmd := GetComs(4, "purgeall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						gr, _ := client.GetGroupIdsJoined()
						all := []string{}
						for _, aa := range gr {
							_, memlist, _ := client.GetChatList(aa)
							lkicks := []string{}
							for _, v := range memlist {
								if MemUser(aa, v) {
									lkicks = append(lkicks, v)
								}
							}
							lkick := []string{}
							for _, ban := range lkicks {
								if MemBan(aa, ban) {
									lkick = append(lkick, ban)
									all = append(all, ban)
								}
							}
							nom := []*oop.Account{}
							ilen := len(lkick)
							xx := 0
							exe := []*oop.Account{}
							for _, c := range oop.GetRoom(aa).Client {
								if !c.Limited {
									exe = append(exe, c)
								}
							}
							if len(exe) != 0 {
								for i := 0; i < ilen; i++ {
									if xx < len(exe) {
										nom = append(nom, exe[xx])
										xx += 1
									} else {
										xx = 0
										nom = append(nom, exe[xx])
									}
								}
								for i := 0; i < ilen; i++ {
									target := lkick[i]
									cl := nom[i]
									go cl.DeleteOtherFromChats(aa, target)
								}
								time.Sleep(1 * time.Second)
							}
						}
						newsend += fmt.Sprintf("Success purgeall %v user in blacklist", len(all))
						logAccess(client, to, sender, "purgeall", all, msg.ToType)
					}
				}
			} else if strings.HasPrefix(cmd, "gleave") {
				rngcmd := GetComs(4, "gleave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									su := "gleave"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := StripOut(str)
									hapuss := oop.Archimed(st, tempgroup)
									if len(hapuss) == 0 {
										client.SendMessage(to, "Please input the right number\nSee group number with command groups")
										return
									}
									names := []string{}
									for _, gid := range hapuss {
										name, mem := client.GetGroupInvitation(gid)
										names = append(names, name)
										anu := []string{}
										for m, _ := range mem {
											if InArray2(Squadlist, m) {
												anu = append(anu, m)
											}
										}
										if len(anu) != 0 {
											for _, mid := range anu {
												cl := GetKorban(mid)
												cl.AcceptGroupInvitationNormal(gid)
												oop.GetRoom(gid).ConvertGo(cl)
											}
										}
										GetSquad(client, gid)
										room := oop.GetRoom(gid)
										bk = room.Client
										for _, cl := range bk {
											go cl.LeaveGroup(gid)
										}
										if LogGroup == gid {
											LogMode = false
											LogGroup = ""
										}
										oop.SquadRoom = oop.RemoveRoom(oop.SquadRoom, room)
									}
									strs := strings.Join(names, ", ")
									client.SendMessage(to, "Bot's leave from group: \n\n"+strs)
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invme ") {
				rngcmd := GetComs(4, "invme")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, "Please input the right number\nSee group number with command groups")
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										GetSquad(client, gid)
										room := oop.GetRoom(gid)
										bk := room.Client
										name, mem, inv := client.GetChatList(gid)
										if InArray2(mem, msg.From_) {
											client.SendMessage(to, "You was on group "+name)
											return
										} else {
											if InArray2(inv, msg.From_) {
												bk[0].CancelChatInvitations(gid, msg.From_)
											}
											for _, cl := range bk {
												if !cl.Limited && !cl.Limitadd {
													AddContact2(cl, msg.From_)
													fl, _ := cl.GetAllContactIds()
													if InArray2(fl, msg.From_) {
														err := cl.InviteIntoGroupNormal(gid, []string{msg.From_})
														if err != nil {
															code := oop.GetCode(err)
															if code != 35 && code != 10 {
																client.SendMessage(to, "You has invited to group "+name)
																return
															}
														} else {
															client.SendMessage(to, "You has invited to group "+name)
															return
														}
													}
												}
											}
											newsend += "Sorry, all bot has invite banned"
										}
									} else {
										newsend += "out of range."
									}
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "decline ") {
				rngcmd := GetComs(1, "decline")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									num, _ := strconv.Atoi(result[1])
									gr := []string{}
									for i := range ClientBot {
										grs, _ := ClientBot[i].GetGroupsInvited()
										if len(grs) != 0 {
											for _, a := range grs {
												if !InArray2(gr, a) {
													gr = append(gr, a)
												}
											}
										}
									}
									grup, _ := client.GetGroups(gr)
									tempgroup := []string{}
									for _, gi := range grup {
										if !InArray2(tempgroup, gi.ChatMid) {
											tempgroup = append(tempgroup, gi.ChatMid)
										}
									}
									if num > 0 && num <= len(tempgroup) {
										exe := []*oop.Account{}
										gen := tempgroup[num-1]
										names, _, _ := client.GetChatList(tempgroup[num-1])
										for i := range ClientBot {
											if ClientMid[ClientBot[i].MID].Limited == false {
												grs, _ := ClientBot[i].GetGroupsInvited()
												if InArray2(grs, gen) {
													exe = append(exe, ClientBot[i])
												}
											}
										}
										if len(exe) != 0 {
											for i := range exe {
												exe[i].RejectChatInvitation(gen)
											}
											newsend += fmt.Sprintf("Successfully declined invitation for: %v\n", names)
										}
									} else {
										newsend += "out of range pendinglist.\n"
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "accept") && cmd != "acceptall" {
				rngcmd := GetComs(4, "accept")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									su := "accept"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := StripOut(str)
									hapuss := oop.Archimed(st, tempginv)
									if len(hapuss) == 0 {
										newsend += "Please input the right number\nSee group number with command groups"
									} else {
										names := []string{}
										for _, gid := range hapuss {
											name, mem := client.GetGroupInvitation(gid)
											names = append(names, name)
											anu := []string{}
											for m, _ := range mem {
												if InArray2(Squadlist, m) {
													anu = append(anu, m)
												}
											}
											if len(anu) != 0 {
												for _, mid := range anu {
													cl := GetKorban(mid)
													cl.AcceptGroupInvitationNormal(gid)
													oop.GetRoom(gid).ConvertGo(cl)
												}
											}
										}
										str := strings.Join(names, ", ")
										newsend += "Bot's join to group \n\n" + str
									}
								}
							}
						}
					}
				}
			} else if cmd == "abort" {
				rngcmd := GetComs(4, "abort")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if remotegrupidto != "" {
							client.SendMessage(remotegrupidto, "Done Have abort.")
						} else {
							newsend += "Done Have abort." + "\n"
						}
						abort()
					}
				}
			} else if cmd == "declineall" {
				rngcmd := GetComs(1, "declineall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						tempgroup := []string{}
						for i := range ClientBot {
							grs, _ := ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !InArray2(tempgroup, gi.ChatMid) {
										tempgroup = append(tempgroup, gi.ChatMid)
									}
									ClientBot[i].RejectChatInvitation(gi.ChatMid)
								}
								time.Sleep(1 * time.Second)
							}

						}
						if len(tempgroup) != 0 {
							newsend += fmt.Sprintf("Successfully declined invitations: (%v)\n", len(tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if cmd == "acceptall" {
				rngcmd := GetComs(3, "acceptall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						tempgroup := []string{}
						for i := range ClientBot {
							grs, _ := ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !InArray2(tempgroup, gi.ChatMid) {
										tempgroup = append(tempgroup, gi.ChatMid)
									}
									ClientBot[i].AcceptGroupInvitationNormal(gi.ChatMid)
									oop.GetRoom(gi.ChatMid).ConvertGo(ClientBot[i])
									time.Sleep(1 * time.Second)
								}
							}
						}
						if len(tempgroup) != 0 {
							newsend += fmt.Sprintf("Success accept bot %v Group\n", len(tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "upstatus") {
				rngcmd := GetComs(3, "upstatus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upstatus"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						MsgBio = str
						ChangeBio = true
						timeabort = time.Now()
						newsend += fmt.Sprintf("Which bot's should be Status %v", str)
					}
				}
			} else if strings.HasPrefix(cmd, "upname") {
				rngcmd := GetComs(3, "upname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						aa := utf8.RuneCountInString(str)
						if aa != 0 && aa <= 20 {
							MsgName = str
							ChangName = true
							timeabort = time.Now()
							newsend += fmt.Sprintf("Which bot's should be Name %v", str)
						}
					}
				}
			} else if cmd == "buyers" {
				rngcmd := GetComs(3, "buyers")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Buyer) != 0 {
							list := " ✠ 𝗯𝘂𝘆𝗲𝗿𝘀 ✠ \n"
							for num, xd := range UserBot.Buyer {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "history" {
				rngcmd := GetComs(4, "history")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						countK := 0
						countinv := 0
						countcancel := 0
						for i := range ClientBot {
							countK = countK + ClientBot[i].Ckick
							countinv = countinv + ClientBot[i].Cinvite
							countcancel = countcancel + ClientBot[i].Ccancel
						}
						list := fmt.Sprintf("History: \n\n Kick: %v \n Cancel: %v \n Invited: %v", countK, countcancel, countinv)
						client.SendMessage(to, list)
					}
				}
			} else if cmd == "clearhide" {
				rngcmd := GetComs(4, "clearhide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(checkHaid) != 0 {
							logAccess(client, to, sender, "clearhid", checkHaid, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v Hidelist\n", len(checkHaid))
							checkHaid = []string{}
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if cmd == "hidelist" {
				rngcmd := GetComs(4, "hidelist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(checkHaid) != 0 {
							list := " ✠ Hide List ✠ \n"
							for num, xd := range checkHaid {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unhide") || strings.HasPrefix(cmd, "delhide") {
				rngcmd := GetComs(4, "unhide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						list := ""
						listuser := []string{}
						nCount1 := 0
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User removed from hidelist:\n\n"
							test2 := "User not exist in hidelist:\n\n"
							for n, xx := range listuser {
								if InArray2(checkHaid, xx) {
									checkHaid = Remove(checkHaid, xx)
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								new := client.Getcontactuser(xx)
								if new != nil {
									list += rengs + ". Closed Account \n"
								} else {
									x, _ := client.GetContact(xx)
									list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

								}
							}
							if nCount1 != 0 {
								client.SendMessage(to, test1+list)
							} else {
								client.SendMessage(to, test2+list)
							}
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unhide"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, checkHaid)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											for _, i := range hapuss {
												if InArray2(checkHaid, i) {
													checkHaid = Remove(checkHaid, i)
													listuser = append(listuser, i)
												}
											}
											if len(listuser) != 0 {
												list += "User removed from hidelist:\n\n"
												for n, xx := range listuser {
													checkHaid = Remove(checkHaid, xx)
													rengs := strconv.Itoa(n + 1)
													new := client.Getcontactuser(xx)
													if new != nil {
														list += rengs + ". Closed Account \n"
													} else {
														x, _ := client.GetContact(xx)
														list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

													}
												}
												client.SendMessage(to, list)
											}
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "hide") && cmd != "hidelist" {
				rngcmd := GetComs(4, "hide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						nCount1 := 0
						list := ""
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User added to hidelist:\n\n"
							test2 := "User already exist in hidelist:\n\n"
							for n, xx := range listuser {
								if !InArray2(checkHaid, xx) {
									checkHaid = append(checkHaid, xx)
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								new := client.Getcontactuser(xx)
								if new != nil {
									list += rengs + ". Closed Account \n"
								} else {
									x, _ := client.GetContact(xx)
									list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

								}
							}
							if nCount1 != 0 {
								client.SendMessage(to, test1+list)
							} else {
								client.SendMessage(to, test2+list)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "owner") && cmd != "owners" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unowner") {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Owner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "clearowner" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Owner) != 0 {
							logAccess(client, to, sender, "clearowner", UserBot.Owner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v ownerlist\n", len(UserBot.Owner))
							UserBot.ClearOwner()
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if cmd == "logmode on" {
				rngcmd := GetComs(4, "logmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if LogGroup == to {
							LogMode = true
							newsend += "Already enabled.\n"
						} else {
							LogMode = true
							LogGroup = to
							newsend += "Logmode is enabled.\n"
						}
					}
				}
			} else if cmd == "logmode off" {
				rngcmd := GetComs(4, "logmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if LogGroup == to {
							LogMode = false
							LogGroup = ""
							newsend += "Logmode is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setsname ") {
				rngcmd := GetComs(3, "setsname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setsname ", "", 1)
						if result == "," || result == "" {
							MsSname = ","
						} else {
							MsSname = result
						}
						newsend += "Sname set to: " + Sname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setrname ") {
				rngcmd := GetComs(3, "setrname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setrname ", "", 1)
						if result == "," || result == "" {
							MsRname = ","
						} else {
							MsRname = result
						}
						newsend += "Succes update Rname to " + Rname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgrespon") {
				rngcmd := GetComs(3, "msgrespon")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgrespon"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						MsgRespon = str
						newsend += "Message respon set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "upgname") {
				rngcmd := GetComs(3, "upgname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upgname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						client.UpdateChatName(to, str)
						newsend += "group name has been changed to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlogo") {
				rngcmd := GetComs(3, "setlogo")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "setlogo"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						Data.Logobot = str
						newsend += "Menu logo set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgwelcome") {
				rngcmd := GetComs(3, "msgwelcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgwelcome"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.WelcomeMsg = str
						newsend += "Message Welcome set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgleave") {
				rngcmd := GetComs(3, "msgleave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgleave"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLeave = str
						newsend += "Message Leave set to: " + str + "\n"
					}
				}

			} else if strings.HasPrefix(cmd, "msgunban ") {
				rngcmd := GetComs(3, "msgunban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgunban ", "", 1)
						MsgBan = result
						newsend += "Message unban set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglurk") {
				rngcmd := GetComs(3, "msglurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msglurk"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLurk = str
						newsend += "Message sider set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgfresh ") {
				rngcmd := GetComs(3, "msgfresh")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgfresh ", "", 1)
						MsFresh = result
						newsend += "Message fresh set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglimit ") {
				rngcmd := GetComs(3, "msglimit")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msglimit ", "", 1)
						MsLimit = result
						newsend += "Message limit set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setkick ") {
				rngcmd := GetComs(3, "setkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						anjay := strings.Split((cmd), " ")
						num, err := strconv.Atoi(anjay[1])
						if err != nil {
							newsend += "Please use number!\n"
						} else {
							MaxKick = num
							newsend += "Limiter kick set to " + anjay[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcancel ") {
				rngcmd := GetComs(3, "setcancel")
				if GetCodeprem(rngcmd, sender, to) {
					anjay := strings.Split((cmd), " ")
					num, err := strconv.Atoi(anjay[1])
					if err != nil {
						newsend += "Please use number!\n"
					} else {
						MaxCancel = num
						newsend += "Limiter cancel set to " + anjay[1] + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlimiter ") {
				rngcmd := GetComs(3, "setlimiter")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						no, err := strconv.Atoi(result[1])
						if err != nil {
							newsend += "Please use number!\n"
						} else {
							MaxKick = no
							MaxCancel = no
							newsend += "Limiter successs set to " + result[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "fuck") && cmd != "fucklist" {
				rngcmd := GetComs(4, "fuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mute") && cmd != "mutelist" {
				rngcmd := GetComs(4, "mute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 11
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "whois") {
				rngcmd := GetComs(4, "whois")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 12
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "name") {
				rngcmd := GetComs(4, "name")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 16
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mid") {
				rngcmd := GetComs(4, "mid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 14
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmute") {
				rngcmd := GetComs(4, "unmute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unmute"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "owners" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Owner) != 0 {
							list := " 👑 𝗼𝘄𝗻𝗲𝗿𝘀 👑 \n"
							for num, xd := range UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmaster") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unmaster"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Master)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungowner") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "ungowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, room.Gowner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "gojoin" {
				rngcmd := GetComs(4, "gojoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m, _ := range mem {
							if InArray2(Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "master") && cmd != "masters" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gowner") && cmd != "gowners" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcmd ") {
				rngcmd := GetComs(4, "setcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "setcmd ", "")
						texts := strings.Split(txt, " ")
						if len(texts) > 1 {
							new := Upsetcmd(texts[0], texts[1])
							if new != "" {
								newsend += new
							} else {
								newsend += "Cmd not found.\n"
							}
						} else {
							newsend += "Cmd not found.\n"
						}
					}
				}
			} else if cmd == "restartcmd" {
				rngcmd := GetComs(4, "restartcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						Commands.Botname = ""
						Commands.Upallimage = ""
						Commands.Upallcover = ""
						Commands.Unsend = ""
						Commands.Upvallimage = ""
						Commands.Upvallcover = ""
						Commands.Appname = ""
						Commands.Useragent = ""
						Commands.Hostname = ""
						Commands.Friends = ""
						Commands.Adds = ""
						Commands.Limits = ""
						Commands.Addallbots = ""
						Commands.Addallsquads = ""
						Commands.Leave = ""
						Commands.Respon = ""
						Commands.Ping = ""
						Commands.Count = ""
						Commands.Limitout = ""
						Commands.Access = ""
						Commands.Allbanlist = ""
						Commands.Gaccess = ""
						Commands.Checkram = ""
						Commands.Backups = ""
						Commands.Upimage = ""
						Commands.Upcover = ""
						Commands.Upvimage = ""
						Commands.Upvcover = ""
						Commands.Bringall = ""
						Commands.Purgeall = ""
						Commands.Banlist = ""
						Commands.Clearban = ""
						Commands.Stayall = ""
						Commands.Clearchat = ""
						Commands.Here = ""
						Commands.Speed = ""
						Commands.Status = ""
						Commands.Tagall = ""
						Commands.Kick = ""
						Commands.Max = ""
						Commands.None = ""
						Commands.Kickall = ""
						Commands.Cancelall = ""
						newsend += "Done restart all Cmd.\n"
					}
				}
			} else if cmd == "cleargowner" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							logAccess(client, to, sender, "cleargowner", room.Gowner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gownerlist\n", len(room.Gowner))
							room.Gowner = []string{}
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmaster" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Master) != 0 {
							newsend += fmt.Sprintf("Cleared %v masterlist\n", len(UserBot.Master))
							logAccess(client, to, sender, "clearmaster", UserBot.Master, msg.ToType)
							UserBot.ClearMaster()
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if cmd == "clearfuck" {
				rngcmd := GetComs(4, "clearfuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Fucklist) != 0 {
							logAccess(client, to, sender, "clearfuck", Banned.Fucklist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(Banned.Fucklist))
							Banned.Fucklist = []string{}
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmute" {
				rngcmd := GetComs(4, "clearmute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Mutelist) != 0 {
							logAccess(client, to, sender, "clearmute", Banned.Mutelist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(Banned.Mutelist))
							Banned.Mutelist = []string{}
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "clear allprotect" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						oop.ClearProtect()
						newsend += "Cleared allprotected.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "perm ") {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						ditha := strings.ReplaceAll(cmd, "perm ", "")
						cmdLil := strings.Split(ditha, " ")
						Addpermcmd(client, to, cmdLil[0], cmdLil[1])
					}
				}
			} else if cmd == "permlist" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						list := PerCheckList()
						if list != "" {
							newsend += list
						} else {
							newsend += "Not have perm in list.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autojoin") {
				rngcmd := GetComs(4, "autojoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "autojoin"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "invite" {
							if Autojoin != "invite" {
								Autojoin = "invite"
								newsend += "Autojoin Invite enabled.\n"
							} else {
								newsend += "Autojoin Already Invite.\n"
							}
						} else if str == "qr" {
							if Autojoin != "qr" {
								Autojoin = "qr"
								newsend += "Autojoin qr enabled.\n"
							} else {
								newsend += "Autojoin Already qr.\n"
							}
						} else if str == "off" {
							if Autojoin != "off" {
								Autojoin = "off"
								newsend += fmt.Sprintf("Autojoin %s disabled.\n", Autojoin)
							} else {
								newsend += "Autojoin Already disabled.\n"
							}
						}
					}
				}
			} else if cmd == "mutelist" {
				rngcmd := GetComs(4, "mutelist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Mutelist) != 0 {
							list := "Mutelist:"
							client.SendPollMention(to, list, Banned.Mutelist)
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "fucklist" {
				rngcmd := GetComs(4, "fucklist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Fucklist) != 0 {
							list := "Fucklist:"
							client.SendPollMention(to, list, Banned.Fucklist)
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "groupcast") {
				rngcmd := GetComs(4, "groupcast")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "groupcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							gr, _ := client.GetGroupIdsJoined()
							for _, gi := range gr {
								client.SendMessage(gi, str)
							}
							newsend += "Success broadcast to " + strconv.Itoa(len(gr)) + " group\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gourl ") {
				rngcmd := GetComs(4, "gourl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						num, err := strconv.Atoi(result[1])
						if err != nil {
							newsend += "invalid number.\n"
						} else {
							gr := []string{}
							for i := range ClientBot {
								grs, _ := ClientBot[i].GetGroupIdsJoined()
								for _, a := range grs {
									if !InArray2(gr, a) {
										gr = append(gr, a)
									}
								}
							}
							groups, _ := client.GetGroups(gr)
							tempgroup := []string{}
							for _, gi := range groups {
								tempgroup = append(tempgroup, gi.ChatMid)
							}
							if num > 0 && num <= len(tempgroup) {
								gid := tempgroup[num-1]
								tick, err := client.ReissueChatTicket(gid)
								if err == nil {
									var err error
									mes := make(chan bool)
									go func() {
										err = client.UpdateChatQrV2(gid, false)
										if err != nil {
											mes <- false
										} else {
											mes <- true
										}
									}()

									newsend += "https://line.me/R/ti/g/" + tick + "\n"
								}
							} else {
								newsend += "out of range.\n"
							}
						}
					}
				}
			} else if cmd == "groups" {
				rngcmd := GetComs(3, "groups")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						gr := []string{}
						for _, p := range ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsJoined()
								for _, a := range grs {
									if !InArray2(gr, a) {
										gr = append(gr, a)
									}
								}
							}
						}
						nm := []string{}
						grup, _ := client.GetGroups(gr)
						ci := []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups := []*talkservice.Chat{}
						tempgroup = []string{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !InArrayChat(groups, gi) {
										groups = append(groups, gi)
										tempgroup = append(tempgroup, gi.ChatMid)
									}
								}
							}
						}
						for c, a := range groups {
							name, mem := a.ChatName, a.Extra.GroupExtra.MemberMids
							c += 1
							jm := 0
							for mid, _ := range mem {
								if InArray2(Squadlist, mid) {
									jm++
								}
							}
							name = fmt.Sprintf("%v. %s (%v/%v)", c, name, jm, len(mem))
							nm = append(nm, name)
							GetSquad(client, a.ChatMid)
						}
						stf := "All Group List:\n\n"
						str := strings.Join(nm, "\n")
						anu := []string{}
						for _, p := range ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsInvited()
								for _, a := range grs {
									if !InArray2(gr, a) && !InArray2(anu, a) {
										anu = append(anu, a)
									}
								}
							}
						}
						grup, _ = client.GetGroups(anu)
						ci = []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups = []*talkservice.Chat{}
						tempginv = []string{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !InArrayChat(groups, gi) {
										groups = append(groups, gi)
										tempginv = append(tempginv, gi.ChatMid)
									}
								}
							}
						}
						nm = []string{}
						nn := 1
						for _, a := range groups {
							name, mem, inv := a.ChatName, a.Extra.GroupExtra.MemberMids, a.Extra.GroupExtra.InviteeMids
							if name != "" {
								jm := 0
								for mid, _ := range inv {
									if InArray2(Squadlist, mid) {
										jm++
									}
								}
								if jm != 0 {
									name = fmt.Sprintf("%v. %s (invited) (%v/%v)", nn, name, jm, len(mem))
									nm = append(nm, name)
									GetSquad(client, a.ChatMid)
									nn++
								} else {
									tempginv = Remove(tempginv, a.ChatMid)
								}
							} else {
								tempginv = Remove(tempginv, a.ChatMid)
							}
						}
						var strs, strsa = "", ""
						if len(nm) != 0 {
							strs = "\n\nAll Group Invitation:\n\n"
							strsa = strings.Join(nm, "\n")
						}
						newsend += stf + str + strs + strsa
					}
				}
			} else if strings.HasPrefix(cmd, "nukejoin ") {
				rngcmd := GetComs(3, "nukejoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "nukejoin ", "", 1)
						if spl == "on" {
							NukeJoin = true
							newsend += "Nukejoin is enabled.\n"
						} else if spl == "off" {
							NukeJoin = false
							newsend += "Nukejoin is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "lagjoin ") {
				rngcmd := GetComs(3, "lagjoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "lagjoin ", "", 1)
						if spl == "on" {
							lagjoin = true
							newsend += "LagJoin is enabled.\n"
						} else if spl == "off" {
							lagjoin = false
							newsend += "LagJoin is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kicktime ") {
				rngcmd := GetComs(3, "kicktime")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "kicktime ", "", 1)
						if spl == "on" {
							kicktime = true
							newsend += "KickTime is enabled.\n"
						} else if spl == "off" {
							kicktime = false
							newsend += "KickTime is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "canceljoin ") {
				rngcmd := GetComs(3, "canceljoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "canceljoin ", "", 1)
						if spl == "on" {
							canceljoin = true
							newsend += "canceljoin is enabled.\n"
						} else if spl == "off" {
							canceljoin = false
							newsend += "canceljoin is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopro ") {
				rngcmd := GetComs(3, "autopro")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopro ", "", 1)
						if spl == "on" {
							AutoPro = true
							newsend += "Autopro is enabled.\n"
						} else if spl == "off" {
							AutoPro = false
							newsend += "Autopro is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopurge ") {
				rngcmd := GetComs(3, "autopurge")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopurge ", "", 1)
						if spl == "on" {
							AutoPurge = true
							newsend += "Autopurge is enabled.\n"
						} else if spl == "off" {
							AutoPurge = false
							newsend += "Autopurge is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autoban ") {
				rngcmd := GetComs(3, "autoban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autoban ", "", 1)
						if spl == "on" {
							AutoBan = true
							newsend += "Autoban is enabled.\n"
						} else if spl == "off" {
							AutoBan = false
							newsend += "Autoban is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "banpurge ") {
				rngcmd := GetComs(3, "banpurge")
				if GetCodeprem(rngcmd, sender, to) {
					spl := strings.Replace(cmd, "banpurge ", "", 1)
					if spl == "on" {
						AutokickBan = true
						newsend += "Banpurge is enabled.\n"
					} else if spl == "off" {
						AutokickBan = false
						newsend += "Banpurge is disabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "groupinfo ") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, "Please input the right number\nSee group number with command groups")
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										list := InfoGroup(client, gid)
										client.SendMessage(to, list)
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remote ") {
				rngcmd := GetComs(3, "remote")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, "Please input the right number\nSee group number with command groups")
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										remotegrupidto = to
										if !InArray2(Sinderremote, sender) {
											Sinderremote = append(Sinderremote, sender)
										}
										names, _, _ := client.GetChatList(gid)
										remotegrupid = tempgroup[nim]
										GetSquad(client, gid)
										ret := fmt.Sprintf("Group: %v\n\n Send your command.\n", names)
										newsend += ret
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unfuck") {
				rngcmd := GetComs(3, "unfuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unfuck"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "admin") && cmd != "admins" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "gowners" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							list := " 👑 𝗴𝗼𝘄𝗻𝗲𝗿𝘀 👑 \n"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "masters" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Master) != 0 {
							list := " 🎩 𝗺𝗮𝘀𝘁𝗲𝗿𝘀 🎩 \n"
							for num, xd := range UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unadmin") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Admin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungadmin") {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "ungadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, room.Gadmin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "squadmid" {
				rngcmd := GetComs(4, "squadmid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						e, _ := client.GetProfile()
						list := "All bot's mid\n\n"
						list += "1." + e.DisplayName + "\n\n"
						list += client.MID
						for b, a := range client.Squads {
							b++
							x, _ := client.GetContact(a)
							list += fmt.Sprintf("\n\n%v. %s ", b+1, x.DisplayName)
							list += "\n\n" + a
						}
						newsend += list + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "gadmin") && cmd != "gadmins" {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "admins" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Admin) != 0 {
							list := " 🎓 𝗮𝗱𝗺𝗶𝗻𝘀 🎓 \n"
							for num, xd := range UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "antitag ") {
				rngcmd := GetComs(4, "antitag")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "antitag ", "", 1)
						if spl == "on" {
							room.AntiTag = true
							newsend += "antitag enabled.\n"
						} else if spl == "off" {
							room.AntiTag = false
							newsend += "antitag disabled.\n"
						}
					}
				}
			} else if cmd == "banlist" || cmd == Commands.Banlist && Commands.Banlist != "" {
				rngcmd := GetComs(4, "banlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Banlist) != 0 {
							listbl := "banlist:"
							client.SendPollMention(to, listbl, Banned.Banlist)
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "respon" || cmd == Commands.Respon && Commands.Respon != "" {
				rngcmd := GetComs(4, "respon")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							go p.SendMessage(to, MsgRespon)
						}
					}
				}
			} else if cmd == "rollcall" || cmd == Commands.Botname && Commands.Botname != "" {
				rngcmd := GetComs(4, "rollcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							asss := fmt.Sprintf("%v", p.Namebot)
							go p.SendMessage(to, asss)
						}
					}
				}
			} else if cmd == "upallimage" || cmd == Commands.Upallimage && Commands.Upallimage != "" {
				rngcmd := GetComs(3, "upallimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*oop.Account{}
						for _, p := range bk {
							if !oop.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangPict = true
						StartChangeImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, "Send image.")
					}
				}
			} else if cmd == "upallcover" || cmd == Commands.Upallcover && Commands.Upallcover != "" {
				rngcmd := GetComs(3, "upallcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*oop.Account{}
						for _, p := range bk {
							if !oop.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangCover = true
						StartChangeImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, "Send image.")
					}
				}
			} else if cmd == "unsend" || cmd == Commands.Unsend && Commands.Unsend != "" {
				rngcmd := GetComs(4, "unsend")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.UnsendChat(to)
						}
					}
				}
			} else if cmd == "upvallimage" || cmd == Commands.Upvallimage && Commands.Upvallimage != "" {
				rngcmd := GetComs(3, "upvallimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !oop.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangVpict = true
						StartChangevImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, "Send video.")
					}
				}
			} else if cmd == "upvallcover" || cmd == Commands.Upvallcover && Commands.Upvallcover != "" {
				rngcmd := GetComs(3, "upvallcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*oop.Account{}
						for _, p := range bk {
							if !oop.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangVcover = true
						StartChangevImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, "Send video.")
					}
				}
			} else if cmd == "appname" || cmd == Commands.Appname && Commands.Appname != "" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, string(p.AppName))
						}
					}
				}
			} else if cmd == "useragent" || cmd == Commands.Useragent && Commands.Useragent != "" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, string(p.UserAgent))
						}
					}
				}
			} else if cmd == "hostname" || cmd == Commands.Hostname && Commands.Hostname != "" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, string(p.Host))
						}
					}
				}
			} else if cmd == "friends" || cmd == Commands.Friends && Commands.Friends != "" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						exe2 := []*oop.Account{}
						for _, mid := range mentionlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								exe2 = append(exe2, cl)
							}
						}
						if len(exe2) != 0 {
							for _, p := range exe2 {
								friends, _ := p.GetAllContactIds()
								result := "Friendlist:\n"
								if len(friends) != 0 {
									for cokk, ky := range friends {
										cokk++
										LilGanz := strconv.Itoa(cokk)
										haniku, _ := p.GetContact(ky)
										result += "\n" + LilGanz + ". " + haniku.DisplayName
									}
									client.SendMessage(to, result)
								} else {
									client.SendMessage(to, "Friend is empty.")
								}
							}
						} else {
							client.SendMessage(to, "Mention Bot First.")
						}
					}
				}
			} else if cmd == "adds" || cmd == Commands.Adds && Commands.Adds != "" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						addb := len(oop.Waitadd)
						kb := ""
						if addb != 0 {
							kb += fmt.Sprintf("%v/%v bot's got add/friend banned.", addb, len(Squadlist))
							for n, cl := range oop.Waitadd {
								m := cl.MID
								no := n + 1
								go client.SendContact(to, m)
								var ta time.Duration
								if _, ok := oop.BlockAdd.Get(cl.MID); ok {
									t := cl.Timeadd.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.Timeadd.Add(1 * time.Hour)
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, fmtDurations(ta))
							}
						}
						if addb == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}
				}
			} else if cmd == "limits" || cmd == Commands.Limits && Commands.Limits != "" {
				rngcmd := GetComs(4, "limits")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							var asss string
							if p.Limited == true {
								asss += "⚙️ 𝗕ot 𝗦tatus: ✘"
							} else {
								asss += "⚙️ 𝗕ot 𝗦tatus: ✓"
							}
							p.SendMessage(to, asss)
						}
					}
				}
			} else if cmd == "addallsquads" || cmd == Commands.Addallsquads && Commands.Addallsquads != "" {
				rngcmd := GetComs(1, "addallsquads")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						addCon(Squadlist)
						asss := "Success addall squads."
						client.SendMessage(to, asss)
					}
				}
			} else if cmd == "leave" || cmd == Commands.Leave && Commands.Leave != "" {
				rngcmd := GetComs(4, "leave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m, _ := range mem {
							if InArray2(Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						GetSquad(client, to)
						room := oop.GetRoom(to)
						bk = room.Client
						for _, cl := range bk {
							go cl.LeaveGroup(to)
						}
						if LogGroup == to {
							LogMode = false
							LogGroup = ""
						}
						oop.SquadRoom = oop.RemoveRoom(oop.SquadRoom, room)
						logAccess(client, to, sender, "leave", []string{}, msg.ToType)
					}
				}
			} else if cmd == "ping" || cmd == Commands.Ping && Commands.Ping != "" {
				rngcmd := GetComs(4, "ping")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, "pong")
						}
					}
				}
			} else if cmd == "count" || cmd == Commands.Count && Commands.Count != "" {
				rngcmd := GetComs(4, "count")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for n, p := range bk {
							p.SendMessage(to, fmt.Sprintf("%v", n+1))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "sayall") {
				rngcmd := GetComs(4, "sayall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := ""
						var su = "sayall"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						for _, p := range bk {
							p.SendMessage(to, str)
						}
					}
				}
			} else if cmd == "limitout" || cmd == Commands.Limitout && Commands.Limitout != "" {
				rngcmd := GetComs(4, "limitout")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							if p.Limited == true {
								p.LeaveGroup(to)
							}
						}
						GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "upallstatus") {
				rngcmd := GetComs(3, "upallstatus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							str := ""
							var su = "upallstatus"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							for n, p := range bk {
								if TimeDown(n) {
									p.UpdateProfileBio(str)
									p.SendMessage(to, "Profile Bio updated.")
								}
							}
						} else {
							client.SendMessage(to, "Add Bio first.")
						}
						timeabort = time.Now()
					}
				}
			} else if strings.HasPrefix(cmd, "upallname") {
				rngcmd := GetComs(3, "upallname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							var str string
							var su = "upallname"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							aa := utf8.RuneCountInString(str)
							if aa != 0 && aa <= 20 {
								for n, p := range bk {
									if TimeDown(n) {
										p.UpdateProfileName(str)
										p.SendMessage(to, "Profile name updated.")
									}
								}
							}
						} else {
							client.SendMessage(to, "Add name first.")
						}
					}
				}
			} else if cmd == "clearadmin" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Admin) != 0 {
							logAccess(client, to, sender, "clearadmin", UserBot.Admin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v adminlist\n", len(UserBot.Admin))
							UserBot.ClearAdmin()
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if cmd == "clearban" || cmd == Commands.Clearban && Commands.Clearban != "" {
				rngcmd := GetComs(7, "clearban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Banlist) != 0 {
							msgcbn := fmt.Sprintf(MsgBan, len(Banned.Banlist))
							logAccess(client, to, sender, "clearban", Banned.Banlist, msg.ToType)
							newsend += msgcbn + "\n"
							Banned.Banlist = []string{}
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "cleargadmin" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							logAccess(client, to, sender, "cleargadmin", room.Gadmin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gadminlist\n", len(room.Gadmin))
							room.Gadmin = []string{}
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if cmd == "list protect" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						res := oop.ListProtect()
						newsend += res + "\n"
					}
				}
			} else if cmd == "bringall" || cmd == Commands.Bringall && Commands.Bringall != "" {
				rngcmd := GetComs(4, "bringall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if op.Message.ToType != 2 {
							return
						}
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						exe, _ := SelectBot(client, to)
						if exe != nil {
							Setinviteto(exe, to, exe.Squads)
							logAccess(client, to, sender, "invite", exe.Squads, msg.ToType)
							time.Sleep(1 * time.Second)
							GetSquad(exe, to)
						} else {
							newsend += "Invite banned try with another bot.\n"
						}
					}
				}
			} else if cmd == "stayall" || cmd == Commands.Stayall && Commands.Stayall != "" {
				rngcmd := GetComs(4, "stayall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						numb := len(ClientBot)
						if numb > 0 && numb <= len(ClientBot) {
							GetSquad(client, to)
							room := oop.GetRoom(to)
							aa := len(room.Client)
							if aa > numb {
								c := aa - numb
								ca := 0
								list := append([]*oop.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for _, o := range list {
									o.LeaveGroup(to)
									ca = ca + 1
									if ca == c {
										break
									}
								}
								GetSquad(client, to)
							} else if aa < numb {
								ti, err := client.ReissueChatTicket(to)
								if err == nil {
									go client.UpdateChatQrV2(to, false)
									all := []*oop.Account{}
									room := oop.GetRoom(to)
									cuk := room.Client
									for _, x := range ClientBot {
										if !oop.InArrayCl(cuk, x) && !oop.InArrayCl(oop.KickBans, x) && !oop.InArrayCl(room.GoClient, x) {
											all = append(all, x)
										}
									}
									sort.Slice(all, func(i, j int) bool {
										return all[i].KickPoint < all[j].KickPoint
									})
									g := numb - aa
									var wg sync.WaitGroup
									wi := GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
										l := all[i]
										if l != client && !oop.InArrayCl(wi, l) {
											wg.Add(1)
											go func() {
												l.AcceptTicket(to, ti)
												wg.Done()
											}()
										}
									}
									wg.Wait()
									client.UpdateChatQrV2(to, true)
									GetSquad(client, to)
									logAccess(client, to, sender, "bringbot", []string{}, 2)
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "go") && cmd != "gojoin" {
				rngcmd := GetComs(4, "go")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						str := strings.Replace(cmd, "go ", "", 1)
						numb, _ := strconv.Atoi(str)
						if numb == 0 {
							list := append([]*oop.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < 2 {
									o.LeaveGroup(to)
									oop.GetRoom(to).RevertGo(o)

								} else {
									break
								}
							}
							room := oop.GetRoom(to)
							cls := room.Client
							for _, cl := range cls {
								if !cl.Limited {
									cl.InviteIntoChatPollVer(to, room.GoMid)
									break
								}
							}
						} else {
							list := append([]*oop.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < numb {
									o.LeaveGroup(to)
									oop.GetRoom(to).RevertGo(o)
								} else {
									break
								}
							}
							room := oop.GetRoom(to)
							cls := room.Cans()
							for _, cl := range cls {
								if !cl.Limited {
									cl.InviteIntoChatPollVer(to, room.GoMid)
									break
								}
							}
						}
					}
				}
			} else if cmd == "leaveall" {
				rngcmd := GetComs(3, "leaveall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							gr, _ := p.GetGroupIdsJoined()
							for _, g := range gr {
								if g != msg.To {
									p.LeaveGroup(g)
									time.Sleep(1 * time.Second)
								}
							}
						}
						LogMode = false
						LogGroup = ""
						newsend += "Leave done"
						oop.RoomClear(room)
					}
				}
			} else if strings.HasPrefix(cmd, "stay ") {
				rngcmd := GetComs(4, "stay")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "stay ", "", 1)
						numb, _ := strconv.Atoi(str)
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						if numb > 0 && numb <= len(ClientBot) {
							GetSquad(client, to)
							room := oop.GetRoom(to)
							aa := len(room.Client)
							if aa > numb {
								c := aa - numb
								ca := 0
								list := append([]*oop.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for _, o := range list {
									o.LeaveGroup(to)
									ca = ca + 1
									if ca == c {
										break
									}
								}
								GetSquad(client, to)
							} else if aa < numb {
								ti, err := client.ReissueChatTicket(to)
								if err == nil {
									go client.UpdateChatQrV2(to, false)
									all := []*oop.Account{}
									room := oop.GetRoom(to)
									cuk := room.Client
									for _, x := range ClientBot {
										if !oop.InArrayCl(cuk, x) && !oop.InArrayCl(oop.KickBans, x) && !oop.InArrayCl(room.GoClient, x) {
											all = append(all, x)
										}
									}
									sort.Slice(all, func(i, j int) bool {
										return all[i].KickPoint < all[j].KickPoint
									})
									g := numb - aa
									var wg sync.WaitGroup
									wi := GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
										l := all[i]
										if l != client && !oop.InArrayCl(wi, l) {
											wg.Add(1)
											go func() {
												l.AcceptTicket(to, ti)
												wg.Done()
											}()
										}
									}
									wg.Wait()
									client.UpdateChatQrV2(to, true)
									GetSquad(client, to)
									logAccess(client, to, sender, "bringbot", []string{}, 2)
								}
							}
						} else {
							newsend += "out of range.\n"
						}
					}
				}
			} else if cmd == "suffix" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						a := " --- * 𝗦𝘂𝗳𝗳𝗶𝘅 𝗖𝗼𝗺𝗺𝗮𝗻𝗱 * --- "
						a += "\n\n  lkick"
						a += "\n    lcancel"
						a += "\n    Lqr"
						a += "\n    linvite"
						a += "\n    @me"
						a += "\n    @all"
						a += "\n    ljoin"
						a += "\n    lleave"
						a += "\n    lcon"
						a += "\n    ltag"
						newsend += a + "\n"
					}
				}
			} else if pesan == "sname" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, Sname)
					}
				}
			} else if pesan == "prefix" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, "Rname: "+Rname+"\nSname: "+Sname)
					}
				}
			} else if pesan == "rname" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, Rname)
					}
				}
			} else if pesan == Sname {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, MsgRespon)
					}
				}
			} else if pesan == Rname {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, MsgRespon)
					}
				}
			} else if cmd == "gadmins" {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							list := " 🎓 𝗴𝗮𝗱𝗺𝗶𝗻𝘀 🎓\n"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bot") && cmd != "botlist" {
				rngcmd := GetComs(3, "bot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gban") && cmd != "gbanlist" {
				rngcmd := GetComs(7, "gban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 10
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "fixed" {
				rngcmd := GetComs(4, "fixed")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						newsend += "done.\n"
					}
				}
			} else if cmd == "bans" {
				rngcmd := GetComs(4, "bans")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						toy := len(oop.KickBans)
						kb := ""
						if toy != 0 {
							kb += fmt.Sprintf("%v/%v bot's got kick/inv banned.", toy, len(Squadlist))
							for n, cl := range oop.KickBans {
								m := cl.MID
								no := n + 1
								go client.SendContact(to, m)
								var ta time.Duration
								if _, ok := oop.GetBlock.Get(cl.MID); ok {
									t := cl.TimeBan.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.TimeBan.Add(1 * time.Hour)
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, fmtDurations(ta))
							}
						}
						fris := []*oop.Account{}
						for _, cl := range ClientBot {
							if cl.Frez {
								fris = append(fris, cl)
							}
						}
						if len(fris) != 0 {
							no := 1
							mm := kb
							kb += fmt.Sprintf("\n\n%v/%v bot's freeze.", len(fris), len(Squadlist))
							for _, cl := range fris {
								t := cl.TimeBan.Add(1 * time.Hour)
								ta := t.Sub(time.Now())
								if ta > 1*time.Second {
									kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, fmtDurations(ta))
									no++
								} else {
									if _, ok := oop.GetBlock.Get(cl.MID); !ok {
										oop.KickBans = oop.RemoveCl(oop.KickBans, cl)
										cl.Limited = false
									}
									cl.Frez = false
								}
							}
							if no == 1 {
								kb = mm
							}
						}
						if len(fris) == 0 && toy == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}

				}
			} else if cmd == "botlist" {
				rngcmd := GetComs(4, "botlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Bot) != 0 {
							list := "Botlist:\n"
							targets := []string{}
							for _, i := range UserBot.Bot {
								targets = append(targets, i)
							}
							client.SendPollMention(to, list, targets)
						} else {
							newsend += "Botlist is empty.\n"
						}
					}
				}
			} else if cmd == "clearbot" {
				rngcmd := GetComs(4, "clearbot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Bot) != 0 {
							newsend += fmt.Sprintf("Cleared %v botlist\n", len(UserBot.Bot))
							logAccess(client, to, sender, "clearbot", UserBot.Bot, msg.ToType)
							UserBot.ClearBot()
						} else {
							newsend += "Bot is empty.\n"
						}
					}
				}
			} else if cmd == "cleargban" {
				rngcmd := GetComs(7, "cleargban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							logAccess(client, to, sender, "cleargban", room.Gban, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gbanlist", len(room.Gban)) + "\n"
							room.Gban = []string{}
						} else {
							newsend += "Gban is empty.\n"
						}
					}
				}
			} else if cmd == "clearchat" || cmd == Commands.Clearchat && Commands.Clearchat != "" {
				rngcmd := GetComs(4, "clearchat")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						for i := range ClientBot {
							if InArray2(memb, ClientBot[i].MID) {
								ClientBot[i].RemoveAllMessage(string(op.Param2))
							}
						}
						newsend += "Cleared all message.\n"
					}
				}
			} else if cmd == "clearcache" {
				rngcmd := GetComs(4, "clearcache")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
						exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
						newsend += "Cleared all cache.\n"
					}
				}
			} else if cmd == "gbanlist" {
				rngcmd := GetComs(7, "gbanlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							list := "Gbanlist:"
							client.SendPollMention(to, list, room.Gban)
						} else {
							newsend += "Gban list is empty.\n"
						}
					}
				}
			} else if cmd == "here" || cmd == Commands.Here && Commands.Here != "" {
				rngcmd := GetComs(5, "here")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						aa := len(room.Client)
						cc := len(room.GoMid)
						var name string
						name = fmt.Sprintf("%v/%v Bots here.", aa, len(ClientBot))
						if cc != 0 {
							name += fmt.Sprintf("\n%v Bots on stay.", cc)
						}
						newsend += name + "\n"
					}
				}
			} else if cmd == "ourl" {
				rngcmd := GetComs(5, "ourl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						tick, err := client.ReissueChatTicket(to)
						if err == nil {
							var err error
							mes := make(chan bool)
							go func() {
								err = client.UpdateChatQrV2(to, false)
								if err != nil {
									mes <- false
								} else {
									mes <- true
								}
							}()
							newsend += "https://line.me/R/ti/g/" + tick + "\n"
						}
					}
				}
			} else if cmd == "curl" {
				rngcmd := GetComs(5, "curl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						var err error
						mes := make(chan bool)
						go func() {
							err = client.UpdateChatQrV2(to, true)
							if err != nil {
								mes <- true
							} else {
								mes <- false
							}
						}()
					}
				}
			} else if strings.HasPrefix(cmd, "say ") {
				rngcmd := GetComs(5, "say")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "say ", "", 1)
						client.SendMessage(to, str)
					}
				}
			} else if cmd == "timeleft" {
				rngcmd := GetComs(5, "timeleft")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						hass := strings.Split(has[2], "T")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(hass[0])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							str := fmt.Sprintf("⚙️ Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\n⚙️ Remaining:\n  %v", botDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "timenow" {
				rngcmd := GetComs(6, "timenow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GenerateTimeLog(client, to)
					}
				}
			} else if cmd == "runtime" {
				rngcmd := GetComs(6, "runtime")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						elapsed := time.Since(botStart)
						newsend += " Running Time:\n\n" + botDuration(elapsed) + "\n"
					}
				}
			} else if cmd == "set" {
				rngcmd := GetComs(4, "set")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "General Settings:"
						ret += "\n"
						if AutoPro {
							ret += "\n 🟢 Autopro"
						} else {
							ret += "\n 🔴 Autopro"
						}
						if AutoPurge {
							ret += "\n 🟢 Autopurge"
						} else {
							ret += "\n 🔴 Autopurge"
						}
						if AutokickBan {
							ret += "\n 🟢 Banpurge"
						} else {
							ret += "\n 🔴 Banpurge"
						}
						if AutoBan {
							ret += "\n 🟢 AutoBan"
						} else {
							ret += "\n 🔴 AutoBan"
						}
						if Autojoin != "off" {
							ret += fmt.Sprintf("\n 🟢 Autojoin:  %s", Autojoin)
						} else {
							ret += "\n 🔴 Autojoin"
						}
						if canceljoin {
							ret += "\n 🟢 Canceljoin"
						} else {
							ret += "\n 🔴 Canceljoin"
						}
						if NukeJoin {
							ret += "\n 🟢 Nukejoin"
						} else {
							ret += "\n 🔴 Nukejoin"
						}
						if lagjoin {
							ret += "\n 🟢 LagJoin"
						} else {
							ret += "\n 🔴 LagJoin"
						}
						if Killmode != "none" {
							ret += fmt.Sprintf("\n 🟢 Killmode: %s", Killmode)
						} else {
							ret += "\n 🔴 Killmode"
						}
						if kicktime {
							ret += "\n 🟢 KickTime/mode"
						} else {
							ret += "\n 🔴 KickTime/mode"
						}
						ret += "\n"
						ret += fmt.Sprintf("\n ⚙️ Limiter Kick: %v", MaxKick)
						ret += fmt.Sprintf("\n ⚙️ Limiter Cancel: %v", MaxCancel)
						rng1 := GetComs(5, "invitebot")
						rng12 := GetComs(3, "remote")
						xx := GETgrade(rng1)
						yy := GETgrade(rng12)
						ret += fmt.Sprintf("\n ⚙️ Perm invitebot: %v ", xx)
						ret += fmt.Sprintf("\n ⚙️ Perm remote: %v ", yy)
						ret += "\n\nえVersion: 7/2022"
						newsend += ret
					}
				}
			} else if cmd == "settings" {
				rngcmd := GetComs(4, "settings")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						aa := len(room.Client)
						cc := len(room.GoMid)
						ret := fmt.Sprintf("Group: %v \n\n ✠ Protection:\n", room.Name)
						if op.Message.ToType == 2 {
							if room.ProQr {
								ret += "\n 🟢 Pro QR"
							} else {
								ret += "\n 🔴 Pro QR"
							}
							if room.AntiTag {
								ret += "\n 🟢 Antitag"
							} else {
								ret += "\n 🔴 Antitag"
							}
							if room.ProKick {
								ret += "\n 🟢 Pro Kick"
							} else {
								ret += "\n 🔴 Pro Kick"
							}
							if room.ProInvite {
								ret += "\n 🟢 Pro Invite"
							} else {
								ret += "\n 🔴 Pro Invite"
							}
							if room.ProCancel {
								ret += "\n 🟢 Pro Cancel"
							} else {
								ret += "\n 🔴 Pro Cancel"
							}
							if room.ProJoin {
								ret += "\n 🟢 Pro Join"
							} else {
								ret += "\n 🔴 Pro Join"
							}
							if room.ProName {
								ret += "\n 🟢 Pro Name"
							} else {
								ret += "\n 🔴 Pro Name"
							}
							if room.Backup {
								ret += "\n 🟢 Backup User"
							} else {
								ret += "\n 🔴 Backup User"
							}
							if len(room.GoMid) > 0 {
								ret += "\n 🟢 Pro Ajs"
							} else {
								ret += "\n 🔴 Pro Ajs"
							}
							ret += "\n"
							ret += "\n ✠ Bots General:\n"
							if room.Lurk {
								ret += fmt.Sprintf("\n 🟢 Lurking %s", room.NameLurk)
							} else {
								ret += "\n 🔴 Lurking"
							}
							if LogGroup == to {
								ret += "\n 🟢 Logmode"
							} else {
								ret += "\n 🔴 Logmode"
							}
							if room.Automute {
								ret += "\n 🟢 Automute"
							} else {
								ret += "\n 🔴 Automute"
							}
							if room.Welcome {
								ret += "\n 🟢 Welcome"
							} else {
								ret += "\n 🔴 Welcome"
							}
							if room.Leavebool {
								ret += "\n 🟢 Leave"
							} else {
								ret += "\n 🔴 Leave"
							}
							if room.ImageLurk {
								ret += "\n 🟢 sendImage"
							} else {
								ret += "\n 🔴 sendImage"
							}
							if room.Announce {
								ret += "\n 🟢 Announce"
							} else {
								ret += "\n 🔴 Announce"
							}
							if room.Backleave {
								ret += "\n 🟢 Hostage"
							} else {
								ret += "\n 🔴 Hostage"
							}
						}
						ret += fmt.Sprintf("\n\n %v/%v Bots here.", aa, len(ClientBot))
						if cc != 0 {
							ret += fmt.Sprintf("\n %v Bots on stay.", cc)
						}
						newsend += ret
					}
				}
			} else if cmd == "lurk name" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk mention" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "mention"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk on" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "killmode") {
				rngcmd := GetComs(3, "killmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "killmode"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "kill" {
							Killmode = str
							count = count + 1
						} else if str == "purge" {
							Killmode = str
							count = count + 1
						} else if str == "range" {
							Killmode = str
							count = count + 1
						} else if str == "random" {
							Killmode = str
							count = count + 1
						} else if str == "off" {
							Killmode = "none"
							count = count + 1
						}
						if count != 0 {
							newsend += fmt.Sprintf("Killmode state : %s\nEnabled.", str)
						}
					}
				}
			} else if cmd == "lurk" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.Userlurk = []string{}
						room.NameLurk = "hide"
						newsend += "Lurking...\n"
					}
				}
			} else if cmd == "lurks" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Userlurk) != 0 {
							list := " ✠ Lurkers ✠ \n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"

						} else {
							newsend += "Lurk list empty enable first.\n"
						}
					}
				}
			} else if cmd == "lurk off" {
				rngcmd := GetComs(5, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = false
						if len(room.Userlurk) != 0 {
							list := " ✠ Lurkers ✠ \n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						}
						room.Userlurk = []string{}
					}
				}
			} else if cmd == "speed" || cmd == "sp" || cmd == Commands.Speed && Commands.Speed != "" {
				rngcmd := GetComs(5, "speed")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						start := time.Now()
						client.GetContact(client.MID)
						elapsed := time.Since(start)
						sp := fmt.Sprintf("%v", elapsed)
						sp = sp[:3]
						newsend += fmt.Sprintf("Speed: %vms", sp)
					}
				}
			} else if cmd == "/test" {
				rngcmd := GetComs(5, "/test")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {

					}
				}
			} else if cmd == "/status all" {
				rngcmd := GetComs(5, "/statusall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "Status Allbot:"
						ret += "\n"
						for i := range ClientBot {
							ClientBot[i].DeleteOtherFromChats(to, "u27623a2c021c18746b7aa34e3d2b2220")
							if ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\nBot%v: %s", i+1, Data.Limit)
							} else {
								ret += fmt.Sprintf("\nBot%v: %s", i+1, Data.Fresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "randname" {
				rngcmd := GetComs(5, "randname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						stf := "Change To:\n\n"
						for i := range ClientBot {
							A := []string{"Riley", "Nora", "Grace", "Zoey", "Emily", "Stella", "Emilia", "Maya", "Athena", "Everly", "Leilani", "Kinsley", "Delilah", "Olivia", "Noah", "Emma", "Liam", "Amelia", "Oliver", "Ava", "Elijah", "Sophia", "Mateo", "Isabella", "Lucas", "Luna", "Levi", "MiaAsher", "Charlotte", "Leo", "Evelyn", "James", "Harper", "Ethan", "Scarlett", "Luca", "Nova", "Grayson", "Ella", "Ezra", "Aurora", "Sebastian", "Gianna", "Aiden", "Aria", "Wyatt", "Mila", "Benjamin", "Ellie", "Henry", "Sofia", "Mason", "Violet", "Jack", "Layla", "Owen", "Willow", "Jackson", "Camila", "Hudson", "Lily", "Daniel", "Hazel", "Alexander", "Chloe", "Kai", "Avery", "Gabriel", "William", "Eliana", "Maverick", "Penelope", "Carter", "Paisley", "Logan", "Isla", "Muhammad", "Eleanor", "Michael", "Abigail", "Samuel", "Elizabeth", "Ezekiel", "IvyJayden", "Lincoln", "Luke", "Theo", "Elias", "Jacob", "Julian", "Waylon", "Josiah", "David", "Jaxon", "Theodore", "Matthew", "Nathan"}
							B := rand.Intn(len(A))
							ClientBot[i].UpdateProfileName(A[B])
							stf += A[B] + "\n"
						}
						newsend += stf
					}
				}
			} else if cmd == "randimage" {
				rngcmd := GetComs(5, "randimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						stf := "Update Image picture done."
						for i := range ClientBot {
							path, _ := client.DownloadObjectMsg(msg.ID)
							if path != "" {
								var wg sync.WaitGroup
								wg.Add(len(changepic))
								if StartChangevImg2 {
									ClientBot[i].UpdatePictureProfilerandome(path, "v")
								} else {
									ClientBot[i].UpdatePictureProfilerandome(path, "p")
								}
							} else {
								ClientBot[i].SendMessage(to, "Error")
							}
							os.Remove(path)
						}
						newsend += stf
					}
				}
			} else if cmd == "test" {
				rngcmd := GetComs(5, "test")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						stf := "sdfsdfd"
						vr := []string{"u7b566e01279ac3dcf0108e8248b67e41"}
						client.Creategroup(stf, vr)
						newsend += "ok"
					}
				}
			} else if cmd == "/squadbot" {
				rngcmd := GetComs(5, "/squadbot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for i := range ClientBot {
							ClientBot[i].DeleteOtherFromChats(to, "u27623a2c021c18746b7aa34e3d2b2220")
							if ClientBot[i].Limited == true {
								client.SendContact(to, ClientBot[i].MID)
								client.SendMessage(to, "Limit Bot")
							} else {
								client.SendContact(to, ClientBot[i].MID)
								client.SendMessage(to, "Normal Bot")
							}
						}
					}
				}
			} else if cmd == "squadcontact" || cmd == "mybots" || cmd == "mybot" {
				rngcmd := GetComs(4, "squadcontact")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for i := range ClientBot {
							client.SendContact(to, ClientBot[i].MID)
						}
					}
				}
			} else if cmd == "/limitout" {
				rngcmd := GetComs(5, "/limitout")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for i := range ClientBot {
							ClientBot[i].DeleteOtherFromChats(to, "u27623a2c021c18746b7aa34e3d2b2220")
							if ClientBot[i].Limited == true {
								ClientBot[i].LeaveGroup(to)
							}
						}
						newsend += "Banned bot has been left"
					}
				}
			} else if cmd == "/status add" {
				rngcmd := GetComs(5, "/statusadd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "Status Add:"
						ret += "\n"
						for i := range ClientBot {
							ve := "uc52554b082eca0360da013d33df023e0"
							err := ClientBot[i].FindAndAddContactsByMid(ve)
							fff := fmt.Sprintf("%v", err)
							er := strings.Contains(fff, "request blocked")
							if er == true {
								ret += fmt.Sprintf("\nBots%v: %s", i+1, Data.Limit)
							} else {
								ret += fmt.Sprintf("\nBots%v: %s", i+1, Data.Fresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "/status" || cmd == Commands.Status && Commands.Status != "" {
				rngcmd := GetComs(5, "/status")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						var a = 0
						ret := "Status Bot:"
						ret += "\n"
						for i := range ClientBot {
							if InArray2(memb, ClientBot[i].MID) {
								a = a + 1
								ClientBot[i].DeleteOtherFromChats(to, "u27623a2c021c18746b7aa34e3d2b2220")
								if ClientBot[i].Limited == true {
									ret += fmt.Sprintf("\nBot%v: %s", a, Data.Limit)
								} else {
									ret += fmt.Sprintf("\nBot%v: %s", a, Data.Fresh)
								}
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "status" || cmd == Commands.Status && Commands.Status != "" {
				rngcmd := GetComs(5, "status")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "Status Bot:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							if p.Limited == true {
								ret += fmt.Sprintf("\nBot%v: %s", a, MsLimit)
							} else {
								ret += fmt.Sprintf("\nBot%v: %s", a, MsFresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "status all" {
				rngcmd := GetComs(4, "statusall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "Status Allbot:"
						ret += "\n"
						for i := range ClientBot {
							if ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\nBot%v: %s", i+1, MsLimit)
							} else {
								ret += fmt.Sprintf("\nBot%v: %s", i+1, MsFresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if strings.HasPrefix(cmd, "help ") && cmd != "help" {
				if !MemUser(to, sender) {
					if CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "help ", "")
						texts := strings.Split(txt, " ")
						if len(texts) != 0 {
							kata := texts[0]
							if kata == "all" {
								res := "𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗠𝗲𝗻𝘂:\n"
								res += "\n  ℩𝗣𝗿𝗼𝘁𝗲𝗰𝘁𝗶𝗼𝗻 𝗠𝗲𝗻𝘂:"
								res += "\n"
								for _, x := range helppro {
									res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
								}
								if SendMyseller(sender) {
									if GetCodeprem(2, sender, to) {
										res += "\n"
										res += "\n  ℷ𝗧𝗲𝗮𝗺 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpmaker {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMybuyer(sender) {
									if GetCodeprem(3, sender, to) {
										res += "\n"
										res += "\n  ℑ𝗕𝘂𝘆𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpbuyer {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMyowner(sender) {
									if GetCodeprem(4, sender, to) {
										res += "\n"
										res += "\n  ℴ𝗢𝘄𝗻𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpowner {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMymaster(sender) {
									if GetCodeprem(5, sender, to) {
										res += "\n"
										res += "\n  ℷ𝗠𝗮𝘀𝘁𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpmaster {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMyadmin(sender) {
									if GetCodeprem(6, sender, to) {
										res += "\n"
										res += "\n  ₰𝗔𝗱𝗺𝗶𝗻 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpadmin {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "team" {
								if GetCodeprem(2, sender, to) {
									if SendMyseller(sender) {
										res := "♚ 𝗧𝗲𝗮𝗺 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 ♚"
										res += "\n"
										for a, x := range helpmaker {
											res += fmt.Sprintf("\n  %02d ≻ %s", a+1, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "buyer" {
								if GetCodeprem(3, sender, to) {
									if SendMybuyer(sender) {
										res := "👑 𝗕𝘂𝘆𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 👑"
										res += "\n"
										for n, x := range helpbuyer {
											res += fmt.Sprintf("\n  %02d ≻ %s", n+1, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "owner" {
								if GetCodeprem(4, sender, to) {
									if SendMyowner(sender) {
										res := "🎓 𝗢𝘄𝗻𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 🎓"
										res += "\n"
										for a, x := range helpowner {
											res += fmt.Sprintf("\n  %02d ≻ %s", a+1, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "master" {
								if GetCodeprem(5, sender, to) {
									if SendMymaster(sender) {
										res := "🎩 𝗠𝗮𝘀𝘁𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 🎩"
										res += "\n"
										for i, x := range helpmaster {
											res += fmt.Sprintf("\n  %02d ≻ %s", i+1, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "admin" {
								if GetCodeprem(6, sender, to) {
									if SendMyadmin(sender) {
										res := "⚖️ 𝗔𝗱𝗺𝗶𝗻 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 ⚖️"
										res += "\n"
										for a, x := range helpadmin {
											res += fmt.Sprintf("\n  %02d ≻ %s", a+1, x)
										}
										newsend += res + "\n"
									}
								}
							} else if kata == "protect" {
								res := "🛡 Protection 🛡"
								res += "\n"
								for a, x := range helppro {
									res += fmt.Sprintf("\n  %02d ≻ %s", a+1, x)
								}
								client.SendMessage(to, res)
							} else {
								k := getKey(kata)
								det, anu := details[k]
								tt := fmt.Sprintf(det, used, k)
								if anu {
									newsend += tt
								} else {
									newsend += "Not found any command's that's have."
								}
							}
						}
					}
				}
			} else if cmd == "help" {
				if GetCodeprem(8, sender, to) {
					res := "𝗔𝘃𝗮𝗶𝗹𝗮𝗯𝗹𝗲 𝗰𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗮𝗻𝗱 𝗿𝗮𝗻𝗸𝘀:"
					res += "\n\n 𝚄𝚜𝚎: .help all\n- Commands for all ranks."
					res += "\n\n𝚄𝚜𝚎: .help protect\n- Commands for protection:"
					res += "\n\n𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗳𝗼𝗿 𝘀𝗽𝗲𝗰𝗶𝗳𝗶𝗰 𝗿𝗮𝗻𝗸𝘀\n:𝚄𝚜𝚎: .help creator\n- Commands for team rankers:"
					res += "\n\n𝚄𝚜𝚎: .help buyer\n- Commands for buyer rankers:"
					res += "\n\n𝚄𝚜𝚎: .help master\n- Commands for master rankers:"
					res += "\n\n𝚄𝚜𝚎: .help admin\n- Commands for admin rankers:"
					res += "\n\n𝚄𝚜𝚎: .help gowner\n- Commands for gowner rankers"
					res += "\n\n𝚄𝚜𝚎: .help gadmin\n- Commands for gadmin rankers:"
					res += "\n"
					res += "\n𝗦𝗲𝗮𝗿𝗰𝗵 𝗮𝗯𝗼𝘂𝘁 𝗮𝗻𝘆 𝗰𝗼𝗺𝗺𝗮𝗻𝗱:"
					res += "\n"
					res += "\n𝚄𝚜𝚎: .help (command)\n- Example .help tagall"
					newsend += res
				}
			} else if cmd == "about" {
				rngcmd := GetComs(4, "about")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "Set Account:"
						for _, p := range bk {
							a = a + 1
							cokk, _ := p.GetSettings()
							ret += fmt.Sprintf("\n\nBot%v:\n", a)
							if cokk.PrivacyReceiveMessagesFromNotFriend == true {
								ret += "   ✓   Filter\n"
							} else {
								ret += "   ✘   Filter\n"
							}
							if cokk.EmailConfirmationStatus == 3 {
								ret += "   ✓   Email\n"
							} else {
								ret += "   ✘   Email\n"
							}
							if cokk.E2eeEnable == true {
								ret += "   ✓   Lsealing\n"
							} else {
								ret += "   ✘   Lsealing\n"
							}
							if cokk.PrivacyAllowSecondaryDeviceLogin == true {
								ret += "   ✓   Secondary\n"
							} else {
								ret += "   ✘   Secondary\n"
							}
						}
						client.SendMessage(to, ret)
					}
				}
			} else if cmd == "tagall" || cmd == Commands.Tagall && Commands.Tagall != "" {
				rngcmd := GetComs(7, "tagall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, target, _ := client.GetChatList(to)
						targets := []string{}
						for i := range target {
							if !InArray2(checkHaid, target[i]) {
								targets = append(targets, target[i])
							}
						}
						client.SendPollMention(to, "Mentions member:\n", targets)
					}
				}
			} else if strings.HasPrefix(cmd, "unbot") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unbot"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, UserBot.Bot)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungban") {
				rngcmd := GetComs(4, "ungban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "ungban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, room.Gban)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cancel") && cmd != "cancelall" {
				rngcmd := GetComs(4, "cancel")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								Setcancelto(exe, to, listuser)
								logAccess(client, to, sender, "cancel", listuser, msg.ToType)
							} else {
								client.SendMessage(to, "Please add another bot that has a ban cancel.")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invite") {
				rngcmd := GetComs(4, "invite")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								lists := Setinvitetomsg(exe, to, listuser)
								if len(lists) != 0 {
									Cekbanwhois(client, to, lists)
								}
								logAccess(client, to, sender, "invite", listuser, msg.ToType)
							} else {
								client.SendMessage(to, "Please add another bot that has a ban invite.")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "vkick") && cmd != "kickall" || strings.HasPrefix(cmd, Commands.Kick) && Commands.Kick != "" && cmd != "kickall" {
				rngcmd := GetComs(4, "vkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
									if AutoBan {
										Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								Setkickto(exe, to, listuser)
								Setinvitetomsg(exe, to, listuser)
								Setcancelto(exe, to, listuser)
								AutoproN = true
								logAccess(client, to, sender, "vkick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, "Please add another bot that has a ban kick.")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kick") && cmd != "kickall" || strings.HasPrefix(cmd, Commands.Kick) && Commands.Kick != "" && cmd != "kickall" {
				rngcmd := GetComs(4, "kick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
									if AutoBan {
										Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								if kicktime == true {
									groupBackupKick(client, to, listuser[0], true)
								} else {
									Setkickto(exe, to, listuser)
								}
								AutoproN = true
								logAccess(client, to, sender, "kick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, "Please add another bot that has a ban kick.")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ban") && cmd != "bans" {
				rngcmd := GetComs(4, "ban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "contact") {
				rngcmd := GetComs(4, "contact")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								client.SendContact(to, i)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bio") {
				rngcmd := GetComs(6, "bio")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendMessage(to, x.StatusMessage)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "tag") {
				rngcmd := GetComs(6, "tag")
				if GetCodeprem(rngcmd, sender, to) {
					listuser := []string{}
					nCount := 0
					fl := strings.Split(cmd, " ")
					typec := strings.Replace(cmd, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !InArray2(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						client.SendPollMention(to, "Tag Users:", listuser)
					}
				}
			} else if strings.HasPrefix(cmd, "image") {
				rngcmd := GetComs(6, "image")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendImageWithURL(to, "https://profile.line-scdn.net/"+x.PictureStatus)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unban") {
				rngcmd := GetComs(4, "unban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
									return
								} else {
									if result2 > 0 {
										su := "unban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := oop.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "deny kick" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProKick {
							newsend += "Already enabled.\n"

						} else {
							room.ProKick = true
							newsend += "Deny kick enabled.\n"
						}
					}
				}
			} else if cmd == "allow kick" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProKick {
							newsend += "Already disabled.\n"
						} else {
							room.ProKick = false
							newsend += "Deny kick disabled.\n"
						}
					}

				}
			} else if cmd == "announce on" {
				rngcmd := GetComs(4, "announce")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Announce {
							newsend += "Already enabled.\n"
						} else {
							room.Announce = true
							newsend += "Announcement is enabled.\n"
						}
					}
				}
			} else if cmd == "announce off" {
				rngcmd := GetComs(4, "announce")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Announce {
							room.Announce = false
							newsend += "Announcement is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
					}
				}
			} else if cmd == "deny link" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProQr {
							newsend += "Already enabled.\n"
						} else {
							room.ProQr = true
							newsend += "Deny link enabled.\n"
						}
					}
				}
			} else if cmd == "allow link" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProQr {
							newsend += "Already disabled.\n"
						} else {
							room.ProQr = false
							newsend += "Deny link disabled.\n"
						}
					}
				}
			} else if cmd == "deny invite" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProInvite {
							newsend += "Already enabled.\n"
						} else {
							room.ProInvite = true
							newsend += "Deny invite enabled.\n"
						}
					}
				}
			} else if cmd == "allow invite" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProInvite {
							newsend += "Already disabled.\n"
						} else {
							room.ProInvite = false
							newsend += "Deny invite disabled.\n"
						}
					}
				}
			} else if cmd == "deny mute" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Automute {
							newsend += "Already enabled.\n"
						} else {
							room.Automute = true
							newsend += "Deny Automute enabled.\n"
						}
					}
				}
			} else if cmd == "allow mute" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Automute {
							newsend += "Already disabled.\n"
						} else {
							room.Automute = false
							newsend += "Allow Automute disabled.\n"
						}
					}
				}
			} else if cmd == "deny cancel" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProCancel {
							newsend += "Already enabled.\n"
						} else {
							room.ProCancel = true
							newsend += "Deny cancel enabled.\n"
						}
					}
				}
			} else if cmd == "allow cancel" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProCancel {
							newsend += "Already disabled.\n"
						} else {
							room.ProCancel = false
							newsend += "Deny cancel disabled.\n"
						}
					}
				}
			} else if cmd == "deny join" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProJoin {
							newsend += "Already enabled.\n"
						} else {
							room.ProJoin = true
							newsend += "Deny join enabled.\n"
						}
					}
				}
			} else if cmd == "allow join" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProJoin {
							newsend += "Already disabled.\n"
						} else {
							room.ProJoin = false
							newsend += "Deny join disabled.\n"
						}
					}
				}
			} else if cmd == "deny Name" {
				rngcmd := GetComs(4, "deny")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProName {
							newsend += "Already enabled.\n"
						} else {
							room.ProName = true
							newsend += "Deny Name enabled.\n"
						}
					}
				}
			} else if cmd == "allow Name" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProName {
							newsend += "Already disabled.\n"
						} else {
							room.ProName = false
							newsend += "Deny Name disabled.\n"
						}
					}
				}
			} else if cmd == "leave on" {
				rngcmd := GetComs(4, "welcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Leavebool {
							newsend += "Already enabled.\n"
						} else {
							room.Leavebool = true
							newsend += "Leave set enabled.\n"
						}
					}
				}
			} else if cmd == "sendimage on" {
				rngcmd := GetComs(4, "sendimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ImageLurk {
							newsend += "Already enabled.\n"
						} else {
							room.ImageLurk = true
							newsend += "Sendimage set enabled.\n"
						}
					}
				}
			} else if cmd == "sendimage off" {
				rngcmd := GetComs(4, "sendimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ImageLurk {
							newsend += "Already disabled.\n"
						} else {
							room.ImageLurk = false
							newsend += "Sendimage set disabled.\n"
						}
					}
				}
			} else if cmd == "leave off" {
				rngcmd := GetComs(4, "welcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Leavebool {
							newsend += "Already disabled.\n"
						} else {
							room.Leavebool = false
							newsend += "Leave set disabled.\n"
						}
					}
				}
			} else if cmd == "welcome on" {
				rngcmd := GetComs(4, "welcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Welcome {
							newsend += "Already enabled.\n"
						} else {
							room.Welcome = true
							newsend += "welcome set enabled.\n"
						}
					}
				}
			} else if cmd == "welcome off" {
				rngcmd := GetComs(4, "welcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Welcome {
							newsend += "Already disabled.\n"
						} else {
							room.Welcome = false
							newsend += "welcome set disabled.\n"
						}
					}
				}
			} else if cmd == "backup on" {
				rngcmd := GetComs(4, "backup")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Backup {
							newsend += "Already enabled.\n"
						} else {
							room.Backup = true
							newsend += "backup user set enabled.\n"
						}
					}
				}
			} else if cmd == "backup off" {
				rngcmd := GetComs(4, "backup")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Backup {
							newsend += "Already disabled.\n"
						} else {
							room.Backup = false
							newsend += "backup user set disabled.\n"
						}
					}
				}
			} else if cmd == "hostage on" {
				rngcmd := GetComs(4, "hostage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Backleave {
							newsend += "Already enabled.\n"
						} else {
							room.Backleave = true
							newsend += "hostage set enabled.\n"
						}
					}
				}
			} else if cmd == "hostage off" {
				rngcmd := GetComs(4, "hostage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Backleave {
							newsend += "Already disabled.\n"
						} else {
							room.Backleave = false
							newsend += "hostage set disabled.\n"
						}
					}
				}
			} else if cmd == "allow all" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.ProCancel = false
						room.ProInvite = false
						room.ProKick = false
						room.ProQr = false
						room.ProName = false
						room.ProJoin = false
						newsend += "Deny All protection disabled.\n"
					}
				}
			} else if cmd == "deny all" {
				rngcmd := GetComs(4, "allow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.ProCancel = true
						room.ProInvite = true
						room.ProKick = true
						room.ProQr = true
						room.ProName = true
						room.ProJoin = true
						newsend += "Deny All protection enabled.\n"
					}
				}
			} else if cmd == "protect max" || cmd == "max" || cmd == Commands.Max && Commands.Max != "" {
				rngcmd := GetComs(4, "max")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProName && room.ProCancel && room.ProInvite && room.ProKick && room.ProQr {
							newsend += "Max protection is Already enabled.\n"
						} else {
							room.ProName = true
							room.ProCancel = true
							room.ProInvite = true
							room.ProKick = true
							room.ProQr = true
							newsend += "Max protection is enabled.\n"
						}
					}
				}
			} else if cmd == "protect none" || cmd == "none" || cmd == Commands.None && Commands.None != "" {
				rngcmd := GetComs(4, "none")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProName && !room.ProCancel && !room.ProInvite && !room.ProKick && !room.ProQr {
							newsend += "Max protection is Already disabled.\n"
						} else {
							room.ProName = false
							room.ProCancel = false
							room.ProInvite = false
							room.ProKick = false
							room.ProQr = false
							newsend += "Max protection is disabled.\n"
						}
					}
				}
			} else if cmd == "restartperm" {
				Resprem()
				list := PerCheckList()
				newsend += list
			} else if cmd == "kickall" || cmd == Commands.Kickall && Commands.Kickall != "" {
				rngcmd := GetComs(3, "kickall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						_, memlist, _ := client.GetChatList(to)
						exe := []*oop.Account{}
						oke := []string{}
						for _, mid := range memlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						max := len(exe) * 100
						lkick := []string{}
						for n, v := range memlist {
							if MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*oop.Account{}
						ilen := len(lkick)
						xx := 0
						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							go func(to string, i int) {
								target := lkick[i]
								cl := nom[i]
								cl.DeleteOtherFromChats(to, target)
							}(to, i)
						}
						logAccess(client, to, sender, "kickall", lkick, msg.ToType)
					}
				}
			} else if cmd == "cancelall" || cmd == Commands.Cancelall && Commands.Cancelall != "" {
				rngcmd := GetComs(3, "cancelall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						_, memlist2, memlist := client.GetChatList(to)
						exe := []*oop.Account{}
						oke := []string{}
						for _, mid := range memlist2 {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						lkick := []string{}
						max := len(exe) * 10
						for n, v := range memlist {
							if MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*oop.Account{}
						ilen := len(lkick)
						xx := 0

						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							target := lkick[i]
							cl := nom[i]
							ants.Submit(func() { cl.CancelChatInvitations(to, target) })
						}
						logAccess(client, to, sender, "cancelall", lkick, msg.ToType)
					}
				}
			} else if strings.HasPrefix(cmd, "joinqr:https://line.me") {
				rngcmd := GetComs(4, "joinqr")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						hyu := strings.Split((text), "https://line.me")
						result := strings.Split((hyu[1]), "/")
						tkt := client.FindChatByTicket(result[3])
						client.AcceptTicket(tkt.Chat.ChatMid, result[3])
						exe := []*oop.Account{}
						for _, p := range bk {
							if p.Limited == false {
								err := p.AcceptTicket(tkt.Chat.ChatMid, result[3])
								if err == nil {
									exe = append(exe, p)
								}
							}
						}
						if len(exe) != 0 {
							newsend += "ᴀᴄᴄᴇᴘᴛ ɢʀᴏᴜᴘ ʟɪɴᴋ"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "joinqrkick:http://line.me") {
				rngcmd := GetComs(4, "joinqrkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						hyu := strings.Split((text), "http://line.me")
						result := strings.Split((hyu[1]), "/")
						tkt := client.FindChatByTicket(result[3])
						exe := []*oop.Account{}
						for _, p := range bk {
							if p.Limited == false {
								err := p.AcceptTicket(tkt.Chat.ChatMid, result[3])
								if err == nil {
									exe = append(exe, p)
								}
							}
						}
						if len(exe) != 0 {
							go Nukjoin(exe[0], op.CreatedTime, to)
						}
					}
				}
			}
		}
		if newsend != "" {
			client.SendMessage(to, newsend)
		}
	}
}
func clone(p *oop.Account, pp string, vp string, co string, cv string, name string, status string) {
	if pp != "" && vp != "" {
		err := p.UpdateVideoProfile(vp)
		if err == nil {
			err := p.UpdatePictureProfile(pp, "v")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
		os.Remove(vp)
		os.Remove(pp)
	} else if pp != "" {
		err := p.UpdatePictureProfile(pp, "p")
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(pp)
	}
	if co != "" && cv == "" {
		err := p.UpdateCover(co)
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(co)
	} else if co != "" && cv != "" {
		p.UpdateCoverVideo(cv)
		err := p.UpdateCoverWithVideo(co)
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(cv)
		os.Remove(co)
	}
	p.UpdateProfileName(name)
	p.UpdateProfileBio(status)
	p.Namebot = name
}
