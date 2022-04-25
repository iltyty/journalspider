import xml.dom.minidom

import hanlp
import xml.etree.cElementTree as et


RES_DIR = "./data/"
RES_FILE = RES_DIR + 'news.xml'


class News:
    def __init__(self, title, content) -> None:
        self.title = title
        self.content = content
        self.entities = []


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
        if not news.content:
            continue
        sep_res = list(set(HanLP(news.content)))
        ner_res = ner(sep_res)
        for r in ner_res:
            if r[1] in ['PERSON', 'LOCATION', 'ORGANIZATION']:
                news.entities.append((r[0], r[1]))


def save_res(news_list: list[News], file="./data/ner.xml"):
    doc = xml.dom.minidom.Document()
    root = doc.createElement("data")
    doc.appendChild(root)

    for news in news_list:
        news_node = doc.createElement('news')
        news_node.setAttribute('title', news.title)

        for entity in news.entities:
            entity_node = doc.createElement('entity')
            entity_node.setAttribute('type', entity[1])
            entity_node.appendChild(doc.createTextNode(entity[0]))
            news_node.appendChild(entity_node)

        root.appendChild(news_node)

    with open(file, 'w', encoding='utf-8') as f:
        doc.writexml(f, addindent='\t', newl='\n', encoding='utf-8')


def main():
    news_list = parse_xml(RES_FILE)
    word_sep(news_list)
    save_res(news_list)


if __name__ == '__main__':
    main()
