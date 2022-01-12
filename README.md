# NORMA Gateway

Norma is de Engelse benaming van het sterrenbeeld winkelhaak. Norma is geimplementeerd als prototype om specificaties te verifiëren op volledigheid en correctheid. Het betreffen de specificaties voor gevalideerde vragen van het programma KIK-V. 

Norma Gateway is geimplementeerd voor het datastation van een bronhouder van de data. Het datastation is een concept waarin leverancierneutraal data beschikbaar wordt gesteld door een zorginstelling, voor haarzelf en haar omgeving. Met de Norma gateway kan een bronhouder gevalideerde vragen ontvangen van een afnemer. De gateway verifieert het access token van de afnemer en het credential waarmee de gevalideerde vraag is geimplementeerd.

Voor de authenticatie en verificatie wordt gebruik gemaakt van de software van [Stichting Nuts](https://github.com/nuts-foundation/nuts-node). 

**NOTE:**
:warning: De software is "as-is".
Er wordt geen ondersteuning op verleend en is niet voor bedoeld voor productie!
---

## Installeren

```bash
git clone git@github.com:hietkamp/norma-gw.git
```

## Opstarten

De gateway maakt gebruik van Kafka Streams en de Nuts node. Daarnaast is een triplestore noodzakelijk voor de werking van het datastation. Deze afhankelijkheden kunnen gestart worden in docker. Gebruik hiervoor het commando.

```bash
./manage start
```
Start vervolgens de gateway.

```bash
./manage go
```
In de ext directory zijn dockerfiles opgenomen voor elasticsearch en flink. Dit betreft een experiment om de kafka streams in beeld te brengen via kibana. Dat experiment is niet gereed.