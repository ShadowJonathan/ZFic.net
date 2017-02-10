package ZFic

import (
	"fmt"
	"time"
)

func Reap() {
	var Dead []string
	var Sessions = *Sesses
	for s, ses := range Sessions {
		if time.Since(ses.Expires) > 0 {
			Dead = append(Dead, s)
		}
	}
	if len(Dead) > 0 {
		for _, name := range Dead {
			delete(Sessions, name)
			fmt.Println(name + " has been expired")
		}
	}
	SaveSessions()
}

func Kill(session string) {
	var Sessions = *Sesses
	delete(Sessions, session)
	fmt.Println(session + " has been killed")
	SaveSessions()
}

func StartReaper() {
	Reap()
	go ReapLoop()
}

func ReapLoop() {
	for {
		time.Sleep(10 * time.Minute)
		Reap()
	}
}

func DieDieDie() int {
	var Dead []string
	var Sessions = *Sesses
	for s, _ := range Sessions {
		Dead = append(Dead, s)
	}
	var sescount = len(Dead)
	if len(Dead) > 0 {
		for _, name := range Dead {
			delete(Sessions, name)
			fmt.Println(name + " has been killed with many others")
		}
	}
	return sescount
}
