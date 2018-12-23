// This file is automatically generated by qtc from "article.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line article.qtpl:1
package templates

//line article.qtpl:1
import "fmt"

//line article.qtpl:2
import "strconv"

//line article.qtpl:3
import "meguca/common"

//line article.qtpl:4
import "meguca/lang"

//line article.qtpl:5
import "meguca/imager/assets"

//line article.qtpl:6
import "meguca/util"

//line article.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line article.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line article.qtpl:8
func streamrenderArticle(qw422016 *qt422016.Writer, p common.Post, c articleContext) {
	//line article.qtpl:9
	id := strconv.FormatUint(p.ID, 10)

	//line article.qtpl:10
	ln := lang.Get()

	//line article.qtpl:10
	qw422016.N().S(`<article id="p`)
	//line article.qtpl:11
	qw422016.N().S(id)
	//line article.qtpl:11
	qw422016.N().S(`"`)
	//line article.qtpl:11
	qw422016.N().S(` `)
	//line article.qtpl:11
	streampostClass(qw422016, p, c.op)
	//line article.qtpl:11
	qw422016.N().S(`>`)
	//line article.qtpl:12
	streamdeletedToggle(qw422016)
	//line article.qtpl:12
	qw422016.N().S(`<header class="spaced"><input type="checkbox" class="mod-checkbox hidden">`)
	//line article.qtpl:15
	streamrenderSticky(qw422016, c.sticky)
	//line article.qtpl:16
	streamrenderLocked(qw422016, c.locked)
	//line article.qtpl:17
	if c.subject != "" {
		//line article.qtpl:18
		if c.board != "" {
			//line article.qtpl:18
			qw422016.N().S(`<b class="board">/`)
			//line article.qtpl:20
			qw422016.N().S(c.board)
			//line article.qtpl:20
			qw422016.N().S(`/</b>`)
			//line article.qtpl:22
		}
		//line article.qtpl:22
		qw422016.N().S(`<h3>「`)
		//line article.qtpl:24
		qw422016.E().S(c.subject)
		//line article.qtpl:24
		qw422016.N().S(`」</h3>`)
		//line article.qtpl:26
	}
	//line article.qtpl:26
	qw422016.N().S(`<b class="name spaced`)
	//line article.qtpl:27
	if p.Auth != "" {
		//line article.qtpl:27
		qw422016.N().S(` `)
		//line article.qtpl:27
		qw422016.N().S(`admin`)
		//line article.qtpl:27
	}
	//line article.qtpl:27
	if p.Sage {
		//line article.qtpl:27
		qw422016.N().S(` `)
		//line article.qtpl:27
		qw422016.N().S(`sage`)
		//line article.qtpl:27
	}
	//line article.qtpl:27
	qw422016.N().S(`">`)
	//line article.qtpl:28
	if p.Name != "" || p.Trip == "" {
		//line article.qtpl:28
		qw422016.N().S(`<span>`)
		//line article.qtpl:30
		if p.Name != "" {
			//line article.qtpl:31
			qw422016.E().S(p.Name)
			//line article.qtpl:32
		} else {
			//line article.qtpl:33
			qw422016.N().S(ln.Common.Posts["anon"])
			//line article.qtpl:34
		}
		//line article.qtpl:34
		qw422016.N().S(`</span>`)
		//line article.qtpl:36
	}
	//line article.qtpl:37
	if p.Trip != "" {
		//line article.qtpl:37
		qw422016.N().S(`<code>!`)
		//line article.qtpl:39
		qw422016.E().S(p.Trip)
		//line article.qtpl:39
		qw422016.N().S(`</code>`)
		//line article.qtpl:41
	}
	//line article.qtpl:42
	if p.PosterID != "" {
		//line article.qtpl:42
		qw422016.N().S(`<span>(`)
		//line article.qtpl:44
		qw422016.E().S(p.PosterID)
		//line article.qtpl:44
		qw422016.N().S(`)</span>`)
		//line article.qtpl:46
	}
	//line article.qtpl:47
	if p.Auth != "" {
		//line article.qtpl:47
		qw422016.N().S(`<span>##`)
		//line article.qtpl:49
		qw422016.N().S(` `)
		//line article.qtpl:49
		qw422016.N().S(ln.Common.Posts[p.Auth])
		//line article.qtpl:49
		qw422016.N().S(`</span>`)
		//line article.qtpl:51
	}
	//line article.qtpl:51
	qw422016.N().S(`</b>`)
	//line article.qtpl:53
	if p.Flag != "" {
		//line article.qtpl:54
		title, ok := countryMap[p.Flag]

		//line article.qtpl:55
		if !ok {
			//line article.qtpl:56
			title = p.Flag

			//line article.qtpl:57
		}
		//line article.qtpl:57
		qw422016.N().S(`<img class="flag" src="/assets/flags/`)
		//line article.qtpl:58
		qw422016.N().S(p.Flag)
		//line article.qtpl:58
		qw422016.N().S(`.svg" title="`)
		//line article.qtpl:58
		qw422016.N().S(title)
		//line article.qtpl:58
		qw422016.N().S(`">`)
		//line article.qtpl:59
	}
	//line article.qtpl:59
	qw422016.N().S(`<time>`)
	//line article.qtpl:61
	qw422016.N().S(formatTime(p.Time))
	//line article.qtpl:61
	qw422016.N().S(`</time><nav>`)
	//line article.qtpl:64
	url := "#p" + id

	//line article.qtpl:65
	if c.index {
		//line article.qtpl:66
		url = util.ConcatStrings("/all/", id, "?last=100", url)

		//line article.qtpl:67
	}
	//line article.qtpl:67
	qw422016.N().S(`<a href="`)
	//line article.qtpl:68
	qw422016.N().S(url)
	//line article.qtpl:68
	qw422016.N().S(`">No.</a><a class="quote" href="`)
	//line article.qtpl:71
	qw422016.N().S(url)
	//line article.qtpl:71
	qw422016.N().S(`">`)
	//line article.qtpl:72
	qw422016.N().S(id)
	//line article.qtpl:72
	qw422016.N().S(`</a></nav>`)
	//line article.qtpl:75
	if c.index && c.subject != "" {
		//line article.qtpl:75
		qw422016.N().S(`<span>`)
		//line article.qtpl:77
		streamexpandLink(qw422016, "all", id)
		//line article.qtpl:78
		streamlast100Link(qw422016, "all", id)
		//line article.qtpl:78
		qw422016.N().S(`</span>`)
		//line article.qtpl:80
	}
	//line article.qtpl:81
	streamcontrolLink(qw422016)
	//line article.qtpl:82
	if c.op == p.ID {
		//line article.qtpl:83
		streamthreadWatcherToggle(qw422016, p.ID)
		//line article.qtpl:84
	}
	//line article.qtpl:84
	qw422016.N().S(`</header>`)
	//line article.qtpl:86
	var src string

	//line article.qtpl:87
	if p.Image != nil {
		//line article.qtpl:88
		img := *p.Image

		//line article.qtpl:89
		src = assets.SourcePath(img.FileType, img.SHA1)

		//line article.qtpl:89
		qw422016.N().S(`<figcaption class="spaced"><a class="image-toggle act" hidden></a><span class="spaced image-search-container">`)
		//line article.qtpl:93
		streamimageSearch(qw422016, c.root, img)
		//line article.qtpl:93
		qw422016.N().S(`</span><span class="fileinfo">`)
		//line article.qtpl:96
		if img.Artist != "" {
			//line article.qtpl:96
			qw422016.N().S(`<span class="media-artist">`)
			//line article.qtpl:98
			qw422016.E().S(img.Artist)
			//line article.qtpl:98
			qw422016.N().S(`</span>`)
			//line article.qtpl:100
		}
		//line article.qtpl:101
		if img.Title != "" {
			//line article.qtpl:101
			qw422016.N().S(`<span class="media-title">`)
			//line article.qtpl:103
			qw422016.E().S(img.Title)
			//line article.qtpl:103
			qw422016.N().S(`</span>`)
			//line article.qtpl:105
		}
		//line article.qtpl:106
		if img.Audio {
			//line article.qtpl:106
			qw422016.N().S(`<span class="has-audio">♫</span>`)
			//line article.qtpl:110
		}
		//line article.qtpl:111
		if img.Length != 0 {
			//line article.qtpl:111
			qw422016.N().S(`<span class="media-length">`)
			//line article.qtpl:113
			l := img.Length

			//line article.qtpl:114
			if l < 60 {
				//line article.qtpl:115
				qw422016.N().S(fmt.Sprintf("0:%02d", l))
				//line article.qtpl:116
			} else {
				//line article.qtpl:117
				min := l / 60

				//line article.qtpl:118
				qw422016.N().S(fmt.Sprintf("%02d:%02d", min, l-min*60))
				//line article.qtpl:119
			}
			//line article.qtpl:119
			qw422016.N().S(`</span>`)
			//line article.qtpl:121
		}
		//line article.qtpl:122
		if img.APNG {
			//line article.qtpl:122
			qw422016.N().S(`<span class="is-apng">APNG</span>`)
			//line article.qtpl:126
		}
		//line article.qtpl:126
		qw422016.N().S(`<span class="filesize">`)
		//line article.qtpl:128
		qw422016.N().S(readableFileSize(img.Size))
		//line article.qtpl:128
		qw422016.N().S(`</span>`)
		//line article.qtpl:130
		if img.Dims != [4]uint16{} {
			//line article.qtpl:130
			qw422016.N().S(`<span class="dims">`)
			//line article.qtpl:132
			qw422016.N().S(strconv.FormatUint(uint64(img.Dims[0]), 10))
			//line article.qtpl:132
			qw422016.N().S(`x`)
			//line article.qtpl:134
			qw422016.N().S(strconv.FormatUint(uint64(img.Dims[1]), 10))
			//line article.qtpl:134
			qw422016.N().S(`</span>`)
			//line article.qtpl:136
		}
		//line article.qtpl:136
		qw422016.N().S(`</span>`)
		//line article.qtpl:138
		name := imageName(img.FileType, img.Name)

		//line article.qtpl:138
		qw422016.N().S(`<a href="`)
		//line article.qtpl:139
		qw422016.N().S(assets.RelativeSourcePath(img.FileType, img.SHA1))
		//line article.qtpl:139
		qw422016.N().S(`" download="`)
		//line article.qtpl:139
		qw422016.N().S(name)
		//line article.qtpl:139
		qw422016.N().S(`">`)
		//line article.qtpl:140
		qw422016.N().S(name)
		//line article.qtpl:140
		qw422016.N().S(`</a></figcaption>`)
		//line article.qtpl:143
	}
	//line article.qtpl:143
	qw422016.N().S(`<div class="post-container">`)
	//line article.qtpl:145
	if p.Image != nil {
		//line article.qtpl:146
		img := *p.Image

		//line article.qtpl:146
		qw422016.N().S(`<figure><a target="_blank" href="`)
		//line article.qtpl:148
		qw422016.N().S(src)
		//line article.qtpl:148
		qw422016.N().S(`">`)
		//line article.qtpl:149
		switch {
		//line article.qtpl:150
		case img.ThumbType == common.NoFile:
			//line article.qtpl:151
			var file string

			//line article.qtpl:152
			switch img.FileType {
			//line article.qtpl:153
			case common.MP4, common.MP3, common.OGG, common.FLAC:
				//line article.qtpl:154
				file = "audio"

			//line article.qtpl:155
			default:
				//line article.qtpl:156
				file = "file"

				//line article.qtpl:157
			}
			//line article.qtpl:157
			qw422016.N().S(`<img src="/assets/`)
			//line article.qtpl:158
			qw422016.N().S(file)
			//line article.qtpl:158
			qw422016.N().S(`.png" width="150" height="150">`)
		//line article.qtpl:159
		case img.Spoiler:
			//line article.qtpl:162
			qw422016.N().S(`<img src="/assets/spoil/default.jpg" width="150" height="150">`)
		//line article.qtpl:164
		default:
			//line article.qtpl:164
			qw422016.N().S(`<img src="`)
			//line article.qtpl:165
			qw422016.N().S(assets.ThumbPath(img.ThumbType, img.SHA1))
			//line article.qtpl:165
			qw422016.N().S(`" width="`)
			//line article.qtpl:165
			qw422016.N().D(int(img.Dims[2]))
			//line article.qtpl:165
			qw422016.N().S(`" height="`)
			//line article.qtpl:165
			qw422016.N().D(int(img.Dims[3]))
			//line article.qtpl:165
			qw422016.N().S(`">`)
			//line article.qtpl:166
		}
		//line article.qtpl:166
		qw422016.N().S(`</a></figure>`)
		//line article.qtpl:169
	}
	//line article.qtpl:169
	qw422016.N().S(`<blockquote>`)
	//line article.qtpl:171
	streambody(qw422016, p, c.op, c.board, c.index, c.rbText, c.pyu)
	//line article.qtpl:171
	qw422016.N().S(`</blockquote>`)
	//line article.qtpl:173
	for _, e := range p.Moderation {
		//line article.qtpl:173
		qw422016.N().S(`<b class="admin post-moderation">`)
		//line article.qtpl:175
		streampostModeration(qw422016, e)
		//line article.qtpl:175
		qw422016.N().S(`<br></b>`)
		//line article.qtpl:178
	}
	//line article.qtpl:178
	qw422016.N().S(`</div>`)
	//line article.qtpl:180
	if c.omit != 0 {
		//line article.qtpl:180
		qw422016.N().S(`<span class="omit" data-omit="`)
		//line article.qtpl:181
		qw422016.N().D(c.omit)
		//line article.qtpl:181
		qw422016.N().S(`" data-image-omit="`)
		//line article.qtpl:181
		qw422016.N().D(c.imageOmit)
		//line article.qtpl:181
		qw422016.N().S(`">`)
		//line article.qtpl:182
		qw422016.N().S(pluralize(c.omit, "post"))
		//line article.qtpl:183
		qw422016.N().S(` `)
		//line article.qtpl:183
		qw422016.N().S(ln.Common.Posts["and"])
		//line article.qtpl:183
		qw422016.N().S(` `)
		//line article.qtpl:184
		qw422016.N().S(pluralize(c.imageOmit, "image"))
		//line article.qtpl:185
		qw422016.N().S(` `)
		//line article.qtpl:185
		qw422016.N().S(`omitted`)
		//line article.qtpl:185
		qw422016.N().S(` `)
		//line article.qtpl:185
		qw422016.N().S(`<span class="act"><a href="`)
		//line article.qtpl:187
		qw422016.N().S(strconv.FormatUint(c.op, 10))
		//line article.qtpl:187
		qw422016.N().S(`">`)
		//line article.qtpl:188
		qw422016.N().S(ln.Common.Posts["seeAll"])
		//line article.qtpl:188
		qw422016.N().S(`</a></span></span>`)
		//line article.qtpl:192
	}
	//line article.qtpl:193
	if bls := c.backlinks[p.ID]; len(bls) != 0 {
		//line article.qtpl:193
		qw422016.N().S(`<span class="backlinks spaced">`)
		//line article.qtpl:195
		for _, l := range bls {
			//line article.qtpl:195
			qw422016.N().S(`<em>`)
			//line article.qtpl:197
			streampostLink(qw422016, l, c.index || l.OP != c.op, c.index)
			//line article.qtpl:197
			qw422016.N().S(`</em>`)
			//line article.qtpl:199
		}
		//line article.qtpl:199
		qw422016.N().S(`</span>`)
		//line article.qtpl:201
	}
	//line article.qtpl:201
	qw422016.N().S(`</article>`)
//line article.qtpl:203
}

