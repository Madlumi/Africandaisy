.PHONY: res
run: res
	go run src/main.go
res:
	ffmpeg -y -i res/cubeAnim/cube%04d.png -filter_complex tile=16x1 res/cube.png	
