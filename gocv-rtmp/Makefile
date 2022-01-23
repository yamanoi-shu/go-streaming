run:
	go run main.go | ffmpeg  \
		-video_size 1280x720 \
		-y \
    -f rawvideo \
    -vcodec rawvideo \
    -pix_fmt bgr24 \
    -i - \
    -c:v libx264 \
    -pix_fmt yuv420p \
    -preset ultrafast \
    -f flv \
		{rtmp url}
