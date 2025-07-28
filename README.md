# Talken

Talken is a **peer-to-peer, offline-first messaging system**.
It is designed to work **without any central server**.
It can work inside a **local WiFi network, a personal hotspot, or over Bluetooth**.
The goal is to make devices communicate directly with each other and, when necessary, relay messages through other devices nearby.

This is not a traditional chat app with a server. Talken is a **core library** that client apps (mobile, desktop, CLI) can embed. The clients will provide the user interface and storage, while Talken focuses on:

* Peer discovery
* Peer-to-peer message delivery
* Routing through other peers (multi-hop)
* Key generation and basic encryption
* Offline-first communication

**Important:**
At this stage, Talken is in **early development**.
Only documentation, design notes and planning are available.
The actual core implementation is coming step by step.
The project board for development is here:
[https://github.com/users/m-mdy-m/projects/5](https://github.com/users/m-mdy-m/projects/5)

---

## Why Talken?

Talken is designed for situations where there is **no internet** or no reliable connection:

* Camping, events, isolated environments
* Places with network shutdowns
* Private local networks

In the future, Talken may also work over the internet, but for now the focus is **offline-only**.

---

## How it works (concept)

* Each device runs Talken.
* Each device generates a **public/private key pair** that identifies it.
* Devices discover each other on the local network (LAN or hotspot).
* If two devices are in range, they can exchange messages directly.
* If they are **not** in range, the message can hop through other devices until it reaches the destination.

### Multi-hop example

```
A ---- B ---- D ---- C
```

If A wants to send a message to C:

1. A checks its list of neighbors. A only knows B.
2. A asks B if it knows C.
3. B knows D, and D knows C.
   So the message is sent step by step:

```
A -> B -> D -> C
```

Each peer forwards the message to the next peer it knows.

If the recipient is offline, a peer can store the message temporarily and deliver it later (store-and-forward).

---

### ASCII flowchart (very simplified)

```
[ A ]  --->  [ B ]  --->  [ D ]  --->  [ C ]
   \______________________________________/
            multi-hop relay
```

## Key ideas

* There is **no central server**.
* The network is formed by devices themselves.
* Routing is dynamic. Every device keeps a list of nearby peers and updates it over time.
* When the path is not direct, messages are sent through intermediate peers.

---

## Project status

Right now, **Talken is in the planning and design stage.**
The code is not ready yet. The project structure, issues and drafts are being organized here:

[https://github.com/users/m-mdy-m/projects/5](https://github.com/users/m-mdy-m/projects/5)

The board is organized into these columns:

* Backlog
* Ready
* In progress
* In review
* Done

---

## Contribution

You can open issues for:

* Missing features
* Design suggestions
* Questions about the architecture

Pull requests are welcome once the core structure is ready.

---

## License

MIT License. See LICENSE for details.
