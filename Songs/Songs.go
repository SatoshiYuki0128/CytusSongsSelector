package Songs

import "CytusSongsSelector/Models"

// 10 Freedom D↓ve
// 9 L2 ver.B
// 8 First Gate OVERDRIVE
// 7 Axion
// 6 Bloody purity
// 5 Colorful Skies
// 4 Holy Knight

// Chapter1 難易度 6.4
func Chapter1(songList []Models.Song) []Models.Song {
	const chapter = "1"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "Light up my LOVE", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Ververg", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "(5)Ververg II (畫面變暗時點擊火把)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Chemical Star", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Visions", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Les Parfums de L'Amour", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "Retrospective", Difficulty: 5})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "The Silence", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "D R G", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Secret Garden", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "Hot Air Balloon", Difficulty: 5})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "OPERATORS", Difficulty: 4})
	return songList
}

// Chapter2 難易度 7.8
func Chapter2(songList []Models.Song) []Models.Song {
	const chapter = "2"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "Iris", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Sanctity", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Sacred", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(5)Green Eyes", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Nocturnal Type", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(4)Precipitation", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(5)Precipitation II (雙指下滑)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(4)Hard Landing", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(7)Entrance", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(6)Precipitation at the Entrance I (將左側翅膀向右拉至兩翅膀根重合)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(7)Precipitation at the Entrance II (將右側翅膀向左拉至兩翅膀根重合)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "DISASTER", Difficulty: 6})
	return songList
}

// Chapter3 難易度 7.2
func Chapter3(songList []Models.Song) []Models.Song {
	const chapter = "3"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(3)The Riddle Story", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Libera Me", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "COSMO", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Prismatic Lollipops", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "otome", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "Spectrum", Difficulty: 5})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(6)Halcyon", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "The black Case", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Saika", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Saika II (點擊「彩」「華」兩字，使兩字背景從紅色變成藍色)", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "(5)CYTUS", Difficulty: 8})
	return songList
}

// Chapter4 難易度 7.4
func Chapter4(songList []Models.Song) []Models.Song {
	const chapter = "4"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(4)Evil force", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "New world", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Landscape", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Future world", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "(5)Parousia", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "Skuld", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(5)Darkness", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "Beyond", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(5)AREA 184", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(5)Sweetness and Love", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "THE SILENCE", Difficulty: 7})
	return songList
}

// Chapter5 難易度 7.8
func Chapter5(songList []Models.Song) []Models.Song {
	const chapter = "5"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(4)Holy Knight", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Dino", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Majestic Phoenix", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Sleepless Jasmine", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Chocological", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(5)Recollections", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(4)Total Sphere", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(4)Just a trip", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(4)Zauberkugel", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(6)Biotonic", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "(7)VANESSA", Difficulty: 9})
	return songList
}

// Chapter6 難易度 7.6
func Chapter6(songList []Models.Song) []Models.Song {
	const chapter = "6"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "Dragon Warrior", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Selfish Gene", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Realize", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(5)Colorful Skies", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "It's a wonderful world", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(6)Bloody purity", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(4)Logical steps", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(4)Niflheimr", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(4)OLD GOLD", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(6)The blocks we loved", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "THE LOST", Difficulty: 7})
	return songList
}

// Chapter7 難易度 8.1
func Chapter7(songList []Models.Song) []Models.Song {
	const chapter = "7"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(6)Black Lair", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "(4)The Last Illusion", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "(5)Galaxy Collapse", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(7)L", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(8)L2 ver.A (Loneliness) (按住L至變紅)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(9)L2 ver.B (Liberation) (按住L至變藍)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "(5)Gate of Expectancy", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(6)Rainbow Night Sky Highway", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Quantum Labyrinth", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "Musik", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Hercule", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(4)Aquatic Poseidon", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "LOOM", Difficulty: 7})
	return songList
}

