# gocv-hls

## 起動方法

```go run main.go | ffmpeg -f rawvideo -pixel_format bgr24 -framerate 2 -video_size 1280x720 -i - foo.mp4 -vcodec libx264 -acodec copy -pix_fmt yuv420p -color_range 2 -hls_time 1 -hls_list_size 5 -use_localtime 1 -hls_segment_filename '%Y%m%d-%s.ts' ./playlist.m3u8```
