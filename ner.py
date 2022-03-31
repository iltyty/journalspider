import hanlp
import xml.etree.cElementTree as et


RES_DIR = "./data/"
PEOPLE_DAILY = RES_DIR + "people_daily.xml"
HUNAN_DAILY = RES_DIR + "hunan_daily.xml"
JF_DAILY = RES_DIR + "jiefang_daily.xml"
BJ_NEWS  = RES_DIR + "xinjing.xml"
BJ_YOUTH = RES_DIR + "qingnian.xml"


class News:
    def __init__(self, title, content) -> None:
        self.title = title
        self.content = content


def parse_xml(path: str) -> list[News]:
    news_list = []
    tree = et.ElementTree(file=path)
    root = tree.getroot()

    for news in root:
        title = news.attrib['title']
        content = news.find('content').text
        news_list.append(News(title, content))
    return news_list


def word_sep(news_list: list[News]):
    tok = hanlp.load(hanlp.pretrained.tok.COARSE_ELECTRA_SMALL_ZH)
    # tok = hanlp.load(hanlp.pretrained.tok.PKU_NAME_MERGED_SIX_MONTHS_CONVSEG)
    ner = hanlp.load(hanlp.pretrained.ner.MSRA_NER_ELECTRA_SMALL_ZH)
    # ner = hanlp.load(hanlp.pretrained.ner.MSRA_NER_BERT_BASE_ZH)
    
    HanLP = hanlp.pipeline() \
        .append(hanlp.utils.rules.split_sentence) \
        .append(tok) \
        .append(lambda sents: sum(sents, []))
    
    for news in news_list:
        sep_res = list(set(HanLP(news.content)))
        ner_res = ner(sep_res)
        for r in ner_res:
            if r[1] in ['PERSON', 'LOCATION', 'ORGANIZATION']:
                print(r[0], r[1])
        # print([x[0] for x in ner(sep_res)])


def main():
    news_list = parse_xml(PEOPLE_DAILY)
    word_sep(news_list[:1])

if __name__ == '__main__':
    main()
