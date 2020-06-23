package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"reality.rehab/hgol/board"
)

func main() {
	//testString := "The ancient pond. A frog leaps in. Water's sound."
	//testString := "fuck the police. this is the life. i do not know."

	poemFile, err := os.Open("./haiku.json")
	if err != nil {
		fmt.Print("uhoh")
	}
	defer poemFile.Close()

	var poems []string
	byteValue, _ := ioutil.ReadAll(poemFile)
	json.Unmarshal(byteValue, &poems)

	rand.Seed(time.Now().Unix())
	poem := poems[rand.Intn(len(poems))]

	world := board.New(poem)

	generations := len(world.BinaryString)

	fmt.Print(generations)

	fmt.Print("\nInitial world:\n")

	world.String()

	for i := 0; i < generations; i++ {
		world.Progress()
	}

	fmt.Print("Final world:\n")
	world.String()

	fmt.Print(poem)
	fmt.Print('\n')

	world.Draw("basho.png")

	post(poem)
}

func post(poem string) {
	api := anaconda.NewTwitterApiWithCredentials()

	data, err := ioutil.ReadFile("basho.png")

	if err != nil {
		fmt.Println(err)
	}

	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))

	if err != nil {
		fmt.Println(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))
	result, err := api.PostTweet(poem, v)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
