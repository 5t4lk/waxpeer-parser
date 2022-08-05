package csgobackpack

type Welcome struct {
	Success   bool                 `json:"success"`
	Currency  string               `json:"currency"`
	Timestamp int64                `json:"timestamp"`
	ItemsList map[string]ItemsList `json:"items_list"`
}

type ItemsList struct {
	Name          string      `json:"name"`
	Marketable    int64       `json:"marketable"`
	Tradable      int64       `json:"tradable"`
	Classid       string      `json:"classid"`
	IconURL       string      `json:"icon_url"`
	IconURLLarge  string      `json:"icon_url_large"`
	Type          Type        `json:"type"`
	Rarity        Rarity      `json:"rarity"`
	RarityColor   RarityColor `json:"rarity_color"`
	Price         Price       `json:"price,omitempty"`
	FirstSaleDate string      `json:"first_sale_date,omitempty"`
	WeaponType    WeaponType  `json:"weapon_type,omitempty"`
	GunType       GunType     `json:"gun_type,omitempty"`
	Exterior      Exterior    `json:"exterior,omitempty"`
	KnifeType     KnifeType   `json:"knife_type,omitempty"`
	Souvenir      int64       `json:"souvenir,omitempty"`
	Tournament    Tournament  `json:"tournament,omitempty"`
	Stattrak      int64       `json:"stattrak,omitempty"`
	Sticker       int64       `json:"sticker,omitempty"`
}

type Price struct {
	The24_Hours The24__Hours `json:"24_hours,omitempty"`
	The7_Days   The24__Hours `json:"7_days,omitempty"`
	The30_Days  The24__Hours `json:"30_days,omitempty"`
	AllTime     The24__Hours `json:"all_time"`
}

type The24__Hours struct {
	Average           float64 `json:"average"`
	Median            float64 `json:"median"`
	Sold              string  `json:"sold"`
	StandardDeviation string  `json:"standard_deviation"`
	LowestPrice       float64 `json:"lowest_price"`
	HighestPrice      float64 `json:"highest_price"`
}

type Exterior string

const (
	BattleScarred Exterior = "Battle-Scarred"
	FactoryNew    Exterior = "Factory New"
	FieldTested   Exterior = "Field-Tested"
	MinimalWear   Exterior = "Minimal Wear"
	NotPainted    Exterior = "Not Painted"
	WellWorn      Exterior = "Well-Worn"
)

type GunType string

const (
	Ak47         GunType = "AK-47"
	Aug          GunType = "AUG"
	Awp          GunType = "AWP"
	CZ75Auto     GunType = "CZ75-Auto"
	DesertEagle  GunType = "Desert Eagle"
	DualBerettas GunType = "Dual Berettas"
	Famas        GunType = "FAMAS"
	FiveSeveN    GunType = "Five-SeveN"
	G3Sg1        GunType = "G3SG1"
	GalilAR      GunType = "Galil AR"
	Glock18      GunType = "Glock-18"
	M249         GunType = "M249"
	M4A1S        GunType = "M4A1-S"
	M4A4         GunType = "M4A4"
	MAC10        GunType = "MAC-10"
	Mag7         GunType = "MAG-7"
	Mp5SD        GunType = "MP5-SD"
	Mp7          GunType = "MP7"
	Mp9          GunType = "MP9"
	Negev        GunType = "Negev"
	Nova         GunType = "Nova"
	P2000        GunType = "P2000"
	P250         GunType = "P250"
	P90          GunType = "P90"
	PPBizon      GunType = "PP-Bizon"
	R8Revolver   GunType = "R8 Revolver"
	SawedOff     GunType = "Sawed-Off"
	Scar20       GunType = "SCAR-20"
	Sg553        GunType = "SG 553"
	Ssg08        GunType = "SSG 08"
	Tec9         GunType = "Tec-9"
	Ump45        GunType = "UMP-45"
	UspS         GunType = "USP-S"
	Xm1014       GunType = "XM1014"
)

type KnifeType string

