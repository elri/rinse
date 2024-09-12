package rinse

var LanguageTika = map[string]string{
	"da": "dan", // Danish
	"de": "deu", // German
	"et": "est", // Estonian
	"el": "ell", // Greek
	"en": "eng", // English
	"es": "spa", // Spanish
	"fi": "fin", // Finnish
	"fr": "fra", // French
	"hu": "hun", // Hungarian
	"is": "isl", // Icelandic
	"it": "ita", // Italian
	"nl": "nld", // Dutch
	"no": "nor", // Norwegian
	"pl": "pol", // Polish
	"pt": "por", // Portuguese
	"ru": "rus", // Russian
	"sv": "swe", // Swedish
	"th": "tha", // Thai
}

var LanguageCode = map[string]string{
	"afr":      "Afrikaans",             // 	afr.traineddata
	"amh":      "Amharic",               // 	amh.traineddata
	"ara":      "Arabic",                // 	ara.traineddata
	"asm":      "Assamese",              // 	asm.traineddata
	"aze":      "Azerbaijani",           // 	aze.traineddata
	"aze_cyrl": "Azerbaijani, Cyrillic", //  	aze_cyrl.traineddata
	"bel":      "Belarusian",            // 	bel.traineddata
	"ben":      "Bengali",               // 	ben.traineddata
	"bod":      "Tibetan",               // 	bod.traineddata
	"bos":      "Bosnian",               // 	bos.traineddata
	"bul":      "Bulgarian",             // 	bul.traineddata
	"cat":      "Catalan, Valencian",    // ; 	cat.traineddata
	"ceb":      "Cebuano",               // 	ceb.traineddata
	"ces":      "Czech",                 // 	ces.traineddata
	"chi_sim":  "Chinese, Simplified",   //  	chi_sim.traineddata
	"chi_tra":  "Chinese, Traditional",  //  	chi_tra.traineddata
	"chr":      "Cherokee",              // 	chr.traineddata
	"cym":      "Welsh",                 // 	cym.traineddata
	"dan":      "Danish",                // 	dan.traineddata
	"deu":      "German",                // 	deu.traineddata
	"dzo":      "Dzongkha",              // 	dzo.traineddata
	"ell":      "Greek, Modern",         // , Modern (1453-)	ell.traineddata
	"eng":      "English",               // 	eng.traineddata
	"enm":      "English, Old",          // , Middle (1100-1500)	enm.traineddata
	"epo":      "Esperanto",             // 	epo.traineddata
	"est":      "Estonian",              // 	est.traineddata
	"eus":      "Basque",                // 	eus.traineddata
	"fas":      "Persian",               // 	fas.traineddata
	"fin":      "Finnish",               // 	fin.traineddata
	"fra":      "French",                // 	fra.traineddata
	"frk":      "German, Fraktur",       //  	frk.traineddata
	"frm":      "French, Old",           // , Middle (ca. 1400-1600)	frm.traineddata
	"gle":      "Irish",                 // 	gle.traineddata
	"glg":      "Galician",              // 	glg.traineddata
	"grc":      "Greek, Ancient",        // , Ancient (-1453)	grc.traineddata
	"guj":      "Gujarati",              // 	guj.traineddata
	"hat":      "Haitian, Creole",       // ; Haitian Creole	hat.traineddata
	"heb":      "Hebrew",                // 	heb.traineddata
	"hin":      "Hindi",                 // 	hin.traineddata
	"hrv":      "Croatian",              // 	hrv.traineddata
	"hun":      "Hungarian",             // 	hun.traineddata
	"iku":      "Inuktitut",             // 	iku.traineddata
	"ind":      "Indonesian",            // 	ind.traineddata
	"isl":      "Icelandic",             // 	isl.traineddata
	"ita":      "Italian",               // 	ita.traineddata
	"ita_old":  "Italian, Old",          //  – Old	ita_old.traineddata
	"jav":      "Javanese",              // 	jav.traineddata
	"jpn":      "Japanese",              // 	jpn.traineddata
	"kan":      "Kannada",               // 	kan.traineddata
	"kat":      "Georgian",              // 	kat.traineddata
	"kat_old":  "Georgian, Old",         //  – Old	kat_old.traineddata
	"kaz":      "Kazakh",                // 	kaz.traineddata
	"khm":      "Khmer, Central",        //  Khmer	khm.traineddata
	"kir":      "Kirghiz",               // ; Kyrgyz	kir.traineddata
	"kor":      "Korean",                // 	kor.traineddata
	"kur":      "Kurdish",               // 	kur.traineddata
	"lao":      "Lao",                   // 	lao.traineddata
	"lat":      "Latin",                 // 	lat.traineddata
	"lav":      "Latvian",               // 	lav.traineddata
	"lit":      "Lithuanian",            // 	lit.traineddata
	"mal":      "Malayalam",             // 	mal.traineddata
	"mar":      "Marathi",               // 	mar.traineddata
	"mkd":      "Macedonian",            // 	mkd.traineddata
	"mlt":      "Maltese",               // 	mlt.traineddata
	"msa":      "Malay",                 // 	msa.traineddata
	"mya":      "Burmese",               // 	mya.traineddata
	"nep":      "Nepali",                // 	nep.traineddata
	"nld":      "Dutch, Flemish",        // ; Flemish	nld.traineddata
	"nor":      "Norwegian",             // 	nor.traineddata
	"ori":      "Oriya",                 // 	ori.traineddata
	"pan":      "Panjabi",               // ; Punjabi	pan.traineddata
	"pol":      "Polish",                // 	pol.traineddata
	"por":      "Portuguese",            // 	por.traineddata
	"pus":      "Pushto",                // ; Pashto	pus.traineddata
	"ron":      "Romanian, Moldavian",   // ; Moldavian; Moldovan	ron.traineddata
	"rus":      "Russian",               // 	rus.traineddata
	"san":      "Sanskrit",              // 	san.traineddata
	"sin":      "Sinhala",               // ; Sinhalese	sin.traineddata
	"slk":      "Slovak",                // 	slk.traineddata
	"slv":      "Slovenian",             // 	slv.traineddata
	"spa":      "Spanish",               // ; Castilian	spa.traineddata
	"spa_old":  "Spanish, Old",          // ; Castilian Old	spa_old.traineddata
	"sqi":      "Albanian",              // 	sqi.traineddata
	"srp":      "Serbian",               // 	srp.traineddata
	"srp_latn": "Serbian, Latin",        //  Latin	srp_latn.traineddata
	"swa":      "Swahili",               // 	swa.traineddata
	"swe":      "Swedish",               // 	swe.traineddata
	"syr":      "Syriac",                // 	syr.traineddata
	"tam":      "Tamil",                 // 	tam.traineddata
	"tel":      "Telugu",                // 	tel.traineddata
	"tgk":      "Tajik",                 // 	tgk.traineddata
	"tgl":      "Tagalog",               // 	tgl.traineddata
	"tha":      "Thai",                  // 	tha.traineddata
	"tir":      "Tigrinya",              // 	tir.traineddata
	"tur":      "Turkish",               // 	tur.traineddata
	"uig":      "Uighur",                // ; Uyghur	uig.traineddata
	"ukr":      "Ukrainian",             // 	ukr.traineddata
	"urd":      "Urdu",                  // 	urd.traineddata
	"uzb":      "Uzbek",                 // 	uzb.traineddata
	"uzb_cyrl": "Uzbek, Cyrillic",       //  Cyrillic	uzb_cyrl.traineddata
	"vie":      "Vietnamese",            // 	vie.traineddata
	"yid":      "Yiddish",               // 	yid.traineddata
}

func (rns *Rinse) LanguageName(code string) string {
	return LanguageCode[code]
}
