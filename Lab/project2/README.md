# Technical Manual
This is a system written in Golang, it works with three main technologies

1. Telegram bot and the Telegram's bot API for Golang, we aim to be able to send a message/command to our telegram bot, retrieve the message and perform actions on a Trello's board or start some applications using `Robotgo` depending on the command that was sent 
2. Trello's REST API, We will enable to create a new card/history inside a List/column
3. Robotgo 

## Environment Set Up
I developed most of the project using windows, I just installed Golang but when reached the point where I needed to use Robotgo Golang package I needed to install GCC as stated in the [requirements section](https://github.com/go-vgo/robotgo?tab=readme-ov-file#requirements), I decided to move to a virtual machine so I choosed to use VMWare Workstation 17.6.3 and Linux Mint Cinnamon, so these are the steps to set up the virtual machine

1. Download VMWare Workstation, search exactly that in a web search engine and just enter the website shown, there should be a button that says "download" or something similar, you will be directed to a website from an entity called "broadcom", you have to register and login, there you should find a section that says "downloadable software" or something similar and there you can search for a text that says "vm ware" using ctrl + f and when you click it, choose workstation or workstation pro, in the next screen you have to check the "accept terms and conditions", it might act weird for some reason but you should be able to refresh the page or go back and forward again, once you accept you will be able to download by clicking the cloud icon, then you have to enter some personal info and the download should start

2. Download Linux Mint Cinnamon, just download the iso image from their website

3. Create a virtual machine I choosed a VMware ESX version 8, if you get an error you will have to edit virtual machine settings and in the Processors section disable "Virtualize Intel VT -x/EPT or AMD-V/RVI" 

4. Install Go following the instructions from their website, you might need to run any code they give you using the super user like this `sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz`, I decided to add Golang binary to my bash profile instead of modifying `PATH` env variable directly, so I ran `nano ~/.bashrc` and pasted at the bottom `PATH=$PATH:/usr/local/go/bin`, then just reload the termina with `source ~/.bashrc`

5. The only requirement to use `robotgo` in this context is to install this package in Linux `sudo apt install gcc libc6-dev`, there is other packages you could install as well but that depends on what you're trying to do with `robotgo`

## Telegram Bot and API
This is the backbone of our app, a user sents a command, we poll it using Telegram's bot REST API using the wrapper we will talk about below and then interacting with Trello's REST API we will perform actions on Trello's board, also we will be able to start some applications on our computer by using the library `robotgo`

### Create your own bot
Basically, follow [this](https://core.telegram.org/bots/features#creating-a-new-bot) tutorial

1. Access the Telegram app, and look for `BotFather`, there you can send the command `/start` and then `/newbot`, then choose a name for your bot and then a username for your bot, this will give you an API token in the following format
```
110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
```

2. Save the token in a safe place, you mustlikely you should not share this token and neither push it to a version control tool like git even if your reporitory is private in tools like github, instead you should do something similar to what I'm doing here, create an `.env` file in the same directory as where your go file with the `package main` that has the `main` function is, this is the entry point to your app. In this `.env` file add the key with the following format

```
TG_API_KEY=110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
```

### Using the Telegram Bot API
There is two ways you can start getting the messages users sents to your bot, and it may deppend on your use case. Your options are

1. Using a poll based solution, this means you basically will set a client and this client will create a channel by which messages can be sent and received. This is probably usefull if you want to run the logic that will serve your bot on a personal machine like a desktop or laptop

2. Using a webhook, here you have to tell telegram the URL you want them to forward your messages to, this could be an on premise or a cloud based server, this way you could run logic using your company's database or services

I used the poll base solution, [this library](https://github.com/go-telegram-bot-api/telegram-bot-api) written in Golang is a wrapper around telegram's [REsT bot API](https://core.telegram.org/bots/api), this wrapper is fetching updates all the time inside an infinite loop in a goroutine and communicating responses through a channel. **Be aware you can't have bot working at the same time**, you can not have a webhook and getting updates through Telegram's Bot REST API at the same time.

Once you have your token set up, you can try it in your terminal/console using `curl`

```
export TELEGRAM_TOKEN=110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw &&
curl https://api.telegram.org/bot$TELEGRAM_TOKEN/getMe

```

Simple examples on how you can ennable your poll based solution, using this library, to handle commands sent from users can be found in their repository provided in prev notes

## Trello REST API
Trello is a Kanban based project management software, it helps you organize a project's tasks

1. To be able use trello first you have to register, I registered using my Google's account

2. You should follow the initial tutorial but it will basically explain you that in trello it exists workspaces, these are like repositories or groups of boards, basically a way to limit access, set scope, group boards, boards is what we will be actually working with, you can either create a new board or just use the one that Trello creates for you when you create your profile, to create a new board you can click the trello icon on the left top corner, in there you can even create a new workspace so then you can scope/contain a new board inside that workspace and therefore to people inside that workspace only. On the top you have a bar where you can enter a name, click "create" and then click "create new board". Your board contains columns but Trello calls them "lists", these columns contains "cards", they represent the different phases a task can go through for example from start to end it could go from "to do/backlog > in progress > done" but the naming is up to you also you can add or remove columns but probably you want to define these lists based on the Agile/Kanban/Scrum metodologies.

3. With a board created you have to register to their "Power-up" to be able to use their REST API [here](https://trello.com/power-ups/admin), you can check a tutorial on how to register [here](https://www.youtube.com/watch?v=ndLSAD3StH8&t=6s), their official documentaiton is [this](https://developer.atlassian.com/cloud/trello/guides/power-ups/managing-power-ups/#power-up-admin-portal). You will be given an "api key", a "secret" and a "token", again you probably wouldn't want to share this with anyone you don't trust nor submit them to a repo even if it is private, store them in your `.env` file in the following format and names

    ```
    TRELLO_API_KEY=348j0235c10z12345678ec222bin3326
    TRELLO_SECRET=9b777na70m9c4m235ca39d9d447m3e40009b8e17888o133pa88a9e444ilf9zz
    TRELLO_TOKEN=ATTA71114444m56o9388i111g15d765558y1fr8jj7b12uh6c8d47ff46bv19477777fL8E6B555
    ```

4. In order to be able to interact with your board you need to get the board ID, you can get it by navigating to the board and for example in this url `https://trello.com/b/xitasdwe/ia` the ID will be `xitasdwe`, add it to your `.env` file

    ```
    TRELLO_BOARD_ID=xitasdwe
    ```

    You can test your board with the following command, this will give you the lists in your board

    ```
    export TRELLO_API_KEY=348j0235c10z12345678ec222bin3326 &&
    export TRELLO_TOKEN=ATTA71114444m56o9388i111g15d765558y1fr8jj7b12uh6c8d47ff46bv19477777fL8E6B555 &&
    curl --request GET \
    --url "https://api.trello.com/1/boards/xitfceCe/lists?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
    --header 'Accept: application/json'
    ```

In our app We will use [this endpoint](https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-lists-get) to get the lists(todo, in_progress and done) in our board and [this endpoint](https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-put) to update a card, this will let us update the list a card belongs to which will let us move a card to a different list

## Robotgo
This library is used to do RPA and GUI automation in Golang, we will use it to serve a command called `/startdaily` in our Telegram's bot which will take control of the mouse and keyboard for a few seconds to open a few tools, we will use it to open Mozilla, click on one of our bookmarks which is the google meet screen, then press `ctrl + t` to open a new tab, and open another bookmark, this time it is the trellos board we are interacting with and then open the terminal to change to "documents" directory and from there launch "vs code". You can find the documentation in their [repository](https://github.com/go-vgo/robotgo?tab=readme-ov-file#requirements)

This function was helpful to get the coordinates to automate the mouse clicks, it will print on console the coordinates of your mouse every second

```
func main() {
    for {
        fmt.Println(robotgo.Location())
        robotgo.Sleep(1)
    }
}
```

## Building the app
1. If statring a program from 0 first create the go module with `go mod init <a_name_for_your_module>`

2. create a go file like `app.go` add the main package on the top with `package main`

3. add any import with
```
import (
    "fmt"
    "os"
)
```

4. add the main function
```
func main() {
    fmt.Println("Hello World")
}
```

5. To run the app, just run this command on your terminal, change direactory to the same directory as your file with the main package, do not include the ">" sign, this is just to indicate different lines or commands
```
> cd Documents
> go run .
```

6. If you want to create an executable to start it as a background process you can do that by running the following command
```
go build .
```

7. If you need to download packages, just import them and then run `go mod tidy`

## Deleting the Project
1. Delete your bot using the `BotFather` chat in Telegram

2. Delete the power-up Key in Trello

# User Manual
## MauIARob
This is a telegram bot that lets you control a Trello board by creating new cards and move them around the different lists in the board. To interact with it, follow these instructions:

1. Search for the bot by name of "MauIARob" in your telegram

### Commands Available

* You can start the tools for your daily with `/iniciar_daily`. This will start two tabs in Mozilla, one is Google meet and the other is the trello's board we are interacting with as well as to open the directory that contains the source code of our program in VS code

* You can create new cards in Trello with the following command
    ```
    /historia
    <A title for your card>
    <the description for this task>
    ```

    Make sure you follow the same format, first line is the command name, second line is the title of the card and third line is the description of the card/task. When this creates a card successfuly you will get two messages, one is a confirmation and the second is a command you can use to move the card you created to another list, you just have to copy it and define the list you want to move your task to

* You can move your card to another list/column using the following format
    ```
    /move
    list:
    <name of the list, either todo, in_progress or done>
    ID:
    <the card ID>
    Name:
    <Name of the card>
    ```

    Again you have to make sure to follow this pattern, this means you'd have 7 lines in this command

