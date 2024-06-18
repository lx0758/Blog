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
        catalogueResult: MutableList<CatalogueItem>?,
    ) {

        private val root: CatalogueItem = CatalogueItem("", "", 0, null, catalogueResult ?: ArrayList())
        private var last: CatalogueItem = root

        fun getCatalogueList(): List<CatalogueItem> = root.childs!!

        /**
         * 处理某些文章跳级使用标题标签, 导致无法正常解析标题列表的错误
         *
         * #     ........ 1
         * ###   ........ 2
         * ##### ........ 3
         * ##    ........ 2
         * ####  ........ 3
         * ###   ........ 3
         * #     ........ 1
         **/
        fun append(heading: Heading) {
            if (heading.level == last.rawLevel) {
                last = addCatalogueItem(last.parent!!, heading)
            } else if (heading.level > last.rawLevel) {
                last = addCatalogueItem(last, heading)
            } else {
                var current = last
                while (heading.level <= current.rawLevel) {
                    current = current.parent ?: break
                }
                last = addCatalogueItem(current, heading)
            }
        }

        private fun addCatalogueItem(parent: CatalogueItem, heading: Heading): CatalogueItem {
            val title = getHeadingTitle(heading)
            val anchor = idGenerator.generateId(title)
            val item = CatalogueItem(title, anchor, heading.level, parent)

            if (parent.childs == null) {
                parent.childs = ArrayList()
            }
            parent.childs!!.add(item)

            return item
        }

        private fun getHeadingTitle(heading: Heading): String {
            val builder = StringBuilder()
            heading.accept(object : AbstractVisitor() {
                override fun visit(text: Text) {
                    builder.append(text.literal)
                }

                override fun visit(code: Code) {
                    builder.append(code.literal)
                }
            })
            return builder.toString()
        }
    }
}
