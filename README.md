This is intended to be a markdown parser that I can use to pass a string to and get a nice, html, highlighted file. WIP

#### Example request

````bash
 curl -X POST \
 'https://v8xtuqb7a2.execute-api.us-east-2.amazonaws.com/' \
 --header 'Accept: */*' \
 --header 'Content-Type: text/plain' \
 --data-raw '
 ```typescript
const loremFetcher = async (id: string) => {
 const ipsum = await fetch(`.../${id}`);
 const text = await ipsum.text();
 return text;
};'
````
