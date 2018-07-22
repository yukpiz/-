#/bin/bash

#pandoc -s -t revealjs -o ./outputs/vimtutor.html vimtutor.md --slide-level=2
#pandoc -s -t revealjs -V theme:black -V width:1920 -V margin:0 -o ./outputs/example.html example.md --slide-level=1 --highlight-style=espresso

OUTPUT_DIR=./outputs/
THEME=black
WIDTH=1920
HEIGHT=1280
MARGIN=0
HIGHLIGHT_STYLE=espresso
SLIDE_LEVEL=1
HISTORY=true

echo "===> GENERATE SLIDES SCRIPT ===>"
echo "  Theme:           ${THEME}"
echo "  Width:           ${WIDTH}"
echo "  Margin:          ${MARGIN}"
echo "  Highlight Style: ${HIGHLIGHT_STYLE}"
echo "  Slide Level:     ${SLIDE_LEVEL}"
echo "  History:         ${HISTORY}"
echo ""

INDEX=1
for FILE in ${DIRPATH}*.md; do
  echo "=== FILE INDEX ${INDEX} =========="
  echo "Improt File: ${FILE}"
  echo "Export File: ${FILE%.*}.html"
  echo ""
  let INDEX++
  pandoc -s -t revealjs -V padding:0 -V theme:${THEME} -V width:${WIDTH} -V height:${HEIGHT} -V margin:${MARGIN} -V history:${HISTORY} -o ${OUTPUT_DIR}${FILE%.*}.html ${FILE} --slide-level=${SLIDE_LEVEL} --highlight-style=${HIGHLIGHT_STYLE}
done