package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"

	"github.com/bashery/botline/oop"
)

//
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
