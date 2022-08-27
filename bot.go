package main

import (
	"botline/oop"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	talkservice "botline/linethrift"

	"github.com/panjf2000/ants"
	"github.com/shirou/gopsutil/mem"
	"github.com/tidwall/gjson"
)

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
	var bk = []*oop.Account{}
	room = oop.GetRoom(to)
	bk = room.Client
	if len(bk) == 0 {
		GetSquad(client, to)
		room = oop.GetRoom(to)
		bk = room.Client
	}
	sort.Slice(room.Ava, func(i, j int) bool {
		return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
	})
	if room.AntiTag && MemUser(to, msg.From_) && len(mentionlist) != 0 && !room.Automute {
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
				rngcmd := GetComs(0, "clearcreator")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(0, "uncreator")
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
				rngcmd := GetComs(1, "seller")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(2, "sellers")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(1, "clearseller")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(1, "unseller")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(7, "expel")
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
							Checklistexpel(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "access" || cmd == Commands.Access && Commands.Access != "" {
				rngcmd := GetComs(3, "access")
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
				rngcmd := GetComs(2, "buyer")
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
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setdate ") {
				if GetCodeprem(0, sender, to) {
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
							Dalltime = times
							str := fmt.Sprintf("⚙️ Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\n⚙️ Remaining:\n  %v", botDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "addweek" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 7 * mont
						t := batas.Add(mont)
						Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addday" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						t := batas.Add(mont)
						Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addyear" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 365 * mont
						t := batas.Add(mont)
						Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addmonth" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 30 * mont
						t := batas.Add(mont)
						Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "reboot" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						client.SendMessage(to, "Rebooting...")
						ReloginProgram()
					}
				}
			} else if strings.HasPrefix(cmd, "unbuyer") {
				rngcmd := GetComs(2, "unbuyer")
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
				rngcmd := GetComs(2, "clearbuyer")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(5, "unsend")
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
			} else if cmd == "Purgeallbans" || cmd == Commands.Purgeall && Commands.Purgeall != "" {
				rngcmd := GetComs(4, "Purgeallbans")
				if GetCodeprem(rngcmd, sender, to) {
					KIckbansPurges(client, to)
				}
			} else if strings.HasPrefix(cmd, "gleave") {
				rngcmd := GetComs(3, "gleave")
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
										for m := range mem {
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
										if Data.LogGroup == gid {
											LogMode = false
											Data.LogGroup = ""
										}
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
				rngcmd := GetComs(3, "invme")
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
				rngcmd := GetComs(3, "decline")
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
				rngcmd := GetComs(3, "accept")
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
											for m := range mem {
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
				rngcmd := GetComs(6, "abort")
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
				rngcmd := GetComs(3, "declineall")
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
			} else if cmd == "stats" {
				rngcmd := GetComs(3, "stats")
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
				rngcmd := GetComs(3, "clearhide")
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
				rngcmd := GetComs(3, "hidelist")
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
				rngcmd := GetComs(3, "unhide")
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
				rngcmd := GetComs(3, "hide")
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
				rngcmd := GetComs(3, "owner")
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
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unowner") {
				rngcmd := GetComs(3, "unowner")
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
				rngcmd := GetComs(3, "clearowner")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "logmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if Data.LogGroup == to {
							LogMode = true
							newsend += "Already enabled.\n"
						} else {
							LogMode = true
							Data.LogGroup = to
							newsend += "Logmode is enabled.\n"
						}
					}
				}
			} else if cmd == "logmode off" {
				rngcmd := GetComs(3, "logmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if Data.LogGroup == to {
							LogMode = false
							Data.LogGroup = ""
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
				rngcmd := GetComs(4, "msgwelcome")
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

			} else if strings.HasPrefix(cmd, "msgclearban ") {
				rngcmd := GetComs(3, "msgclearban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgclearban ", "", 1)
						MsgBan = result
						newsend += "Message clearban set to: " + result + "\n"
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
			} else if strings.HasPrefix(cmd, "msgstatus ") {
				rngcmd := GetComs(3, "msgstatus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgstatus ", "", 1)
						MsFresh = result
						newsend += "Message status set to: " + result + "\n"
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
				rngcmd := GetComs(5, "mute")
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
				rngcmd := GetComs(5, "whois")
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
				rngcmd := GetComs(6, "name")
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
				rngcmd := GetComs(6, "mid")
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
				rngcmd := GetComs(5, "unmute")
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
				rngcmd := GetComs(4, "owners")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(4, "unmaster")
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
				rngcmd := GetComs(5, "ungowner")
				if GetCodeprem(rngcmd, sender, to) {
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
						for m := range mem {
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
				rngcmd := GetComs(4, "master")
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
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gowner") && cmd != "gowners" {
				rngcmd := GetComs(5, "gowner")
				if GetCodeprem(rngcmd, sender, to) {
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
			} else if cmd == "clearlistcmd" {
				rngcmd := GetComs(3, "clearlistcmd")
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
						Commands.Allgaccess = ""
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
						newsend += "Done clearall Cmd.\n"
					}
				}
			} else if cmd == "cleargowner" {
				rngcmd := GetComs(4, "cleargowner")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "clearmaster")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "clearfuck")
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
				rngcmd := GetComs(3, "clearmute")
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
				rngcmd := GetComs(3, "clearallprotect")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						oop.ClearProtect()
						newsend += "Cleared allprotected.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "perm ") {
				rngcmd := GetComs(3, "perm")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ditha := strings.ReplaceAll(cmd, "perm ", "")
						cmdLil := strings.Split(ditha, " ")
						Addpermcmd(client, to, cmdLil[0], cmdLil[1])
					}
				}
			} else if cmd == "permlist" {
				rngcmd := GetComs(3, "permlist")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "autojoin")
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
				rngcmd := GetComs(5, "mutelist")
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
				rngcmd := GetComs(3, "groupcast")
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
				rngcmd := GetComs(3, "gourl")
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
							for mid := range mem {
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
								for mid := range inv {
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
				rngcmd := GetComs(3, "groupinfo")
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
				rngcmd := GetComs(4, "unfuck")
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
				rngcmd := GetComs(5, "admin")
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
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "gowners" {
				rngcmd := GetComs(7, "gowners")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(5, "masters")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(5, "unadmin")
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
				rngcmd := GetComs(7, "ungadmin")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(6, "gadmin")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(6, "admins")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(5, "banlist")
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
				rngcmd := GetComs(6, "respon")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							go p.SendMessage(to, MsgRespon)
						}
					}
				}
			} else if cmd == "rollcall" || cmd == Commands.Botname && Commands.Botname != "" {
				rngcmd := GetComs(5, "rollcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							asss := fmt.Sprintf("%v", p.Namebot)
							go p.SendMessage(to, asss)
						}
					}
				}
			} else if cmd == "upallimage" || cmd == Commands.Upallimage && Commands.Upallimage != "" {
				rngcmd := GetComs(2, "upallimage")
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
				rngcmd := GetComs(2, "upallcover")
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
				rngcmd := GetComs(5, "unsend")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.UnsendChat(to)
						}
					}
				}
			} else if cmd == "upvallimage" || cmd == Commands.Upvallimage && Commands.Upvallimage != "" {
				rngcmd := GetComs(2, "upvallimage")
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
				rngcmd := GetComs(2, "upvallcover")
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
				rngcmd := GetComs(3, "friends")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "adds")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							asss := ""
							if p.Limitadd {
								asss += "⚙️ 𝗕ot 𝗦tatus: ✘"
							} else {
								asss += "⚙️ 𝗕ot 𝗦tatus: ✓"
							}
							p.SendMessage(to, asss)
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
			} else if cmd == "addallbots" || cmd == Commands.Addallbots && Commands.Addallbots != "" {
				rngcmd := GetComs(2, "addallbots")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Bot) != 0 {
							asss := ""
							for _, p := range bk {
								for _, mid := range UserBot.Bot {
									if oop.IsFriends(p, mid) == false {
										time.Sleep(5 * time.Second)
										p.FindAndAddContactsByMid(mid)
									}
								}
							}
							asss += "Success addall bots."
							client.SendMessage(to, asss)
						} else {
							client.SendMessage(to, "Bot list empty.")
						}
					}
				}
			} else if cmd == "addallsquads" || cmd == Commands.Addallsquads && Commands.Addallsquads != "" {
				rngcmd := GetComs(1, "addallsquads")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						asss := ""
						for _, p := range bk {
							for _, mid := range p.Squads {
								if oop.IsFriends(p, mid) == false {
									time.Sleep(1 * time.Second)
									p.FindAndAddContactsByMid(mid)
								}
							}
						}
						asss += "Success addall squads."
						client.SendMessage(to, asss)
					}
				}
			} else if cmd == "leave" || cmd == Commands.Leave && Commands.Leave != "" {
				rngcmd := GetComs(4, "leave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m := range mem {
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
						if Data.LogGroup == to {
							LogMode = false
							Data.LogGroup = ""
						}
						logAccess(client, to, sender, "leave", []string{}, msg.ToType)
					}
				}
			} else if cmd == "ping" || cmd == Commands.Ping && Commands.Ping != "" {
				rngcmd := GetComs(8, "ping")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, "pong")
						}
					}
				}
			} else if cmd == "count" || cmd == Commands.Count && Commands.Count != "" {
				rngcmd := GetComs(6, "count")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for n, p := range bk {
							p.SendMessage(to, fmt.Sprintf("%v", n+1))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "sayall") {
				rngcmd := GetComs(5, "sayall")
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
				rngcmd := GetComs(5, "limitout")
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
				rngcmd := GetComs(2, "upallname")
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
				rngcmd := GetComs(3, "clearadmin")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "clearban")
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
				rngcmd := GetComs(7, "cleargadmin")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(3, "listprotect")
				if GetCodeprem(rngcmd, sender, to) {
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
						exe, _ := SelectBot(client, to)
						if exe != nil {
							Setinviteto(exe, to, exe.Squads)
							logAccess(client, to, sender, "invite", exe.Squads, msg.ToType)
							GetSquad(exe, to)
						} else {
							newsend += "Invite banned try with another bot.\n"
						}
					}
				}
			} else if cmd == "stayall" || cmd == Commands.Stayall && Commands.Stayall != "" {
				rngcmd := GetComs(5, "stayall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
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
					client.SendMessage(to, "Wait...")
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
						Data.LogGroup = ""
						newsend += "Leave done"
					}
				}
			} else if strings.HasPrefix(cmd, "stay ") {
				rngcmd := GetComs(4, "stay")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "stay ", "", 1)
						numb, _ := strconv.Atoi(str)
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
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, Sname)
					}
				}
			} else if pesan == "prefix" {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, "Rname: "+Rname+"\nSname: "+Sname)
					}
				}
			} else if pesan == "rname" {
				if GetCodeprem(8, sender, to) {
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
				rngcmd := GetComs(7, "bot")
				if GetCodeprem(rngcmd, sender, to) {
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
				rngcmd := GetComs(5, "bans")
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
				rngcmd := GetComs(3, "clearbot")
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
				rngcmd := GetComs(4, "cleargban")
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
				rngcmd := GetComs(8, "here")
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
				rngcmd := GetComs(8, "say")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "say ", "", 1)
						client.SendMessage(to, str)
					}
				}
			} else if cmd == "timeleft" {
				rngcmd := GetComs(3, "timeleft")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Dalltime)
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
				rngcmd := GetComs(4, "timenow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GenerateTimeLog(client, to)
					}
				}
			} else if cmd == "runtime" {
				rngcmd := GetComs(4, "runtime")
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
						if Killmode != "none" {
							ret += fmt.Sprintf("\n 🟢 Killmode: %s", Killmode)
						} else {
							ret += "\n 🔴 Killmode"
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
						newsend += ret
					}
				}
			} else if cmd == "settings" {
				rngcmd := GetComs(5, "settings")
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
							if Data.LogGroup == to {
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
							if room.ImageLurk {
								ret += "\n 🟢 sendImage"
							} else {
								ret += "\n 🔴 sendImage"
							}
							if room.Leavebool {
								ret += "\n 🟢 Leave"
							} else {
								ret += "\n 🔴 Leave"
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
				rngcmd := GetComs(4, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk mention" {
				rngcmd := GetComs(4, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "mention"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk on" {
				rngcmd := GetComs(4, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "killmode") {
				rngcmd := GetComs(4, "killmode")
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
				rngcmd := GetComs(6, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.Userlurk = []string{}
						room.NameLurk = "hide"
						newsend += "Lurking...\n"
					}
				}
			} else if cmd == "lurks" {
				rngcmd := GetComs(6, "lurk")
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
				rngcmd := GetComs(6, "lurk")
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
			} else if cmd == "speed" || cmd == Commands.Speed && Commands.Speed != "" {
				rngcmd := GetComs(4, "speed")
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
			} else if cmd == "statusall" {
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
				rngcmd := GetComs(8, "tagall")
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
				rngcmd := GetComs(3, "unbot")
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
				rngcmd := GetComs(7, "ungban")
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
				rngcmd := GetComs(6, "cancel")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
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
				rngcmd := GetComs(6, "invite")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
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
			} else if strings.HasPrefix(cmd, "kick") && cmd != "kickall" || strings.HasPrefix(cmd, Commands.Kick) && Commands.Kick != "" && cmd != "kickall" {
				rngcmd := GetComs(6, "kick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
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
								AutoproN = true
								logAccess(client, to, sender, "kick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, "Please add another bot that has a ban kick.")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ban") && cmd != "bans" {
				rngcmd := GetComs(5, "ban")
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
				rngcmd := GetComs(5, "contact")
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
			} else if strings.HasPrefix(cmd, "unban") {
				rngcmd := GetComs(5, "unban")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(5, "announce")
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
				rngcmd := GetComs(5, "announce")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(6, "deny")
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
				rngcmd := GetComs(6, "allow")
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
				rngcmd := GetComs(5, "leave")
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
				rngcmd := GetComs(5, "welcome")
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
				rngcmd := GetComs(5, "welcome")
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
				rngcmd := GetComs(5, "welcome")
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
				rngcmd := GetComs(3, "backup")
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
				rngcmd := GetComs(3, "backup")
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
				rngcmd := GetComs(5, "hostage")
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
				rngcmd := GetComs(5, "hostage")
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
				rngcmd := GetComs(6, "allowall")
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
				rngcmd := GetComs(6, "denyall")
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
				rngcmd := GetComs(7, "protectmax")
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
				rngcmd := GetComs(7, "protectnone")
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
			} else if strings.HasPrefix(cmd, "joinqr:http://line.me") {
				rngcmd := GetComs(4, "joinqr")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						hyu := strings.Split((text), "http://line.me")
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
				rngcmd := GetComs(3, "joinqrkick")
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
							go Nukjoin(exe[0], op.CreatedTime, tkt.Chat.ChatMid)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "clone") {
				rngcmd := GetComs(3, "clone")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(mentionlist) == 1 {
							cok := strings.Split((cmd), " ")
							if len(cok) > 1 {
								targets := ""
								var pp, vp, co, cv, name, stats string
								cok := strings.Split((cmd), " ")
								if len(cok) > 1 {
									ann := cok[1]
									var prof *talkservice.Contact
									if ann == "@me" {
										prof, _ = client.GetContact(msg.From_)
										targets = msg.From_
									}
									if prof != nil {
										name = prof.DisplayName
										stats = prof.StatusMessage
										if prof.VideoProfile != "" {
											ps, err := client.Downloads("http://dl.profile.line-cdn.net"+prof.PicturePath+"/vp", "mp4")
											if err != nil {
												client.SendMessage(to, "Download video profile error.")
											} else {
												vp = ps
											}
										}
										if prof.PicturePath != "" {
											ps, err := client.Downloads("http://dl.profile.line.naver.jp"+prof.PicturePath, "jpg")
											if err != nil {
												client.SendMessage(to, "Download picture profile error.")
											} else {
												pp = ps
											}
										}
										profs := client.GetProfileDetail(msg.From_)
										pss, err := client.Downloads("https://obs.line-scdn.net/r/myhome/c/"+gjson.Get(profs, "result.objectId").String(), "jpg")
										if err == nil {
											co = pss
										}
										pss, err = client.Downloads("https://obs.line-scdn.net/r/myhome/vc/"+gjson.Get(profs, "result.objectId").String(), "mp4")
										if err == nil {
											cv = pss
										}
										if len(mentionlist) != 0 {
											clon := false
											for _, target := range mentionlist {
												if target != targets && InArray2(Squadlist, target) {
													idx := GetKorban(target)
													clone(idx, pp, vp, co, cv, name, stats)
													idx.SendMention(to, "Cloning @! profile done.", []string{targets})
													clon = true
												}
											}
											if !clon {
												if pp != "" {
													os.Remove(pp)
												}
												if vp != "" {
													os.Remove(vp)
												}
												if co != "" {
													os.Remove(co)
												}
												if cv != "" {
													os.Remove(cv)
												}
											}
										} else {
											if pp != "" {
												os.Remove(pp)
											}
											if vp != "" {
												os.Remove(vp)
											}
											if co != "" {
												os.Remove(co)
											}
											if cv != "" {
												os.Remove(cv)
											}
										}
									}
								}
							}
						}
					}
				}

			}
		}
		if newsend != "" {
			client.SendMessage(to, newsend)
			SaveBackup()
		}
	}
}
