package main

import (
	"fmt"
	"math/rand"
	"time"
)

//type Chapter struct {
//	Code string
//	Name string
//}

type Song struct {
	Chapter    string
	Sequence   int
	Name       string
	Difficulty int
}

func main() {
	//song := allRandom()
	//printSong(song)
	songList := getSongByDifficulty(9)
	printSong(randomSong(songList))
}

func randomSong(allSongs []Song) Song {
	rand.Seed(time.Now().UnixNano())
	rand.Intn(len(allSongs))
	return allSongs[rand.Intn(len(allSongs))]
}

func getSongByDifficulty(difficulty int) []Song {
	allSongs := getAllSong()
	var targetSongs []Song
	for i := range allSongs {
		if allSongs[i].Difficulty == difficulty {
			targetSongs = append(targetSongs, allSongs[i])
		}
	}
	return targetSongs
}

func printSong(song Song) {
	fmt.Println()
	fmt.Printf("章節: %s\n", song.Chapter)
	fmt.Printf("順序: %d\n", song.Sequence)
	fmt.Printf("曲名: %s\n", song.Name)
	fmt.Printf("難易度: %d\n", song.Difficulty)
	fmt.Println()
}

func getAllSong() []Song {
	var songList []Song
	songList = chapter1(songList)
	songList = chapter2(songList)
	songList = chapter3(songList)
	songList = chapter4(songList)
	songList = chapter5(songList)
	songList = chapter6(songList)
	songList = chapter7(songList)
	songList = chapter8(songList)
	songList = chapter9(songList)
	songList = chapter10(songList)
	//songList = chapter0(songList)
	songList = chapterS(songList)
	songList = chapterK(songList)
	songList = chapterM(songList)
	return songList
}

// 10 Freedom D↓ve
// 9 L2 ver.B
// 8 First Gate OVERDRIVE
// 7 Axion
// 6 Bloody purity
// 5 Colorful Skies
// 4 Holy Knight

// 難易度 6.4
func chapter1(songList []Song) []Song {
	const chapter = "1"
	songList = append(songList, Song{chapter, 1, "Light up my LOVE", 7})
	songList = append(songList, Song{chapter, 2, "Ververg", 7})
	songList = append(songList, Song{chapter, 2, "(5)Ververg II (畫面變暗時點擊火把)", 9})
	songList = append(songList, Song{chapter, 3, "Chemical Star", 6})
	songList = append(songList, Song{chapter, 4, "Visions", 7})
	songList = append(songList, Song{chapter, 5, "Les Parfums de L'Amour", 6})
	songList = append(songList, Song{chapter, 6, "Retrospective", 5})
	songList = append(songList, Song{chapter, 7, "The Silence", 7})
	songList = append(songList, Song{chapter, 8, "D R G", 7})
	songList = append(songList, Song{chapter, 9, "Secret Garden", 7})
	songList = append(songList, Song{chapter, 10, "Hot Air Balloon", 5})
	songList = append(songList, Song{chapter, 0, "OPERATORS", 4})
	return songList
}

// 難易度 7.8
func chapter2(songList []Song) []Song {
	const chapter = "2"
	songList = append(songList, Song{chapter, 1, "Iris", 7})
	songList = append(songList, Song{chapter, 2, "Sanctity", 7})
	songList = append(songList, Song{chapter, 3, "Sacred", 7})
	songList = append(songList, Song{chapter, 4, "(5)Green Eyes", 8})
	songList = append(songList, Song{chapter, 5, "Nocturnal Type", 7})
	songList = append(songList, Song{chapter, 6, "(4)Precipitation", 8})
	songList = append(songList, Song{chapter, 6, "(5)Precipitation II (雙指下滑)", 9})
	songList = append(songList, Song{chapter, 7, "(4)Hard Landing", 8})
	songList = append(songList, Song{chapter, 8, "(7)Entrance", 9})
	songList = append(songList, Song{chapter, 8, "(6)Precipitation at the Entrance I (將左側翅膀向右拉至兩翅膀根重合)", 9})
	songList = append(songList, Song{chapter, 8, "(7)Precipitation at the Entrance II (將右側翅膀向左拉至兩翅膀根重合)", 9})
	songList = append(songList, Song{chapter, 0, "DISASTER", 6})
	return songList
}

