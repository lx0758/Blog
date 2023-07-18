package com.liux.blog.markdown

import org.commonmark.node.Heading
import org.commonmark.node.Link
import org.commonmark.node.Node
import org.commonmark.renderer.html.AttributeProvider
import org.commonmark.renderer.html.AttributeProviderContext
import org.commonmark.renderer.html.AttributeProviderFactory

/**
 * 1. 给超链接设置新标签打开属性
 * 2. 给标题添加锚点
 */
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

    class Factory : AttributeProviderFactory {
        override fun create(context: AttributeProviderContext): AttributeProvider {
            return NexTAttributeProvider()
        }
    }
}