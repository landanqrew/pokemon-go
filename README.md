# pokemon-go
a classic pokedex CLI tool (written in Go)

## Overview

This CLI tool allows users to interact with the PokeAPI to explore Pokémon data. Below is a list of available commands:

| Command   | Description                                                     | Arguments        |
|-----------|-----------------------------------------------------------------|------------------|
| `exit`    | Exit the Pokedex                                                |                  |
| `help`    | Displays a help message                                         |                  |
| `map`     | Displays a map of locations                                     |                  |
| `mapb`    | Displays the previous page of the map                           |                  |
| `explore` | Displays a list of Pokémon in a specific location               | `location`       |
| `catch`   | Attempts to catch a Pokémon and add it to your Pokedex          | `pokemonName`    |
| `pokedex` | Displays a list of Pokémon you have caught                      |                  |
| `list`    | Displays a list of Pokémon you have caught (alias for `pokedex`)|                  |
| `inspect` | Displays details of a Pokémon caught in your Pokedex            | `pokemonName`    |

## Lessons Learned

Building this CLI application helped to build my understanding programming and system design in a couple specific ways:

### 1. Caching with `pokecache`

The `internal/pokecache` package implements an in-memory cache to store API responses, significantly reducing redundant network requests and improving performance.

*   **Implementation Details:**
    *   `Cache` struct: Stores `Entries` (a `map[string]CacheEntry`) and a `sync.RWMutex` for concurrent access.
    *   `Add(key string, value []byte)`: Adds an entry to the cache with the current timestamp. It uses a write lock (`c.mu.Lock()`) to ensure thread-safe addition.
    *   `Get(key string)`: Retrieves an entry. It first attempts to read with a read lock (`c.mu.RLock()`) for concurrency. If a cache miss occurs, it fetches data from the API *without* holding the lock, then acquires a write lock (`c.mu.Lock()`) to add the new entry, employing a double-checked locking pattern to avoid race conditions.
    *   `reapLoop(interval time.Duration)`: A goroutine runs periodically to clean up expired cache entries. It acquires a write lock over the entire loop to prevent race conditions during iteration and deletion. This design prevents `map concurrent write` panics.

### 2. HTTP Client with `api.Client`

The `internal/api` package provides a robust HTTP client for interacting with the PokeAPI, leveraging the caching mechanism.

*   **Implementation Details:**
    *   `Client` struct: Holds a reference to `pokecache.Cache` and the `BaseURL` for API requests.
    *   `NewClient(cache *pokecache.Cache, baseURL string)`: Constructor for the API client, injecting the cache dependency.
    *   `GetResponse(relPath string)`: This method is the primary way to fetch data. It first attempts to get the response from the `Cache`. If the data is not in the cache or has expired, it makes an HTTP request via `web.GetResponseBytesBaseUrl()` and then adds the response to the cache. This design prioritizes cache hits while ensuring fresh data when necessary.

### 3. State Management with `state.AppState`

The `internal/state` package manages the application's global state, including pagination for map navigation and the user's Pokedex.

*   **Implementation Details:**
    *   `AppState`: A global variable of type `*State` that holds the current application state. This may not be the best way to implement state, but I found that it worked, so I stuck with it.
    *   `State` struct: Contains `LocationNamePage` for map pagination and `Pokedex` (a `map[string]pokemon.Pokemon`) to store caught Pokémon.
    *   Methods like `GetLocationNamePage()`, `IncrementLocationPage()`, `DecrementLocationPage()`, and `ResetLocationPage()` manage navigation state.
    *   `AddPokemon(pokemonName string, pokemon pokemon.Pokemon)` and `GetPokemon(pokemonName string)`: Handle adding and retrieving Pokémon from the user's Pokedex.
    *   `ListPokemon()`: Provides a list of caught Pokémon names.
    *   Initial thoughts on placing `Pokedex` directly in `config.Config` led to an `import cycle`. Moving it to a dedicated `state` package resolved this, emphasizing the importance of separating configuration from mutable application state to maintain a clean dependency graph.
