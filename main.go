package main

import (
	"CytusSongsSelector/Controllers"
	"CytusSongsSelector/Models"
	"CytusSongsSelector/Songs"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/song")
	{
		v1.POST("/", Controllers.AddSongController)
		v1.GET("/:difficulty", Controllers.GetSongsByDifficultyController)
	}

	router.Run()
}

func randomSong(allSongs []Models.Song) Models.Song {
	rand.Seed(time.Now().UnixNano())
	rand.Intn(len(allSongs))
	return allSongs[rand.Intn(len(allSongs))]
}

func getSongByDifficulty(difficulty int) []Models.Song {
	allSongs := getAllSong()
	var targetSongs []Models.Song
	for i := range allSongs {
		if allSongs[i].Difficulty == difficulty {
			targetSongs = append(targetSongs, allSongs[i])
		}
	}
	return targetSongs
}

func printSong(song Models.Song) {
	fmt.Println()
	fmt.Printf("章節: %s\n", song.Chapter)
	fmt.Printf("順序: %d\n", song.Sequence)
	fmt.Printf("曲名: %s\n", song.Name)
	fmt.Printf("難易度: %d\n", song.Difficulty)
	fmt.Println()
}

func getAllSong() []Models.Song {
	var songList []Models.Song
	songList = Songs.Chapter1(songList)
	songList = Songs.Chapter2(songList)
	songList = Songs.Chapter3(songList)
	songList = Songs.Chapter4(songList)
	songList = Songs.Chapter5(songList)
	songList = Songs.Chapter6(songList)
	songList = Songs.Chapter7(songList)
	songList = Songs.Chapter8(songList)
	songList = Songs.Chapter9(songList)
	songList = Songs.Chapter10(songList)
	songList = Songs.ChapterS(songList)
	songList = Songs.ChapterK(songList)
	songList = Songs.ChapterM(songList)
	return songList
}
