import requests as re
import bs4
import argparse
import sys

parser = argparse.ArgumentParser(
    prog="parse",
    description="Скачивает ссылки на статьи с сайта menunedeli.ru"
)
parser.add_argument("-c", "--count", type=int, default=1000, help="Количество ссылок для скачивания")
parser.add_argument("-s", "--save", type=str, default="links.txt", help="Файл, куда сохранять ссылки")

args = parser.parse_args()

saveTo = args.save
UpToLinks = args.count
catalogFormat = "https://menunedeli.ru/novye-stati/page/{}"
with open(saveTo, "w") as f:
    links = 0
    page = 1
    while True:
        catalogPage = re.get(catalogFormat.format(page))

        bs = bs4.BeautifulSoup(catalogPage.content, "lxml")

        for a in bs.find_all('article'):
            for link in a.find_all('meta', attrs={'itemprop':'url'}):
                if link['content'].startswith("https://menunedeli.ru/recipe"):
                    print(link['content'], file=f)
                    links += 1
                    if links >= UpToLinks:
                        exit(0)
        page += 1
    