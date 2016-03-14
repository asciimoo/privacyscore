package penalty

type Score int

type PenaltyType int

const (
	P_COOKIE            PenaltyType = 1
	P_EXTERNAL_LINK     PenaltyType = 2
	P_HTTP_LINK         PenaltyType = 3
	P_EXTERNAL_RESOURCE PenaltyType = 4
	P_NO_HTTPS          PenaltyType = 5
	P_JS                PenaltyType = 6
	P_NO_SECURE_HEADER  PenaltyType = 7
)

type Penalty struct {
	Description string
	DetailLink  string
	Notes       []string
	Value       Score
}

func New(p PenaltyType, value Score) *Penalty {
	desc := ""
	link := ""
	switch p {
	case P_COOKIE:
		desc = "Automatically sets cookies"
		link = "https://en.wikipedia.org/wiki/Internet_privacy#HTTP_cookies"
	case P_EXTERNAL_LINK:
		desc = "Leaks HTTP referrer to foreign host"
		link = "https://randomoracle.wordpress.com/2013/11/23/privacy-and-http-referer-header-12/"
	case P_HTTP_LINK:
		desc = "Has link to unencrypted service (no HTTPS)"
		link = "https://en.wikipedia.org/wiki/HTTP_Secure"
	case P_EXTERNAL_RESOURCE:
		desc = "Loads external resource"
		link = "https://jonathanmayer.org/papers_data/trackingsurvey12.pdf"
	case P_NO_HTTPS:
		desc = "Uses unencrypted transport layer (no HTTPS)"
		link = "https://en.wikipedia.org/wiki/HTTP_Secure"
	case P_JS:
		desc = "Uses JavaScript"
		link = "todo"
	case P_NO_SECURE_HEADER:
		desc = "Missing secure HTTP header"
		link = "https://scotthelme.co.uk/hardening-your-http-response-headers/"
	}
	return &Penalty{desc, link, make([]string, 0, 8), value}
}
