# PokéAPI v2 Overview

This document provides a comprehensive overview of the objects and their corresponding endpoints available in the PokéAPI v2.

## API Information
*   **Consumption-Only**: Only HTTP GET method is available.
*   **No Authentication**: No API keys or tokens are required.
*   **Rate Limiting**: Removed since November 2018, but caching is encouraged for fair use.

## Core Concepts

### Resource Lists/Pagination
*   **Named Endpoints**: `GET https://pokeapi.co/api/v2/{endpoint}/`
    *   Returns a paginated list of resources (default 20).
    *   Example: `https://pokeapi.co/api/v2/ability/`
    *   Query parameters: `limit`, `offset`
    *   Response includes `count`, `next`, `previous`, `results` (list of `NamedAPIResource` with `name` and `url`).
*   **Unnamed Endpoints**: `GET https://pokeapi.co/api/v2/{endpoint}/`
    *   Similar to named endpoints but results contain `APIResource` (only `url`).
    *   Endpoints like `characteristic`, `contest-effect`, `evolution-chain`, `machine`, `super-contest-effect` are unnamed.

### Common Models
*   **`APIResource`**: `{ "url": "string" }`
*   **`NamedAPIResource`**: `{ "name": "string", "url": "string" }`

## Example Usage

To get a list of Pokémon abilities (a named endpoint):

**Request:**
```
GET https://pokeapi.co/api/v2/ability/?limit=5
```

**Expected Response (JSON):**
```json
{
  "count": 248,
  "next": "https://pokeapi.co/api/v2/ability/?limit=5&offset=5",
  "previous": null,
  "results": [
    {
      "name": "stench",
      "url": "https://pokeapi.co/api/v2/ability/1/"
    },
    {
      "name": "drizzle",
      "url": "https://pokeapi.co/api/v2/ability/2/"
    },
    {
      "name": "speed-boost",
      "url": "https://pokeapi.co/api/v2/ability/3/"
    },
    {
      "name": "battle-armor",
      "url": "https://pokeapi.co/api/v2/ability/4/"
    },
    {
      "name": "sturdy",
      "url": "https://pokeapi.co/api/v2/ability/5/"
    }
  ]
}
```

## Available Objects and Endpoints

### 1. Berries
*   **Endpoint**: `GET https://pokeapi.co/api/v2/berry/{id or name}/`
    *   **Objects**: `Berry`, `BerryFirmness`, `BerryFlavor`
*   **Berry Firmnesses**: `GET https://pokeapi.co/api/v2/berry-firmness/{id or name}/`
*   **Berry Flavors**: `GET https://pokeapi.co/api/v2/berry-flavor/{id or name}/`

### 2. Contests
*   **Endpoint**: `GET https://pokeapi.co/api/v2/contest-type/{id or name}/`
    *   **Objects**: `ContestType`, `ContestEffect`, `SuperContestEffect`
*   **Contest Effects**: `GET https://pokeapi.co/api/v2/contest-effect/{id}/`
*   **Super Contest Effects**: `GET https://pokeapi.co/api/v2/super-contest-effect/{id}/`

### 3. Encounters
*   **Endpoint**: `GET https://pokeapi.co/api/v2/encounter-method/{id or name}/`
    *   **Objects**: `EncounterMethod`, `EncounterCondition`, `EncounterConditionValue`
*   **Encounter Conditions**: `GET https://pokeapi.co/api/v2/encounter-condition/{id or name}/`
*   **Encounter Condition Values**: `GET https://pokeapi.co/api/v2/encounter-condition-value/{id or name}/`

### 4. Evolution
*   **Endpoint**: `GET https://pokeapi.co/api/v2/evolution-chain/{id}/`
    *   **Objects**: `EvolutionChain`, `EvolutionTrigger`
*   **Evolution Triggers**: `GET https://pokeapi.co/api/v2/evolution-trigger/{id or name}/`

### 5. Games
*   **Endpoint**: `GET https://pokeapi.co/api/v2/generation/{id or name}/`
    *   **Objects**: `Generation`, `Pokedex`, `Version`, `VersionGroup`
*   **Pokedexes**: `GET https://pokeapi.co/api/v2/pokedex/{id or name}/`
*   **Version**: `GET https://pokeapi.co/api/v2/version/{id or name}/`
*   **Version Groups**: `GET https://pokeapi.co/api/v2/version-group/{id or name}/`

