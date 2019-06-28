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
	CnName   string
}

const (
	Cnt_Angola              = "Angola"
	Cnt_Afghanistan         = "Afghanistan"
	Cnt_Albania             = "Albania"
	Cnt_Algeria             = "Algeria"
	Cnt_Andorra             = "Andorra"
	Cnt_Anguilla            = "Anguilla"
	Cnt_AntiguaBarbuda      = "Antigua and Barbuda"
	Cnt_Argentina           = "Argentina"
	Cnt_Armenia             = "Armenia"
	Cnt_Ascension           = "Ascension"
	Cnt_Australia           = "Australia"
	Cnt_Austria             = "Austria"
	Cnt_Azerbaijan          = "Azerbaijan"
	Cnt_Bahamas             = "Bahamas"
	Cnt_Bahrain             = "Bahrain"
	Cnt_Bangladesh          = "Bangladesh"
	Cnt_Barbados            = "Barbados"
	Cnt_Belarus             = "Belarus"
	Cnt_Belgium             = "Belgium"
	Cnt_Belize              = "Belize"
	Cnt_Benin               = "Benin"
	Cnt_BermudaIs           = "Bermuda Is."
	Cnt_Bolivia             = "Bolivia "
	Cnt_Botswana            = "Botswana"
	Cnt_Brazil              = "Brazil"
	Cnt_Brunei              = "Brunei"
	Cnt_Bulgaria            = "Bulgaria"
	Cnt_BurkinaFaso         = "Burkina-faso"
	Cnt_Burma               = "Burma"
	Cnt_Burundi             = "Burundi"
	Cnt_Cameroon            = "Cameroon"
	Cnt_Canada              = "Canada"
	Cnt_CaymanIs            = "Cayman Is."
	Cnt_CentralAfricanRep   = "Central African Republic"
	Cnt_Chad                = "Chad"
	Cnt_Chile               = "Chile"
	Cnt_China               = "China"
	Cnt_Colombia            = "Colombia"
	Cnt_Congo               = "Congo"
	Cnt_CookIs              = "Cook Is."
	Cnt_CostaRica           = "Costa Rica"
	Cnt_Cuba                = "Cuba"
	Cnt_Cyprus              = "Cyprus"
	Cnt_CzechRep            = "Czech Republic"
	Cnt_Denmark             = "Denmark"
	Cnt_Djibouti            = "Djibouti"
	Cnt_DominicaRep         = "Dominica Rep."
	Cnt_Ecuador             = "Ecuador"
	Cnt_Egypt               = "Egypt"
	Cnt_EISalvador          = "EI Salvador"
	Cnt_Estonia             = "Estonia"
	Cnt_Ethiopia            = "Ethiopia"
	Cnt_Fiji                = "Fiji"
	Cnt_Finland             = "Finland"
	Cnt_France              = "France"
	Cnt_FrenchGuiana        = "French Guiana"
	Cnt_Gabon               = "Gabon"
	Cnt_Gambia              = "Gambia"
	Cnt_Georgia             = "Georgia"
	Cnt_Germany             = "Germany"
	Cnt_Ghana               = "Ghana"
	Cnt_Gibraltar           = "Gibraltar"
	Cnt_Greece              = "Greece"
	Cnt_Grenada             = "Grenada"
	Cnt_Guam                = "Guam"
	Cnt_Guatemala           = "Guatemala"
	Cnt_Guinea              = "Guinea"
	Cnt_Guyana              = "Guyana"
	Cnt_Haiti               = "Haiti"
	Cnt_Honduras            = "Honduras"
	Cnt_Hongkong            = "Hongkong"
	Cnt_Hungary             = "Hungary"
	Cnt_Iceland             = "Iceland"
	Cnt_India               = "India"
	Cnt_Indonesia           = "Indonesia"
	Cnt_Iran                = "Iran"
	Cnt_Iraq                = "Iraq"
	Cnt_Ireland             = "Ireland"
	Cnt_Israel              = "Israel"
	Cnt_Italy               = "Italy"
	Cnt_IvoryCoast          = "Ivory Coast"
	Cnt_Jamaica             = "Jamaica"
	Cnt_Japan               = "Japan"
	Cnt_Jordan              = "Jordan"
	Cnt_Kampuchea           = "Kampuchea (Cambodia)"
	Cnt_Kazakstan           = "Kazakstan"
	Cnt_Kenya               = "Kenya"
	Cnt_Korea               = "Korea"
	Cnt_Kuwait              = "Kuwait"
	Cnt_Kyrgyzstan          = "Kyrgyzstan"
	Cnt_Laos                = "Laos"
	Cnt_Latvia              = "Latvia"
	Cnt_Lebanon             = "Lebanon"
	Cnt_Lesotho             = "Lesotho"
	Cnt_Liberia             = "Liberia"
	Cnt_Libya               = "Libya"
	Cnt_Liechtenstein       = "Liechtenstein"
	Cnt_Lithuania           = "Lithuania"
	Cnt_Luxembourg          = "Luxembourg"
	Cnt_Macao               = "Macao"
	Cnt_Madagascar          = "Madagascar"
	Cnt_Malawi              = "Malawi"
	Cnt_Malaysia            = "Malaysia"
	Cnt_Maldives            = "Maldives"
	Cnt_Mali                = "Mali"
	Cnt_Malta               = "Malta"
	Cnt_MarianaIs           = "Mariana Is"
	Cnt_Martinique          = "Martinique"
	Cnt_Mauritius           = "Mauritius"
	Cnt_Mexico              = "Mexico"
	Cnt_MoldovaRep          = "Republic of Moldova"
	Cnt_Monaco              = "Monaco"
	Cnt_Mongolia            = "Mongolia"
	Cnt_MontserratIs        = "Montserrat Is"
	Cnt_Morocco             = "Morocco"
	Cnt_Mozambique          = "Mozambique"
	Cnt_Namibia             = "Namibia"
	Cnt_Nauru               = "Nauru"
	Cnt_Nepal               = "Nepal"
	Cnt_NetheriandsAntilles = "Netheriands Antilles"
	Cnt_Netherlands         = "Netherlands"
	Cnt_NewZealand          = "New Zealand"
	Cnt_Nicaragua           = "Nicaragua"
	Cnt_Niger               = "Niger"
	Cnt_Nigeria             = "Nigeria"
	Cnt_NorthKorea          = "North Korea"
	Cnt_Norway              = "Norway"
	Cnt_Oman                = "Oman"
	Cnt_Pakistan            = "Pakistan"
	Cnt_Panama              = "Panama"
	Cnt_PapuaNewCuinea      = "Papua New Cuinea"
	Cnt_Paraguay            = "Paraguay"
	Cnt_Peru                = "Peru"
	Cnt_Philippines         = "Philippines"
	Cnt_Poland              = "Poland"
	Cnt_FrenchPolynesia     = "French Polynesia"
	Cnt_Portugal            = "Portugal"
	Cnt_PuertoRico          = "Puerto Rico"
	Cnt_Qatar               = "Qatar"
	Cnt_Reunion             = "Reunion"
	Cnt_Romania             = "Romania"
	Cnt_Russia              = "Russia"
	Cnt_SaintLueia          = "Saint Lueia"
	Cnt_SaintVincent        = "Saint Vincent"
	Cnt_SamoaEastern        = "Samoa Eastern"
	Cnt_SamoaWestern        = "Samoa Western"
	Cnt_SanMarino           = "San Marino"
	Cnt_SaoTomePrincipe     = "Sao Tome and Principe"
	Cnt_SaudiArabia         = "Saudi Arabia"
	Cnt_Senegal             = "Senegal"
	Cnt_Seychelles          = "Seychelles"
	Cnt_SierraLeone         = "Sierra Leone"
	Cnt_Singapore           = "Singapore"
	Cnt_Slovakia            = "Slovakia"
	Cnt_Slovenia            = "Slovenia"
	Cnt_SolomonIs           = "Solomon Is"
	Cnt_Somali              = "Somali"
	Cnt_SouthAfrica         = "South Africa"
	Cnt_Spain               = "Spain"
	Cnt_SriLanka            = "Sri Lanka"
	Cnt_StLucia             = "St.Lucia"
	Cnt_StVincent           = "St.Vincent"
	Cnt_Sudan               = "Sudan"
	Cnt_Suriname            = "Suriname"
	Cnt_Swaziland           = "Swaziland"
	Cnt_Sweden              = "Sweden"
	Cnt_Switzerland         = "Switzerland"
	Cnt_Syria               = "Syria"
	Cnt_Taiwan              = "Taiwan"
	Cnt_Tajikstan           = "Tajikstan"
	Cnt_Tanzania            = "Tanzania"
	Cnt_Thailand            = "Thailand"
	Cnt_Togo                = "Togo"
	Cnt_Tonga               = "Tonga"
	Cnt_TrinidadTobago      = "Trinidad and Tobago"
	Cnt_Tunisia             = "Tunisia"
	Cnt_Turkey              = "Turkey"
	Cnt_Turkmenistan        = "Turkmenistan"
	Cnt_Uganda              = "Uganda"
	Cnt_Ukraine             = "Ukraine"
	Cnt_UnitedArabEmirates  = "United Arab Emirates"
	Cnt_UnitedKiongdom      = "United Kiongdom "
	Cnt_USA                 = "United States of America"
	Cnt_Uruguay             = "Uruguay"
	Cnt_Uzbekistan          = "Uzbekistan"
	Cnt_Venezuela           = "Venezuela"
	Cnt_Vietnam             = "Vietnam"
	Cnt_Yemen               = "Yemen"
	Cnt_Yugoslavia          = "Yugoslavia"
	Cnt_Zimbabwe            = "Zimbabwe"
	Cnt_Zaire               = "Zaire"
	Cnt_Zambia              = "Zambia"
)

