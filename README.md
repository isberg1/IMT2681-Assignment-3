
Connect a chat-board(Slack) with a Professional chat-bot(Dialogflow), and make an API that the chat-bot uses
to get user replies based on our business logic. Statistics for the use of different features will be stored in a
MongoDB database. The application will be deployed on Heroku. but it will be possible to use it locally on our own
server as long as you have public IP-address. The local version will be deployed using docker.

The business logic in this project will be to get a parameter form the chat-bot. Use it to get some information
form the internet. And reply back to the chat-bot.  

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

for full log of all tasks and worktime spent by each team member. Se the repo wiki work log sections.

what we succeeded at:
* Connect a chat-board(Slack) with a chat-bot(Dialogflow)
* Make an API that the chat-bot uses to get user replies
* Make 5 API request handlers
* Use a Database to store statistic
* Use Docker to make a locally deplorably version of the app
*

what we failed at
* Cloud functions



# Idea
  // How original the idea is? Useful?  Or is it something that has been taken out of Online Tutorial?   E.g. TodoList should
  // score 10 out of 100, unless it has unique features not seen anywhere else.											

# Scope
The Plan for the project was for the team members to have 2 tasks each. set up 1 API request handler. And 1 other thing




// How well the project has been specifiied on the Wiki? Was it easy to understand the scope and what the project will do? Were the requirements of the user separated from the technical details of the implementation? Was it all easy to understand?											

# Work



// Given the size of the group and the project scope, how would you judge the execution? Have the group done sufficient amount of work? It has been expected to do approximately 20 hours per person, therefore for 5 people group expect 100 hours of work. Are all the elements of the project in?											

# Code/work quality   


//Same as with the previous assignments: linter, metalinter, tests, test coverage, git usage, branching and commits, continuous integration, quality of the Readme, code organisation and commenting, architectural choices, time tracking, project group work/processes/communication tools, and so on.											








# API for Dialogflow chat-bot
Assignment 3 (Project) in the course IMT2681-2018 (Cloud Technologies) at NTNU Gjøvik.

This application is a RESTful API that can querie other public APIs based on the parameter specified.     
This API is used by [dialogflow](https://dialogflow.com/) to post different messaged to a Slack webhook.

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
"joke"    For a chuch Norris Joke
"dad"     For a dad Joke


```

### log
 Returns all logs of all events that has beed logged (mainly errors).  
```
GET: /log
```

### Old posts
Returns all API requests preveausly sent to this API.
```
GET: /OldPosts
```

###  Statistics
Returns statistics for the different API usage.
```
GET: /statistics
```

### Website
A website where you can test the available APIs we have.    
[Link](https://warm-meadow-53471.herokuapp.com/website.html)

***

## How this app is deployed
 * The app runs in Heroku at https://warm-meadow-53471.herokuapp.com/
 * Database ?
 * OpenStack? Docker?

## Additional information
Created by Alexander Jakobsen, Martin Brådalen, Mats Ove Mandt Skjærstein
