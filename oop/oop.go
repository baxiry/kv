package oop

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"bline/channel"
	"bline/hashmap"
	"bline/modcompact"
	"bline/thrift"

	"github.com/tidwall/gjson"
	"golang.org/x/net/http2"

	talkservice "bline/linethrift"

	"github.com/imroc/req"
)

func HashToMap(mas *hashmap.HashMap) map[string]interface{} {
	ama := map[string]interface{}{}
	aa := mas.Listing() //.Listing()
	for _, ma := range aa {
		ama[ma.Key.(string)] = ma.Value
	}
	return ama
}

type (
	tagdata struct {
		S string `json:"S"`
		E string `json:"E"`
		M string `json:"M"`
	}
	mentions struct {
		MENTIONEES []struct {
			Start string `json:"S"`
			End   string `json:"E"`
			Mid   string `json:"M"`
		} `json:"MENTIONEES"`
	}
	SaveJoin struct {
		User []string
		Time []int64
	}
	Helper struct {
		Rngcmd map[string]int
	}
	ProfileCoverStruct struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Result  struct {
			HomeId       string `json:"homeId"`
			HomeType     string `json:"homeType"`
			HasNewPost   bool   `json:"hasNewPost"`
			CoverObsInfo struct {
				ObsNamespace string `json:"obsNamespace"`
				ServiceName  string `json:"serviceName"`
				ObjectId     string `json:"objectId"`
			} `json:"coverObsInfo"`
			VideoCoverObsInfo struct {
				ObsNamespace string `json:"obsNamespace"`
				ServiceName  string `json:"serviceName"`
				ObjectId     string `json:"objectId"`
			} `json:"videoCoverObsInfo"`
			PostCount         int `json:"postCount"`
			FollowSummaryInfo struct {
				FollowingCount int  `json:"followingCount"`
				FollowerCount  int  `json:"followerCount"`
				Following      bool `json:"following"`
				AllowFollow    bool `json:"allowFollow"`
				ShowFollowList bool `json:"showFollowList"`
			} `json:"followSummaryInfo"`
			GiftShopInfo struct {
				GiftShopScheme         string `json:"giftShopScheme"`
				BirthdayGiftShopScheme string `json:"birthdayGiftShopScheme"`
				GiftShopUrl            string `json:"giftShopUrl"`
				IsGiftShopAvailable    bool   `json:"isGiftShopAvailable"`
			} `json:"giftShopInfo"`
			UserStyleMedia struct {
				MenuInfo struct {
					LatestEditTime int64 `json:"latestEditTime"`
				} `json:"menuInfo"`
				AvatarMenuInfo struct {
					LatestEditTime int64 `json:"latestEditTime"`
				} `json:"avatarMenuInfo"`
			} `json:"userStyleMedia"`
			Meta struct {
			} `json:"meta"`
		} `json:"result"`
	}
	Account struct {
		AuthToken     string
		AppName       string
		UserAgent     string
		Host          string
		MID           string
		Shnall        string
		Limited       bool
		Frez          bool
		Limitadd      bool
		Waitadd       bool
		Seq           int32
		Carrier       string
		Akick         int
		Ainvite       int
		CustomPoint   int
		KickCount     int
		CancelCount   int
		InvCount      int
		Curtime       int64
		TempKick      int
		TempInv       int
		KickPoint     int
		Add           int
		Acancel       int
		Namebot       string
		IpProxy       string
		Numar         int
		Ckick         int
		SessionPoll   *thrift.THttpClient
		Transport     *http.Transport
		Timeadd       time.Time
		TimeBan       time.Time
		Lastkick      time.Time
		Lastinvite    time.Time
		Lastcancel    time.Time
		Lastadd       time.Time
		CountDay      int
		Locale        string
		HttpHeader    http.Header
		hc            *http.Client
		UrS4          string
		UrP5          string
		Cinvite       int
		Ccancel       int
		SHani         int
		Count         int32
		Revision      int64
		GRevision     int64
		Ctx           context.Context
		reqSeqMessage int32
		IRevision     int64
		Cpoll         int
		Squads        []string
		Backup        []string
	}
	BanUser struct {
		Banlist  []string
		Fucklist []string
		Mutelist []string
	}
	Access struct {
		Creator []string
		Seller  []string
		Buyer   []string
		Owner   []string
		Master  []string
		Admin   []string
		Bot     []string
	}
	LineRoom struct {
		Name        string
		Id          string
		Lurk        bool
		Announce    bool
		Userlurk    []string
		NameLurk    string
		Leavebool   bool
		Backleave   bool
		MsgLeave    string
		MsgLurk     string
		ProKick     bool
		ProQr       bool
		ProName     bool
		ImageLurk   bool
		ProInvite   bool
		ProJoin     bool
		ProCancel   bool
		Limit       bool
		Welcome     bool
		WelcomeMsg  string
		AntiTag     bool
		Automute    bool
		LeaveBack   []string
		Gowner      []string
		ListInvited []string
		Gadmin      []string
		Gban        []string
		Exe         *Account
		Bot         []string
		GoMid       []string
		Client      []*Account
		Ava         []*Ava
		GoClient    []*Account
		HaveClient  []*Account
		Invite      int
		Kick        int
		Cancel      int
		Fight       time.Time
		Leave       time.Time
		Backup      bool
		Qr          bool
		Purge       bool
	}
	Ava struct {
		Client *Account
		Exist  bool
		Mid    string
	}
	Command struct {
		Botname      string
		Upallimage   string
		Upallcover   string
		Unsend       string
		Upvallimage  string
		Upvallcover  string
		Appname      string
		Useragent    string
		Hostname     string
		Friends      string
		Adds         string
		Limits       string
		Addallbots   string
		Addallsquads string
		Leave        string
		Respon       string
		Ping         string
		Count        string
		Limitout     string
		Access       string
		Allbanlist   string
		Allgaccess   string
		Gaccess      string
		Checkram     string
		Backups      string
		Upimage      string
		Upcover      string
		Upvimage     string
		Upvcover     string
		Bringall     string
		Purgeall     string
		Banlist      string
		Clearban     string
		Stayall      string
		Clearchat    string
		Here         string
		Speed        string
		Status       string
		Tagall       string
		Kick         string
		Max          string
		None         string
		Kickall      string
		Cancelall    string
	}
)

var (
	Random      bool
	Waitadd     = []*Account{}
	BlockAdd    = &hashmap.HashMap{}
	CoverVideo  = ""
	PathProxy   = ""
	GetBlock    = &hashmap.HashMap{}
	GetBlockAdd = &hashmap.HashMap{}
	limiterBot  = map[string]time.Time{}
	KickBans    = []*Account{}
	SquadRoom   = []*LineRoom{}
	Room        = &LineRoom{}
	err         error
	AccessUser  = []*Access{}
	ListProxy   = []string{}
)

func Loadproxy(base string) {
	fn := base + ".json"
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	asu := strings.Join(txtlines, "\n")
	getcan := gjson.Get(asu, "proxy")
	if getcan.Exists() {
		for _, a := range getcan.Array() {
			ListProxy = append(ListProxy, a.String())
		}
	}
}

//---------Access----------------
func (self *Access) AddCreator(user string) {
	defer PanicOnly()
	if !InArray(self.Creator, user) {
		self.Creator = append(self.Creator, user)
	}
}
func (self *Access) DelCreator(user string) {
	defer PanicOnly()
	if InArray(self.Creator, user) {
		self.Creator = Remove(self.Creator, user)
	}
}
func (self *Access) GetCreator(user string) bool {
	defer PanicOnly()
	if InArray(self.Creator, user) {
		return true
	}
	return false
}
func (self *Access) ClearCreator() {
	defer PanicOnly()
	self.Creator = []string{}
}

//------------------------
func (self *Access) AddSeller(user string) {
	defer PanicOnly()
	if !InArray(self.Seller, user) {
		self.Seller = append(self.Seller, user)
	}
}
func (self *Access) DelSeller(user string) {
	defer PanicOnly()
	if InArray(self.Seller, user) {
		self.Seller = Remove(self.Seller, user)
	}
}
func (self *Access) GetSeller(user string) bool {
	defer PanicOnly()
	if InArray(self.Seller, user) {
		return true
	}
	return false
}
func (self *Access) ClearSeller() {
	defer PanicOnly()
	self.Seller = []string{}
}

//------------------------
func (self *Access) AddBuyer(user string) {
	defer PanicOnly()
	if !InArray(self.Buyer, user) {
		self.Buyer = append(self.Buyer, user)
	}
}
func (self *Access) DelBuyer(user string) {
	defer PanicOnly()
	if InArray(self.Buyer, user) {
		self.Buyer = Remove(self.Buyer, user)
	}
}
func (self *Access) GetBuyer(user string) bool {
	defer PanicOnly()
	if InArray(self.Buyer, user) {
		return true
	}
	return false
}
func (self *Access) ClearBuyer() {
	defer PanicOnly()
	self.Buyer = []string{}
}

//------------------------
func (self *Access) AddOwner(user string) {
	defer PanicOnly()
	if !InArray(self.Owner, user) {
		self.Owner = append(self.Owner, user)
	}
}
func (self *Access) DelOwner(user string) {
	defer PanicOnly()
	if InArray(self.Owner, user) {
		self.Owner = Remove(self.Owner, user)
	}
}
func (self *Access) GetOwner(user string) bool {
	defer PanicOnly()
	if InArray(self.Owner, user) {
		return true
	}
	return false
}
func (self *Access) ClearOwner() {
	defer PanicOnly()
	self.Owner = []string{}
}

//------------------------
func (self *Access) AddMaster(user string) {
	defer PanicOnly()
	if !InArray(self.Master, user) {
		self.Master = append(self.Master, user)
	}
}
func (self *Access) DelMaster(user string) {
	defer PanicOnly()
	if InArray(self.Master, user) {
		self.Master = Remove(self.Master, user)
	}
}
func (self *Access) GetMaster(user string) bool {
	defer PanicOnly()
	if InArray(self.Master, user) {
		return true
	}
	return false
}
func (self *Access) ClearMaster() {
	defer PanicOnly()
	self.Master = []string{}
}

//------------------------
func (self *Access) AddAdmin(user string) {
	defer PanicOnly()
	if !InArray(self.Admin, user) {
		self.Admin = append(self.Admin, user)
	}
}
func (self *Access) DelAdmin(user string) {
	defer PanicOnly()
	if InArray(self.Admin, user) {
		self.Admin = Remove(self.Admin, user)
	}
}
func (self *Access) GetAdmin(user string) bool {
	defer PanicOnly()
	if InArray(self.Admin, user) {
		return true
	}
	return false
}
func (self *Access) ClearAdmin() {
	defer PanicOnly()
	self.Admin = []string{}
}

//------------------------
func (self *Access) AddBot(user string) {
	defer PanicOnly()
	if !InArray(self.Bot, user) {
		self.Bot = append(self.Bot, user)
	}
}
func (self *Access) DelBot(user string) {
	defer PanicOnly()
	if InArray(self.Bot, user) {
		self.Bot = Remove(self.Bot, user)
	}
}
func (self *Access) GetBot(user string) bool {
	defer PanicOnly()
	if InArray(self.Bot, user) {
		return true
	}
	return false
}
func (self *Access) ClearBot() {
	defer PanicOnly()
	self.Bot = []string{}
}

