package main

import (
	"fmt"
	//"bufio"
	//"os"
)

type storyPage struct {
	text string
	nextPage *storyPage


}

func playStory(page *storyPage){
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It is a dark and stromy night",nil}
	page2 := storyPage{" you are all alone, need to find helmet",nil}
	page3 := storyPage{"there is a troll",nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)

}
