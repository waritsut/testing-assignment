# Testing Assignment Project 
##### By Warit Suttichaitrakul ([GitHub](https://github.com/waritsut))

## Start Container 

build image and run container by Docker Compose in the background (Frontend and Backend)

```
docker-compose up -d
```

## Q1) Find value in the data set
```
X = 3 ,Y = 31 ,Z = 51
```
#### API
path | method | payload | description 
--- | --- | --- | --- 
/sequenceNumbers | GET | - | Get the answer of question 1
/sequenceNumbers/:index | GET | - | Get the value by index


## Q2) Cashier System

#### Calculate the change system flow
[![pic1.jpg](https://i.postimg.cc/vTjKj3z0/pic1.jpg)](https://postimg.cc/4H6wHvjc)

#### API
path | method | payload | description 
--- | --- | --- | --- 
/cashiers | GET | - | Get balance in the cashier
/cashiers/resettings | PUT | - | Reset the cashier

path | method |  description 
--- | --- |  --- 
/cashiers/changes | PATCH | Calulate the change

Example of the calulate the change API payload
```
{
    "itemPrice": 90,
    "receivedCash": 1000,
    "cash" : {
        "oneThousandNote": 1,
        "fiveHundredNote": 0,
        "oneHundredNote": 0,
        "fiftyNote": 0,
        "twentyNote": 0,
        "tenCoin": 0,
        "fiveCoin": 0,
        "oneCoin": 0,
        "twentyFiveSatang": 0
    }

}
```
 

#### Project Architecture
[![pic2.jpg](https://i.postimg.cc/Vk3YMFPB/pic2.jpg)](https://postimg.cc/9DY5HZMr)


## Project Snapshot
[![pic3.png](https://i.postimg.cc/0QYvm6pt/pic3.png)](https://postimg.cc/ftR1QRqX)
[![pic4.png](https://i.postimg.cc/v8vMZV8T/pic4.png)](https://postimg.cc/yWW2fWZC)