// regionsCache regions information cache
var regionsCache = make(map[string]*Region)

func init() {
	regionsCache[Cnt_Angola] = &Region{"AO", "244", "-7", "安哥拉"}
	regionsCache[Cnt_Afghanistan] = &Region{"AF", "93", "0", "阿富汗"}
	regionsCache[Cnt_Albania] = &Region{"AL", "355", "-7", "阿尔巴尼亚"}
	regionsCache[Cnt_Algeria] = &Region{"DZ", "213", "-8", "阿尔及利亚"}
	regionsCache[Cnt_Andorra] = &Region{"AD", "376", "-8", "安道尔共和国"}
	regionsCache[Cnt_Anguilla] = &Region{"AI", "1264", "-12", "安圭拉岛"}
	regionsCache[Cnt_AntiguaBarbuda] = &Region{"AG", "1268", "-12", "安提瓜和巴布达"}
	regionsCache[Cnt_Argentina] = &Region{"AR", "54", "-11", "阿根廷"}
	regionsCache[Cnt_Armenia] = &Region{"AM", "374", "-6", "亚美尼亚"}
	regionsCache[Cnt_Ascension] = &Region{"", "247", "-8", "阿森松"}
	regionsCache[Cnt_Australia] = &Region{"AU", "61", "2", "澳大利亚"}
	regionsCache[Cnt_Austria] = &Region{"AT", "43", "-7", "奥地利"}
	regionsCache[Cnt_Azerbaijan] = &Region{"AZ", "994", "-5", "阿塞拜疆"}
	regionsCache[Cnt_Bahamas] = &Region{"BS", "1242", "-13", "巴哈马"}
	regionsCache[Cnt_Bahrain] = &Region{"BH", "973", "-5", "巴林"}
	regionsCache[Cnt_Bangladesh] = &Region{"BD", "880", "-2", "孟加拉国"}
	regionsCache[Cnt_Barbados] = &Region{"BB", "1246", "-12", "巴巴多斯"}
	regionsCache[Cnt_Belarus] = &Region{"BY", "375", "-6", "白俄罗斯"}
	regionsCache[Cnt_Belgium] = &Region{"BE", "32", "-7", "比利时"}
	regionsCache[Cnt_Belize] = &Region{"BZ", "501", "-14", "伯利兹"}
	regionsCache[Cnt_Benin] = &Region{"BJ", "229", "-7", "贝宁"}
	regionsCache[Cnt_BermudaIs] = &Region{"BM", "1441", "-12", "百慕大群岛"}
	regionsCache[Cnt_Bolivia] = &Region{"BO", "591", "-12", "玻利维亚"}
	regionsCache[Cnt_Botswana] = &Region{"BW", "267", "-6", "博茨瓦纳"}
	regionsCache[Cnt_Brazil] = &Region{"BR", "55", "-11", "巴西"}
	regionsCache[Cnt_Brunei] = &Region{"BN", "673", "0", "文莱"}
	regionsCache[Cnt_Bulgaria] = &Region{"BG", "359", "-6", "保加利亚"}
	regionsCache[Cnt_BurkinaFaso] = &Region{"BF", "226", "-8", "布基纳法索"}
	regionsCache[Cnt_Burma] = &Region{"MM", "95", "-1.3", "缅甸"}
	regionsCache[Cnt_Burundi] = &Region{"BI", "257", "-6", "布隆迪"}
	regionsCache[Cnt_Cameroon] = &Region{"CM", "237", "-7", "喀麦隆"}
	regionsCache[Cnt_Canada] = &Region{"CA", "1", "-13", "加拿大"}
	regionsCache[Cnt_CaymanIs] = &Region{"", "1345", "-13", "开曼群岛"}
	regionsCache[Cnt_CentralAfricanRep] = &Region{"CF", "236", "-7", "中非共和国"}
	regionsCache[Cnt_Chad] = &Region{"TD", "235", "-7", "乍得"}
	regionsCache[Cnt_Chile] = &Region{"CL", "56", "-13", "智利"}
	regionsCache[Cnt_China] = &Region{"CN", "86", "0", "中国"}
	regionsCache[Cnt_Colombia] = &Region{"CO", "57", "0", "哥伦比亚"}
	regionsCache[Cnt_Congo] = &Region{"CG", "242", "-7", "刚果"}
	regionsCache[Cnt_CookIs] = &Region{"CK", "682", "-18.3", "库克群岛"}
	regionsCache[Cnt_CostaRica] = &Region{"CR", "506", "-14", "哥斯达黎加"}
	regionsCache[Cnt_Cuba] = &Region{"CU", "53", "-13", "古巴"}
	regionsCache[Cnt_Cyprus] = &Region{"CY", "357", "-6", "塞浦路斯"}
	regionsCache[Cnt_CzechRep] = &Region{"CZ", "420", "-7", "捷克"}
	regionsCache[Cnt_Denmark] = &Region{"DK", "45", "-7", "丹麦"}
	regionsCache[Cnt_Djibouti] = &Region{"DJ", "253", "-5", "吉布提"}
	regionsCache[Cnt_DominicaRep] = &Region{"DO", "1890", "-13", "多米尼加共和国"}
	regionsCache[Cnt_Ecuador] = &Region{"EC", "593", "-13", "厄瓜多尔"}
	regionsCache[Cnt_Egypt] = &Region{"EG", "20", "-6", "埃及"}
	regionsCache[Cnt_EISalvador] = &Region{"SV", "503", "-14", "萨尔瓦多"}
	regionsCache[Cnt_Estonia] = &Region{"EE", "372", "-5", "爱沙尼亚"}
	regionsCache[Cnt_Ethiopia] = &Region{"ET", "251", "-5", "埃塞俄比亚"}
	regionsCache[Cnt_Fiji] = &Region{"FJ", "679", "4", "斐济"}
	regionsCache[Cnt_Finland] = &Region{"FI", "358", "-6", "芬兰"}
	regionsCache[Cnt_France] = &Region{"FR", "33", "-8", "法国"}
	regionsCache[Cnt_FrenchGuiana] = &Region{"GF", "594", "-12", "法属圭亚那"}
	regionsCache[Cnt_Gabon] = &Region{"GA", "241", "-7", "加蓬"}
	regionsCache[Cnt_Gambia] = &Region{"GM", "220", "-8", "冈比亚"}
	regionsCache[Cnt_Georgia] = &Region{"GE", "995", "0", "格鲁吉亚"}
	regionsCache[Cnt_Germany] = &Region{"DE", "49", "-7", "德国"}
	regionsCache[Cnt_Ghana] = &Region{"GH", "233", "-8", "加纳"}
	regionsCache[Cnt_Gibraltar] = &Region{"GI", "350", "-8", "直布罗陀"}
	regionsCache[Cnt_Greece] = &Region{"GR", "30", "-6", "希腊"}
	regionsCache[Cnt_Grenada] = &Region{"GD", "1809", "-14", "格林纳达"}
	regionsCache[Cnt_Guam] = &Region{"GU", "1671", "2", "关岛"}
	regionsCache[Cnt_Guatemala] = &Region{"GT", "502", "-14", "危地马拉"}
	regionsCache[Cnt_Guinea] = &Region{"GN", "224", "-8", "几内亚"}
	regionsCache[Cnt_Guyana] = &Region{"GY", "592", "-11", "圭亚那"}
	regionsCache[Cnt_Haiti] = &Region{"HT", "509", "-13", "海地"}
	regionsCache[Cnt_Honduras] = &Region{"HN", "504", "-14", "洪都拉斯"}
	regionsCache[Cnt_Hongkong] = &Region{"HK", "852", "0", "香港"}
	regionsCache[Cnt_Hungary] = &Region{"HU", "36", "-7", "匈牙利"}
	regionsCache[Cnt_Iceland] = &Region{"IS", "354", "-9", "冰岛"}
	regionsCache[Cnt_India] = &Region{"IN", "91", "-2.3", "印度"}
	regionsCache[Cnt_Indonesia] = &Region{"ID", "62", "-0.3", "印度尼西亚"}
	regionsCache[Cnt_Iran] = &Region{"IR", "98", "-4.3", "伊朗"}
	regionsCache[Cnt_Iraq] = &Region{"IQ", "964", "-5", "伊拉克"}
	regionsCache[Cnt_Ireland] = &Region{"IE", "353", "-4.3", "爱尔兰"}
	regionsCache[Cnt_Israel] = &Region{"IL", "972", "-6", "以色列"}
	regionsCache[Cnt_Italy] = &Region{"IT", "39", "-7", "意大利"}
	regionsCache[Cnt_IvoryCoast] = &Region{"", "225", "-6", "科特迪瓦"}
	regionsCache[Cnt_Jamaica] = &Region{"JM", "1876", "-12", "牙买加"}
	regionsCache[Cnt_Japan] = &Region{"JP", "81", "1", "日本"}
	regionsCache[Cnt_Jordan] = &Region{"JO", "962", "-6", "约旦"}
	regionsCache[Cnt_Kampuchea] = &Region{"KH", "855", "-1", "柬埔寨"}
	regionsCache[Cnt_Kazakstan] = &Region{"KZ", "327", "-5", "哈萨克斯坦"}
	regionsCache[Cnt_Kenya] = &Region{"KE", "254", "-5", "肯尼亚"}
	regionsCache[Cnt_Korea] = &Region{"KR", "82", "1", "韩国"}
	regionsCache[Cnt_Kuwait] = &Region{"KW", "965", "-5", "科威特"}
	regionsCache[Cnt_Kyrgyzstan] = &Region{"KG", "331", "-5", "吉尔吉斯坦"}
	regionsCache[Cnt_Laos] = &Region{"LA", "856", "-1", "老挝"}
	regionsCache[Cnt_Latvia] = &Region{"LV", "371", "-5", "拉脱维亚"}
	regionsCache[Cnt_Lebanon] = &Region{"LB", "961", "-6", "黎巴嫩"}
	regionsCache[Cnt_Lesotho] = &Region{"LS", "266", "-6", "莱索托"}
	regionsCache[Cnt_Liberia] = &Region{"LR", "231", "-8", "利比里亚"}
	regionsCache[Cnt_Libya] = &Region{"LY", "218", "-6", "利比亚"}
	regionsCache[Cnt_Liechtenstein] = &Region{"LI", "423", "-7", "列支敦士登"}
	regionsCache[Cnt_Lithuania] = &Region{"LT", "370", "-5", "立陶宛"}
	regionsCache[Cnt_Luxembourg] = &Region{"LU", "352", "-7", "卢森堡"}
	regionsCache[Cnt_Macao] = &Region{"MO", "853", "0", "澳门"}
	regionsCache[Cnt_Madagascar] = &Region{"MG", "261", "-5", "马达加斯加"}
	regionsCache[Cnt_Malawi] = &Region{"MW", "265", "-6", "马拉维"}
	regionsCache[Cnt_Malaysia] = &Region{"MY", "60", "-0.5", "马来西亚"}
	regionsCache[Cnt_Maldives] = &Region{"MV", "960", "-7", "马尔代夫"}
	regionsCache[Cnt_Mali] = &Region{"ML", "223", "-8", "马里"}
	regionsCache[Cnt_Malta] = &Region{"MT", "356", "-7", "马耳他"}
	regionsCache[Cnt_MarianaIs] = &Region{"", "1670", "1", "马里亚那群岛"}
	regionsCache[Cnt_Martinique] = &Region{"", "596", "-12", "马提尼克"}
	regionsCache[Cnt_Mauritius] = &Region{"MU", "230", "-4", "毛里求斯"}
	regionsCache[Cnt_Mexico] = &Region{"MX", "52", "-15", "墨西哥"}
	regionsCache[Cnt_MoldovaRep] = &Region{"MD", "373", "-5", "摩尔多瓦"}
	regionsCache[Cnt_Monaco] = &Region{"MC", "377", "-7", "摩纳哥"}
	regionsCache[Cnt_Mongolia] = &Region{"MN", "976", "0", "蒙古"}
	regionsCache[Cnt_MontserratIs] = &Region{"MS", "1664", "-12", "蒙特塞拉特岛"}
	regionsCache[Cnt_Morocco] = &Region{"MA", "212", "-6", "摩洛哥"}
	regionsCache[Cnt_Mozambique] = &Region{"MZ", "258", "-6", "莫桑比克"}
	regionsCache[Cnt_Namibia] = &Region{"NA", "264", "-7", "纳米比亚"}
	regionsCache[Cnt_Nauru] = &Region{"NR", "674", "4", "瑙鲁"}
	regionsCache[Cnt_Nepal] = &Region{"NP", "977", "-2.3", "尼泊尔"}
	regionsCache[Cnt_NetheriandsAntilles] = &Region{"", "599", "-12", "荷属安的列斯"}
	regionsCache[Cnt_Netherlands] = &Region{"NL", "31", "-7", "荷兰"}
	regionsCache[Cnt_NewZealand] = &Region{"NZ", "64", "4", "新西兰"}
	regionsCache[Cnt_Nicaragua] = &Region{"NI", "505", "-14", "尼加拉瓜"}
	regionsCache[Cnt_Niger] = &Region{"NE", "227", "-8", "尼日尔"}
	regionsCache[Cnt_Nigeria] = &Region{"NG", "234", "-7", "尼日利亚"}
	regionsCache[Cnt_NorthKorea] = &Region{"KP", "850", "1", "朝鲜"}
	regionsCache[Cnt_Norway] = &Region{"NO", "47", "-7", "挪威"}
	regionsCache[Cnt_Oman] = &Region{"OM", "968", "-4", "阿曼"}
	regionsCache[Cnt_Pakistan] = &Region{"PK", "92", "-2.3", "巴基斯坦"}
	regionsCache[Cnt_Panama] = &Region{"PA", "507", "-13", "巴拿马"}
	regionsCache[Cnt_PapuaNewCuinea] = &Region{"PG", "675", "2", "巴布亚新几内亚"}
	regionsCache[Cnt_Paraguay] = &Region{"PY", "595", "-12", "巴拉圭"}
	regionsCache[Cnt_Peru] = &Region{"PE", "51", "-13", "秘鲁"}
	regionsCache[Cnt_Philippines] = &Region{"PH", "63", "0", "菲律宾"}
	regionsCache[Cnt_Poland] = &Region{"PL", "48", "-7", "波兰"}
	regionsCache[Cnt_FrenchPolynesia] = &Region{"PF", "689", "3", "法属玻利尼西亚"}
	regionsCache[Cnt_Portugal] = &Region{"PT", "351", "-8", "葡萄牙"}
	regionsCache[Cnt_PuertoRico] = &Region{"PR", "1787", "-12", "波多黎各"}
	regionsCache[Cnt_Qatar] = &Region{"QA", "974", "-5", "卡塔尔"}
	regionsCache[Cnt_Reunion] = &Region{"", "262", "-4", "留尼旺"}
	regionsCache[Cnt_Romania] = &Region{"RO", "40", "-6", "罗马尼亚"}
	regionsCache[Cnt_Russia] = &Region{"RU", "7", "-5", "俄罗斯"}
	regionsCache[Cnt_SaintLueia] = &Region{"LC", "1758", "-12", "圣卢西亚"}
	regionsCache[Cnt_SaintVincent] = &Region{"VC", "1784", "-12", "圣文森特岛"}
	regionsCache[Cnt_SamoaEastern] = &Region{"", "684", "-19", "东萨摩亚(美)"}
	regionsCache[Cnt_SamoaWestern] = &Region{"", "685", "-19", "西萨摩亚"}
	regionsCache[Cnt_SanMarino] = &Region{"SM", "378", "-7", "圣马力诺"}
	regionsCache[Cnt_SaoTomePrincipe] = &Region{"ST", "239", "-8", "圣多美和普林西比"}
	regionsCache[Cnt_SaudiArabia] = &Region{"SA", "966", "-5", "沙特阿拉伯"}
	regionsCache[Cnt_Senegal] = &Region{"SN", "221", "-8", "塞内加尔"}
	regionsCache[Cnt_Seychelles] = &Region{"SC", "248", "-4", "塞舌尔"}
	regionsCache[Cnt_SierraLeone] = &Region{"SL", "232", "-8", "塞拉利昂"}
	regionsCache[Cnt_Singapore] = &Region{"SG", "65", "0.3", "新加坡"}
	regionsCache[Cnt_Slovakia] = &Region{"SK", "421", "-7", "斯洛伐克"}
	regionsCache[Cnt_Slovenia] = &Region{"SI", "386", "-7", "斯洛文尼亚"}
	regionsCache[Cnt_SolomonIs] = &Region{"SB", "677", "3", "所罗门群岛"}
	regionsCache[Cnt_Somali] = &Region{"SO", "252", "-5", "索马里"}
	regionsCache[Cnt_SouthAfrica] = &Region{"ZA", "27", "-6", "南非"}
	regionsCache[Cnt_Spain] = &Region{"ES", "34", "-8", "西班牙"}
	regionsCache[Cnt_SriLanka] = &Region{"LK", "94", "0", "斯里兰卡"}
	regionsCache[Cnt_StLucia] = &Region{"LC", "1758", "-12", "圣卢西亚"}
	regionsCache[Cnt_StVincent] = &Region{"VC", "1784", "-12", "圣文森特"}
	regionsCache[Cnt_Sudan] = &Region{"SD", "249", "-6", "苏丹"}
	regionsCache[Cnt_Suriname] = &Region{"SR", "597", "-11.3", "苏里南"}
	regionsCache[Cnt_Swaziland] = &Region{"SZ", "268", "-6", "斯威士兰"}
	regionsCache[Cnt_Sweden] = &Region{"SE", "46", "-7", "瑞典"}
	regionsCache[Cnt_Switzerland] = &Region{"CH", "41", "-7", "瑞士"}
	regionsCache[Cnt_Syria] = &Region{"SY", "963", "-6", "叙利亚"}
	regionsCache[Cnt_Taiwan] = &Region{"TW", "886", "0", "台湾省"}
	regionsCache[Cnt_Tajikstan] = &Region{"TJ", "992", "-5", "塔吉克斯坦"}
	regionsCache[Cnt_Tanzania] = &Region{"TZ", "255", "-5", "坦桑尼亚"}
	regionsCache[Cnt_Thailand] = &Region{"TH", "66", "-1", "泰国"}
	regionsCache[Cnt_Togo] = &Region{"TG", "228", "-8", "多哥"}
	regionsCache[Cnt_Tonga] = &Region{"TO", "676", "4", "汤加"}
	regionsCache[Cnt_TrinidadTobago] = &Region{"TT", "1809", "-12", "特立尼达和多巴哥"}
	regionsCache[Cnt_Tunisia] = &Region{"TN", "216", "-7", "突尼斯"}
	regionsCache[Cnt_Turkey] = &Region{"TR", "90", "-6", "土耳其"}
	regionsCache[Cnt_Turkmenistan] = &Region{"TM", "993", "-5", "土库曼斯坦"}
	regionsCache[Cnt_Uganda] = &Region{"UG", "256", "-5", "乌干达"}
	regionsCache[Cnt_Ukraine] = &Region{"UA", "380", "-5", "乌克兰"}
	regionsCache[Cnt_UnitedArabEmirates] = &Region{"AE", "971", "-4", "阿拉伯联合酋长国"}
	regionsCache[Cnt_UnitedKiongdom] = &Region{"GB", "44", "-8", "英国"}
	regionsCache[Cnt_USA] = &Region{"US", "1", "-13", "美国"}
	regionsCache[Cnt_Uruguay] = &Region{"UY", "598", "-10.3", "乌拉圭"}
	regionsCache[Cnt_Uzbekistan] = &Region{"UZ", "233", "-5", "乌兹别克斯坦"}
	regionsCache[Cnt_Venezuela] = &Region{"VE", "58", "-12.3", "委内瑞拉"}
	regionsCache[Cnt_Vietnam] = &Region{"VN", "84", "-1", "越南"}
	regionsCache[Cnt_Yemen] = &Region{"YE", "967", "-5", "也门"}
	regionsCache[Cnt_Yugoslavia] = &Region{"YU", "381", "-7", "南斯拉夫"}
	regionsCache[Cnt_Zimbabwe] = &Region{"ZW", "263", "-6", "津巴布韦"}
	regionsCache[Cnt_Zaire] = &Region{"ZR", "243", "-7", "扎伊尔"}
	regionsCache[Cnt_Zambia] = &Region{"ZM", "260", "-6", "赞比亚"}
}

// GetRegion get region information by country
func GetRegion(country string) *Region {
	region := regionsCache[country]
	if region != nil {
		return &Region{
			Code:     region.Code,
			Phone:    region.Phone,
			TimeDiff: region.TimeDiff,
			CnName:   region.CnName,
		}
	}
	return nil
}

// GetRegionByCode get country and region information by code and phone
func GetRegionByCode(code string, phone ...string) (string, *Region) {
	regions, lastcountry := make(map[string]*Region), ""
	for country, region := range regionsCache {
		if region.Code == code {
			regions[country] = region
			lastcountry = country
		}
	}

	if len(phone) > 0 && phone[0] != "" {
		for country, region := range regions {
			if region.Phone == phone[0] {
				return country, &Region{
					Code:     region.Code,
					Phone:    region.Phone,
					TimeDiff: region.TimeDiff,
					CnName:   region.CnName,
				}
			}
		}
	} else if lastcountry != "" {
		region := regions[lastcountry]
		return lastcountry, &Region{
			Code:     region.Code,
			Phone:    region.Phone,
			TimeDiff: region.TimeDiff,
			CnName:   region.CnName,
		}
	}
	return "", nil
}
