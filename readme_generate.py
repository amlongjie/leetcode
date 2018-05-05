# encoding:utf-8
import json

def get_difficulty(num):
    if num == 1:
        return "Easy"
    if num == 2:
        return "Medium"
    if num == 3:
        return "Hard"
    return "Unknown"

data = json.load(open("current.json"))
tmp_data = [{'id': x['stat']['frontend_question_id'], 'title': x['stat']['question__title'], 'status': x['status'], 'difficulty':x['difficulty']['level']}
            for x in data['stat_status_pairs']]

tmp_data.sort(key=lambda k: k['id'])

f = open('./README.md', 'w')

f.write("# LeetCode 刷题记录    ")

f.write("![](http://progressed.io/bar/%s?title=completed)\n" % int(data['num_solved'] / data['num_total']))

f.write("AC=%s Total=%s \n\n" % (data['num_solved'], data['num_total']))

f.write("id  | title | difficulty |status\n")
f.write("-----  | ----- | ----- | -----\n")
for tmp in tmp_data:
    f.write("%s|%s|%s|%s\n" % (tmp['id'], tmp['title'],
    get_difficulty(int(tmp['difficulty'])),
    ":white_check_mark:" if tmp['status'] == "ac" else ":x:"))
f.close()




