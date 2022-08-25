package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	talkservice "github.com/bashery/linethrift"

	"github.com/bashery/botline/oop"
	"github.com/kardianos/osext"
	"github.com/tidwall/gjson"
)

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
