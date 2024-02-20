package bot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	domainName string = ""
	BotToken   string
	AppID      string
	// TODO: Change this to the guild id of the server the bot joins for each discord server
	GuildID string = "1194818825783349288"
)

func Run() {
	// we create a session basically a bot user and give it all the capiability to do stuff
	session, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Couldn't connect to discord ", err)
	} else {
		fmt.Println("Bot is Connected to Discord")
	}
	// we register slash commands here in discord
	registerCommands(session)

	// this is a handler that will be fired on a certain discord websocket api event (WSAPI EVENT)
	session.AddHandler(newMessage)

	// handler for interaction
	session.AddHandler(interactionHandler)

	// its like my bot will do everything without privilages
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	session.Identify.Intents = discordgo.IntentsGuildMembers

	// Opens the session
	if err := session.Open(); err != nil {
		log.Fatal("Couldn't Open the session", err)
	}
	defer session.Close()

	fmt.Println("Bot is running...")
	// Run until bot is terminated

	/*
		We want the bot to know when to terminate itself, It won't know it by itself.
		So, The notify function will check the termination signal from everyhwere it can be ctrl-c or shutdown of my computer or crash.
		Since the Notify function requires a channel through which it will send the `os.Signal` of termination
	*/
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}

func registerCommands(session *discordgo.Session) {
	// this is where we register commands
	_, err := session.ApplicationCommandBulkOverwrite(AppID,
		GuildID,
		[]*discordgo.ApplicationCommand{{
			Name:        "hello",
			Description: "Lists the mentors available for a specific domain",
		}, {
			Name:        "commands",
			Description: "Lists the mentors available for a specific domain",
		}, {
			Name:        "mentor",
			Description: "Lists the mentors available for a specific domain",
		}})
	if err != nil {
		log.Fatal(err)
	}
}

func interactionHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	data := interaction.ApplicationCommandData()

	// embed := &discordgo.MessageEmbed{
	// 	Title:       domainName,
	// 	Description: "#MENTORID is available to mentor you, You can message him and ask him for assistance",
	// 	Color:       rand.Intn(16777215),
	// }

	// Respond to messages
	// find if the mentor role exists or not
	// get mentor id from the roles

	switch data.Name {
	case "hello":
		err := session.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{{Title: "Hello, How can I help you", Description: "I'm your personal TA for your own convienience", Color: rand.Intn(16777215)}},
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

	case "commands":
		err := session.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{{Title: "Commands", Description: "Commands: \n write use slash command `mentor` <Domain Name> to get a Mentor's Assistance. \n *Following are the Domains in which you can get assistance:* \n- Opensource\n- System Programming\n- Web Development\n- App Development", Color: rand.Intn(16777215)}},
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

	case "mentor":
		mentor := getMentorsForDomain(session, interaction, domainName)
		err := session.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{{Title: domainName, Description: fmt.Sprintf("%s", mentor), Color: rand.Intn(16777215)}},
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getMentorsForDomain(session *discordgo.Session, interaction *discordgo.InteractionCreate, domainName string) string {
	members, _ := session.GuildMembers(interaction.GuildID, "", 100)
	fmt.Println(session.RequestGuildMembers(interaction.GuildID, "", 0, "", false))
	fmt.Println(members)
	for _, member := range members {
		for _, role := range member.Roles {
			fmt.Println(role)
			if role == "Mentor" || role == "mentor" {
				return fmt.Sprintf("@%s will help you soon", member.User.ID)
			}
		}
	}
	return "No one with Role `Mentor` or `mentor` was found"
}

// this functions checks every message on discord channel
func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	//checking if the message that is send by is by the bot or not. If yes, Do nothing
	if message.Author.ID == session.State.User.ID {
		return
	}
	args := strings.Split(message.Content, " ")
	if args[0] != "!mentor" {
		return
	}

	if len(args) == 1 || len(args) > 2 {
		_, err := session.ChannelMessageSend(message.ChannelID, "Specify the Domain in this format \n ```!mentor Domain-Name-Without-Space``` \n for Better Assistance")
		if err != nil {
			log.Fatal("Some problem occurred in sending messages", err)
		}
		return
	}

	// TODO: Let users add the domain name with spaces
	domainName = args[1]
}
