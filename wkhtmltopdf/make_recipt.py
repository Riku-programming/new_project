import os
import pdfkit

# wkhtmltopdfのインストールパスを指定(要書換★) --- (*1)
wkhtmltopdf = '/usr/local/bin/wkhtmltopdf'

# ファイルのパスを指定
root = os.path.dirname(__file__)
if root == '': root = '.'
template_html = root + '/template.html'
temp_file = root + '/__tmp.html'
output_file = root + '/output.pdf'

# ユーザーに質問する --- (*2)
import tkinter as tk
import tkinter.simpledialog as sd

tk.Tk().withdraw()  # ダイアログを使う場合
name = sd.askstring("名前", "名前を入力してください")
price = sd.askstring("値段", "値段を入力してください")

# テンプレートを読む --- (*3)
with open(template_html, 'rt', encoding='utf-8') as fp:
    text = fp.read()
# 日付や名前、金額を置換する --- (*4)
from datetime import datetime

text = text.replace('__DATE__',
                    datetime.now().strftime('%Y年%m月%d日'))
text = text.replace('__NAME__', name)
text = text.replace('__PRICE__', price)
# 一時ファイルを出力 --- (*5)
with open(temp_file, 'wt', encoding='utf-8') as fp:
    fp.write(text)

# 変換オプションなどを指定
conf = pdfkit.configuration(wkhtmltopdf=wkhtmltopdf)
options = {'page-size': 'A4', 'encoding': "UTF-8"}

# HTML/CSSファイルをPDF出力 --- (*6)
pdfkit.from_file(temp_file, output_file, options=options, configuration=conf)
print("ok")
