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

- Service Hook URL: https://myhost/vss/release  
- Resource details to send: All 
- Messages to send: All 
- Detailed messages to send: All 

![VSTS Service Hook creation](https://raw.githubusercontent.com/payvision-development/scribe/master/img/service-hook-configuration.gif)

Create one service hook for each release deployment event:

<p align="center">
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/service-hooks.png" alt="Service Hooks">
  <br>
</p>

For more information about how to configure VSTS service hooks events go to: https://docs.microsoft.com/en-us/vsts/service-hooks/services/webhooks?view=vsts

## Freshservice

This configuration will result in a detailed and realtime updated Freshservice Change, most of the descriptions constains links to each item in Azure DevOps / TFS portal.

<p align="center">
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/freshservice-change.png" alt="Freshservice Change">
  <br>
</p>

## Build and execution

Build the docker image:

    docker build -t scribe .

Or download the image from  `docker pull payvisiondevelopment/scribe:1.1.0`

Run the image with the required environment variables:

```shell
docker run --rm -it 
    -e SCRIBE_USER="user"
    -e SCRIBE_PASS="pass"
    -e SCRIBE_FRESHSERVICE_URL="https://foo.freshservice.com"
    -e SCRIBE_FRESHSERVICE_EMAIL="hulk@outerspace.com"
    -e SCRIBE_FRESHSERVICE_APIKEY="key"
    -e SCRIBE_VSTS_APIKEY="key"
    -p 8080:8080 scribe
```

Check the health endpoint:

```json
GET /status HTTP/1.1

{
    "Service": "Scribe",
    "Description": "VSTS Release event integration with Freshservice",
    "Status": "OK",
    "Version": "1.1.0",
    "Info": {
        "Started": "2018-01-01T00:00:00.000000000+01:00",
        "Events": 0
    }
}
```