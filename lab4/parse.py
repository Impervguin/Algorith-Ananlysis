import requests as re
import bs4

saveTo = "links.txt"
UpToLinks = 1000
catalogFormat = "https://menunedeli.ru/novye-stati/page/{}"
with open(saveTo, "w") as f:
    links = 0
    page = 1
    while True:
        catalogPage = re.get(catalogFormat.format(page))

        bs = bs4.BeautifulSoup(catalogPage.content)

        for a in bs.find_all('article'):
            for link in a.find_all('meta', attrs={'itemprop':'url'}):
                print(link['content'], file=f)
                links += 1
                if links > UpToLinks:
                    exit(0)
        page += 1
    