// 難易度 7.2
func chapter3(songList []Song) []Song {
	const chapter = "3"
	songList = append(songList, Song{chapter, 1, "(3)The Riddle Story", 8})
	songList = append(songList, Song{chapter, 2, "Libera Me", 7})
	songList = append(songList, Song{chapter, 3, "COSMO", 7})
	songList = append(songList, Song{chapter, 4, "Prismatic Lollipops", 7})
	songList = append(songList, Song{chapter, 5, "otome", 7})
	songList = append(songList, Song{chapter, 6, "Spectrum", 5})
	songList = append(songList, Song{chapter, 7, "(6)Halcyon", 9})
	songList = append(songList, Song{chapter, 8, "The black Case", 7})
	songList = append(songList, Song{chapter, 9, "Saika", 7})
	songList = append(songList, Song{chapter, 9, "Saika II (點擊「彩」「華」兩字，使兩字背景從紅色變成藍色)", 7})
	songList = append(songList, Song{chapter, 0, "(5)CYTUS", 8})
	return songList
}

// 難易度 7.4
func chapter4(songList []Song) []Song {
	const chapter = "4"
	songList = append(songList, Song{chapter, 1, "(4)Evil force", 8})
	songList = append(songList, Song{chapter, 2, "New world", 6})
	songList = append(songList, Song{chapter, 3, "Landscape", 7})
	songList = append(songList, Song{chapter, 4, "Future world", 7})
	songList = append(songList, Song{chapter, 5, "(5)Parousia", 9})
	songList = append(songList, Song{chapter, 6, "Skuld", 6})
	songList = append(songList, Song{chapter, 7, "(5)Darkness", 8})
	songList = append(songList, Song{chapter, 8, "Beyond", 6})
	songList = append(songList, Song{chapter, 9, "(5)AREA 184", 9})
	songList = append(songList, Song{chapter, 10, "(5)Sweetness and Love", 8})
	songList = append(songList, Song{chapter, 0, "THE SILENCE", 7})
	return songList
}

// 難易度 7.8
func chapter5(songList []Song) []Song {
	const chapter = "5"
	songList = append(songList, Song{chapter, 1, "(4)Holy Knight", 8})
	songList = append(songList, Song{chapter, 2, "Dino", 7})
	songList = append(songList, Song{chapter, 3, "Majestic Phoenix", 7})
	songList = append(songList, Song{chapter, 4, "Sleepless Jasmine", 6})
	songList = append(songList, Song{chapter, 5, "Chocological", 6})
	songList = append(songList, Song{chapter, 6, "(5)Recollections", 9})
	songList = append(songList, Song{chapter, 7, "(4)Total Sphere", 8})
	songList = append(songList, Song{chapter, 8, "(4)Just a trip", 8})
	songList = append(songList, Song{chapter, 9, "(4)Zauberkugel", 9})
	songList = append(songList, Song{chapter, 10, "(6)Biotonic", 9})
	songList = append(songList, Song{chapter, 0, "(7)VANESSA", 9})
	return songList
}

// 難易度 7.6
func chapter6(songList []Song) []Song {
	const chapter = "6"
	songList = append(songList, Song{chapter, 1, "Dragon Warrior", 7})
	songList = append(songList, Song{chapter, 2, "Selfish Gene", 6})
	songList = append(songList, Song{chapter, 3, "Realize", 7})
	songList = append(songList, Song{chapter, 4, "(5)Colorful Skies", 9})
	songList = append(songList, Song{chapter, 5, "It's a wonderful world", 6})
	songList = append(songList, Song{chapter, 6, "(6)Bloody purity", 9})
	songList = append(songList, Song{chapter, 7, "(4)Logical steps", 8})
	songList = append(songList, Song{chapter, 8, "(4)Niflheimr", 8})
	songList = append(songList, Song{chapter, 9, "(4)OLD GOLD", 8})
	songList = append(songList, Song{chapter, 10, "(6)The blocks we loved", 9})
	songList = append(songList, Song{chapter, 0, "THE LOST", 7})
	return songList
}