//------------------------
func (self *BanUser) AddBan(user string) {
	defer PanicOnly()
	if !InArray(self.Banlist, user) && user != "" {
		self.Banlist = append(self.Banlist, user)
	}
}
func (self *BanUser) DelBan(user string) {
	defer PanicOnly()
	if InArray(self.Banlist, user) {
		self.Banlist = Remove(self.Banlist, user)
	}
}
func (self *BanUser) GetBan(user string) bool {
	defer PanicOnly()
	if InArray(self.Banlist, user) {
		return true
	}
	return false
}

func (self *BanUser) AddFuck(user string) {
	defer PanicOnly()
	if !InArray(self.Fucklist, user) && user != "" {
		self.Fucklist = append(self.Fucklist, user)
	}
}
func (self *BanUser) DelFuck(user string) {
	defer PanicOnly()
	if InArray(self.Fucklist, user) {
		self.Fucklist = Remove(self.Fucklist, user)
	}
}
func (self *BanUser) GetFuck(user string) bool {
	defer PanicOnly()
	if InArray(self.Fucklist, user) {
		return true
	}
	return false
}
func (self *BanUser) AddMute(user string) {
	defer PanicOnly()
	if !InArray(self.Mutelist, user) && user != "" {
		self.Mutelist = append(self.Mutelist, user)
	}
}
func (self *BanUser) DelMute(user string) {
	defer PanicOnly()
	if InArray(self.Mutelist, user) {
		self.Mutelist = Remove(self.Mutelist, user)
	}
}
func (self *BanUser) GetMute(user string) bool {
	defer PanicOnly()
	if InArray(self.Mutelist, user) {
		return true
	}
	return false
}
func (cl *Account) SendImageWithURL(to string, url string) {
	res, err := http.Get("https://minz-restapi.xyz/sendimgurl?authtoken=" + cl.AuthToken + "&to=" + to + "&url=" + url)
	if err != nil {
		fmt.Println("Failed")
		return
	}
	defer res.Body.Close()
}

func (s *Account) CountKick() {
	var asu int
	var cokss int
	cokss = s.SHani + 1
	asu = s.Ckick + 1
	s.Ckick = asu
	s.SHani = cokss
}
func (s *Account) CCancel() {
	var asu int
	asu = s.Ccancel + 1
	s.Ccancel = asu
}
func (s *Account) CInvite() {
	var asu int
	asu = s.Cinvite + 1
	s.Cinvite = asu
}
func InArray(ArrList []string, rstr string) bool {
	for _, x := range ArrList {
		if x == rstr {
			return true
		}
	}
	return false
}

func InArrayInt(arr []int, str int) bool {
	for _, tar := range arr {
		if tar == str {
			return true
		}
	}
	return false
}

func CheckEqual(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if InArray(list2, v) {
			return true
		}
	}
	return false
}

func Randint(min int, max int) int {
	return rand.Intn(max-min) + min
}

func IsMember(members map[string]int64, mid string) bool {
	for x := range members {
		if x == mid {
			return true
		}
	}
	return false
}
func Checkmulti(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if InArray(list2, v) {
			return true
		}
	}
	return false
}
func IndexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
func (self *Account) Getcontactuser(id string) (err error) {
	defer PanicOnly()
	client := self.Talk()
	_, err = client.GetContact(self.Ctx, id)
	if err != nil {
		return err
	}
	return err
}
func (self *Account) GetGroup(groupId string) (r *talkservice.Group) {
	defer PanicOnly()
	client := self.Talk()
	r, _ = client.GetGroup(self.Ctx, groupId)
	return r
}

func (client *Account) SendPollMention(to string, jenis string, memlist []string) {
	defer PanicOnly()
	ta := false
	tx := ""
	tag := []string{}
	z := len(memlist) / 20
	y := z + 1
	for i := 0; i < y; i++ {
		if !ta {
			tx += fmt.Sprintf("%s\n", jenis)
			ta = true
		}
		if i == z {
			tag = memlist[i*20:]
			no := i * 20
			no += 1
			for i := 0; i < len(tag); i++ {
				iki := no + i
				tx += fmt.Sprintf("\n%v. @!", iki)
			}
		} else {
			tag = memlist[i*20 : (i+1)*20]
			no := i * 20
			no += 1
			for i := 0; i < len(tag); i++ {
				iki := no + i
				if iki < 10 {
					tx += fmt.Sprintf("\n%v.  @!", iki)
				} else {
					tx += fmt.Sprintf("\n%v. @!", iki)
				}

			}
		}
		if len(tag) != 0 {
			client.SendMention(to, tx, tag)
		}
		tx = ""
	}
}
func (self *Account) LeaveGroup(groupId string) (err error) {
	req := &talkservice.DeleteSelfFromChatRequest{
		ChatMid:                      groupId,
		ReqSeq:                       self.Seq,
		LastSeenMessageDeliveredTime: 0,
		LastSeenMessageId:            "",
		LastMessageDeliveredTime:     0,
		LastMessageId:                "",
	}
	_, err = self.Talk().DeleteSelfFromChat(context.TODO(), req)
	self.Seq++
	return err
}

func (s *Account) UnsendChatnume(toId string, text string) (err error) {
	client := s.Talk()
	err = client.UnsendMessage(s.Ctx, int32(0), text)
	return err
}
func (s *Account) RemoveAllMessage(lastMessageId string) {
	client := s.Talk()
	client.RemoveAllMessages(s.Ctx, s.Seq, lastMessageId)
}
func InArrayCl(List []*Account, self *Account) bool {
	for _, x := range List {
		if x == self {
			return true
		}
	}
	return false
}
func CheckErr(self *Account, e error, s string, t string) int {
	val := GetCode(e)
	if val == 35 {
		if !InArrayCl(KickBans, self) {
			KickBans = append(KickBans, self)
			self.TimeBan = time.Now()
		}
		self.Limited = true
		if _, ok := GetBlock.Get(self.MID); !ok {
			GetBlock.Set(self.MID, time.Now())
		}
	} else if val == 20 {
		if strings.Contains(e.Error(), "suspended") {
			if !InArrayCl(KickBans, self) {
				KickBans = append(KickBans, self)
				self.TimeBan = time.Now()
			}
			self.Limited = true
			self.Frez = true
		}
	}
	return val
}

func GetCode(e error) int {
	jos := e.Error()
	if strings.Contains(jos, "ABUSE_BLOCK") {
		return 35
	} else if strings.Contains(jos, "INTERNAL_ERROR") || strings.Contains(jos, "suspended") {
		return 20
	} else {
		return 0
	}
}

func (s *Account) UpdateProfilePicture(path, tipe string) error {
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	dataa := ""
	nama := filepath.Base(path)
	if tipe == "vp" {
		dataa = fmt.Sprintf(`{"name": "%s", "oid": "%s", "type": "image", "ver": "2.0", "cat": "vp.mp4"}`, nama, s.MID)
	} else {
		dataa = fmt.Sprintf(`{"name": "%s", "oid": "%s", "type": "image", "ver": "2.0"}`, nama, s.MID)
	}
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	cl := s.hc
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/talk/p/upload.nhn", bytes.NewBuffer(bytess))
	for k, v := range map[string]string{
		"User-Agent":         s.UserAgent,
		"X-Line-Application": s.AppName,
		"X-Line-Access":      s.AuthToken,
		"x-lal":              s.Locale,
		"x-lpv":              "1",
	} {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))
	res, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
