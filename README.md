# SIMPLE WEATHER CLI Using GO

## Introduction

A simple CLI application using golang which fetches weather data form [`https://www.weatherapi.com`](https://www.weatherapi.com) and displays in a tabular format in terminal

## Setup

1. Clone repository and change directory

   ```bash
     git clone https://github.com/nabinthapaa/weather-cli
     cd weather-cli
   ```

2. Create `.env` file at the root

   ```env
     WEATHER_API=<api key from https://www.weatherapi.com>
   ```

3. Build the project and run

   ```bash
     go build -o main main.go
     ./main
   ```

4. To check the arguments that can be pass run

   ```bash
     ./main --help
   ```
