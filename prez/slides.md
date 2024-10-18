---
# You can also start simply with 'default'
theme: dracula
title: We don't talk about Bruno
highlighter: shiki
drawings:
  enabled: false
lineNumbers: true

transition: slide-left
class: text-center
hideInToc: true
mdc: true
layout: cover
---

# Ne parlons pas de Bruno

BDX I/O 2024

<div class="w-md ma">
    <img src="/images/bruno.jpg" alt="Bruno" />
</div>

<div class="abs-br m-6 flex gap-2">
    <a href="https://github.com/ImFlog/bruno-talk" target="_blank" alt="GitHub" title="Open in GitHub"
      class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
      <carbon-logo-github />
    </a>
</div>


---
layout: profile
speaker: Florian Garcia
job: Tech Lead
company: Synapse Medicine
tags: [Backend, Cloud, Health tech]
website:
url: https://synapse-medicine.com
name: synapse-medicine
image: /images/florian.png
twitter: ImFlog
---

---
layout: two-cols
---

# ~~Who~~ What is Bruno ?

[//]: # (NB: I'm really bad at CSS)

<br />
<br />

* Open source API client
* Like Postman, Insomnia ...
* Written in Javascript
* Offline only
* Cute dog logo

::right::

<br />
<br />
<br />
<br />

<div class="w-xs ma">
    <img src="/images/logo-bruno.png" alt="Bruno" style="width: 250px"/>
</div>

---
layout: default
---

# Demo

[//]: # (TODO: Add a link of the repository + a photo of the UI ?)

* Variables, environments and secrets
* Authentication
* Scripting
* Testing
* CLI

---
layout: section
---

# What about Postman ?

---
layout: default
zoom: 0.6
---

## Quick comparison

|                             | Postman                                                                                                                                                            | Bruno                                                          |
|-----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|
| Request protocols           | HTTP, GraphQL, gRPC, WebSocket, Socket.IO, MQTT, SOAP.                                                                                                             | HTTP and GraphQL.                                              | 
| Collection format           | Open source JSON format.                                                                                                                                           | Plain text markup language called Bru.                         |
| Collection storage          | Proprietary cloud                                                                                                                                                  | In your file system.                                           |
| Collection run              | Free and basic plan at 25 runs per month. Professional at 250 runs, unlimited for Enterprise.                                                                      | Unlimited runs, anywhere.                                      |
| API requests transfert      | Through Postman servers in the web app. [Some reports](https://community.postman.com/t/working-in-offline-mode/20174/51) of proxied requests from the desktop app. | Directly from your computer.                                   |
| Offline ?                   | [No](https://community.postman.com/t/working-in-offline-mode/20174/37)                                                                                             | Yes.                                                           |
| Collaboration               | Paid feature to share with more than 3 users.                                                                                                                      | Free, just use git                                             |
| Scripting with NPM modules  | Using workarounds like pulling from a CDN or copying the library code in a variable.                                                                               | Use a package.json, like usual                                 |
| CLI                         | Postman CLI or Newman. Both have limits on Postman server API calls amount based on purchased plan.                                                                | [Bruno CLI](https://docs.usebruno.com/bru-cli/overview). Free. |
| Local performance testing ? | Yes with run and user limits.                                                                                                                                      | No. But planned in the golden edition.                         |

**But** Postman has features that Bruno doesn't have (yet).

---
layout: two-cols
---

## Pricing ?

![PostmanPricing](/images/pricing-postman.png)

::right::

![BrunoPricing](/images/pricing-bruno.png)

<!--
Il y a aussi une version pour une utilisation personnelle qui coûte 19E en one-time payment pour 2 ordi et 2 ans de maj.
-->

