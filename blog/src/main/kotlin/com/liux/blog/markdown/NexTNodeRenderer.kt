package com.liux.blog.markdown

import org.commonmark.node.AbstractVisitor
import org.commonmark.node.CustomBlock
import org.commonmark.node.FencedCodeBlock
import org.commonmark.node.Node
import org.commonmark.renderer.NodeRenderer
import org.commonmark.renderer.html.HtmlNodeRendererContext
import org.commonmark.renderer.html.HtmlNodeRendererFactory

class NexTNodeRenderer(
    htmlNodeRendererContext: HtmlNodeRendererContext,
) : AbstractVisitor(), NodeRenderer {

    private val html = htmlNodeRendererContext.writer

    override fun getNodeTypes(): MutableSet<Class<out Node>> {
        return hashSetOf(
            TocBlock::class.java,
            ShowTocBlock::class.java,

            MoreBlock::class.java,
            SuspendMoreBlock::class.java,
            AnchorMoreBlock::class.java,

            FencedCodeBlock::class.java,
        )
    }

    override fun render(node: Node) {
        node.accept(this)
    }

    override fun visit(customBlock: CustomBlock) {
        when(customBlock) {
            is ShowTocBlock -> {
                rendererCatalogues(customBlock.catalogues)
            }
            is SuspendMoreBlock -> {
                html.raw(CUSTOM_NODE_MORE)
            }
            is AnchorMoreBlock -> {
                html.tag("a", mapOf("id" to "more", "href" to "#more"))
                html.tag("/a")
            }
            else -> {}
        }
    }

    private fun rendererCatalogues(catalogues: Collection<Catalogue>?) {
        if (catalogues == null) return
        html.tag("ol")
        catalogues.forEach {
            html.tag("li")
            html.tag("strong")
            html.tag("a", mapOf("href" to "#${it.title}"))
            html.raw(it.title)
            html.tag("/a")
            html.tag("/strong")
            html.tag("/li")
            rendererCatalogues(it.childs)
        }
        html.tag("/ol")
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

    class Factory : HtmlNodeRendererFactory {
        override fun create(context: HtmlNodeRendererContext): NodeRenderer {
            return NexTNodeRenderer(context)
        }
    }
}