### 6. Items
*   **Endpoint**: `GET https://pokeapi.co/api/v2/item/{id or name}/`
    *   **Objects**: `Item`, `ItemAttribute`, `ItemCategory`, `ItemFlingEffect`, `ItemPocket`
*   **Item Attributes**: `GET https://pokeapi.co/api/v2/item-attribute/{id or name}/`
*   **Item Categories**: `GET https://pokeapi.co/api/v2/item-category/{id or name}/`
*   **Item Fling Effects**: `GET https://pokeapi.co/api/v2/item-fling-effect/{id or name}/`
*   **Item Pockets**: `GET https://pokeapi.co/api/v2/item-pocket/{id or name}/`

### 7. Locations
*   **Endpoint**: `GET https://pokeapi.co/api/v2/location/{id or name}/`
    *   **Objects**: `Location`, `LocationArea`, `PalParkArea`, `Region`
*   **Location Areas**: `GET https://pokeapi.co/api/v2/location-area/{id or name}/`
*   **Pal Park Areas**: `GET https://pokeapi.co/api/v2/pal-park-area/{id or name}/`
*   **Regions**: `GET https://pokeapi.co/api/v2/region/{id or name}/`

### 8. Machines
*   **Endpoint**: `GET https://pokeapi.co/api/v2/machine/{id}/`
    *   **Objects**: `Machine`

### 9. Moves
*   **Endpoint**: `GET https://pokeapi.co/api/v2/move/{id or name}/`
    *   **Objects**: `Move`, `MoveAilment`, `MoveBattleStyle`, `MoveCategory`, `MoveDamageClass`, `MoveLearnMethod`, `MoveTarget`
*   **Move Ailments**: `GET https://pokeapi.co/api/v2/move-ailment/{id or name}/`
*   **Move Battle Styles**: `GET https://pokeapi.co/api/v2/move-battle-style/{id or name}/`
*   **Move Categories**: `GET https://pokeapi.co/api/v2/move-category/{id or name}/`
*   **Move Damage Classes**: `GET https://pokeapi.co/api/v2/move-damage-class/{id or name}/`
*   **Move Learn Methods**: `GET https://pokeapi.co/api/v2/move-learn-method/{id or name}/`
*   **Move Targets**: `GET https://pokeapi.co/api/v2/move-target/{id or name}/`

### 10. Pokémon
*   **Endpoint**: `GET https://pokeapi.co/api/v2/pokemon/{id or name}/`
    *   **Objects**: `Pokemon`, `Ability`, `Characteristic`, `EggGroup`, `Gender`, `GrowthRate`, `Nature`, `PokeathlonStat`, `PokemonLocationArea`, `PokemonColor`, `PokemonForm`, `PokemonHabitat`, `PokemonShape`, `PokemonSpecies`, `Stat`, `Type`
*   **Abilities**: `GET https://pokeapi.co/api/v2/ability/{id or name}/`
*   **Characteristics**: `GET https://pokeapi.co/api/v2/characteristic/{id}/`
*   **Egg Groups**: `GET https://pokeapi.co/api/v2/egg-group/{id or name}/`
*   **Genders**: `GET https://pokeapi.co/api/v2/gender/{id or name}/`
*   **Growth Rates**: `GET https://pokeapi.co/api/v2/growth-rate/{id or name}/`
*   **Natures**: `GET https://pokeapi.co/api/v2/nature/{id or name}/`
*   **Pokeathlon Stats**: `GET https://pokeapi.co/api/v2/pokeathlon-stat/{id or name}/`
*   **Pokemon Location Areas**: `GET https://pokeapi.co/api/v2/location-area/{id or name}/pokemon/`
*   **Pokemon Colors**: `GET https://pokeapi.co/api/v2/pokemon-color/{id or name}/`
*   **Pokemon Forms**: `GET https://pokeapi.co/api/v2/pokemon-form/{id or name}/`
*   **Pokemon Habitats**: `GET https://pokeapi.co/api/v2/pokemon-habitat/{id or name}/`
*   **Pokemon Shapes**: `GET https://pokeapi.co/api/v2/pokemon-shape/{id or name}/`
*   **Pokemon Species**: `GET https://pokeapi.co/api/v2/pokemon-species/{id or name}/`
*   **Stats**: `GET https://pokeapi.co/api/v2/stat/{id or name}/`
*   **Types**: `GET https://pokeapi.co/api/v2/type/{id or name}/`

### 11. Utility
*   **Endpoint**: `GET https://pokeapi.co/api/v2/language/{id or name}/`
    *   **Objects**: `Language`
