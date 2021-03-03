import re
import time

import requests
from sqlalchemy import Column, Integer, String, Time
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

base_url = "https://gateway.36kr.com/api/mis/nav/ifm/subNav/flow"
headers = {
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36",
    "content-type": "application/json"
}
initPageCallback = "eyJmaXJzdElkIjozMjMwNzM4LCJsYXN0SWQiOjMyMzAxNjAsImZpcnN0Q3JlYXRlVGltZSI6MTYxNDc1NzMzNjY0NSwibGFzdENyZWF0ZVRpbWUiOjE2MTQ3MzYxNDEwMTV9"
body = {
    "partner_id": "web",
    "timestamp": time.time(),
    "param": {
        "pageCallback": initPageCallback,
        "pageEvent": 1,
        "pageSize": 30,
        "platformId": 2,
        "siteId": 1,
        "subnavNick": "technology",
        "subnavType": 1
    }
}
Base = declarative_base()


class Technology(Base):
    __tablename__ = 'technologies'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    summary = Column(String)
    image = Column(String)
    image_type = Column(Integer)
    context = Column(String)
    user_id = Column(String)
    create_at = Column(Time)
    update_at = Column(Time)

    def __repr__(self):
        return "<Technology(id='%s',name='%s', summary='%s', image='%s',image_type='%s',context='%s', user_id='%s', create_at='%s',update_at='%s')>" % (
            self.id, self.name, self.summary, self.image, self.image_type, self.context, self.user_id, self.create_at,
            self.update_at)


def getList(body):
    results = requests.post(base_url, headers=headers, json=body).json()
    newPageCallback = results["data"]["pageCallback"]
    list = results["data"]["itemList"]
    print(newPageCallback)
    return list, newPageCallback


def insertTech(engine, tech):
    Session = sessionmaker(bind=engine)
    session = Session()
    session.add(tech)
    session.commit()
    # try:
    #
    # except:
    #     # 发生错误时回滚
    #     print("have err")
    #     db.rollback()


if __name__ == '__main__':
    engine = create_engine("mysql+pymysql://root:leeyfMysql100%@106.15.198.212:3306/tech_platform", max_overflow=5)
    list, newPageCallback = getList(body)
    for i in list:
        print(i)
    body["param"]["pageCallback"] = newPageCallback
    j = 0
    t = 0
    while j < 10:
        j += 1
        list, newPageCallback = getList(body)
        body["param"]["pageCallback"] = newPageCallback
        print(newPageCallback)
        for i in list:
            t += 1
            print(t)
            item = i["templateMaterial"]
            name = item["widgetTitle"]
            summary = item["summary"]
            create_at = item["publishTime"]
            image = item["widgetImage"]
            context_url = "https://www.36kr.com/p/" + str(item["itemId"])
            text = requests.get(context_url).text
            # text = text.replace("\\", "")
            list = re.compile("<p>(.*?)</p>").findall(text)
            context_text = ""
            for l in list:
                context_text += '<p>' + l + '</p>\n'
            context = context_text
            user_id = "371310968510091264"
            ed_tech = Technology(name=name,summary=summary,image=image,image_type=1,context=context,user_id=user_id,create_at=time.localtime(create_at/1000),update_at=time.localtime())
            insertTech(engine, ed_tech)
