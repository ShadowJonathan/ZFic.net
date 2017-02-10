package ZFic

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"strconv"

	"time"

	"net/http"

	"github.com/russross/blackfriday"
)

func GetMD(input []byte) []byte {
	return blackfriday.MarkdownCommon(input)
}

func HandlePage(icfg string) []byte {
	CFGDATA, err := ioutil.ReadFile(icfg)
	if err != nil {
		panic(err)
	}
	var Icfg = &TopLayer{}
	err = json.Unmarshal(CFGDATA, Icfg)
	if err != nil {
		panic(err)
	}
	TL, err := ioutil.ReadFile(tp + Icfg.TP)
	if err != nil {
		panic(err)
	}
	var total []byte
	if Icfg.Inception {
		total = HandleLayer(Icfg.I, TL)
	} else {
		total = TL
	}
	return total
}

func HandleLayer(L *sublayer, UL []byte) []byte {
	LAYER, err := ioutil.ReadFile(ZFroot + L.TP)
	if err != nil {
		panic(err)
	}
	var total []byte
	if L.IsMD {
		total = bytes.Replace(UL, []byte("TPDATA"), GetMD(LAYER), 1)
	} else {
		total = bytes.Replace(UL, []byte("TPDATA"), LAYER, 1)
	}
	if L.Inception {
		return HandleLayer(L.I, total)
	}
	return total
}

func GetUser(username string, password string) (bool, *User) {
	var emu = &User{}
	var users = *DB
	for _, U := range users {
		if username == U.Username && U.Password == password {
			return true, U
		}
	}
	return false, emu
}

func GetUserSimple(user string) (bool, *User) {
	var emu = &User{}
	var get = *DB
	users := get
	for _, U := range users {
		if user == strconv.FormatInt(int64(U.ID), 10) {
			u := *U
			US := RemoveInfo(u)
			us := &US
			// @.@
			// this took me WAY too long to figure out and make it work, im guessing there is a better way
			// if i dont "redirect", it just removes the password and username of the original database, thats not good >.<
			return true, us
		}
	}
	return false, emu
}

func LoadDataBase() {
	users, err := ioutil.ReadFile(fileroot + "/Users.Ur")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(users, DB)
	if err == nil {
		return
	} else {
		fmt.Println(err)
	}
}

func LoadSessions() {
	sess, err := ioutil.ReadFile(fileroot + "/Sessions.ses")
	if err != nil {
		fmt.Println(err)
	}
	Sesses = &map[string]*Session{}
	err = json.Unmarshal(sess, Sesses)
	if err == nil {
		return
	} else {
		fmt.Println(err)
	}
}

func SaveSessions() {

	var Sessions = *Sesses
	sess, _ := json.Marshal(Sessions)
	ioutil.WriteFile(fileroot+"/Sessions.ses", sess, 9001)
}

func GetNewSession(user string, w http.ResponseWriter) {

	var Sessions = *Sesses
	s, err := GenerateRandomString(32)
	if err != nil {
		log.Fatal(err)
	} else {
		var Ses = &Session{}
		Ses.UserID = user
		Ses.Expires = time.Now().AddDate(0, 0, 7) //sessions expire in 7 days
		Sessions[s] = Ses
		http.SetCookie(w, &http.Cookie{
			Name:    "zfsn",
			Value:   s,
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 7),
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "zfid",
			Value:   user,
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 7),
		})
	}
	SaveSessions()
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func GSFI(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func GetDateTime(T TimeandDate) string {
	return GSFI(T.Day) + " " + T.Month.String() + " " + GSFI(T.Year) + " " + GSFI(T.Hour) + ":" + GSFI(T.Min)
}

func RemoveInfo(U User) User {
	var u = U
	u.Username = ""
	u.Password = ""
	return u
}

func Nulldate() time.Time {
	return time.Date(1970, time.January, 1, 1, 1, 1, 1, time.UTC)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request, cookie string) {
	c, err := r.Cookie(cookie)
	if err != nil {
		c = &http.Cookie{Name: cookie}
	}
	c.Expires = Nulldate()
	c.Path = "/"
	http.SetCookie(w, c)
}

func HandleError(ErrorCode int, w http.ResponseWriter) {
	switch ErrorCode {
	case 404:
		w.Write(HandlePage(fileroot + "/error/fourohfour.icfg"))
	}
}

func GetFromSession(r *http.Request) (*User, bool) {
	var EmUs = &User{}
	SessionCookie, err := r.Cookie("zfsn")
	if err != nil {
		return EmUs, false
	}
	var Sessions = *Sesses
	for a, b := range Sessions {
		if a == SessionCookie.Value {
			_, u := GetUserSimple(b.UserID)
			return u, true
		}
	}
	return EmUs, false
}
