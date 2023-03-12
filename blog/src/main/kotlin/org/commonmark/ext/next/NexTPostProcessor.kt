package org.commonmark.ext.next

import org.commonmark.node.Document
import org.commonmark.node.Heading
import org.commonmark.node.Node
import org.commonmark.parser.PostProcessor

class NexTPostProcessor : PostProcessor {

    override fun process(node: Node): Node {
        if (node !is Document) return node
        return NexTDocument(
            node,
            handleCatalogues(node)
        )
    }

    /**
     * 1. 生成目录树
     * 2. 向 Toc 节点配置目录树
     */
    private fun handleCatalogues(document: Document): Collection<Catalogue> {
        val catalogues = ArrayList<Catalogue>()
        var lastCatalogueLevel = 0
        var outCatalogueLevel = 0
        var node: Node? = document.firstChild
        while (node != null) {
            when (node) {
                is Heading -> {
                    val level = node.level
                    val title = node.parseTitle()

                    // 修正某些文章跳级使用标题
                    // 导致无法正常解析标题的错误
                    if (level != lastCatalogueLevel) {
                        if (level > outCatalogueLevel + 1) {
                            outCatalogueLevel += 1
                        } else {
                            outCatalogueLevel = level
                        }
                        lastCatalogueLevel = level
                    }

                    setLastCatalogueByLevel(catalogues, outCatalogueLevel, title)
                }
                is ShowTocBlock -> {
                    node.catalogues = catalogues
                }
            }
            node = node.next
        }

        return catalogues
    }

    private fun setLastCatalogueByLevel(catalogues: MutableList<Catalogue>, level: Int, title: String) {
        var childs: MutableList<Catalogue> = catalogues
        var index = level
        while (index > 1) {
            val last = childs.last()
            if (last.childs == null) {
                last.childs = ArrayList()
            }
            childs = last.childs!!
            index--
        }
        childs.add(Catalogue(title, null))
    }
}
