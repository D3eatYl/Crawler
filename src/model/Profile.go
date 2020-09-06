package model

type LeakZoneProfile struct {
	Name			string
	Url				string
	UserTitle		string
	MemberFor		string
	Posts			string
	Threads			string
	Credits			string
	LeecherStatus	string
}

type GreySecProfile struct {
	Name			string
	Url 			string
	SmallText		string
	Posts			string
	Threads			string
	Joined			string
	Reputation		string
}

type RstForumsProfile struct {
	Name			string
	Url				string
	Member			string
	ContentCount	string
	Joined			string
	DaysWon			string
	Reputation		string
}