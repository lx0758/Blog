package com.liux.blog.markdown

import com.liux.blog.markdown.MarkdownHelper.Companion.FLAG_MORE_SUSPEND
import com.liux.blog.markdown.MarkdownHelper.Companion.FLAG_SHOW_INFO
import com.liux.blog.markdown.MarkdownHelper.Companion.FLAG_SHOW_MORE_ANCHOR
import com.liux.blog.markdown.MarkdownHelper.Companion.FLAG_SHOW_TOC

data class Flag(
    private val flag: Int
) {
    fun showInfo() = flag and FLAG_SHOW_INFO == FLAG_SHOW_INFO
    fun showToc() = flag and FLAG_SHOW_TOC == FLAG_SHOW_TOC
    fun showMoreAnchor() = flag and FLAG_SHOW_MORE_ANCHOR == FLAG_SHOW_MORE_ANCHOR

    fun isMoreSuspend() = flag and FLAG_MORE_SUSPEND == FLAG_MORE_SUSPEND
}