import requests
import lxml.html as lh
import pandas as pd

res = {}
for i in range(1,26):
    url = 'https://codeforces.com/contest/1118/standings/page/' + str(i)
    page = requests.get(url)
    #Create a handle, page, to handle the contents of the website
    # page = requests.get(url)
    #Store the contents of the website under doc
    doc = lh.fromstring(page.content)
    #Parse data that are stored between <tr>..</tr> of HTML
    tr_elements = doc.xpath('//tr')


    #For each row, store each first element (header) and an empty list
    for i in range(len(tr_elements)):
        tmp = []
        for t in tr_elements[i][:2]:
            name=t.text_content().strip()
            tmp.append(name)
        if len(tmp) >= 2:
            res[tmp[1]] = tmp[0]


file = open('/Users/lisanaaa/Downloads/contest.txt', 'w')
for k,v in res.items():
    tmp = k + ' ' + v + '\n'
    file.write(tmp)
file.close()





