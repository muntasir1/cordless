# Cordless

| OS | Build-Status |
| - |:- |
| linux | [![CircleCI](https://circleci.com/gh/Bios-Marcel/cordless.svg?style=svg)](https://circleci.com/gh/Bios-Marcel/cordless) |
| darwin | TODO |
| windows | [![Build status](https://ci.appveyor.com/api/projects/status/svv866htsr33hdoh/branch/master?svg=true)](https://ci.appveyor.com/project/Bios-Marcel/cordless/branch/master) |

## Overview

* How to install it
  * [Using pre-built binaries](https://github.com/Bios-Marcel/cordless#using-pre-built-binaries)
  * [Building it from source](https://github.com/Bios-Marcel/cordless#building-it-from-source)
* [Configuration](https://github.com/Bios-Marcel/cordless#configuration)
* [Features](https://github.com/Bios-Marcel/cordless#features)
* [Troubleshooting](https://github.com/Bios-Marcel/cordless#troubleshooting)

Cordless is supposed to be a custom [Discord](https://discordapp.com) client
that aims to have a low memory footprint and be aimed at powerusers.

This project was mainly inspired by [Southclaws](https://github.com/Southclaws)
[Cordless](https://github.com/Southclaws/cordless), which he sadly didn't
develop any further.

The application only uses the official Discord API and doesn't send data to
any third party.

This application is currently a WIP and will change rather fast.

![Cordless 19th January 2019](https://i.imgur.com/ftmvcYM.png)

## How to install it

### Using pre-built binaries

Currently every commit triggers a build for windows and linux, those builds
each produce ready to use binaries. The builds can be found at the top, by
clicking on the respective builds badges.

### Building it from source

First you have to grab the code via:

```shell
go get github.com/Bios-Marcel/cordless
```

In order to execute this command
[you need to have go installed](https://golang.org/doc/install).

In order to execute the application, simply run the executable, which lies at
`$GOPATH/bin/cordless`. In order to be able to run this from your commandline,
`$GOPATH/bin` has to be in your `PATH` variable.

### Configuration

Cordless doesn't need anything besides your Discord token. The application will
ask you to input the token in case it can't be found.

In order to retrieve your token, simply follow the steps in the graphic below:

![Steps to retrieve discord token - by ripcord](https://cancel.fm/ripcord/static/app_misc/discord_token_howto_en-US.png)

After retrieving the token, you have to insert it into the input prompt.

Alternative, you can manually insert it into your configuration file.
In order to find the location of the configuration file, simply run
cordless via your terminal and check the output, as it will tell you, where
its configuration file lies.

The token will be saved on your machine, but it won't be encrypted, so be
careful with your configuration file.

## Features

* Guilds
  * Enter channels
  * See channels
  * See members
* Channels
  * See NSFW flag
  * See group
  * See topic
* Members
  * See Nickname
  * See hoist group
* Chatting
  * 1on1 conversation with a friend
  * Groupchats
  * Talk in a channel
  * Timestamp Just `HH:MM(:SS)`
* See all your friends
* Messages
  * Send messages
  * Mention people using their full username
  * Mention a channel
  * Edit last message
  * Delete last message
  * Multiline

## Troubleshooting

If you happen to encounter a crash or a bug, please submit a bug request.

In case that you simply can't use any shortcuts that the application has, this
might be due to your terminal emulator accepting those instead of letting
cordless handle them.