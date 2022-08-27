package main

import (
	"botline/hashmap"
	talkservice "botline/linethrift"
	"botline/oop"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type mentions struct {
	MENTIONEES []struct {
		Start string `json:"S"`
		End   string `json:"E"`
		Mid   string `json:"M"`
	} `json:"MENTIONEES"`
}

type kickop struct {
	Kick  []string
	Inv   []string
	Opinv []int64
}

type Stickers struct {
	Id  string
	Pid string
}

type clustering struct {
	mem string
	tm  int64
	fr  []string
}

type DATA struct {
	Authoken      []string            `json:"Authoken"`
	CreatorBack   []string            `json:"CreatorBack"`
	BuyerBack     []string            `json:"BuyerBack"`
	OwnerBack     []string            `json:"OwnerBack"`
	MasterBack    []string            `json:"MasterBack"`
	AdminBack     []string            `json:"AdminBack"`
	ResponBack    string              `json:"ResponBack"`
	RnameBack     string              `json:"RnameBack"`
	SnameBack     string              `json:"SnameBack"`
	BotBack       []string            `json:"BotBack"`
	SellerBack    []string            `json:"SellerBack"`
	BanBack       []string            `json:"BanBack"`
	FuckBack      []string            `json:"FuckBack"`
	Logobot       string              `json:"Logobot"`
	LogGroup      string              `json:"LogGroup"`
	MuteBack      []string            `json:"MuteBack"`
	AnnunceBack   []string            `json:"AnnunceBack"`
	ProQrBack     []string            `json:"ProQrBack"`
	ProjoinBack   []string            `json:"ProjoinBack"`
	ProInviteBack []string            `json:"ProInviteBack"`
	ProCancelBack []string            `json:"ProCancelBack"`
	ProkickBack   []string            `json:"ProkickBack"`
	GbanBack      map[string][]string `json:"GbanBack"`
	GadminBack    map[string][]string `json:"GadminBack"`
	GownerBack    map[string][]string `json:"GownerBack"`
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
	SetHelper        = &oop.Helper{}
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
	MsgRespon        = "𝙷𝚒,𝙳𝚊𝚍𝚢..."
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
	canceljoin       = false
	AutoBan          = true
	Autojoin         = "qr"
	backlist         = &hashmap.HashMap{}
	cekoptime        = []int64{}
	Ceknuke          = &hashmap.HashMap{}
	Cekstaybot       = &hashmap.HashMap{}
	Commands         = &oop.Command{}
	Waitlistin       = map[string][]string{}
	AutoproN         = false
	LogMode          = false
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
	Dalltime         = "2022-07-20T00:00:00Z"
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
	GO               = getArg() // argsRaw[1]
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
	MaxCancel        = 100
	MaxKick          = 100
	cekGo            = []int64{}
	UpdatePicture    = map[string]bool{}
	UpdateCover      = map[string]bool{}
	UpdateVProfile   = map[string]bool{}
	UpdateVCover     = map[string]bool{}
	Qrwar            = false
	filterWar        = &kickop{Kick: []string{}, Inv: []string{}, Opinv: []int64{}}
	ColorCyan        = "\033[36m"
	ColorReset       = "\033[0m"
	tempban          = []string{}
	Data             DATA
	remotegrupid            = ""
	LastActive              = &hashmap.HashMap{}
	used                    = ""
	IPServer         string = "192.198.1.1"
	Killmode                = "kill"
	filtermsg               = &hashmap.HashMap{}
	Opinvite                = []int64{}
	stringToInt             = []rune("01")
	DATABASE                = "db/" + GO + ".json"
	MAKERS                  = []string{"u27603a2c021c18746b7aa34e3d2b2220", "u7b566e01279ac3dcf0108e8248b67e41"}
	HostName                = []string{
		"legy-jp-addr-long",
		"legy-jp-short",
		"legy-jp-long",
		"gwx",
		"gm2",
		"ga2",
		"gd2",
		"gfp",
		"gf",
		"legy-backup",
		"legy-gslb",
		"legy-jp-addr",
		"gw",
		"legy-jp",
		"gwz",
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
		"canceljoin":   "'%s%s' on/off\ncancelall member's while bot has invited..",
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

func getArg() string {

	args := os.Args
	if len(os.Args) <= 1 {
		fmt.Println("\033[0;31m not enoght args")
		fmt.Println("\033[37m try :\n\t  \033[33m <app-name> <arg>")
		fmt.Println("\033[37m for example:\n\t \033[33m  ./botline dir-name")
		fmt.Println("\033[37m or\n\t \033[33m go run *go dir-name", "")
		fmt.Println()
		os.Exit(0)
	}
	return args[1]
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
