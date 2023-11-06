run:
	./data.sh
	go run main.go queue.go resize.go
	python3 ./plot.py
	latexmk -pdf -pdflatex="pdflatex -interaction=nonstopmode" -use-make ./report.tex
	rm -f ./report.aux ./report.fdb_latexmk ./report.fls ./report.log ./report.out ./report.synctex.gz

clean:
	rm -f *.aux *.fdb_latexmk *.fls *.log *.out *.synctex.gz
	rm -rf ./figure ./latency ./throughput ./output
	rm -rf ./tiny-imagenet-200