func (p *Account) DownloadObjectMsg(msgid string, t ...string) (string, error) {
	var tipe string
	if len(t) == 0 {
		tipe = "jpg"
	} else {
		tipe = t[0]
	}

	client := p.hc
	req, _ := http.NewRequest("GET", "https://obs-sg.line-apps.com/talk/m/download.nhn?oid="+msgid, nil)
	req.Header.Set("User-Agent", p.UserAgent)
	req.Header.Set("X-Line-Application", p.AppName)
	req.Header.Set("X-Line-Access", p.AuthToken)
	res, err := client.Do(req)
	if err != nil {
		println("[DownloadObjectMsg]", err)
	}
	defer res.Body.Close()
	file, err := os.Create("download/" + msgid + "-dl." + tipe)
	if err != nil {
		return "", err
	}
	io.Copy(file, res.Body)
	file.Close()
	return file.Name(), nil
}
func (self *Account) FindAndAddContactsByMidnew2(mid string) (r map[string]*talkservice.Contact, err error) {
	if self.MID == mid || self.Limitadd {
		return map[string]*talkservice.Contact{}, nil
	} else if self.Frez {
		return map[string]*talkservice.Contact{}, nil
	}
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	res, err := TS.FindAndAddContactsByMid(self.Ctx, self.Seq, mid, 0, "")
	if err != nil {
		e := GetCode(err)
		if e == 35 {
			if !InArrayCl(Waitadd, self) {
				Waitadd = append(Waitadd, self)
				self.Timeadd = time.Now()
			}
			self.Limitadd = true
			if _, ok := BlockAdd.Get(self.MID); !ok {
				BlockAdd.Set(self.MID, time.Now())
			}
		}
	}
	return res, err
}
func (self *Account) FindAndAddContactsByMidnew(mid string) (r map[string]*talkservice.Contact, err error) {
	if self.MID == mid || self.Limitadd {
		return map[string]*talkservice.Contact{}, nil
	} else if self.Frez {
		return map[string]*talkservice.Contact{}, nil
	}
	var TS *talkservice.TalkServiceClient
	AddCount(self, "add")
	TS = self.Talk()
	res, err := TS.FindAndAddContactsByMid(self.Ctx, self.Seq, mid, 0, "")
	if err != nil {
		e := GetCode(err)
		if e == 35 {
			if !InArrayCl(Waitadd, self) {
				Waitadd = append(Waitadd, self)
				self.Timeadd = time.Now()
			}
			self.Limitadd = true
			if _, ok := BlockAdd.Get(self.MID); !ok {
				BlockAdd.Set(self.MID, time.Now())
			}
		}
	}
	return res, err
}
func AddCount(user *Account, t string) {
	if t == "kick" {
		user.KickCount += 1
		user.TempKick += 1
		user.KickPoint += 1
		user.CountDay += 1
		user.Lastkick = time.Now()
		if user.TempKick >= 40 || user.CountDay >= 150 {
			if !InArrayCl(KickBans, user) {
				KickBans = append(KickBans, user)
				user.TimeBan = time.Now()
			}
			user.Limited = true
		}
	} else if t == "c" {
		user.CancelCount += 1
		user.Lastcancel = time.Now()
	} else if t == "add" {
		user.Add += 1
		user.Lastadd = time.Now()
		if user.Add >= 10 {
			if !InArrayCl(Waitadd, user) {
				Waitadd = append(Waitadd, user)
				user.Timeadd = time.Now()
			}
			user.Limitadd = true
		}
	} else {
		user.InvCount += 1
		user.TempInv += 1
		user.KickPoint += 1
		user.Lastinvite = time.Now()
		if user.TempInv >= 40 {
			if !InArrayCl(KickBans, user) {
				KickBans = append(KickBans, user)
				user.TimeBan = time.Now()
			}
			user.Limited = true
		}
	}
	if Random {
		Random = false
	} else {
		Random = true
	}
}
func (s *Account) UpdateProfilePictureVideo(pict, vid string) error {
	fl, err := os.Open(vid)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	dataa := fmt.Sprintf(`{"name": "%s", "oid": "%s", "ver": "2.0", "type": "video", "cat": "vp.mp4"}`, filepath.Base(vid), s.MID)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	cl := s.hc
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/talk/vp/upload.nhn", bytes.NewBuffer(bytess))
	for k, v := range map[string]string{
		"User-Agent":         s.UserAgent,
		"X-Line-Application": s.AppName,
		"X-Line-Access":      s.AuthToken,
		"x-lal":              s.Locale,
		"x-lpv":              "1",
	} {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))
	res, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return s.UpdateProfilePicture(pict, "vp")
}
func (s *Account) UnsendChat(toId string) (err error) {
	client := s.Talk()
	Nganu, _ := client.GetRecentMessagesV2(s.Ctx, toId, int32(100000000))
	Mid := []string{}
	for _, chat := range Nganu {
		if chat.From_ == s.MID {
			Mid = append(Mid, chat.ID)
		}
	}
	for i := 0; i < len(Mid); i++ {
		err = client.UnsendMessage(s.Ctx, int32(0), Mid[i])
	}
	return err
}
func (cl *Account) TimeLineGet(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	cl.DefaultTimelineHeader(req)
	client := cl.hc
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	cvt, _ := ioutil.ReadAll(resp.Body)
	return string(cvt), err
}
func (cl *Account) GetProfileDetail(mid string) string {
	url := AddParam("https://"+cl.Host+".line.naver.jp/mh/api/v45/post/list.json?", map[string]string{
		"userMid": mid,
	})
	tr, _ := cl.TimeLineGet(url)
	return tr
}
func AddParam(urls string, param map[string]string) string {
	p := url.Values{}
	for k, v := range param {
		p.Add(k, v)
	}
	return urls + p.Encode()
}
func (s *Account) DownloadFileURL(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	y := s.hc
	res, err := y.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	var tp string
	if strings.Contains(res.Header.Get("Content-Type"), "image") {
		tp = "jpg"
	} else if strings.Contains(res.Header.Get("Content-Type"), "video") {
		tp = "mp4"
	} else if strings.Contains(res.Header.Get("Content-Type"), "audio") {
		tp = "mp3"
	} else {
		tp = "bin"
	}
	tmpfile, err := ioutil.TempFile("/tmp", "DL-*."+tp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tmpfile.Close()
	if _, err := io.Copy(tmpfile, res.Body); err != nil {
		fmt.Println(err)
		return "", err
	}
	return tmpfile.Name(), nil
}

func (s *Account) ChangeProfileVideo(to string, msgid string) {
	prof, _ := s.GetProfile()
	path_p, _ := s.DownloadFileURL("https://obs.line-scdn.net/" + prof.PictureStatus)
	if _, err := os.Stat(path_p); os.IsNotExist(err) {
		s.SendMessage(to, "Update profile error.")
		return
	}
	path, _ := s.DownloadObjectMsg(msgid, "bin")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		s.SendMessage(to, "Update profile error.")
		return
	}
	_ = s.UpdateProfilePictureVideo(path_p, path)
	s.SendMessage(to, "Success update profile video.")
}
func (s *Account) UpdateProfileCoverById(objId string, coverVideo bool) error {
	chtoken, _ := s.ApproveChannelAndIssueChannelToken("1341209850")
	header := make(http.Header)
	for k, v := range map[string]string{
		"Content-Type":              "application/json; charset=UTF-8",
		"User-Agent":                s.UserAgent,
		"X-Line-Mid":                s.MID,
		"X-Line-Access":             s.AuthToken,
		"X-Line-Application":        s.AppName,
		"X-Line-ChannelToken":       chtoken.ChannelAccessToken,
		"x-lal":                     s.Locale,
		"x-lpv":                     "1",
		"x-lsr":                     "JP",
		"x-line-bdbtemplateversion": "v1",
		"x-line-global-config":      "discover.enable=true; follow.enable=true",
	} {
		header.Set(k, v)
	}
	data := map[string]string{
		"homeId":        s.MID,
		"coverObjectId": objId,
		"storyShare":    "false",
	}
	if coverVideo == true {
		data["videoCoverObjectId"] = objId
	}
	_, err := req.Post("https://"+s.Host+".line.naver.jp/hm/api/v1/home/cover.json", header, req.BodyJSON(data))
	return err
}
func genObsParam(dict map[string]string) string {
	marshal, _ := json.Marshal(dict)
	return b64.StdEncoding.EncodeToString(marshal)
}
func (s *Account) UploadObjHome(path, tipe, objId string) (string, error) {
	chtoken, _ := s.ApproveChannelAndIssueChannelToken("1341209850")
	header := make(http.Header)
	for k, v := range map[string]string{
		"Content-Type":              "application/json; charset=UTF-8",
		"User-Agent":                s.UserAgent,
		"X-Line-Mid":                s.MID,
		"X-Line-Access":             s.AuthToken,
		"X-Line-Application":        s.AppName,
		"X-Line-ChannelToken":       chtoken.ChannelAccessToken,
		"x-lal":                     s.Locale,
		"x-lpv":                     "1",
		"x-lsr":                     "JP",
		"x-line-bdbtemplateversion": "v1",
		"x-line-global-config":      "discover.enable=true; follow.enable=true",
	} {
		header.Set(k, v)
	}
	var ctype string
	var endpoint string
	if tipe == "image" {
		ctype = "image/jpeg"
		endpoint = "/r/myhome/c/"
	} else {
		ctype = "video/mp4"
		endpoint = "/r/myhome/vc/"
	}
	if objId == "objid" {
		hstr := fmt.Sprintf("Line_%d", time.Now().Unix())
		objId = fmt.Sprintf("%x", md5.Sum([]byte(hstr)))
	}
	file, _ := os.Open(path)
	fi, err := file.Stat()
	if err != nil {
		return "", err
	}
	for k, v := range map[string]string{
		"x-obs-params": genObsParam(map[string]string{
			"name":   fmt.Sprintf("%d", time.Now().Unix()),
			"userid": s.MID,
			"oid":    objId,
			"type":   tipe,
			"ver":    "1.0",
		}),
		"Content-Type":   ctype,
		"Content-Length": fmt.Sprintf("%d", fi.Size()),
	} {
		header.Set(k, v)
	}
	_, err = req.Post("https://obs-sg.line-apps.com"+endpoint+objId, file, header)
	if err != nil {
		return "", err
	}
	return objId, nil
}
func (s *Account) ChangeCoverPicture(to string, msgid string) {
	path, _ := s.DownloadObjectMsg(msgid, "bin")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		s.SendMessage(to, "Update cover error.")
		return
	}
	oid, err := s.UploadObjHome(path, "image", "objid")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = s.UpdateProfileCoverById(oid, false)
	s.SendMessage(to, "Success update cover.")
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
func (cl *Account) DefaultLineHeader(req *http.Request) {
	req.Header.Set("User-Agent", cl.UserAgent)
	req.Header.Set("X-Line-Application", cl.AppName)
	req.Header.Set("X-Line-Access", cl.AuthToken)
	req.Header.Set("X-Line-Carrier", "51010")
	req.Header.Set("x-lal", cl.Locale)
}
func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	st := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = st[rand.Intn(len(st))]
	}
	return string(b)
}
func hed(r *http.Request, heder map[string]string) {
	for k, v := range heder {
		r.Header.Set(k, v)
	}
}
func (cl *Account) ApproveChannelAndIssueChannelToken(chanid string) (*channel.ChannelToken, error) {
	return cl.LoadChannel().ApproveChannelAndIssueChannelToken(cl.Ctx, chanid)
}
func (cl *Account) DefaultTimelineHeader(req *http.Request) {
	chtoken, _ := cl.ApproveChannelAndIssueChannelToken("1341209850")
	mp := map[string]string{
		`Content-Type`:        `application/json`,
		`User-Agent`:          cl.UserAgent,
		`X-Line-Mid`:          cl.MID,
		`X-Line-Carrier`:      "51010",
		`X-Line-Application`:  cl.AppName,
		`X-Line-ChannelToken`: chtoken.ChannelAccessToken,
	}
	hed(req, mp)
}