// 難易度 8.1
func chapter7(songList []Song) []Song {
	const chapter = "7"
	songList = append(songList, Song{chapter, 1, "(6)Black Lair", 9})
	songList = append(songList, Song{chapter, 2, "(4)The Last Illusion", 8})
	songList = append(songList, Song{chapter, 3, "(5)Galaxy Collapse", 8})
	songList = append(songList, Song{chapter, 4, "(7)L", 9})
	songList = append(songList, Song{chapter, 4, "(8)L2 ver.A (Loneliness) (按住L至變紅)", 9})
	songList = append(songList, Song{chapter, 4, "(9)L2 ver.B (Liberation) (按住L至變藍)", 9})
	songList = append(songList, Song{chapter, 5, "(5)Gate of Expectancy", 8})
	songList = append(songList, Song{chapter, 6, "(6)Rainbow Night Sky Highway", 9})
	songList = append(songList, Song{chapter, 7, "Quantum Labyrinth", 7})
	songList = append(songList, Song{chapter, 8, "Musik", 7})
	songList = append(songList, Song{chapter, 9, "Hercule", 7})
	songList = append(songList, Song{chapter, 10, "(4)Aquatic Poseidon", 8})
	songList = append(songList, Song{chapter, 0, "LOOM", 7})
	return songList
}

// 難易度 8.0
func chapter8(songList []Song) []Song {
	const chapter = "8"
	songList = append(songList, Song{chapter, 1, "(6)Masquerade", 9})
	songList = append(songList, Song{chapter, 2, "Her Sword", 7})
	songList = append(songList, Song{chapter, 3, "Morpho", 8})
	songList = append(songList, Song{chapter, 4, "Silt", 8})
	songList = append(songList, Song{chapter, 4, "(7)Slit ooo (在螢幕中間由左上劃向右下)", 9})
	songList = append(songList, Song{chapter, 4, "(8)Silt iii (在螢幕中間由右上劃向左下)", 9})
	songList = append(songList, Song{chapter, 5, "(5)Laplace", 9})
	songList = append(songList, Song{chapter, 6, "(5)Q", 9})
	songList = append(songList, Song{chapter, 7, "Scherzo", 7})
	songList = append(songList, Song{chapter, 8, "(7)AXION", 9})
	songList = append(songList, Song{chapter, 9, "Code 03", 6})
	songList = append(songList, Song{chapter, 10, "Reverence", 8})
	songList = append(songList, Song{chapter, 0, "ANOTHER ME", 6})
	return songList
}

// 難易度 7.9
func chapter9(songList []Song) []Song {
	const chapter = "9"
	songList = append(songList, Song{chapter, 1, "(7)Oriens", 9})
	songList = append(songList, Song{chapter, 2, "Hey wonder", 6})
	songList = append(songList, Song{chapter, 3, "Brionac", 7})
	songList = append(songList, Song{chapter, 4, "First Gate", 8})
	songList = append(songList, Song{chapter, 4, "(8)First Gate OVERDRIVE (點擊標題，標題上彈出OVERDRIVE印章)", 9})
	songList = append(songList, Song{chapter, 5, "Qualia", 7})
	songList = append(songList, Song{chapter, 6, "East West Whobble", 8})
	songList = append(songList, Song{chapter, 7, "Warlords of Atlantis", 8})
	songList = append(songList, Song{chapter, 8, "To Further Dream", 8})
	songList = append(songList, Song{chapter, 9, "COMA", 8})
	songList = append(songList, Song{chapter, 10, "(9)CODE NAME: ZERO", 9})
	songList = append(songList, Song{chapter, 0, "BURIED", 8})
	return songList
}

// 難易度 7.3
func chapter10(songList []Song) []Song {
	const chapter = "10"
	songList = append(songList, Song{chapter, 1, "(9)Freedom Dive", 9})
	songList = append(songList, Song{chapter, 1, "(10)Freedom D↓ve (手指下滑，曲繪標題變為「Freedom D↓ve」)", 9})
	songList = append(songList, Song{chapter, 2, "(7)Halloween Party", 9})
	songList = append(songList, Song{chapter, 3, "YURERO", 7})
	songList = append(songList, Song{chapter, 4, "Twenty One", 7})
	songList = append(songList, Song{chapter, 5, "Solar Wind", 8})
	songList = append(songList, Song{chapter, 6, "Red Eyes", 7})
	songList = append(songList, Song{chapter, 7, "Finite Circuit", 7})
	songList = append(songList, Song{chapter, 8, "Set Free", 6})
	songList = append(songList, Song{chapter, 9, "(5)Do Not Wake", 9})
	songList = append(songList, Song{chapter, 10, "(6)GATORIX", 9})
	songList = append(songList, Song{chapter, 0, "THE NEW WORLD", 1})
	return songList
}

