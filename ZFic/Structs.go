package ZFic

import "time"

//server

type ServerConfig struct {
	Archiveworker bool  // is worker working, yes/no
	Stories       int64 // total ampunt of stories overall
	AWQ           chan ARequest
}

type HttpServer struct {
	port    string
	address string
}

func (h *HttpServer) Port() string    { return h.port }
func (h *HttpServer) Address() string { return h.address }

//request

type ARequest struct {
	TypeReq int // 0: Create, 1: Update
	Create  *CreateReq
	Update  *StoryUpdate
}

type CreateReq struct { //initial story creation request, makes pointers, files, and references for the story
	Title     string
	Summary   string
	Author    StoryAuthor
	Tags      []*StoryTag
	StoryFile string //filepath to first chapter markdown
}

type StoryUpdate struct { //update request
	UType       []int // 0: update tags, 1: update title, 2: update summary, 3: update chapters, 4: delete
	StoryID     int
	NewTags     []*StoryTag
	NewTitle    string
	NewSummary  string
	NewChapters []*StoryChapter
	Notify      bool //is this update "note-worthy" to be sent to email?
}

//story

type Story struct { //this is what will be saved in an archive in an array
	Title         string
	Summary       string
	ID            int
	Author        StoryAuthor
	Chapters      []*StoryChapter
	ChapCount     int `json:"CC"`
	Tags          []*StoryTag
	Published     TimeandDate
	IsUpdatedOnce bool        `json:"IUO"`
	Updated       TimeandDate `json:",omitempty"`
	ExtraInfo     StoryStats
}

type StoryAuthor struct {
	ID     int
	Author Author
}

type StoryTag struct { //used in an array in the archive file, reference tag
	Type      int `json:"T"` //0: character, 1: rating, 2: special (for content filtering), 3: Global tag-type, 4: User-defined Tag (with ID)
	CharID    int `json:",omitempty"`
	UDTID     int `json:",omitempty"` //id for user defined tag
	GDTID     int `json:",omitempty"` //id for globally defined tag
	Rating    int `json:",omitempty"` // 0: everyone, 1: teen, 2: mature
	SpecialID int `json:"SID,omitempty"`
}

type StoryChapter struct {
	StoryID        int
	CName          string //name of the chapter
	RawChapterFile string `json:"RCF,omitempty"` //we save every chapter that is linked in another place "just in case" the html needs to be rebuilded
	Pos            int    //chapter position, start with 1, 0 is the "front page"
	FrontPageFile  string `json:"FPF,omitempty"` //if its the front page, then ignore the chapterfile and use a "special"(?) file
	// chapter summary maybe??
}

type TimeandDate struct {
	Year  int
	Month time.Month
	Day   int
	Hour  int
	Min   int
	Sec   int
}

type StoryStats struct {
	Alerts    []int
	Favorites []int
	Reviews   []Review
}

type Review struct {
	ID        int    `json:",omitempty"`
	Anon      bool   `json:",omitempty"`
	Title     string `json:",omitempty"`
	Content   string
	Posted    TimeandDate
	OnChapter StoryChapter `json:",omitempty"`
}

//tags

type GlobalTag struct {
	ID   int
	name string
}

type SpecialTag struct {
	ID     int
	name   string
	Mature bool `json:",omitempty"`
	Gore   bool `json:",omitempty"`
}

type UserTag struct {
	ID        int
	name      string
	Owner     User
	CreatedOn TimeandDate
}

//user

type User struct {
	ID       int
	Name     string
	Nick     string `json:",omitempty"`
	Username string
	Password string
	Avatar   string    `json:",omitempty"` //path in /pic/<id>/
	Stories  []int     `json:",omitempty"`
	Bio      Biography `json:",omitempty"`
}

type Users struct {
	Users []*User
}

type Author struct {
	User   User
	Subbed []int
	Favs   []int
}

type Biography struct {
	Content     string
	Lastupdated TimeandDate
}

type Session struct {
	UserID  string
	Expires time.Time
}

// settings

type SettingChange struct {
	Type       int //0: name, 1:nickname, 2: Username, 3: Password, 4: Avatar, 5:	Biography
	NameChange string
	NickChange string
	UNchange   string
	PWchange   string
	Avatar     string // path to avatar file after upload
	BioChange  string
}

// template construct

type TopLayer struct {
	TP             string
	Inception      bool        `json:"i,omitempty"`
	I              *sublayer   `json:",omitempty"`
	MultiInception bool        `json:"mi,omitempty"`
	MI             []*sublayer `json:",omitempty"`
}

type sublayer struct {
	TP             string
	Inception      bool        `json:"i,omitempty"`
	I              *sublayer   `json:",omitempty"`
	IsMD           bool        `json:"isMD,omitempty"`
	MultiInception bool        `json:"mi,omitempty"`
	MI             []*sublayer `json:",omitempty"`
}