func (cl *Account) UpdateCoverById(objId, tipe string) error {
	defer PanicOnly()
	var js []byte
	if tipe == "p" {
		js, _ = json.Marshal(map[string]interface{}{
			"homeId":        cl.MID,
			"coverObjectId": objId,
			"storyShare":    false,
			"meta":          map[string]string{},
		})
	} else if tipe == "v" {
		js, _ = json.Marshal(map[string]interface{}{
			"homeId":             cl.MID,
			"coverObjectId":      objId,
			"storyShare":         false,
			"meta":               map[string]string{},
			"videoCoverObjectId": CoverVideo,
		})
		//dataa = fmt.Sprintf(`{"homeId": %s, "coverObjectId": %s, "storyShare": "false", "meta":{}, :"%s", "videoCoverObjectId":%s%}`,cl.Mid,objId,objId)
	} else {
		js, _ = json.Marshal(map[string]interface{}{
			"homeId":        tipe,
			"coverObjectId": objId,
			"storyShare":    false,
			"meta":          map[string]string{},
		})
		req, err := http.NewRequest("POST", "https://"+cl.Host+".line.naver.jp/hm/api/v1/home/groupprofile/defaultimages.json", bytes.NewBuffer(js))
		if err != nil {
			return err
		}
		client := cl.hc
		cl.DefaultTimelineHeader(req)
		for {
			res, err := client.Do(req)
			if err != nil {
				ef := err.Error()
				if strings.Contains(ef, "EOF") {
					continue
				} else {
					return err
				}
			} else {
				defer res.Body.Close()
				return nil
			}
		}
		//return nil
	}

	req, err := http.NewRequest("POST", "https://"+cl.Host+".line.naver.jp/hm/api/v1/home/cover.json", bytes.NewBuffer(js))
	if err != nil {
		return err
	}

	cl.DefaultTimelineHeader(req)

	resp, err := cl.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func SetHeader(z *http.Request) {
	z.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	z.Header.Set("authority", "scrapeme.live")
	z.Header.Set("upgrade-insecure-requests", "1")
	z.Header.Set("dnt", "1")
	z.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	z.Header.Set("sec-fetch-mode", "navigate")
	z.Header.Set("sec-fetch-user", "?1")
	z.Header.Set("sec-fetch-dest", "document")
	z.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
}
func (p *Account) Downloads(url string, tp string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	SetHeader(req)
	y := p.hc
	res, err := y.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if strings.Contains(res.Header.Get("Content-Type"), "image") {
		tp = "jpg"
	} else if strings.Contains(res.Header.Get("Content-Type"), "video") {
		tp = "mp4"
	} else if strings.Contains(res.Header.Get("Content-Type"), "audio") {
		tp = "mp3"
	}
	tmpfile, err := ioutil.TempFile("download", "DL-*."+tp)
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()
	if _, err := io.Copy(tmpfile, res.Body); err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}
func (y *Account) UpdateCover(path string) error {
	objId := RandomString(32)
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	nama := filepath.Base(path)
	dataa := fmt.Sprintf(`
    {
      "name": "%s",
      "oid": "%s",
      "type": "image",
      "userid": "%s",
      "ver": "2.0"
    }`, nama, objId, y.MID)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))

	client := y.hc
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/r/myhome/c/"+objId, bytes.NewBuffer(bytess))
	y.DefaultTimelineHeader(req)
	req.Header.Set("x-obs-params", string(sDec))
	req.Header.Set("X-Line-PostShare", "false")
	req.Header.Set("X-Line-StoryShare", "false")
	req.Header.Set("x-line-signup-region", "ID")
	req.Header.Set("content-type", "image/png")
	req.ContentLength = int64(len(bytess))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return y.UpdateCoverById(objId, "p")
}
func (y *Account) UpdatePictureProfile(path, tipe string) error {
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	tim := fmt.Sprintf("%s%d", filepath.Base(path), time.Now().UnixNano()/1000)
	nama := GetMD5Hash(tim)

	dataa := fmt.Sprintf(`{"name": "%s", "quality": "100", "type": "image"`, nama)
	if tipe == "v" {
		dataa += `, "ver": "2.0", "cat": "vp.mp4"}`
	} else {
		dataa += `, "ver": "2.0"}`
	}
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	client := y.hc
	req, _ := http.NewRequest("POST", "https://obs.line-apps.com/r/talk/p/"+y.MID, bytes.NewBuffer(bytess))
	y.DefaultLineHeader(req)
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))
	for {
		res, err := client.Do(req)
		if err != nil {
			ef := err.Error()
			if strings.Contains(ef, "EOF") {
				continue
			} else {
				return err
			}
		} else {
			res.Body.Close()
			return nil
		}
	}
	//return nil
}
func (y *Account) UpdateCoverVideo(path string) error {
	objId := RandomString(32)
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	nama := filepath.Base(path)
	dataa := fmt.Sprintf(`
    {
      "name": "%s",
      "oid": "%s",
      "type": "video",
      "userid": "%s",
      "ver": "2.0"
    }`, nama, objId, y.MID)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	client := y.hc
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/r/myhome/vc/"+objId, bytes.NewBuffer(bytess))
	y.DefaultTimelineHeader(req)
	req.Header.Set("x-obs-params", string(sDec))
	req.Header.Set("X-Line-PostShare", "false")
	req.Header.Set("X-Line-StoryShare", "false")
	req.Header.Set("x-line-signup-region", "ID")
	req.Header.Set("content-type", "video/mp4")
	req.ContentLength = int64(len(bytess))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	CoverVideo = objId
	return nil
}
func (y *Account) UpdateCoverWithVideo(path string) error {
	objId := RandomString(32)
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	nama := filepath.Base(path)
	dataa := fmt.Sprintf(`
    {
      "name": "%s",
      "oid": "%s",
      "type": "image",
      "userid": "%s",
      "ver": "2.0"
    }`, nama, objId, y.MID)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))

	client := y.hc
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/r/myhome/c/"+objId, bytes.NewBuffer(bytess))
	y.DefaultTimelineHeader(req)
	req.Header.Set("x-obs-params", string(sDec))
	req.Header.Set("X-Line-PostShare", "false")
	req.Header.Set("X-Line-StoryShare", "false")
	req.Header.Set("x-line-signup-region", "ID")
	req.Header.Set("content-type", "image/png")
	req.ContentLength = int64(len(bytess))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return y.UpdateCoverById(objId, "v")
}
func (y *Account) UpdateVideoProfile(vid string) error {
	fl, err := os.Open(vid)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}

	tim := fmt.Sprintf("%s%d", filepath.Base(vid), time.Now().UnixNano()/1000)
	nama := GetMD5Hash(tim)

	dataa := fmt.Sprintf(`{"name": "%s", "ver": "2.0", "type": "video", "cat": "vp.mp4"}`, nama)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))

	client := y.hc
	req, _ := http.NewRequest("POST", "https://obs.line-apps.com/r/talk/vp/"+y.MID, bytes.NewBuffer(bytess))
	y.DefaultLineHeader(req)
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
func Clearcache() {
	fmt.Println("CACHE_CLEARED")
	exec.Command("sync;", "echo", "1", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "2", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "3", ">", "/proc/sys/vm/drop_caches").Run()
}
func setHC() *http.Client {
	return &http.Client{
		Transport: &http.Transport{},
	}
}

func removeEndNewLine(input string) string {
	defer PanicOnly()
	if len(input) == 0 {
		return input
	}
	if input[len(input)-1:] == "\n" {

		return input[:len(input)-1]
	}
	return input
}
func (s *Account) SendMessage(to string, text string) {
	s.Sendnewmsg(to, removeEndNewLine(text))
}

func deBug(where string, err error) bool {
	if err != nil {
		fmt.Printf("\033[33m#%s\nReason:\n%s\n\n\033[39m", where, err)
		return false
	}
	return true
}
func (self *Account) ReissueChatTickets(groupId string) (tiket string, err error) {
	defer PanicOnly()
	req := &talkservice.ReissueChatTicketRequest{
		GroupMid: groupId,
		ReqSeq:   self.Seq,
	}
	self.Seq++
	res, err := self.Talk().ReissueChatTicket(self.Ctx, req)
	if err != nil {
		return "", err
	}
	return res.TicketId, err
}
func (self *Account) GetGroupMember(groupId string) (string, map[string]int64) {
	defer PanicOnly()
	res, err := self.Talk().GetChats(context.TODO(), &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: false,
		WithMembers:  true,
	})
	if err != nil {
		CheckErr(self, err, groupId, "GetGroupMember")
		return "", map[string]int64{}
	}
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		mem := ch.Extra.GroupExtra.MemberMids
		return ch.ChatName, mem
	}
	return "", map[string]int64{}
}
func (s *Account) GetChat(targets []string, opsiMembers bool, opsiPendings bool) (r *talkservice.GetChatsResponse) {
	defer PanicOnly()
	client := s.Talk()
	fst := talkservice.NewGetChatsRequest()
	fst.ChatMids = targets
	fst.WithMembers = opsiMembers
	fst.WithInvitees = opsiPendings
	r, e := client.GetChats(s.Ctx, fst)
	if e != nil {
		CheckErr(s, e, "none", "GetChat")
	}
	return r
}
func (self *Account) GetChatListMem(groupId string) (mem []string) {
	defer PanicOnly()
	res, _ := self.Talk().GetChats(self.Ctx, &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  true,
	})
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		for a := range ch.Extra.GroupExtra.MemberMids {
			mem = append(mem, a)
		}
		return mem
	}
	return []string{}
}

