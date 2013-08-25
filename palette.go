package BrickMosaic

import (
	//  "fmt"
	//	"image"
	"image/color"
)

type BrickColor struct {
	id   int
	name string
	c    color.Color
}

func (c BrickColor) RGBA() (r, g, b, a uint32) {
	return c.c.RGBA()
}

var (
	Red = color.RGBA{uint8(200), 0, 0, 0}

	// TODO(ndunn): Pull in the bricklink colors. http://www.peeron.com/cgi-bin/invcgis/inv/colors?PagerSortDir=f&PagerSortCol=BLName&PagerSortRev=0

	// All color RGB values from
	// http://www.peeron.com/cgi-bin/invcgis/colorguide.cgi
	// cat ~/Dropbox/BrickColors.txt  | awk '{print $2 " = BrickColor{name: \"" $2 "\", c:color.ARGB{R:uint8(" $7 "), G:uint8(" $8 "), B:uint8(" $9 "), A:uint8(0)}}"}' | mate
	White                 = BrickColor{id: 1, name: "White", c: color.RGBA{R: uint8(242), G: uint8(243), B: uint8(242), A: uint8(0)}}
	Grey                  = BrickColor{id: 2, name: "Grey", c: color.RGBA{R: uint8(161), G: uint8(165), B: uint8(162), A: uint8(0)}}
	LightYellow           = BrickColor{id: 3, name: "LightYellow", c: color.RGBA{R: uint8(249), G: uint8(233), B: uint8(153), A: uint8(0)}}
	BrickYellow           = BrickColor{id: 5, name: "BrickYellow", c: color.RGBA{R: uint8(215), G: uint8(197), B: uint8(153), A: uint8(0)}}
	LightGreen            = BrickColor{id: 6, name: "LightGreen", c: color.RGBA{R: uint8(194), G: uint8(218), B: uint8(184), A: uint8(0)}}
	LightReddishViolet    = BrickColor{id: 9, name: "LightReddishViolet", c: color.RGBA{R: uint8(232), G: uint8(186), B: uint8(199), A: uint8(0)}}
	LightOrangeBrown      = BrickColor{id: 12, name: "LightOrangeBrown", c: color.RGBA{R: uint8(203), G: uint8(132), B: uint8(66), A: uint8(0)}}
	Nougat                = BrickColor{id: 18, name: "Nougat", c: color.RGBA{R: uint8(204), G: uint8(142), B: uint8(104), A: uint8(0)}}
	BrightRed             = BrickColor{id: 21, name: "BrightRed", c: color.RGBA{R: uint8(196), G: uint8(40), B: uint8(27), A: uint8(0)}}
	MedReddishViolet      = BrickColor{id: 22, name: "MedReddishViolet", c: color.RGBA{R: uint8(196), G: uint8(112), B: uint8(160), A: uint8(0)}}
	BrightBlue            = BrickColor{id: 23, name: "BrightBlue", c: color.RGBA{R: uint8(13), G: uint8(105), B: uint8(171), A: uint8(0)}}
	BrightYellow          = BrickColor{id: 24, name: "BrightYellow", c: color.RGBA{R: uint8(245), G: uint8(205), B: uint8(47), A: uint8(0)}}
	EarthOrange           = BrickColor{id: 25, name: "EarthOrange", c: color.RGBA{R: uint8(98), G: uint8(71), B: uint8(50), A: uint8(0)}}
	Black                 = BrickColor{id: 26, name: "Black", c: color.RGBA{R: uint8(27), G: uint8(42), B: uint8(52), A: uint8(0)}}
	DarkGrey              = BrickColor{id: 27, name: "DarkGrey", c: color.RGBA{R: uint8(109), G: uint8(110), B: uint8(108), A: uint8(0)}}
	DarkGreen             = BrickColor{id: 28, name: "DarkGreen", c: color.RGBA{R: uint8(40), G: uint8(127), B: uint8(70), A: uint8(0)}}
	MediumGreen           = BrickColor{id: 29, name: "MediumGreen", c: color.RGBA{R: uint8(161), G: uint8(196), B: uint8(139), A: uint8(0)}}
	LightYellowishOrange  = BrickColor{id: 36, name: "LightYellowishOrange", c: color.RGBA{R: uint8(243), G: uint8(207), B: uint8(155), A: uint8(0)}}
	BrightGreen           = BrickColor{id: 37, name: "BrightGreen", c: color.RGBA{R: uint8(75), G: uint8(151), B: uint8(74), A: uint8(0)}}
	DarkOrange            = BrickColor{id: 38, name: "DarkOrange", c: color.RGBA{R: uint8(160), G: uint8(95), B: uint8(52), A: uint8(0)}}
	LightBluishViolet     = BrickColor{id: 39, name: "LightBluishViolet", c: color.RGBA{R: uint8(193), G: uint8(202), B: uint8(222), A: uint8(0)}}
	LightBlue             = BrickColor{id: 45, name: "LightBlue", c: color.RGBA{R: uint8(180), G: uint8(210), B: uint8(227), A: uint8(0)}}
	LightRed              = BrickColor{id: 100, name: "LightRed", c: color.RGBA{R: uint8(238), G: uint8(196), B: uint8(182), A: uint8(0)}}
	MediumRed             = BrickColor{id: 101, name: "MediumRed", c: color.RGBA{R: uint8(218), G: uint8(134), B: uint8(121), A: uint8(0)}}
	MediumBlue            = BrickColor{id: 102, name: "MediumBlue", c: color.RGBA{R: uint8(110), G: uint8(153), B: uint8(201), A: uint8(0)}}
	LightGrey             = BrickColor{id: 103, name: "LightGrey", c: color.RGBA{R: uint8(199), G: uint8(193), B: uint8(183), A: uint8(0)}}
	BrightViolet          = BrickColor{id: 104, name: "BrightViolet", c: color.RGBA{R: uint8(107), G: uint8(50), B: uint8(123), A: uint8(0)}}
	BrightYellowishOrange = BrickColor{id: 105, name: "BrightYellowishOrange", c: color.RGBA{R: uint8(226), G: uint8(155), B: uint8(63), A: uint8(0)}}
	BrightOrange          = BrickColor{id: 106, name: "BrightOrange", c: color.RGBA{R: uint8(218), G: uint8(133), B: uint8(64), A: uint8(0)}}
	BrightBluishGreen     = BrickColor{id: 107, name: "BrightBluishGreen", c: color.RGBA{R: uint8(0), G: uint8(143), B: uint8(155), A: uint8(0)}}
	EarthYellow           = BrickColor{id: 108, name: "EarthYellow", c: color.RGBA{R: uint8(104), G: uint8(92), B: uint8(67), A: uint8(0)}}
	BrightBluishViolet    = BrickColor{id: 110, name: "BrightBluishViolet", c: color.RGBA{R: uint8(67), G: uint8(84), B: uint8(147), A: uint8(0)}}
	MediumBluishViolet    = BrickColor{id: 112, name: "MediumBluishViolet", c: color.RGBA{R: uint8(104), G: uint8(116), B: uint8(172), A: uint8(0)}}
	MedYellowishGreen     = BrickColor{id: 115, name: "MedYellowishGreen", c: color.RGBA{R: uint8(199), G: uint8(210), B: uint8(60), A: uint8(0)}}
	MedBluishGreen        = BrickColor{id: 116, name: "MedBluishGreen", c: color.RGBA{R: uint8(85), G: uint8(165), B: uint8(175), A: uint8(0)}}
	LightBluishGreen      = BrickColor{id: 118, name: "LightBluishGreen", c: color.RGBA{R: uint8(183), G: uint8(215), B: uint8(213), A: uint8(0)}}
	BrYellowishGreen      = BrickColor{id: 119, name: "BrYellowishGreen", c: color.RGBA{R: uint8(164), G: uint8(189), B: uint8(70), A: uint8(0)}}
	LigYellowishGreen     = BrickColor{id: 120, name: "LigYellowishGreen", c: color.RGBA{R: uint8(217), G: uint8(228), B: uint8(167), A: uint8(0)}}
	MedYellowishOrange    = BrickColor{id: 121, name: "MedYellowishOrange", c: color.RGBA{R: uint8(231), G: uint8(172), B: uint8(88), A: uint8(0)}}
	BrReddishOrange       = BrickColor{id: 123, name: "BrReddishOrange", c: color.RGBA{R: uint8(211), G: uint8(111), B: uint8(76), A: uint8(0)}}
	BrightReddishViolet   = BrickColor{id: 124, name: "BrightReddishViolet", c: color.RGBA{R: uint8(146), G: uint8(57), B: uint8(120), A: uint8(0)}}
	LightOrange           = BrickColor{id: 125, name: "LightOrange", c: color.RGBA{R: uint8(234), G: uint8(184), B: uint8(145), A: uint8(0)}}
	Gold                  = BrickColor{id: 127, name: "Gold", c: color.RGBA{R: uint8(220), G: uint8(188), B: uint8(129), A: uint8(0)}}
	DarkNougat            = BrickColor{id: 128, name: "DarkNougat", c: color.RGBA{R: uint8(174), G: uint8(122), B: uint8(89), A: uint8(0)}}
	Silver                = BrickColor{id: 131, name: "Silver", c: color.RGBA{R: uint8(156), G: uint8(163), B: uint8(168), A: uint8(0)}}
	SandBlue              = BrickColor{id: 135, name: "SandBlue", c: color.RGBA{R: uint8(116), G: uint8(134), B: uint8(156), A: uint8(0)}}
	SandViolet            = BrickColor{id: 136, name: "SandViolet", c: color.RGBA{R: uint8(135), G: uint8(124), B: uint8(144), A: uint8(0)}}
	MediumOrange          = BrickColor{id: 137, name: "MediumOrange", c: color.RGBA{R: uint8(224), G: uint8(152), B: uint8(100), A: uint8(0)}}
	SandYellow            = BrickColor{id: 138, name: "SandYellow", c: color.RGBA{R: uint8(149), G: uint8(138), B: uint8(115), A: uint8(0)}}
	EarthBlue             = BrickColor{id: 140, name: "EarthBlue", c: color.RGBA{R: uint8(32), G: uint8(58), B: uint8(86), A: uint8(0)}}
	EarthGreen            = BrickColor{id: 141, name: "EarthGreen", c: color.RGBA{R: uint8(39), G: uint8(70), B: uint8(44), A: uint8(0)}}
	SandBlueMetallic      = BrickColor{id: 145, name: "SandBlueMetallic", c: color.RGBA{R: uint8(121), G: uint8(136), B: uint8(161), A: uint8(0)}}
	SandVioletMetallic    = BrickColor{id: 146, name: "SandVioletMetallic", c: color.RGBA{R: uint8(149), G: uint8(142), B: uint8(163), A: uint8(0)}}
	SandYellowMetallic    = BrickColor{id: 147, name: "SandYellowMetallic", c: color.RGBA{R: uint8(147), G: uint8(135), B: uint8(103), A: uint8(0)}}
	DarkGreyMetallic      = BrickColor{id: 148, name: "DarkGreyMetallic", c: color.RGBA{R: uint8(87), G: uint8(88), B: uint8(87), A: uint8(0)}}
	BlackMetallic         = BrickColor{id: 149, name: "BlackMetallic", c: color.RGBA{R: uint8(22), G: uint8(29), B: uint8(50), A: uint8(0)}}
	LightGreyMetallic     = BrickColor{id: 150, name: "LightGreyMetallic", c: color.RGBA{R: uint8(171), G: uint8(173), B: uint8(172), A: uint8(0)}}
	Sand                  = BrickColor{id: 151, name: "Sand", c: color.RGBA{R: uint8(10), G: uint8(120), B: uint8(144), A: uint8(0)}}
	SandRed               = BrickColor{id: 153, name: "SandRed", c: color.RGBA{R: uint8(149), G: uint8(121), B: uint8(118), A: uint8(0)}}
	DarkRed               = BrickColor{id: 154, name: "DarkRed", c: color.RGBA{R: uint8(123), G: uint8(46), B: uint8(47), A: uint8(0)}}
	Gun                   = BrickColor{id: 168, name: "Gun", c: color.RGBA{R: uint8(15), G: uint8(117), B: uint8(108), A: uint8(0)}}
	Curry                 = BrickColor{id: 180, name: "Curry", c: color.RGBA{R: uint8(215), G: uint8(169), B: uint8(75), A: uint8(0)}}
	LemonMetalic          = BrickColor{id: 200, name: "LemonMetalic", c: color.RGBA{R: uint8(130), G: uint8(138), B: uint8(93), A: uint8(0)}}
	FireYellow            = BrickColor{id: 190, name: "FireYellow", c: color.RGBA{R: uint8(249), G: uint8(214), B: uint8(46), A: uint8(0)}}
	FlameYellowishOrange  = BrickColor{id: 191, name: "FlameYellowishOrange", c: color.RGBA{R: uint8(232), G: uint8(171), B: uint8(45), A: uint8(0)}}
	ReddishBrown          = BrickColor{id: 192, name: "ReddishBrown", c: color.RGBA{R: uint8(105), G: uint8(64), B: uint8(39), A: uint8(0)}}
	FlameReddishOrange    = BrickColor{id: 193, name: "FlameReddishOrange", c: color.RGBA{R: uint8(207), G: uint8(96), B: uint8(36), A: uint8(0)}}
	MediumStoneGrey       = BrickColor{id: 194, name: "MediumStoneGrey", c: color.RGBA{R: uint8(163), G: uint8(162), B: uint8(164), A: uint8(0)}}
	RoyalBlue             = BrickColor{id: 195, name: "RoyalBlue", c: color.RGBA{R: uint8(70), G: uint8(103), B: uint8(164), A: uint8(0)}}
	DarkRoyalBlue         = BrickColor{id: 196, name: "DarkRoyalBlue", c: color.RGBA{R: uint8(35), G: uint8(71), B: uint8(139), A: uint8(0)}}
	BrightReddishLilac    = BrickColor{id: 198, name: "BrightReddishLilac", c: color.RGBA{R: uint8(142), G: uint8(66), B: uint8(133), A: uint8(0)}}
	DarkStoneGrey         = BrickColor{id: 199, name: "DarkStoneGrey", c: color.RGBA{R: uint8(99), G: uint8(95), B: uint8(97), A: uint8(0)}}
	LightStoneGrey        = BrickColor{id: 208, name: "LightStoneGrey", c: color.RGBA{R: uint8(229), G: uint8(228), B: uint8(222), A: uint8(0)}}
	DarkCurry             = BrickColor{id: 209, name: "DarkCurry", c: color.RGBA{R: uint8(176), G: uint8(142), B: uint8(68), A: uint8(0)}}
	FadedGreen            = BrickColor{id: 210, name: "FadedGreen", c: color.RGBA{R: uint8(112), G: uint8(149), B: uint8(120), A: uint8(0)}}
	Turquoise             = BrickColor{id: 211, name: "Turquoise", c: color.RGBA{R: uint8(121), G: uint8(181), B: uint8(181), A: uint8(0)}}
	LightRoyalBlue        = BrickColor{id: 212, name: "LightRoyalBlue", c: color.RGBA{R: uint8(159), G: uint8(195), B: uint8(233), A: uint8(0)}}
	MediumRoyalBlue       = BrickColor{id: 213, name: "MediumRoyalBlue", c: color.RGBA{R: uint8(108), G: uint8(129), B: uint8(183), A: uint8(0)}}
	Rust                  = BrickColor{id: 216, name: "Rust", c: color.RGBA{R: uint8(143), G: uint8(76), B: uint8(42), A: uint8(0)}}
	Brown                 = BrickColor{id: 217, name: "Brown", c: color.RGBA{R: uint8(124), G: uint8(92), B: uint8(69), A: uint8(0)}}
	ReddishLilac          = BrickColor{id: 218, name: "ReddishLilac", c: color.RGBA{R: uint8(150), G: uint8(112), B: uint8(159), A: uint8(0)}}
	Lilac                 = BrickColor{id: 219, name: "Lilac", c: color.RGBA{R: uint8(107), G: uint8(98), B: uint8(155), A: uint8(0)}}
	LightLilac            = BrickColor{id: 220, name: "LightLilac", c: color.RGBA{R: uint8(167), G: uint8(169), B: uint8(206), A: uint8(0)}}
	BrightPurple          = BrickColor{id: 221, name: "BrightPurple", c: color.RGBA{R: uint8(205), G: uint8(98), B: uint8(152), A: uint8(0)}}
	LightPurple           = BrickColor{id: 222, name: "LightPurple", c: color.RGBA{R: uint8(228), G: uint8(173), B: uint8(200), A: uint8(0)}}
	LightPink             = BrickColor{id: 223, name: "LightPink", c: color.RGBA{R: uint8(220), G: uint8(144), B: uint8(149), A: uint8(0)}}
	LightBrickYellow      = BrickColor{id: 224, name: "LightBrickYellow", c: color.RGBA{R: uint8(240), G: uint8(213), B: uint8(160), A: uint8(0)}}
	WarmYellowishOrange   = BrickColor{id: 225, name: "WarmYellowishOrange", c: color.RGBA{R: uint8(235), G: uint8(184), B: uint8(127), A: uint8(0)}}
	CoolYellow            = BrickColor{id: 226, name: "CoolYellow", c: color.RGBA{R: uint8(253), G: uint8(234), B: uint8(140), A: uint8(0)}}
	DoveBlue              = BrickColor{id: 232, name: "DoveBlue", c: color.RGBA{R: uint8(125), G: uint8(187), B: uint8(221), A: uint8(0)}}
	MediumLilac           = BrickColor{id: 268, name: "MediumLilac", c: color.RGBA{R: uint8(52), G: uint8(43), B: uint8(117), A: uint8(0)}}
	Transparent           = BrickColor{id: 40, name: "Transparent", c: color.RGBA{R: uint8(236), G: uint8(236), B: uint8(236), A: uint8(0)}}
	TrRed                 = BrickColor{id: 41, name: "TrRed", c: color.RGBA{R: uint8(205), G: uint8(84), B: uint8(75), A: uint8(0)}}
	TrLgBlue              = BrickColor{id: 42, name: "TrLgBlue", c: color.RGBA{R: uint8(193), G: uint8(223), B: uint8(240), A: uint8(0)}}
	TrBlue                = BrickColor{id: 43, name: "TrBlue", c: color.RGBA{R: uint8(123), G: uint8(182), B: uint8(232), A: uint8(0)}}
	TrYellow              = BrickColor{id: 44, name: "TrYellow", c: color.RGBA{R: uint8(247), G: uint8(241), B: uint8(141), A: uint8(0)}}
	TrFluReddishOrange    = BrickColor{id: 47, name: "TrFluReddishOrange", c: color.RGBA{R: uint8(217), G: uint8(133), B: uint8(108), A: uint8(0)}}
	TrGreen               = BrickColor{id: 48, name: "TrGreen", c: color.RGBA{R: uint8(132), G: uint8(182), B: uint8(141), A: uint8(0)}}
	TrFluGreen            = BrickColor{id: 49, name: "TrFluGreen", c: color.RGBA{R: uint8(248), G: uint8(241), B: uint8(132), A: uint8(0)}}
	PhosphWhite           = BrickColor{id: 50, name: "PhosphWhite", c: color.RGBA{R: uint8(236), G: uint8(232), B: uint8(222), A: uint8(0)}}
	TrBrown               = BrickColor{id: 111, name: "TrBrown", c: color.RGBA{R: uint8(191), G: uint8(183), B: uint8(177), A: uint8(0)}}
	TrMediReddishViolet   = BrickColor{id: 113, name: "TrMediReddishViolet", c: color.RGBA{R: uint8(228), G: uint8(173), B: uint8(200), A: uint8(0)}}
	TrBrightBluishViolet  = BrickColor{id: 126, name: "TrBrightBluishViolet", c: color.RGBA{R: uint8(165), G: uint8(165), B: uint8(203), A: uint8(0)}}
	NeonOrange            = BrickColor{id: 133, name: "NeonOrange", c: color.RGBA{R: uint8(213), G: uint8(115), B: uint8(61), A: uint8(0)}}
	NeonGreen             = BrickColor{id: 134, name: "NeonGreen", c: color.RGBA{R: uint8(216), G: uint8(221), B: uint8(86), A: uint8(0)}}
	TrFluBlue             = BrickColor{id: 143, name: "TrFluBlue", c: color.RGBA{R: uint8(207), G: uint8(226), B: uint8(247), A: uint8(0)}}
	TrFluYellow           = BrickColor{id: 157, name: "TrFluYellow", c: color.RGBA{R: uint8(255), G: uint8(246), B: uint8(123), A: uint8(0)}}
	TrFluRed              = BrickColor{id: 158, name: "TrFluRed", c: color.RGBA{R: uint8(225), G: uint8(164), B: uint8(194), A: uint8(0)}}
	RedFlipFlop           = BrickColor{id: 176, name: "RedFlipFlop", c: color.RGBA{R: uint8(151), G: uint8(105), B: uint8(91), A: uint8(0)}}
	YellowFlipFlop        = BrickColor{id: 178, name: "YellowFlipFlop", c: color.RGBA{R: uint8(180), G: uint8(132), B: uint8(85), A: uint8(0)}}
	SilverFlipFlop        = BrickColor{id: 179, name: "SilverFlipFlop", c: color.RGBA{R: uint8(137), G: uint8(135), B: uint8(136), A: uint8(0)}}

	FullPalette = color.Palette([]color.Color{
		White,
		Grey,
		LightYellow,
		BrickYellow,
		LightGreen,
		LightReddishViolet,
		LightOrangeBrown,
		Nougat,
		BrightRed,
		MedReddishViolet,
		BrightBlue,
		BrightYellow,
		EarthOrange,
		Black,
		DarkGrey,
		DarkGreen,
		MediumGreen,
		LightYellowishOrange,
		BrightGreen,
		DarkOrange,
		LightBluishViolet,
		LightBlue,
		LightRed,
		MediumRed,
		MediumBlue,
		LightGrey,
		BrightViolet,
		BrightYellowishOrange,
		BrightOrange,
		BrightBluishGreen,
		EarthYellow,
		BrightBluishViolet,
		MediumBluishViolet,
		MedYellowishGreen,
		MedBluishGreen,
		LightBluishGreen,
		BrYellowishGreen,
		LigYellowishGreen,
		MedYellowishOrange,
		BrReddishOrange,
		BrightReddishViolet,
		LightOrange,
		Gold,
		DarkNougat,
		Silver,
		SandBlue,
		SandViolet,
		MediumOrange,
		SandYellow,
		EarthBlue,
		EarthGreen,
		SandBlueMetallic,
		SandVioletMetallic,
		SandYellowMetallic,
		DarkGreyMetallic,
		BlackMetallic,
		LightGreyMetallic,
		Sand,
		SandRed,
		DarkRed,
		Gun,
		Curry,
		LemonMetalic,
		FireYellow,
		FlameYellowishOrange,
		ReddishBrown,
		FlameReddishOrange,
		MediumStoneGrey,
		RoyalBlue,
		DarkRoyalBlue,
		BrightReddishLilac,
		DarkStoneGrey,
		LightStoneGrey,
		DarkCurry,
		FadedGreen,
		Turquoise,
		LightRoyalBlue,
		MediumRoyalBlue,
		Rust,
		Brown,
		ReddishLilac,
		Lilac,
		LightLilac,
		BrightPurple,
		LightPurple,
		LightPink,
		LightBrickYellow,
		WarmYellowishOrange,
		CoolYellow,
		DoveBlue,
		MediumLilac,
		Transparent,
		TrRed,
		TrLgBlue,
		TrBlue,
		TrYellow,
		TrFluReddishOrange,
		TrGreen,
		TrFluGreen,
		PhosphWhite,
		TrBrown,
		TrMediReddishViolet,
		TrBrightBluishViolet,
		NeonOrange,
		NeonGreen,
		TrFluBlue,
		TrFluYellow,
		TrFluRed,
		RedFlipFlop,
		YellowFlipFlop,
		SilverFlipFlop,
	})
	LimitedPalette = color.Palette([]color.Color{
		White,
		Grey,
		Black,
		BrightRed,
		BrightBlue,
		BrightYellow,
		DarkGrey,
	})
	GrayPlusPalette = color.Palette([]color.Color{
		White,
		Grey,
		Black,
		DarkGrey,
		LightGrey,
		DarkStoneGrey,
		EarthGreen,
		DarkStoneGrey,
		LightStoneGrey,
		BrightRed,
		BrightBlue,
		BrightYellow,
	})
)
