# winterm-themes
Settings for Windows Terminal and theme converter from iterm2

## Intro
Example [Windows Terminal](https://www.microsoft.com/en-gb/p/windows-terminal-preview/9n0dx20hk701?activetab=pivot:overviewtab) settings. Powershell and WSL are based on the excellent [Remedy theme](https://marketplace.visualstudio.com/items?itemName=robertrossmann.remedy). Remedy has an [iterm2 profile](https://github.com/robertrossmann/vscode-remedy/blob/master/resources/iTerm2/Remedy%20-%20Dark.itermcolors) which can be converted for use with Windows Terminal.

Defaults use the Fira Code Fonts which include the Powerline glyphs. The glyphs are used in [post-git](https://github.com/dahlbyk/posh-git) and [oh-my-posh](https://github.com/JanDeDobbeleer/oh-my-posh) for PowerShell promt and [powerline-go](https://github.com/justjanne/powerline-go) for WSL.

The customised 

Cmd uses [Firewatch theme](https://github.com/mbadolato/iTerm2-Color-Schemes/blob/master/windowsterminal/Firewatch.json).

There are also black and white themes which is used when writing for books, light text on dark screenshots doesn't print well.


## Pre-reqs

1. Download [Fira Code Font](https://github.com/tonsky/FiraCode) - note Light [isn't working](https://github.com/tonsky/FiraCode/issues/941) with ligatures
2. Download latest [Remedy iterm2](https://github.com/robertrossmann/vscode-remedy/blob/master/resources/iTerm2/Remedy%20-%20Dark.itermcolors) and place into convert_color folder
3. Install GO
sudo apt install golang-go


## Convert file and add theme to profiles.json

https://superuser.com/questions/1455847/how-do-i-get-guid-of-apps-installed-from-microsoft-store


### How to get a Windows Store GUID
https://www.ibm.com/support/knowledgecenter/en/SS8H2S/com.ibm.mc.doc/winphone_enrollment_source/concepts/appid_manual_steps.htm

(Adding Distros after the fact)


