package com.liux.blog.markdown.render

import org.commonmark.node.Heading
import org.commonmark.node.Link
import org.commonmark.node.Node
import org.commonmark.renderer.html.AttributeProvider

/**
 * 1. 给超链接设置新标签打开属性
 * 2. 获取标题锚点
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
                val anchor = attributes?.get("id") ?: return
                attributes["href"] = "#${anchor}"
            }
        }
    }
}