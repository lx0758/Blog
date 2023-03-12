package org.commonmark.ext.next

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
        val contents = fencedCodeBlock.literal.trim().lines()
        val language = fencedCodeBlock.info
        val lineMap = mapOf("class" to "line")

        html.line()
        html.tag("figure", mapOf("class" to "highlight $language"))
        html.tag("table")
        html.tag("tr")

        html.tag("td", mapOf("class" to "gutter"))
        html.tag("pre")
        for (line in 1..contents.size) {
            html.tag("span", lineMap)
            html.text(line.toString())
            html.tag("/span")
            html.tag("br")
        }
        html.tag("/pre")
        html.tag("/td")

        html.tag("td", mapOf("class" to "code"))
        html.tag("pre")
        contents.forEach {
            html.tag("span", lineMap)
            html.text(it)
            html.tag("/span")
            html.tag("br")
        }
        html.tag("/pre")
        html.tag("/td")

        html.tag("/tr")
        html.tag("/table")
        html.tag("/figure")
        html.line()
    }

    class Factory : HtmlNodeRendererFactory {
        override fun create(context: HtmlNodeRendererContext): NodeRenderer {
            return NexTNodeRenderer(context)
        }
    }
}