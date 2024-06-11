# Biome - A Lightweight Container Orchestrator
![biome](https://github.com/Sanniv2002/Biome/assets/100380315/2545aff1-b965-4721-854a-5614eaffd092)

## What is Biome?
Deploying your containerized server has never been easier. Simply upload your container to [DockerHub](https://hub.docker.com/), send a POST request to Biome's Gateway, and voilà—your server is live on the web!
But that's not all. Biome does more than just deploy. It handles scaling, routing, and even those pesky container crashes, so you can focus on what you do best—coding.
Let Biome handle the heavy lifting. After all, superheroes deserve a break too.

Here’s a pro tip from the creeper whisperers: you should prefer running a single image Biome. It’s like focusing on building one epic castle instead of spreading yourself thin across multiple builds. This way, you'll encounter fewer faults and keep those pesky bugs at bay—no TNT required!

## Why Biome is fast and lightweight.
Biome's internal servers are loosely coupled and communicate using the [gRPC](https://grpc.io/) protocol, ensuring efficient and reliable interactions. Biome is designed to maximize concurrency and minimize execution time and CPU idle times by leveraging multiple Go routines. This means your container orchestration is not only swift but also optimized for performance.

# Biome's Architecture
![archBackend](https://github.com/Sanniv2002/Biome/assets/100380315/8e81491e-6cae-4179-baba-08ee0b9768df)


## Meet Biome's Microservices and CRON Jobs.

## API Gateway

Welcome to the entry point for your containerized application! This is your master control panel for starting and stopping containers. Direct communication with Biome's internal APIs is off-limits; the Gateway is your sole interface for managing your container services. Need to start or stop your container service? This is your command center. Send your requests here, and the Gateway will take care of the rest, seamlessly orchestrating your containers without exposing the intricate inner workings of Biome.

## Container Init Service

Meet Biome's Container Conductor, the microservice with a mission! Just like an Enderman in Minecraft, this service is here to start and suspend containers with a touch of magic. You know how Endermen are always building and moving blocks around? Well, our microservice does the same—minus the creepy teleportation and block stealing.

## Store Service

Introducing the brain and memory warehouse of Biome, the esteemed cornerstone of Biome's operations! This service serves as the conductor of our centralized storage system, orchestrating the flow of data, CPU, memory stats, and retrieval of all vital information essential for keeping Biome in full swing. Without this repository, Biome would indeed be akin to a cosmic void—empty and directionless!

## Monitoring CRON

The logging and monitoring system for container states. This essential component diligently oversees the status of all containers, ensuring operational efficiency and reliability.

## Autoscaler CRON

This vital component dynamically adjusts container instances, scaling up or down based on their memory and CPU usage. Its primary function is to optimize resource allocation, ensuring efficient operation.

## Recovery CRON

A vigilant guardian against unexpected crashes and shutdowns. This essential component monitors and swiftly restarts containers, ensuring uninterrupted operation and maintaining system reliability.

## Portal

Meet [Portal](https://github.com/Sanniv2002/Portal), Biome's dedicated load balancer. This essential component, powered by the Round Robin algorithm, ensures equitable distribution of workloads across all servers.

> [!CAUTION]
> As Biome takes its first steps in development, this initial version was brought to life within a week, albeit with the acknowledgment that bugs may be present. I remain committed to addressing and resolving any issues swiftly as I continue to refine and enhance Biome's capabilities.

> [!NOTE]  
> Feel free to clone the repository and reach out if you're interested in contributing to Biome's development.
