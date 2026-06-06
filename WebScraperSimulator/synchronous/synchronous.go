package synchronous

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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




func SynchronousBehaviourScraping(){
		for idx, val := range listOfWebsitesToBeScraped {
		
		scrapedData := ScrapeWebData(val)
		if scrapedData.sucess {
				fname := strconv.Itoa(idx)+".txt"
				wroteData := WriteDataFile(string(scrapedData.data),fname )
				if wroteData {
					fmt.Println("Sucessfully scraped:",val)
				}else{
					fmt.Println("Some error occured while executing the functions.")
				}
		}else{
			fmt.Println("Message: ",string(scrapedData.data))
		}
		

	}
}

type ScrapeData struct{
	sucess bool
	data []byte
	url string
}
func ScrapeWebData(url string) ScrapeData{
	cmd  := exec.Command("curl" , "-f" ,url)
	res , err:= cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error While Scraping data from : ", url)
		
		return ScrapeData{false,[]byte("failed scraping data from the website"),url}
	}
	return ScrapeData{true,res,url}
}


func WriteDataFile(data string , fname string ) bool{
	currDir , err := os.Getwd()

	if err != nil {
		return  false
	}
	dirPath := filepath.Join(currDir,"scraped_data")

	os.Mkdir(dirPath , 0755)

	filename := filepath.Join(dirPath , fname)

	errwriterr := os.WriteFile(filename , []byte(data) , 0644)

	if errwriterr != nil {
		return  false
	}
		
	return true
}


