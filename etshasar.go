package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func download(sii_decrypt_location string){
	// download SII_Decrypt.exe
	fmt.Println("Downloading SII_Decrypt.exe")
	// https://raw.githubusercontent.com/LastPlayerTR/ets2_damagereset/main/SII_Decrypt.exe
	url := "https://raw.githubusercontent.com/LastPlayerTR/ets2_damagereset/main/SII_Decrypt.exe"
	// get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("cannot download SII_Decrypt.exe: ", err)
		fmt.Scanln() // wait for Enter Key
		return
	}
	defer resp.Body.Close()
	// create the file
	out, err := os.Create(sii_decrypt_location)
	if err != nil {
		fmt.Println("cannot create SII_Decrypt.exe: ", err)
		fmt.Scanln() // wait for Enter Key
		return
	}
	defer out.Close()
	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("cannot write SII_Decrypt.exe: ", err)
		fmt.Scanln() // wait for Enter Key
		return
	}
	fmt.Println("SII_Decrypt.exe downloaded")
	fmt.Println("Please restart the app!")

}

func main() {



	fmt.Println("Starting...")
	

	ex, err := os.Executable()
	if err != nil {
		fmt.Println("os exe okunamadi: ", err)
		fmt.Scanln() // wait for Enter Key
		return
	}
	exPath := filepath.Dir(ex)
	fmt.Printf("Current directory: %s\n", exPath)

	sii_decrypt_location := filepath.Join(exPath , "SII_Decrypt.exe")

	file, err := os.Open(sii_decrypt_location)
	if err != nil {
		fmt.Println("SII Decrypt not found downloading... ", err)
		download(sii_decrypt_location)
		fmt.Scanln() // wait for Enter Key
		return
	}
	defer file.Close()

	// current directory

	fmt.Printf(filepath.Join(exPath , "game.sii"))

	game_sii_location := ""
	
	if(len(os.Args) < 2) {
		fmt.Println("argument not found checking self directory...")
		game_sii_location = filepath.Join(exPath , "game.sii")
		
	}else{
		game_sii_location = os.Args[1]

	}



	cmd := exec.Command( sii_decrypt_location , game_sii_location)


	_cmd := cmd.Run()
	if _cmd != nil {
		// if exit code is 1
		if(_cmd.Error() == "exit status 1") {
			fmt.Println("game.sii already decoded:")
		}else{
			fmt.Println("SII_Decrypt.exe error: ", _cmd)
			fmt.Scanln() // wait for Enter Key
			return
		}

	}


	//game.sii
	file_game , err_game := os.Open(game_sii_location)
	if err_game != nil {
		fmt.Println("game.sii not found: ", err_game)
		fmt.Scanln() // wait for Enter Key

		return
	}
	// game.sii should be lower than 50 mb
	file_game_stat, _ := file_game.Stat()
	if file_game_stat.Size() > 50000000 {
		fmt.Println("game.sii bigger than 50mb")
		fmt.Scanln() // wait for Enter Key

		return
	}

	
	// read all content of game.sii
	content, err_read := io.ReadAll(file_game)
	if err_read != nil {
		fmt.Println("cannot read game.sii: ", err_read)
		fmt.Scanln() // wait for Enter Key

		return
	}
	defer file_game.Close()
	// convert content to string
	content_str := string(content)

	
	// read every line of content_str
	lines := strings.Split(content_str, "\n")

	// find the line that contains engine_wear , transmission_wear, cabin_wear, engine_wear_unfixable, transmission_wear_unfixable, cabin_wear_unfixable,chassis_wear , chassis_wear_unfixable , wheels_wear_unfixable[

	// loop lines
	if(len(lines) < 1) {
		fmt.Println("game.sii is empty")
		fmt.Scanln() // wait for Enter Key
		return
	}


	for i, line := range lines {
		stringd := strings.Split(line, ":")
		if len(stringd) < 2 {
			continue
		}
		if len(stringd) > 2 {
			continue
		}
		if(stringd[0] == " engine_wear") {
			line = " engine_wear: 0"
		}
		if( stringd[0] == " transmission_wear") {
			line = " transmission_wear: 0"
			//
		}
		if( stringd[0] == " cabin_wear") {
			line = " cabin_wear: 0"
			//
		}
		if( stringd[0] == " engine_wear_unfixable") {
			line = " engine_wear_unfixable: 0"

			//
		}
		if( stringd[0] == " transmission_wear_unfixable") {
			line = " transmission_wear_unfixable: 0"
			//
		}
		if( stringd[0] == " cabin_wear_unfixable") {
			line = " cabin_wear_unfixable: 0"
			//
		}
		if stringd[0] == " chassis_wear" {
			line = " chassis_wear: 0"
		}
		if stringd[0] == " chassis_wear_unfixable" {
			line = " chassis_wear_unfixable: 0"
		}


		
		if(strings.Contains(line, "wheels_wear_unfixable[")) {
			ss := strings.Split(line, ":")
			line = ss[0] + ": 0"
			//
		}
		if(strings.Contains(line, "wheels_wear[")) {
			ss := strings.Split(line, ":")
			line = ss[0] + ": 0"
			//
		}

		lines[i] = line



	}

	// convert lines to content_str
	content_str = strings.Join(lines, "\n")

	// write the new content to game.sii
	file_game, err_game = os.Create(game_sii_location)
	if err_game != nil {
		fmt.Println("game.sii cannot write: ", err_game)
		fmt.Scanln() // wait for Enter Key
		return
	}
	defer file_game.Close()
	_, err_write := file_game.WriteString(content_str)
	if err_write != nil {
		fmt.Println("game.sii cannot write: ", err_write)
		fmt.Scanln() // wait for Enter Key
		return
	}
	fmt.Println("game.sii resetted")

	fmt.Println("Press Enter to exit")
	

    fmt.Scanln() // wait for Enter Key

}