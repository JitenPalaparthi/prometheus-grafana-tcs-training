# Add telegram for alerts

    1. Create a bot 
        - open telegram app goto @BotFather
        - in the BotFather chat window type /newbot
        - it  prompts "Alright, a new bot. How are we going to call it? Please choose a name  for your bot." then
        - give a name for the bot. For examle "alertmanager-test".
        - It prompts "Good. Now let's choose a username for your bot. It must end in `bot`. Like this, for example: TetrisBot or tetris_bot."
        - give a name to the bot. example "prometheus_alertmanager_bot"
        - You see a message something like "Done! Congratulations on your new bot.  Just make sure the bot is fully operational before you do this.Use this token to access the HTTP API: 6351822082:AAGf7owyEk1zvAGPoOprfSN4--1MdP4Cwn4
        Keep your token secure and store it safely, it can be used by anyone to control your bot. [Note: The token is important]
    
    2- Search for the bot that just has been created in the search bar
        - search for @alertmanager-test [Note: Search with the name that was given while creating a bot]
    
    3- Create a new group (Any groupname)
        - add newly created bot to the group.
    
    4- Fetch Chat_ID
        - run below command to get Chat_ID
        - curl https://api.telegram.org/bot<token>/getUpdates/
        - example curl https://api.telegram.org/bot6351822082:AAGf7owyEk1zvAGPoOprfSN4--1MdP4Cwn4/getUpdates/

# to install alaertmanager


# to configure alertmanager

    - create a prometheus_rules.yml file
    - change prometheus.yml file to enable rules.
    - create alertmanager.yml file with telegram configuration

# to test the application

    - alertmanager alerts when the node_exporter is down. So run node_exporter and stop it.It should send alerts to the telegram channel