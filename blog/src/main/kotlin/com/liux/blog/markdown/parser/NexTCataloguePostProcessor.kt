package com.liux.blog.markdown.parser

import com.liux.blog.markdown.bean.CatalogueItem
import org.commonmark.ext.heading.anchor.IdGenerator
import org.commonmark.node.*
import org.commonmark.parser.PostProcessor

class NexTCataloguePostProcessor(
    private val catalogueStore: CatalogueStore,
) : PostProcessor {
    override fun process(node: Node): Node {
        if (node is Document) {
            traversalDocument(node)
        }
        return node
    }

    private fun traversalDocument(document: Document) {
        var node: Node? = document.firstChild
        while (node != null) {
            handlerNode(node)
            node = node.next
        }
    }

    private fun handlerNode(node: Node) {
        if (node is Heading) {
            catalogueStore.append(node)
        }
    }

    class CatalogueStore(
        private val idGenerator: IdGenerator,
        private val catalogueResult : MutableList<CatalogueItem>?,
    ) {
        private var lastHeadingLevel = 0
        private var lastCatalogueLevel = 0

        private val catalogueList: MutableList<CatalogueItem> = catalogueResult ?: ArrayList()

        fun getCatalogueList() : List<CatalogueItem> = catalogueList

        fun append(heading: Heading) {
            heading.apply {
                // 修正某些文章跳级使用标题, 导致无法正常解析标题的错误
                if (level != lastHeadingLevel) {
                    if (level > lastCatalogueLevel + 1) {
                        lastCatalogueLevel += 1
                    } else {
                        lastCatalogueLevel = level
                    }
                    lastHeadingLevel = level
                }
            }
            val title = heading.let {
                val builder = StringBuilder()
                it.accept(object : AbstractVisitor() {
                    override fun visit(text: Text) {
                        builder.append(text.literal)
                    }
                    override fun visit(code: Code) {
                        builder.append(code.literal)
                    }
                })
                builder.toString()
            }
            val anchor = idGenerator.generateId(title)
            val parent = catalogueList.let {
                var childs = it
                var index = lastCatalogueLevel
                while (index > 1) {
                    val last = childs.last()
                    if (last.childs == null) {
                        last.childs = ArrayList()
                    }
                    childs = last.childs!!
                    index--
                }
                childs
            }

            parent.add(
                CatalogueItem(title, anchor)
            )
        }
    }
}
