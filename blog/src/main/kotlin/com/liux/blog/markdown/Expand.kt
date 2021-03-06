package com.liux.blog.markdown

import org.commonmark.node.Code
import org.commonmark.node.Heading
import org.commonmark.node.Text

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