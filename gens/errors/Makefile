bindata:
	@cd ../../template; go-bindata -o ../gens/errors/bindata.go  ./...; \
	sed -i '' 's/package main/package generrors/g' ../gens/errors/bindata.go;
