package main

import (
	"botline/oop"
	"fmt"
	"time"
)

func BackSeave() {
	fmt.Println("start Backup Data *__*")
	TimeBackup = time.Time{}
	MsSname = Data.SnameBack
	MsRname = Data.RnameBack
	MsgRespon = Data.ResponBack
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
