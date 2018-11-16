# VSTS Release Scribe

[![Build Status](https://travis-ci.org/payvision-development/scribe.svg?branch=master)](https://travis-ci.org/payvision-development/scribe)

<p align="center">
  <br>
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/scribe.png" alt="Scribe" width="250">
  <br><br>
</p>

The goal of this project is to capture VSTS service hooks events, specifically those related with the release process and register the progress in the Freshservice platform.

Documentation: [Execution Flow](https://github.com/payvision-development/scribe/wiki/Execution-Flow)

## Configuration

Web Hooks provides a way to send a JSON representation of an event to any service, to start sending this events go to your VSTS project service hooks page and configure the VSTS event to fire it when any of this event types occurs:

- Release deployment started: `ms.vss-release.deployment-started-event`
- Release deployment approval pending: `ms.vss-release.deployment-approval-pending-event`
- Release deployment approval completed: `ms.vss-release.deployment-approval-completed-event`
- Release deployment completed: `ms.vss-release.deployment-completed-event`

Service Hook configuration:

- Service Hook URL: https://scribe.app/vss-release  
- Resource details to send: All 
- Messages to send: All 
- Detailed messages to send: All 

For more information about how to configure VSTS service hooks events go to: https://docs.microsoft.com/en-us/vsts/service-hooks/services/webhooks?view=vsts

## Build and execution

Build the docker image:

    docker build -t scribe .

Run the image with the required environment variables:

    docker run --rm -it 
        -e SCRIBE_FRESHSERVICE_URL="https://foo.freshservice.com"
        -p 8000:8000 scribe
