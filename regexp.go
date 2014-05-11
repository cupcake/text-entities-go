package text

import "regexp"

// Do not modify this file, generate it with 'go run gen_regexp.go'
var (
	asciiDomain         = regexp.MustCompile("(?:[\\-0-9A-Za-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]+(?i:\\.))+(?:(?i:A)(?:(?i:C)(?:(?i:ADEMY)|(?i:COUNTANTS)|(?i:TOR))|(?i:ERO)|(?i:GENCY)|(?i:IRFORCE)|(?i:R)(?:(?i:CHI)|(?i:PA))|(?i:S)(?:(?i:IA)|(?i:SOCIATES))|(?i:XA))|(?i:B)(?:(?i:A)(?:(?i:R)(?:(?:)|(?i:GAINS))|(?i:YERN))|(?i:E)(?:(?i:RLIN)|(?i:ST))|(?i:I)(?:(?i:D)|(?i:KE)|(?i:Z))|(?i:L)(?:(?i:ACK)(?:(?:)|(?i:FRIDAY))|(?i:UE))|(?i:OUTIQUE)|(?i:U)(?:(?i:ILD)(?:(?:)|(?i:ERS))|(?i:ZZ)))|(?i:C)(?:(?i:A)(?:(?i:B)|(?i:M)(?:(?i:ERA)|(?i:P))|(?i:PITAL)|(?i:R)(?:(?i:DS)|(?i:E)(?:(?:)|(?i:ER)(?:(?:)|(?i:S))))|(?i:SH)|(?i:T)(?:(?:)|(?i:ERING)))|(?i:E)(?:(?i:NTER)|(?i:O))|(?i:H)(?:(?i:EAP)|(?i:RISTMAS))|(?i:ITIC)|(?i:L)(?:(?i:AIMS)|(?i:EANING)|(?i:INIC)|(?i:OTHING)|(?i:UB))|(?i:O)(?:(?i:DES)|(?i:FFEE)|(?i:L)(?:(?i:LEGE)|(?i:OGNE))|(?i:M)(?:(?:)|(?i:MUNITY)|(?i:P)(?:(?i:ANY)|(?i:UTER)))|(?i:N)(?:(?i:DOS)|(?i:S)(?:(?i:TRUCTION)|(?i:ULTING))|(?i:TRACTORS))|(?i:O)(?:(?i:KING)|[LPlp])|(?i:UNTRY))|(?i:R)(?:(?i:EDIT)(?:(?:)|(?i:CARD))|(?i:UISES)))|(?i:D)(?:(?i:A)(?:(?i:NCE)|(?i:TING))|(?i:E)(?:(?i:MOCRAT)|(?i:NTAL)|(?i:SI))|(?i:I)(?:(?i:AMONDS)|(?i:GITAL)|(?i:RECTORY)|(?i:SCOUNT))|(?i:NP)|(?i:OMAINS))|(?i:E)(?:(?i:DU)(?:(?:)|(?i:CATION))|(?i:MAIL)|(?i:N)(?:(?i:GINEERING)|(?i:TERPRISES))|(?i:QUIPMENT)|(?i:STATE)|(?i:US)|(?i:VENTS)|(?i:X)(?:(?i:CHANGE)|(?i:P)(?:(?i:ERT)|(?i:OSED))))|(?i:F)(?:(?i:A)(?:(?i:IL)|(?i:RM))|(?i:EEDBACK)|(?i:I)(?:(?i:NANC)(?:(?i:E)|(?i:IAL))|(?i:SH)(?:(?:)|(?i:ING))|(?i:TNESS))|(?i:L)(?:(?i:IGHTS)|(?i:ORIST))|(?i:O)(?:(?i:O)|(?i:UNDATION))|(?i:ROGANS)|(?i:U)(?:(?i:ND)|(?i:RNITURE)|(?i:TBOL)))|(?i:G)(?:(?i:AL)(?:(?:)|(?i:LERY))|(?i:IFT)|(?i:L)(?:(?i:ASS)|(?i:OBO))|(?i:MO)|(?i:O)[PVpv]|(?i:R)(?:(?i:A)(?:(?i:PHICS)|(?i:TIS))|(?i:IPE))|(?i:U)(?:(?i:ITARS)|(?i:RU)))|(?i:H)(?:(?i:AUS)|(?i:O)(?:(?i:L)(?:(?i:DINGS)|(?i:IDAY))|(?i:RSE)|(?i:USE)))|(?i:I)(?:(?i:MMOBILIEN)|(?i:N)(?:(?i:DUSTRIES)|(?i:FO)|(?i:K)|(?i:S)(?:(?i:TITUTE)|(?i:URE))|(?i:T)(?:(?:)|(?i:ERNATIONAL))|(?i:VESTMENTS)))|(?i:J)(?:(?i:ETZT)|(?i:OBS))|(?i:K)(?:(?i:AUFEN)|(?i:I)(?:(?i:M)|(?i:TCHEN)|(?i:WI))|(?i:OELN)|(?i:RED))|(?i:L)(?:(?i:AND)|(?i:EASE)|(?i:I)(?:(?i:GHTING)|(?i:M)(?:(?i:ITED)|(?i:O))|(?i:NK))|(?i:ONDON)|(?i:UXURY))|(?i:M)(?:(?i:A)(?:(?i:ISON)|(?i:N)(?:(?i:AGEMENT)|(?i:GO))|(?i:RKETING))|(?i:E)(?:(?i:DIA)|(?i:ET)|(?i:NU))|(?i:I)(?:(?i:AMI)|(?i:L))|(?i:O)(?:(?i:BI)|(?i:DA)|(?i:E)|(?i:NASH)|(?i:SCOW))|(?i:USEUM))|(?i:N)(?:(?i:A)(?:(?i:GOYA)|(?i:ME))|(?i:E)(?:(?i:T)|(?i:USTAR))|(?i:INJA)|(?i:YC))|(?i:O)(?:(?i:KINAWA)|(?i:NL)|(?i:RG))|(?i:P)(?:(?i:AR)(?:(?i:IS)|(?i:T)(?:(?i:NERS)|(?i:S)))|(?i:HOTO)(?:(?:)|(?i:GRAPHY)|(?i:S))|(?i:I)(?:(?i:C)(?:(?i:S)|(?i:TURES))|(?i:NK))|(?i:LUMBING)|(?i:OST)|(?i:RO)(?:(?:)|(?i:DUCTIONS)|(?i:PERTIES))|(?i:UB))|(?i:Q)(?:(?i:PON)|(?i:UEBEC))|(?i:R)(?:(?i:E)(?:(?i:CIPES)|(?i:D)|(?i:ISEN)|(?i:N)(?:(?:)|(?i:TALS))|(?i:P)(?:(?i:AIR)|(?i:ORT))|(?i:ST)|(?i:VIEWS))|(?i:ICH)|(?i:O)(?:(?i:CKS)|(?i:DEO))|(?i:UHR)|(?i:YUKYU))|(?i:S)(?:(?i:AARLAND)|(?i:CHULE)|(?i:E)(?:(?i:RVICES)|(?i:XY))|(?i:H)(?:(?i:IKSHA)|(?i:OES))|(?i:INGLES)|(?i:O)(?:(?i:CIAL)|(?i:HU)|(?i:L)(?:(?i:AR)|(?i:UTIONS))|(?i:Y))|(?i:U)(?:(?i:PP)(?:(?i:L)(?:(?i:IES)|(?i:Y))|(?i:ORT))|(?i:RGERY))|(?i:YSTEMS))|(?i:T)(?:(?i:A)(?:(?i:TTOO)|(?i:X))|(?i:E)(?:(?i:CHNOLOGY)|(?i:L))|(?i:I)(?:(?i:ENDA)|(?i:PS))|(?i:O)(?:(?i:DAY)|(?i:KYO)|(?i:OLS)|(?i:WN)|(?i:YS))|(?i:RA)(?:(?i:DE)|(?i:INING)|(?i:VEL)))|(?i:UN)(?:(?i:IVERSITY)|(?i:O))|(?i:V)(?:(?i:ACATIONS)|(?i:E)(?:(?i:GAS)|(?i:NTURES))|(?i:I)(?:(?i:AJES)|(?i:LLAS)|(?i:SION))|(?i:O)(?:(?i:DKA)|(?i:T)(?:(?i:E)|(?i:ING)|(?i:O))|(?i:YAGE)))|(?i:W)(?:(?i:A)(?:(?i:NG)|(?i:TCH))|(?i:E)(?:(?i:BCAM)|(?i:D))|(?i:I)(?:(?i:EN)|(?i:KI))|(?i:ORKS)|(?i:T)[CFcf])|(?i:X)(?:(?i:XX)|(?i:YZ))|(?i:YOKOHAMA)|(?i:ZONE)|(?i:A)[C-GIL-OQ-UW-XZc-gil-oq-uw-xz\u017f]|(?i:B)[A-BD-JM-OR-TV-WY-Za-bd-jm-or-tv-wy-z\u017f]|(?i:C)[AC-DF-IK-ORU-Zac-df-ik-oru-z\u212a]|(?i:D)[EJ-KMOZej-kmoz\u212a]|(?i:E)[CEGR-Ucegr-u\u017f]|(?i:F)[I-KMORi-kmor\u212a]|(?i:G)[A-BD-IL-NP-UWYa-bd-il-np-uwy\u017f]|(?i:H)[KM-NRT-Ukm-nrt-u\u212a]|(?i:I)[D-EL-OQ-Td-el-oq-t\u017f]|(?i:J)[EMO-Pemo-p]|(?i:K)[EG-IM-NPRWY-Zeg-im-nprwy-z]|(?i:L)[A-CIKR-VYa-cikr-vy\u017f\u212a]|(?i:M)[AC-EG-HK-Zac-eg-hk-z\u017f\u212a]|(?i:N)[ACE-GILO-PRUZace-gilo-pruz]|(?i:OM)|(?i:P)[AE-HK-NR-TWYae-hk-nr-twy\u017f\u212a]|(?i:QA)|(?i:R)[EOSUWeosuw\u017f]|(?i:S)[A-EG-ORT-VX-Za-eg-ort-vx-z\u212a]|(?i:T)[C-DF-HJ-PRTV-WZc-df-hj-prtv-wz\u212a]|(?i:U)[AGKSY-Zagksy-z\u017f\u212a]|(?i:V)[ACEGINUaceginu]|(?i:W)[FSfs\u017f]|(?i:Y)[ETet]|(?i:Z)[AMWamw]|(?i:\u96c6\u56e2)|(?i:\u5728\u7ebf)|(?i:\ud55c\uad6d)|(?i:\u09ad\u09be\u09b0\u09a4)|(?i:\u516c)[\u53f8\u76ca]|(?i:\u79fb\u52a8)|(?i:\u6211\u7231\u4f60)|(?i:\u041c\u041e\u0421\u041a\u0412\u0410)|(?i:\u049a\u0410\u0417)|(?i:\u041e\u041d\u041b\u0410\u0419\u041d)|(?i:\u0421)(?:(?i:\u0410\u0419\u0422)|(?i:\u0420\u0411))|(?i:\u041e\u0420\u0413)|(?i:\uc0bc\uc131)|(?i:\u0b9a\u0bbf\u0b99\u0bcd\u0b95\u0baa\u0bcd\u0baa\u0bc2\u0bb0\u0bcd)|(?i:\u5546\u57ce)|(?i:\u0414\u0415\u0422\u0418)|(?i:\u4e2d)(?:(?i:\u6587\u7f51)|[\u4fe1\u56fd\u570b])|(?i:\u0c2d\u0c3e\u0c30\u0c24\u0c4d)|(?i:\u0dbd\u0d82\u0d9a\u0dcf)|(?i:\u0aad\u0abe\u0ab0\u0aa4)|(?i:\u092d\u093e\u0930\u0924)|(?i:\u0938\u0902\u0917\u0920\u0928)|(?i:\u7f51\u7edc)|(?i:\u0423\u041a\u0420)|(?i:\u9999\u6e2f)|(?i:\u53f0)[\u6e7e\u7063]|(?i:\u041c\u041e\u041d)|(?i:\u0627\u0644\u062c\u0632\u0627\u0626\u0631)|(?i:\u0639\u0645\u0627\u0646)|(?i:\u0627)(?:(?i:\u06cc\u0631\u0627\u0646)|(?i:\u0645\u0627\u0631\u0627\u062a))|(?i:\u0628\u0627\u0632\u0627\u0631)|(?i:\u0627\u0644\u0627\u0631\u062f\u0646)|(?i:\u0628\u06be\u0627\u0631\u062a)|(?i:\u0627\u0644)(?:(?i:\u0645\u063a\u0631\u0628)|(?i:\u0633\u0639\u0648\u062f\u064a\u0629))|(?i:\u0645\u0644\u064a\u0633\u064a\u0627)|(?i:\u0634\u0628\u0643\u0629)|(?i:\u673a\u6784)|(?i:\u7ec4\u7ec7\u673a\u6784)|(?i:\u0e44\u0e17\u0e22)|(?i:\u0633\u0648\u0631\u064a\u0629)|(?i:\u0420\u0424)|(?i:\u062a\u0648\u0646\u0633)|(?i:\u307f\u3093\u306a)|(?i:\u4e16\u754c)|(?i:\u0a2d\u0a3e\u0a30\u0a24)|(?i:\u7f51\u5740)|(?i:\u6e38\u620f)|(?i:\u0645\u0635\u0631)|(?i:\u0642\u0637\u0631)|(?i:\u0b87)(?:(?i:\u0bb2\u0b99\u0bcd\u0b95\u0bc8)|(?i:\u0ba8\u0bcd\u0ba4\u0bbf\u0baf\u0bbe))|(?i:\u65b0\u52a0\u5761)|(?i:\u0641\u0644\u0633\u0637\u064a\u0646)|(?i:\u653f\u52a1)|(?i:XN--)[0-9A-Za-z\u017f\u212a]+)")
	hashtagPattern      = regexp.MustCompile("(?:^|$|[^&0-9A-Z_a-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u0400-\u0527\u0591-\u05bf\u05c1-\u05c2\u05c4-\u05c5\u05c7\u05d0-\u05ea\u05f0-\u05f4\u0610-\u061a\u0620-\u065f\u066e-\u06d3\u06d5-\u06dc\u06de-\u06e8\u06ea-\u06ef\u06fa-\u06fc\u06ff\u0750-\u077f\u08a0\u08a2-\u08ac\u08e4-\u08fe\u0e01-\u0e3a\u0e40-\u0e4e\u1100-\u11ff\u1e00-\u1eff\u1fbe\\x{200c}\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f\u2de0-\u2dff\u3003\u3005\u303b\u3041-\u3096\u3099-\u309e\u30a1-\u30fa\u30fc-\u30fe\\x{3130}-\u3185\u3400-\\x{4dbf}\u4e00-\\x{9fff}\ua640-\ua69f\ua960-\\x{a97f}\uac00-\\x{d7ff}\\x{fb12}-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe70-\ufe74\ufe76-\ufefc\uff10-\uff19\uff21-\uff3a\uff41-\uff5a\uff66-\uff9f\uffa1-\uffdc\U0002a700-\\x{2b81f}\U0002f800-\\x{2fa1f}])([#\uff03]([0-9A-Z_a-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u0400-\u0527\u0591-\u05bf\u05c1-\u05c2\u05c4-\u05c5\u05c7\u05d0-\u05ea\u05f0-\u05f4\u0610-\u061a\u0620-\u065f\u066e-\u06d3\u06d5-\u06dc\u06de-\u06e8\u06ea-\u06ef\u06fa-\u06fc\u06ff\u0750-\u077f\u08a0\u08a2-\u08ac\u08e4-\u08fe\u0e01-\u0e3a\u0e40-\u0e4e\u1100-\u11ff\u1e00-\u1eff\u1fbe\\x{200c}\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f\u2de0-\u2dff\u3003\u3005\u303b\u3041-\u3096\u3099-\u309e\u30a1-\u30fa\u30fc-\u30fe\\x{3130}-\u3185\u3400-\\x{4dbf}\u4e00-\\x{9fff}\ua640-\ua69f\ua960-\\x{a97f}\uac00-\\x{d7ff}\\x{fb12}-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe70-\ufe74\ufe76-\ufefc\uff10-\uff19\uff21-\uff3a\uff41-\uff5a\uff66-\uff9f\uffa1-\uffdc\U0002a700-\\x{2b81f}\U0002f800-\\x{2fa1f}]*[A-Z_a-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u0400-\u0527\u0591-\u05bf\u05c1-\u05c2\u05c4-\u05c5\u05c7\u05d0-\u05ea\u05f0-\u05f4\u0610-\u061a\u0620-\u065f\u066e-\u06d3\u06d5-\u06dc\u06de-\u06e8\u06ea-\u06ef\u06fa-\u06fc\u06ff\u0750-\u077f\u08a0\u08a2-\u08ac\u08e4-\u08fe\u0e01-\u0e3a\u0e40-\u0e4e\u1100-\u11ff\u1e00-\u1eff\u1fbe\\x{200c}\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f\u2de0-\u2dff\u3003\u3005\u303b\u3041-\u3096\u3099-\u309e\u30a1-\u30fa\u30fc-\u30fe\\x{3130}-\u3185\u3400-\\x{4dbf}\u4e00-\\x{9fff}\ua640-\ua69f\ua960-\\x{a97f}\uac00-\\x{d7ff}\\x{fb12}-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe70-\ufe74\ufe76-\ufefc\uff10-\uff19\uff21-\uff3a\uff41-\uff5a\uff66-\uff9f\uffa1-\uffdc\U0002a700-\\x{2b81f}\U0002f800-\\x{2fa1f}][0-9A-Z_a-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u0400-\u0527\u0591-\u05bf\u05c1-\u05c2\u05c4-\u05c5\u05c7\u05d0-\u05ea\u05f0-\u05f4\u0610-\u061a\u0620-\u065f\u066e-\u06d3\u06d5-\u06dc\u06de-\u06e8\u06ea-\u06ef\u06fa-\u06fc\u06ff\u0750-\u077f\u08a0\u08a2-\u08ac\u08e4-\u08fe\u0e01-\u0e3a\u0e40-\u0e4e\u1100-\u11ff\u1e00-\u1eff\u1fbe\\x{200c}\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f\u2de0-\u2dff\u3003\u3005\u303b\u3041-\u3096\u3099-\u309e\u30a1-\u30fa\u30fc-\u30fe\\x{3130}-\u3185\u3400-\\x{4dbf}\u4e00-\\x{9fff}\ua640-\ua69f\ua960-\\x{a97f}\uac00-\\x{d7ff}\\x{fb12}-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe70-\ufe74\ufe76-\ufefc\uff10-\uff19\uff21-\uff3a\uff41-\uff5a\uff66-\uff9f\uffa1-\uffdc\U0002a700-\\x{2b81f}\U0002f800-\\x{2fa1f}]*))")
	invalidBeforeDomain = regexp.MustCompile("[\\--/_](?-m:$)")
	invalidHashtagEnd   = regexp.MustCompile("\\A(?:[#\uff03]|://)")
	invalidWithoutPath  = regexp.MustCompile("\\A(?:[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}][^\\t-\\r -!#-,\\.-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}]*)?[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}](?i:\\.)(?:(?i:A)[C-GIL-OQ-UW-XZc-gil-oq-uw-xz\u017f]|(?i:B)[A-BD-JM-OR-TV-WY-Za-bd-jm-or-tv-wy-z\u017f]|(?i:C)[AC-DF-IK-ORU-Zac-df-ik-oru-z\u212a]|(?i:D)[EJ-KMOZej-kmoz\u212a]|(?i:E)[CEGR-Ucegr-u\u017f]|(?i:F)[I-KMORi-kmor\u212a]|(?i:G)[A-BD-IL-NP-UWYa-bd-il-np-uwy\u017f]|(?i:H)[KM-NRT-Ukm-nrt-u\u212a]|(?i:I)[D-EL-OQ-Td-el-oq-t\u017f]|(?i:J)[EMO-Pemo-p]|(?i:K)[EG-IM-NPRWY-Zeg-im-nprwy-z]|(?i:L)[A-CIKR-VYa-cikr-vy\u017f\u212a]|(?i:M)[AC-EG-HK-Zac-eg-hk-z\u017f\u212a]|(?i:N)[ACE-GILO-PRUZace-gilo-pruz]|(?i:OM)|(?i:P)[AE-HK-NR-TWYae-hk-nr-twy\u017f\u212a]|(?i:QA)|(?i:R)[EOSUWeosuw\u017f]|(?i:S)[A-EG-ORT-VX-Za-eg-ort-vx-z\u212a]|(?i:T)[C-DF-HJ-PRTV-WZc-df-hj-prtv-wz\u212a]|(?i:U)[AGKSY-Zagksy-z\u017f\u212a]|(?i:V)[ACEGINUaceginu]|(?i:W)[FSfs\u017f]|(?i:Y)[ETet]|(?i:Z)[AMWamw])(?-m:$)")
	urlPattern          = regexp.MustCompile("([^#-\\$0-9@-Za-z\u017f\\x{202a}-\\x{202e}\u212a\\x{feff}\uff03\uff20\\x{fffe}-\\x{ffff}]|^)(((?i:HTTP)(?i:S)?(?i:://))?((?:(?:[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}][^\\t-\\r -!#-,\\.-/:-@\\[-\\^\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}]*)?[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}](?i:\\.))*(?:[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}][^\\t-\\r -!#-,\\.-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}]*)?[^\\t-\\r -!#-/:-@\\[-_\\{-~\\x85\\xa0\\x{1680}\\x{180e}\\x{2000}-\\x{200a}\\x{2028}-\\x{202f}\\x{205f}\\x{3000}\\x{feff}\\x{fffe}-\\x{ffff}](?i:\\.)(?:(?i:A)(?:(?i:C)(?:(?i:ADEMY)|(?i:COUNTANTS)|(?i:TOR))|(?i:ERO)|(?i:GENCY)|(?i:IRFORCE)|(?i:R)(?:(?i:CHI)|(?i:PA))|(?i:S)(?:(?i:IA)|(?i:SOCIATES))|(?i:XA))|(?i:B)(?:(?i:A)(?:(?i:R)(?:(?:)|(?i:GAINS))|(?i:YERN))|(?i:E)(?:(?i:RLIN)|(?i:ST))|(?i:I)(?:(?i:D)|(?i:KE)|(?i:Z))|(?i:L)(?:(?i:ACK)(?:(?:)|(?i:FRIDAY))|(?i:UE))|(?i:OUTIQUE)|(?i:U)(?:(?i:ILD)(?:(?:)|(?i:ERS))|(?i:ZZ)))|(?i:C)(?:(?i:A)(?:(?i:B)|(?i:M)(?:(?i:ERA)|(?i:P))|(?i:PITAL)|(?i:R)(?:(?i:DS)|(?i:E)(?:(?:)|(?i:ER)(?:(?:)|(?i:S))))|(?i:SH)|(?i:T)(?:(?:)|(?i:ERING)))|(?i:E)(?:(?i:NTER)|(?i:O))|(?i:H)(?:(?i:EAP)|(?i:RISTMAS))|(?i:ITIC)|(?i:L)(?:(?i:AIMS)|(?i:EANING)|(?i:INIC)|(?i:OTHING)|(?i:UB))|(?i:O)(?:(?i:DES)|(?i:FFEE)|(?i:L)(?:(?i:LEGE)|(?i:OGNE))|(?i:M)(?:(?:)|(?i:MUNITY)|(?i:P)(?:(?i:ANY)|(?i:UTER)))|(?i:N)(?:(?i:DOS)|(?i:S)(?:(?i:TRUCTION)|(?i:ULTING))|(?i:TRACTORS))|(?i:O)(?:(?i:KING)|[LPlp])|(?i:UNTRY))|(?i:R)(?:(?i:EDIT)(?:(?:)|(?i:CARD))|(?i:UISES)))|(?i:D)(?:(?i:A)(?:(?i:NCE)|(?i:TING))|(?i:E)(?:(?i:MOCRAT)|(?i:NTAL)|(?i:SI))|(?i:I)(?:(?i:AMONDS)|(?i:GITAL)|(?i:RECTORY)|(?i:SCOUNT))|(?i:NP)|(?i:OMAINS))|(?i:E)(?:(?i:DU)(?:(?:)|(?i:CATION))|(?i:MAIL)|(?i:N)(?:(?i:GINEERING)|(?i:TERPRISES))|(?i:QUIPMENT)|(?i:STATE)|(?i:US)|(?i:VENTS)|(?i:X)(?:(?i:CHANGE)|(?i:P)(?:(?i:ERT)|(?i:OSED))))|(?i:F)(?:(?i:A)(?:(?i:IL)|(?i:RM))|(?i:EEDBACK)|(?i:I)(?:(?i:NANC)(?:(?i:E)|(?i:IAL))|(?i:SH)(?:(?:)|(?i:ING))|(?i:TNESS))|(?i:L)(?:(?i:IGHTS)|(?i:ORIST))|(?i:O)(?:(?i:O)|(?i:UNDATION))|(?i:ROGANS)|(?i:U)(?:(?i:ND)|(?i:RNITURE)|(?i:TBOL)))|(?i:G)(?:(?i:AL)(?:(?:)|(?i:LERY))|(?i:IFT)|(?i:L)(?:(?i:ASS)|(?i:OBO))|(?i:MO)|(?i:O)[PVpv]|(?i:R)(?:(?i:A)(?:(?i:PHICS)|(?i:TIS))|(?i:IPE))|(?i:U)(?:(?i:ITARS)|(?i:RU)))|(?i:H)(?:(?i:AUS)|(?i:O)(?:(?i:L)(?:(?i:DINGS)|(?i:IDAY))|(?i:RSE)|(?i:USE)))|(?i:I)(?:(?i:MMOBILIEN)|(?i:N)(?:(?i:DUSTRIES)|(?i:FO)|(?i:K)|(?i:S)(?:(?i:TITUTE)|(?i:URE))|(?i:T)(?:(?:)|(?i:ERNATIONAL))|(?i:VESTMENTS)))|(?i:J)(?:(?i:ETZT)|(?i:OBS))|(?i:K)(?:(?i:AUFEN)|(?i:I)(?:(?i:M)|(?i:TCHEN)|(?i:WI))|(?i:OELN)|(?i:RED))|(?i:L)(?:(?i:AND)|(?i:EASE)|(?i:I)(?:(?i:GHTING)|(?i:M)(?:(?i:ITED)|(?i:O))|(?i:NK))|(?i:ONDON)|(?i:UXURY))|(?i:M)(?:(?i:A)(?:(?i:ISON)|(?i:N)(?:(?i:AGEMENT)|(?i:GO))|(?i:RKETING))|(?i:E)(?:(?i:DIA)|(?i:ET)|(?i:NU))|(?i:I)(?:(?i:AMI)|(?i:L))|(?i:O)(?:(?i:BI)|(?i:DA)|(?i:E)|(?i:NASH)|(?i:SCOW))|(?i:USEUM))|(?i:N)(?:(?i:A)(?:(?i:GOYA)|(?i:ME))|(?i:E)(?:(?i:T)|(?i:USTAR))|(?i:INJA)|(?i:YC))|(?i:O)(?:(?i:KINAWA)|(?i:NL)|(?i:RG))|(?i:P)(?:(?i:AR)(?:(?i:IS)|(?i:T)(?:(?i:NERS)|(?i:S)))|(?i:HOTO)(?:(?:)|(?i:GRAPHY)|(?i:S))|(?i:I)(?:(?i:C)(?:(?i:S)|(?i:TURES))|(?i:NK))|(?i:LUMBING)|(?i:OST)|(?i:RO)(?:(?:)|(?i:DUCTIONS)|(?i:PERTIES))|(?i:UB))|(?i:Q)(?:(?i:PON)|(?i:UEBEC))|(?i:R)(?:(?i:E)(?:(?i:CIPES)|(?i:D)|(?i:ISEN)|(?i:N)(?:(?:)|(?i:TALS))|(?i:P)(?:(?i:AIR)|(?i:ORT))|(?i:ST)|(?i:VIEWS))|(?i:ICH)|(?i:O)(?:(?i:CKS)|(?i:DEO))|(?i:UHR)|(?i:YUKYU))|(?i:S)(?:(?i:AARLAND)|(?i:CHULE)|(?i:E)(?:(?i:RVICES)|(?i:XY))|(?i:H)(?:(?i:IKSHA)|(?i:OES))|(?i:INGLES)|(?i:O)(?:(?i:CIAL)|(?i:HU)|(?i:L)(?:(?i:AR)|(?i:UTIONS))|(?i:Y))|(?i:U)(?:(?i:PP)(?:(?i:L)(?:(?i:IES)|(?i:Y))|(?i:ORT))|(?i:RGERY))|(?i:YSTEMS))|(?i:T)(?:(?i:A)(?:(?i:TTOO)|(?i:X))|(?i:E)(?:(?i:CHNOLOGY)|(?i:L))|(?i:I)(?:(?i:ENDA)|(?i:PS))|(?i:O)(?:(?i:DAY)|(?i:KYO)|(?i:OLS)|(?i:WN)|(?i:YS))|(?i:RA)(?:(?i:DE)|(?i:INING)|(?i:VEL)))|(?i:UN)(?:(?i:IVERSITY)|(?i:O))|(?i:V)(?:(?i:ACATIONS)|(?i:E)(?:(?i:GAS)|(?i:NTURES))|(?i:I)(?:(?i:AJES)|(?i:LLAS)|(?i:SION))|(?i:O)(?:(?i:DKA)|(?i:T)(?:(?i:E)|(?i:ING)|(?i:O))|(?i:YAGE)))|(?i:W)(?:(?i:A)(?:(?i:NG)|(?i:TCH))|(?i:E)(?:(?i:BCAM)|(?i:D))|(?i:I)(?:(?i:EN)|(?i:KI))|(?i:ORKS)|(?i:T)[CFcf])|(?i:X)(?:(?i:XX)|(?i:YZ))|(?i:YOKOHAMA)|(?i:ZONE)|(?i:A)[C-GIL-OQ-UW-XZc-gil-oq-uw-xz\u017f]|(?i:B)[A-BD-JM-OR-TV-WY-Za-bd-jm-or-tv-wy-z\u017f]|(?i:C)[AC-DF-IK-ORU-Zac-df-ik-oru-z\u212a]|(?i:D)[EJ-KMOZej-kmoz\u212a]|(?i:E)[CEGR-Ucegr-u\u017f]|(?i:F)[I-KMORi-kmor\u212a]|(?i:G)[A-BD-IL-NP-UWYa-bd-il-np-uwy\u017f]|(?i:H)[KM-NRT-Ukm-nrt-u\u212a]|(?i:I)[D-EL-OQ-Td-el-oq-t\u017f]|(?i:J)[EMO-Pemo-p]|(?i:K)[EG-IM-NPRWY-Zeg-im-nprwy-z]|(?i:L)[A-CIKR-VYa-cikr-vy\u017f\u212a]|(?i:M)[AC-EG-HK-Zac-eg-hk-z\u017f\u212a]|(?i:N)[ACE-GILO-PRUZace-gilo-pruz]|(?i:OM)|(?i:P)[AE-HK-NR-TWYae-hk-nr-twy\u017f\u212a]|(?i:QA)|(?i:R)[EOSUWeosuw\u017f]|(?i:S)[A-EG-ORT-VX-Za-eg-ort-vx-z\u212a]|(?i:T)[C-DF-HJ-PRTV-WZc-df-hj-prtv-wz\u212a]|(?i:U)[AGKSY-Zagksy-z\u017f\u212a]|(?i:V)[ACEGINUaceginu]|(?i:W)[FSfs\u017f]|(?i:Y)[ETet]|(?i:Z)[AMWamw]|(?i:\u96c6\u56e2)|(?i:\u5728\u7ebf)|(?i:\ud55c\uad6d)|(?i:\u09ad\u09be\u09b0\u09a4)|(?i:\u516c)[\u53f8\u76ca]|(?i:\u79fb\u52a8)|(?i:\u6211\u7231\u4f60)|(?i:\u041c\u041e\u0421\u041a\u0412\u0410)|(?i:\u049a\u0410\u0417)|(?i:\u041e\u041d\u041b\u0410\u0419\u041d)|(?i:\u0421)(?:(?i:\u0410\u0419\u0422)|(?i:\u0420\u0411))|(?i:\u041e\u0420\u0413)|(?i:\uc0bc\uc131)|(?i:\u0b9a\u0bbf\u0b99\u0bcd\u0b95\u0baa\u0bcd\u0baa\u0bc2\u0bb0\u0bcd)|(?i:\u5546\u57ce)|(?i:\u0414\u0415\u0422\u0418)|(?i:\u4e2d)(?:(?i:\u6587\u7f51)|[\u4fe1\u56fd\u570b])|(?i:\u0c2d\u0c3e\u0c30\u0c24\u0c4d)|(?i:\u0dbd\u0d82\u0d9a\u0dcf)|(?i:\u0aad\u0abe\u0ab0\u0aa4)|(?i:\u092d\u093e\u0930\u0924)|(?i:\u0938\u0902\u0917\u0920\u0928)|(?i:\u7f51\u7edc)|(?i:\u0423\u041a\u0420)|(?i:\u9999\u6e2f)|(?i:\u53f0)[\u6e7e\u7063]|(?i:\u041c\u041e\u041d)|(?i:\u0627\u0644\u062c\u0632\u0627\u0626\u0631)|(?i:\u0639\u0645\u0627\u0646)|(?i:\u0627)(?:(?i:\u06cc\u0631\u0627\u0646)|(?i:\u0645\u0627\u0631\u0627\u062a))|(?i:\u0628\u0627\u0632\u0627\u0631)|(?i:\u0627\u0644\u0627\u0631\u062f\u0646)|(?i:\u0628\u06be\u0627\u0631\u062a)|(?i:\u0627\u0644)(?:(?i:\u0645\u063a\u0631\u0628)|(?i:\u0633\u0639\u0648\u062f\u064a\u0629))|(?i:\u0645\u0644\u064a\u0633\u064a\u0627)|(?i:\u0634\u0628\u0643\u0629)|(?i:\u673a\u6784)|(?i:\u7ec4\u7ec7\u673a\u6784)|(?i:\u0e44\u0e17\u0e22)|(?i:\u0633\u0648\u0631\u064a\u0629)|(?i:\u0420\u0424)|(?i:\u062a\u0648\u0646\u0633)|(?i:\u307f\u3093\u306a)|(?i:\u4e16\u754c)|(?i:\u0a2d\u0a3e\u0a30\u0a24)|(?i:\u7f51\u5740)|(?i:\u6e38\u620f)|(?i:\u0645\u0635\u0631)|(?i:\u0642\u0637\u0631)|(?i:\u0b87)(?:(?i:\u0bb2\u0b99\u0bcd\u0b95\u0bc8)|(?i:\u0ba8\u0bcd\u0ba4\u0bbf\u0baf\u0bbe))|(?i:\u65b0\u52a0\u5761)|(?i:\u0641\u0644\u0633\u0637\u064a\u0646)|(?i:\u653f\u52a1)|(?i:XN--)[0-9A-Za-z\u017f\u212a]+))(?:(?i::)[0-9]+)?((?i:/)(?:[!#-'\\*-;=@-\\[\\]_a-z\\|~\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]*(?:(?i:\\()[!#-'\\*-;=@-\\[\\]_a-z\\|~\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]+(?i:\\))[!#-'\\*-;=@-\\[\\]_a-z\\|~\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]*)*[#\\+\\-/-9=A-Z_a-z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]|(?i:\\()[!#-'\\*-;=@-\\[\\]_a-z\\|~\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]+(?i:\\))|(?i:@)[!#-'\\*-;=@-\\[\\]_a-z\\|~\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u024f\u0253-\u0254\u0256-\u0257\u0259\u025b\u0260\u0263\u0268-\u0269\u026f\u0272\u0275\u0280\u0283\u0288-\u028c\u0292\u02bb\u0300-\u036f\u0399\u03b9\u1e00-\u1eff\u1fbe\u212a-\u212b\u2c65-\u2c66\u2c7e-\u2c7f]+(?i:/))*)?(?:(?i:\\?)[!#-;=\\?-\\[\\]_a-z\\|~\u017f\u212a]*[#&/-9=A-Z_a-z\u017f\u212a])?)")
)