const (
	Bayonet        KnifeType = "Bayonet"
	BowieKnife     KnifeType = "Bowie Knife"
	ButterflyKnife KnifeType = "Butterfly Knife"
	ClassicKnife   KnifeType = "Classic Knife"
	FalchionKnife  KnifeType = "Falchion Knife"
	FlipKnife      KnifeType = "Flip Knife"
	GutKnife       KnifeType = "Gut Knife"
	HuntsmanKnife  KnifeType = "Huntsman Knife"
	Karambit       KnifeType = "Karambit"
	M9Bayonet      KnifeType = "M9 Bayonet"
	NavajaKnife    KnifeType = "Navaja Knife"
	NomadKnife     KnifeType = "Nomad Knife"
	ParacordKnife  KnifeType = "Paracord Knife"
	ShadowDaggers  KnifeType = "Shadow Daggers"
	SkeletonKnife  KnifeType = "Skeleton Knife"
	StilettoKnife  KnifeType = "Stiletto Knife"
	SurvivalKnife  KnifeType = "Survival Knife"
	TalonKnife     KnifeType = "Talon Knife"
	UrsusKnife     KnifeType = "Ursus Knife"
)

type Rarity string

const (
	BaseGrade       Rarity = "Base Grade"
	Classified      Rarity = "Classified"
	ConsumerGrade   Rarity = "Consumer Grade"
	Contraband      Rarity = "Contraband"
	Covert          Rarity = "Covert"
	Distinguished   Rarity = "Distinguished"
	Exceptional     Rarity = "Exceptional"
	Exotic          Rarity = "Exotic"
	Extraordinary   Rarity = "Extraordinary"
	HighGrade       Rarity = "High Grade"
	IndustrialGrade Rarity = "Industrial Grade"
	Master          Rarity = "Master"
	MilSpecGrade    Rarity = "Mil-Spec Grade"
	Remarkable      Rarity = "Remarkable"
	Restricted      Rarity = "Restricted"
	Superior        Rarity = "Superior"
)

type RarityColor string

const (
	B0C3D9    RarityColor = "b0c3d9"
	D32Ce6    RarityColor = "d32ce6"
	E4Ae39    RarityColor = "e4ae39"
	Eb4B4B    RarityColor = "eb4b4b"
	The4B69Ff RarityColor = "4b69ff"
	The5E98D9 RarityColor = "5e98d9"
	The8847Ff RarityColor = "8847ff"
)

type Tournament string

const (
	The2013DreamHackWinter     Tournament = "2013 DreamHack Winter"
	The2014DreamHackWinter     Tournament = "2014 DreamHack Winter"
	The2014EMSOneKatowice      Tournament = "2014 EMS One Katowice"
	The2014ESLOneCologne       Tournament = "2014 ESL One Cologne"
	The2015DreamHackClujNapoca Tournament = "2015 DreamHack Cluj-Napoca"
	The2015ESLOneCologne       Tournament = "2015 ESL One Cologne"
	The2015ESLOneKatowice      Tournament = "2015 ESL One Katowice"
	The2016ESLOneCologne       Tournament = "2016 ESL One Cologne"
	The2016MLGColumbus         Tournament = "2016 MLG Columbus"
	The2017ELEAGUEAtlanta      Tournament = "2017 ELEAGUE Atlanta"
	The2017PGLKrakow           Tournament = "2017 PGL Krakow"
	The2018FACEITLondon        Tournament = "2018 FACEIT London"
	The2019IEMKatowice         Tournament = "2019 IEM Katowice"
	The2019StarLadderBerlin    Tournament = "2019 StarLadder Berlin"
	The2020Rmr                 Tournament = "2020 RMR"
	The2021PGLStockholm        Tournament = "2021 PGL Stockholm"
)

type Type string

const (
	Collectible Type = "Collectible"
	Container   Type = "Container"
	Gift        Type = "Gift"
	Gloves      Type = "Gloves"
	Graffiti    Type = "Graffiti"
	Key         Type = "Key"
	MusicKit    Type = "Music Kit"
	Pass        Type = "Pass"
	Weapon      Type = "Weapon"
)

type WeaponType string

const (
	Knife       WeaponType = "Knife"
	Machinegun  WeaponType = "Machinegun"
	Pistol      WeaponType = "Pistol"
	Rifle       WeaponType = "Rifle"
	Shotgun     WeaponType = "Shotgun"
	Smg         WeaponType = "SMG"
	SniperRifle WeaponType = "Sniper Rifle"
)