//line article.qtpl:203
func writerenderArticle(qq422016 qtio422016.Writer, p common.Post, c articleContext) {
	//line article.qtpl:203
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:203
	streamrenderArticle(qw422016, p, c)
	//line article.qtpl:203
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:203
}

//line article.qtpl:203
func renderArticle(p common.Post, c articleContext) string {
	//line article.qtpl:203
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:203
	writerenderArticle(qb422016, p, c)
	//line article.qtpl:203
	qs422016 := string(qb422016.B)
	//line article.qtpl:203
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:203
	return qs422016
//line article.qtpl:203
}

// Render image search links according to file type

//line article.qtpl:206
func streamimageSearch(qw422016 *qt422016.Writer, root string, img common.Image) {
	//line article.qtpl:207
	if img.ThumbType == common.NoFile || img.FileType == common.PDF {
		//line article.qtpl:208
		return
		//line article.qtpl:209
	}
	//line article.qtpl:211
	url := root + assets.ImageSearchPath(img.ImageCommon)

	//line article.qtpl:211
	qw422016.N().S(`<a class="image-search google" target="_blank" rel="nofollow" href="https://www.google.com/searchbyimage?image_url=`)
	//line article.qtpl:212
	qw422016.N().S(url)
	//line article.qtpl:212
	qw422016.N().S(`">G</a><a class="image-search iqdb" target="_blank" rel="nofollow" href="http://iqdb.org/?url=`)
	//line article.qtpl:215
	qw422016.N().S(url)
	//line article.qtpl:215
	qw422016.N().S(`">Iq</a><a class="image-search saucenao" target="_blank" rel="nofollow" href="http://saucenao.com/search.php?db=999&url=`)
	//line article.qtpl:218
	qw422016.N().S(url)
	//line article.qtpl:218
	qw422016.N().S(`">Sn</a><a class="image-search whatAnime" target="_blank" rel="nofollow" href="https://trace.moe/?url=`)
	//line article.qtpl:221
	qw422016.N().S(url)
	//line article.qtpl:221
	qw422016.N().S(`">Wa</a>`)
	//line article.qtpl:224
	switch img.FileType {
	//line article.qtpl:225
	case common.JPEG, common.PNG, common.GIF, common.WEBM:
		//line article.qtpl:225
		qw422016.N().S(`<a class="image-search desustorage" target="_blank" rel="nofollow" href="https://desuarchive.org/_/search/image/`)
		//line article.qtpl:226
		qw422016.N().S(img.MD5)
		//line article.qtpl:226
		qw422016.N().S(`">Ds</a>`)
		//line article.qtpl:229
	}
	//line article.qtpl:230
	switch img.FileType {
	//line article.qtpl:231
	case common.JPEG, common.PNG:
		//line article.qtpl:231
		qw422016.N().S(`<a class="image-search exhentai" target="_blank" rel="nofollow" href="http://exhentai.org/?fs_similar=1&fs_exp=1&f_shash=`)
		//line article.qtpl:232
		qw422016.N().S(img.SHA1)
		//line article.qtpl:232
		qw422016.N().S(`">Ex</a>`)
		//line article.qtpl:235
	}
//line article.qtpl:236
}

//line article.qtpl:236
func writeimageSearch(qq422016 qtio422016.Writer, root string, img common.Image) {
	//line article.qtpl:236
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:236
	streamimageSearch(qw422016, root, img)
	//line article.qtpl:236
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:236
}

//line article.qtpl:236
func imageSearch(root string, img common.Image) string {
	//line article.qtpl:236
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:236
	writeimageSearch(qb422016, root, img)
	//line article.qtpl:236
	qs422016 := string(qb422016.B)
	//line article.qtpl:236
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:236
	return qs422016
//line article.qtpl:236
}
