# Windows Terminal Custom Theme and Prompt


## Intro
With meetups going virtual all over the world, we've all been given access to sessions which we would not normally be able to attend. During these sessions there have been many questions on how to cusomise terminals, both prompts and colour schemes.

There are [docs](https://docs.microsoft.com/en-us/windows/terminal/customize-settings/color-schemes) on how to customise your theme and [sites](https://terminalsplash.com/) to download them from. But what if you want to use different fonts and a scheme that's not available to just copy?

![My Windows Terminal Ubuntu](/images/my-bash-winterm.png)

This is [Windows Terminal](https://www.microsoft.com/en-gb/p/windows-terminal/9n0dx20hk701?activetab=pivot:overviewtab) with the awesome [Remedy Theme](https://marketplace.visualstudio.com/items?itemName=robertrossmann.remedy) from VSCode, which isn't readily available for Windows Terminal. The font is [Fira Code](https://github.com/tonsky/FiraCode) which contains [ligature](https://en.wikipedia.org/wiki/Orthographic_ligature) support, converting multi-character representations to [glyphs](https://github.com/tonsky/FiraCode#whats-in-the-box). This is how the prompt displays icons for Git, using [powerline-go](https://github.com/justjanne/powerline-go) setup for bash.

This page is a walkthrough of how to set this up.

## Pre-reqs

1. Download [Windows Terminal](https://www.microsoft.com/en-gb/p/windows-terminal/9n0dx20hk701?activetab=pivot:overviewtab)
2. Download [Fira Code Font](https://github.com/tonsky/FiraCode) - note Light [isn't working](https://github.com/tonsky/FiraCode/issues/941) with ligatures
3. Download latest [Remedy iterm2](https://github.com/robertrossmann/vscode-remedy/blob/master/resources/iTerm2/Remedy%20-%20Dark.itermcolors) and place into convert_color folder
4. Download [Visual Studio Code](https://code.visualstudio.com/download)

## Introduction to settings and using a readily available scheme from GitHub

![My Windows Terminal Cmd](/images/my-cmd-winterm.png)

This is a themed command prompt profile, it's a good place to start as it uses a scheme readily available to copy from github called [Firewatch](https://github.com/mbadolato/iTerm2-Color-Schemes/blob/master/windowsterminal/Firewatch.json). If you click on the down arrow at the top of your windows terminal, your Windows Terminal menu opens. At the bottom of the menu is Settings:

![MWindows Terminal Settings](/images/winterm-settings.png)

Choose Settings and the settings.json file will open for editing in VSCode. There are some great [msdocs](https://docs.microsoft.com/en-us/windows/terminal/customize-settings/global-settings) on how to edit you settings. Here you are going start by adding a new scheme.  Schemes are added into the schemes section of the settings.json file, which is near the bottom of the page: 


![Windows Terminal Settings](/images/winterm-schemes.png)


Copy the [Firewatch](https://github.com/mbadolato/iTerm2-Color-Schemes/blob/master/windowsterminal/Firewatch.json) scheme, incuding the curly braces and paste it into the settings.json between the square brackets, it should now look like this:


![Windows Terminal Settings with Firewatch](/images/winterm-settings-firewatch.png)

If you don't like the look of that sceheme, go up a level in the GitHub repo, there are hundreds more, with screenshots [here](https://github.com/mbadolato/iTerm2-Color-Schemes)!

The scheme is named "Firewatch", this name is used to reference the scheme. The reference can be used in the default profile, which will apply it to all profiles, or a single profile. A profile is each terminal type that has been setup in Windows Terminal, such as Command Prompt, Powershell or [WSL](https://docs.microsoft.com/en-us/windows/wsl/faq).

In this example the firewatch scheme is only to be applied to the command prompt. Scroll up in the settings.json file until you see *// Make changes here to the cmd.exe profile.*:

![Windows Terminal Settings Cmd](/images/winterm-settings-cmd.png)

To reference the firewatch scheme add ``` "colorScheme": "Firewatch" ``` into the cmd.exe profile, make sure you add a comma after ```"hidden": false,```:

![Windows Terminal Settings Cmd](/images/winterm-settings-cmd-firewatch.png)

Save the file, if you have the Command Prompt open in your windows terminal it will change automatically. Note that settings at the profile level overrride those at the default level.

## Setting up your own theme and converting one from iTerm2

The schemes structure uses a JSON variant of the iTerm2 format which can be ported for use by many terminals as listed on the [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes) GitHub page. The scheme displayed on the first screenshot in this page is from the [Remedy](https://marketplace.visualstudio.com/items?itemName=robertrossmann.remedy) VSCode theme. The [Remedy GitHub repo](https://github.com/robertrossmann/vscode-remedy) contains the [.itermcolors](https://github.com/robertrossmann/vscode-remedy/blob/master/resources/iTerm2/Remedy%20-%20Dark.itermcolors) file which enables the same theme in iTerm2. You can also export schemes from iTerm2 .itermcolor files.

The .itermcolors file looks complicated, but it can be ported for use in Windows Terminal using an itermcolors to hex convertor. There are many available, 

The [color schemes](https://docs.microsoft.com/en-us/windows/terminal/customize-settings/color-schemes) section of msdocs gives the format required to create your own scheme. You can paste a second section into the schemes array (the square brackets) and give it a name other than "Firewatch". For example, this could be "powershellScheme", add the named reference to another profile using ``` "colorScheme": "powershellScheme" ```. As you edit the hex presentation of the colours for powershellScheme.


There are also black and white themes which is used when writing for books, light text on dark screenshots doesn't print well.

Settings for Windows Terminal and theme converter from iterm2

``` bash
sudo apt update
sudo apt install golang-go
go version
```


## Convert file and add theme to profiles.json

Windows Terminal supports many iTerm2 

## Create GUIDS

https://www.thomasmaurer.ch/2016/03/create-guid-using-powershell/

## SSH With terminal

https://www.thomasmaurer.ch/2020/05/how-to-ssh-into-an-azure-vm-from-windows-terminal-menu/


