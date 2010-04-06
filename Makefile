all: fuzz-client fuzz-client-arm
	adb push fuzz-client-arm /data/local
fuzz-client: fuzz.8
	8l -o fuzz-client fuzz.8
fuzz-client-arm: fuzz.5
	5l -o fuzz-client-arm fuzz.5

fuzz.5: fuzz.go 
	5g fuzz.go
fuzz.8: fuzz.go
	8g fuzz.go

clean:
	rm -f *.5 *.8 fuzz-client fuzz-client-arm


