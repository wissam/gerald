# Gerald
Gerald is a twitch bot written in Go.
It was supposed to be a starting project to help me learn Go and get back to
coding since I am an infrastructure person and I have not programmed in years.
Since then, the idea kept on getting bigger in my head and been experimenting
with different parts of the language.

At the moment it is still in an experimental/prototype stage, the code is all
over the place but this is changing bit by bit. 
When I am done with the experimental/prototyping stage, I am squashing all the
commits into one and calling it: prototype done. Since the majority of my
commits at the moment is something like "saving".

# Project Structure

The project consists of 4 parts (in my head at the moment), the main bot, a
small webhook microservice, a redis server for pub/sub and a database. 

## Webhook
This is a small web server that will be processing events from twitch, ie:
broadcast starting and broadcast ending. This will be useful to start processes
with the bot or even make it join the channel.


## messaging queue (redis pub/sub)
Just a message broker to be able to signal the bot from other parts of the
system, currently from the webhook server only.


## Bot
The main part of the system, where are all action will happen, and by action I
mean a simple emote counter just as something to do for now. This will teach me
goroutine, channels, handling reconnection/timesouts, best practices, database
connection and handling, connecting and using libs and apis, etc.


## Database
SQlite for now on my computer, to be phased out soon and replaced with a mysql
db, or maybe postgres , will see. I usually avoid working with databases even
from the infrastructure side. This has to change, it is limiting me.

# DEVOPSing

## cicd
Github actions, why? because I never used it before. So why not :).
I am gonna do a dual ci with gitlab as well at one point then will see where is
it easier to manage the cd part.

## Multiple environments
This will have a dev, staging and prod environments. Not because it needs it
with me as a lone dev/ops but simply it is good practice.

## Docker
Services will be dockerised where applicable

## Kubernetes
The deployment will be on a hosted kubernetes cluster ie: gke.

## Cloud provider
I am gonna go with google cloud, I got lots of experience with AWS, been
working on it since I started with "devops", I need to expand my capabilities
too.

