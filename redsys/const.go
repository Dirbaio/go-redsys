package redsys

type Language int

const (
	LanguageDefault    Language = 0   // Por defecto
	LanguageSpanish    Language = 1   // Español
	LanguageEnglish    Language = 2   // English - Inglés
	LanguageCatalan    Language = 3   // Català
	LanguageFrench     Language = 4   // Français - Frances
	LanguageGerman     Language = 5   // Deutsch - Aleman
	LanguageDutch      Language = 6   // Nederlands - Holandes
	LanguageItalian    Language = 7   // Italiano
	LanguageSwedish    Language = 8   // Svenska - Sueco
	LanguagePortuguese Language = 9   // Português
	LanguageValencian  Language = 10  // Valencià
	LanguagePolish     Language = 11  // Polski - Polaco
	LanguageGalician   Language = 12  // Galego
	LanguageEuskera    Language = 13  // Euskara
	LanguageBulgarian  Language = 100 // български език - Bulgaro
	LanguageChinese    Language = 156 // Chino
	LanguageCroatian   Language = 191 // Hrvatski - Croata
	LanguageCzech      Language = 203 // Čeština - Checo
	LanguageDanish     Language = 208 // Dansk - Danes
	LanguageEstonian   Language = 233 // Eesti keel - Estonio
	LanguageFinnish    Language = 246 // Suomi - Finlandes
	LanguageGreek      Language = 300 // ελληνικά - Griego
	LanguageHungarian  Language = 348 // Magyar - Hungaro
	LanguageJapanese   Language = 392 // Japonés
	LanguageLatvian    Language = 428 // Latviešu valoda - Leton
	LanguageLithuanian Language = 440 // Lietuvių kalba - Lituano
	LanguageMaltese    Language = 470 // Malti - Maltés
	LanguageRomanian   Language = 642 // Română - Rumano
	LanguageRussian    Language = 643 // ру́сский язы́к – Ruso
	LanguageSlovak     Language = 703 // Slovenský jazyk - Eslovaco
	LanguageSlovenian  Language = 705 // Slovenski jezik - Esloveno
	LanguageTurkish    Language = 792 // Türkçe - Turco
)

type OperationType int

const (
	OperationTypeCharge OperationType = 0

	OperationTypeRefund OperationType = 3

	OperationTypePreauth        OperationType = 1
	OperationTypePreauthConfirm OperationType = 2
	OperationTypePreauthVoid    OperationType = 9

	OperationTypeSplitAuth    OperationType = 7
	OperationTypeSplitConfirm OperationType = 8
)

type Currency int

