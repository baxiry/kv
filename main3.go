package main

import (
	"fmt"
	"runtime"
	"sort"
	"time"

	"github.com/bashery/botline/oop"
	//	"github.com/bashery/botline/oop"
)

func getKey(cmd string) string {
	mp := oop.HashToMap(CmdHelper)
	for k, v := range mp {
		if v.(string) == cmd {
			return k
		}
	}
	return cmd
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

// here func getKey

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
