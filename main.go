package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/bashery/botline/oop"
	"github.com/panjf2000/ants"

	"github.com/bashery/botline/hashmap"
	talkservice "github.com/bashery/linethrift"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("do not forget arguement")
		os.Exit(1)
	}
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

//
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
