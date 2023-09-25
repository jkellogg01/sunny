# Sunny

## Description

Sunny is a simple CLI tool that tells you the weather.

That's it, really.

Sunny is my first real Go project and has helped me learn some things (certainly not everything) about working in the Go language.

## Usage

Sunny is relatively simple to use, but does require a little bit of fanagling[sic] before it can realize its full potential.

### Configuration

First, you need a config.json in the main directory of the program. This will contain both your default city and your API key; the former is optional but adds to the intended ease of use, the latter is essential.

config.json should contain two key-value pairs:

```json
{
  "api_key": "{your api key}",
  "home_city": "{city}"
}
```

Sunny uses the [OpenWeather API](https://openweathermap.org/), and their accompanying geocoding API. These APIs are free to use, and will both work using the same API key.

The `"home_city"` value simply provides a default city name to search when the `-c` flag is not used to specify a city to search for; it can be ommitted but can add a large degree of convenience especially if your home city's name takes a particularly long time to type or requires you to add multiple [specifiers](#city-specifiers).

### Running the Program

Sunny is very simple by nature of my inability to do anything much more complicated. As such, given that you have completed your config.json correctly, all you need to do to run sunny is:

```zsh
$ sunny
```

This will retrieve data about whichever city you have set as your default/home city.
You can also retrieve data about a specific city with the `-c` flag:

```zsh
$ sunny -c denver
```

As of now, that's really all there is to it.

### City Specifiers

Many cities share names, and as such you may need to specify a state or country. These values should be separated by commas, not spaces, and any names containing spaces should use quotes: `{city},{state},{country}`

For example:

```zsh
$ sunny -c paris
```

will retrieve data about the weather in Paris, France. Someone looking for the weather in Paris, Texas, would need to be a little bit more specific:

```zsh
$ sunny -c paris,tx,usa
```

The way this is interfacing with OpenWeather's geocoding API is relatively indelicate at the moment, so a good rule of thumb is to follow what is laid out by [their own documentation](https://openweathermap.org/current#geocoding).

## Installation

I'm new enough to the go language that I'm not really of much help in this department. I will be doing some research but a better description of this process would be greatly appreciated as well.

## Contribution

Given that this is a learning project for me, contributions other than in the aforementioned installation section of this README file will not be accepted for the time being. I'm well aware that there are many things I could be doing better here, and I am more interested in learning how to do so than I am in the app working optimally in the shortest time frame possible.
