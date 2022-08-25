package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bashery/botline/oop"
	talkservice "github.com/bashery/linethrift"
)

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
