// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute"
)

var spinVditorDOMTests = []*parseTest{

	{"147", "<p data-block=\"0\"><mark data-marker=\"==\">markf<wbr></mark></p>", "<p data-block=\"0\"><mark data-marker=\"==\">markf<wbr></mark></p>"},
	{"146", "<p data-block=\"0\">==mark==<wbr></p>", "<p data-block=\"0\"><mark data-marker=\"==\">mark</mark><wbr></p>"},
	{"145", "<p data-block=\"0\">foo<vditor-comment data-id=\"1\">bar<wbr>baz</vditor-comment>foo<vditor-comment data-id=\"2\">bar</vditor-comment>baz</p>", "<p data-block=\"0\">foo\u200b<vditor-comment data-id=\"1\">bar<wbr>baz\u200b</vditor-comment>foo\u200b<vditor-comment data-id=\"2\">bar\u200b</vditor-comment>baz</p>"},
	{"144", "<p data-block=\"0\"><vditor-comment data-id=\"1\">foo<wbr></vditor-comment></p>", "<p data-block=\"0\">\u200b\u200b<vditor-comment data-id=\"1\">foo<wbr>\u200b</vditor-comment></p>"},
	{"143", "<p data-block=\"0\">​<img src=\"bar\" alt=\"text\" data-type=\"link-ref\" data-link-label=\"foo\"><wbr></p><div data-block=\"0\" data-type=\"link-ref-defs-block\">[foo]: bar\n</div>", "<p data-block=\"0\">\u200b<img src=\"bar\" alt=\"text\" data-type=\"link-ref\" data-link-label=\"foo\" /><wbr></p><div data-block=\"0\" data-type=\"link-ref-defs-block\">[foo]: bar\n</div>"},
	{"142", "<p data-block=\"0\">![text][foo]<wbr></p><div data-block=\"0\" data-type=\"link-ref-defs-block\">[foo]: bar", "<p data-block=\"0\">\u200b<img src=\"bar\" alt=\"text\" data-type=\"link-ref\" data-link-label=\"foo\" /><wbr></p><div data-block=\"0\" data-type=\"link-ref-defs-block\">[foo]: bar\n</div>"},
	{"141", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">​</li><li data-marker=\"*\">b<wbr></li></ul></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo\n*<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">b<wbr></li></ul></li></ul>"},
	{"140", "<h2 data-block=\"0\" id=\"wysiwyg-fo_0\" data-marker=\"#\" data-id=\"#custom-id\">foo<wbr></h2>", "<h2 data-block=\"0\" data-id=\"#custom-id\" id=\"wysiwyg-#custom-id\" data-marker=\"#\">foo<wbr></h2>"},
	{"139", "<p data-block=\"0\">​<sup data-type=\"footnotes-ref\" data-footnotes-label=\"^foo\" class=\"vditor-tooltipped vditor-tooltipped__s\" aria-label=\"\">1<wbr></sup></p><div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^foo\"></li></ol></div>", "<p data-block=\"0\">\u200b<sup data-type=\"footnotes-ref\" data-footnotes-label=\"^foo\" class=\"vditor-tooltipped vditor-tooltipped__s\" aria-label=\"\">1</sup>\u200b<wbr></p><div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^foo\"></li></ol></div>"},
	{"138", "<p data-block=\"0\">​<code data-marker=\"``\">​`code`</code>​f<wbr></p>", "<p data-block=\"0\">\u200b<code data-marker=\"``\">\u200b`code`</code>\u200b f<wbr></p>"},
	{"137", "<p>\\<wbr>$foo$</p>", "<p data-block=\"0\"><span data-type=\"backslash\"><span>\\</span>$</span><wbr>foo$</p>"},
	{"136", "<ul data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"><p data-block=\"0\">foo</p><div class=\"vditor-wysiwyg__block\" data-type=\"math-block\" data-block=\"0\"><pre><code data-type=\"math-block\">bar<wbr></code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"1\"><code class=\"language-math\"><div class=\"vditor-math\" data-math=\"ba\"><span class=\"katex-display\"><span class=\"katex\"><span class=\"katex-html\" aria-hidden=\"true\"><span class=\"base\"><span class=\"strut\"></span><span class=\"mord mathdefault\">b</span><span class=\"mord mathdefault\">a</span></span></span></span></span></div></code></pre></div></li></ul>", "<ul data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"><p data-block=\"0\">foo</p><div class=\"vditor-wysiwyg__block\" data-type=\"math-block\" data-block=\"0\"><pre><code data-type=\"math-block\">bar<wbr></code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">bar</code></pre></div></li></ul>"},
	{"135", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<wbr><div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code>bar\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"1\"><div class=\"vditor-copy\"><textarea></textarea><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"vditor-tooltipped vditor-tooltipped__w\" onclick=\"this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制')\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 32 32\" width=\"32px\" height=\"32px\"> <path d=\"M28.681 11.159c-0.694-0.947-1.662-2.053-2.724-3.116s-2.169-2.030-3.116-2.724c-1.612-1.182-2.393-1.319-2.841-1.319h-11.5c-1.379 0-2.5 1.121-2.5 2.5v23c0 1.378 1.121 2.5 2.5 2.5h19c1.378 0 2.5-1.122 2.5-2.5v-15.5c0-0.448-0.137-1.23-1.319-2.841zM24.543 9.457c0.959 0.959 1.712 1.825 2.268 2.543h-4.811v-4.811c0.718 0.556 1.584 1.309 2.543 2.268v0zM28 29.5c0 0.271-0.229 0.5-0.5 0.5h-19c-0.271 0-0.5-0.229-0.5-0.5v-23c0-0.271 0.229-0.5 0.5-0.5 0 0 11.499-0 11.5 0v7c0 0.552 0.448 1 1 1h7v15.5zM18.841 1.319c-1.612-1.182-2.393-1.319-2.841-1.319h-11.5c-1.378 0-2.5 1.121-2.5 2.5v23c0 1.207 0.86 2.217 2 2.45v-25.45c0-0.271 0.229-0.5 0.5-0.5h15.215c-0.301-0.248-0.595-0.477-0.873-0.681z\"></path> </svg></span></div><code class=\"hljs nginx\"><span class=\"hljs-attribute\">bar</span>\n</code></pre></div></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<wbr><div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\" style=\"display: none\"><code>bar\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code>bar\n</code></pre></div></li></ul>"},
	{"134", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img <wbr> src=\"\"&gt;</code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img <wbr> src=&quot;&quot;&gt;</code></span>\u200b</p>"},
	{"133", "<p data-block=\"0\">​f<wbr>​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img src=\"\"&gt;</code></span>​</p>", "<p data-block=\"0\">f<wbr><span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot;&gt;</code></span>\u200b</p>"},
	{"132", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img src=\"\" <wbr>&gt;</code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot; <wbr>&gt;</code></span>\u200b</p>"},
	{"131", "<p data-block=\"0\">&lt;img src=\"\"<wbr>&gt;1</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot;<wbr>&gt;</code></span>\u200b1</p>"},
	{"130", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img s<wbr>&gt;</code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img s<wbr>&gt;</code></span>\u200b</p>"},
	{"129", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img <wbr>&gt;</code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img <wbr>&gt;</code></span>\u200b</p>"},
	{"128", "<p data-block=\"0\">&lt;img <wbr>src=\"\"&gt;</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img <wbr>src=&quot;&quot;&gt;</code></span>\u200b</p>"},
	{"127", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;bar <wbr>&gt;</code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;bar <wbr>&gt;</code></span>\u200b</p>"},
	{"126", "<p data-block=\"0\">&lt;img src=\"\"&gt;<wbr></p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot;&gt;</code></span>\u200b<wbr></p>"},
	{"125", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img src=\"\"&gt;<wbr></code></span>​</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot;&gt;</code></span>\u200b<wbr></p>"},
	{"124", "<p data-block=\"0\">&lt;img src=\"\"<wbr></p>", "<p data-block=\"0\">&lt;img src=&quot;&quot;<wbr></p>"},
	{"123", "<p data-block=\"0\">​<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">​&lt;img src=\"\" <wbr> &gt;</code></span>​ 1</p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"html-inline\"><code data-type=\"html-inline\">\u200b&lt;img src=&quot;&quot; <wbr> &gt;</code></span>\u200b 1</p>"},
	{"122", "<p data-block=\"0\">&lt;img src 。<wbr>", "<p data-block=\"0\">&lt;img src 。<wbr></p>"},
	{"121", "<p data-block=\"0\">&amp;<wbr></p>", "<p data-block=\"0\">&amp;<wbr></p>"},
	{"120", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\"> test<wbr></li></ul><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\"> test</li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> test<wbr></li><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> test</li></ul>"},
	{"119", "&parx", "<p data-block=\"0\">&amp;parx</p>"},
	{"118", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\"><p>[ ]<wbr></p></li></ul>", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\" class=\"vditor-task\"><input type=\"checkbox\" /> <wbr></li></ul>"},
	{"117", "<p data-block=\"0\">foo<wbr><a>​</a></p>", "<p data-block=\"0\">foo<wbr></p>"},
	{"116", "<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1.\">foo<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1.\">bar</li></ol><ol data-block=\"0\"><li data-marker=\"1.\">‸foo2<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1.\">bar2</li></ol></li></ol></li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">foo<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">bar</li><li data-marker=\"2.\"><wbr>foo2<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">bar2</li></ol></li></ol></li></ol>"},
	{"115", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul><ul data-block=\"0\"><li data-marker=\"*\"><wbr>foo2<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar2</li></ul></li></ul></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar</li><li data-marker=\"*\"><wbr>foo2<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar2</li></ul></li></ul></li></ul>"},
	{"114", "<p data-block=\"0\">```<wbr>a b\nc</p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code class=\"language-a\"><wbr>c\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-a\">c\n</code></pre></div>"},
	{"113", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\"><p>[ <wbr>]</p></li></ul>", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\" class=\"vditor-task\"><input type=\"checkbox\" /> <wbr></li></ul>"},
	{"112", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"></li><li data-marker=\"*\"><p>f<wbr></p></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">\u200b</li><li data-marker=\"*\">f<wbr></li></ul>"},
	{"111", "<h1 data-block=\"0\" data-marker=\"#\">foo {#custom-id}<wbr></h1>", "<h1 data-block=\"0\" data-id=\"#custom-id\" id=\"wysiwyg-#custom-id\" data-marker=\"#\">foo<wbr></h1>"},
	{"110", "<div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^111\"><p data-block=\"0\">1<wbr></p></li></ol></div>", "<div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^111\"><p data-block=\"0\">1<wbr></p></li></ol></div>"},
	// 109：重复的脚注定义 marker 会被去重，重现步骤：在脚注定义中换行
	{"109", "<div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^1\"><p data-block=\"0\">foo</p></li><li data-type=\"footnotes-li\" data-marker=\"^1\"><p data-block=\"0\"><wbr>bar</p></li></ol></div>", "<div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^1\"><p data-block=\"0\">foo</p></li></ol></div>"},
	{"108", "<p data-block=\"0\"><wbr>## heading</p>", "<h2 data-block=\"0\" id=\"wysiwyg-heading\" data-marker=\"#\"><wbr>heading</h2>"},
	{"107", "<div class=\"toc-div\" data-type=\"toc-block\"><span class=\"toc-h1\"><a class=\"toc-a\" href=\"#foo\">foo</a></span><br></div>\n\n<h1 data-block=\"0\" data-marker=\"#\">foo</h1>", "<div class=\"vditor-toc\" data-block=\"0\" data-type=\"toc-block\" contenteditable=\"false\"><span data-type=\"toc-h\">foo</span><br></div><p data-block=\"0\"></p><h1 data-block=\"0\" id=\"wysiwyg-foo\" data-marker=\"#\">foo</h1>"},
	{"106", "<p data-block=\"0\"><sup data-type=\"footnotes-ref\" data-footnotes-label=\"^1\">1</sup></p><div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^1\">bar</li></ol></div>", "<p data-block=\"0\">\u200b<sup data-type=\"footnotes-ref\" data-footnotes-label=\"^1\" class=\"vditor-tooltipped vditor-tooltipped__s\" aria-label=\"bar\">1</sup>\u200b</p><div data-block=\"0\" data-type=\"footnotes-block\"><ol data-type=\"footnotes-defs-ol\"><li data-type=\"footnotes-li\" data-marker=\"^1\"><p data-block=\"0\">bar</p></li></ol></div>"},
	{"105", "<p data-block=\"0\"><span data-type=\"link-ref\" data-link-text=\"1\" data-link-label=\"1\">1</span></p><p data-block=\"0\" data-type=\"link-ref-defs\">[1]: f<wbr></p>", "<p data-block=\"0\">\u200b<span data-type=\"link-ref\" data-link-label=\"1\">1</span>\u200b</p><div data-block=\"0\" data-type=\"link-ref-defs-block\">[1]: f<wbr>\n</div>"},
	{"104", "<a href=\"\" title=\"baz\">foo</a>", "<p data-block=\"0\"><a href=\"\"baz\"\">foo</a></p>"},
	{"103", "<p data-block=\"0\"><strong data-marker=\"**\">foo\n<em data-marker=\"*\">ba<wbr></em></strong></p>", "<p data-block=\"0\"><strong data-marker=\"**\">foo\n<em data-marker=\"*\">ba<wbr></em></strong></p>"},
	{"102", "<p data-block=\"0\"><strong data-marker=\"**\">foo<em>\u200b\nb<wbr></em></strong></p>", "<p data-block=\"0\"><strong data-marker=\"**\">foo\n<em data-marker=\"*\">b<wbr></em></strong></p>"},
	// 101：hr 已由 Vditor 处理
	{"101", `<ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*"><ul data-tight="true" data-marker="-" data-block="0"><li data-marker="-"><p>- -<wbr></p></li></ul></li></ul>`, "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"><ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\"><ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\">-<wbr></li></ul></li></ul></li></ul>"},
	{"100", "<em data-marker=\"*\">foo\nbar</em>\t\n---", "<h2 data-block=\"0\" id=\"wysiwyg-foobar\" data-marker=\"-\"><em data-marker=\"*\">foo\nbar</em></h2>"},
	{"99", `<ol data-block="0"><li><p>f<wbr></p></li></ol>`, "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">f<wbr></li></ol>"},
	{"98", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bcode</code>​    <wbr><code data-marker=\"`\">\u200bcode</code>\u200b</p>", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bcode</code>\u200b    <wbr><code data-marker=\"`\">\u200bcode</code>\u200b</p>"},
	{"97", "<p data-block=\"0\"><strong data-marker=\"**\">foo\n</strong>b<wbr></p>", "<p data-block=\"0\"><strong data-marker=\"**\">foo</strong>\nb<wbr></p>"},
	{"96", "<p data-block=\"0\">​<code data-marker=\"`\">\u200bcode<wbr></code><code data-marker=\"`\">\u200bspan</code><span>\u200b</span></p>", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bcode<wbr>span</code>\u200b</p>"},
	{"95", "<p data-block=\"0\"><strong><em><wbr>\u200b</em></strong></p>", ""},
	{"94", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bcode\nspan<wbr></code>\u200b</p>", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bcode span<wbr></code>\u200b</p>"},
	{"93", "<p data-block=\"0\"><strong data-marker=\"**\"><wbr></strong></p>", "<p data-block=\"0\"><wbr></p>"},
	{"92", "<p data-block=\"0\"><strong><em>\u200b<wbr></em></strong></p>", ""},
	{"91", "<p data-block=\"0\"><wbr>    ***</p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code><wbr>***\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code>***\n</code></pre></div>"},
	{"90", "<p data-block=\"0\">    ***<wbr></p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code>***<wbr>\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code>***\n</code></pre></div>"},
	{"89", `<ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*" class="vditor-task"><p><wbr><input type="checkbox"> foo</p></li></ul>`, "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> <wbr>foo</li></ul>"},
	// 88：hr 已由 Vditor 处理
	{"88", `<ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*"><p>test</p></li></ul><ul data-tight="true" data-marker="-" data-block="0"><li data-marker="-"><p>--<wbr></p></li></ul>`, "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">test</li></ul><ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\">--<wbr></li></ul>"},
	{"87", `<ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*" class="vditor-task"><p><input type="checkbox"> foo</p><ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*" class="vditor-task"><p><input type="checkbox"> bar</p></li></ul><p data-block="0">b<wbr></p></li></ul>`, "<ul data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><p data-block=\"0\"><input type=\"checkbox\" /> foo</p><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> bar</li></ul><p data-block=\"0\">b<wbr></p></li></ul>"},
	{"86", "<p data-block=\"0\"><strong><s>foo</s></strong>bar<wbr></p>", "<p data-block=\"0\"><strong data-marker=\"**\"><s data-marker=\"~~\">foo</s></strong>bar<wbr></p>"},
	{"85", "<p data-block=\"0\"><span data-marker=\"**\"><b>foo </b><span>b<wbr></span></span>", "<p data-block=\"0\"><strong data-marker=\"**\">foo</strong> b<wbr></p>"},
	{"84", "<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1)\"><p>f<wbr></p></li></ol>", "<ol data-tight=\"true\" data-marker=\"1)\" data-block=\"0\"><li data-marker=\"1)\">f<wbr></li></ol>"},
	{"83", "<p data-block=\"0\">1) f<wbr></p>", "<ol data-tight=\"true\" data-marker=\"1)\" data-block=\"0\"><li data-marker=\"1)\">f<wbr></li></ol>"},
	{"82", "<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"2.\"><p>bar<wbr></p></li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">bar<wbr></li></ol>"},
	{"81", "<p data-block=\"0\"><strong>\u200b<em>\u200b<s>\u200b1</s></em></strong><wbr></p>", "<p data-block=\"0\"><em data-marker=\"*\"><strong data-marker=\"**\"><s data-marker=\"~~\">1</s></strong></em><wbr></p>"},
	{"80", "<s><em>\u200b</em></s>", ""},
	{"79", "<p data-block=\"0\"><b>&#8203;</b></p>", ""},
	{"78", "<p data-block=\"0\">``foo``<wbr></p>", "<p data-block=\"0\">\u200b<code data-marker=\"``\">\u200bfoo</code>\u200b<wbr></p>"},
	{"77", `<p data-block="0">1<wbr><span style="background-color: var(--textarea-focus-background-color); color: var(--textarea-text-color);">​</span><span class="vditor-wysiwyg__block" data-type="math-inline" style="background-color: var(--textarea-focus-background-color); color: var(--textarea-text-color);"><code data-type="math-inline">foo</code><span class="vditor-wysiwyg__preview" data-render="false"><span class="vditor-math" data-math="foo"><span class="katex"><span class="katex-html" aria-hidden="true"><span class="base"><span class="strut" style="height:0.85396em;vertical-align:-0.19444em;"></span><span class="mord mathdefault" style="margin-right:0.05724em;">f</span><span class="mord mathdefault" style="margin-right:0.05724em;">o</span><span class="mord mathdefault" style="margin-right:0.05724em;">o</span></span></span></span></span></span></span><span style="background-color: var(--textarea-focus-background-color); color: var(--textarea-text-color);">​</span></p>`, "<p data-block=\"0\">1<wbr><span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\" style=\"display: none\">\u200bfoo</code><span class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span></span>\u200b</p>"},
	{"76", "<ul><li data-marker=\"1.\"><p>12<wbr></p></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">12<wbr></li></ul>"},
	{"75", "<ol><li data-marker=\"*\"><p>foo<wbr></p></li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">foo<wbr></li></ol>"},
	{"74", "<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1.\" class=\"vditor-task\"><p><input checked=\"\" type=\"checkbox\"> f<wbr></p></li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> f<wbr></li></ol>"},
	{"73", "<ol data-tight=\"true\" data-block=\"0\"><li data-marker=\"1.\"><p>[x] foo<wbr></p></li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> foo<wbr></li></ol>"},
	{"72", "<p data-block=\"0\">foo\n-<wbr></p>", "<h2 data-block=\"0\" id=\"wysiwyg-foo\" data-marker=\"-\">foo<wbr></h2>"},
	{"71", "<p data-block=\"0\">foo<span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\">\u200bbar</code></span> </p>", "<p data-block=\"0\">foo<span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\" style=\"display: none\">\u200bbar</code><span class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-math\">bar</code></span></span>\u200b</p>"},
	{"70", "<p data-block=\"0\"> <span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\">\u200bfoo</code></span> </p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\" style=\"display: none\">\u200bfoo</code><span class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span></span>\u200b</p>"},
	{"69", "<p data-block=\"0\"><a href=\"/bar\"><code>foo</code></a><wbr></p>", "<p data-block=\"0\"><a href=\"/bar\">\u200b<code data-marker=\"`\">\u200bfoo</code></a><wbr></p>"},
	{"68", `<p data-block="0">|foo|bar|<wbr></p>`, "<p data-block=\"0\">|foo|bar|<wbr></p>"},
	{"67", `<ul data-tight="true" data-marker="*" data-block="0"><li data-marker="*"><p>[ ]<wbr></p></li></ul>`, "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> <wbr></li></ul>"},
	{"66", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><p><input type=\"checkbox\" checked=\"checked\"><wbr> foo</p></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> <wbr>foo</li></ul>"},
	{"65", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"><p>foo<em data-marker=\"*\">bar</em></p></li><li data-marker=\"*\"><p><em data-marker=\"*\"><wbr><br></em></p></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<em data-marker=\"*\">bar</em></li><li data-marker=\"*\"><wbr></li></ul>"},
	{"64", "<p data-block=\"0\">[foo<wbr>](/bar)", "<p data-block=\"0\"><a href=\"/bar\">foo<wbr></a></p>"},
	{"63", "<p data-block=\"0\">![foo<wbr>](/bar)", "<p data-block=\"0\"><img src=\"/bar\" alt=\"foo\" /></p>"},
	{"62", "<p data-block=\"0\"><strong data-marker=\"__\"><wbr><br></strong></p>", "<p data-block=\"0\"><wbr></p>"},
	{"61", "<p data-block=\"0\">_foo_<wbr></p>", "<p data-block=\"0\"><em data-marker=\"_\">foo</em><wbr></p>"},
	{"60", "<p data-block=\"0\">foo\n=<wbr></p>", "<h1 data-block=\"0\" id=\"wysiwyg-foo\" data-marker=\"=\">foo<wbr></h1>"},
	{"59", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\"><p>foo</p><ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\"><p>bar<wbr><p></li></ul></li></ul>", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\">foo<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\">bar<wbr></li></ul></li></ul>"},
	{"58", "<p data-block=\"0\">![](/bar)<wbr></p>", "<p data-block=\"0\"><img src=\"/bar\" alt=\"\" /><wbr></p>"},
	{"57", "<p data-block=\"0\">/<span data-type=\"backslash\"><span>\\</span>_</span><span data-type=\"backslash\"><span>\\</span>_</span>foo__.</p>", "<p data-block=\"0\">/<span data-type=\"backslash\"><span>\\</span>_</span><span data-type=\"backslash\"><span>\\</span>_</span>foo__.</p>"},
	{"56", "<p data-block=\"0\">[foo](bar<wbr>)</p>", "<p data-block=\"0\"><a href=\"bar\">foo<wbr></a></p>"},
	{"55", "<p data-block=\"0\">[foo<wbr>](bar)</p>", "<p data-block=\"0\"><a href=\"bar\">foo<wbr></a></p>"},
	{"54", "<table data-block=\"0\"><thead><tr><th>col1</th></tr></thead><tbody><tr><td><wbr>\n</td></tr></tbody></table>", "<table data-block=\"0\"><thead><tr><th>col1</th></tr></thead><tbody><tr><td><wbr> </td></tr><tr><td> </td></tr></tbody></table>"},
	{"53", "<table data-block=\"0\"><thead><tr><th>col1</th></tr></thead><tbody><tr><td><wbr><br></td></tr></tbody></table>", "<table data-block=\"0\"><thead><tr><th>col1</th></tr></thead><tbody><tr><td><wbr> </td></tr></tbody></table>"},
	{"52", "<p data-block=\"0\">---<wbr></p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"yaml-front-matter\" data-block=\"0\"><pre><code data-type=\"yaml-front-matter\"><wbr>\n</code></pre></div>"},
	{"51", "<p data-block=\"0\">### <wbr></p>", "<p data-block=\"0\">### <wbr></p>"},
	{"50", "<details open=\"\">\n<summary>foo</summary><ul data-tight=\"true\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul></details>", "<div class=\"vditor-wysiwyg__block\" data-type=\"html-block\" data-block=\"0\"><pre><code>&lt;details open=&quot;&quot;&gt;\n&lt;summary&gt;foo&lt;/summary&gt;</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><details open=\"\">\n<summary>foo</summary></pre></div><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul><div class=\"vditor-wysiwyg__block\" data-type=\"html-block\" data-block=\"0\"><pre><code>&lt;/details&gt;</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"></details></pre></div>"},
	{"49", "<details>\n<summary>foo</summary><ul data-tight=\"true\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul></details>", "<div class=\"vditor-wysiwyg__block\" data-type=\"html-block\" data-block=\"0\"><pre><code>&lt;details&gt;\n&lt;summary&gt;foo&lt;/summary&gt;</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><details>\n<summary>foo</summary></pre></div><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul><div class=\"vditor-wysiwyg__block\" data-type=\"html-block\" data-block=\"0\"><pre><code>&lt;/details&gt;</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"></details></pre></div>"},
	{"48", "<p data-block=\"0\"><a href=\"中文\">link</a><wbr></p>", "<p data-block=\"0\"><a href=\"中文\">link</a><wbr></p>"},
	{"47", "<p data-block=\"0\">`1<wbr>`</p>", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200b1<wbr></code>\u200b</p>"},
	{"46", "<p>- [x] f<wbr></p>", "<ul data-tight=\"true\" data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> f<wbr></li></ul>"},
	{"45", "<ul data-tight=\"true\"><li data-marker=\"-\" class=\"vditor-task\"><input type=\"checkbox\"> foo</li></ul><p>- [ ] b<wbr></p>", "<ul data-marker=\"-\" data-block=\"0\"><li data-marker=\"-\" class=\"vditor-task\"><p data-block=\"0\"><input type=\"checkbox\" /> foo</p></li><li data-marker=\"-\" class=\"vditor-task\"><p data-block=\"0\"><input type=\"checkbox\" /> b<wbr></p></li></ul>"},
	{"44", "<p data-block=\"0\">* [ ]<wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> <wbr></li></ul>"},
	{"43", "<p>* [ <wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">[ <wbr></li></ul>"},
	{"42", "<p>* [<wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">[<wbr></li></ul>"},
	{"40", "<h3>隐藏细节</h3><div class=\"vditor-wysiwyg__block\" data-type=\"html-block\"><pre><code>&lt;details&gt;\n&lt;summary&gt;\n这里是摘要部分。\n&lt;/summary&gt;\n这里是细节部分。\n&lt;/details&gt;<br></code></pre><div class=\"vditor-wysiwyg__preview\" contenteditable=\"false\" data-render=\"1\"></div></div><p>1<wbr></p>", "<h3 data-block=\"0\" id=\"wysiwyg-隐藏细节\" data-marker=\"#\">隐藏细节</h3><div class=\"vditor-wysiwyg__block\" data-type=\"html-block\" data-block=\"0\"><pre><code>&lt;details&gt;\n&lt;summary&gt;\n这里是摘要部分。\n&lt;/summary&gt;\n这里是细节部分。\n&lt;/details&gt;</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><details>\n<summary>\n这里是摘要部分。\n</summary>\n这里是细节部分。\n</details></pre></div><p data-block=\"0\">1<wbr></p>"},
	{"39", "<p>*foo<wbr>*bar</p>", "<p data-block=\"0\"><em data-marker=\"*\">foo<wbr></em>bar</p>"},
	{"38", "<p>[foo](b<wbr>)</p>", "<p data-block=\"0\"><a href=\"b\">foo<wbr></a></p>"},
	{"37", "<blockquote><p><wbr></p></blockquote>", ""},
	{"36", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-marker=\"```\"><pre><code class=\"language-go\">foo</code></pre></div>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\" style=\"display: none\"><code class=\"language-go\">foo\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-go\">foo\n</code></pre></div>"},
	{"35", "<p><em data-marker=\"*\">foo</em></p><p><em data-marker=\"*\"><wbr><br></em></p>", "<p data-block=\"0\"><em data-marker=\"*\">foo</em></p><p data-block=\"0\"><wbr></p>"},
	{"34", "<p> <span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\">foo</code></span> </p>", "<p data-block=\"0\">\u200b<span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\" style=\"display: none\">\u200bfoo</code><span class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span></span>\u200b</p>"},
	{"33", "<p><code>foo</code><wbr></p>", "<p data-block=\"0\">\u200b<code data-marker=\"`\">\u200bfoo</code>\u200b<wbr></p>"},
	{"32", "<p>```<wbr></p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code><wbr>\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code></code></pre></div>"},
	{"31", "<p>\u200bfoo<wbr></p>", "<p data-block=\"0\">foo<wbr></p>"},
	{"30", "<p>1. Node.js</p><p>2. Go<wbr></p>", "<ol data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\"><p data-block=\"0\">Node.js</p></li><li data-marker=\"2.\"><p data-block=\"0\">Go<wbr></p></li></ol>"},
	{"29", "<p><wbr><br></p>", "<p data-block=\"0\"><wbr></p>"},
	{"28", "<p data-block=\"0\">❤️<wbr></p>", "<p data-block=\"0\">❤️<wbr></p>"},
	{"27", "<p><wbr></p>", "<p data-block=\"0\"><wbr></p>"},
	{"26", "<p>![alt](src \"title\")</p>", "<p data-block=\"0\"><img src=\"src\" alt=\"alt\" title=\"title\" /></p>"},
	{"25", "<pre><code class=\"language-java\"><wbr>\n</code></pre>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\" data-block=\"0\" data-marker=\"```\"><pre class=\"vditor-wysiwyg__pre\"><code class=\"language-java\"><wbr>\n</code></pre><pre class=\"vditor-wysiwyg__preview\" data-render=\"2\"><code class=\"language-java\">\n</code></pre></div>"},
	{"24", "<ul data-tight=\"true\"><li data-marker=\"*\"><wbr><br></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\"><wbr></li></ul>"},
	{"23", "<ol><li data-marker=\"1.\">foo</li></ol>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">foo</li></ol>"},
	{"22", "<ul><li data-marker=\"*\">foo</li><li data-marker=\"*\"><ul><li data-marker=\"*\">bar</li></ul></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo</li><li data-marker=\"*\"><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">bar</li></ul></li></ul>"},
	{"21", "<p>[foo](/bar \"baz\")</p>", "<p data-block=\"0\"><a href=\"/bar\" title=\"baz\">foo</a></p>"},
	{"20", "<p>[foo](/bar)</p>", "<p data-block=\"0\"><a href=\"/bar\">foo</a></p>"},
	{"19", "<p>[foo]()</p>", "<p data-block=\"0\"><a href=\"\">foo</a></p>"},
	{"18", "<p>[](/bar)</p>", "<p data-block=\"0\">[](/bar)</p>"},
	{"17", "<p>[]()</p>", "<p data-block=\"0\">[]()</p>"},
	{"16", "<p>[](</p>", "<p data-block=\"0\">[](</p>"},
	{"15", "<p><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" /></p>", "<p data-block=\"0\"><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" /></p>"},
	{"14", ":octocat:", "<p data-block=\"0\"><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" /></p>"},
	{"13", "<p>1、foo</p>", "<ol data-tight=\"true\" data-marker=\"1.\" data-block=\"0\"><li data-marker=\"1.\">foo</li></ol>"},
	{"12", "<p><s data-marker=\"~~\">Hi</s> Hello, world!</p>", "<p data-block=\"0\"><s data-marker=\"~~\">Hi</s> Hello, world!</p>"},
	{"11", "<p><del data-marker=\"~\">Hi</del> Hello, world!</p>", "<p data-block=\"0\"><s data-marker=\"~\">Hi</s> Hello, world!</p>"},
	{"10", "<ul data-tight=\"true\"><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> foo<wbr></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> foo<wbr></li></ul>"},
	{"9", "<ul data-tight=\"true\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> foo<wbr></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> foo<wbr></li></ul>"},
	{"8", "> <wbr>", "<p data-block=\"0\">&gt; <wbr></p>"},
	{"7", "><wbr>", "<p data-block=\"0\">&gt;<wbr></p>"},
	{"6", "<p>> foo<wbr></p>", "<blockquote data-block=\"0\"><p data-block=\"0\">foo<wbr></p></blockquote>"},
	{"5", "<p>foo</p><p><wbr><br></p>", "<p data-block=\"0\">foo</p><p data-block=\"0\"><wbr></p>"},
	{"4", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<wbr></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\"><li data-marker=\"*\">foo<wbr></li></ul>"},
	{"3", "<p><em data-marker=\"*\">foo<wbr></em></p>", "<p data-block=\"0\"><em data-marker=\"*\">foo<wbr></em></p>"},
	{"2", "<p>foo<wbr></p>", "<p data-block=\"0\">foo<wbr></p>"},
	{"1", "<p><strong data-marker=\"**\">foo</strong></p>", "<p data-block=\"0\"><strong data-marker=\"**\">foo</strong></p>"},
	{"0", "<p>foo</p>", "<p data-block=\"0\">foo</p>"},
}

func TestSpinVditorDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.ToC = true
	luteEngine.Sanitize = true
	luteEngine.Mark = true

	for _, test := range spinVditorDOMTests {
		html := luteEngine.SpinVditorDOM(test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}