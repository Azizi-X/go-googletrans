package translator

var defaultUserAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv109.0) Gecko/20100101 Firefox/115.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64; rv109.0) Gecko/20100101 Firefox/115.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv109.0) Gecko/20100101 Firefox/115.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.2 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; rv109.0) Gecko/20100101 Firefox/115.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv109.0) Gecko/20100101 Firefox/116.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv109.0) Gecko/20100101 Firefox/115.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.1 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv109.0) Gecko/20100101 Firefox/114.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64; rv102.0) Gecko/20100101 Firefox/102.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.67",
	"Mozilla/5.0 (X11; Linux x86_64; rv109.0) Gecko/20100101 Firefox/114.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 OPR/100.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv102.0) Gecko/20100101 Firefox/102.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv109.0) Gecko/20100101 Firefox/114.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.79",
	"Mozilla/5.0 (X11; Linux x86_64; rv109.0) Gecko/20100101 Firefox/116.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.75 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 OPR/99.0.0.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Safari/605.1.15",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; rv102.0) Gecko/20100101 Firefox/102.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5666.197 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; rv109.0) Gecko/20100101 Firefox/116.0",
	"Mozilla/5.0 (Windows NT 10.0; rv114.0) Gecko/20100101 Firefox/114.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv109.0) Gecko/20100101 Firefox/114.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6.1 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv109.0) Gecko/20100101 Firefox/116.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.86",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv109.0) Gecko/20100101 Firefox/113.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 YaBrowser/23.5.4.674 Yowser/2.5 Safari/537.36",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv109.0) Gecko/20100101 Firefox/116.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
}

var defaultServiceUrls = []string{
	"translate.google.ac", "translate.google.ad", "translate.google.ae",
	"translate.google.al", "translate.google.am", "translate.google.as",
	"translate.google.at", "translate.google.az", "translate.google.ba",
	"translate.google.be", "translate.google.bf", "translate.google.bg",
	"translate.google.bi", "translate.google.bj", "translate.google.bs",
	"translate.google.bt", "translate.google.by", "translate.google.ca",
	"translate.google.cat", "translate.google.cc", "translate.google.cd",
	"translate.google.cf", "translate.google.cg", "translate.google.ch",
	"translate.google.ci", "translate.google.cl", "translate.google.cm",
	"translate.google.cn", "translate.google.co.ao", "translate.google.co.bw",
	"translate.google.co.ck", "translate.google.co.cr", "translate.google.co.id",
	"translate.google.co.il", "translate.google.co.in", "translate.google.co.jp",
	"translate.google.co.ke", "translate.google.co.kr", "translate.google.co.ls",
	"translate.google.co.ma", "translate.google.co.mz", "translate.google.co.nz",
	"translate.google.co.th", "translate.google.co.tz", "translate.google.co.ug",
	"translate.google.co.uk", "translate.google.co.uz", "translate.google.co.ve",
	"translate.google.co.vi", "translate.google.co.za", "translate.google.co.zm",
	"translate.google.co.zw", "translate.google.co", "translate.google.com.af",
	"translate.google.com.ag", "translate.google.com.ai", "translate.google.com.ar",
	"translate.google.com.au", "translate.google.com.bd", "translate.google.com.bh",
	"translate.google.com.bn", "translate.google.com.bo", "translate.google.com.br",
	"translate.google.com.bz", "translate.google.com.co", "translate.google.com.cu",
	"translate.google.com.cy", "translate.google.com.do", "translate.google.com.ec",
	"translate.google.com.eg", "translate.google.com.et", "translate.google.com.fj",
	"translate.google.com.gh", "translate.google.com.gi", "translate.google.com.gt",
	"translate.google.com.hk", "translate.google.com.jm", "translate.google.com.kh",
	"translate.google.com.kw", "translate.google.com.lb", "translate.google.com.lc",
	"translate.google.com.ly", "translate.google.com.mm", "translate.google.com.mt",
	"translate.google.com.mx", "translate.google.com.my", "translate.google.com.na",
	"translate.google.com.ng", "translate.google.com.ni", "translate.google.com.np",
	"translate.google.com.om", "translate.google.com.pa", "translate.google.com.pe",
	"translate.google.com.pg", "translate.google.com.ph", "translate.google.com.pk",
	"translate.google.com.pr", "translate.google.com.py", "translate.google.com.qa",
	"translate.google.com.sa", "translate.google.com.sb", "translate.google.com.sg",
	"translate.google.com.sl", "translate.google.com.sv", "translate.google.com.tj",
	"translate.google.com.tr", "translate.google.com.tw", "translate.google.com.ua",
	"translate.google.com.uy", "translate.google.com.vc", "translate.google.com.vn",
	"translate.google.com", "translate.google.cv", "translate.google.cx",
	"translate.google.cz", "translate.google.de", "translate.google.dj",
	"translate.google.dk", "translate.google.dm", "translate.google.dz",
	"translate.google.ee", "translate.google.es", "translate.google.eu",
	"translate.google.fi", "translate.google.fm", "translate.google.fr",
	"translate.google.ga", "translate.google.ge", "translate.google.gf",
	"translate.google.gg", "translate.google.gl", "translate.google.gm",
	"translate.google.gp", "translate.google.gr", "translate.google.gy",
	"translate.google.hn", "translate.google.hr", "translate.google.ht",
	"translate.google.hu", "translate.google.ie", "translate.google.im",
	"translate.google.io", "translate.google.iq", "translate.google.is",
	"translate.google.it", "translate.google.je", "translate.google.jo",
	"translate.google.kg", "translate.google.ki", "translate.google.kz",
	"translate.google.la", "translate.google.li", "translate.google.lk",
	"translate.google.lt", "translate.google.lu", "translate.google.lv",
	"translate.google.md", "translate.google.me", "translate.google.mg",
	"translate.google.mk", "translate.google.ml", "translate.google.mn",
	"translate.google.ms", "translate.google.mu", "translate.google.mv",
	"translate.google.mw", "translate.google.ne", "translate.google.nf",
	"translate.google.nl", "translate.google.no", "translate.google.nr",
	"translate.google.nu", "translate.google.pl", "translate.google.pn",
	"translate.google.ps", "translate.google.pt", "translate.google.ro",
	"translate.google.rs", "translate.google.ru", "translate.google.rw",
	"translate.google.sc", "translate.google.se", "translate.google.sh",
	"translate.google.si", "translate.google.sk", "translate.google.sm",
	"translate.google.sn", "translate.google.so", "translate.google.sr",
	"translate.google.st", "translate.google.td", "translate.google.tg",
	"translate.google.tk", "translate.google.tl", "translate.google.tm",
	"translate.google.tn", "translate.google.to", "translate.google.tt",
	"translate.google.us", "translate.google.vg", "translate.google.vu",
	"translate.google.ws",
}

