run:
	mkdir streams/`date '+%Y%m%d%H_%M_%S'`
	go run main.go | ffmpeg -f rawvideo -pixel_format bgr24 -framerate 2 -video_size 1280x720 -i -  -vcodec libx264 -acodec copy -pix_fmt yuv420p -color_range 2 -g 50 -keyint_min 50 -sc_threshold 0 -segment_time 1 -hls_list_size 0 -use_localtime 1 -hls_segment_filename streams/`date '+%Y%m%d%H_%M_%S'`/%s.ts streams/`date '+%Y%m%d%H_%M_%S/playlist.m3u8'`
