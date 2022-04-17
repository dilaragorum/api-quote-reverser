### Quote Reverser (s)

Consume a quotes api for fetching quotes, group all sentences by author and reverse all the sentences.

Quotes API: [https://type.fit/api/quotes](https://type.fit/api/quotes)

**Given**

```json
[
  {
    "text": "foo",
    "author": "Thomas Edison"
  },
  {
    "text": "bar",
    "author": "Lao Tzu"
  },
  {
    "text": "baz",
    "author": "Thomas Edison"
  }
]
```

**Expected**

```json
[
  {
    "author": "Thomas Edison",
    "quotes": ["oof", "zab"]
  },
  {
    "author": "Lao Tzu",
    "quotes": ["rab"]
  }
]
```