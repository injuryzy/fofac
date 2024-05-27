package excel

import "testing"

var (
	r = map[string][][]string{"test": {
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
		[]string{"www.9uyy.", " 154.214.13.107", "80", "nginx", "9uyy.com", "悠久影院 - 热播高清全集电视剧 - 最新电影抢先看 - 免费看电影网", "CN HK Hong Kong 桂ICP备18010014号", "http", "111", "2222", "33333"},
	}}
)

func TestName(t *testing.T) {
	WriteXlsx(r, "data.xlsx")
}
