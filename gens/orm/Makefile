bindata:
	@cd ../../template; go-bindata -o ../gens/orm/bindata.go  ./...; \
	sed -i '' 's/package main/package genorm/g' ../gens/orm/bindata.go;
