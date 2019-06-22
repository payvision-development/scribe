# Scribe

[![Build Status](https://travis-ci.org/payvision-development/scribe.svg?branch=master)](https://travis-ci.org/payvision-development/scribe)

<p align="center">
  <br>
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/scribe.png" alt="Scribe" width="200">
  <br><br>
</p>

The goal of this project is to capture Azure Devops service hooks events, specifically those related with the release pipeline from [Azure Devops release pipelines](https://docs.microsoft.com/en-us/azure/devops/pipelines/release/?view=azure-devops) keeping track of the progress and recording changes in real time in [Freshservice](https://freshservice.com) platform auditing all the continuously delivery process.

### Documentation

* [Build and execution](https://github.com/payvision-development/scribe/wiki/Build-and-execution)
  * [Debugging in Visual Studio Code](https://github.com/payvision-development/scribe/wiki/Debugging-in-Visual-Studio-Code)
* [Execution flow](https://github.com/payvision-development/scribe/wiki/Execution-Flow)
* [Configuration](https://github.com/payvision-development/scribe/wiki/Configuration)
* [Freshservice](https://github.com/payvision-development/scribe/wiki/Freshservice)

### Example

A new release is created from Azure Devops release pipelines, usually you define the release pipeline using stages.

<p align="center">
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/azure-devops-release.png" alt="azure-devops-release">
</p>

Each stage is recorded and will result in a detailed and always updated Freshservice Change for each deployment, most of the descriptions constains links to each item in Azure DevOps portal.

<p align="center">
  <img src="https://raw.githubusercontent.com/payvision-development/scribe/master/img/freshservice-change-example.png" alt="freshservice-change">
</p>
