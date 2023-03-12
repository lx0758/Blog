package org.commonmark.ext.next

import org.commonmark.node.*

const val CUSTOM_NODE_TOC = "[toc]"
const val CUSTOM_NODE_MORE = "<!--more-->"

const val FLAG_TOC = 1
const val FLAG_MORE_SUSPEND = 1 shl 1
const val FLAG_MORE_ANCHOR = 1 shl 2

fun Heading.parseTitle(): String {
    val springBuilder = StringBuilder()
    var node = firstChild
    while (node != null) {
        if (springBuilder.isNotEmpty()) springBuilder.append(' ')
        when (node) {
            is Text -> springBuilder.append(node.literal)
            is Code -> springBuilder.append(node.literal)
        }
        node = node.next
    }
    return springBuilder.toString()
}

data class Catalogue(
    var title: String,
    var childs: MutableList<Catalogue>? = null
)

data class Flag(
    private val flag: Int
) {
    fun isShowToc(): Boolean = flag and FLAG_TOC == FLAG_TOC

    fun isMoreSuspend(): Boolean = flag and FLAG_MORE_SUSPEND == FLAG_MORE_SUSPEND
    
    fun isMoreAnchor(): Boolean = flag and FLAG_MORE_ANCHOR == FLAG_MORE_ANCHOR
}

class NexTDocument(
    val original: Document,
    val catalogues: Collection<Catalogue>
) : Document()

open class TocBlock : CustomBlock()

class ShowTocBlock : TocBlock() {
    var catalogues: Collection<Catalogue>? = null
}

open class MoreBlock : CustomBlock()

class SuspendMoreBlock : MoreBlock()

class AnchorMoreBlock : MoreBlock()
