package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	talkservice "github.com/bashery/linethrift"

	"github.com/bashery/botline/hashmap"
	"github.com/bashery/botline/oop"
)

func getArg() string {
	args := os.Args
	if len(os.Args) <= 1 {
		fmt.Println("\033[0;31m not enoght args")
		fmt.Println("\033[37m try :\n\t  \033[33m <app-name> <arg>")
		fmt.Println("\033[37m for example:\n\t \033[33m  ./botline dir-name")
		fmt.Println("\033[37m or\n\t \033[33m go run *go dir-name", "")
		os.Exit(0)
	}
	return args[1]
}

var (
	GO               = getArg()
	Whitelist        = []string{}
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
	AutoBan          = false
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
	DATABASE    = "db/" + getArg() + ".json"
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
		"Clearlistcmd",
		"Clearmute",
		"Clone:",
		"Decline",
		"Declineall",
		"Expel",
		"Friends",
		"Friendlist",
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
		"Msgclearban",
		"Msglimit",
		"Msglurk",
		"Msgrespon",
		"Msgstatus",
		"Nukejoin",
		"Owner",
		"Perm",
		"Permlist:",
		"Refresh",
		"Remote",
		"Stats",
		"Setcancel",
		"setlogo",
		"Setkick",
		"Setlimiter",
		"Setrname",
		"Setsname",
		"Statusall",
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
		"Allgaccess",
		"Antitag",
		"Bringall",
		"Clearcache",
		"Clearchat",
		"Cleargban",
		"Cleargowner",
		"Contactsquad",
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
		"shutdown":     "'%s%s'\n\nShutting down the bot's.",
		"perm":         "'%s%s .<grade>.<command>'\n\nAvailable grade buyer/owner/master/admin",
		"nukejoin":     "'%s%s' on/off\nkickall member's while bot has invited..",
		"announce":     "'%s%s on/off'\n\nEnable detect announce.",
		"hostage":      "'%s%s on/off'\n\nEnable auto invite leave member.",
		"accept":       "'%s%s <number>'\n\nAccept group invitation by number.",
		"reject":       "'%s%s <number>'\n\nReject group invitation by number.",
		"welcome":      "'%s%s on/off'\n\nEnable welcome message.",
		"leave":        "'%s%s on/off'\n\nEnable leave message.",
		"setcmd":       "'%s%s <state> <command>'\n\nUsed to enabling/disabling command\nAvailable state lock/unlock/disable/enable.",
		"fixed":        "'%s%s'\n\nIf bot's error, please use this command to autofix.",
		"logmode":      "'%s%s <state>'\n\nUsed to see bot's activity.\nAvailable state on/off",
		"go":           "'%s%s <number>'\n\nSet bot to stay on group invitation.\nDefault is 2 bot.",
		"unseller":     "'%s%s <range/lcon/lkick/etc>'Used to expel seller.\nAvailable range '<', '>', '-', ',' with number.",
		"unbuyer":      "'%s%s <range/lcon/lkick/etc>'Used to expel buyer.\nAvailable range '<', '>', '-', ',' with number.",
		"unowner":      "'%s%s <range/lcon/lkick/etc>'Used to expel owner.\nAvailable range '<', '>', '-', ',' with number.",
		"unadmin":      "'%s%s <range/lcon/lkick/etc>'Used to expel admin.\nAvailable range '<', '>', '-', ',' with number.",
		"unmaster":     "'%s%s <range/lcon/lkick/etc>'Used to expel master.\nAvailable range '<', '>', '-', ',' with number.",
		"ungowner":     "'%s%s <range/lcon/lkick/etc>'Used to expel gowner.\nAvailable range '<', '>', '-', ',' with number.",
		"ungadmin":     "'%s%s <range/lcon/lkick/etc>'Used to expel gadmin.\nAvailable range '<', '>', '-', ',' with number.",
		"clearseller":  "'%s%s'\n\nClearing all sellers.",
		"clearbuyer":   "'%s%s'\n\nClearing all buyer list.",
		"clearowner":   "'%s%s'\n\nClearing all owner list.",
		"clearmaster":  "'%s%s'\n\nClearing all master list.",
		"clearadmin":   "'%s%s'\n\nClearing all admin list.",
		"cleargadmin":  "'%s%s'\n\nClearing all gadmin list.",
		"cleargowner":  "'%s%s'\n\nClearing all gowner list.",
		"clearbot":     "'%s%s'\n\nClearing all bot list.",
		"clearban":     "'%s%s'\n\nClearing all ban list.",
		"clearfuck":    "'%s%s'\n\nClearing all fuck list.",
		"clearmute":    "'%s%s'\n\nClearing all mute list.",
		"cleargban":    "'%s%s'\n\nClearing all gban list.",
		"clearchat":    "'%s%s'\n\nClearing all squad messages.",
		"upvallimage":  "'%s%s'\n\nUpdating all bot's video profile.",
		"upvimage":     "'%s%s'\n\nUpdating all bot's video profile.",
		"upallimage":   "'%s%s'\n\nUpdating all bot's picture profile.",
		"upimage":      "'%s%s'\n\nUpdating bot's profile picture.",
		"upvallcover":  "'%s%s'\n\nUpdating all bot's video cover.",
		"upvcover":     "'%s%s @tag bot'\n\nUpdating bot's video cover.",
		"upcover":      "'%s%s' @tag\n\nUpdating bot's cover picture.",
		"upallcover":   "'%s%s'\n\nUpdating all bot's cover picture.",
		"upname":       "'%s%s newname'\n\nUpdating bot's displayname.",
		"upallname":    "'%s%s newname'\n\nUpdating all bot's displayname.",
		"leaveall":     "'%s%s'\n\nleave all bot's from all group's.",
		"groups":       "'%s%s'\n\nsee bot group's.",
		"stayall":      "'%s%s'\n\naccepting all group invitation.",
		"setcom":       "'%s%s .key .value'\n\nChange command.",
		"upstatus":     "'%s%s <status message>'\n\nUpdating bot's profile bio.",
		"upallstatus":  "'%s%s <status message>'\n\nUpdating all bot's profile bio.",
		"kick":         "'%s%s @tag/lcon/lkick/etc'\n\nKick member's.",
		"prefix":       "'%s%s on/off'\n\nEnable/disable prefix.",
		"list protect": "'%s%s'\n\nShow all protection group's.",
		"invme":        "'%s%s gnumber'\n\nInvite user to the destination group.",
		"autojoin":     "'%s%s qr/invite/off'\n\nForcing bot's to joinall while invited.",
		"autoban":      "'%s%s on/off'\n\nAuto banned user.",
		"sellers":      "'%s%s'\n\nShow seller list.",
		"buyers":       "'%s%s'\n\nShow buyer list.",
		"owners":       "'%s%s'\n\nShow owner list.",
		"masters":      "'%s%s'\n\nShow master list.",
		"admins":       "'%s%s'\n\nShow admin list.",
		"gowners":      "'%s%s'\n\nShow gowner list.",
		"gadmins":      "'%s%s'\n\nShow gadmin list.",
		"botlist":      "'%s%s'\n\nShow bot list.",
		"banlist":      "'%s%s'\n\nShow ban list.",
		"fucklist":     "'%s%s'\n\nShow fuck list.",
		"mutelist":     "'%s%s'\n\nShow mutelist list.",
		"gbanlist":     "'%s%s'\n\nShow gban list.",
		"hides":        "'%s%s'\n\nShow Invisible user.",
		"hide":         "'%s%s @tag/lcon/lkick/etc'\n\nAdded user to invisible list.",
		"kickall":      "'%s%s'\n\nKick all group member's.",
		"group info":   "'%s%s'\n\nShow all group member's./pendings/access",
		"autopurge":    "'%s%s on/off'\n\nEnable autopurge.",
		"lurk":         "'%s%s on/off'\n\nEnable lurking mode.",
		"lurkmsg":      "'%s%s <message>'\n\nSet lurk message.\nUse @! for placing user tagging.",
		"antitag":      "'%s%s on/off'\n\nEnable antitag.",
		"killmode":     "'%s%s kill/purge/on/off/range'\n\nKiller mode to kick all banlist/squad.",
		"autopro":      "'%s%s on/off'\n\nAuto protect max while bot's join.",
		"setlimit":     "'%s%s number'\n\nSet max kick in killmode /bot.",
		"stay":         "'%s%s number'\n\nSet amount of bot's in group invite via link invitation.",
		"bringall":     "'%s%s'\n\nBring all bot's by invitation.",
		"bring":        "'%s%s number'\n\nSet amount of bot's in group via invitation.",
		"here":         "'%s%s'\n\nShow amount of bot's in group.",
		"friends":      "'%s%s'\n\nShow all bot's friends.",
		"msgrespon":    "'%s%s respon'\n\nSet bot's response.",
		"msgwelcome":   "'%s%s <message>'\n\nSet welcome message each group.\nParameter for changing need to adding @user for replacing username and @group for replacing groupname.",
		"setrname":     "'%s%s newrname'\n\nChange the rname prefix.",
		"setsname":     "'%s%s newsname'\n\nChange the sname prefix.",
		"invite":       "'%s%s @tag/lcon/lkick/etc'\n\nInvite target to the group's.",
		"clone":        "'%s%s @tag/lcon/lkick/etc @tagbot'\n\nCloning targte profile.",
		"gaccess":      "'%s%s'\n\nSee all group access list.",
		"limitout":     "'%s%s'\n\nLeave the kicbanned bot's.",
		"say":          "'%s%s word'\n\nThe bot's would said the word.",
		"sayall":       "'%s%s word'\n\nAll bot's would said the word.",
		"expel":        "'%s%s @tag/lcon/lkick/etc'\n\nUsed to expel user access.",
		"respon":       "'%s%s'\n\nBot response.",
		"ping":         "'%s%s'\n\nBot response.",
		"permlist":     "'%s%s key'\n\nGet the command value.",
		"settings":     "'%s%s'\n\nShow the group preset status in group.",
		"set":          "'%s%s'\n\nShow the bot's set.",
		"help":         "'%s%s'\n\nShow the help command.",
		"deny":         "'%s%s invite/kick/qr/join/cancel/off/all/max'\n\nEnable the protection.",
		"allow":        "'%s%s invite/kick/qr/join/cancel/all'\n\nDisable the protection.",
		"ourl":         "'%s%s'\n\nOpen group links.",
		"curl":         "'%s%s'\n\nClose group links.",
		"mysquad":      "'%s%s'\n\nSend squad contact's",
		"count":        "'%s%s'\n\nShow bot's number.",
		"speed":        "'%s%s'\n\nShow bot response speed.",
		"unsend":       "'%s%s count'\n\nUnsend recent bot's message.\nIf count not definde, it would unsend all recent message.",
		"tagall":       "'%s%s'\n\nTagging all member's.",
		"ftagall":      "'%s%s'\n\nTagging all member's with sticker.",
		"access":       "'%s%s'\n\nShow all bot access.",
		"bans":         "'%s%s'\n\nShow the bot's status.",
		"runtime":      "'%s%s'\n\nShiw the bot's time alive.",
		"timeleft":     "'%s%s'\n\nShow the bot's timeleft.",
		"linvite":      "'%s%s'\n\nShow the last invited in group.",
		"lkick":        "'%s%s'\n\nShow the last kicked in group.",
		"lmid":         "'%s%s'\n\nShow the last mid in group.",
		"lcon":         "'%s%s'\n\nShow the last contact in group.",
		"ltag":         "'%s%s'\n\nShow the last tag in group.",
		"lban":         "'%s%s'\n\nShow the last banned in group.",
		"lcancel":      "'%s%s'\n\nShow the last cancel in group.",
		"lqr":          "'%s%s'\n\nShow the last upded qr in group.",
		"ljoin":        "'%s%s'\n\nShow the last join in group.",
		"lleave":       "'%s%s'\n\nShow the last leave in group.",
		"abort":        "'%s%s'\n\nAborting command.",
		"groupcast":    "'%s%s <your word>'\n\nBroadcasting message to all groups.",
		"contact":      "'%s%s @tag/lcon/lkick/etc'\n\nUsed to get contact's.",
		"rollcall":     "'%s%s'\n\nShow bot's name.",
		"gojoin":       "'%s%s'\n\nJoining bot's from invitation list.",
		"mid":          "'%s%s @tag/lcon/lkick/etc'\n\nGet midlist.",
		"name":         "'%s%s @tag/lcon/lkick/etc'\n\nGet namelist.",
		"purgeall":     "'%s%s'\n\nPurge all banlist in all group.",
		"squadmid":     "'%s%s'\n\nShow all bots mid.",
		"whois":        "'%s%s @tag/lcon/lkick/etc'\n\nSee user info.",
		"cancel":       "'%s%s @tag/lcon/lkick/etc'\n\nCancel group invitation.",
		"remote":       "'%s%s:'\n\nthe right number\nSee group number with command groups.\nExample:\n  remote: 2 gmember.\nund send command.",
	}
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
	rngcmd := GetComs(3, "clone")
	rngcmd = GetComs(3, "joinqrkick")
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
	rngcmd = GetComs(3, "msgclearban")
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
	rngcmd = GetComs(3, "clearlistcmd")
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
	//return false
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
	//return false
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