func (self *Account) GetChatListinv(groupId string) (inv []string) {
	defer PanicOnly()
	res, _ := self.Talk().GetChats(self.Ctx, &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  true,
	})
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		for a := range ch.Extra.GroupExtra.InviteeMids {
			inv = append(inv, a)
		}
		return inv
	}
	return []string{}
}
func IsPending(client *Account, to string, mid string) bool {
	defer PanicOnly()
	pend := client.GetChatListinv(to)
	if len(pend) != 0 {
		for i := range pend {
			if pend[i] == mid {
				return true
				//break
			}
		}
	}
	return false
}
func IsMembers(client *Account, to string, mid string) bool {
	defer PanicOnly()
	memlist := client.GetChatListMem(to)
	for i := range memlist {
		if memlist[i] == mid {
			return true
			//break
		}
	}
	return false
}
func (self *Account) FindAndAddContactsByMid4(mid string) (r map[string]*talkservice.Contact, err error) {
	if self.MID == mid || self.Limitadd {
		return map[string]*talkservice.Contact{}, nil
	} else if self.Frez {
		return map[string]*talkservice.Contact{}, nil
	}
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	self.Add += 1
	self.Lastadd = time.Now()
	if self.Add >= 10 {
		if !InArrayCl(Waitadd, self) {
			Waitadd = append(Waitadd, self)
			self.Timeadd = time.Now()
			BlockAdd.Set(self.MID, time.Now())
		}
		self.Limitadd = true
		return map[string]*talkservice.Contact{}, errors.New("limit goblok")
	}
	res, err := TS.FindAndAddContactsByMid(context.TODO(), self.Seq, mid, talkservice.ContactType_MID, `{"screen":"friendAdd:recommend","spec":"native"}`)
	if err != nil {
		e := GetCode(err)
		fmt.Println(err)
		if e == 35 {
			if _, ok := GetBlockAdd.Get(self.MID); !ok {
				GetBlockAdd.Set(self.MID, time.Now())
			}
			self.Limitadd = true
		}
	}
	return res, err
}
func AddContact3(cl *Account, con string) int {
	fl, _ := cl.GetAllContactIds()
	if !InArray(fl, con) {
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

func SetRoutine(ms time.Duration) (context.Context, context.CancelFunc) {
	ctx, deadline := context.WithDeadline(context.Background(), time.Now().Add(ms))
	return ctx, deadline
}
func (s *Account) SendMention(toID string, msgText string, mids []string) {
	defer PanicOnly()
	client := s.Talk()
	arr := []*tagdata{}
	mentionee := "@lined"
	texts := strings.Split(msgText, "@!")
	textx := ""
	for i := 0; i < len(mids); i++ {
		textx += texts[i]
		arr = append(arr, &tagdata{S: strconv.Itoa(len(textx)), E: strconv.Itoa(len(textx) + 6), M: mids[i]})
		textx += mentionee
	}
	textx += texts[len(texts)-1]
	allData, _ := json.MarshalIndent(arr, "", " ")
	msg := talkservice.NewMessage()
	msg.ContentType = 0
	msg.To = toID
	msg.Text = textx
	msg.ContentMetadata = map[string]string{"MENTION": "{\"MENTIONEES\":" + string(allData) + "}"}
	msg.RelatedMessageId = "0"
	_, e := client.SendMessage(s.Ctx, s.Seq, msg)
	deBug("SendMention", e)
}
func (self *Account) SendMention3(to string, text string, mids []string) (err error) {
	defer PanicOnly()
	if to != self.MID {
		arr := []*tagdata{}
		mentionee := "@Linebots2022"
		texts := strings.Split(text, "%v")
		if len(mids) == 0 || len(texts) < len(mids) {
			return errors.New("Invalid mids.")
		}
		textx := ""
		for i := 0; i < len(mids); i++ {
			textx += texts[i]
			uni := strconv.QuoteToASCII(string(textx))
			asswdx := utf8.RuneCountInString(textx)
			if strings.Contains(uni, "U0") {
				asswdx += strings.Count(uni, "U0")
			}
			asqqq := asswdx + 13
			slen := fmt.Sprintf("%v", asswdx)
			elen := fmt.Sprintf("%v", asqqq)
			arr = append(arr, &tagdata{S: slen, E: elen, M: mids[i]})
			textx += mentionee
		}
		textx += texts[len(texts)-1]
		arrData, _ := json.MarshalIndent(arr, "", " ")
		metas := map[string]string{"MENTION": "{\"MENTIONEES\":" + string(arrData) + "}"}
		M := &talkservice.Message{
			From_:            self.MID,
			To:               to,
			ToType:           2,
			Text:             textx,
			ContentType:      0,
			ContentMetadata:  metas,
			RelatedMessageId: "0",
		}
		self.Seq++
		ctx, cancel := SetRoutine(10 * time.Second)
		defer cancel()
		_, err = self.Talk().SendMessage(ctx, self.Seq, M)
	}
	return err
}
func (self *Account) SendMention4(to string, text string, mids []string) (err error) {
	defer PanicOnly()
	if to != self.MID {
		arr := []*tagdata{}
		mentionee := "@Linebots2022"
		texts := strings.Split(text, "!@")
		if len(mids) == 0 || len(texts) < len(mids) {
			return errors.New("Invalid mids.")
		}
		textx := ""
		for i := 0; i < len(mids); i++ {
			textx += texts[i]
			uni := strconv.QuoteToASCII(string(textx))
			asswdx := utf8.RuneCountInString(textx)
			if strings.Contains(uni, "U0") {
				asswdx += strings.Count(uni, "U0")
			}
			asqqq := asswdx + 13
			slen := fmt.Sprintf("%v", asswdx)
			elen := fmt.Sprintf("%v", asqqq)
			arr = append(arr, &tagdata{S: slen, E: elen, M: mids[i]})
			textx += mentionee
		}
		textx += texts[len(texts)-1]
		arrData, _ := json.MarshalIndent(arr, "", " ")
		metas := map[string]string{"MENTION": "{\"MENTIONEES\":" + string(arrData) + "}"}
		M := &talkservice.Message{
			From_:            self.MID,
			To:               to,
			ToType:           2,
			Text:             textx,
			ContentType:      0,
			ContentMetadata:  metas,
			RelatedMessageId: "0",
		}
		self.Seq++
		ctx, cancel := SetRoutine(10 * time.Second)
		defer cancel()
		_, err = self.Talk().SendMessage(ctx, self.Seq, M)
	}
	return err
}
func (self *Account) SendMention2(to string, text string, mids []string) (err error) {
	defer PanicOnly()
	if to != self.MID {
		arr := []*tagdata{}
		mentionee := "@LineD2022"
		texts := strings.Split(text, "TAGHERE")
		if len(mids) == 0 || len(texts) < len(mids) {
			return errors.New("Invalid mids.")
		}
		textx := ""
		for i := 0; i < len(mids); i++ {
			textx += texts[i]
			uni := strconv.QuoteToASCII(string(textx))
			asswdx := utf8.RuneCountInString(textx)
			if strings.Contains(uni, "U0") {
				asswdx += strings.Count(uni, "U0")
			}
			asqqq := asswdx + 13
			slen := fmt.Sprintf("%v", asswdx)
			elen := fmt.Sprintf("%v", asqqq)
			arr = append(arr, &tagdata{S: slen, E: elen, M: mids[i]})
			textx += mentionee
		}
		textx += texts[len(texts)-1]
		arrData, _ := json.MarshalIndent(arr, "", " ")
		metas := map[string]string{"MENTION": "{\"MENTIONEES\":" + string(arrData) + "}"}
		M := &talkservice.Message{
			From_:            self.MID,
			To:               to,
			ToType:           2,
			Text:             textx,
			ContentType:      0,
			ContentMetadata:  metas,
			RelatedMessageId: "0",
		}
		self.Seq++
		ctx, cancel := SetRoutine(10 * time.Second)
		defer cancel()
		_, err = self.Talk().SendMessage(ctx, self.Seq, M)
	}
	return err
}
func (self *Account) UnFriend(mid string) {
	err := self.Talk().UpdateContactSetting(context.TODO(), self.Seq, mid, talkservice.ContactSetting_CONTACT_SETTING_DELETE, "True")
	if err != nil {
		CheckErr(self, err, "none", "UnFriend")
	}
}
func (self *Account) AcceptTicket(groupMid string, ticketId string) (err error) {
	if self.Limited == false {
		_, err = self.Talk().AcceptChatInvitationByTicket(self.Ctx, &talkservice.AcceptChatInvitationByTicketRequest{
			ChatMid:  groupMid,
			ReqSeq:   self.Seq,
			TicketId: ticketId,
		})
		if err != nil {
			CheckErr(self, err, groupMid, "joinQr")
		}
		self.Seq++
	}
	return err
}

func (self *Account) CancelChatInvitations(groupId string, contactIds string) (err error) {
	if self.Limited == false {
		self.Talk().CancelChatInvitation(self.Ctx, &talkservice.CancelChatInvitationRequest{
			ChatMid:        groupId,
			ReqSeq:         self.Seq,
			TargetUserMids: []string{contactIds},
		})
		AddCount(self, "c")
		GetRoom(groupId).Fight = time.Now()
		if err != nil {
			CheckErr(self, err, groupId, "Cancel")
		}
		self.Seq++
	}
	return err
}
func (self *Account) InviteIntoGroupNormal(groupId string, contactIds []string) (err error) {
	if self.Limited == false {
		_, err = self.Talk().InviteIntoChat(self.Ctx, &talkservice.InviteIntoChatRequest{
			ChatMid:        groupId,
			ReqSeq:         self.Seq,
			TargetUserMids: contactIds,
		})
		AddCount(self, "invite")
		GetRoom(groupId).Fight = time.Now()
		if err != nil {
			CheckErr(self, err, groupId, "invite")
		}
		self.Seq++
	}
	return err
}
func (self *Account) NormalDeleteOtherFromChats(to string, contactIds string) (err error) {
	_, err = self.Talk().DeleteOtherFromChat(self.Ctx, &talkservice.DeleteOtherFromChatRequest{
		ReqSeq:         0,
		ChatMid:        to,
		TargetUserMids: []string{contactIds},
	})
	if err != nil {
		CheckErr(self, err, to, "checkban")
	}
	return err
}
func (self *Account) DeleteOtherFromChats(groupId string, contactIds string) (err error) {
	if self.Limited == false {
		_, err = self.Talk().DeleteOtherFromChat(self.Ctx, &talkservice.DeleteOtherFromChatRequest{
			ChatMid:        groupId,
			ReqSeq:         self.Seq,
			TargetUserMids: []string{contactIds},
		})
		AddCount(self, "kick")
		GetRoom(groupId).Fight = time.Now()
		if err != nil {
			CheckErr(self, err, groupId, "kick")
		}
		self.Seq++
	}
	return err
}

func (s *Account) GetGroupWithoutMembers(groupId string) (r *talkservice.Group) {
	client := s.Talk()
	r, _ = client.GetGroupWithoutMembers(s.Ctx, groupId)
	return r
}

func (s *Account) UpdateGroup(groupOBJ *talkservice.Group) {
	client := s.Talk()
	e := client.UpdateGroup(s.Ctx, s.Seq, groupOBJ)
	deBug("ReissueChatTicket", e)
}

func (self *Account) GetGroupMembers(groupId string) (error, string, map[string]int64) {
	res, err := self.Talk().GetChats(self.Ctx, &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: false,
		WithMembers:  true,
	})
	if err != nil {
		return err, "", map[string]int64{}
	}
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		mem := ch.Extra.GroupExtra.MemberMids
		return err, ch.ChatName, mem
	}
	return err, "", map[string]int64{}
}

func (self *Account) GetChatList(groupId string) (name string, mem, inv []string) {
	defer PanicOnly()
	res, _ := self.Talk().GetChats(self.Ctx, &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  true,
	})
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		for a := range ch.Extra.GroupExtra.MemberMids {
			mem = append(mem, a)
		}
		for a := range ch.Extra.GroupExtra.InviteeMids {
			inv = append(inv, a)
		}
		return ch.ChatName, mem, inv
	}
	return "", []string{}, []string{}
}
func (self *Account) GetGroups(groupId []string) (r []*talkservice.Chat, err error) {
	defer PanicOnly()
	tux := [][]string{}
	if len(groupId) > 100 {
		for {
			if len(groupId) != 0 {
				if len(groupId) < 99 {
					tux = append(tux, groupId)
					groupId = []string{}
				} else {
					tux = append(tux, groupId[:99])
					groupId = groupId[99:]
				}
			} else {
				break
			}
		}
	} else {
		tux = append(tux, groupId)
	}
	for _, lis := range tux {
		res, _ := self.Talk().GetChats(context.TODO(), &talkservice.GetChatsRequest{
			ChatMids:     lis,
			WithInvitees: true,
			WithMembers:  true,
		})
		r = append(r, res.Chats...)
	}
	return r, err
}
func (cl *Account) Sendnewmsg(to string, text string) (*talkservice.Message, error) {
	defer PanicOnly()
	M := &talkservice.Message{
		To:               to,
		ContentType:      0,
		Text:             text,
		RelatedMessageId: "0",
	}
	return cl.Talk().SendMessage(cl.Ctx, 0, M)
}
func (self *Account) GetGroupIdsJoined() ([]string, error) {
	defer PanicOnly()
	req := &talkservice.GetAllChatMidsRequest{
		WithInvitedChats: false,
		WithMemberChats:  true,
	}
	res, err := self.Talk().GetAllChatMids(context.TODO(), req, talkservice.SyncReason_UNKNOWN)
	return res.MemberChatMids, err
}
func (self *Account) UpdateProfileName(name string) (err error) {
	if len(name) < 1 {
		return err
	}
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	self.Seq++
	for {
		err = TS.UpdateProfileAttribute(context.TODO(), self.Seq, 2, name)
		if err == nil {
			self.Namebot = name
			return err
		} else if strings.Contains(err.Error(), "EOF") {
			continue
		} else {
			return err
		}
	}
	//return err
}
func (p *Account) GetE2EEPublicKeys() (r []*talkservice.E2EEPublicKey, err error) {
	r, err = p.Talk().GetE2EEPublicKeys(p.Ctx)
	return r, err
}

func (client *Account) RemoveLeterSelling() {
	a, _ := client.GetE2EEPublicKeys()
	for _, x := range a {
		err := client.RemoveE2EEPublicKey(x)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *Account) RemoveE2EEPublicKey(a *talkservice.E2EEPublicKey) (err error) {
	err = p.Talk().RemoveE2EEPublicKey(p.Ctx, a)
	return err
}
func (self *Account) GetChatListMap(groupId string) (mem, inv map[string]int64) {
	res, err := self.Talk().GetChats(context.TODO(), &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  true,
	})
	if err != nil {
		return mem, inv
	}
	if len(res.Chats) != 0 {
		ch := res.Chats[0]
		return ch.Extra.GroupExtra.MemberMids, ch.Extra.GroupExtra.InviteeMids
	}
	return map[string]int64{}, map[string]int64{}
}
func RemoveCl(items []*Account, item *Account) []*Account {
	defer PanicOnly()
	newitems := []*Account{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}
	return newitems
}
func Randint2(a []int) []int {
	defer PanicOnly()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}
