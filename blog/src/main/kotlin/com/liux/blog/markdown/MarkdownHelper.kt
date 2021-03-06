package com.liux.blog.markdown

import com.liux.blog.bean.vo.CatalogueVO
import org.commonmark.ext.autolink.AutolinkExtension
import org.commonmark.ext.gfm.strikethrough.StrikethroughExtension
import org.commonmark.ext.gfm.tables.TablesExtension
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

class MarkdownHelper {
    companion object {
        private val EXTENSIONS = arrayListOf(
            TablesExtension.create(),
            AutolinkExtension.create(),
            StrikethroughExtension.create(),
            NexTExtension.create(),
        )
        private val PARSER = Parser.builder().extensions(EXTENSIONS).build()
        private val RENDERER = HtmlRenderer.builder().extensions(EXTENSIONS).build()
        fun parse(content: String, catalogues: MutableList<CatalogueVO>? = null): String {
            val nexTDocument = PARSER.parse(content) as NexTDocument

            catalogues?.addAll(nexTDocument.catalogues)

            return RENDERER.render(nexTDocument.original)
        }
    }
}
