package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

	"botline/hashmap"
	talkservice "botline/linethrift"
	"botline/oop"

	"github.com/kardianos/osext"
	"github.com/panjf2000/ants"
	"github.com/tidwall/gjson"
)

func StartlistFixed() {
	SetHelper.Rngcmd = make(map[string]int, 1)
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
	now := time.Now()
	if now.String() <= Dalltime {
	} else {
		fmt.Println("License Expired: ", Dalltime)
		os.Exit(1)
	}
	if IPServer == "172.104.76.121" {
		fmt.Println("Welcome,\nyou are subscribed to\nthe license with a server address:", IPServer)
	} else {
		fmt.Println("\033[33m\nYour ip [\033[39m" + IPServer + "\033[33m] is not registered !\n\nPlease Contact LineID :-kont\nor\nhttps://line.me/ti/p/XGUUubUdhu \n\n\033[39m")
		os.Exit(1)
	}
	fmt.Println("\n_Started Login_:")
	go gracefulShutdown()
	StartlistFixed()
	for no, tok := range Data.Authoken {
		time.Sleep(250 * time.Millisecond)
		sort := rand.Intn(9999-1000) + 1000
		app := fmt.Sprintf("ANDROID\t11.6.1\tAndroid OS\t7.1.%v", sort)
		mids := strings.Split(tok, ":")
		mid := mids[0]
		var ua = "Line/11.6.1"

		cl, err := oop.CreateNewLogin(tok, no, mid, app, ua, HostName[rand.Intn(len(HostName))])
		if err == nil {
			cl.Carrier = carierMap["ANDROID"]
			fmt.Println("\n\n  ↳ DisplayName : " + cl.Namebot + "\n  ↳ Mid : " + cl.MID + "\n  ↳ AppName : " + cl.AppName + "\n  ↳ UserAgent : " + cl.UserAgent + "\n  ↳ Bots No: " + fmt.Sprintf("%v", no+1))
			ClientBot = append(ClientBot, cl)
			ClientMid[cl.MID] = cl
			Squadlist = append(Squadlist, cl.MID)
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
	for m := range oop.HashToMap(oop.GetBlock) {
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
		for i := range ClientBot {
			for _, x := range Squadlist {
				if !InArray2(ClientBot[i].Squads, x) && x != ClientBot[i].MID {
					ClientBot[i].Squads = append(ClientBot[i].Squads, x)
				}
			}
		}
		Resprem()
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
	for mid := range memlist {
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
	for mid := range memlist {
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
	for mid := range memlist {
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
	for mid := range memlist {
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
			if oop.IsFriends(client, cc) == false {
				client.FindAndAddContactsByMid(cc)
				time.Sleep(250 * time.Millisecond)
			}
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
	for _, cc := range invits {
		if oop.IsMembers(client, to, cc) == false && oop.IsPending(client, to, cc) == false {
			if !MemBan(to, cc) {
				if oop.IsFriends(client, cc) == false {
					client.FindAndAddContactsByMid(cc)
					time.Sleep(250 * time.Millisecond)
				}
				news = append(news, cc)
			} else {
				bans = append(bans, cc)

			}
		}
	}
	if len(news) != 0 {
		client.InviteIntoChatPollVer(to, news)
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
	for mid := range memlist {
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
				g = []*Stickers{{Id: ids[0], Pid: pids[0]}}
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
	if Commands.Allgaccess != "" {
		list += fmt.Sprintf(" - Allgaccess: %s\n", Commands.Allgaccess)
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
	room := oop.GetRoom(Data.LogGroup)
	if len(room.Client) != 0 {
		exe, err := SelectBot(room.Client[0], Data.LogGroup)
		if err == nil {
			if exe != nil {
				exe.SendMention(Data.LogGroup, ts, []string{from})
			}
		} else {
			LogMode = false
			Data.LogGroup = ""
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
	} else if text == "allgaccess" {
		Commands.Allgaccess = text2
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
	d := fmt.Sprintf("%v", Dalltime)
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
	bod := string(Dalltime)
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
							} else if UserBot.GetBot(user) {
								client.AcceptGroupInvitationNormal(Group)
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
									tempban = invited
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
					}
					LogGet(op)
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
								for mid := range memlist {
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
					}
					LogGet(op)
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
								if InArray2(tempban, user) {
									if MemUser(Group, user) {
										if cekjoin(user) {
											kickPelaku(client, Group, user)
											deljoin(user)
											tempban = Remove(tempban, user)
										}
									} else {
										tempban = Remove(tempban, user)
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
					if MemUser(Group, user) {
						savejoin(user, Optime)
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
					}
					LogGet(op)
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
					}
					LogGet(op)
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
								for mid := range memlist {
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
					}
					LogGet(op)
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
					}
					LogGet(op)
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
					}
					LogGet(op)
				} else if op.Type == 128 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
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
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
					}
					LogGet(op)
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
	for mid := range memlist {
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
	if len(cans) != 0 {
		no := 0
		ah := 0
		if len(mem) > 50 {
			mem = mem[:50]
		}
		for _, target := range mem {
			go func(target string, no int) {
				go cans[no].CancelChatInvitations(to, target)
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
}
func cancelall(client *oop.Account, mem []string, to string) {
	defer panicHandle("cancelall")
	runtime.GOMAXPROCS(cpu)
	Room := oop.GetRoom(to)
	Cans := Room.HaveClient
	if len(Cans) != 0 {
		no := 0
		ah := 0
		if len(mem) > 50 {
			mem = mem[:50]
		}
		for _, target := range mem {
			go func(target string, no int) {
				Cans[no].CancelChatInvitations(to, target)
			}(target, no)
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
	for mid := range memlist {
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
	for mid := range memlist {
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
	for mid := range memlist {
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
	for mid := range memlist {
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
		if con != cl.MID && !cl.Waitadd && !cl.Limitadd {
			_, err := cl.FindAndAddContactsByMid3(con)
			if err != nil {
				println(fmt.Sprintf("%v", err.Error()))
				return 0
			}
			cl.Waitadd = false
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
		room := oop.GetRoom(Data.LogGroup)
		if len(room.Client) != 0 {
			exe, err := SelectBot(room.Client[0], Data.LogGroup)
			if err == nil {
				if exe != nil {
					exe.SendMessage(Data.LogGroup, tx)
				}
			} else {
				LogMode = false
				Data.LogGroup = ""
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
				} else {
					cll = ClientBot[0]
				}
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
				filterWar.clear()
				Nkick = &hashmap.HashMap{}
				filterop = &hashmap.HashMap{}
				oplist = []int64{}
				Ceknuke = &hashmap.HashMap{}
				cekoptime = []int64{}
				PurgeOP = []int64{}
				filtermsg = &hashmap.HashMap{}
				opjoin = []string{}
				Cekpurge = []int64{}
				tempban = []string{}
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
					} else {
						cll = ClientBot[0]
					}
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
				}
			}
		}
	}
	for _, cl := range oop.Waitadd {
		if now.Sub(cl.Timeadd) >= 24*time.Hour {
			if _, ok := oop.GetBlockAdd.Get(cl.MID); !ok {
				cl.Limitadd = false
				cl.Add = 0
				cl.Lastadd = now
				oop.Waitadd = oop.RemoveCl(oop.Waitadd, cl)
				oop.GetBlockAdd.Del(cl.MID)
			}
		}
		if now.Sub(cl.Timeadd) >= 1*time.Hour {
			if _, ok := oop.BlockAdd.Get(cl.MID); !ok {
				cl.Limitadd = false
				cl.Add = 0
				cl.Lastadd = now
				oop.Waitadd = oop.RemoveCl(oop.Waitadd, cl)
				oop.BlockAdd.Del(cl.MID)
			}
		}
	}
	for _, cl := range ClientBot {
		if now.Sub(cl.Lastadd) >= 10*time.Minute {
			cl.Add = 0
			cl.Lastadd = now
		}
		if now.Sub(cl.Lastkick) >= 1*time.Hour {
			cl.TempKick = 0
			cl.TempInv = 0
		}
		if now.Sub(cl.TimeBan) <= 1*time.Second {
			oop.KickBans = oop.RemoveCl(oop.KickBans, cl)
			cl.Limited = false
			cl.TempKick = 0
			cl.TempInv = 0
			cl.Frez = false
			oop.GetBlock.Del(cl.MID)
		}
		if cl.Cpoll >= 10 {
			cl.Cpoll = 0
			//cl.Talk = cl.LoadClient()
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
		tempban = []string{}
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
		for target := range memlists {
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
	for mid := range memlists {
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
	for mid := range memlist {
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
