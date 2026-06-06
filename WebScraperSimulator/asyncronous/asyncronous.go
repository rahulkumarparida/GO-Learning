package asyncronous


import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)




var listOfWebsitesToBeScraped = []string{
	"https://github.com/",
	"https://google.com",
	"https://techheart.life/",
	"https://example.com/",
    "https://npr.org",               
    "https://duckduckgo.com",    
    "https://berkshirehathaway.com",   
    "https://news.ycombinator.com/",        
    "http://paulgraham.com",              
    "https://danluu.com",                  
    "http://textfiles.com",                
    "https://wiby.me",                   
    "http://cern.ch", 
	"https://cnn.com",               
    "https://wttr.in",                     
    "https://google.com",         
    "https://motherfuckingwebsite.com",   
    "http://gutenberg.org",                
    "http://catb.org",   
}





func RoutineScrapingBehaviour(){
	scrapeChan := make(chan []byte,1)
	wg := sync.WaitGroup{}

	for idx, val := range listOfWebsitesToBeScraped {
		fname := strconv.Itoa(idx)+".txt"		

		wg.Add(2)
		go ScrapeWebData(val , scrapeChan , &wg)
		go WriteDataFile(fname,scrapeChan, &wg)
	}
	wg.Wait()
	close(scrapeChan)	

}

type ScrapeData struct{
	sucess bool
	data []byte
	url string
}
func ScrapeWebData(url string , cp chan<- []byte , wg *sync.WaitGroup) ScrapeData{
	defer wg.Done()

	cmd  := exec.Command("curl" , "-f" ,url)
	res , err:= cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error While Scraping data from : ", url)
		return ScrapeData{false,[]byte("failed scraping data from the website"),url}
	}
	cp<- res
	return  ScrapeData{true,res,url}

}


func WriteDataFile( fname string , cp <-chan []byte , wg *sync.WaitGroup) bool{
	defer wg.Done()

	currDir , err := os.Getwd()
	data := string(<-cp)
	if err != nil {
		return  false
	}
	dirPath := filepath.Join(currDir,"scraped_data")

	os.Mkdir(dirPath , 0755)

	filename := filepath.Join(dirPath , fname)

	errwriterr := os.WriteFile(filename , []byte(data) , 0644)

	if errwriterr != nil {
		return false
	}
		
	return true
}


