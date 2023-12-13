
# go-stream-mp3

This fork contains changes that allow the decoder to be used with streaming data. In particular, the decoder no longer tries to read all the data from the file to determine its size.

### Original readme:

># This project is no longer maintained.
>
># go-mp3
>
>[![Go Reference](https://pkg.go.dev/badge/github.com/hajimehoshi/go-mp3.svg)](https://pkg.go.dev/github.com/hajimehoshi/go-mp3) 
>
>An MP3 decoder in pure Go based on [PDMP3](https://github.com/technosaurus/PDMP3).
>
>[Slide at golang.tokyo #11](https://docs.google.com/presentation/d/e/2PACX-1vTTXf-LWNRvMVGQ7GI4Wh8EKohot_9CMtlF4dswpYGpuYKOek5NeNP-_QZnNcRFZp9Cwm0pCcykjqDN/pub?>start=false&loop=false&delayms=3000)