// Source: https://sis-d.redsys.es/pagosonline/codigos-autorizacion.html#monedas
// Slightly cleaned up (some duplicate or wrong entries.)
// The codes seem to be the same as ISO 4217 https://en.wikipedia.org/wiki/ISO_4217
const (
	CurrencyAED Currency = 784 // UAE DIRHAM AED
	CurrencyAFN Currency = 971 // AFGHANISTAN AFGHANI AFN
	CurrencyALL Currency = 8   // LEK ALL
	CurrencyAMD Currency = 51  // ARMENIAN DRAM AMD
	CurrencyANG Currency = 532 // NETHERLANDS ANTILLIA ANG
	CurrencyAOA Currency = 973 // KWANZA ANGOLA AOA
	CurrencyAOK Currency = 24  // ANGOLA KWANZA AOK
	CurrencyARP Currency = 32  // ARGENTINE AUSTRAL ARP
	CurrencyAUD Currency = 36  // AUSTRALIAN DOLLAR AUD
	CurrencyAWG Currency = 533 // ARUBA AWG
	CurrencyAZM Currency = 31  // AZERBAIJANIAN MANAT AZM
	CurrencyAZN Currency = 944 // AZERBAIJANIAN MANAT AZN
	CurrencyBAD Currency = 70  // DINAR BAD
	CurrencyBAM Currency = 977 // BOSNIAN MARKA BAM
	CurrencyBBD Currency = 52  // BARBADOS DOLLAR BBD
	CurrencyBDT Currency = 50  // TAKA BDT
	CurrencyBGL Currency = 100 // LEV BGL
	CurrencyBGN Currency = 975 // NEW LEV BGN
	CurrencyBHD Currency = 48  // BAHRAINI DINAR BHD
	CurrencyBIF Currency = 108 // BURUNDI FRANC BIF
	CurrencyBMD Currency = 60  // BERMUDAN DOLLAR BMD
	CurrencyBND Currency = 96  // BRUNEI DOLLAR BND
	CurrencyBOP Currency = 68  // BOLIVIAN PESO BOP
	CurrencyBRC Currency = 76  // CRUZEIRO BRC
	CurrencyBRL Currency = 986 // BRAZILIAN REAL BRL
	CurrencyBSD Currency = 44  // BAHAMIAN DOLLAR BSD
	CurrencyBTN Currency = 64  // NGULTRUM BTN
	CurrencyBUK Currency = 104 // KYAT BUK
	CurrencyBWP Currency = 72  // PULA BWP
	CurrencyBYB Currency = 112 // BELARUSSIAN RUBLE BYB
	CurrencyBYR Currency = 974 // BELARUSSIAN RUBLE BYR
	CurrencyBZD Currency = 84  // BELIZE DOLLAR
	CurrencyCAD Currency = 124 // CANADIAN DOLLAR CAD
	CurrencyCDF Currency = 976 // FRANCO DEL CONGO CDF
	CurrencyCHF Currency = 756 // SWISS FRANC CHF
	CurrencyCLP Currency = 152 // CHILEAN PESO CLP
	CurrencyCNH Currency = 157 // CHINESE RENMIMBI CNH
	CurrencyCNY Currency = 156 // YUAN RENMINBI CNY
	CurrencyCOP Currency = 170 // COLOMBIAN PESO COP
	CurrencyCRC Currency = 188 // COSTA RICA COLON CRC
	CurrencyCSK Currency = 200 // KORUNA CSK
	CurrencyCUP Currency = 192 // CUBAN PESO CUP
	CurrencyCVE Currency = 132 // CAPE VERDE ESCUDO CVE
	CurrencyCYP Currency = 196 // CYPRUS POUND CYP
	CurrencyCZK Currency = 203 // CZECH KORUNA CZK
	CurrencyDDM Currency = 278 // MARK DER DDR DDM
	CurrencyDJF Currency = 262 // DJIBOUTI FRANC DJF
	CurrencyDKK Currency = 208 // DANISH KRONE DKK
	CurrencyDOP Currency = 214 // DOMINICAN PESO DOP
	CurrencyDZD Currency = 12  // ALGERIAN DINAR DZD
	CurrencyECS Currency = 218 // SUCRE ECS
	CurrencyEEK Currency = 233 // ESTONIAN KROON EEK
	CurrencyEGP Currency = 818 // EGYPTIAN POUND EGP
	CurrencyERN Currency = 232 // ERITREAN NAKTAN ERN
	CurrencyETB Currency = 230 // ETHIOPIAN BIRR ETB
	CurrencyEUR Currency = 978 // EURO
	CurrencyFJD Currency = 242 // FIJI DOLLAR FJD
	CurrencyFKP Currency = 238 // FALKLAND ISLANDS FKP
	CurrencyGBP Currency = 826 // POUND STERLING GBP
	CurrencyGEL Currency = 981 // GEORGIAN LARI GEL
	CurrencyGHC Currency = 288 // GHANA CEDI GHC
	CurrencyGHS Currency = 936 // GHANA CEDI GHS
	CurrencyGIP Currency = 292 // GIBRALTAR POUND GIP
	CurrencyGMD Currency = 270 // DALASI GMD
	CurrencyGNF Currency = 624 // GUINEAN FRANC
	CurrencyGNS Currency = 324 // SYLI GNS
	CurrencyGQE Currency = 226 // EKWELE GQE
	CurrencyGTQ Currency = 320 // QUETZAL GTQ
	CurrencyGYD Currency = 328 // GUYANA DOLLAR GYD
	CurrencyHKD Currency = 344 // HONG KONG DOLLAR HKD
	CurrencyHNL Currency = 340 // LEMPIRA HNL
	CurrencyHRK Currency = 191 // CROATIAN KUNA HRKY
	CurrencyHTG Currency = 332 // GOURDE HTG
	CurrencyHUF Currency = 348 // FORINT HUF
	CurrencyIDR Currency = 360 // RUPIAH IDR
	CurrencyILS Currency = 376 // ISRAEL SHEKEL ILS
	CurrencyINR Currency = 356 // INDIAN RUPEE INR
	CurrencyIQD Currency = 368 // IRAQI DINAR IQD
	CurrencyIRA Currency = 365 // IRANIAN AIRLINE RATE IRA
	CurrencyIRR Currency = 364 // IRANIAL RIAL IRR
	CurrencyISK Currency = 352 // ICELAND KRONA ISK
	CurrencyJMD Currency = 388 // JAMAICAN DOLLAR JMD
	CurrencyJOD Currency = 400 // JORDANIAN DINAR JOD
	CurrencyJPY Currency = 392 // YEN JPY
	CurrencyKES Currency = 404 // KENYAN SHILLING KES
	CurrencyKGS Currency = 417 // KYRGYZSTAN SON KGS
	CurrencyKHR Currency = 116 // RIEL KHR
	CurrencyKMF Currency = 174 // COMOROS FRANC KMF
	CurrencyKPW Currency = 408 // NORTH KOREAN WON KPW
	CurrencyKRW Currency = 410 // KOREAN WON KRW
	CurrencyKWD Currency = 414 // KUWAITI DINAR KWD
	CurrencyKYD Currency = 136 // CAYMAN ISLANDS DOLLA KYDY
	CurrencyKZT Currency = 398 // TENGE KZT
	CurrencyLAK Currency = 418 // KIP LAK
	CurrencyLBP Currency = 422 // LEBANESE POUND LBP
	CurrencyLKR Currency = 144 // SRI LANKA RUPEE LKR
	CurrencyLRD Currency = 430 // LIBERIAN DOLLAR LRD
	CurrencyLSM Currency = 426 // LESOTHO LOTI LSM
	CurrencyLTL Currency = 440 // LITHUANIAN LITAS LTL
	CurrencyLVL Currency = 428 // LATVIAN LAT LVL
	CurrencyLYD Currency = 434 // LIBYAN DINAR LYD
	CurrencyMAD Currency = 504 // MORROCAN DIRHAM MAD
	CurrencyMDL Currency = 498 // MOLDOVIAN LEU MDL
	CurrencyMGA Currency = 969 // ARIARY MGA
	CurrencyMGF Currency = 450 // MALAGASY FRANC MGF
	CurrencyMKD Currency = 807 // MACEDONIAN DENAR MKD
	CurrencyMLF Currency = 466 // MALI MLF
	CurrencyMNT Currency = 496 // TUGRIK MNT
	CurrencyMON Currency = 30  // PROBANDO DESA MON
	CurrencyMOP Currency = 446 // PATACA MOP
	CurrencyMRO Currency = 478 // OUGUIYA MRO
	CurrencyMTL Currency = 470 // MALTESE LIRA MTL
	CurrencyMUR Currency = 480 // MAURITIUS RUPEE MUR
	CurrencyMVR Currency = 462 // MALDIVE RUPEE MVR
	CurrencyMWK Currency = 454 // MALAWI KWACHA MWK
	CurrencyMXP Currency = 484 // MEXICAN PESO MXP
	CurrencyMYR Currency = 458 // MALASYAN RINGGIT MYR
	CurrencyMZM Currency = 508 // METICAL MZM
	CurrencyMZN Currency = 943 // MOZAMBIQUE METICAL MZN
	CurrencyNAD Currency = 516 // NAMIBIAN DOLLAR NAD
	CurrencyNGN Currency = 566 // NAIRA NGN
	CurrencyNIC Currency = 558 // CORDOBA NIC
	CurrencyNOK Currency = 578 // NORWEGIAN KRONE NOK
	CurrencyNPR Currency = 524 // NEPALESE RUPEE NPR
	CurrencyNTZ Currency = 536 // YUGOSLAVIAN NEW DIAN NTZ
	CurrencyNZD Currency = 554 // NEW ZEALAND DOLLAR NZD
	CurrencyOMR Currency = 512 // RIAL OMANI OMR
	CurrencyPAB Currency = 590 // BALBOA PAB
	CurrencyPCI Currency = 582 // PACIFIC ISLAND PCI
	CurrencyPEI Currency = 604 // PERU INTI PEI
	CurrencyPGK Currency = 598 // KINA PGK
	CurrencyPHP Currency = 608 // PHILIPPINE PESO PHP
	CurrencyPKR Currency = 586 // PAKISTAN RUPEE PKR
	CurrencyPLN Currency = 985 // NEW POLISH ZLOTY PLN
	CurrencyPLZ Currency = 616 // ZLOTY PLZ
	CurrencyPTL Currency = 793 // PSEUDO TURKISH LIRA PTL
	CurrencyPYG Currency = 600 // GUARANI PYG
	CurrencyQAR Currency = 634 // QATARI RIAL QAR
	CurrencyROL Currency = 642 // LEU ROL
	CurrencyRON Currency = 946 // NEW LEU RON
	CurrencyRSD Currency = 941 // DINAR SERBIO RSD
	CurrencyRUB Currency = 643 // RUSSIAN ROUBLE RUB
	CurrencyRUR Currency = 810 // RUSSIAN ROUBLE RUR
	CurrencyRWF Currency = 646 // RWANDA FRANC RWF
	CurrencySAR Currency = 682 // SAUDI RIYAL SAR
	CurrencySBD Currency = 90  // SOLOMON ISLANDS DOLL SBD
	CurrencySCR Currency = 690 // SEYCHELLES RUPEE SCR
	CurrencySDA Currency = 737 // SUDAN AIRLINES SDA
	CurrencySDP Currency = 736 // SUDANESE POUND SDP
	CurrencySEK Currency = 752 // SWEDISH KRONA SEK
	CurrencySGD Currency = 702 // SINGAPORE DOLLAR SGD
	CurrencySHP Currency = 654 // ST.HELENA POUND SHP
	CurrencySIT Currency = 705 // SLOVENIAN TOLAR SIT
	CurrencySKK Currency = 703 // SLOVAK KORUNA SKK
	CurrencySLL Currency = 694 // LEONE SLL
	CurrencySOS Currency = 706 // SOMALI SHILLING SOS
	CurrencySRD Currency = 968 // SURINAME DOLLAR SRD
	CurrencySRG Currency = 740 // SURINAM GUILDER SRG
	CurrencySSP Currency = 728 // SOUTH SUDANESE POUND SSP
	CurrencySTD Currency = 678 // DOBRA STD
	CurrencySVC Currency = 222 // EL SALVADOR COLON SVC
	CurrencySYP Currency = 760 // SYRIAN POUND SYP
	CurrencySZL Currency = 748 // LILANGENI SZL
	CurrencyTHB Currency = 764 // BAHT THB
	CurrencyTJR Currency = 762 // TAJIK RUBLE TJR
	CurrencyTJS Currency = 972 // TAJIKISTAN SOMONI TJS
	CurrencyTMM Currency = 795 // MANAT TMM
	CurrencyTMT Currency = 934 // NEW MANAT TMT
	CurrencyTND Currency = 788 // TUNISIAN DINAR TND
	CurrencyTOP Currency = 776 // PA'ANGA TOP
	CurrencyTPE Currency = 626 // TIMOR ESCUDO TPE
	CurrencyTRL Currency = 792 // TURKISH LIRA TRL
	CurrencyTRY Currency = 949 // TURKISH LIRA TRY
	CurrencyTTD Currency = 780 // TRINIDAD Y TOBAGO DO TTD
	CurrencyTWD Currency = 901 // NEW TAIWAN DOLLAR TWD
	CurrencyTZS Currency = 834 // TANZANIAN SHILLING TZS
	CurrencyUAH Currency = 980 // HRYVNIA UAH
	CurrencyUAK Currency = 804 // KARBOVANET UAK
	CurrencyUGS Currency = 800 // UGANDA SHILLING UGS
	CurrencyUSD Currency = 840 // DOLAR U.S.A. USD
	CurrencyUYP Currency = 858 // URUGUAYAN PESO UYP
	CurrencyUZS Currency = 860 // UZBEKISTAN SUM UZS
	CurrencyVEB Currency = 862 // BOLIVAR VEB
	CurrencyVND Currency = 704 // DONG VND
	CurrencyVUV Currency = 548 // VANUATU VATU VUV
	CurrencyWST Currency = 882 // TALA WST
	CurrencyXAF Currency = 950 // CFA FRANC BEAC XAF
	CurrencyXCD Currency = 951 // EAST CARIBBEAN DOLLA XCD
	CurrencyXEU Currency = 954 // E.C.U. EUROPEAN CUR XEU
	CurrencyXOF Currency = 952 // CFA FRANC BCEAO XOF
	CurrencyXPF Currency = 953 // CFP FRANC XPF
	CurrencyYDD Currency = 720 // YEMENI DINAR YDD
	CurrencyYER Currency = 886 // YEMINI RIAL YER
	CurrencyYUD Currency = 890 // NEW YUGOSLAVIAN DOLL YUD
	CurrencyYUG Currency = 891 // NEW DINAR YUG
	CurrencyZAL Currency = 991 // RAND FINANCIER ZAL
	CurrencyZAR Currency = 710 // RAND ZAR
	CurrencyZMK Currency = 892 // KWACHA ZMK
	CurrencyZMW Currency = 967 // KWACHA ZMW
	CurrencyZRZ Currency = 180 // ZAIRE ZRZ
	CurrencyZWD Currency = 716 // ZIMBABWE DOLLAR ZWD
)
