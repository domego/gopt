bindata:
	@cd ../../template; go-bindata -o ../gens/types/bindata.go  ./...; \
	sed -i '' 's/package main/package gentypes/g' ../gens/types/bindata.go;
