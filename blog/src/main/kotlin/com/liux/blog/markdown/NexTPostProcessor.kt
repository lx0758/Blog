package com.liux.blog.markdown

import com.liux.blog.bean.vo.CatalogueVO
import org.commonmark.node.Document
import org.commonmark.node.Heading
import org.commonmark.node.HtmlBlock
import org.commonmark.node.Node
import org.commonmark.parser.PostProcessor

class NexTPostProcessor : PostProcessor {

    override fun process(node: Node): Node {
        if (node !is Document) return node

        val catalogues = parseCatalogues(node)

        processMore(node)

        return NexTDocument(
            node,
            catalogues
        )
    }

    private fun processMore(document: Document) {
        var node: Node? = document.firstChild
        while (node != null) {
            if (node is HtmlBlock && "<!-- more -->" == node.literal) {
                node.insertAfter(MoreBlock())
                break
            }
            node = node.next
        }
    }

    private fun parseCatalogues(document: Document): List<CatalogueVO> {
        val catalogues = ArrayList<CatalogueVO>()
        var lastLevel = 0
        var outLevel = 0
        var node: Node? = document.firstChild
        while (node != null) {
            if (node is Heading) {
                val level = node.level
                val title = node.parseTitle()

                // 修正某些文章跳级使用标题
                // 导致无法正常解析标题的错误
                if (level != lastLevel) {
                    if (level > outLevel + 1) {
                        outLevel += 1
                    } else outLevel = level
                    lastLevel = level
                }

                setLastCatalogueByLevel(catalogues, outLevel, title)
            }
            node = node.next
        }
        return catalogues
    }

    private fun setLastCatalogueByLevel(catalogues: MutableList<CatalogueVO>, level: Int, title: String) {
        var childs: MutableList<CatalogueVO> = catalogues
        var index = level
        while (index > 1) {
            val last = childs.last()
            if (last.childs == null) {
                last.childs = ArrayList()
            }
            childs = last.childs!!
            index--
        }
        childs.add(CatalogueVO(title, null))
    }
}
