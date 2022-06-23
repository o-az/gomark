This is intended to be a markdown parser that I can use to pass a string to and get a nice, html, highlighted file. WIP

### Accepts `GET` or `POST` requests.

---

## How to use

#### `GET` request requires a `text` query parameter encoded as base64. For example:

```typescript
const loremFetcher = async (id: string) => {
  const ipsum = await fetch(`.../${id}`);
  const text = await ipsum.text();
  return text;
};
```

Becomes: `YGBgdHlwZXNjcmlwdApjb25zdCBsb3JlbUZldGNoZXIgPSBhc3luYyAoaWQ6IHN0cmluZykgPT4gewogIGNvbnN0IGlwc3VtID0gYXdhaXQgZmV0Y2goYC4uLi8ke2lkfWApOwogIGNvbnN0IHRleHQgPSBhd2FpdCBpcHN1bS50ZXh0KCk7CiAgcmV0dXJuIHRleHQ7Cn07CmBgYA==`

And the request:

```bash
curl -X GET -G 'http://...?text=YGBgdHlwZXNjcmlwdApjb25zdCBsb3JlbUZldGNoZXIgPSBhc3luYyAoaWQ6IHN0cmluZykgPT4gewogIGNvbnN0IGlwc3VtID0gYXdhaXQgZmV0Y2goYC4uLi8ke2lkfWApOwogIGNvbnN0IHRleHQgPSBhd2FpdCBpcHN1bS50ZXh0KCk7CiAgcmV0dXJuIHRleHQ7Cn07CmBgYA==`
```

Returns [this html text](https://nopaste.ml/?l=html#XQAAAQCFAwAAAAAAAAAeCEUG0O+oKBdZ2an16qclQGevk+oY5yGa1a7rJq+VXcNSx8JGRDxce6oszg6o/60j2zhqpMctUsvWsKoJtlr26cZWkTYVw2XQPa3U4eBxI3uJ6NgJH4xMXNO4auXb/PXzgrzlTrUHjghDJexfYioWEJf4eW45ViEXOT944ObXWQoqhAC8Fhuqy9w/SD/rM69Jqp9rg6qwKNueJEoJkSK+0X8BonzveHytgLtBtkeLCntpebXRlq6dwI/5C7H57ptkIeDoeVTIL9kJdkIZMrATLunNrVLi78aneCzSC6XlNQQE8WiVCTBpxJA2ATm4M1WY9BK5YZONsYkXxtYsrPmmulPF7peUmsVAZ0FsoQerqu3yauF30gjkNfIKo5An/gIfHuaDU79lYO3xjsjB0HTWRODZVWptwBQ8Xq0jqKJh/yRV4J+GdaWfnsbbhRN8CT3VZ09cpDN9r1qM2BRqORWwt9khyeU5NxPUrgV72YNmQb0/KjYoAdJUbZvrbBNxMrDvCS4AdS39PXX96syHaLpdjLtM/iH8K/8f6dkp7CMsOlnwFyI9kKHfpTVF1P1EWn/Cj115CBvBaT92YOXN4Dugi01vd5ZE200p9PKvuCqvpA5kA8hOX+40Orh8IcFvJV1fQHPYoFewmcTipYslJkZbmPA/Wuh4c4n7m+rI), hich would look like this:
![](https://cdn.discordapp.com/attachments/876884946994741329/989539821473103902/unknown.png)

`POST` request gets more options. Request payload (`application/json`):

````json
{
  "content": "# Testing md to html with snytax highlighting:\n```\nconst loremFetcher = async (id: string) => {\n  const ipsum = await fetch(`.../${id}`);\n  const text = await ipsum.text();\n  return text;\n};\n```",
  "options": {
    "return_type": "html", // or "text"
    "font": true, // default: true
    "syntax_highlight": true // default: true
    // WIP: "extensions": ["gfm", "strikethrough", "typographer", ...]
  }
}
````

Returns:

### WIP
