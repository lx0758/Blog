package com.liux.blog.markdown

import com.liux.blog.markdown.bean.CatalogueItem
import com.liux.blog.markdown.parser.NexTParseExtension
import com.liux.blog.markdown.render.NexTRenderExtension
import org.commonmark.ext.autolink.AutolinkExtension
import org.commonmark.ext.gfm.strikethrough.StrikethroughExtension
import org.commonmark.ext.gfm.tables.TablesExtension
import org.commonmark.ext.heading.anchor.HeadingAnchorExtension
import org.commonmark.ext.heading.anchor.IdGenerator
import org.commonmark.ext.image.attributes.ImageAttributesExtension
import org.commonmark.ext.ins.InsExtension
import org.commonmark.ext.task.list.items.TaskListItemsExtension
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

class MarkdownHelper {
    companion object {
        const val FLAG_SHOW_TOC = 1
        const val FLAG_SHOW_MORE_ANCHOR = 1 shl 1
        const val FLAG_MORE_SUSPEND = 1 shl 2
        const val FLAG_SHOW_INFO = FLAG_SHOW_TOC or FLAG_SHOW_MORE_ANCHOR

        private const val HEADING_ANCHOR_DEFAULT_ID = "id"
        private const val HEADING_ANCHOR_ID_PREFIX = ""
        private const val HEADING_ANCHOR_ID_SUFFIX = ""

        private val EXTENSIONS = arrayListOf(
            InsExtension.create(),
            TablesExtension.create(),
            AutolinkExtension.create(),
            HeadingAnchorExtension.builder()
                .defaultId(HEADING_ANCHOR_DEFAULT_ID)
                .idPrefix(HEADING_ANCHOR_ID_PREFIX)
                .idSuffix(HEADING_ANCHOR_ID_SUFFIX)
                .build(),
            StrikethroughExtension.create(),
            TaskListItemsExtension.create(),
            ImageAttributesExtension.create(),
            NexTRenderExtension.create(),
        )

        private val HTML_RENDERER = HtmlRenderer.builder()
            .extensions(EXTENSIONS)
            .build()

        fun parseHtml(content: String?, flag: Int = 0, catalogueResult: MutableList<CatalogueItem>? = null): String {
            if (content.isNullOrEmpty()) return ""

            val idGenerator = IdGenerator.builder()
                .defaultId(HEADING_ANCHOR_DEFAULT_ID)
                .prefix(HEADING_ANCHOR_ID_PREFIX)
                .suffix(HEADING_ANCHOR_ID_SUFFIX)
                .build()

            val document = Parser.builder()
                .extensions(EXTENSIONS)
                .extensions(arrayListOf(NexTParseExtension(Flag(flag), idGenerator, catalogueResult)))
                .build()
                .parse(content)

            val html = HTML_RENDERER.render(document)

            return html
        }
    }
}