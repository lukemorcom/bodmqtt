# BodMQTT

Simple publisher/broker combo mainly intended for coordinating ESP32 microcontrollers in my house.

Plan: Go publisher aggregates various API data and publishes to specified topics based on config.yml. Runs alongside the broker for ease.


## Configuration

Create a file `config.yaml` in the root where `apis` is a list of the services you want to consume, and `events` is a list of actions that tie APIs to pubsub topics.
``` yaml
apis:
  - name: home_router
    url: 192.168.1.1:80
    strategy: ping

events:
  - name: ping_router
    # The name of the API from above
    api: home_router
    topic: home/status
```


## Strategies

### Ping

Currently only ping is supported, it will publish "up" or "down" based on the outcome of the ping


## Roadmap

- More strategies
