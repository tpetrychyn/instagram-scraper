# instagram-scraper
A golang script to grab an instagram accounts latest post shortcodes

## How to Use
1. Enter an instagram profile name in the profileName const in `main.go`
2. `go build . && ./instagram-scraper`
3. On success, the script will output 12 shortcodes from the Instagram profile
4. The image can be accessed at https://www.instagram.com/p/{shortCode}/media?size=l where size is `t, s, m, l`

#### For more see https://www.instagram.com/developer/embedding/