func GetNewTransport() *http.Transport {
	certs, _ := tls.LoadX509KeyPair("rensmtcert.crt", "rensmtcert.key")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{certs},
			InsecureSkipVerify: true,
		},
	}
	http2.ConfigureTransport(tr)
	return tr
}
func (cl *Account) SettingUpHeader() {
	cl.HttpHeader = map[string][]string{"Content-Type": {"application/x-thrift"}}
	cl.HttpHeader.Add("user-agent", cl.UserAgent)
	cl.HttpHeader.Add("x-line-application", cl.AppName)
	cl.HttpHeader.Add("x-line-access", cl.AuthToken)
}
func (cl *Account) SetLocaleAccount() {
	sets, err := cl.GetSettings()
	if err == nil {
		cl.Locale = sets.PreferenceLocale
		cl.HttpHeader.Add("x-lal", cl.Locale)
	} else {
		cl.Locale = "in_ID"
		cl.HttpHeader.Add("x-lal", "in_ID")
	}
}
func CreateNewLogin(token string, num int, mids string, app string, ua string, host string) (*Account, error) {
	Hosts4 := fmt.Sprintf("https://%s.line.naver.jp/S4", host)
	SostP5 := fmt.Sprintf("https://%s.line.naver.jp/P5", host)
	SotShnal := fmt.Sprintf("https://%s.line.naver.jp/CH4", host)
	Urs4, _ := url.Parse(Hosts4)
	Urp5, _ := url.Parse(SostP5)
	Shna, _ := url.Parse(SotShnal)
	s := new(Account)
	s = &Account{
		AuthToken:     token,
		AppName:       app,
		UserAgent:     ua,
		Host:          host,
		MID:           mids,
		Limited:       false,
		Frez:          false,
		Limitadd:      false,
		Waitadd:       false,
		Seq:           0,
		Akick:         0,
		KickPoint:     0,
		Ainvite:       0,
		Transport:     GetNewTransport(),
		Acancel:       0,
		Namebot:       "",
		Numar:         num,
		hc:            setHC(),
		Shnall:        Shna.String(),
		UrS4:          Urs4.String(),
		UrP5:          Urp5.String(),
		Ckick:         0,
		Cinvite:       0,
		TimeBan:       time.Now(),
		Ccancel:       0,
		SHani:         0,
		Count:         50,
		CustomPoint:   0,
		GRevision:     0,
		Cpoll:         0,
		Ctx:           context.Background(),
		reqSeqMessage: 0,
		IRevision:     0,
		Squads:        []string{},
		Backup:        []string{},
	}
	s.SettingUpHeader()
	s.SetLocaleAccount()
	s.Revision, _ = s.GetLastOpRevision()
	prof, err := s.GetProfile()
	if err == nil {
		s.Namebot = prof.DisplayName
	}
	return s, err
}

func GetE2EE() (string, string, string) {
	out, _ := exec.Command("python3", "enc.py", "1").Output()
	//stdout, err := cmd.StdoutPipe()
	s := strings.Split(string(out), "\n")
	return s[0], s[1], s[2]
}

func GetIntBytes(n int) (valo []byte) {
	var bits = 64
	zigzag := ((n << 1) ^ (n >> (bits - 1)))
	for {
		if zigzag&-128 == 0 {
			valo = append(valo, byte(zigzag))
			break
		} else {
			valo = append(valo, byte((zigzag&0xff)|0x80))
			zigzag >>= 7
		}
	}
	return valo
}
func PanicOnly() {
	if r := recover(); r != nil {
		return
	}
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
func IsFriends(client *Account, from string) bool {
	defer PanicOnly()
	friendsip, _ := client.GetAllContactIds()
	for _, a := range friendsip {
		if a == from {
			return true
			//break
		}
	}
	return false
}
func (self *Account) GetGroupIdsInvited() (r []string, err error) {
	defer PanicOnly()
	req := &talkservice.GetAllChatMidsRequest{
		WithInvitedChats: true,
		WithMemberChats:  false,
	}
	rs, err := self.Talk().GetAllChatMids(context.TODO(), req, talkservice.SyncReason_UNKNOWN)
	return rs.InvitedChatMids, err
}

func (cl *Account) SendContact(to string, mid string) error {
	fOB := []byte{130, 33, 3, 11, 115, 101, 110, 100, 77, 101, 115, 115, 97, 103, 101, 21, 0, 28, 40, 33}
	fOB = append(fOB, GetStringBytes(to)...)
	fOB = append(fOB, []byte{213, 26, 59, 1, 136, 3, 109, 105, 100, 33}...)
	fOB = append(fOB, GetStringBytes(mid)...)
	fOB = append(fOB, []byte{0, 0}...)
	HTTP, _ := thrift.NewTHttpClient(cl.UrS4, cl.Transport)
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.AppName)
	transport.SetHeader("x-line-access", cl.AuthToken)
	transport.SetHeader("x-lal", cl.Locale)
	transport.Write(fOB)
	return transport.Flush(cl.Ctx)
}

func (cl *Account) UpdateChatQrV2(chatId string, typevar bool) error {
	HTTP, _ := thrift.NewTHttpClient(cl.UrS4, cl.Transport)
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("User-Agent", cl.UserAgent)
	transport.SetHeader("X-Line-Application", cl.AppName)
	transport.SetHeader("X-Line-Access", cl.AuthToken)
	transport.SetHeader("x-lal", cl.Locale)
	var x string
	if typevar {
		x = "!"
	}
	transport.Write([]byte("\x82!\x00\nupdateChat\x1c\x15\x00\x1c(!" + chatId + "l\x1c" + x + "\x00\x00\x00\x15\x08\x00\x00"))
	return transport.Flush(cl.Ctx)
}
func (cl *Account) SettingUpFetch() {
	cl.SessionPoll = thrift.ModHttpClient(cl.UrP5, cl.Transport, cl.HttpHeader)
	cl.SessionPoll.SetMoreCompact(true)
}
func (cl *Account) FetchOps(count int32) (res []*talkservice.Operation, err *modcompact.ExceptionMod) {
	var fOB = []byte{130, 33, 1, 8, 102, 101, 116, 99, 104, 79, 112, 115, 38}
	fOB = append(fOB, GetIntBytes(int(cl.Revision))...)
	fOB = append(fOB, 21)
	fOB = append(fOB, GetIntBytes(int(count))...)
	fOB = append(fOB, 22)
	fOB = append(fOB, GetIntBytes(int(cl.GRevision))...)
	fOB = append(fOB, 22)
	fOB = append(fOB, GetIntBytes(int(cl.IRevision))...)
	fOB = append(fOB, 0)
	cl.SettingUpFetch()
	cl.SessionPoll.Write(fOB)
	Times := time.Now().Unix()
	b, errS := cl.SessionPoll.FlushMod(cl.Ctx)
	if errS != nil {
		cl.Transport = GetNewTransport()
		cl.Curtime = Times
	} else if len(b) > 0 {
		tmcp := modcompact.TMoreCompactProtocolGoods(b)
		res, err = tmcp.GETOPS()
	}
	if len(res) == 0 || (Times-cl.Curtime) > 300 {
		cl.Transport = GetNewTransport()
		cl.Curtime = Times
	}
	return res, err
}

func (cl *Account) Talk() *talkservice.TalkServiceClient {
	compact := thrift.NewTCompactProtocol(thrift.ModHttpClient(cl.UrS4, cl.Transport, cl.HttpHeader))
	return talkservice.NewTalkServiceClient(thrift.NewTStandardClient(compact, compact))
}

func (cl *Account) LoadChannel() *channel.ChannelServiceClient {
	compact := thrift.NewTCompactProtocol(thrift.ModHttpClient(cl.Shnall, cl.Transport, cl.HttpHeader))
	return channel.NewChannelServiceClient(thrift.NewTStandardClient(compact, compact))
}
func (cl *Account) UpdateChatName(chatId string, name string) error {
	var fOB = []byte{130, 33, 1, 10, 117, 112, 100, 97, 116, 101, 67, 104, 97, 116, 28, 21, 0, 28, 21, 0, 24, 33}
	fOB = append(fOB, []byte(chatId)...)
	fOB = append(fOB, []byte{22, 0, 18, 22, 0, 24, byte(len(name))}...)
	fOB = append(fOB, []byte(name)...)
	fOB = append(fOB, []byte{24, 0, 0, 21, 2, 0, 0}...)
	HTTP, _ := thrift.NewTHttpClient(cl.UrS4, cl.Transport)
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("User-Agent", cl.UserAgent)
	transport.SetHeader("X-Line-Application", cl.AppName)
	transport.SetHeader("X-Line-Access", cl.AuthToken)
	transport.SetHeader("x-lal", cl.Locale)
	transport.Write(fOB)
	return transport.Flush(cl.Ctx)
}
func random() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(30000) + 30000)
}

func MaxRevision(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func FetchRev(client *Account, op *talkservice.Operation) {
	if op.Param2 != "" {
		a := strings.Split(op.Param2, "\x1e")
		res, err := strconv.ParseInt(a[0], 10, 64)
		if err == nil {
			client.GRevision = MaxRevision(client.GRevision, res)
		} else {
			println(op.Param2)
		}
	}
	if op.Param1 != "" {
		a := strings.Split(op.Param1, "\x1e")
		res, err := strconv.ParseInt(a[0], 10, 64)
		if err == nil {
			client.IRevision = MaxRevision(client.IRevision, res)
		} else {
			println(op.Param2)
		}
	}
}
func (cl *Account) CorrectRevision(op *talkservice.Operation, local bool, global bool, individual bool) {
	if global {
		if op.Revision == -1 && op.Param2 != "" {
			s := strings.Split(op.Param2, "\x1e")
			cl.GRevision, _ = strconv.ParseInt(s[0], 10, 64)
		}
	}
	if individual {
		if op.Revision == -1 && op.Param1 != "" {
			s := strings.Split(op.Param1, "\x1e")
			cl.IRevision, _ = strconv.ParseInt(s[0], 10, 64)
		}
	}
	if local {
		if op.Revision > cl.Revision {
			cl.Revision = op.Revision
		}
	}
}
func (cl *Account) GetLastOpRevision() (r int64, err error) {
	return cl.Talk().GetLastOpRevision(cl.Ctx)
}
func (self *Account) UpdateProfileBio(bio string) (err error) {
	if len(bio) < 1 {
		return err
	}
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	self.Seq++
	for {
		err = TS.UpdateProfileAttribute(context.TODO(), self.Seq, 16, bio)
		if err == nil {
			return err
		} else if strings.Contains(err.Error(), "EOF") {
			continue
		} else {
			return err
		}
	}
	//return err
}
func (cl *Account) GetContact(mid string) (*talkservice.Contact, error) {
	return cl.Talk().GetContact(cl.Ctx, mid)
}

func (cl *Account) SendText(to string, text string) (*talkservice.Message, error) {
	M := &talkservice.Message{
		To:               to,
		ContentType:      0,
		Text:             text,
		RelatedMessageId: "0",
	}
	return cl.Talk().SendMessage(context.TODO(), 0, M)
}
func GetStringBytes(str string) []byte {
	var va []byte
	for a := range str {
		va = append(va, byte(int(str[a])))
	}
	return va
}

func (cl *Account) FindAndAddContactsByMid(mid string) error {
	_, err := cl.Talk().FindAndAddContactsByMid(cl.Ctx, 0, mid, 0, "")
	return err
}
func (self *Account) GetGroup3(groupId string) (r *talkservice.Chat, err error) {
	res, err := self.Talk().GetChats(context.TODO(), &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  true,
	})
	if err != nil {
		CheckErr(self, err, groupId, "GetGroup3")
	}
	return res.Chats[0], err
}
func (self *Account) GetRecentMessagesV2(gid string) (r []*talkservice.Message, err error) {
	ls, err := self.Talk().GetRecentMessagesV2(self.Ctx, gid, int32(100000000))
	if err != nil {
		CheckErr(self, err, gid, "GetRecentMessagesV2")
	}
	return ls, err
}

