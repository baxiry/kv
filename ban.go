package main

import (
	"botline/oop"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

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
		//return false
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
	//return false
}

func checkip(ip string) bool {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	localip := strings.Split((localAddr).String(), ":")
	IPServer = localip[0]
	if err != nil {
		os.Exit(1)
		fmt.Println(ip)
	} else if string(localip[0]) != ip {
		fmt.Println("\033[33m\nYour ip [\033[39m" + localip[0] + "\033[33m] is not registered !\n\nPlease Contact LineID :  hh7o-\n\n\033[39m")
		os.Exit(1)
		fmt.Println(ip)
	} else if string(localip[0]) == ip {
		fmt.Println(ip)
		return true
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
