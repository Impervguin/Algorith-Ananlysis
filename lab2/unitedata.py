import csv
import os
import re

OUT_FILE = 'data.tsv'

files = list(filter(lambda x: re.match(".*data[0-9]*.tsv", x) != None, os.listdir("./")))
lines = []
for file in files:
    with open(f"./{file}", "r", newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f, delimiter="\t")
        for row in reader:
            lines.append(row)
            row['size'] = int(row['size'])
    print(f"Read {file} successfully.")

lines.sort(key=lambda x: x['size'])

with open(OUT_FILE, "w", newline="", encoding="utf-8") as f:
    writer = csv.DictWriter(f, fieldnames=lines[0].keys(), delimiter="\t")
    writer.writeheader()
    writer.writerows(lines)

print(f"Sorted data saved to {OUT_FILE}.")