func (cl *Account) UpdateProfileAttribute(a talkservice.ProfileAttribute, v string) error {
	return cl.Talk().UpdateProfileAttribute(cl.Ctx, 0, a, v)
}

func (cl *Account) UpdateProfileAttributes(a talkservice.ProfileAttribute, v string) error {
	D := make(map[talkservice.ProfileAttribute]*talkservice.ProfileContent)
	D[a] = &talkservice.ProfileContent{Value: v}
	return cl.Talk().UpdateProfileAttributes(cl.Ctx, 0, &talkservice.UpdateProfileAttributesRequest{ProfileAttributes: D})
}

func (cl *Account) GetProfile() (*talkservice.Profile, error) {
	return cl.Talk().GetProfile(cl.Ctx, 3)
}

func (cl *Account) GetAllContactIds() ([]string, error) {
	return cl.Talk().GetAllContactIds(cl.Ctx, 1)
}

func (cl *Account) AcceptChatInvitationByTicket(to string, ticket string) error {
	_, err = cl.Talk().AcceptChatInvitationByTicket(cl.Ctx, &talkservice.AcceptChatInvitationByTicketRequest{
		ReqSeq:   0,
		ChatMid:  to,
		TicketId: ticket,
	})
	return err
}

func (cl *Account) DeleteSelfFromChat(to string) error {
	_, err = cl.Talk().DeleteSelfFromChat(cl.Ctx, &talkservice.DeleteSelfFromChatRequest{
		ReqSeq:  0,
		ChatMid: to,
	})
	return err
}

func (s *Account) FindChatByTicket(ticketId string) (r *talkservice.FindChatByTicketResponse) {
	client := s.Talk()
	v := talkservice.NewFindChatByTicketRequest()
	v.TicketId = ticketId
	r, e := client.FindChatByTicket(s.Ctx, v)
	deBug("ReissueChatTicket", e)
	return r
}
func (self *Account) GetSettings() (r *talkservice.Settings, err error) {
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	res, err := TS.GetSettings(self.Ctx, talkservice.SyncReason_UNKNOWN)
	return res, err
}
func (cl *Account) InviteIntoChatPollVer(to string, mid []string) {
	if len(mid) > 3 {
		var b []string
		for x, v := range mid {
			x++
			b = append(b, v)
			if x%3 == 0 {
				go cl.InviteIntoGroupNormal(to, b)
				b = []string{}
			}
		}
		if len(b) > 0 {
			go cl.InviteIntoGroupNormal(to, b)
		}
		time.Sleep(5 * time.Millisecond)
	} else {
		cl.InviteIntoGroupNormal(to, mid)
	}
}

func (self *Account) GetChats(groups []string) (*talkservice.GetChatsResponse, error) {
	return self.Talk().GetChats(self.Ctx, &talkservice.GetChatsRequest{
		ChatMids:     groups,
		WithMembers:  true,
		WithInvitees: true,
	})
}
func (self *Account) RejectChatInvitation(chatMid string) error {
	req := talkservice.NewRejectChatInvitationRequest()
	req.ReqSeq = self.Seq
	self.Seq++
	req.ChatMid = chatMid
	_, err := self.Talk().RejectChatInvitation(self.Ctx, req)
	return err
}
func (self *Account) GetGroupsInvited() (r []string, err error) {
	req := &talkservice.GetAllChatMidsRequest{
		WithInvitedChats: true,
		WithMemberChats:  false,
	}
	rs, err := self.Talk().GetAllChatMids(self.Ctx, req, talkservice.SyncReason_UNKNOWN)
	return rs.InvitedChatMids, err
}
func (self *Account) ReissueChatTicket(groupId string) (tiket string, err error) {
	req := &talkservice.ReissueChatTicketRequest{
		GroupMid: groupId,
		ReqSeq:   self.Seq,
	}
	self.Seq++
	res, err := self.Talk().ReissueChatTicket(self.Ctx, req)
	if err != nil {
		CheckErr(self, err, groupId, "ReissueChatTicket")
		return "", err
	}
	return res.TicketId, err
}
func (self *Account) GetGroupInvitation(groupId string) (string, map[string]int64) {
	res, err := self.Talk().GetChats(context.TODO(), &talkservice.GetChatsRequest{
		ChatMids:     []string{groupId},
		WithInvitees: true,
		WithMembers:  false,
	})
	if err != nil {
		CheckErr(self, err, groupId, "GET_GROUP")
		return "", map[string]int64{}
	}
	ch := res.Chats[0]
	mem := ch.Extra.GroupExtra.InviteeMids
	return ch.ChatName, mem
}
func (cl *Account) GetSameJoiningTime(to string, enemy string) []string {
	defer PanicOnly()
	memlist, _ := cl.GetChatListMap(to)
	var together []string
	var tj string
	if IsMember(memlist, enemy) {
		tj = strconv.FormatInt(memlist[enemy], 10)[:9]
		for k, v := range memlist {
			if strconv.FormatInt(v, 10)[:9] == tj {
				together = append(together, k)
			}
		}
	}
	return together
}
func (cl *Account) GetTargetKickall(to string) []string {
	c, err := cl.GetChats([]string{to})
	if err != nil {
		return []string{}
	}
	zxc := c.Chats[0].Extra.GroupExtra.MemberMids
	var together []string
	for k := range zxc {
		if k != cl.MID {
			together = append(together, k)
		}
	}
	return together
}

func (cl *Account) GetTargetCancelall(to string) []string {
	c, err := cl.GetChats([]string{to})
	if err != nil {
		return []string{}
	}
	zxc := c.Chats[0].Extra.GroupExtra.InviteeMids
	var together []string
	for k := range zxc {
		if k != cl.MID {
			together = append(together, k)
		}
	}
	return together
}

func GETBackup(Midlist []string, mid string) int {
	for num, x := range Midlist {
		if x == mid {
			if num+1 == len(Midlist) {
				return 0
			}
			return num + 1
		}
	}
	return 404
}
func (self *Account) GetRecommendationIds() (r []string, err error) {
	res, err := self.Talk().GetRecommendationIds(self.Ctx, talkservice.SyncReason_UNKNOWN)
	return res, err
}
func (self *Account) GetContacts(id []string) (r []*talkservice.Contact, err error) {
	res, err := self.Talk().GetContacts(self.Ctx, id)
	return res, err
}
func (self *Account) AcceptGroupInvitationNormal(groupId string) (err error) {
	self.Talk().AcceptChatInvitation(self.Ctx, &talkservice.AcceptChatInvitationRequest{
		ChatMid: groupId,
		ReqSeq:  self.Seq,
	})
	return nil
}
func (self *Account) FindAndAddContactsByMid3(mid string) (r map[string]*talkservice.Contact, err error) {
	if self.MID == mid || self.Limitadd {
		return map[string]*talkservice.Contact{}, nil
	} else if self.Frez {
		return map[string]*talkservice.Contact{}, nil
	}
	var TS *talkservice.TalkServiceClient
	TS = self.Talk()
	self.Add += 1
	self.Lastadd = time.Now()
	if self.Add >= 10 {
		if !InArrayCl(Waitadd, self) {
			Waitadd = append(Waitadd, self)
			self.Timeadd = time.Now()
			BlockAdd.Set(self.MID, time.Now())
		}
		self.Limitadd = false
		return map[string]*talkservice.Contact{}, errors.New("limit goblok")
	}
	res, err := TS.FindAndAddContactsByMid(context.TODO(), self.Seq, mid, talkservice.ContactType_MID, `{"screen":"friendAdd:recommend","spec":"native"}`) //`{"screen":"homeTab","spec":"native"}`)
	if err != nil {
		e := GetCode(err)
		fmt.Println(err)
		if e == 35 {
			if _, ok := GetBlockAdd.Get(self.MID); !ok {
				GetBlockAdd.Set(self.MID, time.Now())
			}
			self.Limitadd = false
		}
	}
	return res, err
}

///Rom
func RoomClear(item *LineRoom) {
	defer PanicOnly()
	newitems := []*LineRoom{}
	for _, i := range SquadRoom {
		if i.Id == item.Id {
			newitems = append(newitems, i)
		}
	}
	SquadRoom = newitems
}
func RemoveRoom(items []*LineRoom, item *LineRoom) []*LineRoom {
	defer PanicOnly()
	newitems := []*LineRoom{}
	for _, i := range items {
		if i.Id != item.Id {
			newitems = append(newitems, i)
		}
	}
	return newitems
}
func (self *Account) DeleteOtherFromChat(groupId string, contactIds string) int {
	if self.Limited == false {
		_, err = self.Talk().DeleteOtherFromChat(context.TODO(), &talkservice.DeleteOtherFromChatRequest{
			ChatMid:        groupId,
			ReqSeq:         self.Seq,
			TargetUserMids: []string{contactIds},
		})
		var aa = 0
		AddCount(self, "kick")
		GetRoom(groupId).Fight = time.Now()
		if err != nil {
			aa = CheckErr(self, err, groupId, "KICK")
		}
		self.Seq++
		return aa
	}
	return 35
}
func GetRoom(to string) *LineRoom {
	for _, room := range SquadRoom {
		if room.Id == to {
			return room
		}
	}
	new := &LineRoom{Id: to, Userlurk: []string{}, MsgLeave: "See you nix time", WelcomeMsg: "Hallo Welcome Join", MsgLurk: "Hallo I See You.", Gowner: []string{}, Gadmin: []string{}, Gban: []string{}, ListInvited: []string{}, Bot: []string{}, GoMid: []string{}, ProKick: false, Limit: false, ProQr: false, ProName: false, ProInvite: false, ProJoin: false, ProCancel: false, AntiTag: false, Automute: false, Lurk: false, Announce: false, Qr: false, Purge: false}
	SquadRoom = append(SquadRoom, new)
	return new
}

func Actor(to string) (anu []*Account) {
	for _, room := range SquadRoom {
		if room.Id == to {
			return room.HaveClient
		}
	}
	return anu
}
func Gones(to string, cl *Account) {
	for _, room := range SquadRoom {
		if room.Id == to {
			room.HaveClient = RemoveCl(room.HaveClient, cl)
			return
		}
	}
}

// TODO
// here HashToMap func

func (self *LineRoom) Reset() {
	self.Qr = true
	for _, cl := range self.Ava {
		if !cl.Client.Limited {
			cl.Exist = true
		}
	}
	self.HaveClient = append([]*Account{}, self.Client...)
}
func SetAva(to string, list []string) {
	for _, room := range SquadRoom {
		if room.Id == to {
			for _, cls := range room.Ava {
				if !cls.Client.Limited {
					if InArray(list, cls.Mid) {
						cls.Exist = true
					} else {
						cls.Exist = false
					}
				} else {
					cls.Exist = false
				}
			}
		}
	}
}
func (self *LineRoom) AddSquad(bot []string, cls []*Account, goclint []*Account, midgo []string) {
	self.Bot = bot
	self.Client = cls
	self.HaveClient = []*Account{}
	self.Ava = []*Ava{}
	self.GoMid = midgo
	self.GoClient = goclint
	for _, cl := range cls {
		if !cl.Limited {
			self.Ava = append(self.Ava, &Ava{Client: cl, Exist: true, Mid: cl.MID})
			self.HaveClient = append(self.HaveClient, cl)
		} else {
			self.Ava = append(self.Ava, &Ava{Client: cl, Exist: false, Mid: cl.MID})
		}
	}
}
func (self *LineRoom) DelGo(cl *Account) {
	if InArray(self.GoMid, cl.MID) {
		self.GoMid = Remove(self.GoMid, cl.MID)
	}
	if InArrayCl(self.GoClient, cl) {
		self.GoClient = RemoveCl(self.GoClient, cl)
	}
}

