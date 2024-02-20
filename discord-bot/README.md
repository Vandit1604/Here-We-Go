# Discord Bot in Golang

I'll start with basic bot. After I get the hang of it, I'll start adding features that can help tech communities to help people find mentors:

- Owner of a server can add mentors and their discord IDs
a. Some problems, I can think of is where will this data be stored? Maybe a channel specially for the bot which will contain the list of domains which can be the first message of a channel.

- Members of a discord server can use /mentor <Domain-Name>
a. There should be a limit by default on how many time someone can ask for help. 
- Now the bot will tag the mentor that can help according to some type of algorithm to equally provide mentors
a. How about also sending the person a DM providing the details of mentee(?)
- I basically plan to give a TA/Helper to discord communities

Above is the inital scope but later I'll be adding features like:
- More Domains can be added via the owner

## Flow of the bot:

1. You type !mentor <Domain Name>
2. Bot will send a message in this format
	Hi <Name of Mentee>, If you're interested in getting mentorship in this <Domain Name>
	<Name of Mentor> would love to talk to you about <Domain Name>
	Cheers! 🍀

^^ Message can be better

3. Simultaenously, The Bot can send the mentor a DM about someone asking for help.
	Hi <Name of Mentor>, <Mentee ID > is interested in getting mentorship in this <Domain Name>
	If you had a chat/meeting with him respond with /done or /gifted to close this DM
	Cheers! 🍀

## Adding Mentor Data

I'll use Discord Roles to assign someone as mentor and get that list but wouldn't it be large in a community with over 10k members?

## Testing

Perform a GET request to 

```
https://discord.com/api/v9/applications/{APP_ID}/commands
```