var Languages = map[string]string{
	"auto":  "auto",
	"af":    "afrikaans",
	"sq":    "albanian",
	"am":    "amharic",
	"ar":    "arabic",
	"hy":    "armenian",
	"az":    "azerbaijani",
	"eu":    "basque",
	"be":    "belarusian",
	"bn":    "bengali",
	"bs":    "bosnian",
	"bg":    "bulgarian",
	"ca":    "catalan",
	"ceb":   "cebuano",
	"ny":    "chichewa",
	"zh-cn": "chinese (simplified)",
	"zh-tw": "chinese (traditional)",
	"co":    "corsican",
	"hr":    "croatian",
	"cs":    "czech",
	"da":    "danish",
	"nl":    "dutch",
	"en":    "english",
	"eo":    "esperanto",
	"et":    "estonian",
	"tl":    "filipino",
	"fi":    "finnish",
	"fr":    "french",
	"fy":    "frisian",
	"gl":    "galician",
	"ka":    "georgian",
	"de":    "german",
	"el":    "greek",
	"gu":    "gujarati",
	"ht":    "haitian creole",
	"ha":    "hausa",
	"haw":   "hawaiian",
	"iw":    "hebrew",
	"he":    "hebrew",
	"hi":    "hindi",
	"hmn":   "hmong",
	"hu":    "hungarian",
	"is":    "icelandic",
	"ig":    "igbo",
	"id":    "indonesian",
	"ga":    "irish",
	"it":    "italian",
	"ja":    "japanese",
	"jw":    "javanese",
	"kn":    "kannada",
	"kk":    "kazakh",
	"km":    "khmer",
	"ko":    "korean",
	"ku":    "kurdish (kurmanji)",
	"ky":    "kyrgyz",
	"lo":    "lao",
	"la":    "latin",
	"lv":    "latvian",
	"lt":    "lithuanian",
	"lb":    "luxembourgish",
	"mk":    "macedonian",
	"mg":    "malagasy",
	"ms":    "malay",
	"ml":    "malayalam",
	"mt":    "maltese",
	"mi":    "maori",
	"mr":    "marathi",
	"mn":    "mongolian",
	"my":    "myanmar (burmese)",
	"ne":    "nepali",
	"no":    "norwegian",
	"or":    "odia",
	"ps":    "pashto",
	"fa":    "persian",
	"pl":    "polish",
	"pt":    "portuguese",
	"pa":    "punjabi",
	"ro":    "romanian",
	"ru":    "russian",
	"sm":    "samoan",
	"gd":    "scots gaelic",
	"sr":    "serbian",
	"st":    "sesotho",
	"sn":    "shona",
	"sd":    "sindhi",
	"si":    "sinhala",
	"sk":    "slovak",
	"sl":    "slovenian",
	"so":    "somali",
	"es":    "spanish",
	"su":    "sundanese",
	"sw":    "swahili",
	"sv":    "swedish",
	"tg":    "tajik",
	"ta":    "tamil",
	"te":    "telugu",
	"th":    "thai",
	"tr":    "turkish",
	"uk":    "ukrainian",
	"ur":    "urdu",
	"ug":    "uyghur",
	"uz":    "uzbek",
	"vi":    "vietnamese",
	"cy":    "welsh",
	"xh":    "xhosa",
	"yi":    "yiddish",
	"yo":    "yoruba",
	"zu":    "zulu",
}

func reverseMap(m map[string]string) map[string]string {
	reversed := make(map[string]string, len(m))
	for k, v := range m {
		reversed[v] = k
	}
	return reversed
}

var LanguagesReversed = reverseMap(Languages)