func (self *LineRoom) Choose(client *Account) *Account {
	for _, cl := range self.Ava {
		if cl.Exist && cl.Client.Limited {
			return cl.Client
		}
	}
	return client
}

func (self *LineRoom) Names() []string {
	anu := []string{}
	for _, cl := range self.Ava {
		if cl.Exist {
			anu = append(anu, cl.Client.Namebot)
		}
	}
	return anu
}

func (self *LineRoom) Cans() []*Account {
	anu := []*Account{}
	for _, cl := range self.Ava {
		if cl.Exist {
			anu = append(anu, cl.Client)
		}
	}
	return anu
}

func (self *LineRoom) Invites() []string {
	anu := []string{}
	for _, cl := range self.Ava {
		if !cl.Exist {
			anu = append(anu, cl.Client.MID)
		}
	}
	return anu
}
func RemoveAva(items []*Ava, item *Ava) []*Ava {
	defer PanicOnly()
	newitems := []*Ava{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}
	return newitems
}
func (self *LineRoom) RevertGo(cl *Account) {
	if InArray(self.Bot, cl.MID) {
		self.Bot = Remove(self.Bot, cl.MID)
	}
	if InArrayCl(self.Client, cl) {
		self.Client = RemoveCl(self.Client, cl)
	}
	if InArrayCl(self.HaveClient, cl) {
		self.HaveClient = append(self.HaveClient, cl)
	}
	for _, ava := range self.Ava {
		if ava.Client == cl {
			self.Ava = RemoveAva(self.Ava, ava)
		}
	}
	if !InArray(self.GoMid, cl.MID) {
		self.GoMid = append(self.GoMid, cl.MID)
	}
	if !InArrayCl(self.GoClient, cl) {
		self.GoClient = append(self.GoClient, cl)
	}
}
func (self *LineRoom) ConvertGo(cl *Account) {
	if !InArray(self.Bot, cl.MID) {
		self.Bot = append(self.Bot, cl.MID)
	}
	if !InArrayCl(self.Client, cl) {
		self.Client = append(self.Client, cl)
	}
	if !InArrayCl(self.HaveClient, cl) {
		self.HaveClient = append(self.HaveClient, cl)
	}
	if !InArrayAva(self.Ava, cl) {
		if cl.Limited {
			self.Ava = append(self.Ava, &Ava{Client: cl, Exist: true, Mid: cl.MID})
		} else {
			self.Ava = append(self.Ava, &Ava{Client: cl, Exist: false, Mid: cl.MID})
		}
	}
	if InArray(self.GoMid, cl.MID) {
		self.GoMid = Remove(self.GoMid, cl.MID)
	}
	if InArrayCl(self.GoClient, cl) {
		self.GoClient = RemoveCl(self.GoClient, cl)
	}
}
func InArrayAva(arr []*Ava, str *Account) bool {
	for _, tar := range arr {
		if tar.Client == str {
			return true
		}
	}
	return false
}
func Archimed(s string, list []string) []string {
	ln := len(list)
	ls := []int{}
	ind := []int{}
	hasil := []string{}
	if strings.Contains(s, ",") {
		logics := strings.Split(s, ",")
		for _, logic := range logics {
			if strings.Contains(logic, ">") {
				su := strings.TrimPrefix(logic, ">")
				si, _ := strconv.Atoi(su)
				si -= 1
				for i := (si + 1); i > si && i <= ln; i++ {
					ls = append(ls, i)
				}
			} else if strings.Contains(logic, "<") {
				su := strings.TrimPrefix(logic, "<")
				si, _ := strconv.Atoi(su)
				si -= 1
				for i := 0; i <= si; i++ {
					ls = append(ls, i)
				}
			} else if strings.Contains(logic, "-") {
				las := strings.Split(logic, "-")
				si := las[0]
				siu, _ := strconv.Atoi(si)
				siu -= 1
				sa := las[1]
				sau, _ := strconv.Atoi(sa)
				sau -= 1
				for i := (siu); i >= siu && i <= sau; i++ {
					ls = append(ls, i)
				}
			} else {
				sau, _ := strconv.Atoi(logic)
				sau -= 1
				ls = append(ls, sau)
			}
		}
	} else {
		logic := s
		if strings.Contains(logic, ">") {
			su := strings.TrimPrefix(logic, ">")
			si, _ := strconv.Atoi(su)
			si -= 1
			for i := (si + 1); i > si && i <= ln; i++ {
				ls = append(ls, i)
			}
		} else if strings.Contains(logic, "<") {
			su := strings.TrimPrefix(logic, "<")
			si, _ := strconv.Atoi(su)
			si -= 1
			for i := 0; i <= si; i++ {
				ls = append(ls, i)
			}
		} else if strings.Contains(logic, "-") {
			las := strings.Split(logic, "-")
			si := las[0]
			siu, _ := strconv.Atoi(si)
			siu -= 1
			sa := las[1]
			sau, _ := strconv.Atoi(sa)
			sau -= 1
			for i := (siu); i >= siu && i <= sau; i++ {
				ls = append(ls, i)
			}
		} else {
			sau, _ := strconv.Atoi(logic)
			sau -= 1
			ls = append(ls, sau)
		}
	}
	for _, do := range ls {
		if !InArrayInt(ind, do) && do < ln {
			jo := list[do]
			ind = append(ind, do)
			hasil = append(hasil, jo)
		}
	}
	return hasil
}
func (self *LineRoom) Joins(cl *Account) {
	defer PanicOnly()
	for _, cls := range self.Ava {
		if cls.Client == cl {
			if cl.Limited {
				cls.Exist = true
			}
			return
		}
	}
}
func Checkarri(cl []*Account, self *Account) bool {
	defer PanicOnly()
	for _, cls := range cl {
		if self == cls {
			return true
		}
	}
	return false
}
func Qrend(to string) {
	for _, room := range SquadRoom {
		if room.Id == to {
			room.Qr = true
			return
		}
	}
	new := &LineRoom{Id: to, Userlurk: []string{}, MsgLeave: "See you nix time", WelcomeMsg: "Hallo Welcome Join", MsgLurk: "Hallo I See You.", ListInvited: []string{}, Gowner: []string{}, Gadmin: []string{}, Gban: []string{}, Bot: []string{}, GoMid: []string{}, ProKick: false, Limit: false, ProQr: false, ProName: false, ProInvite: false, ProJoin: false, ProCancel: false, AntiTag: false, Automute: false, Lurk: false, Announce: false, Qr: false, Purge: false}
	SquadRoom = append(SquadRoom, new)
}
func ClearProtect() {
	for _, room := range SquadRoom {
		room.ProKick = false
		room.ProCancel = false
		room.ProName = false
		room.ProInvite = false
		room.ProJoin = false
		room.ProQr = false
	}
}
func (room *LineRoom) AutoBro() {
	room.ProKick = true
	room.ProCancel = true
	room.ProName = true
	room.ProInvite = true
	room.ProQr = true
}
func ListProtect() string {
	ret := "༗ 𝗟𝗶𝘀𝘁 𝗣𝗿𝗼𝘁𝗲𝗰𝘁 ༗"
	for _, room := range SquadRoom {
		ret += fmt.Sprintf("\n\n Group: %s\n", room.Name)
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
		if len(room.GoMid) > 0 {
			ret += "\n 🟢 Pro Ajs"
		} else {
			ret += "\n 🔴 Pro Ajs"
		}
	}
	return ret
}

func Qrstart(to string) {
	for _, room := range SquadRoom {
		if room.Id == to {
			room.Qr = false
			return
		}
	}
	new := &LineRoom{Id: to, Userlurk: []string{}, MsgLeave: "See you nix time", WelcomeMsg: "Hallo Welcome Join", MsgLurk: "Hallo I See You.", ListInvited: []string{}, Gowner: []string{}, Gadmin: []string{}, Gban: []string{}, Bot: []string{}, GoMid: []string{}, ProKick: false, Limit: false, ProQr: false, ProName: false, ProInvite: false, ProJoin: false, ProCancel: false, AntiTag: false, Automute: false, Lurk: false, Announce: false, Qr: false, Purge: false}
	SquadRoom = append(SquadRoom, new)
}
func CheckSquadRoom() []*LineRoom {
	return SquadRoom
}
func (room *LineRoom) WelsomeSet(cl *Account, to, mid string) {
	if room.Welcome {
		list := fmt.Sprintf("Group: %v\n\n", room.Name)
		list += "!@ "
		list += fmt.Sprintf("\n\n%v", room.WelcomeMsg)
		cl.SendMention4(to, list, []string{mid})
	}

}
func (room *LineRoom) LeaveSet(cl *Account, to, mid string) {
	if room.Leavebool {
		list := fmt.Sprintf("Group: %v\n\n", room.Name)
		list += "!@ "
		list += fmt.Sprintf("\n\n%v", room.MsgLeave)
		cl.SendMention4(to, list, []string{mid})
	}

}
func (room *LineRoom) CheckLurk(cl *Account, to, mid string) {
	if !InArray(room.Userlurk, mid) {
		if room.NameLurk == "name" {
			room.Userlurk = append(room.Userlurk, mid)
			x, _ := cl.GetContact(mid)
			lis := "- %v "
			lis += room.MsgLurk
			list := fmt.Sprintf(lis, x.DisplayName)
			cl.SendMessage(to, list)
		} else if room.NameLurk == "mention" {
			room.Userlurk = append(room.Userlurk, mid)
			lis := "%v "
			lis += room.MsgLurk
			cl.SendMention3(to, lis, []string{mid})
		} else if room.NameLurk == "hide" {
			room.Userlurk = append(room.Userlurk, mid)
		}
		if room.ImageLurk == true {
			x, _ := cl.GetContact(mid)
			cl.SendImageWithURL(to, "https://profile.line-scdn.net/"+x.PicturePath)
		}
	}
}

func (room *LineRoom) CheckAnnounce(cl *Account, to string) {
	chat, from := cl.GetChatRoomAnnouncements(to)
	tx := fmt.Sprintf("Announced a message:\nFrom: @! \n\n %s", chat)
	cl.SendMention(to, tx, []string{from})
}
func (cl *Account) GetChatRoomAnnouncements(roomId string) (chat string, mid string) {
	ann, err := cl.Talk().GetChatRoomAnnouncements(cl.Ctx, roomId)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	if len(ann) != 0 {
		chat = ann[0].Contents.Text
		mid = ann[0].CreatorMid
		return chat, mid
	}
	return "", ""
}
func IoGOBot(to string, cl *Account) bool {
	for _, room := range SquadRoom {
		if room.Id == to {
			for _, cls := range room.GoClient {
				if cl == cls {
					return false
				}
			}
			return true
		}
	}
	return true
}
func (self *LineRoom) Act(cl *Account) bool {
	for _, cls := range self.Ava {
		if cls.Client == cl {
			if cls.Exist {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
