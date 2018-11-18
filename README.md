# API for Dialogflow chat-bot
Assignment 3 (Project) in the course IMT2681-2018 (Cloud Technologies) at NTNU Gjøvik.

This application is a RESTful API that can querie other public APIs based on the parameter specified.     
This API is used by [dialogflow](https://dialogflow.com/) to post different messaged to a Slack webhook.

***
## Project Idea
Connect a chat-board(Slack) with a Professional chat-bot(Dialogflow), and make an API that the chat-bot uses
to get user replies based on our business logic. Statistics for the use of different features will be stored in a
MongoDB database. The application will be deployed on Heroku. but it will be possible to use it locally on our own
server as long as you have public IP-address. The local version will be deployed using docker.

The business logic in this project will be to get a parameter form the chat-bot. Use them to get some information
form the internet(get data from online public API's). And reply back to the chat-bot.  

We got the idea for this project form Christopher Frantz.


The Plan for the project was for the team members to have 2 tasks each. set up 1 API request handler. And some other
thing based on available time. It is up to each team member to find there own API request case to implement.
the other thing to be done was originally related to the list of features we had to use(Heroku, Docker, Database, Cloud function, Openstack).

total work all members:
total work Alexander Jakobsen:
total work :
total work :
total work :
total work :

for full log of all tasks and work time spent by each team member. Se the repo wiki "work log" sections.

what we succeeded at:
* Connect a chat-board(Slack) with a chat-bot(Dialogflow)
* Make an API that the chat-bot uses to get user replies
* Make 5 API request handlers
* Use a Database to store statistic and some image storage
* Use Docker to make a locally deplorably version of the app
* make a nicer website(using html and css) for testing API features without using slack


what we failed at
* Cloud functions


we have used discord and in person meetings for group communication. all team members are able to push to github. And view Heroku
project page.



Code quality
* golinter
* go fmt
* gocyclo
* gometalinter -- metalinter

    Everything ok

continuous integration

    automatic deployment form github to heroku. if compile fails then we are notified by mail


in order to run the tests locally you must provide the following environmental variables( reviews will be given our data)

  export MONG_ADDRESS=
  export MONG_DATABASE=
  export MONGO_PASSWORD=
  export MONGO_USER=
  export MONGO_ADDRESS=


cd handlers/
go test . -v -cover

# Using the program

Log in to the Slack Channel you have been invited to
input on of the commands bellow

## available Slack commands
```
"joke"                         For a Chuch Norris Joke
"dad"                          For a dad Joke
"cat gif"                      For a cat git
"hacker gif"                   For a hacker gif
"trending gif"                 For a trending gif
"show dog"                     For a  dog image
"add dog"                      For add a dog image to database
"adopt"                        For remove oldest dog image from database
"how many"                     For count of dogs in database
"show all"                     For show all dog images in database
"search <what to search for>"  For Wikipedia search
```


***


## API calls:
### Help page
Displays this README in html format.

### dialogflow
Used by dialogflow to querie this API.    
Different public APIs are queried based on the POST sent by dialogflow.
```
POST: /dialogflow - Returns a custum API responce based on a parameter "b" in post body

Correct json request:

"queryResult": {
    "queryText": "joke",
    "action": "actionAndParameters",
    "parameters": {
      "b": "joke"
    },
    "allRequiredParamsPresent": true,
    "fulfillmentMessages": [{
      "text": {
        "text": [""]
      }
    }],

acceptable values for parameter b are:
"joke"          For a chuch Norris Joke
"dad"           For a dad Joke
"cat gif"       For a cat git
"hacker gif"    For a hacker gif
"trending gif"  For a trending gif
"show dog"      For a  dog image
"add dog"       For add a dog image to database
"adopt"         For remove oldest dog image from database
"how many"      For count of dogs in database
"show all"      For  show all dog images in database
"search" (queryText: must contain text to search form)  For Wikipedia search
```

### log
 Returns all logs of all events that has been logged since last build/restart (mainly errors).  
```
GET: /log
```

### Old posts
Returns all API requests previously sent to this API (used for debugging and json data structure analyze).
```
GET: /oldPosts
```

###  Statistics
Returns statistics for the different API usage.
```
GET: /statistics
```

### Website
A website where you can test some of the available APIs we have.    
[Link](https://warm-meadow-53471.herokuapp.com/website.html)

***

## How this app is deployed
 * The app runs in Heroku at https://warm-meadow-53471.herokuapp.com/
 * Database ?
 * OpenStack? Docker?

 .env file must have

		MONGO_ADDRESS=<mongodb address:port>
		MONGO_USER=<username>
		MONGO_PASSWORD=<password>
		MONGO_DATABASE=<database name>

Run with: `docker run --name app --rm --env-file ./.env -p 8080:8080 -d celebrian/api:latest`

If you want to build it yourself use: `docker build --rm -t <user>/<image>:<tag> .`
Then: `docker run --name <name> --rm --env-file <path/to/env/file/> -p <port>:8080 -d <user>/<image>:<tag>

## Additional information
Created by Alexander Jakobsen, Martin Brådalen, Mats Ove Mandt Skjærstein with help from Per-Kristian Kongelf Buer
