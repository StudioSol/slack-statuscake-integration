StatusCake check notifications for Slack
========================================

A tiny go app that receives webhooks from StatusCake and re-posts them as Slack formatted hooks.

There is a running version of this code on Heroku that you are welcome to use...

https://slack-statuscake-integration.herokuapp.com/

## Here's how to get setup...

1. Generate an incoming webhook in the Slack integration settings e.g. `https://hooks.slack.com/services/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`
2. Switch the domain from `hooks.slack.com` to `slack-statuscake-integration.herokuapp.com`
3. Create a Group on StatusCake with pointing to this url in field `Webhook URL`. e.g. `https://slack-statuscake-integration.herokuapp.com/services/T024XLT1F/B031BS1D0/C4YkI21H6jPQ59PHLQLD3S21`
4. Select method POST and save a group. 


Alternatively you could host the code yourself.