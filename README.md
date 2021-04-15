# ShoppingCart API 

A simple solution for a shoppingCart API, including some routes, for fetching/creating product, users and coupons.

## Usage

Documentation for usage can be found in [Wiki](https://github.com/MarlonCorreia/shoppingCart-LI/wiki) of the repo. 

## Dependecies

- Docker
- Docker-compose

## How to Run

You can change the mode that you want to run the applicando by changing `GIN_DEBUG_MODE`. This will define if you want to run a development or release mode.

After that, you just need to run:

```bash
docker-compose up --build
```

The service will be available at: 

`127.0.0.1:8080/`
