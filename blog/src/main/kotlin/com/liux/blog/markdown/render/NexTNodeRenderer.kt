package com.liux.blog.markdown.render

import com.liux.blog.markdown.bean.CatalogueItem
import com.liux.blog.markdown.node.*
import org.commonmark.node.AbstractVisitor
import org.commonmark.node.CustomBlock
import org.commonmark.node.FencedCodeBlock
import org.commonmark.node.Node
import org.commonmark.renderer.NodeRenderer
import org.commonmark.renderer.html.HtmlNodeRendererContext

class NexTNodeRenderer(
    htmlNodeRendererContext: HtmlNodeRendererContext
) : AbstractVisitor(), NodeRenderer {

    private val html = htmlNodeRendererContext.writer

    override fun getNodeTypes(): MutableSet<Class<out Node>> {
        return hashSetOf(
            Catalogue::class.java,
            NoneCatalogue::class.java,
            MoreSuspend::class.java,
            MoreAnchor::class.java,
            FencedCodeBlock::class.java,
        )
    }

    override fun render(node: Node) {
        node.accept(this)
    }

    override fun visit(customBlock: CustomBlock) {
        when(customBlock) {
            is Catalogue -> {
                rendererCatalogue(customBlock.catalogues)
            }
            is NoneCatalogue -> {}
            is MoreSuspend -> {
                html.raw("<!--more-->")
            }
            is MoreAnchor -> {
                html.tag("a", mapOf("id" to "more", "href" to "#more"))
                html.tag("/a")
            }
            is NoneMore -> {}
            else -> {}
        }
    }

    override fun visit(fencedCodeBlock: FencedCodeBlock) {
        val language = fencedCodeBlock.info
        html.line()
        html.tag("pre", mapOf("class" to "line-numbers"))
        html.tag("code", mapOf("class" to "language-${language.ifEmpty { "none" }}"))
        html.text(fencedCodeBlock.literal)
        html.tag("/code")
        html.tag("/pre")
        html.line()
    }

    private fun rendererCatalogue(catalogues: List<CatalogueItem>?) {
        if (catalogues == null) return
        html.tag("ol")
        catalogues.forEach { catalogue ->
            html.tag("li")
            html.tag("strong")
            html.tag("a", mapOf("href" to "#${catalogue.anchor}"))
            html.raw(catalogue.title)
            html.tag("/a")
            html.tag("/strong")
            html.tag("/li")
            rendererCatalogue(catalogue.childs)
        }
        html.tag("/ol")
    }
}