import os, tables, unicode

let filename = if paramCount() > 0: paramStr(1) else: "page1.txt"

var cf = initCountTable[string]()

try:
  for line in lines filename:
    for c in runes(line):
      cf.inc(c.toUTF8())
except IOError:
  echo "IOError:" & getCurrentExceptionMsg()
  
cf.sort()

for k, v in cf.pairs():
  echo $k,": ",$v
  