bindata:
	@cd ../../template; go-bindata -o ../gens/gincontroller/bindata.go  ./...; \
	sed -i '' 's/package main/package gengincontroller/g' ../gens/gincontroller/bindata.go;
