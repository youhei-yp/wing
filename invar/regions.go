// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package invar

// Region regions data
type Region struct {
	Code     string
	Phone    string
	TimeDiff string
	EnName   string
	CnName   string
}

var (
	Angola              = &Region{"AO", "244", "-7", "Angola", "安哥拉"}
	Afghanistan         = &Region{"AF", "93", "0", "Afghanistan", "阿富汗"}
	Albania             = &Region{"AL", "355", "-7", "Albania", "阿尔巴尼亚"}
	Algeria             = &Region{"DZ", "213", "-8", "Algeria", "阿尔及利亚"}
	Andorra             = &Region{"AD", "376", "-8", "Andorra", "安道尔共和国"}
	Anguilla            = &Region{"AI", "1264", "-12", "Anguilla", "安圭拉岛"}
	AntiguaBarbuda      = &Region{"AG", "1268", "-12", "Antigua and Barbuda", "安提瓜和巴布达"}
	Argentina           = &Region{"AR", "54", "-11", "Argentina", "阿根廷"}
	Armenia             = &Region{"AM", "374", "-6", "Armenia", "亚美尼亚"}
	Ascension           = &Region{"", "247", "-8", "Ascension", "阿森松"}
	Australia           = &Region{"AU", "61", "2", "Australia", "澳大利亚"}
	Austria             = &Region{"AT", "43", "-7", "Austria", "奥地利"}
	Azerbaijan          = &Region{"AZ", "994", "-5", "Azerbaijan", "阿塞拜疆"}
	Bahamas             = &Region{"BS", "1242", "-13", "Bahamas", "巴哈马"}
	Bahrain             = &Region{"BH", "973", "-5", "Bahrain", "巴林"}
	Bangladesh          = &Region{"BD", "880", "-2", "Bangladesh", "孟加拉国"}
	Barbados            = &Region{"BB", "1246", "-12", "Barbados", "巴巴多斯"}
	Belarus             = &Region{"BY", "375", "-6", "Belarus", "白俄罗斯"}
	Belgium             = &Region{"BE", "32", "-7", "Belgium", "比利时"}
	Belize              = &Region{"BZ", "501", "-14", "Belize", "伯利兹"}
	Benin               = &Region{"BJ", "229", "-7", "Benin", "贝宁"}
	BermudaIs           = &Region{"BM", "1441", "-12", "Bermuda Is.", "百慕大群岛"}
	Bolivia             = &Region{"BO", "591", "-12", "Bolivia ", "玻利维亚"}
	Botswana            = &Region{"BW", "267", "-6", "Botswana", "博茨瓦纳"}
	Brazil              = &Region{"BR", "55", "-11", "Brazil", "巴西"}
	Brunei              = &Region{"BN", "673", "0", "Brunei", "文莱"}
	Bulgaria            = &Region{"BG", "359", "-6", "Bulgaria", "保加利亚"}
	BurkinaFaso         = &Region{"BF", "226", "-8", "Burkina-faso", "布基纳法索"}
	Burma               = &Region{"MM", "95", "-1.3", "Burma", "缅甸"}
	Burundi             = &Region{"BI", "257", "-6", "Burundi", "布隆迪"}
	Cameroon            = &Region{"CM", "237", "-7", "Cameroon", "喀麦隆"}
	Canada              = &Region{"CA", "1", "-13", "Canada", "加拿大"}
	CaymanIs            = &Region{"", "1345", "-13", "Cayman Is.", "开曼群岛"}
	CentralAfricanRep   = &Region{"CF", "236", "-7", "Central African Republic", "中非共和国"}
	Chad                = &Region{"TD", "235", "-7", "Chad", "乍得"}
	Chile               = &Region{"CL", "56", "-13", "Chile", "智利"}
	China               = &Region{"CN", "86", "0", "China", "中国"}
	Colombia            = &Region{"CO", "57", "0", "Colombia", "哥伦比亚"}
	Congo               = &Region{"CG", "242", "-7", "Congo", "刚果"}
	CookIs              = &Region{"CK", "682", "-18.3", "Cook Is.", "库克群岛"}
	CostaRica           = &Region{"CR", "506", "-14", "Costa Rica", "哥斯达黎加"}
	Cuba                = &Region{"CU", "53", "-13", "Cuba", "古巴"}
	Cyprus              = &Region{"CY", "357", "-6", "Cyprus", "塞浦路斯"}
	CzechRep            = &Region{"CZ", "420", "-7", "Czech Republic", "捷克"}
	Denmark             = &Region{"DK", "45", "-7", "Denmark", "丹麦"}
	Djibouti            = &Region{"DJ", "253", "-5", "Djibouti", "吉布提"}
	DominicaRep         = &Region{"DO", "1890", "-13", "Dominica Rep.", "多米尼加共和国"}
	Ecuador             = &Region{"EC", "593", "-13", "Ecuador", "厄瓜多尔"}
	Egypt               = &Region{"EG", "20", "-6", "Egypt", "埃及"}
	EISalvador          = &Region{"SV", "503", "-14", "EI Salvador", "萨尔瓦多"}
	Estonia             = &Region{"EE", "372", "-5", "Estonia", "爱沙尼亚"}
	Ethiopia            = &Region{"ET", "251", "-5", "Ethiopia", "埃塞俄比亚"}
	Fiji                = &Region{"FJ", "679", "4", "Fiji", "斐济"}
	Finland             = &Region{"FI", "358", "-6", "Finland", "芬兰"}
	France              = &Region{"FR", "33", "-8", "France", "法国"}
	FrenchGuiana        = &Region{"GF", "594", "-12", "French Guiana", "法属圭亚那"}
	Gabon               = &Region{"GA", "241", "-7", "Gabon", "加蓬"}
	Gambia              = &Region{"GM", "220", "-8", "Gambia", "冈比亚"}
	Georgia             = &Region{"GE", "995", "0", "Georgia", "格鲁吉亚"}
	Germany             = &Region{"DE", "49", "-7", "Germany", "德国"}
	Ghana               = &Region{"GH", "233", "-8", "Ghana", "加纳"}
	Gibraltar           = &Region{"GI", "350", "-8", "Gibraltar", "直布罗陀"}
	Greece              = &Region{"GR", "30", "-6", "Greece", "希腊"}
	Grenada             = &Region{"GD", "1809", "-14", "Grenada", "格林纳达"}
	Guam                = &Region{"GU", "1671", "2", "Guam", "关岛"}
	Guatemala           = &Region{"GT", "502", "-14", "Guatemala", "危地马拉"}
	Guinea              = &Region{"GN", "224", "-8", "Guinea", "几内亚"}
	Guyana              = &Region{"GY", "592", "-11", "Guyana", "圭亚那"}
	Haiti               = &Region{"HT", "509", "-13", "Haiti", "海地"}
	Honduras            = &Region{"HN", "504", "-14", "Honduras", "洪都拉斯"}
	Hongkong            = &Region{"HK", "852", "0", "Hongkong", "香港"}
	Hungary             = &Region{"HU", "36", "-7", "Hungary", "匈牙利"}
	Iceland             = &Region{"IS", "354", "-9", "Iceland", "冰岛"}
	India               = &Region{"IN", "91", "-2.3", "India", "印度"}
	Indonesia           = &Region{"ID", "62", "-0.3", "Indonesia", "印度尼西亚"}
	Iran                = &Region{"IR", "98", "-4.3", "Iran", "伊朗"}
	Iraq                = &Region{"IQ", "964", "-5", "Iraq", "伊拉克"}
	Ireland             = &Region{"IE", "353", "-4.3", "Ireland", "爱尔兰"}
	Israel              = &Region{"IL", "972", "-6", "Israel", "以色列"}
	Italy               = &Region{"IT", "39", "-7", "Italy", "意大利"}
	IvoryCoast          = &Region{"", "225", "-6", "Ivory Coast", "科特迪瓦"}
	Jamaica             = &Region{"JM", "1876", "-12", "Jamaica", "牙买加"}
	Japan               = &Region{"JP", "81", "1", "Japan", "日本"}
	Jordan              = &Region{"JO", "962", "-6", "Jordan", "约旦"}
	Kampuchea           = &Region{"KH", "855", "-1", "Kampuchea (Cambodia )", "柬埔寨"}
	Kazakstan           = &Region{"KZ", "327", "-5", "Kazakstan", "哈萨克斯坦"}
	Kenya               = &Region{"KE", "254", "-5", "Kenya", "肯尼亚"}
	Korea               = &Region{"KR", "82", "1", "Korea", "韩国"}
	Kuwait              = &Region{"KW", "965", "-5", "Kuwait", "科威特"}
	Kyrgyzstan          = &Region{"KG", "331", "-5", "Kyrgyzstan", "吉尔吉斯坦"}
	Laos                = &Region{"LA", "856", "-1", "Laos", "老挝"}
	Latvia              = &Region{"LV", "371", "-5", "Latvia", "拉脱维亚"}
	Lebanon             = &Region{"LB", "961", "-6", "Lebanon", "黎巴嫩"}
	Lesotho             = &Region{"LS", "266", "-6", "Lesotho", "莱索托"}
	Liberia             = &Region{"LR", "231", "-8", "Liberia", "利比里亚"}
	Libya               = &Region{"LY", "218", "-6", "Libya", "利比亚"}
	Liechtenstein       = &Region{"LI", "423", "-7", "Liechtenstein", "列支敦士登"}
	Lithuania           = &Region{"LT", "370", "-5", "Lithuania", "立陶宛"}
	Luxembourg          = &Region{"LU", "352", "-7", "Luxembourg", "卢森堡"}
	Macao               = &Region{"MO", "853", "0", "Macao", "澳门"}
	Madagascar          = &Region{"MG", "261", "-5", "Madagascar", "马达加斯加"}
	Malawi              = &Region{"MW", "265", "-6", "Malawi", "马拉维"}
	Malaysia            = &Region{"MY", "60", "-0.5", "Malaysia", "马来西亚"}
	Maldives            = &Region{"MV", "960", "-7", "Maldives", "马尔代夫"}
	Mali                = &Region{"ML", "223", "-8", "Mali", "马里"}
	Malta               = &Region{"MT", "356", "-7", "Malta", "马耳他"}
	MarianaIs           = &Region{"", "1670", "1", "Mariana Is", "马里亚那群岛"}
	Martinique          = &Region{"", "596", "-12", "Martinique", "马提尼克"}
	Mauritius           = &Region{"MU", "230", "-4", "Mauritius", "毛里求斯"}
	Mexico              = &Region{"MX", "52", "-15", "Mexico", "墨西哥"}
	MoldovaRep          = &Region{"MD", "373", "-5", "Moldova, Republic of", "摩尔多瓦"}
	Monaco              = &Region{"MC", "377", "-7", "Monaco", "摩纳哥"}
	Mongolia            = &Region{"MN", "976", "0", "Mongolia", "蒙古"}
	MontserratIs        = &Region{"MS", "1664", "-12", "Montserrat Is", "蒙特塞拉特岛"}
	Morocco             = &Region{"MA", "212", "-6", "Morocco", "摩洛哥"}
	Mozambique          = &Region{"MZ", "258", "-6", "Mozambique", "莫桑比克"}
	Namibia             = &Region{"NA", "264", "-7", "Namibia", "纳米比亚"}
	Nauru               = &Region{"NR", "674", "4", "Nauru", "瑙鲁"}
	Nepal               = &Region{"NP", "977", "-2.3", "Nepal", "尼泊尔"}
	NetheriandsAntilles = &Region{"", "599", "-12", "Netheriands Antilles", "荷属安的列斯"}
	Netherlands         = &Region{"NL", "31", "-7", "Netherlands", "荷兰"}
	NewZealand          = &Region{"NZ", "64", "4", "New Zealand", "新西兰"}
	Nicaragua           = &Region{"NI", "505", "-14", "Nicaragua", "尼加拉瓜"}
	Niger               = &Region{"NE", "227", "-8", "Niger", "尼日尔"}
	Nigeria             = &Region{"NG", "234", "-7", "Nigeria", "尼日利亚"}
	NorthKorea          = &Region{"KP", "850", "1", "North Korea", "朝鲜"}
	Norway              = &Region{"NO", "47", "-7", "Norway", "挪威"}
	Oman                = &Region{"OM", "968", "-4", "Oman", "阿曼"}
	Pakistan            = &Region{"PK", "92", "-2.3", "Pakistan", "巴基斯坦"}
	Panama              = &Region{"PA", "507", "-13", "Panama", "巴拿马"}
	PapuaNewCuinea      = &Region{"PG", "675", "2", "Papua New Cuinea", "巴布亚新几内亚"}
	Paraguay            = &Region{"PY", "595", "-12", "Paraguay", "巴拉圭"}
	Peru                = &Region{"PE", "51", "-13", "Peru", "秘鲁"}
	Philippines         = &Region{"PH", "63", "0", "Philippines", "菲律宾"}
	Poland              = &Region{"PL", "48", "-7", "Poland", "波兰"}
	FrenchPolynesia     = &Region{"PF", "689", "3", "French Polynesia", "法属玻利尼西亚"}
	Portugal            = &Region{"PT", "351", "-8", "Portugal", "葡萄牙"}
	PuertoRico          = &Region{"PR", "1787", "-12", "Puerto Rico", "波多黎各"}
	Qatar               = &Region{"QA", "974", "-5", "Qatar", "卡塔尔"}
	Reunion             = &Region{"", "262", "-4", "Reunion", "留尼旺"}
	Romania             = &Region{"RO", "40", "-6", "Romania", "罗马尼亚"}
	Russia              = &Region{"RU", "7", "-5", "Russia", "俄罗斯"}
	SaintLueia          = &Region{"LC", "1758", "-12", "Saint Lueia", "圣卢西亚"}
	SaintVincent        = &Region{"VC", "1784", "-12", "Saint Vincent", "圣文森特岛"}
	SamoaEastern        = &Region{"", "684", "-19", "Samoa Eastern", "东萨摩亚(美)"}
	SamoaWestern        = &Region{"", "685", "-19", "Samoa Western", "西萨摩亚"}
	SanMarino           = &Region{"SM", "378", "-7", "San Marino", "圣马力诺"}
	SaoTomePrincipe     = &Region{"ST", "239", "-8", "Sao Tome and Principe", "圣多美和普林西比"}
	SaudiArabia         = &Region{"SA", "966", "-5", "Saudi Arabia", "沙特阿拉伯"}
	Senegal             = &Region{"SN", "221", "-8", "Senegal", "塞内加尔"}
	Seychelles          = &Region{"SC", "248", "-4", "Seychelles", "塞舌尔"}
	SierraLeone         = &Region{"SL", "232", "-8", "Sierra Leone", "塞拉利昂"}
	Singapore           = &Region{"SG", "65", "0.3", "Singapore", "新加坡"}
	Slovakia            = &Region{"SK", "421", "-7", "Slovakia", "斯洛伐克"}
	Slovenia            = &Region{"SI", "386", "-7", "Slovenia", "斯洛文尼亚"}
	SolomonIs           = &Region{"SB", "677", "3", "Solomon Is", "所罗门群岛"}
	Somali              = &Region{"SO", "252", "-5", "Somali", "索马里"}
	SouthAfrica         = &Region{"ZA", "27", "-6", "South Africa", "南非"}
	Spain               = &Region{"ES", "34", "-8", "Spain", "西班牙"}
	SriLanka            = &Region{"LK", "94", "0", "Sri Lanka", "斯里兰卡"}
	StLucia             = &Region{"LC", "1758", "-12", "St.Lucia", "圣卢西亚"}
	StVincent           = &Region{"VC", "1784", "-12", "St.Vincent", "圣文森特"}
	Sudan               = &Region{"SD", "249", "-6", "Sudan", "苏丹"}
	Suriname            = &Region{"SR", "597", "-11.3", "Suriname", "苏里南"}
	Swaziland           = &Region{"SZ", "268", "-6", "Swaziland", "斯威士兰"}
	Sweden              = &Region{"SE", "46", "-7", "Sweden", "瑞典"}
	Switzerland         = &Region{"CH", "41", "-7", "Switzerland", "瑞士"}
	Syria               = &Region{"SY", "963", "-6", "Syria", "叙利亚"}
	Taiwan              = &Region{"TW", "886", "0", "Taiwan", "台湾省"}
	Tajikstan           = &Region{"TJ", "992", "-5", "Tajikstan", "塔吉克斯坦"}
	Tanzania            = &Region{"TZ", "255", "-5", "Tanzania", "坦桑尼亚"}
	Thailand            = &Region{"TH", "66", "-1", "Thailand", "泰国"}
	Togo                = &Region{"TG", "228", "-8", "Togo", "多哥"}
	Tonga               = &Region{"TO", "676", "4", "Tonga", "汤加"}
	TrinidadTobago      = &Region{"TT", "1809", "-12", "Trinidad and Tobago", "特立尼达和多巴哥"}
	Tunisia             = &Region{"TN", "216", "-7", "Tunisia", "突尼斯"}
	Turkey              = &Region{"TR", "90", "-6", "Turkey", "土耳其"}
	Turkmenistan        = &Region{"TM", "993", "-5", "Turkmenistan", "土库曼斯坦"}
	Uganda              = &Region{"UG", "256", "-5", "Uganda", "乌干达"}
	Ukraine             = &Region{"UA", "380", "-5", "Ukraine", "乌克兰"}
	UnitedArabEmirates  = &Region{"AE", "971", "-4", "United Arab Emirates", "阿拉伯联合酋长国"}
	UnitedKiongdom      = &Region{"GB", "44", "-8", "United Kiongdom ", "英国"}
	USA                 = &Region{"US", "1", "-13", "United States of America", "美国"}
	Uruguay             = &Region{"UY", "598", "-10.3", "Uruguay", "乌拉圭"}
	Uzbekistan          = &Region{"UZ", "233", "-5", "Uzbekistan", "乌兹别克斯坦"}
	Venezuela           = &Region{"VE", "58", "-12.3", "Venezuela", "委内瑞拉"}
	Vietnam             = &Region{"VN", "84", "-1", "Vietnam", "越南"}
	Yemen               = &Region{"YE", "967", "-5", "Yemen", "也门"}
	Yugoslavia          = &Region{"YU", "381", "-7", "Yugoslavia", "南斯拉夫"}
	Zimbabwe            = &Region{"ZW", "263", "-6", "Zimbabwe", "津巴布韦"}
	Zaire               = &Region{"ZR", "243", "-7", "Zaire", "扎伊尔"}
	Zambia              = &Region{"ZM", "260", "-6", "Zambia", "赞比亚"}
)