// Chapter8 難易度 8.0
func Chapter8(songList []Models.Song) []Models.Song {
	const chapter = "8"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(6)Masquerade", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Her Sword", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Morpho", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Silt", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(7)Slit ooo (在螢幕中間由左上劃向右下)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(8)Silt iii (在螢幕中間由右上劃向左下)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "(5)Laplace", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(5)Q", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Scherzo", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(7)AXION", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Code 03", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "Reverence", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "ANOTHER ME", Difficulty: 6})
	return songList
}

// Chapter9 難易度 7.9
func Chapter9(songList []Models.Song) []Models.Song {
	const chapter = "9"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(7)Oriens", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Hey wonder", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Brionac", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "First Gate", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(8)First Gate OVERDRIVE (點擊標題，標題上彈出OVERDRIVE印章)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Qualia", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "East West Whobble", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Warlords of Atlantis", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "To Further Dream", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "COMA", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(9)CODE NAME: ZERO", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "BURIED", Difficulty: 8})
	return songList
}

// Chapter10 難易度 7.3
func Chapter10(songList []Models.Song) []Models.Song {
	const chapter = "10"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(9)Freedom Dive", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(10)Freedom D↓ve (手指下滑，曲繪標題變為「Freedom D↓ve」)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "(7)Halloween Party", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "YURERO", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Twenty One", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Solar Wind", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "Red Eyes", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Finite Circuit", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "Set Free", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(5)Do Not Wake", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(6)GATORIX", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "THE NEW WORLD", Difficulty: 1})
	return songList
}

// ChapterS 難易度 7.6
func ChapterS(songList []Models.Song) []Models.Song {
	const chapter = "S"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "LVBNR5 SCHWARZ", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "Vivere La Vita", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "Rain of Fire", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Molto Allergo", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "(7)Revoluxionist", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "LVBNR5 Weiß", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Chaotic Drive", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "Outsider", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(7)Requiem", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(8)The Purified", Difficulty: 9})
	return songList
}

// ChapterK 難易度 7.4
func ChapterK(songList []Models.Song) []Models.Song {
	const chapter = "K"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "The Way We Were", Difficulty: 6})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "The Sanctuary", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "(6)The Red Coronation", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "Forbidden Codex", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "Knight of Firmament", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "Lord of Crimson Rose", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "Predawn", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "The Fallen Bloom", Difficulty: 7})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "Where You Are Not", Difficulty: 8})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "Music. The eternity of us", Difficulty: 7})
	return songList
}

// ChapterM 難易度 8.8
func ChapterM(songList []Models.Song) []Models.Song {
	const chapter = "M"
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 1, Name: "(7)Stardust Sphere", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "(8)The Ricochet", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 2, Name: "(6)The Long Years (點擊機器人的耳朵)", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 3, Name: "(6)The Sacred Story", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 4, Name: "(8)AREA184 -Platnium Mix-", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 5, Name: "(5)Gardenia", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 6, Name: "(9)Sweetness Overload!!!", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 7, Name: "(6)Les Parfums de Celebrez", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 8, Name: "(5)Afterglow", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 9, Name: "(7)GENESYS", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Sequence: 10, Name: "(8)Storia", Difficulty: 9})
	songList = append(songList, Models.Song{Chapter: chapter, Name: "THE BEGINNING", Difficulty: 7})
	return songList
}

// 難易度 7.7
//func chapter0(songList []Models.Song) []Models.Song {
//	const chapter = "0"
//	songList = append(songList, Models.Song{chapter, 1, "PROCESS", 8})
//	songList = append(songList, Models.Song{chapter, 2, "Endless Journey", 7})
//	songList = append(songList, Models.Song{chapter, 3, "Shoot out", 7})
//	songList = append(songList, Models.Song{chapter, 4, "LNS OP", 6})
//	songList = append(songList, Models.Song{chapter, 5, "Blue Eyes", 9})
//	songList = append(songList, Models.Song{chapter, 6, "Diskord", 7})
//	songList = append(songList, Models.Song{chapter, 7, "Infernus", 8})
//	songList = append(songList, Models.Song{chapter, 8, "Megaera", 8})
//	songList = append(songList, Models.Song{chapter, 9, "Violet", 9})
//	songList = append(songList, Models.Song{chapter, 10, "¡Azucar!", 8})
//	return songList
//}
