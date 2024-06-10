# Biome - A Lightweight Container Orchestrator
![biome](https://github.com/Sanniv2002/Biome/assets/100380315/2545aff1-b965-4721-854a-5614eaffd092)

## What is Biome?
Deploying your containerized server has never been easier. Simply upload your container to [DockerHub](https://hub.docker.com/), send a POST request to Biome's Gateway, and voilà—your server is live on the web!
But that's not all. Biome does more than just deploy. It handles scaling, routing, and even those pesky container crashes, so you can focus on what you do best—coding.
Let Biome handle the heavy lifting. After all, superheroes deserve a break too.

## Why Biome is fast and lightweight.
Biome's internal servers are loosely coupled and communicate using the [gRPC](https://grpc.io/) protocol, ensuring efficient and reliable interactions. Biome is designed to maximize concurrency and minimize execution time and CPU idle times by leveraging multiple Go routines. This means your container orchestration is not only swift but also optimized for performance.

# Meet Biome's Microservices and CRON Jobs.
![pikaso_texttoimage_minecraft-group-photo](https://github.com/Sanniv2002/Biome/assets/100380315/97814ea4-906f-40e1-8131-98ca6ab1b567)

# Container Init Service
![pikaso_texttoimage_Enderman-with-minecraft-scenic-background-and-with](https://github.com/Sanniv2002/Biome/assets/100380315/2bff6eda-6f45-463d-a811-c9d3a519571f = 250x250) Meet Biome's Container Conductor, the microservice with a mission! Just like an Enderman in Minecraft, this service is here to start and suspend containers with a touch of magic. You know how Endermen are always building and moving blocks around? Well, our microservice does the same—minus the creepy teleportation and block stealing.
