<a href="/show/login">Login page</a>
<a href="/show/signup">Sign up page</a>
<a href="/show/user">User page</a>
<a href="/show/Settings">Settings page</a>
<a href="/show/Story">Story page</a>
<a href="/show/mainpage">Main page</a>
<a href="/show/reviewpage">Induvidual review page</a>
func ShowRoom(w http.ResponseWriter, r *http.Request) {
	page := strings.Replace(r.URL.Path, "/show/", "", 1)
	var showroot = fileroot + "/show/"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!!!!!!!!!!!!!!!!\nPANICKED:", r)
			w.Write([]byte("Critical error when displaying web-page: " + r.(error).Error()))
		}
	}()

	switch page {
	case "login":
		w.Write(HandlePage(showroot + "Login.icfg"))
		//nd
	case "signup":
		w.Write(HandlePage(showroot + "Signup.icfg"))
		//nd
	case "user":
		w.Write(HandlePage(showroot + "User.icfg"))
		//nd
	case "settings":
		w.Write(HandlePage(showroot + "Settings.icfg"))
		//nd
	case "story":
		w.Write(HandlePage(showroot + "Story.icfg"))
		//nd
	case "mainpage":
		w.Write(HandlePage(showroot + "Mainpage.icfg"))
		//nd
	case "reviewpage":
		w.Write(HandlePage(showroot + "Review.icfg"))
		//nd
	case "":
		w.Write(HandlePage(showroot + "Initial.icfg"))
		//nd
	}
	//forum, maybe?
}
