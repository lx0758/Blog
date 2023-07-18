package com.liux.blog.markdown

import com.liux.blog.bean.vo.CatalogueVO
import org.commonmark.ext.autolink.AutolinkExtension
import org.commonmark.ext.gfm.strikethrough.StrikethroughExtension
import org.commonmark.ext.gfm.tables.TablesExtension
import org.commonmark.ext.heading.anchor.HeadingAnchorExtension
import org.commonmark.ext.image.attributes.ImageAttributesExtension
import org.commonmark.ext.ins.InsExtension
import org.commonmark.ext.task.list.items.TaskListItemsExtension
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

class MarkdownHelper {
    companion object {
        private val EXTENSIONS = arrayListOf(
            InsExtension.create(),
            TablesExtension.create(),
            AutolinkExtension.create(),
            HeadingAnchorExtension.create(),
            StrikethroughExtension.create(),
            TaskListItemsExtension.create(),
            ImageAttributesExtension.create(),
        )
        private val RENDERER = HtmlRenderer.builder()
            .extensions(EXTENSIONS)
            .extensions(arrayListOf(NexTExtension.create()))
            .build()

        fun parse(content: String?, flag: Int = 0, catalogues: MutableList<CatalogueVO>? = null): String {
            if (content.isNullOrEmpty()) return ""
            var document = Parser.builder()
                .extensions(EXTENSIONS)
                .extensions(arrayListOf(NexTExtension.create(flag)))
                .build()
                .parse(content)
            if (document is NexTDocument) {
                if (catalogues != null) {
                    transformCatalogue(document.catalogues, catalogues)
                }
                document = document.original
            }
            return RENDERER.render(document)
        }

        private fun transformCatalogue(catalogues: Collection<Catalogue>, catalogueVOs: MutableList<CatalogueVO>) {
            for (catalogue in catalogues) {
                val childs = catalogue.childs
                var childVOs: MutableList<CatalogueVO>? = null
                if (childs != null) {
                    childVOs = ArrayList()
                    transformCatalogue(childs, childVOs)
                }
                catalogueVOs.add(CatalogueVO(catalogue.title, childVOs))
            }
        }
    }
}
