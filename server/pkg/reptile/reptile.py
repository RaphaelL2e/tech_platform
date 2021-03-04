import re
import time
import random
import requests
from sqlalchemy import Column, Integer, String, Time
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

base_url = "https://gateway.36kr.com/api/mis/nav/ifm/subNav/flow"
headers = {
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36",
    "content-type": "application/json"
}
initPageCallback = "eyJmaXJzdElkIjozMjMyMDY0LCJsYXN0SWQiOjMyMzEzNzgsImZpcnN0Q3JlYXRlVGltZSI6MTYxNDgzOTkwMDI4NywibGFzdENyZWF0ZVRpbWUiOjE2MTQ4MTU5ODA4MjV9"
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


class Article(Base):
    __tablename__ = 'articles'
    id = Column(Integer, primary_key=True)
    title = Column(String)
    summary = Column(String)
    image = Column(String)
    user_id = Column(String)
    author = Column(String)
    context = Column(String)
    status = Column(Integer)
    create_at = Column(Time)
    update_at = Column(Time)

    def __repr__(self):
        return "<Article(id='%s',title='%s', summary='%s', image='%s',user_id='%s', author='%s', context='%s', create_at='%s',update_at='%s')>" % (
            self.id, self.title, self.summary, self.image, self.user_id, self.author, self.context, self.create_at,
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
    while j < 50:
        j += 1
        list, newPageCallback = getList(body)
        body["param"]["pageCallback"] = newPageCallback
        print(newPageCallback)
        for i in list:
            t += 1
            print(t)
            item = i["templateMaterial"]
            title = item["widgetTitle"]
            summary = item["summary"]
            create_at = item["publishTime"]
            image = item["widgetImage"]
            context_url = "https://www.36kr.com/p/" + str(item["itemId"])
            author = item["authorName"]
            text = requests.get(context_url).text
            # text = text.replace("\\", "")
            list = re.compile("<p>(.*?)</p>").findall(text)
            context_text = ""
            for l in list:
                context_text += '<p>' + l + '</p>\n'
            context = context_text
            user_id = "371310968510091264"
            ed_tech = Article(title=title, summary=summary, image=image, context=context,
                              user_id=user_id, author=author, status=random.randint(-1,1), create_at=time.localtime(create_at / 1000),
                              update_at=time.localtime())
            insertTech(engine, ed_tech)
