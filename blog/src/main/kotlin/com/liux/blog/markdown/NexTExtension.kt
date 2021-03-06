package com.liux.blog.markdown

import org.commonmark.Extension
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

/**
 * https://blog.csdn.net/qq_41609208/article/details/106824386
 */
class NexTExtension : Parser.ParserExtension, HtmlRenderer.HtmlRendererExtension {

    companion object {
        fun create(): Extension {
            return NexTExtension()
        }
    }

    override fun extend(parserBuilder: Parser.Builder) {
        parserBuilder.postProcessor(NexTPostProcessor())
    }

    override fun extend(htmlRendererBuilder: HtmlRenderer.Builder) {
        htmlRendererBuilder.nodeRendererFactory { NexTNodeRenderer(it) }
        htmlRendererBuilder.attributeProviderFactory { NexTAttributeProvider() }
    }
}
