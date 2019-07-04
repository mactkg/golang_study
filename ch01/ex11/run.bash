cd $(dirname $0) && pwd

# To generate url list, Run script bellow at https://www.alexa.com/topsites/countries/JP
# Array.from(document.querySelectorAll('.DescriptionCell > p > a')).map(a => a.getAttribute('href')).map(h => `https://${h.split('/')[2]}`).join("\n")
cat urls | xargs go run main.go 
