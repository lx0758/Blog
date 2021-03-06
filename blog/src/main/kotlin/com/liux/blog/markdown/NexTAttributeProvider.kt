package com.liux.blog.markdown

import org.commonmark.node.Heading
import org.commonmark.node.Link
import org.commonmark.node.Node
import org.commonmark.renderer.html.AttributeProvider

class NexTAttributeProvider : AttributeProvider {
    override fun setAttributes(node: Node?, tagName: String?, attributes: MutableMap<String, String>?) {
        when (node) {
            is Link -> {
                val url = attributes?.get("href") ?: return
                if (!url.startsWith("/")) {
                    attributes["target"] = "_blank"
                }
            }
            is Heading -> {
                val title = node.parseTitle()
                attributes?.set("id", title);
                attributes?.set("href", "#$title");
            }
        }
    }
}