package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"reality.rehab/hgol/board"
)

func main() {
	fmt.Printf("you can either type 'input' or 'random': ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	operation := strings.TrimRight(scanner.Text(), "\n")

	var poem string
	if operation == "input" {
		fmt.Printf("Enter some text: \n")
		scanner.Scan()

		re := regexp.MustCompile("\\\\n")
		text := re.ReplaceAllString(strings.TrimRight(scanner.Text(), "\n"), "\n")
		poem = doInputPoem(text)
	} else if operation == "random" {
		poem = doRandom()
	}
	//poem := "My Life\nCame like dew\nDisappears like dew\nAll of Naniwa\nIs dream after Dream"

	fmt.Print(poem)
	fmt.Printf("It is saved to basho.png. Do you want to post it to Twitter? ")
	scanner.Scan()
	choice := strings.TrimRight(scanner.Text(), "\n")

	if choice == "yes" || choice == "y" {
		fmt.Print("posting\n")
		post(poem)
	} else {
		fmt.Print("good bye\n")
	}
}

func doInputPoem(poem string) string {
	world := board.New(poem)
	generations := len(world.BinaryString)

	for i := 0; i < generations; i++ {
		world.Progress()
	}

	world.Draw("basho.png")

	return poem
}

func doRandom() string {
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

	return doInputPoem(poem)
}

func post(poem string) {
	secritsFile, err := os.Open("./secrits.txt")
	if err != nil {
		fmt.Print("uhoh")
	}
	defer secritsFile.Close()

	var cred []string
	scanner := bufio.NewScanner(secritsFile)
	for scanner.Scan() {
		cred = append(cred, scanner.Text())
	}

	// heh
	api := anaconda.NewTwitterApiWithCredentials(cred[0], cred[1], cred[2], cred[3])

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
