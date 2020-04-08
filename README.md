# Score tracker for discgolf and golf

## TODO
- PWA
- HTML Router

## Create new course

### POST
```json
{
	"basketCount": 2,
	"players": ["Tiger King", "Ying Yang"]
}
```

### Server responses
```json
{
  "id": "2ty",
  "basketCount": 2,
  "active": 1,
  "Baskets": {
    "1": {
      "OrderNum": 1,
      "Par": 0,
      "Scores": {
        "Tiger King": {
          "Score": 0,
          "OB": 0
        },
        "Ying Yang": {
          "Score": 0,
          "OB": 0
        }
      }
    },
    "2": {
      "OrderNum": 2,
      "Par": 0,
      "Scores": {
        "Tiger King": {
          "Score": 0,
          "OB": 0
        },
        "Ying Yang": {
          "Score": 0,
          "OB": 0
        }
      }
    }
  }
}
```
