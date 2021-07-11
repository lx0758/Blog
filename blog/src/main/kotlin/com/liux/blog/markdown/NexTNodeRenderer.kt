package com.liux.blog.markdown

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
        return arrayOf(
            MoreBlock::class.java,
            FencedCodeBlock::class.java,
        ).toHashSet()
    }

    override fun render(node: Node) {
        node.accept(this)
    }

    override fun visit(customBlock: CustomBlock) {
        if (customBlock !is MoreBlock) return

        html.tag("a", mapOf("id" to "more", "href" to "#more"))
        html.tag("/a")
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
}