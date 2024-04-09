package com.liux.blog.markdown.render

import org.commonmark.Extension
import org.commonmark.renderer.html.HtmlRenderer

class NexTRenderExtension: HtmlRenderer.HtmlRendererExtension {

    companion object {
        fun create(): Extension {
            return NexTRenderExtension()
        }
    }

    override fun extend(htmlRendererBuilder: HtmlRenderer.Builder) {
        htmlRendererBuilder.attributeProviderFactory {
            NexTAttributeProvider()
        }
        htmlRendererBuilder.nodeRendererFactory { htmlNodeRendererContext ->
            NexTNodeRenderer(htmlNodeRendererContext)
        }
    }
}
