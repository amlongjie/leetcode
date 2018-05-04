# encoding:utf-8
import json

data = json.load(open("current.json"))
tmp_data = [{'id': x['stat']['frontend_question_id'], 'title': x['stat']['question__title'], 'status': x['status']}
            for x in data['stat_status_pairs']]

tmp_data.sort(key=lambda k: k['id'])

f = open('./README.md', 'w')

f.write("id  | title | status\n")
f.write("-----  | ----- | -----\n")
for tmp in tmp_data:
    f.write("%s|%s|%s\n" % (tmp['id'], tmp['title'], ":white_check_mark:" if tmp['status'] == "ac" else ":x:"))
f.close()