// 難易度 7.7
//func chapter0(songList []Song) []Song {
//	const chapter = "0"
//	songList = append(songList, Song{chapter, 1, "PROCESS", 8})
//	songList = append(songList, Song{chapter, 2, "Endless Journey", 7})
//	songList = append(songList, Song{chapter, 3, "Shoot out", 7})
//	songList = append(songList, Song{chapter, 4, "LNS OP", 6})
//	songList = append(songList, Song{chapter, 5, "Blue Eyes", 9})
//	songList = append(songList, Song{chapter, 6, "Diskord", 7})
//	songList = append(songList, Song{chapter, 7, "Infernus", 8})
//	songList = append(songList, Song{chapter, 8, "Megaera", 8})
//	songList = append(songList, Song{chapter, 9, "Violet", 9})
//	songList = append(songList, Song{chapter, 10, "¡Azucar!", 8})
//	return songList
//}

// 難易度 7.6
func chapterS(songList []Song) []Song {
	const chapter = "S"
	songList = append(songList, Song{chapter, 1, "LVBNR5 SCHWARZ", 7})
	songList = append(songList, Song{chapter, 2, "Vivere La Vita", 7})
	songList = append(songList, Song{chapter, 3, "Rain of Fire", 7})
	songList = append(songList, Song{chapter, 4, "Molto Allergo", 7})
	songList = append(songList, Song{chapter, 5, "(7)Revoluxionist", 9})
	songList = append(songList, Song{chapter, 6, "LVBNR5 Weiß", 6})
	songList = append(songList, Song{chapter, 7, "Chaotic Drive", 7})
	songList = append(songList, Song{chapter, 8, "Outsider", 8})
	songList = append(songList, Song{chapter, 9, "(7)Requiem", 9})
	songList = append(songList, Song{chapter, 10, "(8)The Purified", 9})
	return songList
}

// 難易度 7.4
func chapterK(songList []Song) []Song {
	const chapter = "K"
	songList = append(songList, Song{chapter, 1, "The Way We Were", 6})
	songList = append(songList, Song{chapter, 2, "The Sanctuary", 8})
	songList = append(songList, Song{chapter, 3, "(6)The Red Coronation", 9})
	songList = append(songList, Song{chapter, 4, "Forbidden Codex", 7})
	songList = append(songList, Song{chapter, 5, "Knight of Firmament", 7})
	songList = append(songList, Song{chapter, 6, "Lord of Crimson Rose", 8})
	songList = append(songList, Song{chapter, 7, "Predawn", 7})
	songList = append(songList, Song{chapter, 8, "The Fallen Bloom", 7})
	songList = append(songList, Song{chapter, 9, "Where You Are Not", 8})
	songList = append(songList, Song{chapter, 10, "Music. The eternity of us", 7})
	return songList
}

// 難易度 8.8
func chapterM(songList []Song) []Song {
	const chapter = "M"
	songList = append(songList, Song{chapter, 1, "(7)Stardust Sphere", 9})
	songList = append(songList, Song{chapter, 2, "(8)The Ricochet", 9})
	songList = append(songList, Song{chapter, 2, "(6)The Long Years (點擊機器人的耳朵)", 9})
	songList = append(songList, Song{chapter, 3, "(6)The Sacred Story", 9})
	songList = append(songList, Song{chapter, 4, "(8)AREA184 -Platnium Mix-", 9})
	songList = append(songList, Song{chapter, 5, "(5)Gardenia", 9})
	songList = append(songList, Song{chapter, 6, "(9)Sweetness Overload!!!", 9})
	songList = append(songList, Song{chapter, 7, "(6)Les Parfums de Celebrez", 9})
	songList = append(songList, Song{chapter, 8, "(5)Afterglow", 9})
	songList = append(songList, Song{chapter, 9, "(7)GENESYS", 9})
	songList = append(songList, Song{chapter, 10, "(8)Storia", 9})
	songList = append(songList, Song{chapter, 0, "THE BEGINNING", 7})
	return songList
}
