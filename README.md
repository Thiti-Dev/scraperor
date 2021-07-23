[![CircleCI](https://circleci.com/gh/Thiti-Dev/scraperor/tree/master.svg?style=shield)](https://circleci.com/gh/Thiti-Dev/scraperor/tree/master)

## 🎓 Scraperor [scraping service]

A free services that helps you in seeking any information/data of any website in the world

## 💡 Usage
##### GRAPHQL
```url
https://scraperor.herokuapp.com/query
```
###### Example
```gql
mutation{
  scrapeNow(input:{
    pointer: ".user-profile-bio div",
    target_url: "https://github.com/Thiti-Dev"
  }){
    elements
  }
}
```

##### API-ENDPOINT
```url
{{POST}}: https://scraperor.herokuapp.com/scrape
```
###### Example body
```json
{
    "target_url": "https://github.com/Thiti-Dev",
    "pointer": ".user-profile-bio div"
}
```

## 📕 Important notes
- This is an alpha version so you can use our scraping engine without any registration required
- In the official version, you will have to make a registration before using our service and we'll collect all of the data ex. transmission data from host -> target , collected elements, etc . . .
